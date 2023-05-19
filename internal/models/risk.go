package models

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/common/bigNumber"
)

// TStrategyGroupCritieria is a helper type for parsing the risk framework json
type TStrategyGroupCritieria struct {
	Strategies []common.Address `json:"strategies"`
	Exclude    []string         `json:"exclude"`
	NameLike   []string         `json:"nameLike"`
}

type TStrategyAllocation struct {
	Status          string           `json:"status"`
	CurrentTVL      *bigNumber.Float `json:"currentTVL"` // value in USDC
	AvailableTVL    *bigNumber.Float `json:"availableTVL"`
	CurrentAmount   *bigNumber.Float `json:"currentAmount"` // value in WANT
	AvailableAmount *bigNumber.Float `json:"availableAmount"`
}

type TStrategyGroupFromRisk struct {
	Label                  string                  `json:"label"`
	AuditScore             int                     `json:"auditScore"`
	CodeReviewScore        int                     `json:"codeReviewScore"`
	ComplexityScore        int                     `json:"complexityScore"`
	ProtocolSafetyScore    int                     `json:"protocolSafetyScore"`
	TeamKnowledgeScore     int                     `json:"teamKnowledgeScore"`
	TestingScore           int                     `json:"testingScore"`
	StackingRewardTVLScore int                     `json:"stackingRewardTVLScore,omitempty"`
	ChainID                uint64                  `json:"chainID"`
	Criteria               TStrategyGroupCritieria `json:"criteria"`
	Allocation             *TStrategyAllocation    `json:"allocation"`
}

type TStrategyFromRiskRiskDetails struct {
	TVLImpact              int `json:"TVLImpact"`
	AuditScore             int `json:"auditScore"`
	CodeReviewScore        int `json:"codeReviewScore"`
	ComplexityScore        int `json:"complexityScore"`
	LongevityImpact        int `json:"longevityImpact"`
	ProtocolSafetyScore    int `json:"protocolSafetyScore"`
	TeamKnowledgeScore     int `json:"teamKnowledgeScore"`
	TestingScore           int `json:"testingScore"`
	StackingRewardTVLScore int `json:"stackingRewardTVLScore,omitempty"`
}

type TStrategyFromRisk struct {
	Address     common.Address               `json:"address"`
	ChainID     uint64                       `json:"chainID"`
	RiskGroup   string                       `json:"riskGroup"`
	RiskScore   float64                      `json:"riskScore"`
	RiskDetails TStrategyFromRiskRiskDetails `json:"riskDetails"`
	Allocation  *TStrategyAllocation         `json:"allocation"`
}
