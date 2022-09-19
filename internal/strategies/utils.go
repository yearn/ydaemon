package strategies

import "github.com/yearn/ydaemon/internal/utils/models"

type Controller struct{}

func GetStrategyGroupsFromStrategy(strat models.TStrategyList, groups []*TStrategyGroupFromRisk) []*TStrategyGroupFromRisk {
	var stratGroups []*TStrategyGroupFromRisk
	for _, group := range groups {
		if excludeNameLike(strat, *group) {
			continue
		}
		if includeAddress(strat, *group) || includeNameLike(strat, *group) {
			stratGroups = append(stratGroups, group)
			break
		}
	}
	return stratGroups
}
