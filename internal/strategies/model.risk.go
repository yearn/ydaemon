package strategies

import (
	"github.com/yearn/ydaemon/common/bigNumber"
	"github.com/yearn/ydaemon/common/types/common"
)

// TStrategyGroupCritieria is a helper type for parsing the risk framework json
type TStrategyGroupCritieria struct {
	Strategies []string `json:"strategies"`
	Exclude    []string `json:"exclude"`
	NameLike   []string `json:"nameLike"`
}

type TStrategyAllocation struct {
	Status          string           `json:"status"`
	CurrentTVL      *bigNumber.Float `json:"currentTVL"` // value in USDC
	AvailableTVL    *bigNumber.Float `json:"availableTVL"`
	CurrentAmount   *bigNumber.Float `json:"currentAmount"` // value in WANT
	AvailableAmount *bigNumber.Float `json:"availableAmount"`
}

type TStrategyGroupFromRisk struct {
	Label               string                  `json:"label"`
	AuditScore          int                     `json:"auditScore"`
	CodeReviewScore     int                     `json:"codeReviewScore"`
	ComplexityScore     int                     `json:"complexityScore"`
	ProtocolSafetyScore int                     `json:"protocolSafetyScore"`
	TeamKnowledgeScore  int                     `json:"teamKnowledgeScore"`
	TestingScore        int                     `json:"testingScore"`
	ChainID             uint64                  `json:"chainID"`
	Criteria            TStrategyGroupCritieria `json:"criteria"`
	Allocation          *TStrategyAllocation    `json:"allocation"`
}

type TStrategyFromRiskRiskDetails struct {
	TVLImpact           int `json:"TVLImpact"`
	AuditScore          int `json:"auditScore"`
	CodeReviewScore     int `json:"codeReviewScore"`
	ComplexityScore     int `json:"complexityScore"`
	LongevityImpact     int `json:"longevityImpact"`
	ProtocolSafetyScore int `json:"protocolSafetyScore"`
	TeamKnowledgeScore  int `json:"teamKnowledgeScore"`
	TestingScore        int `json:"testingScore"`
}

type TStrategyFromRisk struct {
	Address     common.Address               `json:"address"`
	ChainID     uint64                       `json:"chainID"`
	RiskGroup   string                       `json:"riskGroup"`
	RiskScore   float64                      `json:"riskScore"`
	RiskDetails TStrategyFromRiskRiskDetails `json:"riskDetails"`
	Allocation  *TStrategyAllocation         `json:"allocation"`
}

/**********************************************************************************************
** Set of functions to store and retrieve the strategies from the cache and/or database and
** being able to access them from the rest of the application.
** The _strategyMap variable is not exported and is only used internally by the functions below.
**********************************************************************************************/
var _strategyRiskGroupMap = make(map[uint64]map[string]*TStrategyGroupFromRisk)

/**********************************************************************************************
** ListStrategiesRiskGroups will, for a given chainID, return the list of all the strategies
** groups stored in the _strategyRiskGroupMap.
** The Inactive group will always be fetched first.
**********************************************************************************************/
func ListStrategiesRiskGroups(chainID uint64) []*TStrategyGroupFromRisk {
	var groups []*TStrategyGroupFromRisk
	for _, group := range _strategyRiskGroupMap[chainID] {
		if group.Label == `Inactive` {
			groups = append([]*TStrategyGroupFromRisk{group}, groups...)
		} else {
			groups = append(groups, group)
		}
	}
	return groups
}

/**********************************************************************************************
** setRiskGroupInMap will put a TStrategyGroupFromRisk in the _strategyRiskGroupMap variable.
**********************************************************************************************/
func setRiskGroupInMap(chainID uint64, riskGroup *TStrategyGroupFromRisk) {
	if _, ok := _strategyRiskGroupMap[chainID]; !ok {
		_strategyRiskGroupMap[chainID] = make(map[string]*TStrategyGroupFromRisk)
	}
	_strategyRiskGroupMap[chainID][riskGroup.Label] = riskGroup
}
