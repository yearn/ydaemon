package main

import (
	"strconv"
	"sync"

	"github.com/yearn/ydaemon/internal/utils/logs"
	"github.com/yearn/ydaemon/internal/utils/store"
)

// var chains = helpers.SUPPORTED_CHAIN_IDS

var chains = []uint64{250}

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

func main() {
	store.OpenDB()
	defer store.CloseDB()

	logs.Info(`Loading store data to yDaemon memory...`)
	loadDaemonsForAllChains()
	logs.Info(`Summoning yDaemon...`)
	summonDaemonsForAllChains()
	logs.Success(`Server ready!`)
	NewRouter().Run()
}
