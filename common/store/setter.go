package store

import (
	"encoding/json"
	"strconv"
	"sync"

	"github.com/dgraph-io/badger/v3"
	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/common/bigNumber"
	"github.com/yearn/ydaemon/common/helpers"
	"github.com/yearn/ydaemon/common/logs"
)

/**************************************************************************************************
** StoreBlockTime will store the blockTime in the _blockTimeSyncMap for fast access during that
** same execution, and will store it in the configured DB for future executions.
**************************************************************************************************/
func StoreBlockTime(chainID uint64, blockNumber uint64, blockTime uint64) {
	syncMap := _blockTimeSyncMap[chainID]
	if syncMap == nil {
		syncMap = &sync.Map{}
		_blockTimeSyncMap[chainID] = syncMap
	}
	syncMap.Store(blockNumber, blockTime)

	logs.Info(`Storing block time for chain ` + strconv.FormatUint(chainID, 10) + ` block ` + strconv.FormatUint(blockNumber, 10) + ` time ` + strconv.FormatUint(blockTime, 10))
	switch _dbType {
	case DBBadger:
		go OpenBadgerDB(chainID, TABLES.BLOCK_TIME).Update(func(txn *badger.Txn) error {
			vaultsBytes, err := json.Marshal(blockTime)
			if err != nil {
				return err
			}
			return txn.Set([]byte(strconv.FormatUint(blockNumber, 10)), vaultsBytes)
		})
	case DBMysql:
		DBbaseSchema := DBBaseSchema{
			UUID:    getUUID(strconv.FormatUint(chainID, 10) + strconv.FormatUint(blockNumber, 10) + strconv.FormatUint(blockTime, 10)),
			Block:   blockNumber,
			ChainID: chainID,
		}
		go DATABASE.Table(`db_block_times`).Save(&DBBlockTime{DBbaseSchema, blockTime})
	}
}

/**************************************************************************************************
** StoreHistoricalPrice will store the price of a token at a specific block in the
** _historicalPriceSyncMap for fast access during that same execution, and will store it in the
** configured DB for future executions.
**************************************************************************************************/
func StoreHistoricalPrice(chainID uint64, blockNumber uint64, tokenAddress common.Address, price *bigNumber.Int) {
	syncMap := _historicalPriceSyncMap[chainID]
	key := strconv.FormatUint(blockNumber, 10) + "_" + tokenAddress.Hex()
	syncMap.Store(key, price)

	switch _dbType {
	case DBBadger:
		OpenBadgerDB(chainID, TABLES.HISTORICAL_PRICES).Update(func(txn *badger.Txn) error {
			vaultsBytes, err := json.Marshal(price.String())
			if err != nil {
				return err
			}
			return txn.Set([]byte(key), vaultsBytes)
		})
	case DBMysql:
		DBbaseSchema := DBBaseSchema{
			UUID:    getUUID(strconv.FormatUint(chainID, 10) + strconv.FormatUint(blockNumber, 10) + tokenAddress.Hex()),
			Block:   blockNumber,
			ChainID: chainID,
		}
		humanizedPrice, _ := helpers.ToNormalizedAmount(price, 6).Float64()
		DATABASE.Table(`db_historical_prices`).Save(&DBHistoricalPrice{
			DBbaseSchema,
			tokenAddress.Hex(),
			price.String(),
			humanizedPrice,
		})
	}
}
