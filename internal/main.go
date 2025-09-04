package internal

import (
	"sync"
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

func initStrategies(chainID uint64, vaultMap map[common.Address]models.TVault) {
	/** 🔵 - Yearn *************************************************************************************
	** initStrategies is only called on initialization. It's first job is to retrieve the strategies
	** from Kong (complete replacement for contract querying), then to schedule retrieval
	**************************************************************************************************/
	strategiesMap := indexer.IndexNewStrategies(chainID, vaultMap)
	fetcher.RetrieveAllStrategies(chainID, strategiesMap)
	logs.Success(chainID, `-`, `InitStrategies (Kong) ✅`, len(strategiesMap))
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
				storage.RefreshVaultMetadata(chainID)
				storage.RefreshStrategyMetadata(chainID)
				storage.RefreshTokenMetadata(chainID)
				storage.RefreshKongData(chainID)
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
				_, _, vaultMap, tokenMap = initVaults(chainID)

				risks.RetrieveAvailableRiskScores(chainID)
				risks.RetrieveAllRiskScores(chainID, vaultMap)

				initStakingPools(chainID)
				initStrategies(chainID, vaultMap)
				/**********************************************************************************************
				** Retrieving prices and strategies for all the given token and strategies on that chain.
				** This is done in parallel to speed up the process and reduce the time it takes to complete.
				** The scheduler is used to retrieve the strategies every 15 minutes.
				** Computing APRS
				**********************************************************************************************/
				prices.RetrieveAllPrices(chainID, tokenMap)
				logs.Success(chainID, `-`, `RetrieveAllPrices ✅`)
				apr.ComputeChainAPY(chainID)
				logs.Success(chainID, `-`, `ComputeChainAPY ✅`)
			},
		),
		gocron.WithStartAt(gocron.WithStartImmediately()),
	)
	scheduler.Start()

	// Load persisted APY data on initialization
	apr.LoadPersistedAPY(chainID)
}
