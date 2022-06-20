package daemons

import (
	"sync"
	"time"
)

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
