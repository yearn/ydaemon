package internal

import (
	"sync"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/yearn/ydaemon/common/logs"
	"github.com/yearn/ydaemon/internal/events"
	"github.com/yearn/ydaemon/internal/fetcher"
	"github.com/yearn/ydaemon/internal/indexer"
	"github.com/yearn/ydaemon/internal/models"
	"github.com/yearn/ydaemon/internal/risk"
	"github.com/yearn/ydaemon/internal/storage"
	"github.com/yearn/ydaemon/processes/apr"
	"github.com/yearn/ydaemon/processes/initDailyBlock"
	"github.com/yearn/ydaemon/processes/prices"
)

var STRATLIST = []models.TStrategy{}

func InitializeV2(chainID uint64, wg *sync.WaitGroup, scheduler *gocron.Scheduler) {
	if wg != nil {
		defer wg.Done()
	}
	go InitializeBribes(chainID)

	underWg := sync.WaitGroup{}
	underWg.Add(2)
	/** 🔵 - Yearn *************************************************************************************
	** InitializeV2 is only called on initialization. It's first job is to retrieve the initial data:
	** - The registries vaults
	** - The vaults
	** - The strategies
	** - The tokens
	**************************************************************************************************/
	registries := indexer.IndexNewVaults(chainID)
	vaultMap := fetcher.RetrieveAllVaults(chainID, registries)

	/**********************************************************************************************
	** Start the OP Staking Indexing process and schedule it to run every hour. This indexer
	** will fetch the relevent data for the OP Staking contracts as well as thoose responsible
	** for the APR calculations.
	**********************************************************************************************/
	indexer.IndexStakingPools(chainID)
	scheduler.Every(1).Hours().StartAt(time.Now().Add(time.Minute * 10)).Do(func() {
		indexer.IndexStakingPools(chainID)
	})

	/**********************************************************************************************
	** Start the VEYFI Staking Indexing process and schedule it to run every hour. This indexer
	** will fetch the relevent data for the VEYFI Staking contracts as well as thoose responsible
	** for the APR calculations.
	**********************************************************************************************/
	indexer.IndexVeYFIStakingContract(chainID)
	scheduler.Every(1).Hours().StartAt(time.Now().Add(time.Minute * 10)).Do(func() {
		indexer.IndexVeYFIStakingContract(chainID)
	})

	/**********************************************************************************************
	** Start the Juiced Staking Indexing process and schedule it to run every hour. This indexer
	** will fetch the relevent data for the Juiced Staking contracts as well as thoose responsible
	** for the APR calculations.
	**********************************************************************************************/
	indexer.IndexJuicedStakingContract(chainID)
	scheduler.Every(1).Hours().StartAt(time.Now().Add(time.Minute * 10)).Do(func() {
		indexer.IndexJuicedStakingContract(chainID)
	})

	/**********************************************************************************************
	** Start the V3 Staking Indexing process and schedule it to run every hour. This indexer
	** will fetch the relevent data for the V3 Staking contracts as well as thoose responsible
	** for the APR calculations.
	**********************************************************************************************/
	indexer.IndexV3StakingContract(chainID)
	scheduler.Every(1).Hours().StartAt(time.Now().Add(time.Minute * 10)).Do(func() {
		indexer.IndexV3StakingContract(chainID)
	})
	logs.Success(chainID, `-`, `Index StakingPools ✅`)

	/**********************************************************************************************
	** Continue indexing with Strategies and APR
	**********************************************************************************************/
	strategiesMap := indexer.IndexNewStrategies(chainID, vaultMap)
	logs.Success(chainID, `-`, `Index New Strategies ✅`)
	tokenMap := fetcher.RetrieveAllTokens(chainID, vaultMap)
	logs.Success(chainID, `-`, `Retrieve All Tokens ✅`)

	go func() {
		prices.RetrieveAllPrices(chainID, tokenMap)
		logs.Success(chainID, `-`, `Retrieve All Prices ✅`)
		underWg.Done()
	}() // Retrieve the prices for all tokens

	go func() {
		fetcher.RetrieveAllStrategies(chainID, strategiesMap)
		logs.Success(chainID, `-`, `Retrieve All Strategies ✅`)
		underWg.Done()
	}() // Retrieve the strategies for all chains

	underWg.Wait()
	apr.ComputeChainAPR(chainID)
	go risk.InitRiskScore(chainID)

	scheduler.Every(10).Hours().StartImmediately().At("12:10").Do(func() {
		initDailyBlock.Run(chainID)
	})

	scheduler.Every(15).Minute().WaitForSchedule().Do(func() {
		indexer.IndexVeYFIStakingContract(chainID)
		vaultMap := fetcher.RetrieveAllVaults(chainID, registries)
		fetcher.RetrieveAllTokens(chainID, vaultMap)

		currentTokenMap, _ := storage.ListERC20(chainID)
		prices.RetrieveAllPrices(chainID, currentTokenMap)
		apr.ComputeChainAPR(chainID)
		go risk.InitRiskScore(chainID)

	})
	scheduler.StartAsync()
}

func InitializeBribes(chainID uint64) {
	allRewardsAdded := events.HandleRewardsAdded(chainID, 0, nil)
	for _, reward := range allRewardsAdded {
		indexer.SetInRewardAddedMap(chainID, reward)
	}
}
