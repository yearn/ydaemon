package vaults

import (
	"strings"

	"github.com/yearn/ydaemon/common/env"
)

type Controller struct{}

func selectStrategiesCondition(s string) string {
	if s == `` {
		return `debtRatio`
	}
	for _, c := range env.STRATEGIES_CONDITIONS {
		if strings.EqualFold(c, s) {
			return c
		}
	}
	return `debtRatio`
}
