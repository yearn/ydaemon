package multicalls

import (
	"math/big"

	"github.com/yearn/ydaemon/common/env"
	"github.com/yearn/ydaemon/common/ethereum"
)

func Perform(chainID uint64, calls []ethereum.Call, blockNumber *big.Int) map[string][]interface{} {
	caller := ethereum.MulticallClientForChainID[chainID]
	chain, ok := env.GetChain(chainID)
	if !ok {
		return nil
	}
	batchSize := chain.MaxBatchSize
	return caller.ExecuteByBatch(calls, batchSize, blockNumber)
}
