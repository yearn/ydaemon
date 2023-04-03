package store

import (
	"strconv"
	"sync"

	"github.com/yearn/ydaemon/common/addresses"
	"github.com/yearn/ydaemon/common/bigNumber"
	"github.com/yearn/ydaemon/common/logs"
	"gorm.io/gorm"
)

/**************************************************************************************************
** ListBlockTime will retrieve the all the blockTimes from the the configured DB and store them in
** the _blockTimeSyncMap for fast access during that same execution.
**************************************************************************************************/
func ListBlockTime(chainID uint64, wg *sync.WaitGroup) {
	if wg != nil {
		defer wg.Done()
	}
	syncMap := _blockTimeSyncMap[chainID]
	if syncMap == nil {
		syncMap = &sync.Map{}
		_blockTimeSyncMap[chainID] = syncMap
	}

	switch _dbType {
	case DBBadger:
		temp := make(map[uint64]uint64)
		ListFromBadgerDB(chainID, TABLES.BLOCK_TIME, &temp)
		if temp != nil && (len(temp) > 0) {
			for k, v := range temp {
				syncMap.Store(k, v)
			}
		}

	case DBMysql:
		var temp []DBBlockTime
		DATABASE.Table(`db_block_times`).
			Where("chain_id = ?", chainID).
			FindInBatches(&temp, 10_000, func(tx *gorm.DB, batch int) error {
				logs.Info(batch*10_000, `block times loaded from DB`)
				for _, v := range temp {
					syncMap.Store(v.Block, v.Timestamp)
				}
				return nil
			})
	}
}

/**************************************************************************************************
** ListHistoricalPrice will retrieve the all the historical prices from the the configured DB and
** store them in the _historicalPriceSyncMap for fast access during that same execution.
**************************************************************************************************/
func ListHistoricalPrice(chainID uint64, wg *sync.WaitGroup) {
	if wg != nil {
		defer wg.Done()
	}
	syncMap := _historicalPriceSyncMap[chainID]
	if syncMap == nil {
		syncMap = &sync.Map{}
		_historicalPriceSyncMap[chainID] = syncMap
	}

	switch _dbType {
	case DBBadger:
		temp := make(map[string]string)
		ListFromBadgerDB(chainID, TABLES.HISTORICAL_PRICES, &temp)
		if temp != nil && (len(temp) > 0) {
			for k, v := range temp {
				syncMap.Store(k, bigNumber.NewInt(0).SetString(v))
			}
		}
	case DBMysql:
		var temp []DBHistoricalPrice

		DATABASE.Table(`db_historical_prices`).
			Where("chain_id = ?", chainID).
			FindInBatches(&temp, 10_000, func(tx *gorm.DB, batch int) error {
				logs.Info(batch*10_000, `historical prices loaded from DB`)
				for _, v := range temp {
					key := strconv.FormatUint(v.Block, 10) + "_" + addresses.ToAddress(v.Token).Hex()
					syncMap.Store(key, bigNumber.NewInt(0).SetString(v.Price))
				}
				return nil
			})
	}
}
