package main

import (
	"sync"
	"time"

	"github.com/yearn/ydaemon/common/logs"
	"github.com/yearn/ydaemon/internal"
	"github.com/yearn/ydaemon/internal/meta"
	"github.com/yearn/ydaemon/internal/strategies"
	"github.com/yearn/ydaemon/processes/partnerFees"
	"github.com/yearn/ydaemon/processes/tokenList"
	"github.com/yearn/ydaemon/processes/vaultsMigrations"
)

func SummonDaemonsw(chainID uint64, parentWg *sync.WaitGroup) {
	if parentWg != nil {
		defer parentWg.Done()
	}
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
		go runDaemon(chainID, &wg, 10*time.Minute, strategies.InitRiskScore)
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
		go NewRouter().Run(`:8080`)

		for _, chainID := range chains {
			wg.Add(1)
			go SummonDaemonsw(chainID, &wg)
		}
		wg.Wait()

		logs.Success(`Server ready on port 8080 !`)
		// pprof.Lookup("goroutine").WriteTo(os.Stdout, 1)

		select {}

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

	case ProcessTokenList:
		logs.Info(`Running yDaemon token list process...`)

		for _, chainID := range chains {
			wg.Add(1)
			go func(chainID uint64) {
				tokenList.BuildTokenList(chainID)
				wg.Done()
			}(chainID)
		}
		wg.Wait()

	case ProcessVaultMigrations:
		logs.Info(`Running yDaemon vault migrations process...`)
		for _, chainID := range chains {
			wg.Add(1)
			go func(chainID uint64) {
				vaultsMigrations.Run(chainID)
				wg.Done()
			}(chainID)
		}
		wg.Wait()
	}
}
