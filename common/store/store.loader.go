package store

import (
	"strconv"
	"strings"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/common/addresses"
	"github.com/yearn/ydaemon/common/bigNumber"
	"github.com/yearn/ydaemon/common/env"
	"github.com/yearn/ydaemon/common/logs"
	"github.com/yearn/ydaemon/internal/models"
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

	case DBSql:
		var temp []DBBlockTime
		DATABASE.Table(`db_block_times`).
			Where("chain_id = ?", chainID).
			FindInBatches(&temp, 10_000, func(tx *gorm.DB, batch int) error {
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
	case DBSql:
		var temp []DBNewVaultsFromRegistry

		DATABASE.Table(`db_new_vaults_from_registries`).
			Where("chain_id = ?", chainID).
			FindInBatches(&temp, 10_000, func(tx *gorm.DB, batch int) error {
				for _, vaultFromDB := range temp {
					key := strconv.FormatUint(vaultFromDB.Block, 10) + "_" + addresses.ToAddress(vaultFromDB.RegistryAddress).Hex() + "_" + addresses.ToAddress(vaultFromDB.VaultsAddress).Hex() + "_" + addresses.ToAddress(vaultFromDB.TokenAddress).Hex() + "_" + vaultFromDB.APIVersion
					vaultFromRegistry := models.TVaultsFromRegistry{
						Address:         addresses.ToAddress(vaultFromDB.VaultsAddress),
						RegistryAddress: addresses.ToAddress(vaultFromDB.RegistryAddress),
						TokenAddress:    addresses.ToAddress(vaultFromDB.TokenAddress),
						BlockHash:       common.HexToHash(vaultFromDB.BlockHash),
						Type:            vaultFromDB.Type,
						APIVersion:      vaultFromDB.APIVersion,
						ChainID:         vaultFromDB.ChainID,
						BlockNumber:     vaultFromDB.Block,
						Activation:      vaultFromDB.Activation,
						ManagementFee:   vaultFromDB.ManagementFee,
						TxIndex:         vaultFromDB.TxIndex,
						LogIndex:        vaultFromDB.LogIndex,
					}
					syncMap.Store(key, vaultFromRegistry)
				}
				return nil
			})
	}
}

/**************************************************************************************************
** LoadERC20 will retrieve the all the ERC20 tokens added to the configured DB and store them in
** the _erc20SyncMap for fast access during that
**************************************************************************************************/
func LoadERC20(chainID uint64, wg *sync.WaitGroup) {
	if wg != nil {
		defer wg.Done()
	}
	syncMap := _erc20SyncMap[chainID]

	switch _dbType {
	case DBBadger:
		temp := make(map[common.Address]models.TERC20Token)
		ListFromBadgerDB(chainID, TABLES.TOKENS, &temp)
		if temp != nil && (len(temp) > 0) {
			for _, v := range temp {
				syncMap.Store(v.Address, v)
			}
		}
	case DBSql:
		var temp []DBERC20
		DATABASE.Table(`db_erc20`).
			Where("chain_id = ?", chainID).
			FindInBatches(&temp, 10_000, func(tx *gorm.DB, batch int) error {
				for _, tokenFromDB := range temp {
					token := models.TERC20Token{
						Address:       addresses.ToAddress(tokenFromDB.Address),
						Type:          tokenFromDB.Type,
						Name:          tokenFromDB.Name,
						Symbol:        tokenFromDB.Symbol,
						DisplayName:   tokenFromDB.DisplayName,
						DisplaySymbol: tokenFromDB.DisplaySymbol,
						Description:   tokenFromDB.Description,
						Icon:          env.BASE_ASSET_URL + strconv.FormatUint(chainID, 10) + `/` + tokenFromDB.Address + `/logo-128.png`,
						Decimals:      tokenFromDB.Decimals,
						ChainID:       tokenFromDB.ChainID,
					}
					allUnderlyingAsString := strings.Split(tokenFromDB.UnderlyingTokensAddresses, ",")
					for _, addressStr := range allUnderlyingAsString {
						token.UnderlyingTokensAddresses = append(token.UnderlyingTokensAddresses, common.HexToAddress(addressStr))
					}
					syncMap.Store(tokenFromDB.Address, token)
				}
				return nil
			})
	}
}

/**************************************************************************************************
** LoadVaults will retrieve the all the vaults from the configured DB and store them in the
** _vaultsSyncMap for fast access during that same execution.
**************************************************************************************************/
func LoadVaults(chainID uint64, wg *sync.WaitGroup) {
	if wg != nil {
		defer wg.Done()
	}
	syncMap := _vaultsSyncMap[chainID]

	switch _dbType {
	case DBBadger:
		temp := make(map[common.Address]models.TVault)
		ListFromBadgerDB(chainID, TABLES.VAULTS, &temp)
		if temp != nil && (len(temp) > 0) {
			for _, v := range temp {
				syncMap.Store(v.Address, v)
			}
		}
	case DBSql:
		var temp []DBVault
		DATABASE.Table(`db_vaults`).
			Where("chain_id = ?", chainID).
			FindInBatches(&temp, 10_000, func(tx *gorm.DB, batch int) error {
				for _, vaultFromDB := range temp {
					vault := models.TVault{
						Address:        common.HexToAddress(vaultFromDB.Address),
						Management:     common.HexToAddress(vaultFromDB.Management),
						Governance:     common.HexToAddress(vaultFromDB.Governance),
						Guardian:       common.HexToAddress(vaultFromDB.Guardian),
						Rewards:        common.HexToAddress(vaultFromDB.Rewards),
						Type:           vaultFromDB.Type,
						Symbol:         vaultFromDB.Symbol,
						DisplaySymbol:  vaultFromDB.DisplaySymbol,
						FormatedSymbol: vaultFromDB.FormatedSymbol,
						Name:           vaultFromDB.Name,
						DisplayName:    vaultFromDB.DisplayName,
						FormatedName:   vaultFromDB.FormatedName,
						Icon:           env.BASE_ASSET_URL + strconv.FormatUint(chainID, 10) + `/` + vaultFromDB.Address + `/logo-128.png`,
						Version:        vaultFromDB.Version,
						ChainID:        vaultFromDB.ChainID,
						Inception:      vaultFromDB.Inception,
						Activation:     vaultFromDB.Activation,
						Decimals:       vaultFromDB.Decimals,
						PerformanceFee: vaultFromDB.PerformanceFee,
						ManagementFee:  vaultFromDB.ManagementFee,
						Endorsed:       vaultFromDB.Endorsed,
					}

					/******************************************************************************
					** As the vaults stored in the DB do not store the underlying token information
					** we need to retrieve it from the _erc20SyncMap and add it to our structure.
					******************************************************************************/
					vault.Token = models.TERC20Token{
						Address: common.HexToAddress(vaultFromDB.Token),
						ChainID: vaultFromDB.ChainID,
					}
					if token, ok := GetERC20(chainID, common.HexToAddress(vaultFromDB.Token)); ok {
						vault.Token.Decimals = token.Decimals
						vault.Token.Symbol = token.Symbol
						vault.Token.Name = token.Name
						vault.Token.Icon = env.BASE_ASSET_URL + strconv.FormatUint(chainID, 10) + `/` + vault.Token.Address.Hex() + `/logo-128.png`
						vault.Token.DisplaySymbol = token.DisplaySymbol
						vault.Token.DisplayName = token.DisplayName
						vault.Token.Description = token.Description
						vault.Token.Type = token.Type
						vault.Token.UnderlyingTokensAddresses = token.UnderlyingTokensAddresses
					}

					/******************************************************************************
					** As the vaults stored in the DB do not store mutable data we will need to
					** retrieve them on the fly from the blockchain.
					** Here are the missing fields:
					******************************************************************************/
					vault.WithdrawalQueue = []common.Address{}
					vault.PricePerShare = bigNumber.NewInt(0)
					vault.DepositLimit = bigNumber.NewInt(0)
					vault.AvailableDepositLimit = bigNumber.NewInt(0)
					vault.TotalAssets = bigNumber.NewInt(0)
					syncMap.Store(vault.Address, vault)
				}
				return nil
			})
	}
}

/**************************************************************************************************
** LoadStrategies will retrieve the all the strategies from the configured DB and store them in the
** _strategiesSyncMap for fast access during that same execution.
**************************************************************************************************/
func LoadStrategies(chainID uint64, wg *sync.WaitGroup) {
	if wg != nil {
		defer wg.Done()
	}
	syncMap := _strategiesSyncMap[chainID]

	switch _dbType {
	case DBBadger:
		// not implemented
	case DBSql:
		var temp []DBStrategy
		DATABASE.Table(`db_strategies`).
			Where("chain_id = ?", chainID).
			FindInBatches(&temp, 10_000, func(tx *gorm.DB, batch int) error {
				for _, stratFromDB := range temp {
					strat := models.TStrategyAdded{
						VaultAddress:      common.HexToAddress(stratFromDB.VaultAddress),
						StrategyAddress:   common.HexToAddress(stratFromDB.StrategyAddress),
						TxHash:            common.HexToHash(stratFromDB.TxHash),
						DebtRatio:         bigNumber.NewInt(0).SetString(stratFromDB.DebtRatio),
						MaxDebtPerHarvest: bigNumber.NewInt(0).SetString(stratFromDB.MaxDebtPerHarvest),
						MinDebtPerHarvest: bigNumber.NewInt(0).SetString(stratFromDB.MinDebtPerHarvest),
						PerformanceFee:    bigNumber.NewInt(0).SetString(stratFromDB.PerformanceFee),
						DebtLimit:         bigNumber.NewInt(0).SetString(stratFromDB.DebtLimit),
						RateLimit:         bigNumber.NewInt(0).SetString(stratFromDB.RateLimit),
						VaultVersion:      stratFromDB.VaultVersion,
						ChainID:           stratFromDB.ChainID,
						BlockNumber:       stratFromDB.BlockNumber,
						TxIndex:           stratFromDB.TxIndex,
						LogIndex:          stratFromDB.LogIndex,
					}
					syncMap.Store(strat.StrategyAddress, strat)
				}
				return nil
			})
	}
}
