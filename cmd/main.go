package main

import (
	"strconv"
	"sync"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/yearn/ydaemon/common/ethereum"
	"github.com/yearn/ydaemon/common/logs"
	"github.com/yearn/ydaemon/internal"
	"github.com/yearn/ydaemon/processes/apr"
	"github.com/yearn/ydaemon/processes/initDailyBlock"
	"github.com/yearn/ydaemon/processes/vaultsMigrations"
)

func processServer(chainID uint64) {
	setStatusForChainID(chainID, `Loading`)
	defer setStatusForChainID(chainID, `OK`)

	scheduler := gocron.NewScheduler(time.UTC)
	logs.Info(`Getting WS client for chain ` + strconv.FormatUint(chainID, 10))
	ethereum.GetWSClient(chainID)
	ethereum.InitBlockTimestamp(chainID)
	internal.InitializeV2(chainID, nil, scheduler)
	triggerInitializedStatus(chainID)
}

var version = "" //version to display based on github tag
func getVersion() string {
	if version == "" {
		return "dev"
	}
	//trim to only get the first 7 characters of the commit hash
	return version[:7]
}

func main() {
	initFlags()
	summonDaemonsForAllChains(chains)
	go listenToSignals()

	var wg sync.WaitGroup
	logs.Info(`Running external processes...`)

	/** 🔵 - Yearn *************************************************************************************
	** This section of the code is responsible for running external processes. Each process is associated
	** with a specific chain ID and is run concurrently for efficiency. The processes include fetching
	** and updating prices, initializing V2, building token lists, running vault migrations, initializing
	** daily blocks, initializing APY, etc.
	**
	** The status of each process is tracked and updated for each chain ID. Once all processes for a
	** chain ID have completed, the status is set to "OK".
	**
	** The server is ready to accept requests once all processes have completed.
	**************************************************************************************************/
	switch process {
	case ProcessServer:
		logs.Info(`Running yDaemon server process...`)
		go NewRouter().Run(`:8080`)
		go triggerTgMessage(`💛 - yDaemon v` + getVersion() + ` is ready to accept requests: https://ydaemon.yearn.fi/`)

		for _, chainID := range chains {
			go processServer(chainID)
		}
		logs.Success(`Server ready on port 8080 !`)
		select {}

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
				apr.Run(chainID)
				wg.Done()
			}(chainID)
		}
		wg.Wait()
	}
}
