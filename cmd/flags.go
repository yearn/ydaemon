package main

import (
	"flag"

	"github.com/yearn/ydaemon/common/logs"
)

var chains = []uint64{}
var endBlock *uint64
var process TProcess

func initFlags() {
	/**********************************************************************************************
	** Flag group: Chains
	** Description: List of chain IDs to run yDaemon for
	** Default: All supported chains (1, 10, 137, 250, 8453, 42161)
	**********************************************************************************************/
	rawChains := flag.String(`chains`, `1,10,100,137,146,250,8453,42161,747474`, `List of chain IDs to run yDaemon for: --chains 1,10,100,137,146,250,8453,42161,747474`)

	/**********************************************************************************************
	** Flag group: BlockNunbers
	** Description: Indicate the start and end block numbers
	** Default: 0 - nil
	**********************************************************************************************/
	endBlock = flag.Uint64(`end`, 0, `End block number`)

	/**********************************************************************************************
	** Flag group: Process
	** Description: Define the process to run
	** Default: daemon
	**********************************************************************************************/
	rawProcess := flag.String(`process`, `daemon`, `Define the process to run: --process daemon`)
	flag.Parse()
	if *endBlock == 0 {
		endBlock = nil
	}

	logs.Info(`Initializing chains...`)
	handleChainsInitialization(rawChains)
	logs.Info(`Initializing process...`)
	handleProcessInitialization(rawProcess)
}
