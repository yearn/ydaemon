package events

import (
	"context"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/common/bigNumber"
	"github.com/yearn/ydaemon/common/contracts"
	"github.com/yearn/ydaemon/common/env"
	"github.com/yearn/ydaemon/common/ethereum"
	"github.com/yearn/ydaemon/common/logs"
	"github.com/yearn/ydaemon/internal/models"
)

/**************************************************************************************************
** Filter all updateManagementFee events and store them in a map of blockNumber => TEventBlock
**
** Arguments:
** - chainID: the chain ID of the network we are working on
** - vaultAddress: the address of the vault we are working on
** - vaultActivation: the block number at which the vault was activated
** - asyncFeeMap: the async ptr to the map of vaultAddress -> blockNumber -> TEventBlock
** - wg: the async ptr to the WaitGroup to sync the goroutines
**
** Returns nothing as the asyncFeeMap is updated via a pointer
**************************************************************************************************/
func filterUpdateManagementFee(vault models.TVault, start uint64, end *uint64, asyncFeeMap *sync.Map) {
	client := ethereum.GetRPC(vault.ChainID)
	currentVault, _ := contracts.NewYvault043(vault.Address, client)

	/**********************************************************************************************
	** First, we need to know when to stop our log fetching. By default, we will fetch until the
	** current block number, aka nil.
	** As using nil may cause some issues, we will specify the current block number instead.
	**********************************************************************************************/
	if end == nil {
		blockEnd, _ := client.BlockNumber(context.Background())
		end = &blockEnd
	}

	/******************************************************************************************
	** Then, we need to know when to start our log fetching. By default, we will fetch from the
	** activation block in order to get all the vaults that were ever deployed since it was
	** deployed.
	******************************************************************************************/
	if start == 0 {
		start = vault.Activation
	}

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
		if log, err := currentVault.FilterUpdateManagementFee(opts); err == nil {
			for log.Next() {
				if log.Error() != nil {
					continue
				}

				eventKey := vault.Address.Hex() + `-` + strconv.FormatUint(uint64(log.Event.Raw.BlockNumber), 10)
				blockData := ethereum.TEventBlock{
					EventType:   `updateManagementFee`,
					TxHash:      log.Event.Raw.TxHash,
					BlockNumber: log.Event.Raw.BlockNumber,
					TxIndex:     log.Event.Raw.TxIndex,
					LogIndex:    log.Event.Raw.Index,
					Value:       bigNumber.SetInt(log.Event.ManagementFee),
				}

				if syncMap, ok := asyncFeeMap.Load(eventKey); ok {
					currentBlockData := append(syncMap.([]ethereum.TEventBlock), blockData)
					asyncFeeMap.Store(eventKey, currentBlockData)
				} else {
					asyncFeeMap.Store(eventKey, []ethereum.TEventBlock{blockData})
				}
			}
		} else {
			logs.Error(`impossible to FilterUpdateManagementFee for Yvault043 ` + vault.Address.Hex() + ` on chain ` + strconv.FormatUint(vault.ChainID, 10) + `: ` + err.Error())

		}
	}
}

/**************************************************************************************************
** For each vault we need to know the fee per block, which is the percentage of gains after each
** harvest that will be sent to the governance. This is a dynamic value, and it can be changed
** by the governance. We need to fetch all the events of type `UpdateManagementFee` and build an
** historical mapping of the fee per block, knowing for each block which fee to use.
**
** Arguments:
** - chainID: the chain ID of the network we are working on
** - vaults: the list of vaults we want to fetch the fee for, as a mapping of vaultAddress -> data
** - strategiesList: the list of strategies we want to fetch the fee for, as a mapping of
**   vaultAddress -> strategyAddress -> TStrategyAdded. This one is optional to add initial fee to
**   the strategyPerformanceFee map
**
** Returns:
** - a map of vaultAddress -> blockNumber -> ManagementFee
**************************************************************************************************/
func HandleUpdateManagementFee(
	chainID uint64,
	vaults map[common.Address]models.TVault,
	start uint64,
	end *uint64,
) map[common.Address]map[uint64][]ethereum.TEventBlock {
	timeBefore := time.Now()
	asyncManagementFeeUpdate := sync.Map{}

	wg := &sync.WaitGroup{}
	for _, v := range vaults {
		wg.Add(1)
		go func(v models.TVault) {
			defer wg.Done()
			filterUpdateManagementFee(
				v,
				start,
				end,
				&asyncManagementFeeUpdate,
			)
		}(v)
	}
	wg.Wait()

	/**********************************************************************************************
	** Once all vaults ManagementFees updates have been retrieved, we need to extract them from the
	** sync.Map.
	**
	** The syncMap variable is setup as follows:
	** - key: vaultAddress-blockNumber
	** - value: []TEventBlock
	**
	** We need to transform it into a map as follows:
	** - vaultAddress -> blockNumber -> []TEventBlock
	**********************************************************************************************/
	managementFeeForVaults := make(map[common.Address]map[uint64][]ethereum.TEventBlock)
	asyncManagementFeeUpdate.Range(func(key, value interface{}) bool {
		eventKey := strings.Split(key.(string), `-`)
		vaultAddress := common.HexToAddress(eventKey[0])
		blockNumber, _ := strconv.ParseUint(eventKey[1], 10, 64)

		if _, ok := managementFeeForVaults[vaultAddress]; !ok {
			managementFeeForVaults[vaultAddress] = make(map[uint64][]ethereum.TEventBlock)
		}
		managementFeeForVaults[vaultAddress][blockNumber] = value.([]ethereum.TEventBlock)
		return true
	})

	logs.Success(`It tooks`, time.Since(timeBefore), `to retrieve the managementFee updates`)
	return managementFeeForVaults
}
