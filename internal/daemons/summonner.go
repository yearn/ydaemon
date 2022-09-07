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
	go runDaemon(chainID, &wgPrimary, time.Hour, subgraphDaemons.FetchTokenList)
	go runDaemon(chainID, &wgPrimary, time.Hour, subgraphDaemons.FetchStrategiesList)
	wgPrimary.Add(2)
	wgPrimary.Wait()

	var wg sync.WaitGroup
	time.Sleep(delay)
	go runDaemon(chainID, &wg, 0, metaDaemons.FetchVaultsFromMeta)
	go runDaemon(chainID, &wg, 0, metaDaemons.FetchTokensFromMeta)
	go runDaemon(chainID, &wg, 0, metaDaemons.FetchStrategiesFromMeta)
	go runDaemon(chainID, &wg, 0, metaDaemons.FetchProtocolsFromMeta)
	go runDaemon(chainID, &wg, 0, partnersDaemons.FetchPartnersFromFiles)
	go runDaemon(chainID, &wg, time.Minute, multicallDaemons.FetchLens)
	go runDaemon(chainID, &wg, 10*time.Minute, multicallDaemons.FetchStrategiesMulticallData)
	go runDaemon(chainID, &wg, 10*time.Minute, FetchVaultsFromV1)
	go runDaemon(chainID, &wg, time.Hour, FetchStrategiesFromRisk)
	wg.Add(9)
	wg.Wait()
}

// LoadDaemons is a function that loads the previous store state for a given chainID
func LoadDaemons(chainID uint64) {
	var wgPrimary sync.WaitGroup
	go subgraphDaemons.LoadTokenList(chainID, &wgPrimary)
	go subgraphDaemons.LoadStrategyList(chainID, &wgPrimary)
	wgPrimary.Add(2)
	wgPrimary.Wait()

	var wg sync.WaitGroup
	go multicallDaemons.LoadLens(chainID, &wg)
	go multicallDaemons.LoadStrategyMulticallData(chainID, &wg)
	go LoadAPIV1Vaults(chainID, &wg)
	go LoadRiskStrategies(chainID, &wg)
	wg.Add(4)
	wg.Wait()
}
