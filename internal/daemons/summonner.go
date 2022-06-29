package daemons

import (
	"sync"
	"time"
)

// SummonDaemons is a function that summons the daemons for a given chainID.
func SummonDaemons(chainID uint64, delay time.Duration) {
	var wg sync.WaitGroup

	time.Sleep(delay)
	go RunMetaVaults(chainID, &wg)
	go RunMetaTokens(chainID, &wg)
	go RunMetaStrategies(chainID, &wg)
	go RunLens(chainID, &wg)
	go RunAPIV1Vaults(chainID, &wg)
	wg.Add(5)
	wg.Wait()
}

// LoadDaemons is a function that loads the previous store state for a given chainID
func LoadDaemons(chainID uint64) {
	var wg sync.WaitGroup

	go LoadMetaVaults(chainID, &wg)
	go LoadMetaTokens(chainID, &wg)
	go LoadMetaStrategies(chainID, &wg)
	go LoadLens(chainID, &wg)
	go LoadAPIV1Vaults(chainID, &wg)
	wg.Add(5)
	wg.Wait()
}
