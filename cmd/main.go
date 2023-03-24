package main

import (
	"sync"

	"github.com/yearn/ydaemon/common/logs"
	"github.com/yearn/ydaemon/internal"
	"github.com/yearn/ydaemon/internal/meta"
	"github.com/yearn/ydaemon/internal/strategies"
	"github.com/yearn/ydaemon/processes/partnerFees"
)

// func sum() {
// 	n := bigNumber.NewFloat(5071.0209 + 6805.6828 + 7267.0110 + 7835.6268 + 37196.2244 + 40731.8349 + 1877155.2153 + 1881310.8654 + 6614548.5187 + 6617517.2512 + 6618142.6883 + 6618217.9946 + 6695299.6065 + 6696034.1870)
// 	logs.Pretty(n.Text('f', -1))
// }

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
			go runDaemonWithBlocks(chainID, *startBlock, endBlock, &wg, 0, partnerFees.Run)
		}
		wg.Wait()
	}
}
