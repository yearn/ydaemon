package main

import (
	"log"
	"os"
	"strconv"
	"sync"

	"github.com/getsentry/sentry-go"
	"github.com/yearn/ydaemon/common/logs"
)

// var chains = env.SUPPORTED_CHAIN_IDS

var chains = []uint64{1}

func waitGroupSummonDaemons(wg *sync.WaitGroup, chainID uint64) {
	SummonDaemons(chainID)
	logs.Success(`Daemons for chainID ` + strconv.Itoa(int(chainID)) + ` summoned successfully!`)
	wg.Done()
}

func summonDaemonsForAllChains() {
	var wg sync.WaitGroup
	for _, chainID := range chains {
		wg.Add(1)
		go waitGroupSummonDaemons(&wg, chainID)
	}
	wg.Wait()
}

func waitGroupLoadDaemons(wg *sync.WaitGroup, chainID uint64) {
	LoadDaemons(chainID)
	logs.Success(`Store data loaded in yDaemon memory for chainID ` + strconv.Itoa(int(chainID)) + `!`)
	wg.Done()
}

func loadDaemonsForAllChains() {
	var wg sync.WaitGroup
	for _, chainID := range chains {
		wg.Add(1)
		go waitGroupLoadDaemons(&wg, chainID)
	}
	wg.Wait()
}

func setupSentry() {
	SENTRY_DSN, exists := os.LookupEnv("SENTRY_DSN")
	if exists {
		err := sentry.Init(sentry.ClientOptions{
			Dsn: SENTRY_DSN,
			// Set TracesSampleRate to 1.0 to capture 100%
			// of transactions for performance monitoring.
			// We recommend adjusting this value in production,
			TracesSampleRate: 1.0,
			// As it's not uncommon to panic with a string, it's
			// recommended to use the AttachStacktrace option
			// during SDK initialization, which will try to
			// provide a useful stack trace for messages as well.
			AttachStacktrace: true,
		})
		if err != nil {
			log.Fatalf("sentry.Init: %s", err)
		} else {
			logs.Info("Sentry initialized")
		}
	} else {
		logs.Warning("SENTRY_DSN not set, Sentry not initialized")
	}
}

func main() {
	setupSentry()

	logs.Info(`Loading store data to yDaemon memory...`)
	loadDaemonsForAllChains()
	logs.Info(`Summoning yDaemon...`)
	summonDaemonsForAllChains()

	logs.Success(`Server ready!`)
	NewRouter().Run()
}
