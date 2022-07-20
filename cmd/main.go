package main

import (
	"strconv"
	"sync"
	"time"

	"github.com/majorfi/ydaemon/internal/controllers"
	"github.com/majorfi/ydaemon/internal/daemons"
	"github.com/majorfi/ydaemon/internal/logs"
	"github.com/majorfi/ydaemon/internal/store"
)

// var chains = []uint64{1, 250, 42161}
var chains = []uint64{1}

func waitGroupSummonDaemons(wg *sync.WaitGroup, chainID uint64, delay time.Duration) {
	daemons.SummonDaemons(chainID, delay)
	logs.Success(`Daemons for chainID ` + strconv.Itoa(int(chainID)) + ` summoned successfully!`)
	wg.Done()
}

func summonDaemonsForAllChains() {
	var wg sync.WaitGroup
	for index, chainID := range chains {
		wg.Add(1)
		go waitGroupSummonDaemons(&wg, chainID, time.Duration(index))
	}
	wg.Wait()
}

func waitGroupLoadDaemons(wg *sync.WaitGroup, chainID uint64) {
	daemons.LoadDaemons(chainID)
	logs.Success(`Store data loaded in yDaemon memory for chainUD ` + strconv.Itoa(int(chainID)) + `!`)
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
	controllers.NewRouter().Run()
}
