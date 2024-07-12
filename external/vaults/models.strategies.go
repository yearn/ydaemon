package vaults

import (
	"github.com/yearn/ydaemon/common/bigNumber"
	"github.com/yearn/ydaemon/internal/models"
)

// TExternalStrategyDetails contains the details about a strategy.
type TExternalStrategyDetails struct {
	TotalDebt      *bigNumber.Int `json:"totalDebt"`
	TotalLoss      *bigNumber.Int `json:"totalLoss"`
	TotalGain      *bigNumber.Int `json:"totalGain"`
	PerformanceFee uint64         `json:"performanceFee"`
	LastReport     uint64         `json:"lastReport"`
	DebtRatio      uint64         `json:"debtRatio,omitempty"` // Only > 0.2.2
	InQueue        bool           `json:"-"`
}

// TStrategy contains all the information useful about the strategies currently active in this vault.
type TStrategy struct {
	Address     string                    `json:"address"`
	Name        string                    `json:"name"`
	Description string                    `json:"description,omitempty"`
	Details     *TExternalStrategyDetails `json:"details,omitempty"`
}

func NewStrategy() TStrategy {
	return TStrategy{}
}

func (v TStrategy) AssignTStrategy(strategy models.TStrategy) TStrategy {
	v.Address = strategy.Address.Hex()
	v.Name = strategy.DisplayName
	if v.Name == "" {
		v.Name = strategy.Name
	}
	v.Description = strategy.Description
	v.Details = &TExternalStrategyDetails{
		TotalDebt:      strategy.LastTotalDebt,
		TotalLoss:      strategy.LastTotalLoss,
		TotalGain:      strategy.LastTotalGain,
		PerformanceFee: strategy.LastPerformanceFee.Uint64(),
		LastReport:     strategy.LastReport.Uint64(),
		DebtRatio:      strategy.LastDebtRatio.Uint64(),
		InQueue:        strategy.IsInQueue,
	}
	return v
}

func (v TStrategy) ShouldBeIncluded(condition string) bool {
	if condition == `all` {
		return true
	} else if condition == `absolute` && v.Details.TotalDebt.Gt(bigNumber.Zero) {
		return true
	} else if condition == `inQueue` && v.Details.InQueue {
		return true
	} else if condition == `debtRatio` && v.Details.DebtRatio > 0 {
		return true
	}
	return false
}
