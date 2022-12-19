package vaults

import (
	"strings"
)

type Controller struct{}

// STRATEGIES_CONDITIONS contains the possible conditions to determine which strategies should
// be used in the return value.
// If the strategy is `absolute`, an active strategy will be isActive, inQueue and with a debt > 0
// If the strategy is `inQueue`, we will check if the vault has the strategy in it's withdrawal queue
// If the strategy is `debtLimit`, we will check if the vault has allocated a debtLimit to the strategy
// If the strategy is `debtRatio`, we will check if the vault has allocated a debtRatio to the strategy
var STRATEGIES_CONDITIONS = []string{`inQueue`, `debtLimit`, `debtRatio`, `absolute`, `all`}

func selectStrategiesCondition(s string) string {
	if s == `` {
		return `debtRatio`
	}
	for _, c := range STRATEGIES_CONDITIONS {
		if strings.EqualFold(c, s) {
			return c
		}
	}
	return `debtRatio`
}

// MIGRABLE_CONDITIONS contains the possible conditions to determine if the migrable vaults should be
// included in the return value.
// If the strategy is `none`, they will not be included
// If the strategy is `all`, they will all be included
// If the strategy is `nodust`, they will be included if they have a TVL > 100$
var MIGRABLE_CONDITIONS = []string{`none`, `all`, `nodust`}

func selectMigrableCondition(s string) string {
	if s == `` {
		return `none`
	}
	for _, c := range MIGRABLE_CONDITIONS {
		if strings.EqualFold(c, s) {
			return c
		}
	}
	return `none`
}
