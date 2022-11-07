package strategies

import (
	"github.com/yearn/ydaemon/common/bigNumber"
	"github.com/yearn/ydaemon/common/models"
	"github.com/yearn/ydaemon/common/types/common"
)

type TStore struct {
	// StrategiesFromRisk holds the data for the strategies from the Yearn Risk Framework for each chain.
	StrategiesFromRisk map[uint64]map[common.Address]TStrategyFromRisk

	// WithdrawalQueueMultiCallData holds the details about the withdrawalQueue order based on a multicall
	WithdrawalQueueMultiCallData map[uint64]map[common.Address]int64

	// StrategyGroupFromRisk holds the risk groups from the risk framework
	StrategyGroupFromRisk map[uint64][]*TStrategyGroupFromRisk

	AggregatedStrategies map[uint64]map[common.Address]*TAggregatedStrategy
}

// Store holds the data for the partners for each chain.
var Store = TStore{}

func init() {
	Store.StrategiesFromRisk = make(map[uint64]map[common.Address]TStrategyFromRisk)
	Store.WithdrawalQueueMultiCallData = make(map[uint64]map[common.Address]int64)
	Store.StrategyGroupFromRisk = make(map[uint64][]*TStrategyGroupFromRisk)
	Store.AggregatedStrategies = make(map[uint64]map[common.Address]*TAggregatedStrategy)
}

type TAggregatedStrategy struct {
	Strategy                common.Address
	Vault                   common.Address
	VaultVersion            string
	Name                    string
	CreditAvailable         *bigNumber.Int `json:"creditAvailable"`
	DebtOutstanding         *bigNumber.Int `json:"debtOutstanding"`
	ExpectedReturn          *bigNumber.Int `json:"expectedReturn"`
	PerformanceFee          *bigNumber.Int `json:"performanceFee"`
	Activation              *bigNumber.Int `json:"activation"`
	DebtRatio               *bigNumber.Int `json:"debtRatio,omitempty"`         // Only > 0.2.2
	DebtLimit               *bigNumber.Int `json:"debtLimit,omitempty"`         // Only = 0.2.2
	RateLimit               *bigNumber.Int `json:"rateLimit,omitempty"`         // Only < 0.3.2
	MinDebtPerHarvest       *bigNumber.Int `json:"minDebtPerHarvest,omitempty"` // Only >= 0.3.2
	MaxDebtPerHarvest       *bigNumber.Int `json:"maxDebtPerHarvest,omitempty"` // Only >= 0.3.2
	LastReport              *bigNumber.Int `json:"lastReport"`
	TotalDebt               *bigNumber.Int `json:"totalDebt"`
	TotalGain               *bigNumber.Int `json:"totalGain"`
	TotalLoss               *bigNumber.Int `json:"totalLoss"`
	EstimatedTotalAssets    *bigNumber.Int `json:"estimatedTotalAssets"`
	KeepCRV                 *bigNumber.Int `json:"keepCRV"`
	DelegatedAssets         *bigNumber.Int `json:"delegatedAssets"`
	WithdrawalQueuePosition *bigNumber.Int `json:"withdrawalQueuePosition"`
	IsActive                bool           `json:"isActive"`
}

func (s *TAggregatedStrategy) SetMulticallData(data models.TStrategyMultiCallData) {
	s.CreditAvailable = data.CreditAvailable
	s.DebtOutstanding = data.DebtOutstanding
	s.ExpectedReturn = data.ExpectedReturn
	s.PerformanceFee = data.PerformanceFee
	s.Activation = data.Activation
	s.DebtRatio = data.DebtRatio
	s.DebtLimit = data.DebtLimit
	s.RateLimit = data.RateLimit
	s.MinDebtPerHarvest = data.MinDebtPerHarvest
	s.MaxDebtPerHarvest = data.MaxDebtPerHarvest
	s.LastReport = data.LastReport
	s.TotalDebt = data.TotalDebt
	s.TotalGain = data.TotalGain
	s.TotalLoss = data.TotalLoss
	s.EstimatedTotalAssets = data.EstimatedTotalAssets
	s.KeepCRV = data.KeepCRV
	s.DelegatedAssets = data.DelegatedAssets
	s.WithdrawalQueuePosition = data.WithdrawalQueuePosition
	s.IsActive = data.IsActive
}
