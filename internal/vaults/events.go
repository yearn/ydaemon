package vaults

import (
	"strconv"
	"sync"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/common/contracts"
	"github.com/yearn/ydaemon/common/ethereum"
	"github.com/yearn/ydaemon/common/traces"
	"github.com/yearn/ydaemon/internal/utils"
)

/**************************************************************************************************
** Filter only the first UpdateManagement events to know when it was emitted and extract the
** blockNumber.
**
** Arguments:
** - chainID: the chain ID of the network we are working on
** - vaultAddress: the address of the vault we are working on
** - asyncActivationMap: the async ptr to the map of vaultAddress -> blockNumber
** - wg: the async ptr to the WaitGroup to sync the goroutines
**
** Returns nothing as the asyncActivationMap is updated via a pointer
**************************************************************************************************/
func filterUpdateManagementOneTime(
	chainID uint64,
	vaultAddress common.Address,
	asyncActivationMap *sync.Map,
	wg *sync.WaitGroup,
) {
	defer wg.Done()
	client := ethereum.GetRPC(chainID)
	currentVault, _ := contracts.NewYvault043(vaultAddress, client)

	if log, err := currentVault.FilterUpdateManagement(&bind.FilterOpts{}); err == nil {
		if log.Next() {
			if log.Error() != nil {
				asyncActivationMap.Store(vaultAddress, 0)
				return
			}
			asyncActivationMap.Store(vaultAddress, log.Event.Raw.BlockNumber)
		}
	} else {
		traces.
			Capture(`error`, `impossible to FilterUpdateManagement for Yvault043 `+vaultAddress.Hex()).
			SetEntity(`vault`).
			SetExtra(`error`, err.Error()).
			SetTag(`chainID`, strconv.FormatUint(chainID, 10)).
			SetTag(`rpcURI`, ethereum.GetRPCURI(chainID)).
			SetTag(`vaultAddress`, vaultAddress.Hex()).
			Send()
	}
}

/**********************************************************************************************
** The Activation method for the vaults is missleading on a blockchain perspective. It's not
** the blockNumber of the vault creation, but the timestamp of the vault creation.
** In order to get the logs from the vault initialization, we need to get the "Start" block
** number.
** To get it, we are querying one of the first event of the vault, which is an
** "UpdateManagement". Then, we just append the value to our vaultsList structure.
**********************************************************************************************/
func RetrieveActivationForAllVaults(
	chainID uint64,
	vaults map[common.Address]utils.TVaultsFromRegistry,
) map[common.Address]utils.TVaultsFromRegistry {
	trace := traces.Init(`app.indexer.vaults.activation_events`).
		SetTag(`chainID`, strconv.FormatUint(chainID, 10)).
		SetTag(`rpcURI`, ethereum.GetRPCURI(chainID)).
		SetTag(`entity`, `vaults`).
		SetTag(`subsystem`, `daemon`)
	defer trace.Finish()

	/**********************************************************************************************
	** Concurrently retrieve all first updateManagement events, waiting for the end of all
	** goroutines via the wg before continuing.
	**********************************************************************************************/
	asyncActivationMap := sync.Map{}
	wg := &sync.WaitGroup{}
	for _, v := range vaults {
		wg.Add(1)
		go filterUpdateManagementOneTime(chainID, v.VaultsAddress, &asyncActivationMap, wg)
	}
	wg.Wait()

	/**********************************************************************************************
	** Once we got all the activation blocks, we need to extract them from the sync.Map.
	**
	** The syncMap variable is setup as follows:
	** - key: vaultAddress
	** - value: blockNumber
	**
	** We need to update, for each corresponding vault, the activation block to the Activation
	** key.
	**********************************************************************************************/
	count := 0
	vaultListWithActivation := make(map[common.Address]utils.TVaultsFromRegistry)
	asyncActivationMap.Range(func(key, value interface{}) bool {
		if currentVault, ok := vaults[key.(common.Address)]; ok {
			currentVault.Activation = value.(uint64)
			vaultListWithActivation[key.(common.Address)] = currentVault
			count++
		}
		return true
	})

	return vaultListWithActivation
}
