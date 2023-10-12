package internal

import (
	"sync"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/yearn/ydaemon/internal/events"
	"github.com/yearn/ydaemon/internal/fetcher"
	"github.com/yearn/ydaemon/internal/indexer"
	"github.com/yearn/ydaemon/internal/models"
	"github.com/yearn/ydaemon/internal/risk"
	"github.com/yearn/ydaemon/processes/apy"
	"github.com/yearn/ydaemon/processes/initDailyBlock"
	"github.com/yearn/ydaemon/processes/prices"
)

var STRATLIST = []models.TStrategy{}
var cron = gocron.NewScheduler(time.UTC)

func InitializeV2(chainID uint64, wg *sync.WaitGroup) {
	defer wg.Done()
	go InitializeBribes(chainID)

	/** 🔵 - Yearn *************************************************************************************
	** InitializeV2 is only called on initialization. It's first job is to retrieve the initial data:
	** - The registries vaults
	** - The vaults
	** - The strategies
	** - The tokens
	**************************************************************************************************/
	registries := indexer.IndexNewVaults(chainID)
	vaultMap := fetcher.RetrieveAllVaults(chainID, registries)
	strategiesMap := indexer.IndexNewStrategies(chainID, vaultMap)
	tokenMap := fetcher.RetrieveAllTokens(chainID, vaultMap)

	prices.RetrieveAllPrices(chainID, tokenMap)           // Retrieve the prices for all tokens
	fetcher.RetrieveAllStrategies(chainID, strategiesMap) // Retrieve the strategies for all chains

	cron.Every(10).Hours().StartImmediately().At("12:10").Do(func() {
		initDailyBlock.Run(chainID)
	})

	cron.Every(15).Minutes().StartImmediately().Do(func() {
		prices.RetrieveAllPrices(chainID, tokenMap)
		events.HandleStakingPoolAdded(chainID, 0, nil)
	})

	cron.Every(15).Minute().Do(func() {
		vaultMap := fetcher.RetrieveAllVaults(chainID, registries)
		indexer.IndexNewStrategies(chainID, vaultMap)

		apy.ComputeChainAPR(chainID)
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
