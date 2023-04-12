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
	"github.com/yearn/ydaemon/internal/strategies"
	"github.com/yearn/ydaemon/internal/vaults"
)

/**********************************************************************************************
** RetrieveAllHarvests will fetch all harvests for one specific chain between two blocks and
** return them in a map of vaults -> strategies -> []harvests.
**********************************************************************************************/
func RetrieveAllHarvests(chainID uint64, start uint64, end *uint64) map[common.Address]map[common.Address][]models.THarvest {
	timeBefore := time.Now()

	/**********************************************************************************************
	** All vaults from Yearn are registered in the registries contracts. A vault can be either
	** Standard, Experimental or Automated.
	** From the registries, we are fetching all vaults along with the block in which they were
	** added to the registry, and we remove the duplicates only to keep the latest version of a
	** same vault. Duplicate can happen when a vault is moved from Experimental to Standard.
	**********************************************************************************************/
	transfersFromVaultsToTreasury := map[common.Address]map[uint64][]ethereum.TEventBlock{}
	transfersFromVaultsToStrategies := map[common.Address]map[common.Address]map[uint64][]ethereum.TEventBlock{}
	managementFees := map[common.Address]map[uint64][]ethereum.TEventBlock{}
	performanceFees := map[common.Address]map[uint64][]ethereum.TEventBlock{}
	strategiesPerformanceFees := map[common.Address]map[common.Address]map[uint64][]ethereum.TEventBlock{}
	strategtReportedEvents := []models.TStrategyReportBase{}

	/**********************************************************************************************
	** Fetching all the strategiesList and relevant transfers to proceed
	**********************************************************************************************/
	vaultsMap := vaults.MapVaults(chainID)
	strategiesList := strategies.ListStrategies(chainID)
	strategiesMap := strategies.ToVaultMap(strategiesList)

	wg := sync.WaitGroup{}
	wg.Add(4)
	/**********************************************************************************************
	** Retrieve all transfers from vaults to strategies. This can only happen in one situation: the
	** vault is sending strategist fees to the strategy for them to be taken by the strategist.
	** We need that to be able to calculate the strategist fees as many variable could make the
	** offchain calculation wrong.
	** Thanks to this number, from offchain totalFees calculation, we can deduct the treasury fees
	**********************************************************************************************/
	go func() {
		defer wg.Done()
		transfersFromVaultsToStrategies = events.HandleTransferFromVaultsToStrategies(chainID, strategiesMap, start, end)
	}()

	/**********************************************************************************************
	** For each vault we need to know the fee per block, which is the percentage of gains after each
	** harvest that will be sent to the governance. This is a dynamic value, and it can be changed
	** by the governance. We need to fetch all the events of type `UpdateManagementFee`,
	** `UpdateStrategyPerformanceFee` and `UpdatePerformanceFee` and build an historical mapping of
	** the fee per block, knowing for each block which fee to use.
	** We want them all, from 0 to the end of the block range because we need to know the fee tier
	** for each block.
	**********************************************************************************************/
	go func() {
		defer wg.Done()
		managementFees, performanceFees, strategiesPerformanceFees = retrieveAllFeesBPS(chainID, vaultsMap, strategiesMap, 0, nil)
	}()

	/**********************************************************************************************
	** Retrieve all transfers from vaults to treasury. This can only happen in one situation: the
	** vault is sending managements fees to the treasury after a harvest.
	** We need that to be able to calculate the actual fees as many variable could make the
	** offchain calculation wrong.
	** Before doing that, we need to fetch all the events of type `UpdateRewards` and build an
	** historical mapping of the treasury address per block, knowing for each block which treasury
	** address to use.
	**********************************************************************************************/
	go func() {
		defer wg.Done()
		events.SetUpdateRewards(chainID, events.HandleUpdateRewards(chainID, vaultsMap, 0, nil))
		transfersFromVaultsToTreasury = events.HandleTransfersFromVaultsToTreasury(chainID, vaultsMap, start, end)
	}()

	/**********************************************************************************************
	** Fetching all the strategyReported events to proceed. The result from this function is a
	** mimic of TStrategyReportBase to avoid import circular dependency. We need to convert it to
	** TStrategyReportBase to be able to use it in the next step.
	**********************************************************************************************/
	go func() {
		defer wg.Done()
		strategyReportedEventsMimic := events.HandleStrategyReported(chainID, vaultsMap, start, end)
		strategtReportedEvents = []models.TStrategyReportBase{}
		for _, v := range strategyReportedEventsMimic {
			strategtReportedEvents = append(strategtReportedEvents, models.TStrategyReportBase{
				Vault:     v.Vault,
				Strategy:  v.Strategy,
				Token:     v.Token,
				Gain:      v.Gain,
				Loss:      v.Loss,
				DebtPaid:  v.DebtPaid,
				TotalGain: v.TotalGain,
				TotalLoss: v.TotalLoss,
				TotalDebt: v.TotalDebt,
				DebtAdded: v.DebtAdded,
				DebtRatio: v.DebtRatio,
				Raw:       v.Raw,
			})
		}
	}()
	wg.Wait()

	/**********************************************************************************************
	** We then need to compute the harvests. We need to iterate over all the events and do some
	** calculations and adjustements.
	**********************************************************************************************/
	getTokensPricesAtBlock(chainID, strategtReportedEvents, vaultsMap)

	harvests := processHarvestEvent(
		chainID,
		strategtReportedEvents,
		managementFees,
		performanceFees,
		strategiesPerformanceFees,
		transfersFromVaultsToStrategies,
		transfersFromVaultsToTreasury,
	)

	allHarvestPerStrategyPerVault := map[common.Address]map[common.Address][]models.THarvest{}
	for _, v := range harvests {
		_allHarvests[v.Vault] = append(_allHarvests[v.Vault], v)
		if allHarvestPerStrategyPerVault[v.Vault] == nil {
			allHarvestPerStrategyPerVault[v.Vault] = map[common.Address][]models.THarvest{}
		}
		allHarvestPerStrategyPerVault[v.Vault][v.Strategy] = append(allHarvestPerStrategyPerVault[v.Vault][v.Strategy], v)
	}
	logs.Success("Harvest initialization done in", time.Since(timeBefore), "for", len(harvests), "harvests")

	return allHarvestPerStrategyPerVault
}

/**************************************************************************************************
** processHarvestEvent will iterate over all the events and calculate the APY for each harvest, the
** fees, etc. This will basically build our harvest object so we can use it in some other parts of
** the code.
**************************************************************************************************/
func processHarvestEvent(
	chainID uint64,
	allEvents []models.TStrategyReportBase,
	managementFeeChanges map[common.Address]map[uint64][]ethereum.TEventBlock,
	performanceFeeChanges map[common.Address]map[uint64][]ethereum.TEventBlock,
	strategiesPerformanceFeeChanges map[common.Address]map[common.Address]map[uint64][]ethereum.TEventBlock,
	transfersFromVaultsToStrategies map[common.Address]map[common.Address]map[uint64][]ethereum.TEventBlock,
	transfersFromVaultsToTreasury map[common.Address]map[uint64][]ethereum.TEventBlock,
) []models.THarvest {
	harvests := []models.THarvest{}
	allstrategiesReportedTimestamp := getLastReportsForStrategy(chainID, allEvents)

	for _, formatedLog := range allEvents {
		currentVault, _ := vaults.GetVault(chainID, formatedLog.Vault)
		currentBlock := ethereum.TEventBlock{
			BlockNumber: formatedLog.Raw.BlockNumber,
			TxIndex:     formatedLog.Raw.TxIndex,
			LogIndex:    formatedLog.Raw.Index,
		}
		transferToStrategist, transferToTreasury := findRelatedTransfers(
			formatedLog,
			transfersFromVaultsToStrategies[currentVault.Address],
			transfersFromVaultsToTreasury[currentVault.Address],
		)

		harvest := &models.THarvest{}
		harvest = New(harvest, formatedLog.Raw)
		harvest.Timestamp = ethereum.GetBlockTime(chainID, formatedLog.Raw.BlockNumber)
		harvest.Vault = currentVault.Address
		harvest.VaultName = currentVault.Name
		harvest.VaultVersion = currentVault.Version
		harvest.Strategy = formatedLog.Strategy
		harvest.Gain = bigNumber.SetInt(formatedLog.Gain)
		harvest.Loss = bigNumber.SetInt(formatedLog.Loss)
		harvest.TotalGain = bigNumber.SetInt(formatedLog.TotalGain)
		harvest.TotalLoss = bigNumber.SetInt(formatedLog.TotalLoss)
		harvest.TotalDebt = bigNumber.SetInt(formatedLog.TotalDebt)
		harvest.DebtPaid = bigNumber.SetInt(formatedLog.DebtPaid)
		harvest.DebtAdded = bigNumber.SetInt(formatedLog.DebtAdded)
		harvest.DebtRatio = bigNumber.SetInt(formatedLog.DebtRatio)
		harvest.Duration = getDurationSinceLastReport(formatedLog, allstrategiesReportedTimestamp[currentVault.Address])

		harvest.Fees.ManagementFeeBPS = ethereum.FindEventBefore(managementFeeChanges[currentVault.Address], currentBlock)
		harvest.Fees.PerformanceFeeBPS = ethereum.FindEventBefore(performanceFeeChanges[currentVault.Address], currentBlock)
		harvest.Fees.StrategistFeeBPS = ethereum.FindEventBefore(strategiesPerformanceFeeChanges[currentVault.Address][formatedLog.Strategy], currentBlock)
		harvest.Fees.TreasuryCollectedFee = transferToTreasury
		harvest.Fees.StrategistCollectedFee = transferToStrategist
		harvest.Fees.TotalCollectedFee = bigNumber.NewInt(0).Add(transferToTreasury, transferToStrategist)
		harvest.Fees.TreasuryFeeRatio = bigNumber.NewFloat(0).Div(
			bigNumber.NewFloat(0).SetInt(harvest.Fees.TreasuryCollectedFee),
			bigNumber.NewFloat(0).SetInt(harvest.Gain),
		)
		harvest.Fees.StrategistFeeRatio = bigNumber.NewFloat(0).Div(
			bigNumber.NewFloat(0).SetInt(harvest.Fees.StrategistCollectedFee),
			bigNumber.NewFloat(0).SetInt(harvest.Gain),
		)
		harvest.Fees.TotalFeeRatio = bigNumber.NewFloat(0).Div(
			bigNumber.NewFloat(0).SetInt(harvest.Fees.TotalCollectedFee),
			bigNumber.NewFloat(0).SetInt(harvest.Gain),
		)

		if tokenPrice, ok := prices.GetPricesOnBlock(chainID, harvest.BlockNumber, currentVault.Address); ok {
			harvest.TokenPrice = tokenPrice
		}

		harvests = append(harvests, *harvest)
	}
	return harvests
}
