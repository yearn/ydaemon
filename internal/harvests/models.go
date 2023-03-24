package harvests

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/yearn/ydaemon/common/bigNumber"
)

type TStrategyReportBase struct {
	Vault     common.Address
	Strategy  common.Address
	Token     common.Address
	Gain      *big.Int
	Loss      *big.Int
	DebtPaid  *big.Int
	TotalGain *big.Int
	TotalLoss *big.Int
	TotalDebt *big.Int
	DebtAdded *big.Int
	DebtRatio *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

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

	// Extra
	TokenPrice *bigNumber.Int
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

var _allHarvests = make(map[common.Address][]THarvest)

func GetAll() map[common.Address][]THarvest {
	return _allHarvests
}

func Get(vaultAddress common.Address) []THarvest {
	if _, ok := _allHarvests[vaultAddress]; !ok {
		return []THarvest{}
	}
	return _allHarvests[vaultAddress]
}
