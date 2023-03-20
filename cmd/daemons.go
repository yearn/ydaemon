package main

import (
	"strconv"
	"sync"
	"time"

	"github.com/yearn/ydaemon/common/env"
	"github.com/yearn/ydaemon/common/helpers"
	"github.com/yearn/ydaemon/common/logs"
	"github.com/yearn/ydaemon/common/traces"
	"github.com/yearn/ydaemon/external/partners"
	"github.com/yearn/ydaemon/external/vaults"
	"github.com/yearn/ydaemon/internal/meta"
	"github.com/yearn/ydaemon/internal/strategies"
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
	if !helpers.Contains(env.SUPPORTED_CHAIN_IDS, chainID) {
		return
	}

	var wg sync.WaitGroup

	// This first work group does not need any other data to be able to work.
	// They can all be summoned at the same time, with no dependencies.
	wg.Add(7)
	{
		go runDaemon(chainID, &wg, 0, meta.RetrieveAllVaultsFromFiles)
		go runDaemon(chainID, &wg, 0, meta.RetrieveAllTokensFromFiles)
		go runDaemon(chainID, &wg, 0, meta.RetrieveAllStrategiesFromFiles)
		go runDaemon(chainID, &wg, 0, meta.RetrieveAllProtocolsFromFiles)
		go runDaemon(chainID, &wg, 0, partners.FetchPartnersFromFiles)
		go runDaemon(chainID, &wg, 0, strategies.RetrieveAllRisksGroupsFromFiles)
		go runDaemon(chainID, &wg, 10*time.Minute, vaults.FetchVaultsFromV1)
	}
	wg.Wait()
}

// LoadDaemons is a function that loads the previous store state for a given chainID
func LoadDaemons(chainID uint64) {
	if !helpers.Contains(env.SUPPORTED_CHAIN_IDS, chainID) {
		return
	}

	vaults.LoadAggregatedVaults(chainID, nil)
}

func waitGroupSummonDaemons(trace *traces.TTrace, wg *sync.WaitGroup, chainID uint64) {
	trace = trace.Child(
		`app.bootstrap.summon.daemon`,
		traces.TTags{Name: `chainID`, Value: strconv.Itoa(int(chainID))},
	)
	defer trace.Finish()

	SummonDaemons(chainID)
	wg.Done()
	logs.Success(`Daemons for chainID ` + strconv.Itoa(int(chainID)) + ` summoned successfully!`)
}

func summonDaemonsForAllChains(trace *traces.TTrace) {
	trace = trace.Child(`app.bootstrap.summon.all`)
	defer trace.Finish()

	var wg sync.WaitGroup
	for _, chainID := range chains {
		wg.Add(1)
		go waitGroupSummonDaemons(trace, &wg, chainID)
	}

	wg.Wait()
}

func waitGroupLoadDaemons(trace *traces.TTrace, wg *sync.WaitGroup, chainID uint64) {
	trace = trace.Child(
		`app.bootstrap.load_state.daemon`,
		traces.TTags{Name: `chainID`, Value: strconv.Itoa(int(chainID))},
	)
	defer trace.Finish()

	LoadDaemons(chainID)
	wg.Done()
	logs.Success(`Store data loaded in yDaemon memory for chainID ` + strconv.Itoa(int(chainID)) + `!`)
}

func loadDaemonsForAllChains(trace *traces.TTrace) {
	trace = trace.Child(`app.bootstrap.load_state.all`)
	defer trace.Finish()

	var wg sync.WaitGroup
	for _, chainID := range chains {
		wg.Add(1)
		go waitGroupLoadDaemons(trace, &wg, chainID)
	}
	wg.Wait()
}
