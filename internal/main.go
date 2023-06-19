package internal

import (
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/internal/events"
	"github.com/yearn/ydaemon/internal/indexer"
	bribes "github.com/yearn/ydaemon/internal/indexer.bribes"
	"github.com/yearn/ydaemon/internal/models"
	"github.com/yearn/ydaemon/internal/prices"
	"github.com/yearn/ydaemon/internal/registries"
	"github.com/yearn/ydaemon/internal/strategies"
	"github.com/yearn/ydaemon/internal/tokens"
	"github.com/yearn/ydaemon/internal/vaults"
)

var STRATLIST = []models.TStrategy{}

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

	var internalWG sync.WaitGroup
	//From the events in the registries, retrieve all the vaults -> Should only be done on init or when a new vault is added
	vaultsMap := registries.RegisterAllVaults(chainID, 0, nil)

	// From our list of vaults, retrieve the ERC20 data for each shareToken, underlyingToken and the underlying of those tokens
	// -> Data store will not change unless extreme event, so this should only be done on init or when a new vault is added
	tokens.RetrieveAllTokens(chainID, vaultsMap)

	// From our list of tokens, retieve the price for each one of them -> Should be done every 1(?) minute for all tokens
	internalWG.Add(1)
	go runRetrieveAllPrices(chainID, &internalWG, 10*time.Minute)
	internalWG.Wait()

	//From our list of vault, perform a multicall to get all vaults data -> Should be done every 5(?) minutes for all vaults
	internalWG.Add(1)
	go runRetrieveAllVaults(chainID, vaultsMap, &internalWG, 12*time.Minute)
	internalWG.Wait()

	strategiesAddedList := events.HandleStrategyAdded(chainID, vaultsMap, 0, nil)

	//From our list of strategies, perform a multicall to get all strategies data -> Should be done every 5(?) minutes for all strategies
	internalWG.Add(1)
	go runRetrieveAllStrategies(chainID, strategiesAddedList, &internalWG, 15*time.Minute)
	internalWG.Wait()

	go registries.IndexNewVaults(chainID)
}

func InitializeBribes(chainID uint64) {
	allRewardsAdded := events.HandleRewardsAdded(chainID, 0, nil)
	for _, reward := range allRewardsAdded {
		bribes.SetInRewardAddedMap(chainID, reward)
	}
}
