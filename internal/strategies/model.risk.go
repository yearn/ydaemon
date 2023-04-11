package strategies

import (
	"github.com/yearn/ydaemon/common/env"
	"github.com/yearn/ydaemon/internal/models"
)

/**********************************************************************************************
** Set of functions to store and retrieve the strategies from the cache and/or database and
** being able to access them from the rest of the application.
** The _strategyMap variable is not exported and is only used internally by the functions below.
**********************************************************************************************/
var _strategyRiskGroupMap = make(map[uint64]map[string]*models.TStrategyGroupFromRisk)

func init() {
	for _, chainID := range env.SUPPORTED_CHAIN_IDS {
		if _, ok := _strategyRiskGroupMap[chainID]; !ok {
			_strategyRiskGroupMap[chainID] = make(map[string]*models.TStrategyGroupFromRisk)
		}
	}
}

/**********************************************************************************************
** ListStrategiesRiskGroups will, for a given chainID, return the list of all the strategies
** groups stored in the _strategyRiskGroupMap.
** The Inactive group will always be fetched first.
**********************************************************************************************/
func ListStrategiesRiskGroups(chainID uint64) []*models.TStrategyGroupFromRisk {
	var groups []*models.TStrategyGroupFromRisk
	for _, group := range _strategyRiskGroupMap[chainID] {
		if group.Label == `Inactive` {
			groups = append([]*models.TStrategyGroupFromRisk{group}, groups...)
		} else {
			groups = append(groups, group)
		}
	}
	return groups
}

/**********************************************************************************************
** setRiskGroupInMap will put a models.TStrategyGroupFromRisk in the _strategyRiskGroupMap variable.
**********************************************************************************************/
func setRiskGroupInMap(chainID uint64, riskGroup *models.TStrategyGroupFromRisk) {
	_strategyRiskGroupMap[chainID][riskGroup.Label] = riskGroup
}
