package main

import (
	"sync"

	"github.com/yearn/ydaemon/common/logs"
	"github.com/yearn/ydaemon/internal"
	"github.com/yearn/ydaemon/internal/meta"
	"github.com/yearn/ydaemon/internal/strategies"
	"github.com/yearn/ydaemon/processes/partnerFees"
)

func SummonDaemonsw(chainID uint64) {
	var wg sync.WaitGroup

	// This first work group does not need any other data to be able to work.
	// They can all be summoned at the same time, with no dependencies.
	wg.Add(1)
	{
		//TODO: REPLACE WITH INTERNAL RELOADING
		go internal.InitializeV2(chainID, &wg)
	}
	wg.Wait()

	wg.Add(1)
	{
		//This can only be run after the internal daemons have been initialized
		go runDaemon(chainID, &wg, 0, strategies.ComputeRiskGroupAllocation)
	}
	wg.Wait()
}

func main() {
	initFlags()
	logs.Info(`Loading initial state in memory`)
	loadDaemonsForAllChains(trace)
	summonDaemonsForAllChains(trace)
	var wg sync.WaitGroup

	switch process {
	case ProcessServer:
		logs.Info(`Running yDaemon server process...`)

		for _, chainID := range chains {
			wg.Add(1)
			SummonDaemonsw(chainID)
			wg.Done()
		}
		wg.Wait()

		logs.Success(`Server ready!`)
		NewRouter().Run()
	case ProcessPartnerFees:
		logs.Info(`Running yDaemon partner fees process...`)

		for _, chainID := range chains {
			wg.Add(1)
			meta.RetrieveAllVaultsFromFiles(chainID)
			meta.RetrieveAllTokensFromFiles(chainID)
			meta.RetrieveAllStrategiesFromFiles(chainID)
			meta.RetrieveAllProtocolsFromFiles(chainID)
			go runDaemonWithBlocks(chainID, *startBlock, endBlock, &wg, 0, partnerFees.Run)
		}
		wg.Wait()
	}
}
