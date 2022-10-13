package ethereum

import (
	"context"
	"math/big"
	"sync"

	"github.com/yearn/ydaemon/common/logs"
	"github.com/yearn/ydaemon/common/store"
)

var blockTimeSyncMap = make(map[uint64]*sync.Map)

/**************************************************************************************************
** LoadBlockTime will retrieve the `ChainBlockTime` values from the local DB and store them in the
** blockTimeSyncMap.
**************************************************************************************************/
func LoadBlockTime(chainID uint64, wg *sync.WaitGroup) {
	defer wg.Done()
	temp := make(map[uint64]uint64)
	if err := store.LoadFromDBForChainID(store.KEYS.ChainBlockTime, chainID, &temp); err != nil {
		return
	}
	if temp != nil && (len(temp) > 0) {
		blockTimeSyncMap[chainID] = &sync.Map{}
		for k, v := range temp {
			blockTimeSyncMap[chainID].Store(k, v)
		}
	}
}

/**************************************************************************************************
** saveInLocalDB will take everything that is in the blockTimeSyncMap and save it in the local DB
** for a specific chainID. This will allow us to use theses data across reboots.
**************************************************************************************************/
func saveInLocalDB(chainID uint64, blockNumber uint64, timestamp uint64) {
	syncMap := blockTimeSyncMap[chainID]
	dbMap := make(map[uint64]uint64)
	syncMap.Range(func(key, value interface{}) bool {
		dbMap[key.(uint64)] = value.(uint64)
		return true
	})

	store.SaveInDBForChainID(store.KEYS.ChainBlockTime, chainID, dbMap)
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
		block, err := client.BlockByNumber(context.Background(), big.NewInt(int64(blockNumber)))
		if err != nil {
			logs.Error(`impossible to get block time: ` + err.Error())
			return 0
		}
		syncMap.Store(blockNumber, block.Time())
		go saveInLocalDB(chainID, blockNumber, block.Time())
		return block.Time()
	}

	return blockTimeData.(uint64)
}
