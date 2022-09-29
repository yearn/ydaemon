package strategies

import "math/big"

// TStrategyGroupCriteria is a helper type for parsing the risk framework json
type TStrategyGroupCritieria struct {
	NameLike   []string `json:"nameLike"`
	Strategies []string `json:"strategies"`
	Exclude    []string `json:"exclude"`
}

type TStrategyGroupAllocation struct {
	CurrentTVL      *big.Float // value in USDC
	AvailableTVL    *big.Float
	CurrentAmount   *big.Float // value in WANT
	AvailableAmount *big.Float
}

type TStrategyAllocation struct {
	CurrentTVL      string `json:"currentTVL"` // value in USDC
	AvailableTVL    string `json:"availableTVL"`
	CurrentAmount   string `json:"currentAmount"` // value in WANT
	AvailableAmount string `json:"availableAmount"`
}

type TStrategyGroupFromRisk struct {
	Label               string                   `json:"label"`
	AuditScore          float64                  `json:"auditScore"`
	CodeReviewScore     float64                  `json:"codeReviewScore"`
	ComplexityScore     float64                  `json:"complexityScore"`
	ProtocolSafetyScore float64                  `json:"protocolSafetyScore"`
	TeamKnowledgeScore  float64                  `json:"teamKnowledgeScore"`
	TestingScore        float64                  `json:"testingScore"`
	Criteria            TStrategyGroupCritieria  `json:"criteria"`
	Allocation          TStrategyGroupAllocation `json:"allocation"`
}

// TStrategyRisk contains the details on the risk about a strategy.
type TStrategyRisk struct {
	RiskGroup           string              `json:"riskGroup"`
	TVLImpact           int                 `json:"TVLImpact"`
	AuditScore          int                 `json:"auditScore"`
	CodeReviewScore     int                 `json:"codeReviewScore"`
	ComplexityScore     int                 `json:"complexityScore"`
	LongevityImpact     int                 `json:"longevityImpact"`
	ProtocolSafetyScore int                 `json:"protocolSafetyScore"`
	TeamKnowledgeScore  int                 `json:"teamKnowledgeScore"`
	TestingScore        int                 `json:"testingScore"`
	Allocation          TStrategyAllocation `json:"allocation"`
}
