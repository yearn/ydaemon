package models

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type TStrategyList struct {
	Strategy     common.Address
	Vault        common.Address
	VaultVersion string
}

type TStrategyMultiCallData struct {
	CreditAvailable      *big.Int `json:"creditAvailable"`
	DebtOutstanding      *big.Int `json:"debtOutstanding"`
	ExpectedReturn       *big.Int `json:"expectedReturn"`
	PerformanceFee       *big.Int `json:"performanceFee"`
	Activation           *big.Int `json:"activation"`
	DebtRatio            *big.Int `json:"debtRatio,omitempty"`         // Only > 0.2.2
	DebtLimit            *big.Int `json:"debtLimit,omitempty"`         // Only = 0.2.2
	RateLimit            *big.Int `json:"rateLimit,omitempty"`         // Only < 0.3.2
	MinDebtPerHarvest    *big.Int `json:"minDebtPerHarvest,omitempty"` // Only >= 0.3.2
	MaxDebtPerHarvest    *big.Int `json:"maxDebtPerHarvest,omitempty"` // Only >= 0.3.2
	LastReport           *big.Int `json:"lastReport"`
	TotalDebt            *big.Int `json:"totalDebt"`
	TotalGain            *big.Int `json:"totalGain"`
	TotalLoss            *big.Int `json:"totalLoss"`
	EstimatedTotalAssets *big.Int `json:"estimatedTotalAssets"`
	KeepCRV              *big.Int `json:"keepCRV"`
	DelegatedAssets      *big.Int `json:"delegatedAssets"`
	IsActive             bool     `json:"isActive"`
}
