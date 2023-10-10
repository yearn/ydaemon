package internal

import (
	"sync"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/yearn/ydaemon/internal/events"
	"github.com/yearn/ydaemon/internal/indexer"
	bribes "github.com/yearn/ydaemon/internal/indexer.bribes"
	"github.com/yearn/ydaemon/internal/models"
	"github.com/yearn/ydaemon/internal/registries"
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

	historicalVaults := registries.IndexNewVaults(chainID)
	tokens.RetrieveAllTokens(chainID, historicalVaults)

	cron.Every(10).Hours().StartImmediately().At("12:10").Do(func() {
		initDailyBlock.Run(chainID)
	})

	cron.Every(15).Minute().StartImmediately().Do(func() {
		vaults.RetrieveAllVaults(chainID, historicalVaults)
		prices.Run(chainID)
		strategiesAddedList := events.HandleStrategyAdded(chainID, historicalVaults, 0, nil)
		events.HandleStakingPoolAdded(chainID, 0, nil)
		strategies.RetrieveAllStrategies(chainID, strategiesAddedList)
		indexer.PostProcessStrategies(chainID)
		apy.ComputeChainAPR(chainID)
		go strategies.InitRiskScore(chainID)
	})

	cron.StartAsync()
}

func InitializeBribes(chainID uint64) {
	allRewardsAdded := events.HandleRewardsAdded(chainID, 0, nil)
	for _, reward := range allRewardsAdded {
		bribes.SetInRewardAddedMap(chainID, reward)
	}
}
