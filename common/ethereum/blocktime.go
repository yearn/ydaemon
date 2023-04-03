package ethereum

import (
	"context"
	"math/big"
	"strconv"

	"github.com/yearn/ydaemon/common/store"
	"github.com/yearn/ydaemon/common/traces"
)

/**************************************************************************************************
** GetBlockTime will try, for a specific blockNumber on a specific chain, to find its execution
** timestamp. This timestamp should be available in the blockTimeSyncMap. If it is not, it will
** fetch it from the blockchain and store it in the blockTimeSyncMap.
**************************************************************************************************/
func GetBlockTime(chainID uint64, blockNumber uint64) (blockTime uint64) {
	blockTimeData, ok := store.GetBlockTime(chainID, blockNumber)
	if !ok {
		client := RPC[chainID]
		block, err := client.HeaderByNumber(context.Background(), big.NewInt(int64(blockNumber)))
		if err != nil {
			traces.
				Capture(`warn`, `impossible to retrieve block `+strconv.FormatUint(blockNumber, 10)+` on chain `+strconv.FormatUint(chainID, 10)).
				SetEntity(`block`).
				SetExtra(`error`, err.Error()).
				SetTag(`chainID`, strconv.FormatUint(chainID, 10)).
				SetTag(`rpcURI`, GetRPCURI(chainID)).
				SetTag(`blockNumber`, strconv.FormatUint(blockNumber, 10)).
				Send()
			return 0
		}
		store.StoreBlockTime(chainID, blockNumber, block.Time)
		return block.Time
	}
	return blockTimeData
}
