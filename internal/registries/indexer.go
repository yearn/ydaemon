package registries

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
	"github.com/yearn/ydaemon/internal/indexer"
	"github.com/yearn/ydaemon/internal/models"
	"github.com/yearn/ydaemon/internal/storage"
)

func filterNewVault(
	chainID uint64,
	registry env.TContractData,
	start uint64,
	end *uint64,
	wg *sync.WaitGroup,
	isDone bool,
) (lastBlock uint64) {
	/**********************************************************************************************
	** As we cannot use the WS connection, we will fallback to the regular RPC connection. This
	** means that we will need to filter the logs from the lastSyncedBlock to the latest block
	** every x seconds to check if there are new vaults.
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

		if chunkEnd >= *end && !isDone {
			wg.Done()
		}
		opts := &bind.FilterOpts{Start: chunkStart, End: &chunkEnd}

		switch registry.Version {
		case 1, 2:
			currentRegistry, _ := contracts.NewYRegistryV2(registry.Address, client)
			if log, err := currentRegistry.FilterNewVault(opts, nil, nil); err == nil {
				for log.Next() {
					if log.Error() != nil {
						continue
					}
					handleV2Vault(chainID, log.Event)
				}
			} else {
				logs.Error(`impossible to FilterNewVault for YRegistryV2 ` + registry.Address.Hex() + ` on chain ` + strconv.FormatUint(chainID, 10) + `: ` + err.Error())
			}
		case 3:
			currentRegistry, _ := contracts.NewYRegistryV3(registry.Address, client)
			if log, err := currentRegistry.FilterNewVault(opts, nil, nil); err == nil {
				for log.Next() {
					if log.Error() != nil {
						continue
					}
					handleV3Vault(chainID, log.Event)
				}
			} else {
				logs.Error(`impossible to FilterNewVault for YRegistryV3 ` + registry.Address.Hex() + ` on chain ` + strconv.FormatUint(chainID, 10) + `: ` + err.Error())
			}
		case 4:
			currentRegistry, _ := contracts.NewYRegistryV4(registry.Address, client)
			if log, err := currentRegistry.FilterNewEndorsedVault(opts, nil, nil); err == nil {
				for log.Next() {
					if log.Error() != nil {
						continue
					}
					handleV4Vault(chainID, log.Event)
				}
			} else {
				logs.Error(`impossible to FilterNewVault for YRegistryV4 ` + registry.Address.Hex() + ` on chain ` + strconv.FormatUint(chainID, 10) + `: ` + err.Error())
			}
		}
	}
	return lastBlock
}

/**************************************************************************************************
** Watch for new vaults added to the registry. This function is called by the infinite loop in
** indexNewVaultsWrapper. It uses the WS connection to listen to the events.
** It will catch the new vaults and start all the subsequent processes to add them correctly in
** yDaemon.
** On errror, it will return the lastSyncedBlock, a boolean to indicate if we should retry and the
** error.
**
** Arguments:
** - chainID: the chain ID of the network we are listening to
** - registryAddress: the address of the registry we are listening to
** - registryVersion: the version of the registry we are listening to
** - lastSyncedBlock: the last block we have synced
**
** Returns:
** - lastSyncedBlock: the last block we have synced
** - shouldRetry: a boolean to indicate if we should retry
** - err: the error if any
**************************************************************************************************/
func watchNewVaults(
	chainID uint64,
	registry env.TContractData,
	lastSyncedBlock uint64,
	wg *sync.WaitGroup,
	newHistoricalVaults *sync.Map,
) (uint64, bool, error) {
	/**********************************************************************************************
	** First thing is to connect to the node via a WS connection. We need to use a WS connection
	** because we need to listen to new events as they are emitted via the node.
	** Not all nodes support WS connections, so we need to check if the node supports it.
	**********************************************************************************************/
	client, _ := ethereum.GetWSClient(chainID)

	switch registry.Version {
	case 1, 2:
		currentRegistry, _ := contracts.NewYRegistryV2(registry.Address, client)
		etherReader := ethereum.Reader{Backend: client}
		contractABI, _ := contracts.YRegistryV2MetaData.GetAbi()
		topics, _ := abi.MakeTopics([][]interface{}{{contractABI.Events[`NewVault`].ID}}...)
		query := goEth.FilterQuery{
			FromBlock: big.NewInt(int64(lastSyncedBlock)),
			Addresses: []common.Address{registry.Address},
			Topics:    topics,
		}
		stream, sub, history, err := etherReader.QueryWithHistory(context.Background(), &query)
		if err != nil {
			return 0, false, err
		}
		defer sub.Unsubscribe()

		/** 🔵 - Yearn *************************************************************************************
		** Handle historical events
		**************************************************************************************************/
		for _, log := range history {
			value, err := currentRegistry.ParseNewVault(log)
			if err != nil {
				continue
			}
			historicalVault := handleV2Vault(chainID, value)
			newHistoricalVaults.Store(historicalVault.Address, historicalVault)
		}
		if wg != nil {
			wg.Done()
		}

		/**********************************************************************************************
		** Listen and handle new events
		**********************************************************************************************/
		for {
			select {
			case log := <-stream:
				value, err := currentRegistry.ParseNewVault(log)
				if err != nil {
					continue
				}
				lastSyncedBlock = value.Raw.BlockNumber
				newVault := handleV2Vault(chainID, value)
				indexer.ProcessNewVault(chainID, map[common.Address]models.TVaultsFromRegistry{
					newVault.Address: newVault,
				})
			case err := <-sub.Err():
				logs.Error(err)
				return lastSyncedBlock, true, err
			}
		}
	case 3:
		currentRegistry, _ := contracts.NewYRegistryV3(registry.Address, client)
		etherReader := ethereum.Reader{Backend: client}
		contractABI, _ := contracts.YRegistryV3MetaData.GetAbi()
		topics, _ := abi.MakeTopics([][]interface{}{{contractABI.Events[`NewVault`].ID}}...)
		query := goEth.FilterQuery{
			FromBlock: big.NewInt(int64(registry.Activation)),
			Addresses: []common.Address{registry.Address},
			Topics:    topics,
		}
		stream, sub, history, err := etherReader.QueryWithHistory(context.Background(), &query)
		if err != nil {
			return 0, false, err
		}
		defer sub.Unsubscribe()

		/** 🔵 - Yearn *************************************************************************************
		** Handle historical events
		**************************************************************************************************/
		for _, log := range history {
			value, err := currentRegistry.ParseNewVault(log)
			if err != nil {
				continue
			}
			historicalVault := handleV3Vault(chainID, value)
			newHistoricalVaults.Store(historicalVault.Address, historicalVault)
		}
		if wg != nil {
			wg.Done()
		}

		/**********************************************************************************************
		** Listen and handle new events
		**********************************************************************************************/
		for {
			select {
			case log := <-stream:
				value, err := currentRegistry.ParseNewVault(log)
				if err != nil {
					continue
				}
				lastSyncedBlock = value.Raw.BlockNumber
				newVault := handleV3Vault(chainID, value)
				indexer.ProcessNewVault(chainID, map[common.Address]models.TVaultsFromRegistry{
					newVault.Address: newVault,
				})
			case err := <-sub.Err():
				logs.Error(err)
				return lastSyncedBlock, true, err
			}
		}
	case 4:
		currentRegistry, _ := contracts.NewYRegistryV4(registry.Address, client)
		etherReader := ethereum.Reader{Backend: client}
		contractABI, _ := contracts.YRegistryV4MetaData.GetAbi()
		topics, _ := abi.MakeTopics([][]interface{}{{contractABI.Events[`NewEndorsedVault`].ID}}...)
		query := goEth.FilterQuery{
			FromBlock: big.NewInt(int64(registry.Activation)),
			Addresses: []common.Address{registry.Address},
			Topics:    topics,
		}
		stream, sub, history, err := etherReader.QueryWithHistory(context.Background(), &query)
		if err != nil {
			return 0, false, err
		}
		defer sub.Unsubscribe()

		/** 🔵 - Yearn *************************************************************************************
		** Handle historical events
		**************************************************************************************************/
		for _, log := range history {
			value, err := currentRegistry.ParseNewEndorsedVault(log)
			if err != nil {
				continue
			}
			historicalVault := handleV4Vault(chainID, value)
			newHistoricalVaults.Store(historicalVault.Address, historicalVault)
		}
		if wg != nil {
			wg.Done()
		}

		/**********************************************************************************************
		** Listen and handle new events
		**********************************************************************************************/
		for {
			select {
			case log := <-stream:
				value, err := currentRegistry.ParseNewEndorsedVault(log)
				if err != nil {
					continue
				}
				lastSyncedBlock = value.Raw.BlockNumber
				newVault := handleV4Vault(chainID, value)
				indexer.ProcessNewVault(chainID, map[common.Address]models.TVaultsFromRegistry{
					newVault.Address: newVault,
				})
			case err := <-sub.Err():
				logs.Error(err)
				return lastSyncedBlock, true, err
			}
		}
	}
	return lastSyncedBlock, true, nil
}

/**************************************************************************************************
** As WS connections can be lost, can be instable or can just not be supported by the node, we need
** to have a fallback to the regular RPC connection.
** This function is called by the infinite loop and will try to index the new vaults using the WS
** but will also provide a way to catch the new vaults via the RPC connection.
**
** Arguments:
** - chainID: the chain ID of the network we are listening to
** - registryAddress: the address of the registry we are listening to
** - registryVersion: the version of the registry we are listening to
** - registryActivation: the activation block of the registry
** - delay: the delay between two standard RPC retries
**
** Returns nothing
**************************************************************************************************/
func indexNewVaultsWrapper(
	chainID uint64,
	registry env.TContractData,
	higherIndexedBlockNumber uint64,
	wg *sync.WaitGroup,
	newHistoricalVaults *sync.Map,
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
			break
		}
		lastSyncedBlock, shouldRetry, err = watchNewVaults(
			chainID,
			registry,
			lastSyncedBlock,
			wg,
			newHistoricalVaults,
		)
		if err != nil {
			logs.Error(`error while indexing NewVault from registry ` + registry.Address.Hex() + ` on chain ` + strconv.FormatUint(chainID, 10) + `: ` + err.Error())
		}
		if !shouldRetry {
			break
		}
	}

	/**********************************************************************************************
	** Default method: use the RPC connection to filter the events from the lastSyncedBlock to the
	** latest block. This is a fallback method in case the WS connection is not available.
	** It's only triggered if delay is greater than 0, allowing us to try to retry this whole
	** function under certain conditions while avoiding multiple calls to the same function.
	**********************************************************************************************/
	for {
		lastBlock := filterNewVault(
			chainID,
			registry,
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

/** 🔵 - Yearn *************************************************************************************
** The function `IndexNewVaults` is responsible for indexing new vaults for a given chain ID.
** It starts by logging a success message indicating that the Indexer Daemon has started for the
** specified chain ID.
**
** A `sync.WaitGroup` is used to ensure that all go routines launched in this function have
** completed before the function returns. This is important because each go routine is responsible
** for indexing new vaults from a different registry. If we didn't wait for all go routines to
** finish, the function could return prematurely, leaving some vaults unindexed.
**
** A `sync.Map` is used to store the new historical vaults that are indexed. We use a `sync.Map`
** instead of a regular map because `sync.Map` is safe to use concurrently from multiple go
** routines. This is crucial in this function because each go routine will be adding entries to
** the map concurrently.
**
** The `indexNewVaultsWrapper` will perform two distinct tasks:
**   - Retrieve the already existing vaults from the registry (aka filtering the logs)
**   - Listen to new vaults added to the registry (aka listening to the events)
** Only the first group is stored in the `sync.Map`.
**************************************************************************************************/
func IndexNewVaults(chainID uint64) (historicalVaults map[common.Address]models.TVaultsFromRegistry) {
	logs.Success(`Indexer Daemon has started for chain ` + strconv.FormatUint(chainID, 10))
	wg := sync.WaitGroup{}            // This WaitGroup will be done when all the historical vaults are indexed
	newHistoricalVaults := sync.Map{} // This Map will store the historical vaults

	// Load the previous historical vaults from JSON
	// This will be used to compare with the new vaults and update the historical vaults accordingly
	logs.Info("Loading existing vaults...")
	previousHistoricalVaults, higherBlockNumber := storage.LoadHistoricalVaultsFromJson(chainID)
	for vaultAddress, vault := range previousHistoricalVaults {
		newHistoricalVaults.Store(vaultAddress, vault)
	}

	// Index the new vaults
	logs.Info("Indexing new vaults...")
	wg.Add(len(env.CHAINS[chainID].Registries))
	for _, registry := range env.CHAINS[chainID].Registries {
		go indexNewVaultsWrapper(chainID, registry, higherBlockNumber, &wg, &newHistoricalVaults)
	}
	wg.Wait()

	logs.Info("Processing new vaults...")
	// Create a map to store the values from newHistoricalVaults
	historicalVaults = make(map[common.Address]models.TVaultsFromRegistry)
	newHistoricalVaults.Range(func(key, value interface{}) bool {
		historicalVaults[key.(common.Address)] = value.(models.TVaultsFromRegistry)
		return true
	})
	indexer.ProcessNewVault(chainID, historicalVaults)

	// Store the new historical vaults to JSON
	storage.StoreHistoricalVaultsToJson(chainID, historicalVaults)
	return historicalVaults
}
