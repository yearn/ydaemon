package harvests

import (
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/common/bigNumber"
	"github.com/yearn/ydaemon/common/contracts"
	"github.com/yearn/ydaemon/internal/events"
	"github.com/yearn/ydaemon/internal/strategies"
	"github.com/yearn/ydaemon/internal/utils"
	"github.com/yearn/ydaemon/internal/vaults"
)

func findRelatedTransfers(
	log *contracts.Yvault043StrategyReportedIterator,
	transfersFromVaultsToStrategies map[common.Address]map[uint64][]utils.TEventBlock,
	transfersFromVaultsToTreasury map[uint64][]utils.TEventBlock,
) (*bigNumber.Int, *bigNumber.Int) {
	currentBlock := utils.TEventBlock{
		BlockNumber: log.Event.Raw.BlockNumber,
		TxIndex:     log.Event.Raw.TxIndex,
		LogIndex:    log.Event.Raw.Index,
	}

	transferToStrategist := utils.FindEventBefore(
		map[uint64][]utils.TEventBlock{
			currentBlock.BlockNumber: transfersFromVaultsToStrategies[log.Event.Strategy][currentBlock.BlockNumber],
		},
		currentBlock,
	)
	transferToTreasury := utils.FindEventBefore(
		map[uint64][]utils.TEventBlock{
			currentBlock.BlockNumber: transfersFromVaultsToTreasury[currentBlock.BlockNumber],
		},
		currentBlock,
	)

	return transferToStrategist, transferToTreasury
}

func durationSinceLastReport(
	log *contracts.Yvault043StrategyReportedIterator,
	allLastReport map[common.Address]map[uint64]uint64,
) *bigNumber.Int {
	previousBlockTimestampUint64 := utils.FindPreviousBlock(allLastReport[log.Event.Strategy], log.Event.Raw.BlockNumber)
	duration := bigNumber.NewInt(0).Sub(
		bigNumber.NewUint64(allLastReport[log.Event.Strategy][log.Event.Raw.BlockNumber]),
		bigNumber.NewUint64(previousBlockTimestampUint64),
	)
	if previousBlockTimestampUint64 == 0 || duration.IsZero() {
		return bigNumber.NewInt(0)
	}
	return duration
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
func retrieveAllFeesBPS(
	chainID uint64,
	vaults map[common.Address]*vaults.TVault,
	strategiesLists ...map[common.Address]map[common.Address]*strategies.TStrategy,
) (
	managementFeeForVaults map[common.Address]map[uint64][]utils.TEventBlock,
	performanceFeeForVaults map[common.Address]map[uint64][]utils.TEventBlock,
	performanceFeeForStrategies map[common.Address]map[common.Address]map[uint64][]utils.TEventBlock,
) {
	wg := &sync.WaitGroup{}
	wg.Add(3)
	go func() {
		defer wg.Done()
		managementFeeForVaults = events.HandleUpdateManagementFee(
			chainID,
			vaults,
			strategiesLists...,
		)
	}()
	go func() {
		defer wg.Done()
		performanceFeeForVaults = events.HandleUpdatePerformanceFee(
			chainID,
			vaults,
			strategiesLists...,
		)
	}()
	go func() {
		defer wg.Done()
		performanceFeeForStrategies = events.HandleUpdateStrategyPerformanceFee(
			chainID,
			vaults,
			strategiesLists...,
		)
	}()
	wg.Wait()
	return
}
