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

var storeRateLimiter = rate.NewLimiter(2, 4)

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
			storeRateLimiter.Wait(context.Background())
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
			storeRateLimiter.Wait(context.Background())
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
	AppendInNewVaultsFromRegistry(chainID, vault)
	key := strconv.FormatUint(vault.BlockNumber, 10) + "_" + vault.RegistryAddress.Hex() + "_" + vault.Address.Hex() + "_" + vault.TokenAddress.Hex() + "_" + vault.APIVersion

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
			storeRateLimiter.Wait(context.Background())
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
** AppendInNewVaultsFromRegistry will add a new vault in the _vaultsSyncMap
**************************************************************************************************/
func AppendInNewVaultsFromRegistry(chainID uint64, vault models.TVaultsFromRegistry) {
	syncMap := _newVaultsFromRegistrySyncMap[chainID]
	key := strconv.FormatUint(vault.BlockNumber, 10) + "_" + vault.RegistryAddress.Hex() + "_" + vault.Address.Hex() + "_" + vault.TokenAddress.Hex() + "_" + vault.APIVersion
	syncMap.Store(key, vault)
}

/**************************************************************************************************
** StoreERC20 will store a new erc20 token in the _erc20SyncMap for fast access during that same
** execution, and will store it in the configured DB for future executions.
**************************************************************************************************/
func StoreERC20(chainID uint64, token models.TERC20Token) {
	AppendInERC20(chainID, token)
	key := token.Address.Hex()

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
			storeRateLimiter.Wait(context.Background())
			DATABASE.
				Table(`db_erc20`).
				FirstOrCreate(newItem)
		}()
	}
}

/**************************************************************************************************
** AppendInERC20 will add a new erc20 token in the _erc20SyncMap
**************************************************************************************************/
func AppendInERC20(chainID uint64, token models.TERC20Token) {
	syncMap := _erc20SyncMap[chainID]
	key := token.Address.Hex()
	syncMap.Store(key, token)
}

/**************************************************************************************************
** StoreVault will store a new vault in the _vaultsSyncMap for fast access during that same
** execution, and will store it in the configured DB for future executions.
**************************************************************************************************/
func StoreVault(chainID uint64, vault models.TVault) {
	AppendInVaultMap(chainID, vault)
	key := vault.Address.Hex() + "_" + vault.Token.Address.Hex() + "_" + strconv.FormatUint(vault.Activation, 10) + "_" + strconv.FormatUint(vault.ChainID, 10)

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
		go func() {
			newItem := &DBVault{}
			newItem.UUID = getUUID(key)
			newItem.Address = vault.Address.Hex()
			newItem.Management = vault.Management.Hex()
			newItem.Governance = vault.Governance.Hex()
			newItem.Guardian = vault.Guardian.Hex()
			newItem.Rewards = vault.Rewards.Hex()
			newItem.Token = vault.Token.Address.Hex()
			newItem.Type = vault.Type
			newItem.Symbol = vault.Symbol
			newItem.DisplaySymbol = vault.DisplaySymbol
			newItem.FormatedSymbol = vault.FormatedSymbol
			newItem.Name = vault.Name
			newItem.DisplayName = vault.DisplayName
			newItem.FormatedName = vault.FormatedName
			newItem.Version = vault.Version
			newItem.ChainID = vault.ChainID
			newItem.Inception = vault.Inception
			newItem.Activation = vault.Activation
			newItem.Decimals = vault.Decimals
			newItem.PerformanceFee = vault.PerformanceFee
			newItem.ManagementFee = vault.ManagementFee
			newItem.Endorsed = vault.Endorsed
			storeRateLimiter.Wait(context.Background())
			DATABASE.
				Table(`db_vaults`).
				FirstOrCreate(newItem)
		}()
	}
}

/**************************************************************************************************
** AppendInVaultMap will add a new vault in the _vaultsSyncMap
**************************************************************************************************/
func AppendInVaultMap(chainID uint64, vault models.TVault) {
	syncMap := _vaultsSyncMap[chainID]
	syncMap.Store(vault.Address, vault)
}

/**************************************************************************************************
** StoreStrategies will store the new strategies in the _strategiesSyncMap for fast access during
** that same execution, and will store it in the configured DB for future executions.
**************************************************************************************************/
func StoreStrategies(chainID uint64, strat models.TStrategyAdded) {
	syncMap := _strategiesSyncMap[chainID]

	key := strat.StrategyAddress.Hex() + "_" + strat.VaultAddress.Hex() + "_" + strat.TxHash.Hex() + "_" + strconv.FormatUint(strat.ChainID, 10)
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
			storeRateLimiter.Wait(context.Background())
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
			storeRateLimiter.Wait(context.Background())
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
	switch _dbType {
	case DBBadger:
		// Not implemented
	case DBSql:

		DATABASE.Table("db_strategy_added_syncs").
			Clauses(clause.OnConflict{
				Columns:   []clause.Column{{Name: "chain_id"}, {Name: "vault"}},
				DoUpdates: clause.Assignments(map[string]interface{}{"block": gorm.Expr("GREATEST(db_strategy_added_syncs.block, EXCLUDED.block)")}),
			}).Create(&itemsToUpsert)
	}
}
