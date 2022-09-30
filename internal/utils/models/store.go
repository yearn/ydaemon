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
	CreditAvailable         *bigNumber.BigInt `json:"creditAvailable"`
	DebtOutstanding         *bigNumber.BigInt `json:"debtOutstanding"`
	ExpectedReturn          *bigNumber.BigInt `json:"expectedReturn"`
	PerformanceFee          *bigNumber.BigInt `json:"performanceFee"`
	Activation              *bigNumber.BigInt `json:"activation"`
	DebtRatio               *bigNumber.BigInt `json:"debtRatio,omitempty"`         // Only > 0.2.2
	DebtLimit               *bigNumber.BigInt `json:"debtLimit,omitempty"`         // Only = 0.2.2
	RateLimit               *bigNumber.BigInt `json:"rateLimit,omitempty"`         // Only < 0.3.2
	MinDebtPerHarvest       *bigNumber.BigInt `json:"minDebtPerHarvest,omitempty"` // Only >= 0.3.2
	MaxDebtPerHarvest       *bigNumber.BigInt `json:"maxDebtPerHarvest,omitempty"` // Only >= 0.3.2
	LastReport              *bigNumber.BigInt `json:"lastReport"`
	TotalDebt               *bigNumber.BigInt `json:"totalDebt"`
	TotalGain               *bigNumber.BigInt `json:"totalGain"`
	TotalLoss               *bigNumber.BigInt `json:"totalLoss"`
	EstimatedTotalAssets    *bigNumber.BigInt `json:"estimatedTotalAssets"`
	KeepCRV                 *bigNumber.BigInt `json:"keepCRV"`
	DelegatedAssets         *bigNumber.BigInt `json:"delegatedAssets"`
	WithdrawalQueuePosition *bigNumber.BigInt `json:"withdrawalQueuePosition"`
	IsActive                bool              `json:"isActive"`
}
