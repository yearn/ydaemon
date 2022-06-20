package main

import (
	"strconv"
	"sync"
	"time"

	"github.com/majorfi/ydaemon/internal/controllers"
	"github.com/majorfi/ydaemon/internal/daemons"
	"github.com/majorfi/ydaemon/internal/logs"
)

func waitGroupSummonDaemons(wg *sync.WaitGroup, chainID uint64, delay time.Duration) {
	daemons.SummonDaemons(chainID, delay)
	chainIDAsString := strconv.Itoa(int(chainID))
	logs.Success(`Daemons for chainID ` + chainIDAsString + ` summoned successfully!`)
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	wg.Add(3)
	logs.Info(`Starting the server...`)
	go waitGroupSummonDaemons(&wg, 1, 0)
	go waitGroupSummonDaemons(&wg, 250, 0)
	go waitGroupSummonDaemons(&wg, 42161, 0)
	wg.Wait()
	logs.Success(`Server ready!`)

	controllers.NewRouter().Run()
}
