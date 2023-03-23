package harvests

import (
	"math/big"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/common/bigNumber"
	"github.com/yearn/ydaemon/common/contracts"
	"github.com/yearn/ydaemon/common/ethereum"
	"github.com/yearn/ydaemon/common/logs"
	"github.com/yearn/ydaemon/internal/events"
	"github.com/yearn/ydaemon/internal/prices"
	"github.com/yearn/ydaemon/internal/strategies"
	"github.com/yearn/ydaemon/internal/utils"
	"github.com/yearn/ydaemon/internal/vaults"
)

var AllHarvests = make(map[common.Address][]THarvest)

func RetrieveAllHarvestPerStrategyPerVault(
	chainID uint64,
	start uint64,
	end *uint64,
) map[common.Address]map[common.Address][]THarvest {
	timeBefore := time.Now()
	/**********************************************************************************************
	** All vaults from Yearn are registered in the registries contracts. A vault can be either
	** Standard, Experimental or Automated.
	** From the registries, we are fetching all vaults along with the block in which they were
	** added to the registry, and we remove the duplicates only to keep the latest version of a
	** same vault. Duplicate can happen when a vault is moved from Experimental to Standard.
	**********************************************************************************************/
	transfersFromVaultsToTreasury := map[common.Address]map[uint64][]utils.TEventBlock{}
	transfersFromVaultsToStrategies := map[common.Address]map[common.Address]map[uint64][]utils.TEventBlock{}
	managementFees := map[common.Address]map[uint64][]utils.TEventBlock{}
	performanceFees := map[common.Address]map[uint64][]utils.TEventBlock{}
	strategiesPerformanceFees := map[common.Address]map[common.Address]map[uint64][]utils.TEventBlock{}
	allHarvests := map[common.Address]map[common.Address]map[uint64]uint64{}

	/**********************************************************************************************
	** Fetching all the strategiesList and relevant transfers to proceed
	**********************************************************************************************/
	vaultsMap := vaults.MapVaults(chainID)
	strategiesList := strategies.ListStrategies(chainID)

	wg := sync.WaitGroup{}
	wg.Add(3)
	go func() {
		defer wg.Done()

		/**********************************************************************************************
		** Retrieve all the strategies ever attached to a vault. This will be used in the next step
		** to retrieve the transfer events for the strategists fees.
		** With this process, we are retrieving the standard blockEvents elements and all the arguments
		** from the `StrategyAdded` event.
		**********************************************************************************************/
		strategiesMap := strategies.SplitStrategiesAddedPerVault(strategiesList)

		/**********************************************************************************************
		** Retrieve all transfers from vaults to strategies. This can only happen in one situation: the
		** vault is sending strategist fees to the strategy for them to be taken by the strategist.
		** We need that to be able to calculate the strategist fees as many variable could make the
		** offchain calculation wrong.
		** Thanks to this number, from offchain totalFees calculation, we can deduct the treasury fees
		**********************************************************************************************/
		transfersFromVaultsToStrategies = events.HandleTransferFromVaultsToStrategies(chainID, strategiesMap, start, end)

		/**********************************************************************************************
		** For each vault we need to know the fee per block, which is the percentage of gains after each
		** harvest that will be sent to the governance. This is a dynamic value, and it can be changed
		** by the governance. We need to fetch all the events of type `UpdateManagementFee`,
		** `UpdateStrategyPerformanceFee` and `UpdatePerformanceFee` and build an historical mapping of
		** the fee per block, knowing for each block which fee to use.
		**********************************************************************************************/
		managementFees, performanceFees, strategiesPerformanceFees = retrieveAllFeesBPS(chainID, vaultsMap, strategiesMap, 0, nil)
	}()

	go func() {
		defer wg.Done()
		/**********************************************************************************************
		** Retrieve all change in the treasury address. This is required in order to ensure that we
		** are not missing any transfer from vaults to the treasury.
		**********************************************************************************************/
		events.StoreUpdateRewards(chainID, events.HandleUpdateRewards(chainID, vaultsMap, 0, nil))

		/**********************************************************************************************
		** Retrieve all transfers from vaults to treasury. This can only happen in one situation: the
		** vault is sending managements fees to the treasury after a harvest.
		** We need that to be able to calculate the actual fees as many variable could make the
		** offchain calculation wrong.
		**********************************************************************************************/
		transfersFromVaultsToTreasury = events.HandleTransfersFromVaultsToTreasury(chainID, vaultsMap, start, end)
	}()

	go func() {
		defer wg.Done()
		/**********************************************************************************************
		** Retrieve all harvest events for a vault. This will enable us to know where to look and to
		** compute the gains, losses and the fees.
		**********************************************************************************************/
		allHarvests = events.HandleStrategyReported(chainID, vaultsMap, start, end)
	}()
	wg.Wait()
	logs.Success("Initialization done in", time.Since(timeBefore))

	syncGroup := &sync.WaitGroup{}
	harvests := []THarvest{}
	for _, vault := range vaultsMap {
		switch vault.Version {
		case `0.2.2`:
			syncGroup.Add(1)
			opts := &bind.FilterOpts{Start: start, End: end}
			if start == 0 {
				opts = &bind.FilterOpts{Start: vault.Activation, End: end}
			}

			go handleEvenStrategyReportedFor022(
				chainID,
				vault,
				managementFees[vault.Address],
				performanceFees[vault.Address],
				strategiesPerformanceFees[vault.Address],
				transfersFromVaultsToStrategies[vault.Address],
				transfersFromVaultsToTreasury[vault.Address],
				allHarvests[vault.Address],
				opts,
				syncGroup,
				&harvests,
			)
		case `0.3.0`:
			syncGroup.Add(1)
			opts := &bind.FilterOpts{Start: start, End: end}
			if start == 0 {
				opts = &bind.FilterOpts{Start: vault.Activation, End: end}
			}

			go handleEvenStrategyReportedFor030(
				chainID,
				vault,
				managementFees[vault.Address],
				performanceFees[vault.Address],
				strategiesPerformanceFees[vault.Address],
				transfersFromVaultsToStrategies[vault.Address],
				transfersFromVaultsToTreasury[vault.Address],
				allHarvests[vault.Address],
				opts,
				syncGroup,
				&harvests,
			)
		default: //case `0.3.1`, `0.3.2`, `0.3.3`, `0.3.4`, `0.3.5`, `0.4.2`, `0.4.3`:
			syncGroup.Add(1)
			opts := &bind.FilterOpts{Start: start, End: end}
			if start == 0 {
				opts = &bind.FilterOpts{Start: vault.Activation, End: end}
			}

			go handleEvenStrategyReportedFor031To043(
				chainID,
				vault,
				managementFees[vault.Address],
				performanceFees[vault.Address],
				strategiesPerformanceFees[vault.Address],
				transfersFromVaultsToStrategies[vault.Address],
				transfersFromVaultsToTreasury[vault.Address],
				allHarvests[vault.Address],
				opts,
				syncGroup,
				&harvests,
			)
		}
	}
	syncGroup.Wait()

	allHarvestPerStrategyPerVault := map[common.Address]map[common.Address][]THarvest{}
	count := 0
	for _, v := range harvests {
		AllHarvests[v.Vault] = append(AllHarvests[v.Vault], v)
		if allHarvestPerStrategyPerVault[v.Vault] == nil {
			allHarvestPerStrategyPerVault[v.Vault] = map[common.Address][]THarvest{}
		}
		allHarvestPerStrategyPerVault[v.Vault][v.Strategy] = append(allHarvestPerStrategyPerVault[v.Vault][v.Strategy], v)
		count++
	}
	return allHarvestPerStrategyPerVault
}

func handleEvenStrategyReportedFor031To043(
	chainID uint64,
	vault *vaults.TVault,
	managementFeeChanges map[uint64][]utils.TEventBlock,
	performanceFeeChanges map[uint64][]utils.TEventBlock,
	strategiesPerformanceFeeChanges map[common.Address]map[uint64][]utils.TEventBlock,
	transfersFromVaultsToStrategies map[common.Address]map[uint64][]utils.TEventBlock,
	transfersFromVaultsToTreasury map[uint64][]utils.TEventBlock,
	allLastReport map[common.Address]map[uint64]uint64,
	opts *bind.FilterOpts,
	syncGroup *sync.WaitGroup,
	harvests *[]THarvest,
) {
	defer syncGroup.Done()

	client := ethereum.RPC[1]
	currentVault, _ := contracts.NewYvault043(vault.Address, client)
	if log, err := currentVault.FilterStrategyReported(opts, nil); err == nil {
		for log.Next() {
			if log.Error() != nil {
				continue
			}
			formatedLogs := TStrategyReportBase{
				Strategy:  log.Event.Strategy,
				Gain:      log.Event.Gain,
				Loss:      log.Event.Loss,
				DebtPaid:  log.Event.DebtPaid,
				TotalGain: log.Event.TotalGain,
				TotalLoss: log.Event.TotalLoss,
				TotalDebt: log.Event.TotalDebt,
				DebtAdded: log.Event.DebtAdded,
				DebtRatio: log.Event.DebtRatio,
				Raw:       log.Event.Raw,
			}

			currentBlock := utils.TEventBlock{
				BlockNumber: log.Event.Raw.BlockNumber,
				TxIndex:     log.Event.Raw.TxIndex,
				LogIndex:    log.Event.Raw.Index,
			}
			transferToStrategist, transferToTreasury := findRelatedTransfers(
				formatedLogs,
				transfersFromVaultsToStrategies,
				transfersFromVaultsToTreasury,
			)

			harvest := &THarvest{}
			harvest.New(log.Event.Raw)
			harvest.Timestamp = ethereum.GetBlockTime(chainID, log.Event.Raw.BlockNumber)
			harvest.Vault = vault.Address
			harvest.VaultName = vault.Name
			harvest.VaultVersion = vault.Version
			harvest.Strategy = log.Event.Strategy
			harvest.Gain = bigNumber.SetInt(log.Event.Gain)
			harvest.Loss = bigNumber.SetInt(log.Event.Loss)
			harvest.TotalGain = bigNumber.SetInt(log.Event.TotalGain)
			harvest.TotalLoss = bigNumber.SetInt(log.Event.TotalLoss)
			harvest.TotalDebt = bigNumber.SetInt(log.Event.TotalDebt)
			harvest.DebtPaid = bigNumber.SetInt(log.Event.DebtPaid)
			harvest.DebtAdded = bigNumber.SetInt(log.Event.DebtAdded)
			harvest.DebtRatio = bigNumber.SetInt(log.Event.DebtRatio)
			harvest.Duration = durationSinceLastReport(formatedLogs, allLastReport)

			harvest.Fees.ManagementFeeBPS = utils.FindEventBefore(managementFeeChanges, currentBlock)
			harvest.Fees.PerformanceFeeBPS = utils.FindEventBefore(performanceFeeChanges, currentBlock)
			harvest.Fees.StrategistFeeBPS = utils.FindEventBefore(strategiesPerformanceFeeChanges[log.Event.Strategy], currentBlock)
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

			if tokenPrice, ok := prices.FindPriceOnBlock(chainID, harvest.BlockNumber, vault.Address); ok {
				harvest.TokenPrice = tokenPrice
			}

			*harvests = append(*harvests, *harvest)
		}
	}
}

func handleEvenStrategyReportedFor030(
	chainID uint64,
	vault *vaults.TVault,
	managementFeeChanges map[uint64][]utils.TEventBlock,
	performanceFeeChanges map[uint64][]utils.TEventBlock,
	strategiesPerformanceFeeChanges map[common.Address]map[uint64][]utils.TEventBlock,
	transfersFromVaultsToStrategies map[common.Address]map[uint64][]utils.TEventBlock,
	transfersFromVaultsToTreasury map[uint64][]utils.TEventBlock,
	allLastReport map[common.Address]map[uint64]uint64,
	opts *bind.FilterOpts,
	syncGroup *sync.WaitGroup,
	harvests *[]THarvest,
) {
	defer syncGroup.Done()

	client := ethereum.RPC[1]
	currentVault, _ := contracts.NewYvault030(vault.Address, client)
	if log, err := currentVault.FilterStrategyReported(opts, nil); err == nil {
		for log.Next() {
			if log.Error() != nil {
				continue
			}
			formatedLogs := TStrategyReportBase{
				Strategy:  log.Event.Strategy,
				Gain:      log.Event.Gain,
				Loss:      log.Event.Loss,
				DebtPaid:  big.NewInt(0),
				TotalGain: log.Event.TotalGain,
				TotalLoss: log.Event.TotalLoss,
				TotalDebt: log.Event.TotalDebt,
				DebtAdded: log.Event.DebtAdded,
				DebtRatio: log.Event.DebtRatio,
				Raw:       log.Event.Raw,
			}

			currentBlock := utils.TEventBlock{
				BlockNumber: log.Event.Raw.BlockNumber,
				TxIndex:     log.Event.Raw.TxIndex,
				LogIndex:    log.Event.Raw.Index,
			}
			transferToStrategist, transferToTreasury := findRelatedTransfers(
				formatedLogs,
				transfersFromVaultsToStrategies,
				transfersFromVaultsToTreasury,
			)

			harvest := &THarvest{}
			harvest.New(log.Event.Raw)
			harvest.Timestamp = ethereum.GetBlockTime(chainID, log.Event.Raw.BlockNumber)
			harvest.Vault = vault.Address
			harvest.VaultName = vault.Name
			harvest.VaultVersion = vault.Version
			harvest.Strategy = log.Event.Strategy
			harvest.Gain = bigNumber.SetInt(log.Event.Gain)
			harvest.Loss = bigNumber.SetInt(log.Event.Loss)
			harvest.TotalGain = bigNumber.SetInt(log.Event.TotalGain)
			harvest.TotalLoss = bigNumber.SetInt(log.Event.TotalLoss)
			harvest.TotalDebt = bigNumber.SetInt(log.Event.TotalDebt)
			harvest.DebtPaid = bigNumber.NewInt(0)
			harvest.DebtAdded = bigNumber.SetInt(log.Event.DebtAdded)
			harvest.DebtRatio = bigNumber.SetInt(log.Event.DebtRatio)
			harvest.Duration = durationSinceLastReport(formatedLogs, allLastReport)

			harvest.Fees.ManagementFeeBPS = utils.FindEventBefore(managementFeeChanges, currentBlock)
			harvest.Fees.PerformanceFeeBPS = utils.FindEventBefore(performanceFeeChanges, currentBlock)
			harvest.Fees.StrategistFeeBPS = utils.FindEventBefore(strategiesPerformanceFeeChanges[log.Event.Strategy], currentBlock)
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

			if tokenPrice, ok := prices.FindPriceOnBlock(chainID, harvest.BlockNumber, vault.Address); ok {
				harvest.TokenPrice = tokenPrice
			}

			*harvests = append(*harvests, *harvest)
		}
	}
}

func handleEvenStrategyReportedFor022(
	chainID uint64,
	vault *vaults.TVault,
	managementFeeChanges map[uint64][]utils.TEventBlock,
	performanceFeeChanges map[uint64][]utils.TEventBlock,
	strategiesPerformanceFeeChanges map[common.Address]map[uint64][]utils.TEventBlock,
	transfersFromVaultsToStrategies map[common.Address]map[uint64][]utils.TEventBlock,
	transfersFromVaultsToTreasury map[uint64][]utils.TEventBlock,
	allLastReport map[common.Address]map[uint64]uint64,
	opts *bind.FilterOpts,
	syncGroup *sync.WaitGroup,
	harvests *[]THarvest,
) {
	defer syncGroup.Done()

	client := ethereum.RPC[1]
	currentVault, _ := contracts.NewYvault022(vault.Address, client)
	if log, err := currentVault.FilterStrategyReported(opts, nil); err == nil {
		for log.Next() {
			if log.Error() != nil {
				continue
			}
			formatedLogs := TStrategyReportBase{
				Strategy:  log.Event.Strategy,
				Gain:      log.Event.Gain,
				Loss:      log.Event.Loss,
				DebtPaid:  big.NewInt(0),
				TotalGain: log.Event.TotalGain,
				TotalLoss: log.Event.TotalLoss,
				TotalDebt: log.Event.TotalDebt,
				DebtAdded: log.Event.DebtAdded,
				DebtRatio: big.NewInt(0),
				Raw:       log.Event.Raw,
			}

			currentBlock := utils.TEventBlock{
				BlockNumber: log.Event.Raw.BlockNumber,
				TxIndex:     log.Event.Raw.TxIndex,
				LogIndex:    log.Event.Raw.Index,
			}
			transferToStrategist, transferToTreasury := findRelatedTransfers(
				formatedLogs,
				transfersFromVaultsToStrategies,
				transfersFromVaultsToTreasury,
			)

			harvest := &THarvest{}
			harvest.New(log.Event.Raw)
			harvest.Timestamp = ethereum.GetBlockTime(chainID, log.Event.Raw.BlockNumber)
			harvest.Vault = vault.Address
			harvest.VaultName = vault.Name
			harvest.VaultVersion = vault.Version
			harvest.Strategy = log.Event.Strategy
			harvest.Gain = bigNumber.SetInt(log.Event.Gain)
			harvest.Loss = bigNumber.SetInt(log.Event.Loss)
			harvest.TotalGain = bigNumber.SetInt(log.Event.TotalGain)
			harvest.TotalLoss = bigNumber.SetInt(log.Event.TotalLoss)
			harvest.TotalDebt = bigNumber.SetInt(log.Event.TotalDebt)
			harvest.DebtPaid = bigNumber.NewInt(0)
			harvest.DebtAdded = bigNumber.SetInt(log.Event.DebtAdded)
			harvest.DebtRatio = bigNumber.NewInt(0)
			harvest.Duration = durationSinceLastReport(formatedLogs, allLastReport)

			harvest.Fees.ManagementFeeBPS = utils.FindEventBefore(managementFeeChanges, currentBlock)
			harvest.Fees.PerformanceFeeBPS = utils.FindEventBefore(performanceFeeChanges, currentBlock)
			harvest.Fees.StrategistFeeBPS = utils.FindEventBefore(strategiesPerformanceFeeChanges[log.Event.Strategy], currentBlock)
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

			if tokenPrice, ok := prices.FindPriceOnBlock(chainID, harvest.BlockNumber, vault.Address); ok {
				harvest.TokenPrice = tokenPrice
			}

			*harvests = append(*harvests, *harvest)
		}
	}
}
