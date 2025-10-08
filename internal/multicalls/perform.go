package multicalls

import (
	"math/big"
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
		logs.Warning("🧮 [MULTICALL START]", "chain", chainID, "calls", callCount)
		start := time.Now()

		batchSize := chain.MaxBatchSize
		result := caller.ExecuteByBatch(calls, batchSize, blockNumber)

		elapsed := time.Since(start)
		logs.Success("🧮 [MULTICALL DONE]", "chain", chainID, "took", elapsed)
		return result
	}

	return nil
}
