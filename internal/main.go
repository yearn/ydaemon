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
var cron = gocron.NewScheduler(time.UTC)

func InitializeV2(chainID uint64, wg *sync.WaitGroup) {
	defer wg.Done()
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
	logs.Success(`We have all the staking pool for chainID`, chainID)
	strategiesMap := indexer.IndexNewStrategies(chainID, vaultMap)
	logs.Success(`We got all the strategies for chainID`, chainID)
	tokenMap := fetcher.RetrieveAllTokens(chainID, vaultMap)

	go func() {
		prices.RetrieveAllPrices(chainID, tokenMap)
		underWg.Done()
	}() // Retrieve the prices for all tokens

	go func() {
		fetcher.RetrieveAllStrategies(chainID, strategiesMap)
		underWg.Done()
	}() // Retrieve the strategies for all chains

	underWg.Wait()
	apr.ComputeChainAPR(chainID)
	go risk.InitRiskScore(chainID)

	cron.Every(10).Hours().StartImmediately().At("12:10").Do(func() {
		initDailyBlock.Run(chainID)
	})

	cron.Every(15).Minute().WaitForSchedule().Do(func() {
		vaultMap := fetcher.RetrieveAllVaults(chainID, registries)
		fetcher.RetrieveAllTokens(chainID, vaultMap)

		currentTokenMap, _ := storage.ListERC20(chainID)
		prices.RetrieveAllPrices(chainID, currentTokenMap)
		apr.ComputeChainAPR(chainID)
		go risk.InitRiskScore(chainID)

	})
	cron.StartAsync()
}

func InitializeBribes(chainID uint64) {
	allRewardsAdded := events.HandleRewardsAdded(chainID, 0, nil)
	for _, reward := range allRewardsAdded {
		indexer.SetInRewardAddedMap(chainID, reward)
	}
}
