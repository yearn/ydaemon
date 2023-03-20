package main

import (
	"flag"
)

var chains = []uint64{}
var shouldEnableSentry *bool
var process TProcess

func init() {
	/**********************************************************************************************
	** Flag group: Chains
	** Description: List of chain IDs to run yDaemon for
	** Default: All supported chains (1, 10, 250, 42161)
	**********************************************************************************************/
	flag.Var(&chainFlag, "chains", `List of chain IDs to run yDaemon for: --chains 1 10 250 42161`)

	/**********************************************************************************************
	** Flag group: Sentry
	** Description: Enable Sentry
	** Default: true
	**********************************************************************************************/
	shouldEnableSentry = flag.Bool(`sentry`, true, `Enable Sentry`)

	/**********************************************************************************************
	** Flag group: Process
	** Description: Define the process to run
	** Default: daemon
	**********************************************************************************************/
	rawProcess := flag.String(`process`, `daemon`, `Define the process to run: --process daemon`)
	flag.Parse()

	handleChainsInitialization()
	handleSentryInitialization()
	handleProcessInitialization(rawProcess)
}
