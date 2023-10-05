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
** Filter all updateStrategyPerformanceFee events and store them in a map of blockNumber =>
** TEventBlock
**
** Arguments:
** - chainID: the chain ID of the network we are working on
** - vaultAddress: the address of the vault we are working on
** - vaultActivation: the block number at which the vault was activated
** - perfMap: the async ptr to the map of vaultAddr -> strategyAddr -> block -> TEventBlock
** - wg: the async ptr to the WaitGroup to sync the goroutines
**
** Returns nothing as the perfMap is updated via a pointer
**************************************************************************************************/
func filterUpdateStrategyPerformanceFee(vault models.TVault, start uint64, end *uint64, perfMap *sync.Map) {
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
		if log, err := currentVault.FilterStrategyUpdatePerformanceFee(opts, nil); err == nil {
			for log.Next() {
				if log.Error() != nil {
					continue
				}

				blockData := ethereum.TEventBlock{
					EventType:   "strategyUpdatePerformanceFee",
					TxHash:      log.Event.Raw.TxHash,
					BlockNumber: log.Event.Raw.BlockNumber,
					TxIndex:     log.Event.Raw.TxIndex,
					LogIndex:    log.Event.Raw.Index,
					Value:       bigNumber.SetInt(log.Event.PerformanceFee),
				}

				eventKey := vault.Address.Hex() + `-` + log.Event.Strategy.Hex() + `-` + strconv.FormatUint(uint64(log.Event.Raw.BlockNumber), 10)
				if syncMap, ok := perfMap.Load(eventKey); ok {
					currentBlockData := append((syncMap.([]ethereum.TEventBlock)), blockData)
					perfMap.Store(eventKey, currentBlockData)
				} else {
					perfMap.Store(eventKey, []ethereum.TEventBlock{blockData})
				}
			}
		} else {
			logs.Error(`impossible to FilterStrategyUpdatePerformanceFee for Vault ` + vault.Address.Hex() + ` on chain ` + strconv.FormatUint(vault.ChainID, 10) + `: ` + err.Error())
		}
	}
}

/**************************************************************************************************
** For each vault we need to know the fee per block, which is the percentage of gains after each
** harvest that will be sent to the governance. This is a dynamic value, and it can be changed
** by the governance. We need to fetch all the events of type `UpdateStrategyPerformanceFee` and
** build an historical mapping of the fee per block, knowing for each block which fee to use.
**
** Arguments:
** - chainID: the chain ID of the network we are working on
** - vaults: the list of vaults we want to fetch the fee for, as a mapping of vaultAddress -> data
** - strategiesList: the list of strategies we want to fetch the fee for, as a mapping of
**   vaultAddress -> strategyAddress -> TStrategyAdded. This one is optional to add initial fee to
**   the strategyPerformanceFee map
**
** Returns:
** - a map of vaultAddress -> strategyAddress -> blockNumber -> PerformanceFee
**************************************************************************************************/
func HandleUpdateStrategyPerformanceFee(
	chainID uint64,
	vaults map[common.Address]models.TVault,
	strategiesMap map[common.Address]map[common.Address]*models.TStrategy,
	start uint64,
	end *uint64,
) map[common.Address]map[common.Address]map[uint64][]ethereum.TEventBlock {
	timeBefore := time.Now()
	syncMap := sync.Map{}

	wg := &sync.WaitGroup{}
	for _, v := range vaults {
		wg.Add(1)
		go func(v models.TVault) {
			defer wg.Done()
			filterUpdateStrategyPerformanceFee(v, start, end, &syncMap)
		}(v)
	}
	wg.Wait()

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
	performanceFeeForStrategies := make(map[common.Address]map[common.Address]map[uint64][]ethereum.TEventBlock)
	syncMap.Range(func(key, value interface{}) bool {
		eventKey := strings.Split(key.(string), `-`)
		vaultAddress := common.HexToAddress(eventKey[0])
		strategyAddress := common.HexToAddress(eventKey[1])
		blockNumber, _ := strconv.ParseUint(eventKey[2], 10, 64)

		// If the mapping for [vaultAddress] doesn't exist, create itg
		if _, ok := performanceFeeForStrategies[vaultAddress]; !ok {
			performanceFeeForStrategies[vaultAddress] = make(map[common.Address]map[uint64][]ethereum.TEventBlock)
		}

		// If the mapping for [vaultAddress][strategyAddress] doesn't exist, create it
		if _, ok := performanceFeeForStrategies[vaultAddress][strategyAddress]; !ok {
			performanceFeeForStrategies[vaultAddress][strategyAddress] = make(map[uint64][]ethereum.TEventBlock)
		}
		// If the mapping for [vaultAddress][strategyAddress][blockNumber] doesn't exist, create it
		if _, ok := performanceFeeForStrategies[vaultAddress][strategyAddress][blockNumber]; !ok {
			performanceFeeForStrategies[vaultAddress][strategyAddress][blockNumber] = make([]ethereum.TEventBlock, 0)
		}

		performanceFeeForStrategies[vaultAddress][strategyAddress][blockNumber] = value.([]ethereum.TEventBlock)
		return true
	})

	/**********************************************************************************************
	** The initial strategy performanceFee does not trigger a StrategyPerformanceFeeUpdated event.
	** Therefore we need to add it from the StrategiesList variable
	**********************************************************************************************/
	for vaultAddress, strategies := range strategiesMap {
		for strategyAddress, strategy := range strategies {
			deployEvent := ethereum.TEventBlock{
				EventType:   `strategyUpdatePerformanceFee`,
				TxHash:      strategy.Initialization.TxHash,
				BlockNumber: strategy.Initialization.BlockNumber,
				TxIndex:     strategy.Initialization.TxIndex,
				LogIndex:    strategy.Initialization.LogIndex,
				Value:       strategy.PerformanceFee,
			}

			if _, ok := performanceFeeForStrategies[vaultAddress]; !ok {
				performanceFeeForStrategies[vaultAddress] = make(map[common.Address]map[uint64][]ethereum.TEventBlock)
			}
			if _, ok := performanceFeeForStrategies[vaultAddress][strategyAddress]; !ok {
				performanceFeeForStrategies[vaultAddress][strategyAddress] = make(map[uint64][]ethereum.TEventBlock)
			}
			if _, ok := performanceFeeForStrategies[vaultAddress][strategyAddress][strategy.Initialization.BlockNumber]; !ok {
				performanceFeeForStrategies[vaultAddress][strategyAddress][strategy.Initialization.BlockNumber] = make([]ethereum.TEventBlock, 0)
			}
			performanceFeeForStrategies[vaultAddress][strategyAddress][strategy.Initialization.BlockNumber] = append(
				performanceFeeForStrategies[vaultAddress][strategyAddress][strategy.Initialization.BlockNumber],
				deployEvent,
			)
		}
	}

	logs.Success(`It tooks`, time.Since(timeBefore), `to retrieve the performanceFeeForStrategies updates`)
	return performanceFeeForStrategies
}
