package daemons

import (
	"sync"
	"time"

	metaDaemons "github.com/yearn/ydaemon/internal/daemons/meta"
	multicallDaemons "github.com/yearn/ydaemon/internal/daemons/multicall"
	partnersDaemons "github.com/yearn/ydaemon/internal/daemons/partners"
	subgraphDaemons "github.com/yearn/ydaemon/internal/daemons/subgraph"
)

// runDaemon is a function that contains the standard flow to run a daemon
func runDaemon(chainID uint64, wg *sync.WaitGroup, delay time.Duration, performAction func(chainID uint64)) {
	isDone := false
	for {
		performAction(chainID)
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

// SummonDaemons is a function that summons the daemons for a given chainID.
func SummonDaemons(chainID uint64, delay time.Duration) {
	var wgPrimary sync.WaitGroup
	wgPrimary.Add(2)
	go runDaemon(chainID, &wgPrimary, time.Hour, subgraphDaemons.FetchTokenList)
	go runDaemon(chainID, &wgPrimary, time.Hour, subgraphDaemons.FetchStrategiesList)
	wgPrimary.Wait()

	var wgSecondary sync.WaitGroup
	time.Sleep(delay)
	wgSecondary.Add(9)
	go runDaemon(chainID, &wgSecondary, 0, metaDaemons.FetchVaultsFromMeta)
	go runDaemon(chainID, &wgSecondary, 0, metaDaemons.FetchTokensFromMeta)
	go runDaemon(chainID, &wgSecondary, 0, metaDaemons.FetchStrategiesFromMeta)
	go runDaemon(chainID, &wgSecondary, 0, metaDaemons.FetchProtocolsFromMeta)
	go runDaemon(chainID, &wgSecondary, 0, partnersDaemons.FetchPartnersFromFiles)
	go runDaemon(chainID, &wgSecondary, 10*time.Minute, multicallDaemons.FetchVaultMulticallData)
	go runDaemon(chainID, &wgSecondary, 10*time.Minute, multicallDaemons.FetchStrategiesMulticallData)
	go runDaemon(chainID, &wgSecondary, 10*time.Minute, FetchVaultsFromV1)
	go runDaemon(chainID, &wgSecondary, time.Hour, FetchStrategiesFromRisk)
	wgSecondary.Wait()

	var wg sync.WaitGroup
	time.Sleep(delay)
	wg.Add(1)
	go runDaemon(chainID, &wg, time.Minute, multicallDaemons.FetchLens)
	wg.Wait()
}

// LoadDaemons is a function that loads the previous store state for a given chainID
func LoadDaemons(chainID uint64) {
	var wgPrimary sync.WaitGroup
	wgPrimary.Add(2)
	go subgraphDaemons.LoadTokenList(chainID, &wgPrimary)
	go subgraphDaemons.LoadStrategyList(chainID, &wgPrimary)
	wgPrimary.Wait()

	var wgSecondary sync.WaitGroup
	wgSecondary.Add(4)
	go multicallDaemons.LoadStrategyMulticallData(chainID, &wgSecondary)
	go multicallDaemons.LoadVaultMulticallData(chainID, &wgSecondary)
	go LoadAPIV1Vaults(chainID, &wgSecondary)
	go LoadRiskStrategies(chainID, &wgSecondary)
	wgSecondary.Wait()

	var wg sync.WaitGroup
	wg.Add(1)
	go multicallDaemons.LoadLens(chainID, &wg)
	wg.Wait()
}
