package daemons

import (
	"sync"
	"testing"

	"github.com/majorfi/ydaemon/internal/utils"
)

// init is fired directly on app start and prepare the multicall clients
func TestDaemons(t *testing.T) {
	var wg sync.WaitGroup

	//Testing for chainID == 1
	go func(wg *sync.WaitGroup) {
		SummonDaemons(1, 0)
		wg.Done()
	}(&wg)

	//Testing for chainID == 4
	go func(wg *sync.WaitGroup) {
		SummonDaemons(4, 0)
		wg.Done()
	}(&wg)

	//Testing for chainID == 250
	go func(wg *sync.WaitGroup) {
		SummonDaemons(250, 0)
		wg.Done()
	}(&wg)

	//Testing for chainID == 42161
	go func(wg *sync.WaitGroup) {
		SummonDaemons(42161, 0)
		wg.Done()
	}(&wg)
	wg.Add(4)
	wg.Wait()

	//Edge case testing with values overwrite
	utils.API_V1_BASE_URL = `https://thisisarandomurithatshouldnotbeset.yearn.finance/`
	utils.META_BASE_URL = `https://thisisanotherurithatshouldnotbeset.yearn.finance/`
	SummonDaemons(1, 0)
}
