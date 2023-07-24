package internal

import (
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/go-co-op/gocron"
	"github.com/yearn/ydaemon/common/logs"
	"github.com/yearn/ydaemon/internal/events"
	"github.com/yearn/ydaemon/internal/indexer"
	bribes "github.com/yearn/ydaemon/internal/indexer.bribes"
	"github.com/yearn/ydaemon/internal/models"
	"github.com/yearn/ydaemon/internal/prices"
	"github.com/yearn/ydaemon/internal/registries"
	"github.com/yearn/ydaemon/internal/strategies"
	"github.com/yearn/ydaemon/internal/tokens"
	"github.com/yearn/ydaemon/internal/vaults"
	"github.com/yearn/ydaemon/processes/apy"
	"github.com/yearn/ydaemon/processes/initDailyBlock"
)

var STRATLIST = []models.TStrategy{}
var cron = gocron.NewScheduler(time.UTC)

func runRetrieveAllPrices(chainID uint64, wg *sync.WaitGroup, delay time.Duration) {
	isDone := false
	for {
		prices.RetrieveAllPrices(chainID)
		if !isDone {
			isDone = true
			wg.Done()
		}
		if delay == 0 {
			return
		}
		time.Sleep(delay)
	}
}
func runRetrieveAllVaults(chainID uint64, vaultsMap map[common.Address]models.TVaultsFromRegistry, wg *sync.WaitGroup, delay time.Duration) {
	isDone := false
	for {
		vaults.RetrieveAllVaults(chainID, vaultsMap)
		if !isDone {
			isDone = true
			wg.Done()
		}
		if delay == 0 {
			return
		}
		time.Sleep(delay)
	}
}
func runRetrieveAllStrategies(chainID uint64, strategiesAddedList []models.TStrategyAdded, wg *sync.WaitGroup, delay time.Duration) {
	isDone := false
	for {
		strategies.RetrieveAllStrategies(chainID, strategiesAddedList)
		indexer.PostProcessStrategies(chainID)
		if !isDone {
			isDone = true
			wg.Done()
		}
		if delay == 0 {
			return
		}
		time.Sleep(delay)
	}
}

func InitializeV2(chainID uint64, wg *sync.WaitGroup) {
	defer wg.Done()
	now := time.Now()
	logs.Info(`STARTING INITIALIZE V2`, chainID, now)
	go InitializeBribes(chainID)

	vaultsMap := registries.RegisterAllVaults(chainID, 0, nil)
	tokens.RetrieveAllTokens(chainID, vaultsMap)

	cron.Every(30).Minute().Do(func() {
		logs.Warning(`RUNNING RETRIEVE ALL PRICES`, chainID)
		prices.RetrieveAllPrices(chainID)
		logs.Success(`RETRIEVE ALL PRICES COMPLETE`, chainID)

		logs.Warning(`RUNNING RETRIEVE ALL VAULTS`, chainID)
		vaults.RetrieveAllVaults(chainID, vaultsMap)
		logs.Success(`RETRIEVE ALL VAULTS COMPLETE`, chainID)

		logs.Warning(`RUNNING HANDLE STRATEGY ADDED`, chainID)
		strategiesAddedList := events.HandleStrategyAdded(chainID, vaultsMap, 0, nil)
		logs.Success(`HANDLE STRATEGY ADDED COMPLETE`, chainID)

		logs.Warning(`RUNNING RETRIEVE ALL STRATEGIES`, chainID)
		strategies.RetrieveAllStrategies(chainID, strategiesAddedList)
		logs.Success(`RETRIEVE ALL STRATEGIES COMPLETE`, chainID)

		indexer.PostProcessStrategies(chainID)
		go func() {
			initDailyBlock.Run(chainID)
			apy.ComputeChainAPR(chainID)
		}()
		logs.Success(`INITIALIZATION COMPLETE`, chainID, time.Since(now))
	})

	cron.Every(10).Minute().Do(func() {
		strategies.InitRiskScore(chainID)
	})

	cron.Every(1).Day().At("12:10").Do(func() {
		initDailyBlock.Run(chainID)
	})
	cron.StartAsync()
	go registries.IndexNewVaults(chainID)
	logs.Success(`ALL COMPLETE`, chainID, time.Since(now))

}

func InitializeBribes(chainID uint64) {
	allRewardsAdded := events.HandleRewardsAdded(chainID, 0, nil)
	for _, reward := range allRewardsAdded {
		bribes.SetInRewardAddedMap(chainID, reward)
	}
}
