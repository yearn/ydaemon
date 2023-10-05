package main

import (
	"flag"
)

var chains = []uint64{}
var shouldEnableSentry *bool
var startBlock *uint64
var endBlock *uint64
var process TProcess

func initFlags() {
	/**********************************************************************************************
	** Flag group: Chains
	** Description: List of chain IDs to run yDaemon for
	** Default: All supported chains (1, 10, 137, 250, 8453, 42161)
	**********************************************************************************************/
	rawChains := flag.String(`chains`, `1,10,137,250,8453,42161`, `List of chain IDs to run yDaemon for: --chains 1,10,137,250,8453,42161`)

	/**********************************************************************************************
	** Flag group: Sentry
	** Description: Enable Sentry
	** Default: true
	**********************************************************************************************/
	shouldEnableSentry = flag.Bool(`sentry`, true, `Enable Sentry`)

	/**********************************************************************************************
	** Flag group: BlockNunbers
	** Description: Indicate the start and end block numbers
	** Default: 0 - nil
	**********************************************************************************************/
	startBlock = flag.Uint64(`start`, 0, `Start block number`)
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

	handleChainsInitialization(rawChains)
	handleSentryInitialization()
	handleProcessInitialization(rawProcess)
}
