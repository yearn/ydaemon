package events

import (
	"context"
	"strconv"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/common/contracts"
	"github.com/yearn/ydaemon/common/env"
	"github.com/yearn/ydaemon/common/ethereum"
	"github.com/yearn/ydaemon/common/logs"
	"github.com/yearn/ydaemon/common/store"
	"github.com/yearn/ydaemon/internal/models"
)

/**************************************************************************************************
** Filter only the first UpdateManagement events to know when it was emitted and extract the
** blockNumber.
**
** Arguments:
** - chainID: the chain ID of the network we are working on
** - vaultAddress: the address of the vault we are working on
** - activationMap: the async ptr to the map of vaultAddress -> blockNumber
** - wg: the async ptr to the WaitGroup to sync the goroutines
**
** Returns nothing as the activationMap is updated via a pointer
**************************************************************************************************/
func filterUpdateManagementOneTime(vault models.TVaultsFromRegistry) {
	client := ethereum.GetRPC(vault.ChainID)
	currentVault, _ := contracts.NewYvault043(vault.Address, client)
	/**********************************************************************************************
	** First, we need to know when to stop our log fetching. By default, we will fetch until the
	** current block number, aka nil.
	** As using nil may cause some issues, we will specify the current block number instead.
	**********************************************************************************************/
	blockEnd, _ := client.BlockNumber(context.Background())
	end := &blockEnd
	start := vault.Activation

	/******************************************************************************************
	** Finally, we will fetch the logs in chunks of MAX_BLOCK_RANGE blocks. This is done to
	** avoid hitting some external node providers' rate limits.
	******************************************************************************************/
	for chunkStart := start; chunkStart < *end; chunkStart += env.CHAINS[vault.ChainID].MaxBlockRange {
		chunkEnd := chunkStart + env.CHAINS[vault.ChainID].MaxBlockRange
		if chunkEnd > *end {
			chunkEnd = *end
		}

		opts := &bind.FilterOpts{Start: chunkStart, End: &chunkEnd}
		if log, err := currentVault.FilterUpdateManagement(opts); err == nil {
			if log.Next() {
				if log.Error() != nil {
					return
				}
				store.StoreVaultActivation(vault.ChainID, []store.DBVaultActivation{
					{
						UUID:    `will-be-replaced`,
						ChainID: vault.ChainID,
						Vault:   vault.Address.Hex(),
						Block:   log.Event.Raw.BlockNumber,
					},
				})
			}
		} else {
			logs.Error(`impossible to FilterUpdateManagement for Yvault043 ` + vault.Address.Hex() + ` on chain ` + strconv.FormatUint(vault.ChainID, 10) + `: ` + err.Error())
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
func HandleUpdateManagementOneTime(
	chainID uint64,
	vaults map[common.Address]models.TVaultsFromRegistry,
) map[common.Address]models.TVaultsFromRegistry {
	activationMap := store.ListVaultActivation(chainID)
	/**********************************************************************************************
	** Concurrently retrieve all first updateManagement events, waiting for the end of all
	** goroutines via the wg before continuing.
	**********************************************************************************************/
	for _, v := range vaults {
		if v.Activation != 0 {
			store.StoreVaultActivation(v.ChainID, []store.DBVaultActivation{
				{
					UUID:    `will-be-replaced`,
					ChainID: v.ChainID,
					Vault:   v.Address.Hex(),
					Block:   v.Activation,
				},
			})
			continue
		}
		if _, ok := activationMap[v.Address]; ok {
			if activationMap[v.Address].Uint64() != 0 {
				continue
			}
		}
		filterUpdateManagementOneTime(v)
	}

	activationMap = store.ListVaultActivation(chainID)

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
	for vault, activationBlock := range activationMap {
		if currentVault, ok := vaults[vault]; ok {
			currentVault.Activation = activationBlock.Uint64()
			vaultListWithActivation[vault] = currentVault
			count++
		}
	}

	return vaultListWithActivation
}
