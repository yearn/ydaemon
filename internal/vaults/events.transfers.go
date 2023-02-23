package vaults

import (
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/common/bigNumber"
	"github.com/yearn/ydaemon/common/contracts"
	"github.com/yearn/ydaemon/common/ethereum"
	"github.com/yearn/ydaemon/common/logs"
	"github.com/yearn/ydaemon/common/traces"
	"github.com/yearn/ydaemon/internal/strategies"
	"github.com/yearn/ydaemon/internal/utils"
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
	activation uint64,
	asyncMapTransfers *sync.Map,
	wg *sync.WaitGroup,
) {
	defer wg.Done()

	client := ethereum.GetRPC(chainID)
	currentVault, _ := contracts.NewYvault043(vaultAddress, client)
	if log, err := currentVault.FilterTransfer(&bind.FilterOpts{Start: activation}, fromAddresses, toAddresses); err == nil {
		for log.Next() {
			if log.Error() != nil {
				continue
			}
			eventKey := (log.Event.Sender.String() + `-` +
				log.Event.Receiver.String() + `-` +
				strconv.FormatUint(uint64(log.Event.Raw.BlockNumber), 10))

			blockData := utils.TEventBlock{
				EventType:   `transfer`,
				TxHash:      log.Event.Raw.TxHash,
				BlockNumber: log.Event.Raw.BlockNumber,
				TxIndex:     log.Event.Raw.TxIndex,
				LogIndex:    log.Event.Raw.Index,
				Value:       bigNumber.SetInt(log.Event.Value),
			}

			if syncMap, ok := asyncMapTransfers.Load(eventKey); ok {
				currentBlockData := append((syncMap.([]utils.TEventBlock)), blockData)
				asyncMapTransfers.Store(eventKey, currentBlockData)
			} else {
				asyncMapTransfers.Store(eventKey, []utils.TEventBlock{blockData})
			}
		}
	} else {
		logs.Error(err.Error())
		traces.
			Capture(`error`, `impossible to FilterTransfer for Yvault043 `+vaultAddress.Hex()).
			SetEntity(`vault`).
			SetExtra(`error`, err.Error()).
			SetTag(`chainID`, strconv.FormatUint(chainID, 10)).
			SetTag(`rpcURI`, ethereum.GetRPCURI(chainID)).
			SetTag(`vaultAddress`, vaultAddress.Hex()).
			Send()
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
** - strategies: list of all TStrategyAdded to work on
**
** Returns:
** - a map of vaultAddress -> strategyAddress -> blockNumber -> TEventBlock
**********************************************************************************************/
func RetrieveAllTransferFromVaultsToStrategies(
	chainID uint64,
	strategies map[common.Address]map[common.Address]*strategies.TStrategy,
) map[common.Address]map[common.Address]map[uint64][]utils.TEventBlock {
	timeBefore := time.Now()

	/**********************************************************************************************
	** Concurrently retrieve all transfers from vaults to strategies, waiting for the end of all
	** goroutines via the wg before continuing.
	**********************************************************************************************/
	syncMap := &sync.Map{}
	wg := &sync.WaitGroup{}
	for _, vault := range strategies {
		for _, strategy := range vault {
			wg.Add(1)
			go filterTransfers(
				chainID,
				strategy.VaultAddress,
				[]common.Address{strategy.VaultAddress},
				[]common.Address{strategy.Address},
				strategy.Initialization.BlockNumber,
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
	asyncMapTransfers := make(map[common.Address]map[common.Address]map[uint64][]utils.TEventBlock)
	syncMap.Range(func(key, value interface{}) bool {
		eventKey := strings.Split(key.(string), `-`)
		vaultAddress := common.HexToAddress(eventKey[0])
		strategyAddress := common.HexToAddress(eventKey[1])
		blockNumber, _ := strconv.ParseUint(eventKey[2], 10, 64)

		// If the mapping for [vaultAddress] doesn't exist, create it
		if _, ok := asyncMapTransfers[vaultAddress]; !ok {
			asyncMapTransfers[vaultAddress] = make(map[common.Address]map[uint64][]utils.TEventBlock)
		}

		// If the mapping for [vaultAddress][strategyAddress] doesn't exist, create it
		if _, ok := asyncMapTransfers[vaultAddress][strategyAddress]; !ok {
			asyncMapTransfers[vaultAddress][strategyAddress] = make(map[uint64][]utils.TEventBlock)
		}

		// If the mapping for [vaultAddress][strategyAddress][blockNumber] doesn't exist, create it
		if _, ok := asyncMapTransfers[vaultAddress][strategyAddress][blockNumber]; !ok {
			asyncMapTransfers[vaultAddress][strategyAddress][blockNumber] = make([]utils.TEventBlock, 0)
		}

		asyncMapTransfers[vaultAddress][strategyAddress][blockNumber] = append(
			asyncMapTransfers[vaultAddress][strategyAddress][blockNumber],
			value.([]utils.TEventBlock)...,
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
func RetrieveAllTransferFromVaultsToTreasury(
	chainID uint64,
	vaults map[common.Address]*TVault,
) map[common.Address]map[uint64][]utils.TEventBlock {
	timeBefore := time.Now()

	/**********************************************************************************************
	** Concurrently retrieve all transfers from vaults to strategies, waiting for the end of all
	** goroutines via the syncGroup before continuing.
	**********************************************************************************************/
	treasury := common.HexToAddress(`0x93a62da5a14c80f265dabc077fcee437b1a0efde`)
	logs.Warning(`using generic treasury address`)

	syncMap := &sync.Map{}
	wg := &sync.WaitGroup{}
	for _, vault := range vaults {
		wg.Add(1)
		go filterTransfers(
			chainID,
			vault.Address,
			[]common.Address{vault.Address},
			[]common.Address{treasury}, //TODO: add vaults that have a different treasury address
			vault.Activation,
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
	transfersFromVaultsToStrategies := make(map[common.Address]map[uint64][]utils.TEventBlock)
	syncMap.Range(func(key, value interface{}) bool {
		eventKey := strings.Split(key.(string), `-`)
		senderAddress := common.HexToAddress(eventKey[0])
		blockNumber, _ := strconv.ParseUint(eventKey[2], 10, 64)

		// If the mapping for [senderAddress] doesn't exist, create it
		if _, ok := transfersFromVaultsToStrategies[senderAddress]; !ok {
			transfersFromVaultsToStrategies[senderAddress] = make(map[uint64][]utils.TEventBlock)
		}

		// If the mapping for [senderAddress][blockNumber] doesn't exist, create it
		if _, ok := transfersFromVaultsToStrategies[senderAddress][blockNumber]; !ok {
			transfersFromVaultsToStrategies[senderAddress][blockNumber] = make([]utils.TEventBlock, 0)
		}

		transfersFromVaultsToStrategies[senderAddress][blockNumber] = append(
			transfersFromVaultsToStrategies[senderAddress][blockNumber],
			value.([]utils.TEventBlock)...,
		)
		count++
		return true
	})

	logs.Success(`It tooks`, time.Since(timeBefore), `to retrieve`, count, `transfers from vaults to treasury`)
	return transfersFromVaultsToStrategies
}
