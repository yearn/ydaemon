package store

import (
	"strconv"
	"sync"

	"github.com/yearn/ydaemon/common/env"
)

var _blockTimeSyncMap = make(map[uint64]*sync.Map)

/**************************************************************************************************
** The init function is a special function triggered directly on execution of the package.
** It is used to initialize the package.
** This init is responsible of loading the store
***************************************************************************************************/
func init() {
	wg := &sync.WaitGroup{}
	for _, chainID := range env.SUPPORTED_CHAIN_IDS {
		wg.Add(1)
		go LoadBlockTime(chainID, wg)
	}
	wg.Wait()
}

/**************************************************************************************************
** LoadBlockTime will retrieve the `ChainBlockTime` values from the local DB and store them in the
** _blockTimeSyncMap.
**************************************************************************************************/
func LoadBlockTime(chainID uint64, wg *sync.WaitGroup) {
	defer wg.Done()
	temp := make(map[uint64]uint64)
	Iterate(chainID, TABLES.BLOCK_TIME, &temp)
	if temp != nil && (len(temp) > 0) {
		for k, v := range temp {
			syncMap := _blockTimeSyncMap[chainID]
			if syncMap == nil {
				syncMap = &sync.Map{}
				_blockTimeSyncMap[chainID] = syncMap
			}
			syncMap.Store(k, v)
		}
	}
}

/**************************************************************************************************
** GetBlockTime will try, for a specific blockNumber on a specific chain, to find its execution
** timestamp in the _blockTimeSyncMap.
**************************************************************************************************/
func GetBlockTime(chainID uint64, blockNumber uint64) (blockTime uint64, ok bool) {
	syncMap := _blockTimeSyncMap[chainID]
	if syncMap == nil {
		syncMap = &sync.Map{}
		_blockTimeSyncMap[chainID] = syncMap
	}

	blockTimeData, ok := syncMap.Load(blockNumber)
	if !ok {
		return 0, false
	}

	return blockTimeData.(uint64), true
}

/**************************************************************************************************
** SaveBlockTime will save a blockTime in the _blockTimeSyncMap.
**************************************************************************************************/
func SaveBlockTime(chainID uint64, blockNumber uint64, blockTime uint64) {
	syncMap := _blockTimeSyncMap[chainID]
	if syncMap == nil {
		syncMap = &sync.Map{}
		_blockTimeSyncMap[chainID] = syncMap
	}
	syncMap.Store(blockNumber, blockTime)
	go SaveInDB(chainID, TABLES.BLOCK_TIME, strconv.FormatUint(blockNumber, 10), blockTime)
}
