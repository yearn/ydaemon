package vaults

import (
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/common/contracts"
	"github.com/yearn/ydaemon/common/ethereum"
	"github.com/yearn/ydaemon/common/logs"
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
	client := ethereum.RPC[chainID]

	currentVault, err := contracts.NewYvault043(vaultAddress, client)
	if err != nil {
		logs.Error(err)
		return
	}
	if log, err := currentVault.FilterUpdateManagement(&bind.FilterOpts{}); err == nil {
		if log.Next() {
			if log.Error() != nil {
				logs.Error(log.Error())
				asyncActivationMap.Store(vaultAddress, 0)
				return
			}
			asyncActivationMap.Store(vaultAddress, log.Event.Raw.BlockNumber)
		}
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
	timeBefore := time.Now()

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

	logs.Success(`It took`, time.Since(timeBefore), `to retrieve`, count, `activation events`)
	return vaultListWithActivation
}
