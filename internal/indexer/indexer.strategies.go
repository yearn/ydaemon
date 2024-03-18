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
	"github.com/yearn/ydaemon/common/contracts"
	"github.com/yearn/ydaemon/common/env"
	"github.com/yearn/ydaemon/common/ethereum"
	"github.com/yearn/ydaemon/common/logs"
	"github.com/yearn/ydaemon/internal/fetcher"
	"github.com/yearn/ydaemon/internal/models"
	"github.com/yearn/ydaemon/internal/storage"
)

var _strategiesAlreadyIndexingForVaults = make(map[uint64]*sync.Map)

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

	/**********************************************************************************************
	** First, we need to know when to stop our log fetching. By default, we will fetch until the
	** current block number, aka nil.
	** As using nil may cause some issues, we will specify the current block number instead.
	**********************************************************************************************/
	if end == nil {
		blockEnd, _ := client.BlockNumber(context.Background())
		end = &blockEnd
	}

	for chunkStart := start; chunkStart < *end; chunkStart += env.CHAINS[chainID].MaxBlockRange {
		chunkEnd := chunkStart + env.CHAINS[chainID].MaxBlockRange
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
						fetcher.RetrieveAllStrategies(chainID, map[common.Address]models.TStrategy{
							newStrategy.Address: newStrategy,
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
						fetcher.RetrieveAllStrategies(chainID, map[common.Address]models.TStrategy{
							newStrategy.Address: newStrategy,
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
						fetcher.RetrieveAllStrategies(chainID, map[common.Address]models.TStrategy{
							newStrategy.Address: newStrategy,
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
						fetcher.RetrieveAllStrategies(chainID, map[common.Address]models.TStrategy{
							newStrategy.Address: newStrategy,
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
						fetcher.RetrieveAllStrategies(chainID, map[common.Address]models.TStrategy{
							newStrategy.Address: newStrategy,
						})
					}
				}
			} else {
				logs.Error(`impossible to FilterStrategyMigrated for NewYvault030 ` + vault.Address.Hex() + ` on chain ` + strconv.FormatUint(chainID, 10) + `: ` + err.Error())
			}
		default:
			// case `3.0.0`, `3.0.1`:
			currentVault, _ := contracts.NewYvault300(vault.Address, client)
			if log, err := currentVault.FilterStrategyChanged(opts, nil, nil); err == nil {
				for log.Next() {
					if log.Error() != nil {
						continue
					}
					newStrategy := handleV300Strategies(chainID, vault.Version, log.Event)
					if storage.StoreStrategyIfMissing(chainID, newStrategy) {
						fetcher.RetrieveAllStrategies(chainID, map[common.Address]models.TStrategy{
							newStrategy.Address: newStrategy,
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
							fetcher.RetrieveAllStrategies(chainID, map[common.Address]models.TStrategy{
								historicalStrategy.Address: historicalStrategy,
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
	chainID uint64,
	vault models.TVault,
	lastSyncedBlock uint64,
	wg *sync.WaitGroup,
	isDone bool,
) (uint64, bool, error) {
	/**********************************************************************************************
	** First thing is to connect to the node via a WS connection. We need to use a WS connection
	** because we need to listen to new events as they are emitted via the node.
	** Not all nodes support WS connections, so we need to check if the node supports it.
	**********************************************************************************************/
	client, _ := ethereum.GetWSClient(chainID)

	switch vault.Version {
	case `0.2.2`:
		currentVault, _ := contracts.NewYvault022(vault.Address, client)
		etherReader := ethereum.Reader{Backend: client}
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
					fetcher.RetrieveAllStrategies(chainID, map[common.Address]models.TStrategy{
						newStrategy.Address: newStrategy,
					})
				}
			case err := <-sub.Err():
				return lastSyncedBlock, true, err
			}
		}
	case `0.3.0`, `0.3.1`:
		currentVault, _ := contracts.NewYvault030(vault.Address, client)
		etherReader := ethereum.Reader{Backend: client}
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
						fetcher.RetrieveAllStrategies(chainID, map[common.Address]models.TStrategy{
							newStrategy.Address: newStrategy,
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
						fetcher.RetrieveAllStrategies(chainID, map[common.Address]models.TStrategy{
							newStrategy.Address: newStrategy,
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
		etherReader := ethereum.Reader{Backend: client}
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
						fetcher.RetrieveAllStrategies(chainID, map[common.Address]models.TStrategy{
							newStrategy.Address: newStrategy,
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
						fetcher.RetrieveAllStrategies(chainID, map[common.Address]models.TStrategy{
							newStrategy.Address: newStrategy,
						})
					}
				}
			case err := <-sub.Err():
				return lastSyncedBlock, true, err
			}
		}
	default:
		// case `3.0.0`, `3.0.1`:
		currentVault, _ := contracts.NewYvault300(vault.Address, client)
		etherReader := ethereum.Reader{Backend: client}
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
		** Listen and handle new events
		**********************************************************************************************/
		for {
			select {
			case log := <-stream:
				if value, err := currentVault.ParseStrategyChanged(log); err == nil {
					lastSyncedBlock = value.Raw.BlockNumber
					newStrategy := handleV300Strategies(chainID, vault.Version, value)
					if storage.StoreStrategyIfMissing(chainID, newStrategy) {
						fetcher.RetrieveAllStrategies(chainID, map[common.Address]models.TStrategy{
							newStrategy.Address: newStrategy,
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
	err := error(nil)

	for {
		if _, err := ethereum.GetWSClient(chainID); err != nil {
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
				time.Sleep(1 * time.Minute)
			}
		}
		lastSyncedBlock, shouldRetry, err = watchNewStrategies(
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
			time.Sleep(30 * time.Second)
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
) (historicalStrategiesMap map[common.Address]models.TStrategy) {
	if _, ok := _strategiesAlreadyIndexingForVaults[chainID]; !ok {
		_strategiesAlreadyIndexingForVaults[chainID] = &sync.Map{}
	}

	wg := sync.WaitGroup{} // This WaitGroup will be done when all the historical strategies are indexed

	/** 🔵 - Yearn *********************************************************************************
	** Loop over all the known vaults for the specified chain ID.
	**********************************************************************************************/
	for _, vault := range vaults {
		/** 🔵 - Yearn *************************************************************************************
		** This block of code is responsible for checking if the strategies for a given vault are already
		** being indexed. If they are, it skips to the next vault. If they are not, it marks them as being
		** indexed to prevent duplicate work.
		**************************************************************************************************/
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
		go indexStrategyWrapper(chainID, vault, highestBlockNumber, &wg)
	}
	wg.Wait()

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
