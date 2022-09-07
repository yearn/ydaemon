package vaultsController

import (
	"strings"

	"github.com/yearn/ydaemon/internal/utils"
)

type Controller struct{}

func selectStrategiesCondition(s string) string {
	if s == `` {
		return `debtLimit`
	}
	for _, c := range utils.STRATEGIES_CONDITIONS {
		if strings.EqualFold(c, s) {
			return c
		}
	}
	return `debtLimit`
}
