package store

import (
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/common/bigNumber"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var _vaultsActivations = make(map[uint64]*sync.Map)

/**************************************************************************************************
** LoadVaultsActivation will retrieve the all the activation of a vault from the configured DB and
** store them in the _vaultsActivations for fast access during that same execution.
**************************************************************************************************/
func LoadVaultsActivation(chainID uint64, wg *sync.WaitGroup) {
	if wg != nil {
		defer wg.Done()
	}
	syncMap := _vaultsActivations[chainID]
	if syncMap == nil {
		syncMap = &sync.Map{}
		_vaultsActivations[chainID] = syncMap
	}

	switch _dbType {
	case DBBadger:
		// not implemented
	case DBSql:
		var temp []DBVaultActivation
		DATABASE.Table(`db_vault_activations`).
			Where("chain_id = ?", chainID).
			FindInBatches(&temp, 10_000, func(tx *gorm.DB, batch int) error {
				for _, dataFromDB := range temp {
					key := common.HexToAddress(dataFromDB.Vault).Hex()
					syncMap.Store(key, dataFromDB)
				}
				return nil
			})
	}
}

/**************************************************************************************************
** StoreVaultActivation will store the vault activation in _vaultsActivations for fast access
** during that same execution, and will store it in the configured DB for future executions.
**************************************************************************************************/
func StoreVaultActivation(chainID uint64, data []DBVaultActivation) {
	switch _dbType {
	case DBBadger:
		// Not implemented
	case DBSql:
		go func() {
			items := make([]DBVaultActivation, len(data))
			syncMap := _vaultsActivations[chainID]
			if syncMap == nil {
				syncMap = &sync.Map{}
				_vaultsActivations[chainID] = syncMap
			}

			for _, d := range data {
				key := common.HexToAddress(d.Vault).Hex()
				syncMap.Store(key, d)
				items = append(items, DBVaultActivation{
					UUID:    getUUID(key),
					Vault:   d.Vault,
					ChainID: d.ChainID,
					Block:   d.Block,
				})
			}
			wait(`StoreVaultActivation`)
			DATABASE.
				Table(`db_vault_activations`).
				Clauses(clause.OnConflict{UpdateAll: true}).
				Create(&items)
		}()
	}
}

/**************************************************************************************************
** GetVaultActivation will try, for a specific chain and vault to find the activation block.
**************************************************************************************************/
func GetVaultActivation(chainID uint64, vaultAddress common.Address, time uint64) (block uint64, ok bool) {
	key := vaultAddress.Hex()
	syncMapResult, ok := _vaultsActivations[chainID].Load(key)
	if !ok {
		return 0, false
	}
	data := syncMapResult.(DBVaultActivation)
	return data.Block, true
}

/**************************************************************************************************
** ListPricePerShare will try, for a specific chain and vault, to retrieve all the price per share
** stored in the cache.
**************************************************************************************************/
func ListVaultActivation(chainID uint64) (
	vaultBlock map[common.Address]*bigNumber.Int,
) {
	vaultBlock = make(map[common.Address]*bigNumber.Int)
	syncMap := _vaultsActivations[chainID]
	if syncMap == nil {
		syncMap = &sync.Map{}
		_vaultsActivations[chainID] = syncMap
	}
	syncMap.Range(func(key, value interface{}) bool {
		data := value.(DBVaultActivation)
		vaultBlock[common.HexToAddress(data.Vault)] = bigNumber.NewInt(0).SetUint64(data.Block)
		return true
	})

	return vaultBlock
}
