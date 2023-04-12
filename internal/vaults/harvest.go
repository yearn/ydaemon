package vaults

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/yearn/ydaemon/common/bigNumber"
	"github.com/yearn/ydaemon/common/contracts"
	"github.com/yearn/ydaemon/common/ethereum"
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
	VaultName    string
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
	transfersFromVaultsToStrategies map[common.Address]map[uint64][]ethereum.TEventBlock,
	transfersFromVaultsToTreasury map[uint64][]ethereum.TEventBlock,
) (*bigNumber.Int, *bigNumber.Int) {
	currentBlock := ethereum.TEventBlock{
		BlockNumber: log.Event.Raw.BlockNumber,
		TxIndex:     log.Event.Raw.TxIndex,
		LogIndex:    log.Event.Raw.Index,
	}

	transferToStrategist := ethereum.FindEventBefore(
		map[uint64][]ethereum.TEventBlock{
			currentBlock.BlockNumber: transfersFromVaultsToStrategies[log.Event.Strategy][currentBlock.BlockNumber],
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

func durationSinceLastReport(
	log *contracts.Yvault043StrategyReportedIterator,
	allLastReport map[common.Address]map[uint64]uint64,
) *bigNumber.Int {
	previousBlockTimestampUint64 := ethereum.FindPreviousBlock(allLastReport[log.Event.Strategy], log.Event.Raw.BlockNumber)
	duration := bigNumber.NewInt(0).Sub(
		bigNumber.NewUint64(allLastReport[log.Event.Strategy][log.Event.Raw.BlockNumber]),
		bigNumber.NewUint64(previousBlockTimestampUint64),
	)
	if previousBlockTimestampUint64 == 0 || duration.IsZero() {
		return bigNumber.NewInt(0)
	}
	return duration
}
