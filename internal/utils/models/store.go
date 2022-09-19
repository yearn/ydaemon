package models

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type TStrategyList struct {
	Strategy     common.Address
	Vault        common.Address
	VaultVersion string
	Name         string
}

type TStrategyMultiCallData struct {
	CreditAvailable         *big.Int `json:"creditAvailable"`
	DebtOutstanding         *big.Int `json:"debtOutstanding"`
	ExpectedReturn          *big.Int `json:"expectedReturn"`
	PerformanceFee          *big.Int `json:"performanceFee"`
	Activation              *big.Int `json:"activation"`
	DebtRatio               *big.Int `json:"debtRatio,omitempty"`         // Only > 0.2.2
	DebtLimit               *big.Int `json:"debtLimit,omitempty"`         // Only = 0.2.2
	RateLimit               *big.Int `json:"rateLimit,omitempty"`         // Only < 0.3.2
	MinDebtPerHarvest       *big.Int `json:"minDebtPerHarvest,omitempty"` // Only >= 0.3.2
	MaxDebtPerHarvest       *big.Int `json:"maxDebtPerHarvest,omitempty"` // Only >= 0.3.2
	LastReport              *big.Int `json:"lastReport"`
	TotalDebt               *big.Int `json:"totalDebt"`
	TotalGain               *big.Int `json:"totalGain"`
	TotalLoss               *big.Int `json:"totalLoss"`
	EstimatedTotalAssets    *big.Int `json:"estimatedTotalAssets"`
	KeepCRV                 *big.Int `json:"keepCRV"`
	DelegatedAssets         *big.Int `json:"delegatedAssets"`
	WithdrawalQueuePosition *big.Int `json:"withdrawalQueuePosition"`
	IsActive                bool     `json:"isActive"`
}

type TStrategyFromRiskRiskScores struct {
	TVLImpact           float64 `json:"TVLImpact"`
	AuditScore          float64 `json:"auditScore"`
	CodeReviewScore     float64 `json:"codeReviewScore"`
	ComplexityScore     float64 `json:"complexityScore"`
	LongevityImpact     float64 `json:"longevityImpact"`
	ProtocolSafetyScore float64 `json:"protocolSafetyScore"`
	TeamKnowledgeScore  float64 `json:"teamKnowledgeScore"`
	TestingScore        float64 `json:"testingScore"`
}
type TStrategyFromRisk struct {
	RiskGroup  string                      `json:"riskGroup"`
	RiskScores TStrategyFromRiskRiskScores `json:"riskScores"`
}
