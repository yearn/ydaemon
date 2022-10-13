package main

import (
	"strconv"
	"sync"

	"github.com/yearn/ydaemon/common/env"
	"github.com/yearn/ydaemon/common/logs"
	"github.com/yearn/ydaemon/common/store"
	"github.com/yearn/ydaemon/internal"
)

var chains = env.SUPPORTED_CHAIN_IDS

// var chains = []uint64{1}

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

func init() {
	store.OpenDB()
}

func main() {
	defer store.CloseDB()

	logs.Info(`Loading store data to yDaemon memory...`)
	loadDaemonsForAllChains()
	logs.Info(`Summoning yDaemon...`)
	summonDaemonsForAllChains()
	logs.Success(`Server ready!`)

	internal.Initialize(1) //For later

	NewRouter().Run()
}
