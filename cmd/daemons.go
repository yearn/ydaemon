package main

import (
	"sync"
	"time"

	"github.com/yearn/ydaemon/common/env"
	"github.com/yearn/ydaemon/common/helpers"
	"github.com/yearn/ydaemon/external/meta"
	"github.com/yearn/ydaemon/external/partners"
	"github.com/yearn/ydaemon/external/prices"
	"github.com/yearn/ydaemon/external/strategies"
	"github.com/yearn/ydaemon/external/vaults"
	"github.com/yearn/ydaemon/internal"
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
func SummonDaemons(chainID uint64) {
	if !helpers.ContainsUint64(env.SUPPORTED_CHAIN_IDS, chainID) {
		return
	}

	var wg sync.WaitGroup

	// This first work group does not need any other data to be able to work.
	// They can all be summoned at the same time, with no dependencies.
	wg.Add(8)
	{
		//TODO: REPLACE WITH INTERNAL RELOADING
		go runDaemon(chainID, &wg, 1*time.Minute, internal.LoadTokens)
		go runDaemon(chainID, &wg, time.Hour, strategies.FetchStrategiesList)
		go runDaemon(chainID, &wg, 0, meta.FetchVaultsFromMeta)
		go runDaemon(chainID, &wg, 0, meta.FetchTokensFromMeta)
		go runDaemon(chainID, &wg, 0, meta.FetchStrategiesFromMeta)
		go runDaemon(chainID, &wg, 0, meta.FetchProtocolsFromMeta)
		go runDaemon(chainID, &wg, 0, partners.FetchPartnersFromFiles)
		go runDaemon(chainID, &wg, 10*time.Minute, vaults.FetchVaultsFromV1)
	}
	wg.Wait()

	wg.Add(2)
	{
		go runDaemon(chainID, &wg, 1*time.Minute, internal.LoadVaults)
		go runDaemon(chainID, &wg, 10*time.Minute, strategies.FetchWithdrawalQueueMulticallData)
	}
	wg.Wait()

	wg.Add(2)
	{
		//Require strategies.FetchWithdrawalQueueMulticallData to be done
		go runDaemon(chainID, &wg, 10*time.Minute, strategies.FetchStrategiesMulticallData)
		//Require vaults.FetchVaultMulticallData to be done
		go runDaemon(chainID, &wg, time.Minute, prices.FetchLens)
	}
	wg.Wait()
	wg.Add(1)
	{
		//Require prices.FetchLens to be done
		go runDaemon(chainID, &wg, time.Hour, strategies.FetchStrategiesFromRisk)
	}

	wg.Wait()
}

// LoadDaemons is a function that loads the previous store state for a given chainID
func LoadDaemons(chainID uint64) {
	if !helpers.ContainsUint64(env.SUPPORTED_CHAIN_IDS, chainID) {
		return
	}

	var wg sync.WaitGroup
	wg.Add(1)
	go strategies.LoadAggregatedStrategies(chainID, &wg)
	wg.Wait()

	wg.Add(1)
	// go strategies.LoadStrategyMulticallData(chainID, &wg)
	// go vaults.LoadVaultMulticallData(chainID, &wg)
	go vaults.LoadAggregatedVaults(chainID, &wg)
	wg.Wait()

	wg.Add(1)
	go prices.LoadLens(chainID, &wg)
	wg.Wait()
}
