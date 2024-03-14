package internal

import (
	"sync"

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
	indexer.IndexStakingPools(chainID)
	logs.Success(chainID, `-`, `Index StakingPools ✅`)
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

	scheduler.Every(1).Minute().StartImmediately().At("12:10").Do(func() {
		logs.Info(chainID, `-`, `Running initDailyBlock Scheduled Task`)
		initDailyBlock.Run(chainID)
	})

	scheduler.Every(5).Minute().WaitForSchedule().Do(func() {
		logs.Info(chainID, `-`, `Running Primary Scheduled Task`)
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
