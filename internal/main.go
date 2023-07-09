package internal

import (
	"os"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/go-co-op/gocron"
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
	"github.com/yearn/ydaemon/processes/dailyBlockNumber"
	"github.com/yearn/ydaemon/processes/vaultPricePerShare"
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
	go InitializeBribes(chainID)

	vaultsMap := registries.RegisterAllVaults(chainID, 0, nil)
	tokens.RetrieveAllTokens(chainID, vaultsMap)

	cron.Every(15).Minute().Do(func() {
		prices.RetrieveAllPrices(chainID)
		vaults.RetrieveAllVaults(chainID, vaultsMap)
		strategiesAddedList := events.HandleStrategyAdded(chainID, vaultsMap, 0, nil)
		strategies.RetrieveAllStrategies(chainID, strategiesAddedList)
		indexer.PostProcessStrategies(chainID)
		go func() {
			dailyBlockNumber.Run(chainID)
			vaultPricePerShare.Run(chainID)
			os.Exit(1)
			apy.ComputeChainAPR(chainID)
		}()
		strategies.InitRiskScore(chainID)
	})

	cron.Every(20).Minute().Do(func() {
		apy.ComputeChainAPR(chainID)
	})

	cron.Every(1).Day().At("12:10").Do(func() {
		vaultPricePerShare.Run(chainID)
	})
	cron.StartAsync()
	go registries.IndexNewVaults(chainID)
}

func InitializeBribes(chainID uint64) {
	allRewardsAdded := events.HandleRewardsAdded(chainID, 0, nil)
	for _, reward := range allRewardsAdded {
		bribes.SetInRewardAddedMap(chainID, reward)
	}
}
