package registries

import (
	"strconv"
	"sync"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/common/contracts"
	"github.com/yearn/ydaemon/common/env"
	"github.com/yearn/ydaemon/common/ethereum"
	"github.com/yearn/ydaemon/common/traces"
	"github.com/yearn/ydaemon/common/types/common"
	"github.com/yearn/ydaemon/internal/utils"
	"github.com/yearn/ydaemon/internal/vaults"
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
	vaultsList *[]utils.TVaultsFromRegistry,
	wg *sync.WaitGroup,
) {
	defer wg.Done()
	if registryVersion == 3 {
		return //No newExperimentalVault events in registry v3
	}

	client := ethereum.RPC[chainID]
	currentVault, _ := contracts.NewYregistryv2(registryAddress.ToAddress(), client) //V1 and V2 share the same ABI

	if log, err := currentVault.FilterNewExperimentalVault(&bind.FilterOpts{Start: registryActivation}, nil, nil); err == nil {
		for log.Next() {
			if log.Error() != nil {
				continue
			}
			*vaultsList = append(*vaultsList, utils.TVaultsFromRegistry{
				RegistryAddress: registryAddress.ToAddress(),
				VaultsAddress:   log.Event.Vault,
				TokenAddress:    log.Event.Token,
				APIVersion:      log.Event.ApiVersion,
				BlockNumber:     log.Event.Raw.BlockNumber,
				Activation:      registryActivation,
				ManagementFee:   200,
				BlockHash:       log.Event.Raw.BlockHash,
				TxIndex:         log.Event.Raw.TxIndex,
				LogIndex:        log.Event.Raw.Index,
				Type:            "Experimental",
			})
		}
	} else {
		traces.
			Capture(`error`, `impossible to FilterNewExperimentalVault for YRegistryV2 `+registryAddress.Hex()).
			SetEntity(`registry`).
			SetExtra(`error`, err.Error()).
			SetTag(`chainID`, strconv.FormatUint(chainID, 10)).
			SetTag(`rpcURI`, ethereum.GetRPCURI(chainID)).
			SetTag(`registryAddress`, registryAddress.Hex()).
			Send()
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
	vaultsList *[]utils.TVaultsFromRegistry,
	wg *sync.WaitGroup,
) {
	defer wg.Done()
	client := ethereum.GetRPC(chainID)

	if registryVersion == 1 || registryVersion == 2 {
		currentVault, _ := contracts.NewYregistryv2(registryAddress.ToAddress(), client) //V1 and V2 share the same ABI
		if log, err := currentVault.FilterNewVault(&bind.FilterOpts{Start: registryActivation}, nil, nil); err == nil {
			for log.Next() {
				if log.Error() != nil {
					continue
				}
				*vaultsList = append(*vaultsList, utils.TVaultsFromRegistry{
					RegistryAddress: registryAddress.ToAddress(),
					VaultsAddress:   log.Event.Vault,
					TokenAddress:    log.Event.Token,
					APIVersion:      log.Event.ApiVersion,
					BlockNumber:     log.Event.Raw.BlockNumber,
					Activation:      registryActivation,
					ManagementFee:   200,
					BlockHash:       log.Event.Raw.BlockHash,
					TxIndex:         log.Event.Raw.TxIndex,
					LogIndex:        log.Event.Raw.Index,
					Type:            "Standard",
				})
			}
		} else {
			traces.
				Capture(`error`, `impossible to FilterNewVault for YRegistryV2 `+registryAddress.Hex()).
				SetEntity(`registry`).
				SetExtra(`error`, err.Error()).
				SetTag(`chainID`, strconv.FormatUint(chainID, 10)).
				SetTag(`rpcURI`, ethereum.GetRPCURI(chainID)).
				SetTag(`registryAddress`, registryAddress.Hex()).
				Send()
		}
	} else if registryVersion == 3 {
		currentVault, _ := contracts.NewYRegistryV3(registryAddress.ToAddress(), client) //V3 is not the same

		if log, err := currentVault.FilterNewVault(&bind.FilterOpts{Start: registryActivation}, nil, nil); err == nil {
			for log.Next() {
				if log.Error() != nil {
					continue
				}
				*vaultsList = append(*vaultsList, utils.TVaultsFromRegistry{
					RegistryAddress: registryAddress.ToAddress(),
					VaultsAddress:   log.Event.Vault,
					TokenAddress:    log.Event.Token,
					APIVersion:      log.Event.ApiVersion,
					BlockNumber:     log.Event.Raw.BlockNumber,
					Activation:      registryActivation,
					ManagementFee:   200,
					BlockHash:       log.Event.Raw.BlockHash,
					TxIndex:         log.Event.Raw.TxIndex,
					LogIndex:        log.Event.Raw.Index,
					Type:            "Standard",
				})
			}
		} else {
			traces.
				Capture(`error`, `impossible to FilterNewVault for YRegistryV3 `+registryAddress.Hex()).
				SetEntity(`registry`).
				SetExtra(`error`, err.Error()).
				SetTag(`chainID`, strconv.FormatUint(chainID, 10)).
				SetTag(`rpcURI`, ethereum.GetRPCURI(chainID)).
				SetTag(`registryAddress`, registryAddress.Hex()).
				Send()
		}
	}
}

/**************************************************************************************************
** For each registry, filter all newExperimentalVault and newVault events and store them in an
** array of TVaultsFromRegistry. This process will allow us to get all the vaults that were ever
** deployed on the network. From that, all the Yearn's data can be retrieved.
**
** Arguments:
** - chainID: the chain ID of the network we are working on
**
** Returns:
** - a map of registryAddress -> TVaultsFromRegistry
**************************************************************************************************/
func RetrieveAllVaults(
	chainID uint64,
) map[ethcommon.Address]utils.TVaultsFromRegistry {
	trace := traces.Init(`app.indexer.registry.new_vaults_events`).
		SetTag(`chainID`, strconv.FormatUint(chainID, 10)).
		SetTag(`rpcURI`, ethereum.GetRPCURI(chainID)).
		SetTag(`entity`, `registries`).
		SetTag(`subsystem`, `daemon`)
	defer trace.Finish()

	vaultsList := []utils.TVaultsFromRegistry{}
	vaultsListExperimental := []utils.TVaultsFromRegistry{}

	wg := &sync.WaitGroup{}
	for _, registry := range env.YEARN_REGISTRIES[chainID] {
		wg.Add(2)
		go filterNewVaults(chainID, registry.Address, registry.Version, registry.Activation, &vaultsList, wg)
		go filterNewExperimentalVault(chainID, registry.Address, registry.Version, registry.Activation, &vaultsListExperimental, wg)
	}
	wg.Wait()

	/**********************************************************************************************
	** Once we got all the vaults, we need to remove the duplicates. This is because some vaults
	** were deployed first as experimental vaults and then as standard vaults. In that case, we
	** keep the standard vault.
	**********************************************************************************************/
	uniqueVaultsList := make(map[ethcommon.Address]utils.TVaultsFromRegistry)
	for _, v := range vaultsList {
		uniqueVaultsList[v.VaultsAddress] = v
	}

	for _, v := range vaultsListExperimental {
		if _, ok := uniqueVaultsList[v.VaultsAddress]; !ok {
			uniqueVaultsList[v.VaultsAddress] = v
		}
	}

	/**********************************************************************************************
	** For some debugging reason, we may want to only work with a subset of the vaults. This is
	** why we have the following code. It will only keep the vaults that are in the list below.
	**********************************************************************************************/
	vaultsToFilter := []string{
		// `0xa354F35829Ae975e850e23e9615b11Da1B3dC4DE`, //yvUSDC v0.4.3
		// `0x19d3364a399d251e894ac732651be8b0e4e85001`, //yvDAI v0.3.0
	}
	if len(vaultsToFilter) > 0 {
		filteredVaultsList := make(map[ethcommon.Address]utils.TVaultsFromRegistry)
		for _, v := range vaultsToFilter {
			filteredVaultsList[ethcommon.HexToAddress(v)] = uniqueVaultsList[ethcommon.HexToAddress(v)]
		}
		uniqueVaultsList = filteredVaultsList
	}

	return vaults.RetrieveActivationForAllVaults(chainID, uniqueVaultsList)
}
