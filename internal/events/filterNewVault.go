package events

import (
	"context"
	"strconv"
	"sync"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/yearn/ydaemon/common/contracts"
	"github.com/yearn/ydaemon/common/env"
	"github.com/yearn/ydaemon/common/ethereum"
	"github.com/yearn/ydaemon/common/logs"
	"github.com/yearn/ydaemon/common/store"
	"github.com/yearn/ydaemon/internal/models"
)

/**************************************************************************************************
** Filter all newExperimentalVault events and store them in an array of TVaultsFromRegistry
**
** Arguments:
** - chainID: the chain ID of the network we are working on
** - registryAddress: the address of the vault we are working on
** - registryVersion: the version of the registry we are working on
** - registryActivation: the block number at which the registry was activated
** - vaultsList: the ptr to the array of TVaultsFromRegistry
** - wg: the async ptr to the WaitGroup to sync the goroutines
**
** Returns nothing as vaultsList is updated via a pointer
**************************************************************************************************/
func filterNewExperimentalVault(
	chainID uint64,
	registry env.TContractData,
	opts *bind.FilterOpts,
	syncMap *sync.Map,
) {
	if registry.Version == 3 {
		return //No newExperimentalVault events in registry v3
	}

	client := ethereum.GetRPC(chainID)
	currentRegistry, _ := contracts.NewYregistryv2(registry.Address, client) //V1 and V2 share the same ABI

	if log, err := currentRegistry.FilterNewExperimentalVault(opts, nil, nil); err == nil {
		for log.Next() {
			if log.Error() != nil {
				continue
			}
			eventKey := log.Event.Vault.Hex() + `-` + log.Event.Token.Hex() + `-` + log.Event.ApiVersion + `-` + strconv.FormatUint(uint64(log.Event.Raw.BlockNumber), 10)
			newVault := models.TVaultsFromRegistry{
				ChainID:         chainID,
				RegistryAddress: registry.Address,
				Address:         log.Event.Vault,
				TokenAddress:    log.Event.Token,
				APIVersion:      log.Event.ApiVersion,
				BlockNumber:     log.Event.Raw.BlockNumber,
				Activation:      0,
				ManagementFee:   200,
				BlockHash:       log.Event.Raw.BlockHash,
				TxIndex:         log.Event.Raw.TxIndex,
				LogIndex:        log.Event.Raw.Index,
				Type:            models.VaultTypeExperimental,
			}
			syncMap.Store(eventKey, newVault)
		}
	} else {
		logs.Error(`impossible to FilterNewExperimentalVault for YRegistryV2 ` + registry.Address.Hex() + ` on chain ` + strconv.FormatUint(chainID, 10) + `: ` + err.Error())
	}
}

/**************************************************************************************************
** Filter all newVault events and store them in an array of TVaultsFromRegistry
**
** Arguments:
** - chainID: the chain ID of the network we are working on
** - registryAddress: the address of the vault we are working on
** - registryVersion: the version of the registry we are working on
** - registryActivation: the block number at which the registry was activated
** - vaultsList: the ptr to the array of TVaultsFromRegistry
** - wg: the async ptr to the WaitGroup to sync the goroutines
**
** Returns nothing as vaultsList is updated via a pointer
**************************************************************************************************/
func filterNewVaults(
	chainID uint64,
	registry env.TContractData,
	opts *bind.FilterOpts,
	syncMap *sync.Map,
) {
	client := ethereum.GetRPC(chainID)

	if registry.Version == 1 || registry.Version == 2 {
		currentVault, _ := contracts.NewYregistryv2(registry.Address, client) //V1 and V2 share the same ABI
		if log, err := currentVault.FilterNewVault(opts, nil, nil); err == nil {
			for log.Next() {
				if log.Error() != nil {
					continue
				}
				eventKey := log.Event.Vault.Hex() + `-` + log.Event.Token.Hex() + `-` + log.Event.ApiVersion + `-` + strconv.FormatUint(uint64(log.Event.Raw.BlockNumber), 10)
				newVault := models.TVaultsFromRegistry{
					ChainID:         chainID,
					RegistryAddress: registry.Address,
					Address:         log.Event.Vault,
					TokenAddress:    log.Event.Token,
					APIVersion:      log.Event.ApiVersion,
					BlockNumber:     log.Event.Raw.BlockNumber,
					Activation:      0,
					ManagementFee:   200,
					BlockHash:       log.Event.Raw.BlockHash,
					TxIndex:         log.Event.Raw.TxIndex,
					LogIndex:        log.Event.Raw.Index,
					Type:            models.VaultTypeStandard,
				}
				syncMap.Store(eventKey, newVault)
			}
		} else {
			logs.Error(`impossible to FilterNewVault for YRegistryV2 ` + registry.Address.Hex() + ` on chain ` + strconv.FormatUint(chainID, 10) + `: ` + err.Error())
		}
	} else if registry.Version == 3 {
		currentVault, _ := contracts.NewYRegistryV3(registry.Address, client) //V3 is not the same
		if log, err := currentVault.FilterNewVault(opts, nil, nil); err == nil {
			for log.Next() {
				if log.Error() != nil {
					continue
				}
				newVault := models.TVaultsFromRegistry{
					ChainID:         chainID,
					RegistryAddress: registry.Address,
					Address:         log.Event.Vault,
					TokenAddress:    log.Event.Token,
					APIVersion:      log.Event.ApiVersion,
					BlockNumber:     log.Event.Raw.BlockNumber,
					Activation:      0,
					ManagementFee:   200,
					BlockHash:       log.Event.Raw.BlockHash,
					TxIndex:         log.Event.Raw.TxIndex,
					LogIndex:        log.Event.Raw.Index,
					Type:            models.VaultTypeStandard,
				}
				if log.Event.VaultType.Uint64() == 2 {
					newVault.Type = models.VaultTypeAutomated
					newVault.ManagementFee = 0
				}
				eventKey := log.Event.Vault.Hex() + `-` + log.Event.Token.Hex() + `-` + log.Event.ApiVersion + `-` + strconv.FormatUint(uint64(log.Event.Raw.BlockNumber), 10)
				syncMap.Store(eventKey, newVault)
			}
		} else {
			logs.Error(`impossible to FilterNewVault for YRegistryV3 ` + registry.Address.Hex() + ` on chain ` + strconv.FormatUint(chainID, 10) + `: ` + err.Error())
		}
	}
}

/**************************************************************************************************
** HandleNewVaults will, for each registry, filter all newExperimentalVault and newVault events and
** store them in an array of TVaultsFromRegistry. This process will allow us to get all the vaults
** that were ever deployed on the network. From that, all the Yearn's data can be retrieved.
**
** Arguments:
** - chainID: the chain ID of the network we are working on
** - start: the block number to start fetching from
** - end: the block number to stop fetching from
**
** Returns:
** - an array of TVaultsFromRegistry for the standard vaults
** - an array of TVaultsFromRegistry for the experimental vaults
**************************************************************************************************/
func HandleNewVaults(
	chainID uint64,
	registriesWithLastEvent map[string]uint64,
	start uint64,
	end *uint64,
) (standardVaultList []models.TVaultsFromRegistry, experimentalVaultList []models.TVaultsFromRegistry) {
	syncMap := sync.Map{}
	syncMapExperimental := sync.Map{}

	/**********************************************************************************************
	** First, we need to know when to stop our log fetching. By default, we will fetch until the
	** current block number, aka nil.
	** As using nil may cause some issues, we will specify the current block number instead.
	**********************************************************************************************/
	if end == nil {
		client := ethereum.GetRPC(chainID)
		blockEnd, _ := client.BlockNumber(context.Background())
		end = &blockEnd
	}

	wg := &sync.WaitGroup{}
	for _, registry := range env.YEARN_REGISTRIES[chainID] {
		_start := start

		/******************************************************************************************
		** Then, we need to know when to start our log fetching. By default, we will fetch from the
		** activation block in order to get all the vaults that were ever deployed since it was
		** deployed.
		** However, if the external database is used, we may want to only fetch the logs that were
		** emitted after the last time we fetched the logs. In that case, we will use the last
		** event block number + 1 as the start block number.
		** If another start block number is specified, we will use it instead.
		******************************************************************************************/
		if _start == 0 {
			lastEvent := registriesWithLastEvent[registry.Address.Hex()]
			if lastEvent > 0 {
				_start = lastEvent + 1 //Adding one to get the next event
			} else {
				_start = registry.Block
			}
		}
		/******************************************************************************************
		** Finally, we will fetch the logs in chunks of MAX_BLOCK_RANGE blocks. This is done to
		** avoid hitting some external node providers' rate limits.
		******************************************************************************************/
		for chunkStart := _start; chunkStart < *end; chunkStart += env.MAX_BLOCK_RANGE[chainID] {
			wg.Add(2)
			chunkEnd := chunkStart + env.MAX_BLOCK_RANGE[chainID]
			if chunkEnd > *end {
				chunkEnd = *end
			}

			opts := &bind.FilterOpts{Start: chunkStart, End: &chunkEnd}
			go func(_registry env.TContractData) {
				defer wg.Done()
				filterNewVaults(chainID, _registry, opts, &syncMap)
			}(registry)

			go func(_registry env.TContractData) {
				defer wg.Done()
				filterNewExperimentalVault(chainID, _registry, opts, &syncMapExperimental)
			}(registry)
		}
		wg.Wait()

		/**********************************************************************************************
		** We are storing in the DB the sync status, indicating we went up to the block number to check
		** for new vaults.
		**********************************************************************************************/
		go store.DATABASE.Table("db_registry_syncs").
			Where("chain_id = ? AND registry = ?", chainID, registry.Address.Hex()).
			Where("block <= ?", end).
			Assign(store.DBRegistrySync{Block: *end}).
			FirstOrCreate(&store.DBRegistrySync{
				ChainID:  chainID,
				Registry: registry.Address.Hex(),
				UUID:     store.GetUUID(registry.Address.Hex() + strconv.FormatUint(chainID, 10)),
			})
	}

	/**********************************************************************************************
	** Once we have all the logs, we will dump them from the sync.Map to the array of
	** TVaultsFromRegistry to use them in the rest of the process.
	**********************************************************************************************/
	syncMap.Range(func(_, value interface{}) bool {
		standardVaultList = append(standardVaultList, value.(models.TVaultsFromRegistry))
		return true
	})
	syncMapExperimental.Range(func(_, value interface{}) bool {
		experimentalVaultList = append(experimentalVaultList, value.(models.TVaultsFromRegistry))
		return true
	})

	return standardVaultList, experimentalVaultList
}

/**************************************************************************************************
** HandleNewStandardVaults will, for one registry, filter all newVault events and store them in an
** array of TVaultsFromRegistry. This process will allow us to get all the vaults that were ever
** deployed on the network. From that, all the Yearn's data can be retrieved.
**
** Arguments:
** - chainID: the chain ID of the network we are working on
** - start: the block number to start fetching from
** - end: the block number to stop fetching from
**
** Returns:
** - an array of TVaultsFromRegistry
**************************************************************************************************/
func HandleNewStandardVaults(
	chainID uint64,
	registry env.TContractData,
	start uint64,
	end *uint64,
) (standardVaultList []models.TVaultsFromRegistry) {
	syncMap := sync.Map{}

	/**********************************************************************************************
	** First, we need to know when to stop our log fetching. By default, we will fetch until the
	** current block number, aka nil.
	** As using nil may cause some issues, we will specify the current block number instead.
	**********************************************************************************************/
	if end == nil {
		client := ethereum.GetRPC(chainID)
		blockEnd, _ := client.BlockNumber(context.Background())
		end = &blockEnd
	}

	/**********************************************************************************************
	** Then, we need to know when to start our log fetching. By default, we will fetch from the
	** registry activation block in order to get all the vaults that were ever deployed since the
	** registry was deployed.
	** However, if the external database is used, we may want to only fetch the logs that were
	** emitted after the last time we fetched the logs. In that case, we will use the last event
	** block number + 1 as the start block number.
	** If another start block number is specified, we will use it instead.
	**********************************************************************************************/
	registriesWithLastEvent := store.ListLastNewVaultEventForRegistries(chainID)
	if start == 0 {
		lastEvent := registriesWithLastEvent[registry.Address.Hex()]
		if lastEvent > 0 {
			start = lastEvent + 1 //Adding one to get the next event
		} else {
			start = registry.Activation
		}
	}

	/**********************************************************************************************
	** Finally, we will fetch the logs in chunks of MAX_BLOCK_RANGE blocks. This is done to avoid
	** hitting some external node providers' rate limits.
	**********************************************************************************************/
	for chunkStart := start; chunkStart < *end; chunkStart += env.MAX_BLOCK_RANGE[chainID] {
		chunkEnd := chunkStart + env.MAX_BLOCK_RANGE[chainID]
		if chunkEnd > *end {
			chunkEnd = *end
		}

		opts := &bind.FilterOpts{Start: chunkStart, End: &chunkEnd}
		filterNewVaults(chainID, registry, opts, &syncMap)
	}

	/**********************************************************************************************
	** Once we have all the logs, we will dump them from the sync.Map to the array of
	** TVaultsFromRegistry to use them in the rest of the process.
	**********************************************************************************************/
	syncMap.Range(func(_, value interface{}) bool {
		standardVaultList = append(standardVaultList, value.(models.TVaultsFromRegistry))
		return true
	})

	/**********************************************************************************************
	** We are storing in the DB the sync status, indicating we went up to the block number to check
	** for new vaults.
	**********************************************************************************************/
	go store.DATABASE.Table("db_registry_syncs").
		Where("chain_id = ? AND registry = ?", chainID, registry.Address.Hex()).
		Where("block <= ?", end).
		Assign(store.DBRegistrySync{Block: *end}).
		FirstOrCreate(&store.DBRegistrySync{
			ChainID:  chainID,
			Registry: registry.Address.Hex(),
			UUID:     store.GetUUID(registry.Address.Hex() + strconv.FormatUint(chainID, 10)),
		})

	return standardVaultList
}
