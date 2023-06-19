package store

import (
	"context"
	"encoding/json"
	"strconv"
	"strings"
	"sync"

	"github.com/dgraph-io/badger/v3"
	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/common/bigNumber"
	"github.com/yearn/ydaemon/common/helpers"
	"github.com/yearn/ydaemon/common/logs"
	"github.com/yearn/ydaemon/internal/models"
	"golang.org/x/time/rate"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var storeRateLimiter = rate.NewLimiter(4, 8)

func wait(name string) {
	// logs.Warning(`Limiter for ` + name + `: ` + strconv.FormatFloat(storeRateLimiter.Tokens(), 'f', 2, 64))
	storeRateLimiter.Wait(context.Background())
}

func StoreRateLimiter() *rate.Limiter {
	return storeRateLimiter
}

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

	syncMapForTime := _timeBlockSyncMap[chainID]
	if syncMapForTime == nil {
		syncMapForTime = &sync.Map{}
		_blockTimeSyncMap[chainID] = syncMapForTime
	}

	syncMap.Store(blockNumber, blockTime)
	syncMapForTime.Store(blockTime, blockNumber)

	logs.Info(`Storing block time for chain ` + strconv.FormatUint(chainID, 10) + ` block ` + strconv.FormatUint(blockNumber, 10) + ` time ` + strconv.FormatUint(blockTime, 10))
	switch _dbType {
	case DBBadger:
		go OpenBadgerDB(chainID, TABLES.BLOCK_TIME).Update(func(txn *badger.Txn) error {
			dataByte, err := json.Marshal(blockTime)
			if err != nil {
				return err
			}
			return txn.Set([]byte(strconv.FormatUint(blockNumber, 10)), dataByte)
		})
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
		go OpenBadgerDB(chainID, TABLES.HISTORICAL_PRICES).Update(func(txn *badger.Txn) error {
			dataByte, err := json.Marshal(price.String())
			if err != nil {
				return err
			}
			return txn.Set([]byte(key), dataByte)
		})
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
		// Not implemented
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
** AppendInHistoricalMap is the same as StoreHistoricalPrice but only to store
**************************************************************************************************/
func AppendInHistoricalMap(chainID uint64, blockNumber uint64, tokenAddress common.Address, price *bigNumber.Int) {
	syncMap := _historicalPriceSyncMap[chainID]
	key := strconv.FormatUint(blockNumber, 10) + "_" + tokenAddress.Hex()
	syncMap.Store(key, price)
}

/**************************************************************************************************
** StoreNewVaultsFromRegistry will store a new vault in the _newVaultsFromRegistrySyncMap for fast
** access during that same execution, and will store it in the configured DB for future executions.
**************************************************************************************************/
func StoreNewVaultsFromRegistry(chainID uint64, vault models.TVaultsFromRegistry) {
	syncMap := _newVaultsFromRegistrySyncMap[chainID]
	key := strconv.FormatUint(vault.BlockNumber, 10) + "_" + vault.RegistryAddress.Hex() + "_" + vault.Address.Hex() + "_" + vault.TokenAddress.Hex() + "_" + vault.APIVersion
	if data, exists := syncMap.Load(key); exists {
		if Compare(Hash(data), Hash(vault)) {
			return
		}
	}

	syncMap.Store(key, vault)
	switch _dbType {
	case DBBadger:
		// Not implemented
	case DBSql:
		go func() {
			DBbaseSchema := DBBaseSchema{
				UUID:    getUUID(key),
				Block:   vault.BlockNumber,
				ChainID: chainID,
			}

			wait(`StoreNewVaultsFromRegistry`)
			DATABASE.
				Table(`db_new_vaults_from_registries`).
				FirstOrCreate(&DBNewVaultsFromRegistry{
					DBbaseSchema,
					vault.RegistryAddress.Hex(),
					vault.Address.Hex(),
					vault.TokenAddress.Hex(),
					vault.BlockHash.Hex(),
					vault.APIVersion,
					vault.Activation,
					vault.ManagementFee,
					vault.TxIndex,
					vault.LogIndex,
					vault.Type,
				})
		}()
	}
}

/**************************************************************************************************
** StoreERC20 will store a new erc20 token in the _erc20SyncMap for fast access during that same
** execution, and will store it in the configured DB for future executions.
**************************************************************************************************/
func StoreERC20(chainID uint64, token models.TERC20Token) {
	syncMap := _erc20SyncMap[chainID]
	key := token.Address.Hex()
	if data, exists := syncMap.Load(key); exists {
		tokenA := data.(models.TERC20Token)
		tokenB := token
		if tokenA.Address == tokenB.Address && tokenA.Decimals == tokenB.Decimals && tokenA.ChainID == tokenB.ChainID {
			return
		}
	}

	syncMap.Store(key, token)

	switch _dbType {
	case DBBadger:
		go OpenBadgerDB(chainID, TABLES.TOKENS).Update(func(txn *badger.Txn) error {
			dataByte, err := json.Marshal(token)
			if err != nil {
				return err
			}
			return txn.Set([]byte(key), dataByte)
		})
	case DBSql:
		go func() {
			allUnderlyingAsString := []string{}
			for _, address := range token.UnderlyingTokensAddresses {
				allUnderlyingAsString = append(allUnderlyingAsString, address.Hex())
			}
			newItem := &DBERC20{
				UUID:                      getUUID(key),
				Address:                   token.Address.Hex(),
				Name:                      token.Name,
				Symbol:                    token.Symbol,
				Type:                      token.Type,
				DisplayName:               token.DisplayName,
				DisplaySymbol:             token.DisplaySymbol,
				Description:               token.Description,
				Decimals:                  token.Decimals,
				ChainID:                   token.ChainID,
				UnderlyingTokensAddresses: strings.Join(allUnderlyingAsString, ","),
			}
			wait(`StoreERC20`)
			DATABASE.
				Table(`db_erc20`).
				Where(`address = ? AND chain_id = ?`, newItem.Address, newItem.ChainID).
				Assign(newItem).
				FirstOrCreate(newItem)
		}()
	}
}

/**************************************************************************************************
** StoreVault will store a new vault in the _vaultsSyncMap for fast access during that same
** execution, and will store it in the configured DB for future executions.
**************************************************************************************************/
func StoreVault(chainID uint64, vault models.TVault) {
	syncMap := _vaultsSyncMap[chainID]
	key := vault.Address.Hex() + "_" + vault.Token.Address.Hex() + "_" + strconv.FormatUint(vault.Activation, 10) + "_" + strconv.FormatUint(vault.ChainID, 10)

	data, exists := syncMap.Load(vault.Address)
	if exists {
		vaultData := data.(models.TVault)
		newItemFromData := &DBVault{
			UUID:       getUUID(key),
			Address:    vaultData.Address.Hex(),
			Management: vaultData.Management.Hex(),
			Governance: vaultData.Governance.Hex(),
			Guardian:   vaultData.Guardian.Hex(),
			Rewards:    vaultData.Rewards.Hex(),
			Token:      vaultData.Token.Address.Hex(),
			Type:       vaultData.Type,
			Version:    vaultData.Version,
			ChainID:    vaultData.ChainID,
			Inception:  vaultData.Inception,
			Activation: vaultData.Activation,
			Decimals:   vaultData.Decimals,
			Endorsed:   vaultData.Endorsed,
		}
		itemToCompare := &DBVault{
			UUID:       getUUID(key),
			Address:    vault.Address.Hex(),
			Management: vault.Management.Hex(),
			Governance: vault.Governance.Hex(),
			Guardian:   vault.Guardian.Hex(),
			Rewards:    vault.Rewards.Hex(),
			Token:      vault.Token.Address.Hex(),
			Type:       vault.Type,
			Version:    vault.Version,
			ChainID:    vault.ChainID,
			Inception:  vault.Inception,
			Activation: vault.Activation,
			Decimals:   vault.Decimals,
			Endorsed:   vault.Endorsed,
		}
		if Compare(Hash(newItemFromData), Hash(itemToCompare)) {
			return
		}
	}

	switch _dbType {
	case DBBadger:
		go OpenBadgerDB(chainID, TABLES.VAULTS).Update(func(txn *badger.Txn) error {
			dataByte, err := json.Marshal(vault)
			if err != nil {
				return err
			}
			return txn.Set([]byte(key), dataByte)
		})
	case DBSql:
		syncMap.Store(vault.Address, vault)
		func() {
			newItem := &DBVault{
				UUID:           getUUID(key),
				Address:        vault.Address.Hex(),
				Management:     vault.Management.Hex(),
				Governance:     vault.Governance.Hex(),
				Guardian:       vault.Guardian.Hex(),
				Rewards:        vault.Rewards.Hex(),
				Token:          vault.Token.Address.Hex(),
				Type:           vault.Type,
				Symbol:         vault.Symbol,
				DisplaySymbol:  vault.DisplaySymbol,
				FormatedSymbol: vault.FormatedSymbol,
				Name:           vault.Name,
				DisplayName:    vault.DisplayName,
				FormatedName:   vault.FormatedName,
				Version:        vault.Version,
				ChainID:        vault.ChainID,
				Inception:      vault.Inception,
				Activation:     vault.Activation,
				Decimals:       vault.Decimals,
				PerformanceFee: vault.PerformanceFee,
				ManagementFee:  vault.ManagementFee,
				Endorsed:       vault.Endorsed,
			}
			wait(`StoreVault`)
			DATABASE.
				Table(`db_vaults`).
				Where(`address = ? AND chain_id = ?`, newItem.Address, newItem.ChainID).
				Assign(newItem).
				FirstOrCreate(newItem)
		}()
	}
}

/**************************************************************************************************
** StoreStrategies will store the new strategies in the _strategiesSyncMap for fast access during
** that same execution, and will store it in the configured DB for future executions.
**************************************************************************************************/
func StoreStrategies(chainID uint64, strat models.TStrategyAdded) {
	syncMap := _strategiesSyncMap[chainID]
	key := strat.StrategyAddress.Hex() + "_" + strat.VaultAddress.Hex() + "_" + strat.TxHash.Hex() + "_" + strconv.FormatUint(strat.ChainID, 10)
	if data, exists := syncMap.Load(key); exists {
		if Compare(Hash(data), Hash(strat)) {
			return
		}
	}

	syncMap.Store(strat.StrategyAddress, strat)

	switch _dbType {
	case DBBadger:
		// Not implemented
	case DBSql:
		go func() {
			newItem := &DBStrategy{}
			newItem.UUID = getUUID(key)
			newItem.VaultAddress = strat.VaultAddress.Hex()
			newItem.StrategyAddress = strat.StrategyAddress.Hex()
			newItem.TxHash = strat.TxHash.Hex()
			newItem.DebtRatio = strat.DebtRatio.String()
			newItem.MaxDebtPerHarvest = strat.MaxDebtPerHarvest.String()
			newItem.MinDebtPerHarvest = strat.MinDebtPerHarvest.String()
			newItem.PerformanceFee = strat.PerformanceFee.String()
			newItem.DebtLimit = strat.DebtLimit.String()
			newItem.RateLimit = strat.RateLimit.String()
			newItem.VaultVersion = strat.VaultVersion
			newItem.ChainID = strat.ChainID
			newItem.BlockNumber = strat.BlockNumber
			newItem.TxIndex = strat.TxIndex
			newItem.LogIndex = strat.LogIndex
			wait(`StoreStrategies`)
			DATABASE.
				Table(`db_strategies`).
				FirstOrCreate(newItem)
		}()
	}
}

/**************************************************************************************************
** StoreSyncRegistry will store the sync status indicating we went up to the block number to check
** for new vaults.
**************************************************************************************************/
func StoreSyncRegistry(chainID uint64, registryAddess common.Address, end *uint64) {
	switch _dbType {
	case DBBadger:
		// Not implemented
	case DBSql:
		go func() {
			wait(`StoreSyncRegistry`)
			DATABASE.Table("db_registry_syncs").
				Where("chain_id = ? AND registry = ?", chainID, registryAddess.Hex()).
				Where("block <= ?", end).
				Assign(DBRegistrySync{Block: *end}).
				FirstOrCreate(&DBRegistrySync{
					ChainID:  chainID,
					Registry: registryAddess.Hex(),
					UUID:     GetUUID(registryAddess.Hex() + strconv.FormatUint(chainID, 10)),
				})
		}()
	}
}

/**************************************************************************************************
** StoreSyncStrategiesAdded will store the sync status indicating we went up to the block number
** to check for new strategies added.
**************************************************************************************************/
func StoreSyncStrategiesAdded(itemsToUpsert []DBStrategyAddedSync) {
	OpenBadgerDB(
		chainID,
		TABLES.STRATEGIES_FROM_VAULT_SYNC,
	).Update(func(txn *badger.Txn) error {
		if getter, err := txn.Get([]byte(vaultAddress.Hex())); err == nil {
			if previousValue, err := getter.ValueCopy(nil); err == nil {
				previousSync := &DBRegistrySync{}
				if err := json.Unmarshal(previousValue, previousSync); err != nil {
					logs.Error(`StoreSyncStrategiesAdded: json.Unmarshal(previousValue, previousSync)`, err)
					return err
				}
				if previousSync.Block > end {
					logs.Info(`StoreSyncStrategiesAdded: previousSync.Block > end`, previousSync.Block, end)
					return nil
				}
			}
		}

		data := &DBStrategyAddedSync{
			ChainID: chainID,
			Block:   end,
			Vault:   vaultAddress.Hex(),
			UUID:    GetUUID(vaultAddress.Hex() + strconv.FormatUint(chainID, 10)),
		}
		dataByte, err := json.Marshal(data)
		if err != nil {
			logs.Error(`StoreSyncStrategiesAdded: json.Marshal(data)`, err)
			return err
		}
		return txn.Set([]byte(data.Vault), dataByte)
	})
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
