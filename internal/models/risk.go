package models

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/common/bigNumber"
)

type TStrategyAllocation struct {
	Status          string           `json:"status"`
	CurrentTVL      *bigNumber.Float `json:"currentTVL"` // value in USDC
	AvailableTVL    *bigNumber.Float `json:"availableTVL"`
	CurrentAmount   *bigNumber.Float `json:"currentAmount"` // value in WANT
	AvailableAmount *bigNumber.Float `json:"availableAmount"`
}

type TStrategyGroupFromRisk struct {
	Label      string               `json:"label"`
	ChainID    uint64               `json:"chainID"`
	Strategies []common.Address     `json:"strategies"`
	Allocation *TStrategyAllocation `json:"allocation"`

	TStrategyFromRiskRiskDetails
}

/**********************************************************************************************
** TStrategyFromRiskRiskDetails is a struct that holds various risk-related scores for a
** particular strategy. These scores are used to evaluate the risk associated with the strategy.
** The scores include AuditScore, CodeReviewScore, ComplexityScore, ProtocolSafetyScore,
** TeamKnowledgeScore, TestingScore, TVLImpact, LongevityImpact, and StackingRewardTVLScore.
** Some scores are retrieved from a JSON file while others are computed based on specific
** parameters such as Total Value Locked (TVL), strategy deployment time, and Staking TVL.
**********************************************************************************************/
type TStrategyFromRiskRiskDetails struct {
	AuditScore             int `json:"auditScore"`                       // AuditScore is retrieved from JSON file
	CodeReviewScore        int `json:"codeReviewScore"`                  // CodeReviewScore is retrieved from JSON file
	ComplexityScore        int `json:"complexityScore"`                  // ComplexityScore is retrieved from JSON file
	ProtocolSafetyScore    int `json:"protocolSafetyScore"`              // ProtocolSafetyScore is retrieved from JSON file
	TeamKnowledgeScore     int `json:"teamKnowledgeScore"`               // TeamKnowledgeScore is retrieved from JSON file
	TestingScore           int `json:"testingScore"`                     // TestingScore is retrieved from JSON file
	TVLImpact              int `json:"TVLImpact"`                        // TVLImpact is computed from TVL. It's strategy specific.
	LongevityImpact        int `json:"longevityImpact"`                  // LongevityImpact is computed from strategy deployment time. It's strategy specific.
	StackingRewardTVLScore int `json:"stackingRewardTVLScore,omitempty"` // StackingRewardTVLScore is computed from Staking TVL. It's strategy specific.
}

type TStrategyFromRisk struct {
	Address     common.Address               `json:"address"`
	ChainID     uint64                       `json:"chainID"`
	RiskGroup   string                       `json:"riskGroup"`
	RiskScore   float64                      `json:"riskScore"`
	RiskDetails TStrategyFromRiskRiskDetails `json:"riskDetails"`
	Allocation  *TStrategyAllocation         `json:"allocation"`
}
