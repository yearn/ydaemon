package harvests

import (
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/common/bigNumber"
	"github.com/yearn/ydaemon/common/ethereum"
	"github.com/yearn/ydaemon/common/logs"
	"github.com/yearn/ydaemon/internal/events"
	"github.com/yearn/ydaemon/internal/models"
	"github.com/yearn/ydaemon/internal/prices"
)

func findRelatedTransfers(
	log models.TStrategyReportBase,
	transfersFromVaultsToStrategies map[common.Address]map[uint64][]ethereum.TEventBlock,
	transfersFromVaultsToTreasury map[uint64][]ethereum.TEventBlock,
) (*bigNumber.Int, *bigNumber.Int) {
	currentBlock := ethereum.TEventBlock{
		BlockNumber: log.Raw.BlockNumber,
		TxIndex:     log.Raw.TxIndex,
		LogIndex:    log.Raw.Index,
	}

	transferToStrategist := ethereum.FindEventBefore(
		map[uint64][]ethereum.TEventBlock{
			currentBlock.BlockNumber: transfersFromVaultsToStrategies[log.Strategy][currentBlock.BlockNumber],
		},
		currentBlock,
	)
	transferToTreasury := ethereum.FindEventBefore(
		map[uint64][]ethereum.TEventBlock{
			currentBlock.BlockNumber: transfersFromVaultsToTreasury[currentBlock.BlockNumber],
		},
		currentBlock,
	)

	return transferToStrategist, transferToTreasury
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
	vaults map[common.Address]models.TVault,
	strategiesMap map[common.Address]map[common.Address]*models.TStrategy,
	start uint64,
	end *uint64,
) (
	managementFeeForVaults map[common.Address]map[uint64][]ethereum.TEventBlock,
	performanceFeeForVaults map[common.Address]map[uint64][]ethereum.TEventBlock,
	performanceFeeForStrategies map[common.Address]map[common.Address]map[uint64][]ethereum.TEventBlock,
) {
	wg := &sync.WaitGroup{}
	wg.Add(3)
	go func() {
		defer wg.Done()
		managementFeeForVaults = events.HandleUpdateManagementFee(chainID, vaults, start, end)
	}()
	go func() {
		defer wg.Done()
		performanceFeeForVaults = events.HandleUpdatePerformanceFee(chainID, vaults, start, end)
	}()
	go func() {
		defer wg.Done()
		performanceFeeForStrategies = events.HandleUpdateStrategyPerformanceFee(chainID, vaults, strategiesMap, start, end)
	}()
	wg.Wait()
	return
}

/**************************************************************************************************
** At some point we will need to know the last report for each strategy. This will allow us to
** calculate the APY for one harvest for a given time. To do that we need to know when each harvest
** happened (in time and not blocknumber).
** This function ease the process by returning a map of:
** map -> [vaultAddress][strategyAddress][blockNumber] = [timestamp]
**************************************************************************************************/
func getLastReportsForStrategy(chainID uint64, allEvents []models.TStrategyReportBase) (allReportsTimestamp map[common.Address]map[common.Address]map[uint64]uint64) {
	allReportsTimestamp = map[common.Address]map[common.Address]map[uint64]uint64{}
	for _, event := range allEvents {
		if allReportsTimestamp[event.Vault] == nil {
			allReportsTimestamp[event.Vault] = map[common.Address]map[uint64]uint64{}
		}
		if allReportsTimestamp[event.Vault][event.Strategy] == nil {
			allReportsTimestamp[event.Vault][event.Strategy] = map[uint64]uint64{}
		}
		allReportsTimestamp[event.Vault][event.Strategy][event.Raw.BlockNumber] = ethereum.GetBlockTime(chainID, event.Raw.BlockNumber)
	}
	return allReportsTimestamp
}

/**************************************************************************************************
** getDurationSinceLastReport will return the duration in seconds between the current report and the
** previous one. This is used to calculate the APY for a given harvest.
**************************************************************************************************/
func getDurationSinceLastReport(log models.TStrategyReportBase, allLastReport map[common.Address]map[uint64]uint64) *bigNumber.Int {
	previousBlockTimestampUint64 := ethereum.FindPreviousBlock(allLastReport[log.Strategy], log.Raw.BlockNumber)
	duration := bigNumber.NewInt(0).Sub(
		bigNumber.NewUint64(allLastReport[log.Strategy][log.Raw.BlockNumber]),
		bigNumber.NewUint64(previousBlockTimestampUint64),
	)
	if previousBlockTimestampUint64 == 0 || duration.IsZero() {
		return bigNumber.NewInt(0)
	}
	return duration
}

/**************************************************************************************************
** getTokensPricesAtBlock will fetch the price of all the tokens used in the harvests at a given
** block. This is done in parallel to speed up the process.
**************************************************************************************************/
func getTokensPricesAtBlock(chainID uint64, allEvents []models.TStrategyReportBase, vaultsMap map[common.Address]models.TVault) {
	timeBefore := time.Now()
	blocks := map[uint64][]common.Address{}

	distinctTokens := map[common.Address]uint64{}
	for _, vault := range vaultsMap {
		distinctTokens[vault.Address] = vault.Activation
	}

	for _, event := range allEvents {
		if blocks[event.Raw.BlockNumber] == nil {
			blocks[event.Raw.BlockNumber] = []common.Address{}
		}
		for tokenAddress, activationBlock := range distinctTokens {
			if event.Raw.BlockNumber >= activationBlock {
				blocks[event.Raw.BlockNumber] = append(blocks[event.Raw.BlockNumber], tokenAddress)
			}
		}
		blocks[event.Raw.BlockNumber] = append(blocks[event.Raw.BlockNumber], event.Token)
	}

	logs.Info("Fetching historical prices for", len(blocks), "blocks")

	wg := &sync.WaitGroup{}
	routines := make(chan bool, 100)
	for block, tokens := range blocks {
		wg.Add(1)
		routines <- true
		go func(block uint64, tokens []common.Address) {
			defer wg.Done()
			prices.FetchPricesOnBlock(chainID, block, tokens)
			<-routines
		}(block, tokens)
	}
	wg.Wait()

	logs.Success("It took", time.Since(timeBefore), "to fetch all the prices")
}
