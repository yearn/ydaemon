package main

import (
	"strconv"
	"sync"
	"time"

	"github.com/yearn/ydaemon/common/env"
	"github.com/yearn/ydaemon/common/logs"
	"github.com/yearn/ydaemon/external/partners"
	"github.com/yearn/ydaemon/internal/meta"
	"github.com/yearn/ydaemon/internal/risk"
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
	if _, ok := env.CHAINS[chainID]; !ok {
		return
	}

	var wg sync.WaitGroup

	// This first work group does not need any other data to be able to work.
	// They can all be summoned at the same time, with no dependencies.
	wg.Add(3)
	{
		go runDaemon(chainID, &wg, 0, meta.RetrieveAllProtocolsFromFiles)
		go runDaemon(chainID, &wg, 0, partners.FetchPartnersFromFiles)
		go runDaemon(chainID, &wg, 0, risk.RetrieveAllRisksGroupsFromFiles)
	}
	wg.Wait()
}

func waitGroupSummonDaemons(wg *sync.WaitGroup, chainID uint64) {
	SummonDaemons(chainID)
	wg.Done()
	logs.Success(`Daemons for chainID ` + strconv.Itoa(int(chainID)) + ` summoned successfully!`)
}

func summonDaemonsForAllChains(chains []uint64) {
	var wg sync.WaitGroup
	for _, chainID := range chains {
		wg.Add(1)
		go waitGroupSummonDaemons(&wg, chainID)
	}

	wg.Wait()
}
