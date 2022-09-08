package main

import (
	"sync"
	"testing"

	"github.com/yearn/ydaemon/internal/utils/helpers"
	"github.com/yearn/ydaemon/internal/utils/store"
)

// init is fired directly on app start and prepare the multicall clients
func TestDaemons(t *testing.T) {
	store.OpenDB()
	defer store.CloseDB()

	var wg sync.WaitGroup

	//Testing for chainID == 1
	go func(wg *sync.WaitGroup) {
		LoadDaemons(1)
		SummonDaemons(1, 0)
		wg.Done()
	}(&wg)

	//Testing for chainID == 4
	go func(wg *sync.WaitGroup) {
		LoadDaemons(4)
		SummonDaemons(4, 0)
		wg.Done()
	}(&wg)

	//Testing for chainID == 10
	go func(wg *sync.WaitGroup) {
		LoadDaemons(10)
		SummonDaemons(10, 0)
		wg.Done()
	}(&wg)

	//Testing for chainID == 250
	go func(wg *sync.WaitGroup) {
		LoadDaemons(250)
		SummonDaemons(250, 0)
		wg.Done()
	}(&wg)

	//Testing for chainID == 42161
	go func(wg *sync.WaitGroup) {
		LoadDaemons(42161)
		SummonDaemons(42161, 0)
		wg.Done()
	}(&wg)
	wg.Add(5)
	wg.Wait()

	//Edge case testing with values overwrite
	helpers.API_V1_BASE_URL = `https://thisisarandomurithatshouldnotbeset.yearn.finance/`
	SummonDaemons(1, 0)
}
