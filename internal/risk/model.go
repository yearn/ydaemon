package risk

import (
	"strconv"
	"sync"

	"github.com/yearn/ydaemon/common/env"
	"github.com/yearn/ydaemon/internal/models"
)

/**********************************************************************************************
** Set of functions to store and retrieve the strategies from the cache and/or database and
** being able to access them from the rest of the application.
** The _strategyMap variable is not exported and is only used internally by the functions below.
**********************************************************************************************/
var _strategyRiskGroupMap = make(map[uint64]*sync.Map)

func init() {
	for chainID := range env.CHAINS {
		if _, ok := _strategyRiskGroupMap[chainID]; !ok {
			_strategyRiskGroupMap[chainID] = &sync.Map{}
		}
	}
}

/**********************************************************************************************
** listStrategiesRiskGroups will, for a given chainID, return the list of all the strategies
** groups stored in the _strategyRiskGroupMap.
** The Inactive group will always be fetched first.
**********************************************************************************************/
func listStrategiesRiskGroups(chainID uint64) []*models.TStrategyGroupFromRisk {
	var groups []*models.TStrategyGroupFromRisk

	//Load the groups from the map
	syncMap := _strategyRiskGroupMap[chainID]
	if syncMap == nil {
		syncMap = &sync.Map{}
		_strategyRiskGroupMap[chainID] = syncMap
	}
	syncMap.Range(func(_, value interface{}) bool {
		groups = append(groups, value.(*models.TStrategyGroupFromRisk))
		return true
	})

	// Do what we want to do with the groups
	for _, group := range groups {
		if group.Label == `Inactive` {
			groups = append([]*models.TStrategyGroupFromRisk{group}, groups...)
		} else {
			groups = append(groups, group)
		}
	}
	return groups
}

/**********************************************************************************************
** setStrategyRiskGroup will put a models.TStrategyRiskAllocation in the _strategyRiskGroupMap
** variable.
**********************************************************************************************/
func setStrategyRiskGroup(chainID uint64, strategyRiskAllocation *models.TStrategyGroupFromRisk) {
	syncMap := _strategyRiskGroupMap[chainID]
	if syncMap == nil {
		syncMap = &sync.Map{}
		_strategyRiskGroupMap[chainID] = syncMap
	}
	key := strategyRiskAllocation.Label + strconv.FormatUint(chainID, 10)
	syncMap.Store(key, strategyRiskAllocation)
}
