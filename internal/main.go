package internal

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/go-co-op/gocron/v2"
	"github.com/yearn/ydaemon/common/logs"
	"github.com/yearn/ydaemon/internal/fetcher"
	"github.com/yearn/ydaemon/internal/indexer"
	"github.com/yearn/ydaemon/internal/models"
	"github.com/yearn/ydaemon/internal/storage"
	"github.com/yearn/ydaemon/processes/apr"
	"github.com/yearn/ydaemon/processes/prices"
	"github.com/yearn/ydaemon/processes/risks"
)

var STRATLIST = []models.TStrategy{}

// ----- Scheduler visibility helpers -----
var jobSeq uint64
var jobInProgress sync.Map // key: fmt.Sprintf("%d:%s", chainID, jobName) -> time.Time

func beginJob(chainID uint64, name string) (id uint64, started time.Time, overlapped bool) {
	id = atomic.AddUint64(&jobSeq, 1)
	key := fmt.Sprintf("%d:%s", chainID, name)
	if prev, ok := jobInProgress.Load(key); ok {
		overlapped = true
		if t, ok2 := prev.(time.Time); ok2 {
			age := time.Since(t)
			logs.Warning(fmt.Sprintf("⛔️ [OVERLAP DETECTED] job=%s chain=%d prev_started_at=%s prev_age=%s", name, chainID, t.UTC().Format(time.RFC3339), age))
		} else {
			logs.Warning(fmt.Sprintf("⛔️ [OVERLAP DETECTED] job=%s chain=%d", name, chainID))
		}
	}
	started = time.Now()
	jobInProgress.Store(key, started)
	logs.Warning(fmt.Sprintf("🚀 [JOB START] job=%s chain=%d jobID=%d", name, chainID, id))
	return
}

func endJob(chainID uint64, name string, id uint64, started time.Time) {
	key := fmt.Sprintf("%d:%s", chainID, name)
	jobInProgress.Delete(key)
	took := time.Since(started)
	logs.Success(fmt.Sprintf("✅ [JOB DONE] job=%s chain=%d jobID=%d took=%s", name, chainID, id, took))
}

func initStakingPools(chainID uint64) {
	/**********************************************************************************************
	** Start the Staking Indexing process and schedule it to run every hour. This indexer
	** will fetch the relevent data for the OP Staking contracts as well as thoose responsible
	** for the APR calculations.
	**********************************************************************************************/
	indexer.IndexStakingPools(chainID)
	indexer.IndexVeYFIStakingContract(chainID)
	indexer.IndexJuicedStakingContract(chainID)
	indexer.IndexV3StakingContract(chainID)
	logs.Success(chainID, `-`, `InitStakingPools ✅`)
}

func initVaults(chainID uint64) (
	map[common.Address]models.TVaultsFromRegistry,
	map[string]models.TStrategy,
	map[common.Address]models.TVault,
	map[common.Address]models.TERC20Token,
) {
	/** 🔵 - Yearn *************************************************************************************
	** InitializeV2 is only called on initialization. It's first job is to retrieve the initial data:
	** - The vaults (from Kong, complete replacement for registry discovery)
	** - The strategies (from Kong, complete replacement for contract querying)
	** - The tokens
	**************************************************************************************************/
	indexer.IndexYearnXPoolTogetherVaults(chainID)
	indexer.IndexYearnXCoveVaults(chainID)
	// Use Kong as complete replacement for registry discovery
	registries := indexer.IndexNewVaults(chainID)
	logs.Success(chainID, `-`, `InitVaults (Kong) ✅`, len(registries))
	vaultMap, strategiesMap := indexer.ProcessNewVault(chainID, registries, fetcher.ProcessNewVaultMethodReplace)
	logs.Success(chainID, `-`, `InitVaults ✅`, len(vaultMap))
	tokenMap := fetcher.RetrieveAllTokens(chainID, vaultMap)
	logs.Success(chainID, `-`, `InitTokens ✅`, len(tokenMap))
	return registries, strategiesMap, vaultMap, tokenMap
}

func InitializeV2(chainID uint64, wg *sync.WaitGroup) {
	if wg != nil {
		defer wg.Done()
	}

	var vaultMap map[common.Address]models.TVault
	var tokenMap map[common.Address]models.TERC20Token

	scheduler, err := gocron.NewScheduler()
	if err != nil {
		logs.Error(chainID, `-`, `Failed to create scheduler: %v`, err)
		return
	}

	// Schedule metadata refresh every 5 minutes
	scheduler.NewJob(
		gocron.DurationJob(
			time.Minute*5,
		),
		gocron.NewTask(
			func() {
				id, started, _ := beginJob(chainID, "META5M")
				defer endJob(chainID, "META5M", id, started)

				logs.Warning(fmt.Sprintf("🧱 [META] Refresh start chain=%d", chainID))
				t0 := time.Now()
				storage.RefreshVaultMetadata(chainID)
				logs.Info(fmt.Sprintf("🧱 [META] vaults done chain=%d took=%s", chainID, time.Since(t0)))
				t1 := time.Now()
				storage.RefreshStrategyMetadata(chainID)
				logs.Info(fmt.Sprintf("🧱 [META] strategies done chain=%d took=%s", chainID, time.Since(t1)))
				t2 := time.Now()
				storage.RefreshTokenMetadata(chainID)
				logs.Info(fmt.Sprintf("🧱 [META] tokens done chain=%d took=%s", chainID, time.Since(t2)))
				logs.Success(fmt.Sprintf("🧱 [META] Refresh done chain=%d", chainID))
			},
		),
		gocron.WithStartAt(gocron.WithStartImmediately()),
	)

	// Schedule snapshot refresh every 30 minutes
	scheduler.NewJob(
		gocron.DurationJob(
			time.Minute*30,
		),
		gocron.NewTask(
			func() {
				id, started, _ := beginJob(chainID, "SNAPSHOT30M")
				defer endJob(chainID, "SNAPSHOT30M", id, started)

				logs.Warning(fmt.Sprintf("🧩 [SNAPSHOT] initVaults start chain=%d", chainID))
				_, _, vaultMap, tokenMap = initVaults(chainID)
				logs.Success(fmt.Sprintf("🧩 [SNAPSHOT] initVaults done chain=%d vaults=%d tokens=%d", chainID, len(vaultMap), len(tokenMap)))

				tRiskAvail := time.Now()
				risks.RetrieveAvailableRiskScores(chainID)
				logs.Info(fmt.Sprintf("🧩 [SNAPSHOT] risks availability chain=%d took=%s", chainID, time.Since(tRiskAvail)))
				tRiskAll := time.Now()
				risks.RetrieveAllRiskScores(chainID, vaultMap)
				logs.Info(fmt.Sprintf("🧩 [SNAPSHOT] risks full chain=%d took=%s", chainID, time.Since(tRiskAll)))

				tStake := time.Now()
				initStakingPools(chainID)
				logs.Info(fmt.Sprintf("🧩 [SNAPSHOT] staking init chain=%d took=%s", chainID, time.Since(tStake)))
				/**********************************************************************************************
				** Retrieving prices and strategies for all the given token and strategies on that chain.
				** This is done in parallel to speed up the process and reduce the time it takes to complete.
				** The scheduler is used to retrieve the strategies every 15 minutes.
				** Computing APRS
				**********************************************************************************************/
				logs.Warning(fmt.Sprintf("💰 [PRICES] start chain=%d tokens=%d", chainID, len(tokenMap)))
				prices.RetrieveAllPrices(chainID, tokenMap)
				logs.Success(fmt.Sprintf("💰 [PRICES] done chain=%d", chainID))

				logs.Warning(fmt.Sprintf("📈 [APY] start chain=%d vaults=%d", chainID, len(vaultMap)))
				apr.ComputeChainAPY(chainID)
				logs.Success(fmt.Sprintf("📈 [APY] done chain=%d", chainID))
			},
		),
		gocron.WithStartAt(gocron.WithStartImmediately()),
	)
	scheduler.Start()

	// Load persisted APY data on initialization
	apr.LoadPersistedAPY(chainID)
}
