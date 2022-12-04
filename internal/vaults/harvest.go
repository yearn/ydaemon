package vaults

import (
	"sync"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/yearn/ydaemon/common/bigNumber"
	"github.com/yearn/ydaemon/common/contracts"
	"github.com/yearn/ydaemon/common/ethereum"
	"github.com/yearn/ydaemon/internal/utils"
)

type THarvestFees struct {
	ManagementFeeBPS       *bigNumber.Int
	PerformanceFeeBPS      *bigNumber.Int
	StrategistFeeBPS       *bigNumber.Int
	TreasuryCollectedFee   *bigNumber.Int
	StrategistCollectedFee *bigNumber.Int
	TotalCollectedFee      *bigNumber.Int
	TreasuryFeeRatio       *bigNumber.Float
	StrategistFeeRatio     *bigNumber.Float
	TotalFeeRatio          *bigNumber.Float
}
type THarvest struct {
	// Extracted from common types.Log
	TxHash      common.Hash
	BlockHash   common.Hash
	BlockNumber uint64
	Timestamp   uint64
	TxIndex     uint
	LogIndex    uint
	Removed     bool

	// Extracted from Yvault043StrategyReported & computed
	Vault        common.Address
	Strategy     common.Address
	VaultVersion string
	Gain         *bigNumber.Int
	Loss         *bigNumber.Int
	TotalGain    *bigNumber.Int
	TotalLoss    *bigNumber.Int
	TotalDebt    *bigNumber.Int
	DebtLimit    *bigNumber.Int // Only V0.2.2
	DebtPaid     *bigNumber.Int // Only >= 0.3.1
	DebtAdded    *bigNumber.Int
	DebtRatio    *bigNumber.Int

	// Computed
	Duration *bigNumber.Int
	Fees     THarvestFees
}

func (harvest *THarvest) New(log types.Log) *THarvest {
	harvest.TxHash = log.TxHash
	harvest.BlockHash = log.BlockHash
	harvest.BlockNumber = log.BlockNumber
	harvest.TxIndex = log.TxIndex
	harvest.LogIndex = log.Index
	harvest.Removed = log.Removed
	return harvest
}

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

func HandleEvenStrategyReportedFor031To043(
	chainID uint64,
	vault utils.TVaultsFromRegistry,
	managementFeeChanges map[uint64][]utils.TEventBlock,
	performanceFeeChanges map[uint64][]utils.TEventBlock,
	strategiesPerformanceFeeChanges map[common.Address]map[uint64][]utils.TEventBlock,
	transfersFromVaultsToStrategies map[common.Address]map[uint64][]utils.TEventBlock,
	transfersFromVaultsToTreasury map[uint64][]utils.TEventBlock,
	allLastReport map[common.Address]map[uint64]uint64,
	syncGroup *sync.WaitGroup,
	harvests *[]THarvest,
) {
	defer syncGroup.Done()

	client := ethereum.RPC[1]
	currentVault, _ := contracts.NewYvault043(vault.VaultsAddress, client)
	if log, err := currentVault.FilterStrategyReported(&bind.FilterOpts{Start: vault.Activation}, nil); err == nil {
		for log.Next() {
			if log.Error() != nil {
				continue
			}

			currentBlock := utils.TEventBlock{
				BlockNumber: log.Event.Raw.BlockNumber,
				TxIndex:     log.Event.Raw.TxIndex,
				LogIndex:    log.Event.Raw.Index,
			}
			transferToStrategist, transferToTreasury := findRelatedTransfers(log, transfersFromVaultsToStrategies, transfersFromVaultsToTreasury)

			harvest := &THarvest{}
			harvest.New(log.Event.Raw)
			harvest.Timestamp = ethereum.GetBlockTime(chainID, log.Event.Raw.BlockNumber)
			harvest.Vault = vault.VaultsAddress
			harvest.VaultVersion = vault.APIVersion
			harvest.Strategy = log.Event.Strategy
			harvest.Gain = bigNumber.SetInt(log.Event.Gain)
			harvest.Loss = bigNumber.SetInt(log.Event.Loss)
			harvest.TotalGain = bigNumber.SetInt(log.Event.TotalGain)
			harvest.TotalLoss = bigNumber.SetInt(log.Event.TotalLoss)
			harvest.TotalDebt = bigNumber.SetInt(log.Event.TotalDebt)
			harvest.DebtPaid = bigNumber.SetInt(log.Event.DebtPaid)
			harvest.DebtAdded = bigNumber.SetInt(log.Event.DebtAdded)
			harvest.DebtRatio = bigNumber.SetInt(log.Event.DebtRatio)
			harvest.Duration = durationSinceLastReport(log, allLastReport)

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

			*harvests = append(*harvests, *harvest)
		}
	}
}
