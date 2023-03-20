package main

import (
	"sync"

	"github.com/yearn/ydaemon/common/logs"
	"github.com/yearn/ydaemon/internal"
	"github.com/yearn/ydaemon/internal/meta"
	"github.com/yearn/ydaemon/internal/strategies"
	"github.com/yearn/ydaemon/processes/partnerFees"
)

func main() {
	logs.Info(`Loading initial state in memory`)
	loadDaemonsForAllChains(trace)
	summonDaemonsForAllChains(trace)
	var wg sync.WaitGroup

	switch process {
	case ProcessServer:
		logs.Info(`Running yDaemon server process...`)

		for _, chainID := range chains {
			wg.Add(1)
			innerWg := sync.WaitGroup{}
			innerWg.Add(1)
			go internal.InitializeV2(chainID, &innerWg)
			innerWg.Wait()

			innerWg.Add(1)
			go runDaemon(chainID, &innerWg, 0, strategies.ComputeRiskGroupAllocation)
			innerWg.Wait()
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
			go runDaemon(chainID, &wg, 0, partnerFees.Run)
		}
		wg.Wait()
	}
}
