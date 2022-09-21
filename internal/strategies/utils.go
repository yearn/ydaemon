package strategies

import (
	"sync"

	"github.com/yearn/ydaemon/internal/utils/models"
)

type Controller struct{}

func GetStrategyGroupsFromStrategy(strat models.TStrategyList, groups []*TStrategyGroupFromRisk) []*TStrategyGroupFromRisk {
	var stratGroups []*TStrategyGroupFromRisk
	var mutex = &sync.Mutex{}
	mutex.Lock()
	for _, group := range groups {
		if excludeNameLike(strat, *group) {
			continue
		}
		if includeAddress(strat, *group) || includeNameLike(strat, *group) {
			stratGroups = append(stratGroups, group)
		}
	}
	mutex.Unlock()
	return stratGroups
}
