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

// Note: Last indexed blocks are now persisted to disk via storage.GetLastIndexedBlock/SetLastIndexedBlock

/**************************************************************************************************
** listStrategiesForVault reads current strategies directly from the vault contract
**************************************************************************************************/
func listStrategiesForVault(
	chainID uint64,
	vault models.TVault,
) []models.TStrategy {
	client := ethereum.GetRPC(chainID)

	strategies := []models.TStrategy{}
	// Only v3.0.0+ vaults reach this function due to filtering in IndexNewStrategies()
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

		currentVault, _ := contracts.NewYvault300(vault.Address, client)
		if log, err := currentVault.FilterStrategyChanged(opts, nil, nil); err == nil {
			for log.Next() {
				if log.Error() != nil {
					continue
				}
				newStrategy := handleV300Strategies(chainID, vault.Version, log.Event)
				if storage.StoreStrategyIfMissing(chainID, newStrategy) {
					logs.Info(`Found strategy change for ` + newStrategy.Address.Hex() + ` on vault ` + vault.Address.Hex() + ` at block ` + strconv.FormatUint(newStrategy.Activation, 10))
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

/** 🔵 - Yearn *************************************************************************************
** The function `IndexNewStrategies` is responsible for indexing new strategies for a given chain ID.
** This runs every 30 minutes to catch up on any new strategy events.
**************************************************************************************************/
func IndexNewStrategies(
	chainID uint64,
	vaults map[common.Address]models.TVault,
) (historicalStrategiesMap map[string]models.TStrategy) {
	// Initialize persisted storage for this chain
	storage.InitializeIndexerStorage(chainID)

	// Get RPC client for current block number
	client := ethereum.GetRPC(chainID)
	currentBlock, err := client.BlockNumber(context.Background())
	if err != nil {
		logs.Error(`Failed to get current block number for chain ` + strconv.FormatUint(chainID, 10))
		historicalStrategiesMap, _ = storage.ListStrategies(chainID)
		return historicalStrategiesMap
	}

	logs.Info(`Starting strategy indexing for chain ` + strconv.FormatUint(chainID, 10) + ` at block ` + strconv.FormatUint(currentBlock, 10))

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
		// Only process v3.0.0+ vaults
		if vault.Version < "3.0.0" {
			continue
		}

		/** 🔵 - Yearn *************************************************************************************
		** Determine the starting block for event filtering. Use persisted position if available,
		** otherwise find the highest strategy activation block.
		**************************************************************************************************/
		startBlock := storage.GetLastIndexedBlock(chainID, vault.Address)
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
			logs.Info(`First run for vault ` + vault.Address.Hex() + ` on chain ` + strconv.FormatUint(chainID, 10) + `, starting from block ` + strconv.FormatUint(startBlock, 10))
		}

		/** 🔵 - Yearn *************************************************************************************
		** Filter new strategy events from startBlock to currentBlock
		**************************************************************************************************/
		if startBlock < currentBlock {
			blockRange := currentBlock - startBlock
			logs.Info(`Scanning blocks ` + strconv.FormatUint(startBlock, 10) + ` to ` + strconv.FormatUint(currentBlock, 10) + ` (` + strconv.FormatUint(blockRange, 10) + ` blocks) for vault ` + vault.Address.Hex() + ` v` + vault.Version)
			filterNewStrategies(chainID, vault, startBlock, currentBlock)
			storage.SetLastIndexedBlock(chainID, vault.Address, currentBlock)
		} else {
			logs.Info(`Vault ` + vault.Address.Hex() + ` already up to date at block ` + strconv.FormatUint(currentBlock, 10))
		}

		/** 🔵 - Yearn *************************************************************************************
		** Also read current strategies directly from the vault contract
		**************************************************************************************************/
		strategies := listStrategiesForVault(chainID, vault)
		if len(strategies) > 0 {
			logs.Info(`Reading ` + strconv.Itoa(len(strategies)) + ` strategies directly from vault contract ` + vault.Address.Hex())
		}
		for _, strat := range strategies {
			if storage.StoreStrategyIfMissing(chainID, strat) {
				logs.Info(`Stored strategy ` + strat.Address.Hex() + ` from direct contract read`)
				strategyKey := strat.Address.Hex() + `_` + strat.VaultAddress.Hex()
				fetcher.RetrieveAllStrategies(chainID, map[string]models.TStrategy{
					strategyKey: strat,
				})
			}
		}

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

	/** 🔵 - Yearn *********************************************************************************
	** Process any strategy migrations
	**********************************************************************************************/
	processMigrations(chainID)
	historicalStrategiesMap, _ = storage.ListStrategies(chainID)

	logs.Info(`Completed strategy indexing for chain ` + strconv.FormatUint(chainID, 10) + `, total strategies: ` + strconv.Itoa(len(historicalStrategiesMap)))
	return historicalStrategiesMap
}
