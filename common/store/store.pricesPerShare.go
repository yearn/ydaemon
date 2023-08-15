package store

import (
	"strconv"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/common/bigNumber"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var _vaultsPricePerShareSyncMap = make(map[uint64]*sync.Map)

/**************************************************************************************************
** LoadPricePerShare will retrieve the all the pricePerShare from the configured DB and store them
** in the _vaultsPricePerShareSyncMap for fast access during that same execution.
**************************************************************************************************/
func LoadPricePerShare(chainID uint64, wg *sync.WaitGroup) {
	if wg != nil {
		defer wg.Done()
	}
	syncMap := _vaultsPricePerShareSyncMap[chainID]
	if syncMap == nil {
		syncMap = &sync.Map{}
		_vaultsPricePerShareSyncMap[chainID] = syncMap
	}

	switch _dbType {
	case DBBadger:
		// not implemented
	case DBSql:
		now := time.Now().Unix()
		var temp []DBVaultPricePerShare
		DATABASE.Table(`db_vault_price_per_shares`).
			Where("chain_id = ?", chainID).
			Where("time > ?", now-86400*31).
			FindInBatches(&temp, 10_000, func(tx *gorm.DB, batch int) error {
				for _, dataFromDB := range temp {
					key := common.HexToAddress(dataFromDB.Vault).Hex() + `_` + strconv.FormatUint(dataFromDB.Time, 10)
					syncMap.Store(key, dataFromDB)
				}
				return nil
			})
	}
}

/**************************************************************************************************
** StorePricePerShare will store the pricePerShare information for a vault at a specific time/block
** in _vaultsPricePerShareSyncMap for fast access during that same execution, and will store it in
** the configured DB for future executions.
**************************************************************************************************/
func StorePricePerShare(chainID uint64, data []DBVaultPricePerShare) {
	switch _dbType {
	case DBBadger:
		// Not implemented
	case DBSql:
		go func() {
			items := make([]DBVaultPricePerShare, len(data))
			syncMap := _vaultsPricePerShareSyncMap[chainID]
			if syncMap == nil {
				syncMap = &sync.Map{}
				_vaultsPricePerShareSyncMap[chainID] = syncMap
			}

			for _, d := range data {
				key := (common.HexToAddress(d.Vault).Hex() +
					`_` +
					strconv.FormatUint(d.Time, 10) +
					`_` +
					strconv.FormatUint(d.ChainID, 10) +
					`_` +
					strconv.FormatUint(d.Block, 10))
				syncMap.Store(key, d)
				items = append(items, DBVaultPricePerShare{
					UUID:                   getUUID(key),
					Vault:                  d.Vault,
					PricePerShare:          d.PricePerShare,
					HumanizedPricePerShare: d.HumanizedPricePerShare,
					Time:                   d.Time,
					ChainID:                d.ChainID,
					Block:                  d.Block,
				})
			}
			wait(`StorePricePerShare`)
			DATABASE.
				Table(`db_vault_price_per_shares`).
				Clauses(clause.OnConflict{UpdateAll: true}).
				Create(&items)
		}()
	}
}

/**************************************************************************************************
** ListPricePerShare will try, for a specific chain and vault, to retrieve all the price per share
** stored in the cache.
**************************************************************************************************/
func ListPricePerShare(chainID uint64, vaultAddress common.Address) (
	withTime map[uint64]*bigNumber.Int,
	withBlock map[uint64]*bigNumber.Int,
) {
	withTime = make(map[uint64]*bigNumber.Int)
	withBlock = make(map[uint64]*bigNumber.Int)
	syncMap := _vaultsPricePerShareSyncMap[chainID]
	if syncMap == nil {
		syncMap = &sync.Map{}
		_vaultsPricePerShareSyncMap[chainID] = syncMap
	}

	syncMap.Range(func(key, value interface{}) bool {
		data := value.(DBVaultPricePerShare)
		if data.Vault != vaultAddress.Hex() {
			return true
		}
		withTime[data.Time] = bigNumber.NewInt(0).SetString(data.PricePerShare)
		withBlock[data.Block] = bigNumber.NewInt(0).SetString(data.PricePerShare)
		return true
	})

	return withTime, withBlock
}

/**************************************************************************************************
** GetPricePerShare will try, for a specific chain, vault and time, to find the price per share.
**************************************************************************************************/
func GetPricePerShare(chainID uint64, vaultAddress common.Address, time uint64) (pps string, ok bool) {
	key := vaultAddress.Hex() + `_` + strconv.FormatUint(time, 10)
	syncMapResult, ok := _vaultsPricePerShareSyncMap[chainID].Load(key)
	if !ok {
		return ``, false
	}
	data := syncMapResult.(DBVaultPricePerShare)
	return data.PricePerShare, true
}
