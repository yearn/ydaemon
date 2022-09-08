package main

import (
	"sync"
	"time"

	"github.com/yearn/ydaemon/internal/meta"
	"github.com/yearn/ydaemon/internal/partners"
	"github.com/yearn/ydaemon/internal/prices"
	"github.com/yearn/ydaemon/internal/strategies"
	"github.com/yearn/ydaemon/internal/tokens"
	"github.com/yearn/ydaemon/internal/vaults"
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
	go runDaemon(chainID, &wgPrimary, time.Hour, tokens.FetchTokenList)
	go runDaemon(chainID, &wgPrimary, time.Hour, strategies.FetchStrategiesList)
	wgPrimary.Wait()

	var wgSecondary sync.WaitGroup
	time.Sleep(delay)
	wgSecondary.Add(9)
	go runDaemon(chainID, &wgSecondary, 0, meta.FetchVaultsFromMeta)
	go runDaemon(chainID, &wgSecondary, 0, meta.FetchTokensFromMeta)
	go runDaemon(chainID, &wgSecondary, 0, meta.FetchStrategiesFromMeta)
	go runDaemon(chainID, &wgSecondary, 0, meta.FetchProtocolsFromMeta)
	go runDaemon(chainID, &wgSecondary, 0, partners.FetchPartnersFromFiles)
	go runDaemon(chainID, &wgSecondary, 10*time.Minute, vaults.FetchVaultMulticallData)
	go runDaemon(chainID, &wgSecondary, 10*time.Minute, strategies.FetchStrategiesMulticallData)
	go runDaemon(chainID, &wgSecondary, 10*time.Minute, vaults.FetchVaultsFromV1)
	go runDaemon(chainID, &wgSecondary, time.Hour, strategies.FetchStrategiesFromRisk)
	wgSecondary.Wait()

	var wg sync.WaitGroup
	time.Sleep(delay)
	wg.Add(1)
	go runDaemon(chainID, &wg, time.Minute, prices.FetchLens)
	wg.Wait()
}

// LoadDaemons is a function that loads the previous store state for a given chainID
func LoadDaemons(chainID uint64) {
	var wgPrimary sync.WaitGroup
	wgPrimary.Add(2)
	go tokens.LoadTokenList(chainID, &wgPrimary)
	go strategies.LoadStrategyList(chainID, &wgPrimary)
	wgPrimary.Wait()

	var wgSecondary sync.WaitGroup
	wgSecondary.Add(4)
	go strategies.LoadStrategyMulticallData(chainID, &wgSecondary)
	go strategies.LoadRiskStrategies(chainID, &wgSecondary)
	go vaults.LoadVaultMulticallData(chainID, &wgSecondary)
	go vaults.LoadAPIV1Vaults(chainID, &wgSecondary)
	wgSecondary.Wait()

	var wg sync.WaitGroup
	wg.Add(1)
	go prices.LoadLens(chainID, &wg)
	wg.Wait()
}
