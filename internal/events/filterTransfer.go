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
** Filter all transfer events to retrieve the transfer value and store it in an array of
** TEventBlock
**
** Arguments:
** - chainID: the chain ID of the network we are working on
** - vaultAddress: the address of the vault we are working on
** - fromAddresses: array of FROM addresses to filter on
** - toAddresses: array of TO addresses to filter on
** - activation: the block number at which the strategy was activated
** - asyncMapTransfers: the ptr to the async map to store the TEventBlock
** - wg: the async ptr to the WaitGroup to sync the goroutines
**
** Returns nothing as asyncMapTransfers is updated via a pointer
**************************************************************************************************/
func filterTransfers(
	chainID uint64,
	vaultAddress common.Address,
	fromAddresses []common.Address,
	toAddresses []common.Address,
	startFallback uint64,
	start uint64,
	end *uint64,
	asyncMapTransfers *sync.Map,
	wg *sync.WaitGroup,
) {
	if wg != nil {
		defer wg.Done()
	}

	client := ethereum.GetRPC(chainID)
	currentVault, _ := contracts.NewYvault043(vaultAddress, client)
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
		start = startFallback
	}

	/******************************************************************************************
	** Finally, we will fetch the logs in chunks of MAX_BLOCK_RANGE blocks. This is done to
	** avoid hitting some external node providers' rate limits.
	******************************************************************************************/
	for chunkStart := start; chunkStart < *end; chunkStart += env.MAX_BLOCK_RANGE[chainID] {
		wg.Add(2)
		chunkEnd := chunkStart + env.MAX_BLOCK_RANGE[chainID]
		if chunkEnd > *end {
			chunkEnd = *end
		}

		opts := &bind.FilterOpts{Start: chunkStart, End: &chunkEnd}
		if log, err := currentVault.FilterTransfer(opts, fromAddresses, toAddresses); err == nil {
			for log.Next() {
				if log.Error() != nil {
					continue
				}
				eventKey := (log.Event.Sender.Hex() + `-` +
					log.Event.Receiver.Hex() + `-` +
					strconv.FormatUint(uint64(log.Event.Raw.BlockNumber), 10))

				blockData := ethereum.TEventBlock{
					EventType:   `transfer`,
					TxHash:      log.Event.Raw.TxHash,
					BlockNumber: log.Event.Raw.BlockNumber,
					TxIndex:     log.Event.Raw.TxIndex,
					LogIndex:    log.Event.Raw.Index,
					Value:       bigNumber.SetInt(log.Event.Value),
				}

				if syncMap, ok := asyncMapTransfers.Load(eventKey); ok {
					currentBlockData := append((syncMap.([]ethereum.TEventBlock)), blockData)
					asyncMapTransfers.Store(eventKey, currentBlockData)
				} else {
					asyncMapTransfers.Store(eventKey, []ethereum.TEventBlock{blockData})
				}
			}
		} else {
			logs.Error(`impossible to FilterTransfer for ` + vaultAddress.Hex() + ` on chain ` + strconv.FormatUint(chainID, 10) + `: ` + err.Error())
		}
	}
}

/**********************************************************************************************
** Retrieve all transfers from vaults to strategies. This can only happen in one situation: the
** vault is sending strategist fees to the strategy for them to be taken by the strategist.
** We need that to be able to calculate the strategist fees as many variable could make the
** offchain calculation wrong.
** Thanks to this number, from offchain totalFees calculation, we can deduct the treasury fees.
**
** Arguments:
** - chainID: the chain ID of the network we are working on
** - vaultStrategiesMap: list of all TStrategyAdded to work on
**
** Returns:
** - a map of vaultAddress -> strategyAddress -> blockNumber -> TEventBlock
**********************************************************************************************/
func HandleTransferFromVaultsToStrategies(
	chainID uint64,
	vaultStrategiesMap map[common.Address]map[common.Address]*models.TStrategy,
	start uint64,
	end *uint64,
) map[common.Address]map[common.Address]map[uint64][]ethereum.TEventBlock {
	timeBefore := time.Now()

	/**********************************************************************************************
	** Concurrently retrieve all transfers from vaults to strategies, waiting for the end of all
	** goroutines via the wg before continuing.
	**********************************************************************************************/
	syncMap := &sync.Map{}
	wg := &sync.WaitGroup{}
	for _, vault := range vaultStrategiesMap {
		for _, strategy := range vault {
			wg.Add(1)
			go filterTransfers(
				chainID,
				strategy.VaultAddress,
				[]common.Address{strategy.VaultAddress},
				[]common.Address{strategy.Address},
				strategy.Initialization.BlockNumber,
				start,
				end,
				syncMap,
				wg,
			)
		}
	}
	wg.Wait()

	/**********************************************************************************************
	** Once all transfers from vaults to strategies have been retrieved, we need to extract them
	** from the sync.Map.
	**
	** The syncMap variable is setup as follows:
	** - key: vaultAddress-strategyAddress-blockNumber
	** - value: []TEventBlock
	**********************************************************************************************/
	count := 0
	asyncMapTransfers := make(map[common.Address]map[common.Address]map[uint64][]ethereum.TEventBlock)
	syncMap.Range(func(key, value interface{}) bool {
		eventKey := strings.Split(key.(string), `-`)
		vaultAddress := common.HexToAddress(eventKey[0])
		strategyAddress := common.HexToAddress(eventKey[1])
		blockNumber, _ := strconv.ParseUint(eventKey[2], 10, 64)

		// If the mapping for [vaultAddress] doesn't exist, create it
		if _, ok := asyncMapTransfers[vaultAddress]; !ok {
			asyncMapTransfers[vaultAddress] = make(map[common.Address]map[uint64][]ethereum.TEventBlock)
		}

		// If the mapping for [vaultAddress][strategyAddress] doesn't exist, create it
		if _, ok := asyncMapTransfers[vaultAddress][strategyAddress]; !ok {
			asyncMapTransfers[vaultAddress][strategyAddress] = make(map[uint64][]ethereum.TEventBlock)
		}

		// If the mapping for [vaultAddress][strategyAddress][blockNumber] doesn't exist, create it
		if _, ok := asyncMapTransfers[vaultAddress][strategyAddress][blockNumber]; !ok {
			asyncMapTransfers[vaultAddress][strategyAddress][blockNumber] = make([]ethereum.TEventBlock, 0)
		}

		asyncMapTransfers[vaultAddress][strategyAddress][blockNumber] = append(
			asyncMapTransfers[vaultAddress][strategyAddress][blockNumber],
			value.([]ethereum.TEventBlock)...,
		)
		count++
		return true
	})

	logs.Success(`It tooks`, time.Since(timeBefore), `to retrieve`, count, `transfers from vaults to strategies`)
	return asyncMapTransfers
}

/**********************************************************************************************
** Retrieve all transfers from vaults to treasury. This can only happen in one situation: the
** vault is sending managements fees to the treasury after a harvest.
** We need that to be able to calculate the actual fees as many variable could make the
** offchain calculation wrong.
**
** Arguments:
** - chainID: the chain ID of the network we are working on
** - vaults: list of all the vaults to work on
**
** Returns:
** - a map of vaultAddress -> blockNumber -> TEventBlock
**********************************************************************************************/
func HandleTransfersFromVaultsToTreasury(
	chainID uint64,
	vaultsMap map[common.Address]*models.TVault,
	start uint64,
	end *uint64,
) map[common.Address]map[uint64][]ethereum.TEventBlock {
	timeBefore := time.Now()

	/**********************************************************************************************
	** Concurrently retrieve all transfers from vaults to strategies, waiting for the end of all
	** goroutines via the syncGroup before continuing.
	**********************************************************************************************/
	syncMap := &sync.Map{}
	wg := &sync.WaitGroup{}
	for _, vault := range vaultsMap {
		wg.Add(1)
		go filterTransfers(
			chainID,
			vault.Address,
			[]common.Address{vault.Address},
			FindTreasuriesForVault(chainID, vault.Address),
			vault.Activation,
			start,
			end,
			syncMap,
			wg,
		)
	}
	wg.Wait()

	/**********************************************************************************************
	** Once all transfers from vaults to treasury have been retrieved, we need to extract them
	** from the sync.Map.
	**
	** The syncMap variable is setup as follows:
	** - key: vaultAddress-blockNumber
	** - value: []TEventBlock
	**********************************************************************************************/
	count := 0
	transfersFromVaultsToTreasury := make(map[common.Address]map[uint64][]ethereum.TEventBlock)
	syncMap.Range(func(key, value interface{}) bool {
		eventKey := strings.Split(key.(string), `-`)
		senderAddress := common.HexToAddress(eventKey[0])
		treasuryAddress := common.HexToAddress(eventKey[1])
		blockNumber, _ := strconv.ParseUint(eventKey[2], 10, 64)

		//We need to check if treasuryAddress was actually the treasury at this block
		if !IsTreasuryAtBlock(chainID, treasuryAddress, blockNumber) {
			return true
		}

		// If the mapping for [senderAddress] doesn't exist, create it
		if _, ok := transfersFromVaultsToTreasury[senderAddress]; !ok {
			transfersFromVaultsToTreasury[senderAddress] = make(map[uint64][]ethereum.TEventBlock)
		}

		// If the mapping for [senderAddress][blockNumber] doesn't exist, create it
		if _, ok := transfersFromVaultsToTreasury[senderAddress][blockNumber]; !ok {
			transfersFromVaultsToTreasury[senderAddress][blockNumber] = make([]ethereum.TEventBlock, 0)
		}

		transfersFromVaultsToTreasury[senderAddress][blockNumber] = append(
			transfersFromVaultsToTreasury[senderAddress][blockNumber],
			value.([]ethereum.TEventBlock)...,
		)
		count++
		return true
	})

	logs.Success(`It tooks`, time.Since(timeBefore), `to retrieve`, count, `transfers from vaults to treasury`)
	return transfersFromVaultsToTreasury
}
