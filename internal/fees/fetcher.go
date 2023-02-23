package fees

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
	"github.com/yearn/ydaemon/internal/strategies"
	"github.com/yearn/ydaemon/internal/utils"
	"github.com/yearn/ydaemon/internal/vaults"
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
func filterUpdateManagementFee(
	chainID uint64,
	vaultAddress common.Address,
	vaultActivation uint64,
	asyncFeeMap *sync.Map,
	wg *sync.WaitGroup,
) {
	defer wg.Done()
	client := ethereum.GetRPC(chainID)

	currentVault, _ := contracts.NewYvault043(vaultAddress, client)
	options := bind.FilterOpts{Start: vaultActivation, End: nil}
	if log, err := currentVault.FilterUpdateManagementFee(&options); err == nil {
		for log.Next() {
			if log.Error() != nil {
				continue
			}

			eventKey := vaultAddress.String() + `-` + strconv.FormatUint(uint64(log.Event.Raw.BlockNumber), 10)
			blockData := utils.TEventBlock{
				EventType:   `updateManagementFee`,
				TxHash:      log.Event.Raw.TxHash,
				BlockNumber: log.Event.Raw.BlockNumber,
				TxIndex:     log.Event.Raw.TxIndex,
				LogIndex:    log.Event.Raw.Index,
				Value:       bigNumber.SetInt(log.Event.ManagementFee),
			}

			if syncMap, ok := asyncFeeMap.Load(eventKey); ok {
				currentBlockData := append(syncMap.([]utils.TEventBlock), blockData)
				asyncFeeMap.Store(eventKey, currentBlockData)
			} else {
				asyncFeeMap.Store(eventKey, []utils.TEventBlock{blockData})
			}
		}
	}
}

/**************************************************************************************************
** Filter all updatePerformanceFee events and store them in a map of blockNumber => TEventBlock
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
func filterUpdatePerformanceFee(
	chainID uint64,
	vaultAddress common.Address,
	vaultActivation uint64,
	asyncFeeMap *sync.Map,
	wg *sync.WaitGroup,
) {
	defer wg.Done()
	client := ethereum.GetRPC(chainID)

	currentVault, _ := contracts.NewYvault043(vaultAddress, client)
	if log, err := currentVault.FilterUpdatePerformanceFee(&bind.FilterOpts{Start: vaultActivation}); err == nil {
		for log.Next() {
			if log.Error() != nil {
				continue
			}

			eventKey := vaultAddress.String() + `-` + strconv.FormatUint(uint64(log.Event.Raw.BlockNumber), 10)
			blockData := utils.TEventBlock{
				EventType:   `updatePerformanceFee`,
				TxHash:      log.Event.Raw.TxHash,
				BlockNumber: log.Event.Raw.BlockNumber,
				TxIndex:     log.Event.Raw.TxIndex,
				LogIndex:    log.Event.Raw.Index,
				Value:       bigNumber.SetInt(log.Event.PerformanceFee),
			}

			if syncMap, ok := asyncFeeMap.Load(eventKey); ok {
				currentBlockData := append(syncMap.([]utils.TEventBlock), blockData)
				asyncFeeMap.Store(eventKey, currentBlockData)
			} else {
				asyncFeeMap.Store(eventKey, []utils.TEventBlock{blockData})
			}
		}
	}
}

/**************************************************************************************************
** Filter all updateStrategyPerformanceFee events and store them in a map of blockNumber =>
** TEventBlock
**
** Arguments:
** - chainID: the chain ID of the network we are working on
** - vaultAddress: the address of the vault we are working on
** - vaultActivation: the block number at which the vault was activated
** - asyncFeeMap: the async ptr to the map of vaultAddr -> strategyAddr -> block -> TEventBlock
** - wg: the async ptr to the WaitGroup to sync the goroutines
**
** Returns nothing as the asyncFeeMap is updated via a pointer
**************************************************************************************************/
func filterUpdateStrategyPerformanceFee(
	chainID uint64,
	vaultAddress common.Address,
	vaultActivation uint64,
	asyncFeeMap *sync.Map,
	wg *sync.WaitGroup,
) {
	defer wg.Done()
	client := ethereum.GetRPC(chainID)

	currentVault, _ := contracts.NewYvault043(vaultAddress, client)
	if log, err := currentVault.FilterStrategyUpdatePerformanceFee(&bind.FilterOpts{Start: vaultActivation}, nil); err == nil {
		for log.Next() {
			if log.Error() != nil {
				continue
			}

			blockData := utils.TEventBlock{
				EventType:   "strategyUpdatePerformanceFee",
				TxHash:      log.Event.Raw.TxHash,
				BlockNumber: log.Event.Raw.BlockNumber,
				TxIndex:     log.Event.Raw.TxIndex,
				LogIndex:    log.Event.Raw.Index,
				Value:       bigNumber.SetInt(log.Event.PerformanceFee),
			}

			eventKey := vaultAddress.String() + `-` + log.Event.Strategy.String() + `-` + strconv.FormatUint(uint64(log.Event.Raw.BlockNumber), 10)
			if syncMap, ok := asyncFeeMap.Load(eventKey); ok {
				currentBlockData := append((syncMap.([]utils.TEventBlock)), blockData)
				asyncFeeMap.Store(eventKey, currentBlockData)
			} else {
				asyncFeeMap.Store(eventKey, []utils.TEventBlock{blockData})
			}

		}
	}
}

/**************************************************************************************************
** For each vault we need to know the fee per block, which is the percentage of gains after each
** harvest that will be sent to the governance. This is a dynamic value, and it can be changed
** by the governance. We need to fetch all the events of type `UpdateManagementFee`,
** `UpdateStrategyPerformanceFee` and `UpdatePerformanceFee` and build an historical mapping of
** the fee per block, knowing for each block which fee to use.
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
** - a map of vaultAddress -> blockNumber -> PerformanceFee
** - a map of vaultAddress -> strategyAddress -> blockNumber -> PerformanceFee
**************************************************************************************************/
func RetrieveAllFeesBPS(
	chainID uint64,
	vaults map[common.Address]*vaults.TVault,
	strategiesLists ...map[common.Address]map[common.Address]*strategies.TStrategy,
) (
	map[common.Address]map[uint64][]utils.TEventBlock,
	map[common.Address]map[uint64][]utils.TEventBlock,
	map[common.Address]map[common.Address]map[uint64][]utils.TEventBlock,
) {
	timeBefore := time.Now()

	asyncManagementFeeUpdate := sync.Map{}
	asyncPerformanceFeeUpdate := sync.Map{}
	asyncStrategiesPerformanceFeeUpdate := sync.Map{}

	wg := &sync.WaitGroup{}
	for _, v := range vaults {
		wg.Add(3)
		go filterUpdateManagementFee(chainID, v.Address, v.Activation, &asyncManagementFeeUpdate, wg)
		go filterUpdatePerformanceFee(chainID, v.Address, v.Activation, &asyncPerformanceFeeUpdate, wg)
		go filterUpdateStrategyPerformanceFee(chainID, v.Address, v.Activation, &asyncStrategiesPerformanceFeeUpdate, wg)
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
	managementFeeForVaults := make(map[common.Address]map[uint64][]utils.TEventBlock)
	asyncManagementFeeUpdate.Range(func(key, value interface{}) bool {
		eventKey := strings.Split(key.(string), `-`)
		vaultAddress := common.HexToAddress(eventKey[0])
		blockNumber, _ := strconv.ParseUint(eventKey[1], 10, 64)

		if _, ok := managementFeeForVaults[vaultAddress]; !ok {
			managementFeeForVaults[vaultAddress] = make(map[uint64][]utils.TEventBlock)
		}
		managementFeeForVaults[vaultAddress][blockNumber] = value.([]utils.TEventBlock)
		return true
	})

	/**********************************************************************************************
	** Once all vaults PerformanceFees updates have been retrieved, we need to extract them from
	** the sync.Map.
	**
	** The syncMap variable is setup as follows:
	** - key: vaultAddress-blockNumber
	** - value: []TEventBlock
	**
	** We need to transform it into a map as follows:
	** - vaultAddress -> blockNumber -> []TEventBlock
	**********************************************************************************************/
	performanceFeeForVaults := make(map[common.Address]map[uint64][]utils.TEventBlock)
	asyncPerformanceFeeUpdate.Range(func(key, value interface{}) bool {
		eventKey := strings.Split(key.(string), `-`)
		vaultAddress := common.HexToAddress(eventKey[0])
		blockNumber, _ := strconv.ParseUint(eventKey[1], 10, 64)

		if _, ok := performanceFeeForVaults[vaultAddress]; !ok {
			performanceFeeForVaults[vaultAddress] = make(map[uint64][]utils.TEventBlock)
		}
		performanceFeeForVaults[vaultAddress][blockNumber] = value.([]utils.TEventBlock)
		return true
	})

	/**********************************************************************************************
	** Once all strategies PerformanceFees updates have been retrieved, we need to extract them
	** from the sync.Map.
	**
	** The syncMap variable is setup as follows:
	** - key: vaultAddress-strategyAddress-blockNumber
	** - value: []TEventBlock
	**
	** We need to transform it into a map as follows:
	** - vaultAddress -> strategyAddress -> blockNumber -> []TEventBlock
	**********************************************************************************************/
	performanceFeeForStrategies := make(map[common.Address]map[common.Address]map[uint64][]utils.TEventBlock)
	asyncStrategiesPerformanceFeeUpdate.Range(func(key, value interface{}) bool {
		eventKey := strings.Split(key.(string), `-`)
		vaultAddress := common.HexToAddress(eventKey[0])
		strategyAddress := common.HexToAddress(eventKey[1])
		blockNumber, _ := strconv.ParseUint(eventKey[2], 10, 64)

		// If the mapping for [vaultAddress] doesn't exist, create it
		if _, ok := performanceFeeForStrategies[vaultAddress]; !ok {
			performanceFeeForStrategies[vaultAddress] = make(map[common.Address]map[uint64][]utils.TEventBlock)
		}

		// If the mapping for [vaultAddress][strategyAddress] doesn't exist, create it
		if _, ok := performanceFeeForStrategies[vaultAddress][strategyAddress]; !ok {
			performanceFeeForStrategies[vaultAddress][strategyAddress] = make(map[uint64][]utils.TEventBlock)
		}
		// If the mapping for [vaultAddress][strategyAddress][blockNumber] doesn't exist, create it
		if _, ok := performanceFeeForStrategies[vaultAddress][strategyAddress][blockNumber]; !ok {
			performanceFeeForStrategies[vaultAddress][strategyAddress][blockNumber] = make([]utils.TEventBlock, 0)
		}

		performanceFeeForStrategies[vaultAddress][strategyAddress][blockNumber] = value.([]utils.TEventBlock)
		return true
	})

	/**********************************************************************************************
	** The initial strategy performanceFee does not trigger a StrategyPerformanceFeeUpdated event.
	** Therefore we need to add it from the StrategiesList variable
	**********************************************************************************************/
	if len(strategiesLists) == 1 {
		strategiesList := strategiesLists[0]

		for vaultAddress, strategies := range strategiesList {
			for strategyAddress, strategy := range strategies {
				deployEvent := utils.TEventBlock{
					EventType:   `strategyUpdatePerformanceFee`,
					TxHash:      strategy.Initialization.TxHash,
					BlockNumber: strategy.Initialization.BlockNumber,
					TxIndex:     strategy.Initialization.TxIndex,
					LogIndex:    strategy.Initialization.LogIndex,
					Value:       strategy.PerformanceFee,
				}

				if _, ok := performanceFeeForStrategies[vaultAddress]; !ok {
					performanceFeeForStrategies[vaultAddress] = make(map[common.Address]map[uint64][]utils.TEventBlock)
				}
				if _, ok := performanceFeeForStrategies[vaultAddress][strategyAddress]; !ok {
					performanceFeeForStrategies[vaultAddress][strategyAddress] = make(map[uint64][]utils.TEventBlock)
				}
				if _, ok := performanceFeeForStrategies[vaultAddress][strategyAddress][strategy.Initialization.BlockNumber]; !ok {
					performanceFeeForStrategies[vaultAddress][strategyAddress][strategy.Initialization.BlockNumber] = make([]utils.TEventBlock, 0)
				}
				performanceFeeForStrategies[vaultAddress][strategyAddress][strategy.Initialization.BlockNumber] = append(
					performanceFeeForStrategies[vaultAddress][strategyAddress][strategy.Initialization.BlockNumber],
					deployEvent,
				)
			}
		}
	}

	logs.Success(`It tooks`, time.Since(timeBefore), `to retrieve the fees BPS updates`)
	return managementFeeForVaults, performanceFeeForVaults, performanceFeeForStrategies
}
