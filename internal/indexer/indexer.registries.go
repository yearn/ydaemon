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
	"github.com/yearn/ydaemon/common/contracts"
	"github.com/yearn/ydaemon/common/env"
	"github.com/yearn/ydaemon/common/ethereum"
	"github.com/yearn/ydaemon/common/logs"
	"github.com/yearn/ydaemon/internal/fetcher"
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
	
	blockRange := *end - start
	logs.Info(`Scanning registry ` + registry.Address.Hex() + ` from block ` + strconv.FormatUint(start, 10) + ` to ` + strconv.FormatUint(*end, 10) + ` (` + strconv.FormatUint(blockRange, 10) + ` blocks)`)

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

		switch registry.Version {
		case 1, 2:
			currentRegistry, _ := contracts.NewYRegistryV2(registry.Address, client)
			if log, err := currentRegistry.FilterNewVault(opts, nil, nil); err == nil {
				for log.Next() {
					if log.Error() != nil {
						continue
					}
					historicalVault := handleV02Vault(chainID, log.Event)
					storage.StoreNewVaultToRegistry(chainID, historicalVault)
					ProcessNewVault(
						chainID,
						map[common.Address]models.TVaultsFromRegistry{historicalVault.Address: historicalVault},
						fetcher.ProcessNewVaultMethodAppend,
					)
				}
			} else {
				logs.Error(`impossible to FilterNewVault for YRegistryV2 ` + registry.Address.Hex() + ` on chain ` + strconv.FormatUint(chainID, 10) + `: ` + err.Error())
			}
		case 3:
			currentRegistry, _ := contracts.NewYRegistryV3(registry.Address, client)
			if log, err := currentRegistry.FilterNewVault(opts, nil, nil); err == nil {
				for log.Next() {
					if log.Error() != nil {
						logs.Error(`Error in YRegistryV3 for ` + registry.Address.Hex() + ` on chain ` + strconv.FormatUint(chainID, 10) + `: ` + log.Error().Error())
						continue
					}
					historicalVault := handleV03Vault(chainID, log.Event)
					storage.StoreNewVaultToRegistry(chainID, historicalVault)
					ProcessNewVault(
						chainID,
						map[common.Address]models.TVaultsFromRegistry{historicalVault.Address: historicalVault},
						fetcher.ProcessNewVaultMethodAppend,
					)
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
					historicalVault := handleV04Vault(chainID, log.Event)
					storage.StoreNewVaultToRegistry(chainID, historicalVault)
					ProcessNewVault(
						chainID,
						map[common.Address]models.TVaultsFromRegistry{historicalVault.Address: historicalVault},
						fetcher.ProcessNewVaultMethodAppend,
					)
				}
			} else {
				logs.Error(`impossible to FilterNewVault for YRegistryV4 ` + registry.Address.Hex() + ` on chain ` + strconv.FormatUint(chainID, 10) + `: ` + err.Error())
			}
		case 5:
			currentRegistry, _ := contracts.NewYRegistryV5(registry.Address, client)
			if log, err := currentRegistry.FilterNewVault(opts, nil, nil); err == nil {
				for log.Next() {
					if log.Error() != nil {
						continue
					}
					historicalVault := handleV05Vault(chainID, log.Event)
					storage.StoreNewVaultToRegistry(chainID, historicalVault)
					ProcessNewVault(
						chainID,
						map[common.Address]models.TVaultsFromRegistry{historicalVault.Address: historicalVault},
						fetcher.ProcessNewVaultMethodAppend,
					)
				}
			} else {
				logs.Error(`impossible to FilterNewVault for YRegistryV5 ` + registry.Address.Hex() + ` on chain ` + strconv.FormatUint(chainID, 10) + `: ` + err.Error())
			}
		case 6:
			currentRegistry, _ := contracts.NewYRegistryGamma(registry.Address, client)
			if log, err := currentRegistry.FilterNewGammaLPCompounder(opts, nil, nil); err == nil {
				for log.Next() {
					if log.Error() != nil {
						continue
					}
					historicalVault := handleV06Vault_Gamma(chainID, log.Event)
					storage.StoreNewVaultToRegistry(chainID, historicalVault)
					ProcessNewVault(
						chainID,
						map[common.Address]models.TVaultsFromRegistry{historicalVault.Address: historicalVault},
						fetcher.ProcessNewVaultMethodAppend,
					)
				}
			} else {
				logs.Error(`impossible to FilterNewVault for YRegistryV6 (Gamma) ` + registry.Address.Hex() + ` on chain ` + strconv.FormatUint(chainID, 10) + `: ` + err.Error())
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
	client *ethclient.Client,
	chainID uint64,
	registry env.TContractData,
	lastSyncedBlock uint64,
	wg *sync.WaitGroup,
	isDone bool,
) (uint64, bool, error) {
	switch registry.Version {
	case 1, 2:
		currentRegistry, _ := contracts.NewYRegistryV2(registry.Address, client)
		etherReader := ethereum.Reader{Backend: client, ChainID: chainID}
		contractABI, _ := contracts.YRegistryV2MetaData.GetAbi()
		topics, _ := abi.MakeTopics([][]interface{}{{
			contractABI.Events[`NewVault`].ID,
			contractABI.Events[`NewExperimentalVault`].ID,
		}}...)
		query := goEth.FilterQuery{
			FromBlock: big.NewInt(int64(lastSyncedBlock)),
			Addresses: []common.Address{registry.Address},
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
		** Handle historical events
		**************************************************************************************************/
		for _, log := range history {
			if value, err := currentRegistry.ParseNewVault(log); err == nil {
				historicalVault := handleV02Vault(chainID, value)
				storage.StoreNewVaultToRegistry(chainID, historicalVault)
				continue
			}
			if value, err := currentRegistry.ParseNewExperimentalVault(log); err == nil {
				historicalVault := handleV02ExperimentalVault(chainID, value)
				storage.StoreNewVaultToRegistry(chainID, historicalVault)
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
				if value, err := currentRegistry.ParseNewVault(log); err == nil {
					newVault := handleV02Vault(chainID, value)
					storage.StoreNewVaultToRegistry(chainID, newVault)
					ProcessNewVault(
						chainID,
						map[common.Address]models.TVaultsFromRegistry{newVault.Address: newVault},
						fetcher.ProcessNewVaultMethodAppend,
					)
					continue
				}

				if value, err := currentRegistry.ParseNewExperimentalVault(log); err == nil {
					newVault := handleV02ExperimentalVault(chainID, value)
					storage.StoreNewVaultToRegistry(chainID, newVault)
					ProcessNewVault(
						chainID,
						map[common.Address]models.TVaultsFromRegistry{newVault.Address: newVault},
						fetcher.ProcessNewVaultMethodAppend,
					)
				}
			case err := <-sub.Err():
				logs.Error(err)
				return lastSyncedBlock, true, err
			}
		}
	case 3:
		currentRegistry, _ := contracts.NewYRegistryV3(registry.Address, client)
		etherReader := ethereum.Reader{Backend: client, ChainID: chainID}
		contractABI, _ := contracts.YRegistryV3MetaData.GetAbi()
		topics, _ := abi.MakeTopics([][]interface{}{{contractABI.Events[`NewVault`].ID}}...)
		query := goEth.FilterQuery{
			FromBlock: big.NewInt(int64(registry.Block)),
			Addresses: []common.Address{registry.Address},
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
		** Handle historical events
		**************************************************************************************************/
		for _, log := range history {
			value, err := currentRegistry.ParseNewVault(log)
			if err != nil {
				continue
			}
			historicalVault := handleV03Vault(chainID, value)
			storage.StoreNewVaultToRegistry(chainID, historicalVault)
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
				value, err := currentRegistry.ParseNewVault(log)
				if err != nil {
					continue
				}
				lastSyncedBlock = value.Raw.BlockNumber
				newVault := handleV03Vault(chainID, value)
				ProcessNewVault(
					chainID,
					map[common.Address]models.TVaultsFromRegistry{newVault.Address: newVault},
					fetcher.ProcessNewVaultMethodAppend,
				)
			case err := <-sub.Err():
				logs.Error(err)
				return lastSyncedBlock, true, err
			}
		}
	case 4:
		currentRegistry, _ := contracts.NewYRegistryV4(registry.Address, client)
		etherReader := ethereum.Reader{Backend: client, ChainID: chainID}
		contractABI, _ := contracts.YRegistryV4MetaData.GetAbi()
		topics, _ := abi.MakeTopics([][]interface{}{{contractABI.Events[`NewEndorsedVault`].ID}}...)
		query := goEth.FilterQuery{
			FromBlock: big.NewInt(int64(registry.Block)),
			Addresses: []common.Address{registry.Address},
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
		** Handle historical events
		**************************************************************************************************/
		for _, log := range history {
			value, err := currentRegistry.ParseNewEndorsedVault(log)
			if err != nil {
				continue
			}
			historicalVault := handleV04Vault(chainID, value)
			storage.StoreNewVaultToRegistry(chainID, historicalVault)
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
				value, err := currentRegistry.ParseNewEndorsedVault(log)
				if err != nil {
					continue
				}
				lastSyncedBlock = value.Raw.BlockNumber
				newVault := handleV04Vault(chainID, value)
				ProcessNewVault(
					chainID,
					map[common.Address]models.TVaultsFromRegistry{newVault.Address: newVault},
					fetcher.ProcessNewVaultMethodAppend,
				)
			case err := <-sub.Err():
				logs.Error(err)
				return lastSyncedBlock, true, err
			}
		}
	case 5:
		currentRegistry, _ := contracts.NewYRegistryV5(registry.Address, client)
		etherReader := ethereum.Reader{Backend: client, ChainID: chainID}
		contractABI, _ := contracts.YRegistryV5MetaData.GetAbi()

		topics, _ := abi.MakeTopics([][]interface{}{{contractABI.Events[`NewVault`].ID}}...)
		query := goEth.FilterQuery{
			FromBlock: big.NewInt(int64(registry.Block)),
			Addresses: []common.Address{registry.Address},
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
		** Handle historical events
		**************************************************************************************************/
		for _, log := range history {
			value, err := currentRegistry.ParseNewVault(log)
			if err != nil {
				continue
			}
			historicalVault := handleV05Vault(chainID, value)
			storage.StoreNewVaultToRegistry(chainID, historicalVault)
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
				value, err := currentRegistry.ParseNewVault(log)
				if err != nil {
					continue
				}
				lastSyncedBlock = value.Raw.BlockNumber
				newVault := handleV05Vault(chainID, value)
				ProcessNewVault(
					chainID,
					map[common.Address]models.TVaultsFromRegistry{newVault.Address: newVault},
					fetcher.ProcessNewVaultMethodAppend,
				)
			case err := <-sub.Err():
				logs.Error(err)
				return lastSyncedBlock, true, err
			}
		}
	case 6:
		currentRegistry, _ := contracts.NewYRegistryGamma(registry.Address, client)
		etherReader := ethereum.Reader{Backend: client, ChainID: chainID}
		contractABI, _ := contracts.YRegistryGammaMetaData.GetAbi()
		topics, _ := abi.MakeTopics([][]interface{}{{contractABI.Events[`NewGammaLPCompounder`].ID}}...)
		query := goEth.FilterQuery{
			FromBlock: big.NewInt(int64(registry.Block)),
			Addresses: []common.Address{registry.Address},
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
		** Handle historical events
		**************************************************************************************************/
		for _, log := range history {
			value, err := currentRegistry.ParseNewGammaLPCompounder(log)
			if err != nil {
				continue
			}
			historicalVault := handleV06Vault_Gamma(chainID, value)
			storage.StoreNewVaultToRegistry(chainID, historicalVault)
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
				value, err := currentRegistry.ParseNewGammaLPCompounder(log)
				if err != nil {
					continue
				}
				lastSyncedBlock = value.Raw.BlockNumber
				newVault := handleV06Vault_Gamma(chainID, value)
				ProcessNewVault(
					chainID,
					map[common.Address]models.TVaultsFromRegistry{newVault.Address: newVault},
					fetcher.ProcessNewVaultMethodAppend,
				)
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
				time.Sleep(1 * time.Hour)
			}
		}

		lastSyncedBlock, shouldRetry, err = watchNewVaults(
			client,
			chainID,
			registry,
			lastSyncedBlock,
			wg,
			isDone,
		)
		isDone = true
		if err != nil {
			logs.Error(`error while indexing NewVault from registry ` + registry.Address.Hex() + ` on chain ` + strconv.FormatUint(chainID, 10) + `: ` + err.Error())
		}
		if shouldRetry {
			time.Sleep(10 * time.Minute)
			continue
		}
		break
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
func IndexNewVaults(chainID uint64) (vaultsFromRegistry map[common.Address]models.TVaultsFromRegistry) {
	shouldSkipIndexing := false
	wg := sync.WaitGroup{} // This WaitGroup will be done when all the historical vaults are indexed

	chain, ok := env.GetChain(chainID)
	if !ok {
		vaultsFromRegistry, _ = storage.ListVaultsFromRegistries(chainID)
		return vaultsFromRegistry
	}
	if chainID == 100 {
		vaultsFromRegistry, _ = storage.ListVaultsFromRegistries(chainID)
		return vaultsFromRegistry
	}

	if shouldSkipIndexing {
		vaultsFromRegistry, _ = storage.ListVaultsFromRegistries(chainID)
		return vaultsFromRegistry
	}
	
	registryCount := len(chain.Registries)
	logs.Info(`Starting vault registry scan for ` + strconv.Itoa(registryCount) + ` registries on chain ` + strconv.FormatUint(chainID, 10))

	/** 🔵 - Yearn *********************************************************************************
	** Loop over all the known registries for the specified chain ID.
	**********************************************************************************************/
	for _, registry := range chain.Registries {
		if registry.Tag == `DISABLED` {
			continue
		}
		/** 🔵 - Yearn *****************************************************************************
		** This block of code is responsible for retrieving the list of vaults for a given registry
		** It then iterates over these vaults to find the one with the highest block number.
		** This is where we will start indexing new vaults from to avoid spending ressources with
		** data we already have.
		******************************************************************************************/
		_, vaultSlice := storage.ListVaultsFromRegistry(chainID, registry.Address)
		highestBlockNumber := uint64(0)
		for _, strategy := range vaultSlice {
			if strategy.BlockNumber > highestBlockNumber {
				highestBlockNumber = strategy.BlockNumber
			}
		}

		/** 🔵 - Yearn *****************************************************************************
		** After retrieving the highest block number we can proceed to index new vaults.
		******************************************************************************************/
		wg.Add(1)
		go filterNewVault(chainID, registry, highestBlockNumber, nil, &wg, false)

		// Disable watching and use plain RPC call indexing
		// go indexNewVaultsWrapper(chainID, registry, highestBlockNumber, &wg)
	}
	wg.Wait()

	/** 🔵 - Yearn *********************************************************************************
	** This part is only accessed once, once the `historical` vaults, aka that are already deployed
	** and indexed, are retrieved. New deployed vaults will not hit this part, but will be handle
	** in the above functions directly
	**********************************************************************************************/
	vaultsFromRegistry, _ = storage.ListVaultsFromRegistries(chainID)
	storage.StoreRegistriesToJson(chainID, vaultsFromRegistry)
	return vaultsFromRegistry
}
