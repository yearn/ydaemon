package registries

import (
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/common/contracts"
	"github.com/yearn/ydaemon/common/ethereum"
	"github.com/yearn/ydaemon/common/logs"
	"github.com/yearn/ydaemon/internal/utils"
	"github.com/yearn/ydaemon/internal/vaults"
)

type TRegistry struct {
	Address    common.Address
	ChainID    uint64
	Activation uint64
}

var REGISTRIES = map[uint64][]TRegistry{
	1: {
		{Address: common.HexToAddress("0xe15461b18ee31b7379019dc523231c57d1cbc18c"), ChainID: 1, Activation: 11563389},
		{Address: common.HexToAddress("0x50c1a2eA0a861A967D9d0FFE2AE4012c2E053804"), ChainID: 1, Activation: 12045555},
	},
}

/**************************************************************************************************
** Filter all newExperimentalVault events and store them in an array of TVaultsFromRegistry
**
** Arguments:
** - chainID: the chain ID of the network we are working on
** - vaultAddress: the address of the vault we are working on
** - registryActivation: the block number at which the registry was activated
** - vaultsList: the ptr to the array of TVaultsFromRegistry
** - wg: the async ptr to the WaitGroup to sync the goroutines
**
** Returns nothing as vaultsList is updated via a pointer
**************************************************************************************************/
func filterNewExperimentalVault(
	chainID uint64,
	vaultAddress common.Address,
	registryActivation uint64,
	vaultsList *[]utils.TVaultsFromRegistry,
	wg *sync.WaitGroup,
) {
	defer wg.Done()
	client := ethereum.RPC[chainID]

	currentVault, err := contracts.NewYregistryv2(vaultAddress, client) //V1 and V2 share the same ABI
	if err != nil {
		logs.Error(err)
		return
	}

	if log, err := currentVault.FilterNewExperimentalVault(&bind.FilterOpts{Start: registryActivation}, nil, nil); err == nil {
		for log.Next() {
			if log.Error() != nil {
				continue
			}
			*vaultsList = append(*vaultsList, utils.TVaultsFromRegistry{
				RegistryAddress: vaultAddress,
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
	}
}

/**************************************************************************************************
** Filter all newVault events and store them in an array of TVaultsFromRegistry
**
** Arguments:
** - chainID: the chain ID of the network we are working on
** - vaultAddress: the address of the vault we are working on
** - registryActivation: the block number at which the registry was activated
** - vaultsList: the ptr to the array of TVaultsFromRegistry
** - wg: the async ptr to the WaitGroup to sync the goroutines
**
** Returns nothing as vaultsList is updated via a pointer
**************************************************************************************************/
func filterNewVaults(
	chainID uint64,
	vaultAddress common.Address,
	registryActivation uint64,
	vaultsList *[]utils.TVaultsFromRegistry,
	wg *sync.WaitGroup,
) {
	defer wg.Done()
	client := ethereum.RPC[chainID]

	currentVault, err := contracts.NewYregistryv2(vaultAddress, client) //V1 and V2 share the same ABI
	if err != nil {
		logs.Error(err)
		return
	}

	if log, err := currentVault.FilterNewVault(&bind.FilterOpts{Start: registryActivation}, nil, nil); err == nil {
		for log.Next() {
			if log.Error() != nil {
				continue
			}
			*vaultsList = append(*vaultsList, utils.TVaultsFromRegistry{
				RegistryAddress: vaultAddress,
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
** - a map of vaultAddress -> TVaultsFromRegistry
**************************************************************************************************/
func RetrieveAllVaults(
	chainID uint64,
) map[common.Address]utils.TVaultsFromRegistry {
	timeBefore := time.Now()

	vaultsList := []utils.TVaultsFromRegistry{}
	vaultsListExperimental := []utils.TVaultsFromRegistry{}

	wg := &sync.WaitGroup{}
	for _, registry := range REGISTRIES[chainID] {
		wg.Add(2)
		go filterNewVaults(chainID, registry.Address, registry.Activation, &vaultsList, wg)
		go filterNewExperimentalVault(chainID, registry.Address, registry.Activation, &vaultsListExperimental, wg)
	}
	wg.Wait()

	/**********************************************************************************************
	** Once we got all the vaults, we need to remove the duplicates. This is because some vaults
	** were deployed first as experimental vaults and then as standard vaults. In that case, we
	** keep the standard vault.
	**********************************************************************************************/
	uniqueVaultsList := make(map[common.Address]utils.TVaultsFromRegistry)
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
		filteredVaultsList := make(map[common.Address]utils.TVaultsFromRegistry)
		for _, v := range vaultsToFilter {
			filteredVaultsList[common.HexToAddress(v)] = uniqueVaultsList[common.HexToAddress(v)]
		}
		uniqueVaultsList = filteredVaultsList
	}

	logs.Success(`It took`, time.Since(timeBefore), `to retrieve`, len(uniqueVaultsList), `vaults`)
	return vaults.RetrieveActivationForAllVaults(chainID, uniqueVaultsList)
}
