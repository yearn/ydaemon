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
type TStrategyFromRiskRiskProfile struct {
	High   float64 `json:"high"`
	Low    float64 `json:"low"`
	Median float64 `json:"median"`
}
type TStrategyFromRisk struct {
	RiskScores  TStrategyFromRiskRiskScores  `json:"riskScores"`
	RiskProfile TStrategyFromRiskRiskProfile `json:"riskProfile"`
	Tokens      []string                     `json:"tokens"`
}

type TPartnersWrapper struct {
	Name    string `json:"name"`
	Vault   string `json:"vault,omitempty"`
	Wrapper string `json:"wrapper"`
	Type    string `json:"type"`
}

type TPartners struct {
	Name       string             `json:"name"`
	StartBlock uint64             `json:"start_block"`
	Treasury   string             `json:"treasury"`
	Wrappers   []TPartnersWrapper `json:"wrappers"`
}
