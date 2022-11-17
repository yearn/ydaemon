package ethereum

import (
	"context"
	"math/big"
	"strconv"
	"sync"

	"github.com/yearn/ydaemon/common/logs"
	"github.com/yearn/ydaemon/common/store"
	"github.com/yearn/ydaemon/common/traces"
)

var blockTimeSyncMap = make(map[uint64]*sync.Map)

/**************************************************************************************************
** LoadBlockTime will retrieve the `ChainBlockTime` values from the local DB and store them in the
** blockTimeSyncMap.
**************************************************************************************************/
func LoadBlockTime(chainID uint64, wg *sync.WaitGroup) {
	defer wg.Done()
	temp := make(map[uint64]uint64)
	store.Iterate(chainID, store.TABLES.BLOCK_TIME, &temp)
	if temp != nil && (len(temp) > 0) {
		blockTimeSyncMap[chainID] = &sync.Map{}
		for k, v := range temp {
			blockTimeSyncMap[chainID].Store(k, v)
		}
	}
}

/**************************************************************************************************
** GetBlockTime will try, for a specific blockNumber on a specific chain, to find its execution
** timestamp. This timestamp should be available in the blockTimeSyncMap. If it is not, it will
** fetch it from the blockchain and store it in the blockTimeSyncMap.
**************************************************************************************************/
func GetBlockTime(chainID uint64, blockNumber uint64) (blockTime uint64) {
	syncMap := blockTimeSyncMap[chainID]
	if syncMap == nil {
		syncMap = &sync.Map{}
		blockTimeSyncMap[chainID] = syncMap
	}

	blockTimeData, ok := syncMap.Load(blockNumber)
	if !ok {
		logs.Warning(`not ok`)
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
		syncMap.Store(blockNumber, block.Time)
		go store.SaveInDB(
			chainID,
			store.TABLES.BLOCK_TIME,
			strconv.FormatUint(blockNumber, 10),
			block.Time,
		)
		return block.Time
	}

	return blockTimeData.(uint64)
}
