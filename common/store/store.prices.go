package store

import (
	"context"
	"strconv"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/common/addresses"
	"github.com/yearn/ydaemon/common/bigNumber"
	"github.com/yearn/ydaemon/common/helpers"
	"github.com/yearn/ydaemon/common/logs"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var _historicalPriceSyncMap = make(map[uint64]*sync.Map)

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
		logs.Warning(`BadgerDB is deprecated for LoadHistoricalPrice`)
	case DBSql:
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
** AppendToHistoricalPrice will add a new time in the _historicalPriceSyncMap
**************************************************************************************************/
func AppendToHistoricalPrice(chainID uint64, blockNumber uint64, tokenAddress common.Address, price *bigNumber.Int) {
	syncMap := _historicalPriceSyncMap[chainID]
	if syncMap == nil {
		syncMap = &sync.Map{}
		_historicalPriceSyncMap[chainID] = syncMap
	}

	key := strconv.FormatUint(blockNumber, 10) + "_" + tokenAddress.Hex()
	syncMap.Store(key, price)
}

/**************************************************************************************************
** StoreHistoricalPrice will store the price of a token at a specific block in the
** _historicalPriceSyncMap for fast access during that same execution, and will store it in the
** configured DB for future executions.
**************************************************************************************************/
func StoreHistoricalPrice(chainID uint64, blockNumber uint64, tokenAddress common.Address, price *bigNumber.Int) {
	AppendToHistoricalPrice(chainID, blockNumber, tokenAddress, price)

	switch _dbType {
	case DBBadger:
		logs.Warning(`BadgerDB is deprecated for StoreHistoricalPrice`)
	case DBSql:
		go func() {
			DBbaseSchema := DBBaseSchema{
				UUID:    getUUID(strconv.FormatUint(chainID, 10) + strconv.FormatUint(blockNumber, 10) + tokenAddress.Hex()),
				Block:   blockNumber,
				ChainID: chainID,
			}
			humanizedPrice, _ := helpers.ToNormalizedAmount(price, 6).Float64()
			wait(`StoreHistoricalPrice`)
			DATABASE.Table(`db_historical_prices`).Save(&DBHistoricalPrice{
				DBbaseSchema,
				tokenAddress.Hex(),
				price.String(),
				humanizedPrice,
			})
		}()
	}
}

/**************************************************************************************************
** StoreManyHistoricalPrice is the same as StoreHistoricalPrice but for many prices at once.
**************************************************************************************************/
func StoreManyHistoricalPrice(items []DBHistoricalPrice) {
	switch _dbType {
	case DBBadger:
		logs.Warning(`BadgerDB is deprecated for StoreManyHistoricalPrice`)
	case DBSql:
		go func() {
			storeRateLimiter.Wait(context.Background())
			DATABASE.
				Table(`db_historical_prices`).
				Clauses(clause.OnConflict{
					UpdateAll: true,
				}).Create(items)
		}()
	}
}

/**************************************************************************************************
** GetBlockPrice will try, for a specific blockNumber on a specific chain, the price of the
** provided token address.
** If the price is missing, this will try to fetch it via the lens oracle contract.
**************************************************************************************************/
func GetBlockPrice(chainID uint64, blockNumber uint64, tokenAddress common.Address) (price *bigNumber.Int, ok bool) {
	key := strconv.FormatUint(blockNumber, 10) + "_" + tokenAddress.Hex()
	tokenPrice, ok := _historicalPriceSyncMap[chainID].Load(key)
	if !ok {
		return bigNumber.NewInt(0), false
	}
	return tokenPrice.(*bigNumber.Int), true
}
