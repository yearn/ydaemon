package models

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/internal/utils/bigNumber"
)

type TStrategyList struct {
	Strategy     common.Address
	Vault        common.Address
	VaultVersion string
	Name         string
}

type TStrategyMultiCallData struct {
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
