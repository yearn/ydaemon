package internal

import (
	"sync"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/yearn/ydaemon/internal/events"
	"github.com/yearn/ydaemon/internal/indexer"
	"github.com/yearn/ydaemon/internal/models"
	"github.com/yearn/ydaemon/internal/strategies"
	"github.com/yearn/ydaemon/internal/tokens"
	"github.com/yearn/ydaemon/internal/vaults"
	"github.com/yearn/ydaemon/processes/apy"
	"github.com/yearn/ydaemon/processes/initDailyBlock"
	"github.com/yearn/ydaemon/processes/prices"
)

var STRATLIST = []models.TStrategy{}
var cron = gocron.NewScheduler(time.UTC)

func InitializeV2(chainID uint64, wg *sync.WaitGroup) {
	defer wg.Done()
	go InitializeBribes(chainID)

	registries := indexer.IndexNewVaults(chainID)
	vaultMap := vaults.RetrieveAllVaults(chainID, registries)
	strategiesSlice := indexer.IndexNewStrategies(chainID, vaultMap)
	tokens.RetrieveAllTokens(chainID, registries)

	cron.Every(10).Hours().StartImmediately().At("12:10").Do(func() {
		initDailyBlock.Run(chainID)
	})

	cron.Every(15).Minutes().StartImmediately().Do(func() {
		prices.Run(chainID)
		events.HandleStakingPoolAdded(chainID, 0, nil)
	})

	cron.Every(15).Minute().Do(func() {
		vaultMap := vaults.RetrieveAllVaults(chainID, registries)
		indexer.IndexNewStrategies(chainID, vaultMap)
		strategies.RetrieveAllStrategies(chainID, strategiesSlice) //TODO: update here

		indexer.PostProcess(chainID)

		apy.ComputeChainAPR(chainID)
		go strategies.InitRiskScore(chainID)
	})

	cron.StartAsync()
}

func InitializeBribes(chainID uint64) {
	allRewardsAdded := events.HandleRewardsAdded(chainID, 0, nil)
	for _, reward := range allRewardsAdded {
		indexer.SetInRewardAddedMap(chainID, reward)
	}
}
