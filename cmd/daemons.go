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
func SummonDaemons(chainID uint64) {
	var wg sync.WaitGroup
	// This first work group does not need any other data to be able to work.
	// They can all be summoned at the same time, with no dependencies.
	wg.Add(8)
	{
		go runDaemon(chainID, &wg, time.Hour, tokens.FetchTokenList)
		go runDaemon(chainID, &wg, time.Hour, strategies.FetchStrategiesList)
		go runDaemon(chainID, &wg, 0, meta.FetchVaultsFromMeta)
		go runDaemon(chainID, &wg, 0, meta.FetchTokensFromMeta)
		go runDaemon(chainID, &wg, 0, meta.FetchStrategiesFromMeta)
		go runDaemon(chainID, &wg, 0, meta.FetchProtocolsFromMeta)
		go runDaemon(chainID, &wg, 0, partners.FetchPartnersFromFiles)
		go runDaemon(chainID, &wg, 10*time.Minute, vaults.FetchVaultsFromV1)
	}
	wg.Wait()

	wg.Add(1)
	{
		go runDaemon(chainID, &wg, 10*time.Minute, strategies.FetchWithdrawalQueueMulticallData)
	}
	wg.Wait()

	wg.Add(2)
	{
		//Require tokens.FetchTokenList to be done
		go runDaemon(chainID, &wg, 10*time.Minute, vaults.FetchVaultMulticallData)
		//Require strategies.FetchWithdrawalQueueMulticallData to be done
		go runDaemon(chainID, &wg, 10*time.Minute, strategies.FetchStrategiesMulticallData)
	}
	wg.Wait()

	wg.Add(1)
	{
		//Require vaults.FetchVaultMulticallData to be done
		go runDaemon(chainID, &wg, time.Minute, prices.FetchLens)
	}
	wg.Wait()
	wg.Add(1)
	{
		//Require strategies.FetchLens to be done
		go runDaemon(chainID, &wg, time.Minute, strategies.FetchStrategiesFromRisk)
	}

	wg.Wait()
}

// LoadDaemons is a function that loads the previous store state for a given chainID
func LoadDaemons(chainID uint64) {
	var wg sync.WaitGroup
	wg.Add(2)
	go tokens.LoadTokenList(chainID, &wg)
	go strategies.LoadStrategyList(chainID, &wg)
	wg.Wait()

	wg.Add(1)
	go strategies.LoadWithdrawalQueueMulticallData(chainID, &wg)
	wg.Wait()

	wg.Add(3)
	go strategies.LoadStrategyMulticallData(chainID, &wg)
	go vaults.LoadVaultMulticallData(chainID, &wg)
	go vaults.LoadAPIV1Vaults(chainID, &wg)
	wg.Wait()

	wg.Add(1)
	go prices.LoadLens(chainID, &wg)
	wg.Wait()
}
