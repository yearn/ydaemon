package multicalls

import (
	"math/big"
	"strconv"
	"time"

	"github.com/yearn/ydaemon/common/env"
	"github.com/yearn/ydaemon/common/ethereum"
	"github.com/yearn/ydaemon/common/logs"
)

func Perform(chainID uint64, calls []ethereum.Call, blockNumber *big.Int) map[string][]interface{} {
	caller := ethereum.MulticallClientForChainID[chainID]
	chain, ok := env.GetChain(chainID)
	if !ok {
		return nil
	}
	
	callCount := len(calls)
	if callCount > 0 {
		logs.Info(`Executing multicall: ` + strconv.Itoa(callCount) + ` calls on chain ` + strconv.FormatUint(chainID, 10))
		start := time.Now()
		
		batchSize := chain.MaxBatchSize
		result := caller.ExecuteByBatch(calls, batchSize, blockNumber)
		
		elapsed := time.Since(start)
		if elapsed > 500*time.Millisecond {
			logs.Info(`Multicall completed in ` + strconv.FormatFloat(elapsed.Seconds(), 'f', 2, 64) + `s`)
		}
		return result
	}
	
	return nil
}
