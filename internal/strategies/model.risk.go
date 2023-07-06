package strategies

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
	for _, chainID := range env.SUPPORTED_CHAIN_IDS {
		if _, ok := _strategyRiskGroupMap[chainID]; !ok {
			_strategyRiskGroupMap[chainID] = &sync.Map{}
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
** setRiskGroupInMap will put a models.TStrategyGroupFromRisk in the _strategyRiskGroupMap variable.
**********************************************************************************************/
func setRiskGroupInMap(chainID uint64, riskGroup *models.TStrategyGroupFromRisk) {
	syncMap := _strategyRiskGroupMap[chainID]
	if syncMap == nil {
		syncMap = &sync.Map{}
		_strategyRiskGroupMap[chainID] = syncMap
	}
	key := riskGroup.Label + strconv.FormatUint(chainID, 10)
	syncMap.Store(key, riskGroup)
}
