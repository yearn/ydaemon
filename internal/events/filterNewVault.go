package events

import (
	"strconv"
	"sync"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/common/contracts"
	"github.com/yearn/ydaemon/common/env"
	"github.com/yearn/ydaemon/common/ethereum"
	"github.com/yearn/ydaemon/common/logs"
	"github.com/yearn/ydaemon/internal/utils"
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
	registryAddress common.Address,
	registryVersion uint64,
	registryActivation uint64,
	opts *bind.FilterOpts,
	syncMap *sync.Map,
	wg *sync.WaitGroup,
) {
	if wg != nil {
		defer wg.Done()
	}
	if registryVersion == 3 {
		return //No newExperimentalVault events in registry v3
	}

	client := ethereum.GetRPC(chainID)
	currentVault, _ := contracts.NewYregistryv2(registryAddress, client) //V1 and V2 share the same ABI

	if log, err := currentVault.FilterNewExperimentalVault(opts, nil, nil); err == nil {
		for log.Next() {
			if log.Error() != nil {
				continue
			}
			eventKey := log.Event.Vault.Hex() + `-` + log.Event.Token.Hex() + `-` + log.Event.ApiVersion + `-` + strconv.FormatUint(uint64(log.Event.Raw.BlockNumber), 10)
			syncMap.Store(eventKey, utils.TVaultsFromRegistry{
				RegistryAddress: registryAddress,
				VaultsAddress:   log.Event.Vault,
				TokenAddress:    log.Event.Token,
				APIVersion:      log.Event.ApiVersion,
				BlockNumber:     log.Event.Raw.BlockNumber,
				Activation:      registryActivation,
				ManagementFee:   200,
				BlockHash:       log.Event.Raw.BlockHash,
				TxIndex:         log.Event.Raw.TxIndex,
				LogIndex:        log.Event.Raw.Index,
				Type:            utils.VaultTypeExperimental,
			})
		}
	} else {
		logs.Error(`impossible to FilterNewExperimentalVault for YRegistryV2 ` + registryAddress.Hex() + ` on chain ` + strconv.FormatUint(chainID, 10) + `: ` + err.Error())
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
	registryAddress common.Address,
	registryVersion uint64,
	registryActivation uint64,
	opts *bind.FilterOpts,
	syncMap *sync.Map,
	wg *sync.WaitGroup,
) {
	if wg != nil {
		defer wg.Done()
	}
	client := ethereum.GetRPC(chainID)

	if registryVersion == 1 || registryVersion == 2 {
		currentVault, _ := contracts.NewYregistryv2(registryAddress, client) //V1 and V2 share the same ABI

		if log, err := currentVault.FilterNewVault(opts, nil, nil); err == nil {
			for log.Next() {
				if log.Error() != nil {
					continue
				}
				eventKey := log.Event.Vault.Hex() + `-` + log.Event.Token.Hex() + `-` + log.Event.ApiVersion + `-` + strconv.FormatUint(uint64(log.Event.Raw.BlockNumber), 10)
				syncMap.Store(eventKey, utils.TVaultsFromRegistry{
					RegistryAddress: registryAddress,
					VaultsAddress:   log.Event.Vault,
					TokenAddress:    log.Event.Token,
					APIVersion:      log.Event.ApiVersion,
					BlockNumber:     log.Event.Raw.BlockNumber,
					Activation:      registryActivation,
					ManagementFee:   200,
					BlockHash:       log.Event.Raw.BlockHash,
					TxIndex:         log.Event.Raw.TxIndex,
					LogIndex:        log.Event.Raw.Index,
					Type:            utils.VaultTypeStandard,
				})
			}
		} else {
			logs.Error(`impossible to FilterNewVault for YRegistryV2 ` + registryAddress.Hex() + ` on chain ` + strconv.FormatUint(chainID, 10) + `: ` + err.Error())

		}
	} else if registryVersion == 3 {
		currentVault, _ := contracts.NewYRegistryV3(registryAddress, client) //V3 is not the same
		if log, err := currentVault.FilterNewVault(opts, nil, nil); err == nil {
			for log.Next() {
				if log.Error() != nil {
					continue
				}
				newVault := utils.TVaultsFromRegistry{
					RegistryAddress: registryAddress,
					VaultsAddress:   log.Event.Vault,
					TokenAddress:    log.Event.Token,
					APIVersion:      log.Event.ApiVersion,
					BlockNumber:     log.Event.Raw.BlockNumber,
					Activation:      registryActivation,
					ManagementFee:   200,
					BlockHash:       log.Event.Raw.BlockHash,
					TxIndex:         log.Event.Raw.TxIndex,
					LogIndex:        log.Event.Raw.Index,
					Type:            utils.VaultTypeStandard,
				}
				if log.Event.VaultType.Uint64() == 2 {
					newVault.Type = utils.VaultTypeAutomated
					newVault.ManagementFee = 0
				}
				eventKey := log.Event.Vault.Hex() + `-` + log.Event.Token.Hex() + `-` + log.Event.ApiVersion + `-` + strconv.FormatUint(uint64(log.Event.Raw.BlockNumber), 10)
				syncMap.Store(eventKey, newVault)
			}
		} else {
			logs.Error(`impossible to FilterNewVault for YRegistryV3 ` + registryAddress.Hex() + ` on chain ` + strconv.FormatUint(chainID, 10) + `: ` + err.Error())
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
	start uint64,
	end *uint64,
) (standardVaultList []utils.TVaultsFromRegistry, experimentalVaultList []utils.TVaultsFromRegistry) {
	syncMap := sync.Map{}
	syncMapExperimental := sync.Map{}

	wg := &sync.WaitGroup{}
	for _, registry := range env.YEARN_REGISTRIES[chainID] {
		wg.Add(2)
		opts := &bind.FilterOpts{Start: start, End: end}
		if start == 0 {
			opts = &bind.FilterOpts{Start: registry.Activation, End: end}
		}
		go filterNewVaults(chainID, registry.Address, registry.Version, registry.Activation, opts, &syncMap, wg)
		go filterNewExperimentalVault(chainID, registry.Address, registry.Version, registry.Activation, opts, &syncMapExperimental, wg)
	}
	wg.Wait()

	syncMap.Range(func(_, value interface{}) bool {
		standardVaultList = append(standardVaultList, value.(utils.TVaultsFromRegistry))
		return true
	})
	syncMapExperimental.Range(func(_, value interface{}) bool {
		experimentalVaultList = append(experimentalVaultList, value.(utils.TVaultsFromRegistry))
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
	registryAddress common.Address,
	registryVersion uint64,
	registryActivation uint64,
	start uint64,
	end *uint64,
) (standardVaultList []utils.TVaultsFromRegistry) {
	syncMap := sync.Map{}

	opts := &bind.FilterOpts{Start: start, End: end}
	if start == 0 {
		opts = &bind.FilterOpts{Start: registryActivation, End: end}
	}
	filterNewVaults(chainID, registryAddress, registryVersion, registryActivation, opts, &syncMap, nil)

	syncMap.Range(func(_, value interface{}) bool {
		standardVaultList = append(standardVaultList, value.(utils.TVaultsFromRegistry))
		return true
	})

	return standardVaultList
}
