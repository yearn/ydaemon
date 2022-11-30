package main

import (
	"strconv"
	"sync"

	"github.com/yearn/ydaemon/common/logs"
	"github.com/yearn/ydaemon/common/traces"
)

// var chains = env.SUPPORTED_CHAIN_IDS

var chains = []uint64{1}

func waitGroupSummonDaemons(trace *traces.TTrace, wg *sync.WaitGroup, chainID uint64) {
	trace = trace.Child(
		`app.bootstrap.summon.daemon`,
		traces.TTags{Name: `chainID`, Value: strconv.Itoa(int(chainID))},
	)
	defer trace.Finish()

	SummonDaemons(chainID)
	wg.Done()
	logs.Success(`Daemons for chainID ` + strconv.Itoa(int(chainID)) + ` summoned successfully!`)
}

func summonDaemonsForAllChains(trace *traces.TTrace) {
	trace = trace.Child(`app.bootstrap.summon.all`)
	defer trace.Finish()

	var wg sync.WaitGroup
	for _, chainID := range chains {
		wg.Add(1)
		go waitGroupSummonDaemons(trace, &wg, chainID)
	}

	wg.Wait()
}

func waitGroupLoadDaemons(trace *traces.TTrace, wg *sync.WaitGroup, chainID uint64) {
	trace = trace.Child(
		`app.bootstrap.load_state.daemon`,
		traces.TTags{Name: `chainID`, Value: strconv.Itoa(int(chainID))},
	)
	defer trace.Finish()

	LoadDaemons(chainID)
	wg.Done()
	logs.Success(`Store data loaded in yDaemon memory for chainID ` + strconv.Itoa(int(chainID)) + `!`)
}

func loadDaemonsForAllChains(trace *traces.TTrace) {
	trace = trace.Child(`app.bootstrap.load_state.all`)
	defer trace.Finish()

	var wg sync.WaitGroup
	for _, chainID := range chains {
		wg.Add(1)
		go waitGroupLoadDaemons(trace, &wg, chainID)
	}
	wg.Wait()
}

func initialize() {
	traces.SetupSentry()
	trace := traces.Init(`app.bootstrap`).SetTag(`subsystem`, `main`)
	defer trace.Finish()

	logs.Info(`Loading store data to yDaemon memory...`)
	loadDaemonsForAllChains(trace)

	logs.Info(`Summoning yDaemon...`)
	summonDaemonsForAllChains(trace)
}

func main() {
	go initialize()

	logs.Success(`Server ready!`)
	NewRouter().Run()
}
