package indexer

import (
	"context"
	"math/big"
	"strconv"
	"sync"
	"time"

	goEth "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/yearn/ydaemon/common/addresses"
	"github.com/yearn/ydaemon/common/contracts"
	"github.com/yearn/ydaemon/common/env"
	"github.com/yearn/ydaemon/common/ethereum"
	"github.com/yearn/ydaemon/common/logs"
	"github.com/yearn/ydaemon/internal/fetcher"
	"github.com/yearn/ydaemon/internal/models"
	"github.com/yearn/ydaemon/internal/storage"
)

var _strategiesAlreadyIndexingForVaults = make(map[uint64]*sync.Map)

/**************************************************************************************************
** listStrategiesForVault...
**************************************************************************************************/
func listStrategiesForVault(
	chainID uint64,
	vault models.TVault,
	wg *sync.WaitGroup,
	isDone bool,
) []models.TStrategy {
	/**********************************************************************************************
	** As we cannot use the WS connection, we will fallback to the regular RPC connection. This
	** means that we will need to filter the logs from the lastSyncedBlock to the latest block
	** every x seconds to check if there are new strategies.
	**********************************************************************************************/
	client := ethereum.GetRPC(chainID)

	if !isDone && wg != nil {
		wg.Done()
	}

	strategies := []models.TStrategy{}
	switch vault.Version {
	case `0.2.2`, `0.3.0`, `0.3.1`, `0.3.2`, `0.3.3`, `0.3.4`, `0.3.5`, `0.4.2`, `0.4.3`:
		/******************************************************************************************
		** Vaults versions from 0.2.2 to 0.4.3 are now deprecated and should not be used anymore.
		** This is only for historical purposes.
		** The strategies indexer will not run for this version.
		******************************************************************************************/
		break
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
** filterNewStrategies...
**************************************************************************************************/
func filterNewStrategies(
	chainID uint64,
	vault models.TVault,
	start uint64,
	end *uint64,
	wg *sync.WaitGroup,
	isDone bool,
) (lastBlock uint64) {
	/**********************************************************************************************
	** As we cannot use the WS connection, we will fallback to the regular RPC connection. This
	** means that we will need to filter the logs from the lastSyncedBlock to the latest block
	** every x seconds to check if there are new strategies.
	**********************************************************************************************/
	client := ethereum.GetRPC(chainID)
	chain, ok := env.GetChain(chainID)
	if !ok {
		return 0
	}

	/**********************************************************************************************
	** First, we need to know when to stop our log fetching. By default, we will fetch until the
	** current block number, aka nil.
	** As using nil may cause some issues, we will specify the current block number instead.
	**********************************************************************************************/
	if end == nil {
		blockEnd, _ := client.BlockNumber(context.Background())
		end = &blockEnd
	}

	for chunkStart := start; chunkStart < *end; chunkStart += chain.MaxBlockRange {
		chunkEnd := chunkStart + chain.MaxBlockRange
		if chunkEnd > *end {
			chunkEnd = *end
			lastBlock = chunkEnd
		}

		if chunkEnd >= *end && !isDone && wg != nil {
			wg.Done()
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
	return lastBlock
}

/**************************************************************************************************
** watchNewStrategies...
**************************************************************************************************/
func watchNewStrategies(
	client *ethclient.Client,
	chainID uint64,
	vault models.TVault,
	lastSyncedBlock uint64,
	wg *sync.WaitGroup,
	isDone bool,
) (uint64, bool, error) {
	switch vault.Version {
	case `0.2.2`:
		currentVault, _ := contracts.NewYvault022(vault.Address, client)
		etherReader := ethereum.Reader{Backend: client, ChainID: chainID}
		contractABI, _ := contracts.Yvault022MetaData.GetAbi()
		topics, _ := abi.MakeTopics([][]interface{}{{contractABI.Events[`StrategyAdded`].ID}}...)
		query := goEth.FilterQuery{
			FromBlock: big.NewInt(int64(lastSyncedBlock)),
			Addresses: []common.Address{vault.Address},
			Topics:    topics,
		}
		stream, sub, history, err := etherReader.QueryWithHistory(context.Background(), &query)
		if err != nil {
			if wg != nil && !isDone {
				wg.Done()
			}
			return 0, false, err
		}
		defer sub.Unsubscribe()

		/** 🔵 - Yearn *************************************************************************************
		** Handle historical events. It's only a storing action as the rest will be performed as a batch,
		** all the one in the history in one go.
		**************************************************************************************************/
		for _, log := range history {
			if value, err := currentVault.ParseStrategyAdded(log); err == nil {
				historicalStrategy := handleV02Strategies(chainID, vault.Version, value)
				storage.StoreStrategyIfMissing(chainID, historicalStrategy)
			}
		}
		if wg != nil && !isDone {
			wg.Done()
		}

		/**********************************************************************************************
		** Listen and handle new events
		**********************************************************************************************/
		for {
			select {
			case log := <-stream:
				value, err := currentVault.ParseStrategyAdded(log)
				if err != nil {
					continue
				}
				lastSyncedBlock = value.Raw.BlockNumber
				newStrategy := handleV02Strategies(chainID, vault.Version, value)
				if storage.StoreStrategyIfMissing(chainID, newStrategy) {
					strategyKey := newStrategy.Address.Hex() + `_` + newStrategy.VaultAddress.Hex()
					fetcher.RetrieveAllStrategies(chainID, map[string]models.TStrategy{
						strategyKey: newStrategy,
					})
				}
			case err := <-sub.Err():
				return lastSyncedBlock, true, err
			}
		}
	case `0.3.0`, `0.3.1`:
		currentVault, _ := contracts.NewYvault030(vault.Address, client)
		etherReader := ethereum.Reader{Backend: client, ChainID: chainID}
		contractABI, _ := contracts.Yvault030MetaData.GetAbi()
		topics, _ := abi.MakeTopics([][]interface{}{{
			contractABI.Events[`StrategyAdded`].ID,
			contractABI.Events[`StrategyMigrated`].ID,
		}}...)
		query := goEth.FilterQuery{
			FromBlock: big.NewInt(int64(vault.Activation)),
			Addresses: []common.Address{vault.Address},
			Topics:    topics,
		}
		stream, sub, history, err := etherReader.QueryWithHistory(context.Background(), &query)
		if err != nil {
			if wg != nil && !isDone {
				wg.Done()
			}
			return 0, false, err
		}
		defer sub.Unsubscribe()

		/** 🔵 - Yearn *************************************************************************************
		** Handle historical events. It's only a storing action as the rest will be performed as a batch,
		** all the one in the history in one go.
		**************************************************************************************************/
		for _, log := range history {
			if value, err := currentVault.ParseStrategyAdded(log); err == nil {
				historicalStrategy := handleV03Strategies(chainID, vault.Version, value)
				storage.StoreStrategyIfMissing(chainID, historicalStrategy)
				continue
			}
			if value, err := currentVault.ParseStrategyMigrated(log); err == nil {
				historicalStrategy := handleV03StrategiesMigration(chainID, value)
				storage.StoreStrategyMigrated(chainID, historicalStrategy)
				continue
			}
		}
		if wg != nil && !isDone {
			wg.Done()
		}

		/**********************************************************************************************
		** Listen and handle new events
		**********************************************************************************************/
		for {
			select {
			case log := <-stream:
				if value, err := currentVault.ParseStrategyAdded(log); err == nil {
					lastSyncedBlock = value.Raw.BlockNumber
					newStrategy := handleV03Strategies(chainID, vault.Version, value)
					if storage.StoreStrategyIfMissing(chainID, newStrategy) {
						strategyKey := newStrategy.Address.Hex() + `_` + newStrategy.VaultAddress.Hex()
						fetcher.RetrieveAllStrategies(chainID, map[string]models.TStrategy{
							strategyKey: newStrategy,
						})
					}
					continue
				}

				if value, err := currentVault.ParseStrategyMigrated(log); err == nil {
					lastSyncedBlock = value.Raw.BlockNumber
					newMigratedStrategy := handleV03StrategiesMigration(chainID, value)
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
					continue
				}
			case err := <-sub.Err():
				return lastSyncedBlock, true, err
			}
		}
	case `0.3.2`, `0.3.3`, `0.3.4`, `0.3.5`, `0.4.2`, `0.4.3`, `0.4.4`, `0.4.5`, `0.4.6`, `0.4.7`:
		currentVault, _ := contracts.NewYvault043(vault.Address, client)
		etherReader := ethereum.Reader{Backend: client, ChainID: chainID}
		contractABI, _ := contracts.Yvault043MetaData.GetAbi()
		topics, _ := abi.MakeTopics([][]interface{}{{
			contractABI.Events[`StrategyAdded`].ID,
			contractABI.Events[`StrategyMigrated`].ID,
		}}...)
		query := goEth.FilterQuery{
			FromBlock: big.NewInt(int64(vault.Activation)),
			Addresses: []common.Address{vault.Address},
			Topics:    topics,
		}
		stream, sub, history, err := etherReader.QueryWithHistory(context.Background(), &query)
		if err != nil {
			if wg != nil && !isDone {
				wg.Done()
			}
			return 0, false, err
		}
		defer sub.Unsubscribe()

		/** 🔵 - Yearn *************************************************************************************
		** Handle historical events. It's only a storing action as the rest will be performed as a batch,
		** all the one in the history in one go.
		**************************************************************************************************/
		for _, log := range history {
			if value, err := currentVault.ParseStrategyAdded(log); err == nil {
				historicalStrategy := handleV04Strategies(chainID, vault.Version, value)
				storage.StoreStrategyIfMissing(chainID, historicalStrategy)
				continue
			}
			if value, err := currentVault.ParseStrategyMigrated(log); err == nil {
				historicalStrategy := handleV04StrategiesMigration(chainID, value)
				storage.StoreStrategyMigrated(chainID, historicalStrategy)
				continue
			}
		}
		if wg != nil && !isDone {
			wg.Done()
		}

		/**********************************************************************************************
		** Listen and handle new events
		**********************************************************************************************/
		for {
			select {
			case log := <-stream:
				if value, err := currentVault.ParseStrategyAdded(log); err == nil {
					lastSyncedBlock = value.Raw.BlockNumber
					newStrategy := handleV04Strategies(chainID, vault.Version, value)
					if storage.StoreStrategyIfMissing(chainID, newStrategy) {
						strategyKey := newStrategy.Address.Hex() + `_` + newStrategy.VaultAddress.Hex()
						fetcher.RetrieveAllStrategies(chainID, map[string]models.TStrategy{
							strategyKey: newStrategy,
						})
					}
					continue
				}

				if value, err := currentVault.ParseStrategyMigrated(log); err == nil {
					lastSyncedBlock = value.Raw.BlockNumber
					newMigratedStrategy := handleV04StrategiesMigration(chainID, value)
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
			case err := <-sub.Err():
				return lastSyncedBlock, true, err
			}
		}
	default:
		// case `3.0.0`, `3.0.1`, `3.0.2`:
		currentVault, _ := contracts.NewYvault300(vault.Address, client)
		etherReader := ethereum.Reader{Backend: client, ChainID: chainID}
		contractABI, _ := contracts.Yvault300MetaData.GetAbi()
		topics, _ := abi.MakeTopics([][]interface{}{{contractABI.Events[`StrategyChanged`].ID}}...)
		query := goEth.FilterQuery{
			FromBlock: big.NewInt(int64(vault.Activation)),
			Addresses: []common.Address{vault.Address},
			Topics:    topics,
		}
		stream, sub, history, err := etherReader.QueryWithHistory(context.Background(), &query)
		if err != nil {
			if wg != nil && !isDone {
				wg.Done()
			}
			return 0, false, err
		}
		defer sub.Unsubscribe()

		/** 🔵 - Yearn *************************************************************************************
		** Handle historical events. It's only a storing action as the rest will be performed as a batch,
		** all the one in the history in one go.
		**************************************************************************************************/
		for _, log := range history {
			if value, err := currentVault.ParseStrategyChanged(log); err == nil {
				if value.ChangeType.Uint64() == 1 {
					historicalStrategy := handleV300Strategies(chainID, vault.Version, value)
					storage.StoreStrategyIfMissing(chainID, historicalStrategy)
				}
				continue
			}
		}
		if wg != nil && !isDone {
			wg.Done()
		}

		/**********************************************************************************************
		** Because now some stategies are not added via an event but directly in the contract, we need
		** to fetch them directly from the contract.
		** Ex: https://etherscan.io/address/0x92545bCE636E6eE91D88D2D017182cD0bd2fC22e#events
		**********************************************************************************************/
		for _, lastActiveStrategy := range vault.LastActiveStrategies {
			newStrategy := models.TStrategy{
				Address:      lastActiveStrategy,
				ChainID:      chainID,
				VaultVersion: vault.Version,
				VaultAddress: vault.Address,
				Activation:   vault.Activation,
			}
			storage.StoreStrategyIfMissing(chainID, newStrategy)
		}

		/**********************************************************************************************
		** Listen and handle new events
		**********************************************************************************************/
		for {
			select {
			case log := <-stream:
				if value, err := currentVault.ParseStrategyChanged(log); err == nil {
					lastSyncedBlock = value.Raw.BlockNumber
					newStrategy := handleV300Strategies(chainID, vault.Version, value)
					if storage.StoreStrategyIfMissing(chainID, newStrategy) {
						strategyKey := newStrategy.Address.Hex() + `_` + newStrategy.VaultAddress.Hex()
						fetcher.RetrieveAllStrategies(chainID, map[string]models.TStrategy{
							strategyKey: newStrategy,
						})
					}
					continue
				}
			case err := <-sub.Err():
				return lastSyncedBlock, true, err
			}
		}
	}
}

/**************************************************************************************************
** indexStrategyWrapper ...
**************************************************************************************************/
func indexStrategyWrapper(
	chainID uint64,
	vault models.TVault,
	higherIndexedBlockNumber uint64,
	wg *sync.WaitGroup,
) {
	isDone := false
	lastSyncedBlock := higherIndexedBlockNumber

	/**********************************************************************************************
	** Initialize the infinite loop to listen to new events. This is a wrapper around the actual
	** indexer to be able to catch the errors, to restart the indexer or just to exit the loop to
	** fallback to another method.
	**********************************************************************************************/
	shouldRetry := true

	for {
		client, err := ethereum.GetWSClient(chainID, true)
		if err != nil {
			/**********************************************************************************************
			** Default method: use the RPC connection to filter the events from the lastSyncedBlock to the
			** latest block. This is a fallback method in case the WS connection is not available.
			** It's only triggered if delay is greater than 0, allowing us to try to retry this whole
			** function under certain conditions while avoiding multiple calls to the same function.
			**********************************************************************************************/
			for {
				lastBlock := filterNewStrategies(
					chainID,
					vault,
					lastSyncedBlock,
					nil,
					wg,
					isDone,
				)
				isDone = true
				lastSyncedBlock = lastBlock
				time.Sleep(1 * time.Hour)
			}
		}

		lastSyncedBlock, shouldRetry, err = watchNewStrategies(
			client,
			chainID,
			vault,
			lastSyncedBlock,
			wg,
			isDone,
		)
		isDone = true
		if err != nil {
			logs.Error(`error while indexing New Strategies from vault ` + vault.Address.Hex() + ` on chain ` + strconv.FormatUint(chainID, 10) + `: ` + err.Error())
		}
		if shouldRetry {
			time.Sleep(10 * time.Minute)
			continue
		}
		break
	}
}

/** 🔵 - Yearn *************************************************************************************
** The function `IndexNewStrategies` is responsible for indexing new strategies for a given chain ID.
**************************************************************************************************/
func IndexNewStrategies(
	chainID uint64,
	vaults map[common.Address]models.TVault,
) (historicalStrategiesMap map[string]models.TStrategy) {
	shouldSkipIndexing := false
	if _, ok := _strategiesAlreadyIndexingForVaults[chainID]; !ok {
		_strategiesAlreadyIndexingForVaults[chainID] = &sync.Map{}
	}

	if shouldSkipIndexing {
		historicalStrategiesMap, _ = storage.ListStrategies(chainID)
		return historicalStrategiesMap
	}

	wg := sync.WaitGroup{} // This WaitGroup will be done when all the historical strategies are indexed

	/** 🔵 - Yearn *********************************************************************************
	** Loop over all the known vaults for the specified chain ID.
	**********************************************************************************************/
	allStrats := []models.TStrategy{}
	for _, vault := range vaults {
		/** 🔵 - Yearn *************************************************************************************
		** This block of code is responsible for checking if the strategies for a given vault are already
		** being indexed. If they are, it skips to the next vault. If they are not, it marks them as being
		** indexed to prevent duplicate work.
		**************************************************************************************************/
		if env.IsRegistryDisabled(chainID, vault.RegistryAddress) {
			continue
		}
		if _, ok := _strategiesAlreadyIndexingForVaults[chainID].Load(vault.Address); ok {
			continue
		}
		_strategiesAlreadyIndexingForVaults[chainID].Store(vault.Address, true)

		if vault.Metadata.IsRetired || vault.Metadata.Migration.Available {
			continue
		}

		/** 🔵 - Yearn *************************************************************************************
		** This block of code is responsible for retrieving the list of strategies for a given vault.
		** It then iterates over these strategies to find the one with the highest block number.
		** This is where we will start indexing new strategies from to avoid spending ressources with data
		** we already have.
		**************************************************************************************************/
		_, strategiesSlice := storage.ListStrategiesForVault(chainID, vault.Address)
		highestBlockNumber := uint64(0)
		for _, strategy := range strategiesSlice {
			if strategy.Activation > highestBlockNumber {
				highestBlockNumber = strategy.Activation
			}
		}

		/** 🔵 - Yearn *************************************************************************************
		** After retrieving the highest block number we can proceed to index new strategies.
		**************************************************************************************************/
		wg.Add(1)
		// go indexStrategyWrapper(chainID, vault, highestBlockNumber, &wg)
		allStrats = append(allStrats, listStrategiesForVault(chainID, vault, &wg, false)...)
	}
	wg.Wait()

	for _, strat := range allStrats {
		if storage.StoreStrategyIfMissing(chainID, strat) {
			strategyKey := strat.Address.Hex() + `_` + strat.VaultAddress.Hex()
			fetcher.RetrieveAllStrategies(chainID, map[string]models.TStrategy{
				strategyKey: strat,
			})
		}
	}

	/** 🔵 - Yearn *********************************************************************************
	** This part is only accessed once, once the `historical` strategies, aka that are already
	** deployed and indexed, are retrieved. New deployed strategies will not hit this part, but will
	** be handled in the above functions directly.
	** This specific part is used to handle the migration of strategies from one address to another
	** address.
	**********************************************************************************************/
	processMigrations(chainID)
	historicalStrategiesMap, _ = storage.ListStrategies(chainID)
	return historicalStrategiesMap
}
