package vaults

import (
	"strings"

	"github.com/yearn/ydaemon/internal/utils/env"
)

type Controller struct{}

func selectStrategiesCondition(s string) string {
	if s == `` {
		return `debtLimit`
	}
	for _, c := range env.STRATEGIES_CONDITIONS {
		if strings.EqualFold(c, s) {
			return c
		}
	}
	return `debtLimit`
}
