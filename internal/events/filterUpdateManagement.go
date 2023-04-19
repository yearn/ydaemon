package events

import (
	"strconv"
	"sync"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/common/contracts"
	"github.com/yearn/ydaemon/common/ethereum"
	"github.com/yearn/ydaemon/common/logs"
	"github.com/yearn/ydaemon/internal/models"
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
	vaultActivation uint64,
	asyncActivationMap *sync.Map,
	wg *sync.WaitGroup,
) {
	if wg != nil {
		defer wg.Done()
	}

	client := ethereum.GetRPC(chainID)
	currentVault, _ := contracts.NewYvault043(vaultAddress, client)
	opts := bind.FilterOpts{Start: vaultActivation, End: nil}
	if log, err := currentVault.FilterUpdateManagement(&opts); err == nil {
		if log.Next() {
			if log.Error() != nil {
				asyncActivationMap.Store(vaultAddress, vaultActivation)
				return
			}
			asyncActivationMap.Store(vaultAddress, log.Event.Raw.BlockNumber)
		} else {
			asyncActivationMap.Store(vaultAddress, vaultActivation)
		}
	} else {
		logs.Error(err)
		asyncActivationMap.Store(vaultAddress, vaultActivation)
		logs.Error(`impossible to FilterUpdateManagement for Yvault043 ` + vaultAddress.Hex() + ` on chain ` + strconv.FormatUint(chainID, 10) + `: ` + err.Error())
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
func HandleUpdateManagementOneTime(
	chainID uint64,
	vaults map[common.Address]models.TVaultsFromRegistry,
) map[common.Address]models.TVaultsFromRegistry {
	/**********************************************************************************************
	** Concurrently retrieve all first updateManagement events, waiting for the end of all
	** goroutines via the wg before continuing.
	**********************************************************************************************/
	asyncActivationMap := sync.Map{}
	wg := &sync.WaitGroup{}
	for _, v := range vaults {
		wg.Add(1)
		go filterUpdateManagementOneTime(chainID, v.VaultsAddress, v.Activation, &asyncActivationMap, wg)
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
	vaultListWithActivation := make(map[common.Address]models.TVaultsFromRegistry)
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
