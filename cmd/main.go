package main

import (
	"github.com/yearn/ydaemon/common/ethereum"
	"github.com/yearn/ydaemon/common/logs"
	"github.com/yearn/ydaemon/internal"
	"github.com/yearn/ydaemon/internal/storage"
)

func processServer(chainID uint64) {
	setStatusForChainID(chainID, `Loading`)
	defer setStatusForChainID(chainID, `OK`)

	ethereum.GetWSClient(chainID, true)
	ethereum.InitBlockTimestamp(chainID)
	internal.InitializeV2(chainID, nil)
	TriggerInitializedStatus(chainID)
}

/**************************************************************************************************
** Main entry point for the daemon, handling everything from initialization to running external
** processes.
**************************************************************************************************/
func main() {
	initFlags()
	ethereum.Initialize()
	storage.InitializeStorage()
	go ListenToSignals()

	logs.Info(`Running yDaemon server process...`)
	go NewRouter().Run(`:8080`)
	go TriggerTgMessage(`💛 - yDaemon v` + GetVersion() + ` is ready to accept requests: https://ydaemon.yearn.fi/`)

	for _, chainID := range chains {
		go processServer(chainID)
	}
	logs.Success(`Server ready on port 8080 !`)
	select {}
}
