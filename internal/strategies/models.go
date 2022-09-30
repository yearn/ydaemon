package strategies

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/internal/utils/bigNumber"
)

// TStrategyDetails contains the details about a strategy.
type TStrategyDetails struct {
	Keeper                  common.Address    `json:"keeper"`
	Strategist              common.Address    `json:"strategist"`
	Rewards                 common.Address    `json:"rewards"`
	HealthCheck             common.Address    `json:"healthCheck"`
	TotalDebt               *bigNumber.BigInt `json:"totalDebt"`
	TotalLoss               *bigNumber.BigInt `json:"totalLoss"`
	TotalGain               *bigNumber.BigInt `json:"totalGain"`
	RateLimit               *bigNumber.BigInt `json:"rateLimit,omitempty"`         // Only < 0.3.2
	MinDebtPerHarvest       *bigNumber.BigInt `json:"minDebtPerHarvest,omitempty"` // Only >= 0.3.2
	MaxDebtPerHarvest       *bigNumber.BigInt `json:"maxDebtPerHarvest,omitempty"` // Only >= 0.3.2
	EstimatedTotalAssets    *bigNumber.BigInt `json:"estimatedTotalAssets"`
	CreditAvailable         *bigNumber.BigInt `json:"creditAvailable"`
	DebtOutstanding         *bigNumber.BigInt `json:"debtOutstanding"`
	ExpectedReturn          *bigNumber.BigInt `json:"expectedReturn"`
	DelegatedAssets         *bigNumber.BigInt `json:"delegatedAssets"`
	DelegatedValue          *bigNumber.BigInt `json:"delegatedValue"`
	Protocols               []string          `json:"protocols"`
	Version                 string            `json:"version"`
	APR                     float64           `json:"apr"`
	PerformanceFee          uint64            `json:"performanceFee"`
	LastReport              uint64            `json:"lastReport"`
	Activation              uint64            `json:"activation"`
	KeepCRV                 uint64            `json:"keepCRV"`
	DebtRatio               uint64            `json:"debtRatio,omitempty"` // Only > 0.2.2
	DebtLimit               uint64            `json:"debtLimit"`
	WithdrawalQueuePosition int64             `json:"withdrawalQueuePosition"`
	DoHealthCheck           bool              `json:"doHealthCheck"`
	InQueue                 bool              `json:"inQueue"`
	EmergencyExit           bool              `json:"emergencyExit"`
	IsActive                bool              `json:"isActive"`
}

// TStrategy contains all the information useful about the strategies currently active in this vault.
type TStrategy struct {
	Address         common.Address    `json:"address"`
	Name            string            `json:"name"`
	Description     string            `json:"description"`
	DelegatedAssets *bigNumber.BigInt `json:"-"` //Used for internal computation
	TotalDebt       *bigNumber.BigInt `json:"-"` //Used for internal computation
	DelegatedValue  string            `json:"-"` //Used for internal computation
	InQueue         bool              `json:"-"` //Used for internal computation
	IsActive        bool              `json:"-"` //Used for internal computation
	Details         *TStrategyDetails `json:"details,omitempty"`
	Risk            *TStrategyRisk    `json:"risk,omitempty"`
}
