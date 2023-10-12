package store

import (
	"strconv"
	"sync"

	"github.com/yearn/ydaemon/common/logs"
	"gorm.io/gorm"
)

var _blockTimeSyncMap = make(map[uint64]*sync.Map)
var _timeBlockSyncMap = make(map[uint64]*sync.Map)

/**************************************************************************************************
** LoadBlockTime will retrieve the all the blockTimes from the the configured DB and store them in
** the _blockTimeSyncMap for fast access during that same execution.
**************************************************************************************************/
func LoadBlockTime(chainID uint64, wg *sync.WaitGroup) {
	if wg != nil {
		defer wg.Done()
	}
	syncMap := _blockTimeSyncMap[chainID]
	if syncMap == nil {
		syncMap = &sync.Map{}
		_blockTimeSyncMap[chainID] = syncMap
	}

	syncMapForTime := _timeBlockSyncMap[chainID]
	if syncMap == nil {
		syncMapForTime = &sync.Map{}
		_timeBlockSyncMap[chainID] = syncMapForTime
	}

	switch _dbType {
	case DBSql:
		var temp []DBBlockTime
		DATABASE.Table(`db_block_times`).
			Where("chain_id = ?", chainID).
			FindInBatches(&temp, 10_000, func(tx *gorm.DB, batch int) error {
				for _, v := range temp {
					syncMap.Store(v.Block, v.Timestamp)
					syncMapForTime.Store(v.Timestamp, v.Block)
				}
				return nil
			})
	}
}

/**************************************************************************************************
** AppendToBlockTime will add a new time in the _blockTimeSyncMap
**************************************************************************************************/
func AppendToBlockTime(chainID uint64, blockNumber uint64, blockTime uint64) {
	syncMap := _blockTimeSyncMap[chainID]
	if syncMap == nil {
		syncMap = &sync.Map{}
		_blockTimeSyncMap[chainID] = syncMap
	}

	syncMapForTime := _timeBlockSyncMap[chainID]
	if syncMapForTime == nil {
		syncMapForTime = &sync.Map{}
		_blockTimeSyncMap[chainID] = syncMapForTime
	}

	syncMap.Store(blockNumber, blockTime)
	syncMapForTime.Store(blockTime, blockNumber)
}

/**************************************************************************************************
** StoreBlockTime will store the blockTime in the _blockTimeSyncMap for fast access during that
** same execution, and will store it in the configured DB for future executions.
**************************************************************************************************/
func StoreBlockTime(chainID uint64, blockNumber uint64, blockTime uint64) {
	AppendToBlockTime(chainID, blockNumber, blockTime)

	logs.Info(`Storing block time for chain ` + strconv.FormatUint(chainID, 10) + ` block ` + strconv.FormatUint(blockNumber, 10) + ` time ` + strconv.FormatUint(blockTime, 10))

	switch _dbType {
	case DBSql:
		go func() {
			DBbaseSchema := DBBaseSchema{
				UUID:    getUUID(strconv.FormatUint(chainID, 10) + strconv.FormatUint(blockNumber, 10) + strconv.FormatUint(blockTime, 10)),
				Block:   blockNumber,
				ChainID: chainID,
			}
			wait(`StoreBlockTime`)
			DATABASE.Table(`db_block_times`).Save(&DBBlockTime{DBbaseSchema, blockTime})
		}()
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
** GetTimeBlock will try, for a specific time on a specific chain, to find its execution
** blockNumber in the _timeBlockSyncMap.
**************************************************************************************************/
func GetTimeBlock(chainID uint64, timestamp uint64) (blockNumber uint64, ok bool) {
	syncMap := _timeBlockSyncMap[chainID]
	if syncMap == nil {
		syncMap = &sync.Map{}
		_timeBlockSyncMap[chainID] = syncMap
	}

	blockTimeData, ok := syncMap.Load(timestamp)
	if !ok {
		return 0, false
	}

	return blockTimeData.(uint64), true
}
