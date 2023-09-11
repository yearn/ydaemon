package main

import (
	"sync"

	"github.com/yearn/ydaemon/common/logs"
	"github.com/yearn/ydaemon/internal"
	"github.com/yearn/ydaemon/processes/apy"
	"github.com/yearn/ydaemon/processes/initDailyBlock"
	"github.com/yearn/ydaemon/processes/tokenList"
	"github.com/yearn/ydaemon/processes/vaultsMigrations"
)

func main() {
	initFlags()
	loadDaemonsForAllChains(chains)
	summonDaemonsForAllChains(chains)
	var wg sync.WaitGroup

	switch process {
	case ProcessServer:
		logs.Info(`Running yDaemon server process...`)
		go NewRouter().Run(`:8080`)

		for _, chainID := range chains {
			setStatusForChainID(chainID, "Loading")
			wg.Add(1)
			go internal.InitializeV2(chainID, &wg)
		}
		wg.Wait()
		for _, chainID := range chains {
			setStatusForChainID(chainID, "OK")
		}

		logs.Success(`Server ready on port 8080 !`)
		select {}

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

	case ProcessInitDailyBlock:
		logs.Info(`Running yDaemon DailyBlock Initializer process...`)
		for _, chainID := range chains {
			wg.Add(1)
			go func(chainID uint64) {
				initDailyBlock.Run(chainID)
				wg.Done()
			}(chainID)
		}
		wg.Wait()

	case ProcessAPY:
		logs.Info(`Running yDaemon APY Initializer process...`)
		for _, chainID := range chains {
			wg.Add(1)
			go func(chainID uint64) {
				apy.Run(chainID)
				wg.Done()
			}(chainID)
		}
		wg.Wait()
	}
}
