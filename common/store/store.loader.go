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
** LoadBlockTime will retrieve the all the blockTimes from the the configured DB and store them in
** the _blockTimeSyncMap for fast access during that same execution.
**************************************************************************************************/
func LoadBlockTime(chainID uint64, wg *sync.WaitGroup) {
	if wg != nil {
		defer wg.Done()
	}
	syncMap := _blockTimeSyncMap[chainID]

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
** LoadHistoricalPrice will retrieve the all the historical prices from the configured DB and
** store them in the _historicalPriceSyncMap for fast access during that same execution.
**************************************************************************************************/
func LoadHistoricalPrice(chainID uint64, wg *sync.WaitGroup) {
	if wg != nil {
		defer wg.Done()
	}
	syncMap := _historicalPriceSyncMap[chainID]

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
				for _, v := range temp {
					key := strconv.FormatUint(v.Block, 10) + "_" + addresses.ToAddress(v.Token).Hex()
					syncMap.Store(key, bigNumber.NewInt(0).SetString(v.Price))
				}
				return nil
			})
	}
}

/**************************************************************************************************
** LoadNewVaultsFromRegistry will retrieve the all the vaults added to the registries from the
** configured DB and store them in the _newVaultsFromRegistrySyncMap for fast access during that
**************************************************************************************************/
func LoadNewVaultsFromRegistry(chainID uint64, wg *sync.WaitGroup) {
	if wg != nil {
		defer wg.Done()
	}
	syncMap := _newVaultsFromRegistrySyncMap[chainID]

	switch _dbType {
	case DBBadger:
		// Not implemented
		logs.Warning(`LoadNewVaultsFromRegistry not implemented for DBBadger. Skipping...`)
	case DBMysql:
		var temp []DBNewVaultsFromRegistry

		DATABASE.Table(`db_new_vaults_from_registries`).
			Where("chain_id = ?", chainID).
			FindInBatches(&temp, 10_000, func(tx *gorm.DB, batch int) error {
				for _, v := range temp {
					key := strconv.FormatUint(v.Block, 10) + "_" + addresses.ToAddress(v.RegistryAddress).Hex() + "_" + addresses.ToAddress(v.VaultsAddress).Hex() + "_" + addresses.ToAddress(v.TokenAddress).Hex() + "_" + v.APIVersion
					syncMap.Store(key, v)
				}
				return nil
			})
	}
}
