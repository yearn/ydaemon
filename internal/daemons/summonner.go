package daemons

import (
	"sync"
	"time"
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
	go runDaemon(chainID, &wgPrimary, time.Hour, FetchTokenList)
	go runDaemon(chainID, &wgPrimary, time.Hour, FetchStrategiesList)
	wgPrimary.Add(2)
	wgPrimary.Wait()

	var wg sync.WaitGroup
	time.Sleep(delay)
	go runDaemon(chainID, &wg, 0, FetchVaultsFromMeta)
	go runDaemon(chainID, &wg, 0, FetchTokensFromMeta)
	go runDaemon(chainID, &wg, 0, FetchStrategiesFromMeta)
	go runDaemon(chainID, &wg, 0, FetchProtocolsFromMeta)
	go runDaemon(chainID, &wg, time.Minute, FetchLens)
	go runDaemon(chainID, &wg, 10*time.Minute, FetchVaultsFromV1)
	go runDaemon(chainID, &wg, 10*time.Minute, FetchStrategiesMulticallData)
	go runDaemon(chainID, &wg, time.Hour, FetchStrategiesFromRisk)
	go runDaemon(chainID, &wg, 0, FetchPartnersFromFiles)
	wg.Add(9)
	wg.Wait()
}

// LoadDaemons is a function that loads the previous store state for a given chainID
func LoadDaemons(chainID uint64) {
	var wgPrimary sync.WaitGroup
	go LoadTokenList(chainID, &wgPrimary)
	go LoadStrategyList(chainID, &wgPrimary)
	wgPrimary.Add(2)
	wgPrimary.Wait()

	var wg sync.WaitGroup
	go LoadLens(chainID, &wg)
	go LoadAPIV1Vaults(chainID, &wg)
	go LoadStrategyMulticallData(chainID, &wg)
	go LoadRiskStrategies(chainID, &wg)
	wg.Add(4)
	wg.Wait()
}
