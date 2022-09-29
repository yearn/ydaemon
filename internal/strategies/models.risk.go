package strategies

import "math/big"

// TStrategyGroupCriteria is a helper type for parsing the risk framework json
type TStrategyGroupCritieria struct {
	NameLike   []string `json:"nameLike"`
	Strategies []string `json:"strategies"`
	Exclude    []string `json:"exclude"`
}

type TStrategyGroupAllocation struct {
	CurrentTVL    *big.Float `json:"CurrentTVL"` // value in WANT
	AvailableTVL  *big.Float `json:"AvailableTVL"`
	CurrentUSDC   *big.Float `json:"CurrentUSDC"` // value in USDC
	AvailableUSDC *big.Float `json:"AvailableUSDC"`
}

type TStrategyAllocation struct {
	CurrentTVL    string `json:"currentTVL"` // value in WANT
	AvailableTVL  string `json:"availableTVL"`
	CurrentUSDC   string `json:"currentUSDC"` // value in USDC
	AvailableUSDC string `json:"availableUSDC"`
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
