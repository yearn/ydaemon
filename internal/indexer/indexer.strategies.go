package indexer

import (
	"context"
	"math/big"
	"strconv"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/common/addresses"
	"github.com/yearn/ydaemon/common/contracts"
	"github.com/yearn/ydaemon/common/env"
	"github.com/yearn/ydaemon/common/ethereum"
	"github.com/yearn/ydaemon/common/logs"
	"github.com/yearn/ydaemon/internal/fetcher"
	"github.com/yearn/ydaemon/internal/models"
	"github.com/yearn/ydaemon/internal/storage"
)

// Track last indexed block per vault to avoid re-processing
var _lastIndexedBlock = make(map[uint64]map[common.Address]uint64)

/**************************************************************************************************
** listStrategiesForVault reads current strategies directly from the vault contract
**************************************************************************************************/
func listStrategiesForVault(
	chainID uint64,
	vault models.TVault,
) []models.TStrategy {
	client := ethereum.GetRPC(chainID)

	strategies := []models.TStrategy{}
	switch vault.Version {
	case `0.2.2`, `0.3.0`, `0.3.1`, `0.3.2`, `0.3.3`, `0.3.4`, `0.3.5`, `0.4.2`, `0.4.3`:
		/******************************************************************************************
		** Vaults versions from 0.2.2 to 0.4.3 are now deprecated and should not be used anymore.
		** This is only for historical purposes.
		** The strategies indexer will not run for this version.
		******************************************************************************************/
		break // UNBREAK THIS TO RUN THE INDEXER FOR THESE VERSIONS
		// currentVault, _ := contracts.NewYvault043(vault.Address, client)
		// for i := 0; i < 10; i++ {
		// 	indexedStrategy, err := currentVault.WithdrawalQueue(nil, big.NewInt(int64(i)))
		// 	if addresses.Equals(indexedStrategy, common.Address{}) {
		// 		break
		// 	}
		// 	if err != nil {
		// 		continue
		// 	}
		// 	strategies = append(strategies, models.TStrategy{
		// 		Address:      indexedStrategy,
		// 		ChainID:      chainID,
		// 		VaultVersion: vault.Version,
		// 		VaultAddress: vault.Address,
		// 		Activation:   vault.Activation,
		// 	})
		// }
	case `0.4.4`, `0.4.5`, `0.4.6`, `0.4.7`:
		currentVault, _ := contracts.NewYvault043(vault.Address, client)
		for i := 0; i < 10; i++ {
			indexedStrategy, err := currentVault.WithdrawalQueue(nil, big.NewInt(int64(i)))
			if addresses.Equals(indexedStrategy, common.Address{}) {
				break
			}
			if err != nil {
				continue
			}
			strategies = append(strategies, models.TStrategy{
				Address:      indexedStrategy,
				ChainID:      chainID,
				VaultVersion: vault.Version,
				VaultAddress: vault.Address,
				Activation:   vault.Activation,
			})
		}
	default:
		// case `3.0.0`, `3.0.1`, `3.0.2`:
		currentVault, _ := contracts.NewYvault300(vault.Address, client)
		for i := 0; i < 10; i++ {
			indexedStrategy, err := currentVault.DefaultQueue(nil, big.NewInt(int64(i)))
			if addresses.Equals(indexedStrategy, common.Address{}) {
				break
			}
			if err != nil {
				continue
			}
			strategies = append(strategies, models.TStrategy{
				Address:      indexedStrategy,
				ChainID:      chainID,
				VaultVersion: vault.Version,
				VaultAddress: vault.Address,
				Activation:   vault.Activation,
			})
		}
	}
	return strategies
}

/**************************************************************************************************
** filterNewStrategies queries blockchain events for new strategies from start to end block
**************************************************************************************************/
func filterNewStrategies(
	chainID uint64,
	vault models.TVault,
	start uint64,
	end uint64,
) {
	/**********************************************************************************************
	** As we cannot use the WS connection, we will fallback to the regular RPC connection. This
	** means that we will need to filter the logs from the lastSyncedBlock to the latest block
	** every x seconds to check if there are new strategies.
	**********************************************************************************************/
	client := ethereum.GetRPC(chainID)
	chain, ok := env.GetChain(chainID)
	if !ok {
		return
	}

	for chunkStart := start; chunkStart < end; chunkStart += chain.MaxBlockRange {
		chunkEnd := chunkStart + chain.MaxBlockRange
		if chunkEnd > end {
			chunkEnd = end
		}
		opts := &bind.FilterOpts{Start: chunkStart, End: &chunkEnd}

		switch vault.Version {
		case `0.2.2`:
			currentVault, _ := contracts.NewYvault022(vault.Address, client)
			if log, err := currentVault.FilterStrategyAdded(opts, nil); err == nil {
				for log.Next() {
					if log.Error() != nil {
						continue
					}
					newStrategy := handleV02Strategies(chainID, vault.Version, log.Event)
					if storage.StoreStrategyIfMissing(chainID, newStrategy) {
						strategyKey := newStrategy.Address.Hex() + `_` + newStrategy.VaultAddress.Hex()
						fetcher.RetrieveAllStrategies(chainID, map[string]models.TStrategy{
							strategyKey: newStrategy,
						})
					}
				}
			} else {
				logs.Error(`impossible to FilterStrategyAdded for NewYvault022 ` + vault.Address.Hex() + ` on chain ` + strconv.FormatUint(chainID, 10) + `: ` + err.Error())
			}
		case `0.3.0`, `0.3.1`:
			currentVault, _ := contracts.NewYvault030(vault.Address, client)
			if log, err := currentVault.FilterStrategyAdded(opts, nil); err == nil {
				for log.Next() {
					if log.Error() != nil {
						continue
					}
					newStrategy := handleV03Strategies(chainID, vault.Version, log.Event)
					if storage.StoreStrategyIfMissing(chainID, newStrategy) {
						strategyKey := newStrategy.Address.Hex() + `_` + newStrategy.VaultAddress.Hex()
						fetcher.RetrieveAllStrategies(chainID, map[string]models.TStrategy{
							strategyKey: newStrategy,
						})
					}
				}
			} else {
				logs.Error(`impossible to FilterStrategyAdded for NewYvault030 ` + vault.Address.Hex() + ` on chain ` + strconv.FormatUint(chainID, 10) + `: ` + err.Error())
			}

			if log, err := currentVault.FilterStrategyMigrated(opts, nil, nil); err == nil {
				for log.Next() {
					if log.Error() != nil {
						continue
					}
					newMigratedStrategy := handleV03StrategiesMigration(chainID, log.Event)
					storage.StoreStrategyMigrated(chainID, newMigratedStrategy)
					processMigrations(chainID)
					if newStrategy, ok := storage.GetStrategy(
						chainID,
						newMigratedStrategy.NewStrategyAddress,
						newMigratedStrategy.VaultAddress,
					); ok {
						strategyKey := newStrategy.Address.Hex() + `_` + newStrategy.VaultAddress.Hex()
						fetcher.RetrieveAllStrategies(chainID, map[string]models.TStrategy{
							strategyKey: newStrategy,
						})
					}
				}
			} else {
				logs.Error(`impossible to FilterStrategyMigrated for NewYvault030 ` + vault.Address.Hex() + ` on chain ` + strconv.FormatUint(chainID, 10) + `: ` + err.Error())
			}
		case `0.3.2`, `0.3.3`, `0.3.4`, `0.3.5`, `0.4.2`, `0.4.3`, `0.4.4`, `0.4.5`, `0.4.6`, `0.4.7`:
			currentVault, _ := contracts.NewYvault043(vault.Address, client)
			if log, err := currentVault.FilterStrategyAdded(opts, nil); err == nil {
				for log.Next() {
					if log.Error() != nil {
						continue
					}
					newStrategy := handleV04Strategies(chainID, vault.Version, log.Event)
					if storage.StoreStrategyIfMissing(chainID, newStrategy) {
						strategyKey := newStrategy.Address.Hex() + `_` + newStrategy.VaultAddress.Hex()
						fetcher.RetrieveAllStrategies(chainID, map[string]models.TStrategy{
							strategyKey: newStrategy,
						})
					}
				}
			} else {
				logs.Error(`impossible to FilterStrategyAdded for NewYvault043 ` + vault.Address.Hex() + ` on chain ` + strconv.FormatUint(chainID, 10) + `: ` + err.Error())
			}

			if log, err := currentVault.FilterStrategyMigrated(opts, nil, nil); err == nil {
				for log.Next() {
					if log.Error() != nil {
						continue
					}
					newMigratedStrategy := handleV04StrategiesMigration(chainID, log.Event)
					storage.StoreStrategyMigrated(chainID, newMigratedStrategy)
					processMigrations(chainID)
					if newStrategy, ok := storage.GetStrategy(
						chainID,
						newMigratedStrategy.NewStrategyAddress,
						newMigratedStrategy.VaultAddress,
					); ok {
						strategyKey := newStrategy.Address.Hex() + `_` + newStrategy.VaultAddress.Hex()
						fetcher.RetrieveAllStrategies(chainID, map[string]models.TStrategy{
							strategyKey: newStrategy,
						})
					}
				}
			} else {
				logs.Error(`impossible to FilterStrategyMigrated for NewYvault030 ` + vault.Address.Hex() + ` on chain ` + strconv.FormatUint(chainID, 10) + `: ` + err.Error())
			}
		default:
			// case `3.0.0`, `3.0.1`, `3.0.2`:
			currentVault, _ := contracts.NewYvault300(vault.Address, client)
			if log, err := currentVault.FilterStrategyChanged(opts, nil, nil); err == nil {
				for log.Next() {
					if log.Error() != nil {
						continue
					}
					newStrategy := handleV300Strategies(chainID, vault.Version, log.Event)
					if storage.StoreStrategyIfMissing(chainID, newStrategy) {
						strategyKey := newStrategy.Address.Hex() + `_` + newStrategy.VaultAddress.Hex()
						fetcher.RetrieveAllStrategies(chainID, map[string]models.TStrategy{
							strategyKey: newStrategy,
						})
					}
				}
			} else {
				logs.Error(`impossible to FilterStrategyAdded for NewYvault043 ` + vault.Address.Hex() + ` on chain ` + strconv.FormatUint(chainID, 10) + `: ` + err.Error())
			}

			if log, err := currentVault.FilterStrategyChanged(opts, nil, nil); err == nil {
				for log.Next() {
					if log.Error() != nil {
						continue
					}
					if log.Event.ChangeType.Uint64() == 1 {
						historicalStrategy := handleV300Strategies(chainID, vault.Version, log.Event)
						if storage.StoreStrategyIfMissing(chainID, historicalStrategy) {
							strategyKey := historicalStrategy.Address.Hex() + `_` + historicalStrategy.VaultAddress.Hex()
							fetcher.RetrieveAllStrategies(chainID, map[string]models.TStrategy{
								strategyKey: historicalStrategy,
							})
						}
					}
				}
			} else {
				logs.Error(`impossible to FilterStrategyMigrated for NewYvault030 ` + vault.Address.Hex() + ` on chain ` + strconv.FormatUint(chainID, 10) + `: ` + err.Error())
			}
		}

	}
}


/** 🔵 - Yearn *************************************************************************************
** The function `IndexNewStrategies` is responsible for indexing new strategies for a given chain ID.
** This runs every 30 minutes to catch up on any new strategy events.
**************************************************************************************************/
func IndexNewStrategies(
	chainID uint64,
	vaults map[common.Address]models.TVault,
) (historicalStrategiesMap map[string]models.TStrategy) {
	// Initialize tracking map if needed
	if _lastIndexedBlock[chainID] == nil {
		_lastIndexedBlock[chainID] = make(map[common.Address]uint64)
	}

	// Get RPC client for current block number
	client := ethereum.GetRPC(chainID)
	currentBlock, err := client.BlockNumber(context.Background())
	if err != nil {
		logs.Error(`Failed to get current block number for chain ` + strconv.FormatUint(chainID, 10))
		historicalStrategiesMap, _ = storage.ListStrategies(chainID)
		return historicalStrategiesMap
	}

	/** 🔵 - Yearn *********************************************************************************
	** Loop over all the known vaults for the specified chain ID.
	**********************************************************************************************/
	for _, vault := range vaults {
		if env.IsRegistryDisabled(chainID, vault.RegistryAddress) {
			continue
		}
		if vault.Metadata.IsRetired || vault.Metadata.Migration.Available {
			continue
		}

		/** 🔵 - Yearn *************************************************************************************
		** Determine the starting block for event filtering. Use cached position if available,
		** otherwise find the highest strategy activation block.
		**************************************************************************************************/
		startBlock := _lastIndexedBlock[chainID][vault.Address]
		if startBlock == 0 {
			// First run: use highest strategy activation
			_, strategiesSlice := storage.ListStrategiesForVault(chainID, vault.Address)
			for _, strategy := range strategiesSlice {
				if strategy.Activation > startBlock {
					startBlock = strategy.Activation
				}
			}
			// If no strategies exist, start from vault activation
			if startBlock == 0 {
				startBlock = vault.Activation
			}
		}

		/** 🔵 - Yearn *************************************************************************************
		** Filter new strategy events from startBlock to currentBlock
		**************************************************************************************************/
		if startBlock < currentBlock {
			filterNewStrategies(chainID, vault, startBlock, currentBlock)
			_lastIndexedBlock[chainID][vault.Address] = currentBlock
		}

		/** 🔵 - Yearn *************************************************************************************
		** Also read current strategies directly from the vault contract
		**************************************************************************************************/
		strategies := listStrategiesForVault(chainID, vault)
		for _, strat := range strategies {
			if storage.StoreStrategyIfMissing(chainID, strat) {
				strategyKey := strat.Address.Hex() + `_` + strat.VaultAddress.Hex()
				fetcher.RetrieveAllStrategies(chainID, map[string]models.TStrategy{
					strategyKey: strat,
				})
			}
		}

		/** 🔵 - Yearn *************************************************************************************
		** For V3 vaults, also handle LastActiveStrategies that might not emit events
		**************************************************************************************************/
		if vault.Version >= "3.0.0" {
			for _, lastActiveStrategy := range vault.LastActiveStrategies {
				newStrategy := models.TStrategy{
					Address:      lastActiveStrategy,
					ChainID:      chainID,
					VaultVersion: vault.Version,
					VaultAddress: vault.Address,
					Activation:   vault.Activation,
				}
				if storage.StoreStrategyIfMissing(chainID, newStrategy) {
					strategyKey := newStrategy.Address.Hex() + `_` + newStrategy.VaultAddress.Hex()
					fetcher.RetrieveAllStrategies(chainID, map[string]models.TStrategy{
						strategyKey: newStrategy,
					})
				}
			}
		}
	}

	/** 🔵 - Yearn *********************************************************************************
	** Process any strategy migrations
	**********************************************************************************************/
	processMigrations(chainID)
	historicalStrategiesMap, _ = storage.ListStrategies(chainID)
	return historicalStrategiesMap
}
