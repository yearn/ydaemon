package multicalls

import (
	"math"
	"math/big"

	"github.com/yearn/ydaemon/common/ethereum"
)

func Perform(chainID uint64, calls []ethereum.Call, blockNumber *big.Int) map[string][]interface{} {
	caller := ethereum.MulticallClientForChainID[chainID]
	return caller.ExecuteByBatch(calls, math.MaxInt64, blockNumber)
}
