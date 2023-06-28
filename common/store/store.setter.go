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
		// LEGACY: Deprecated
		logs.Warning(`BadgerDB is deprecated for StoreBlockTime`)
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
		// LEGACY: Deprecated
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
		// LEGACY: Deprecated
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
** AppendInHistoricalMap is the same as StoreHistoricalPrice but only to store
**************************************************************************************************/
func AppendInHistoricalMap(chainID uint64, blockNumber uint64, tokenAddress common.Address, price *bigNumber.Int) {
	syncMap := _historicalPriceSyncMap[chainID]
	key := strconv.FormatUint(blockNumber, 10) + "_" + tokenAddress.Hex()
	syncMap.Store(key, price)
}

/**************************************************************************************************
** StoreNewVaultsFromRegistry will store a new vault in the _newVaultsFromRegistrySyncMap for fast
** access during that same execution, and will store it in the local DB for future executions.
** We are using the local DB because we don't want to trust the shared DB for this data.
**************************************************************************************************/
func StoreNewVaultsFromRegistry(chainID uint64, vault models.TVaultsFromRegistry) {
	AppendInNewVaultsFromRegistry(chainID, vault)
	key := strconv.FormatUint(vault.BlockNumber, 10) + "_" + vault.RegistryAddress.Hex() + "_" + vault.Address.Hex() + "_" + vault.TokenAddress.Hex() + "_" + vault.APIVersion

	go OpenBadgerDB(chainID, TABLES.VAULTS_FROM_REGISTRY_SYNC).Update(func(txn *badger.Txn) error {
		DBbaseSchema := DBBaseSchema{
			UUID:    getUUID(key),
			Block:   vault.BlockNumber,
			ChainID: chainID,
		}
		data := &DBNewVaultsFromRegistry{
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
		}
		dataByte, err := json.Marshal(data)
		if err != nil {
			return err
		}
		return txn.Set([]byte(key), dataByte)
	})
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
		// LEGACY: Deprecated
		logs.Warning(`BadgerDB is deprecated for StoreERC20`)
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
		// LEGACY: Deprecated
		logs.Warning(`BadgerDB is deprecated for StoreVault`)
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
	if data, exists := syncMap.Load(key); exists {
		if Compare(Hash(data), Hash(strat)) {
			return
		}
	}

	syncMap.Store(strat.StrategyAddress, strat)

	switch _dbType {
	case DBBadger:
		// LEGACY: Deprecated
		logs.Warning(`BadgerDB is deprecated for StoreStrategies`)
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
** AppendInStrategyMap will add a new vault in the _strategiesSyncMap
**************************************************************************************************/
func AppendInStrategyMap(chainID uint64, strat models.TStrategyAdded) {
	syncMap := _strategiesSyncMap[chainID]
	syncMap.Store(strat.StrategyAddress, strat)
}

/**************************************************************************************************
** StoreSyncRegistry will store the sync status indicating we went up to the block number to check
** for new vaults.
**************************************************************************************************/
func StoreSyncRegistry(chainID uint64, registryAddess common.Address, end *uint64) {
	OpenBadgerDB(chainID, TABLES.REGISTRY_SYNC).Update(func(txn *badger.Txn) error {
		if getter, err := txn.Get([]byte(registryAddess.Hex())); err == nil {
			if previousValue, err := getter.ValueCopy(nil); err == nil {
				previousSync := &DBRegistrySync{}
				if err := json.Unmarshal(previousValue, previousSync); err != nil {
					logs.Error(`StoreSyncRegistry: json.Unmarshal(previousValue, previousSync)`, err)
					return err
				}
				if previousSync.Block > *end {
					logs.Info(`StoreSyncRegistry: previousSync.Block > *end`, previousSync.Block, *end)
					return nil
				}
			}
		}

		data := &DBRegistrySync{
			ChainID:  chainID,
			Block:    *end,
			Registry: registryAddess.Hex(),
			UUID:     GetUUID(registryAddess.Hex() + strconv.FormatUint(chainID, 10)),
		}
		dataByte, err := json.Marshal(data)
		if err != nil {
			return err
		}
		return txn.Set([]byte(data.Registry), dataByte)
	})
}

/**************************************************************************************************
** StoreSyncStrategiesAdded will store the sync status indicating we went up to the block number
** to check for new strategies added.
**************************************************************************************************/
func StoreSyncStrategiesAdded(chainID uint64, vaultAddress common.Address, end uint64) {
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
