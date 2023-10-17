package vaults

import (
	"strconv"

	"github.com/yearn/ydaemon/common/bigNumber"
	"github.com/yearn/ydaemon/common/helpers"
	"github.com/yearn/ydaemon/internal/models"
	"github.com/yearn/ydaemon/internal/storage"
)

// TExternalStrategyDetails contains the details about a strategy.
type TExternalStrategyDetails struct {
	Keeper               string         `json:"keeper"`
	Strategist           string         `json:"strategist"`
	Rewards              string         `json:"rewards"`
	HealthCheck          string         `json:"healthCheck"`
	TotalDebt            *bigNumber.Int `json:"totalDebt"`
	TotalLoss            *bigNumber.Int `json:"totalLoss"`
	TotalGain            *bigNumber.Int `json:"totalGain"`
	RateLimit            *bigNumber.Int `json:"rateLimit,omitempty"`         // Only < 0.3.2
	MinDebtPerHarvest    *bigNumber.Int `json:"minDebtPerHarvest,omitempty"` // Only >= 0.3.2
	MaxDebtPerHarvest    *bigNumber.Int `json:"maxDebtPerHarvest,omitempty"` // Only >= 0.3.2
	EstimatedTotalAssets *bigNumber.Int `json:"estimatedTotalAssets"`
	CreditAvailable      *bigNumber.Int `json:"creditAvailable"`
	DebtOutstanding      *bigNumber.Int `json:"debtOutstanding"`
	ExpectedReturn       *bigNumber.Int `json:"expectedReturn"`
	DelegatedAssets      *bigNumber.Int `json:"delegatedAssets"`
	DelegatedValue       string         `json:"delegatedValue"`
	Version              string         `json:"version"`
	Protocols            []string       `json:"protocols"`
	APR                  float64        `json:"apr"`
	PerformanceFee       uint64         `json:"performanceFee"`
	LastReport           uint64         `json:"lastReport"`
	Activation           uint64         `json:"activation"`
	KeepCRV              uint64         `json:"keepCRV"`
	KeepCVX              uint64         `json:"keepCVX"`
	DebtRatio            uint64         `json:"debtRatio,omitempty"` // Only > 0.2.2
	DebtLimit            uint64         `json:"debtLimit"`
	DoHealthCheck        bool           `json:"doHealthCheck"`
	InQueue              bool           `json:"inQueue"`
	EmergencyExit        bool           `json:"emergencyExit"`
	IsActive             bool           `json:"isActive"`
}

// TStrategy contains all the information useful about the strategies currently active in this vault.
type TStrategy struct {
	Address     string                      `json:"address"`
	Name        string                      `json:"name"`
	DisplayName string                      `json:"displayName"`
	Description string                      `json:"description"`
	Details     *TExternalStrategyDetails   `json:"details,omitempty"`
	Risk        *TExternalStrategyRiskScore `json:"risk,omitempty"`
}

func buildDelegated(delegatedBalanceToken *bigNumber.Int, decimals int, humanizedPrice *bigNumber.Float) float64 {
	_, delegatedTVL := helpers.FormatAmount(delegatedBalanceToken.String(), decimals)
	fHumanizedTVLPrice, _ := bigNumber.NewFloat().Mul(delegatedTVL, humanizedPrice).Float64()
	return fHumanizedTVLPrice
}

func NewStrategy() TStrategy {
	return TStrategy{}
}
func (v TStrategy) AssignTStrategy(strategy models.TStrategy) TStrategy {
	delegatedValue := `0`
	if vault, ok := storage.GetVault(strategy.ChainID, strategy.VaultAddress); ok {
		if vaultToken, ok := storage.GetERC20(vault.ChainID, vault.Address); ok {
			if tokenPrice, ok := storage.GetPrice(strategy.ChainID, vault.AssetAddress); ok {
				delegatedValue = strconv.FormatFloat(
					buildDelegated(
						strategy.LastDelegatedAssets,
						int(vaultToken.Decimals),
						tokenPrice.HumanizedPrice,
					), 'f', -1, 64,
				)
			}
		}
	}

	v.Address = strategy.Address.Hex()
	v.Name = strategy.Name
	v.DisplayName = strategy.DisplayName
	v.Description = strategy.Description
	v.Details = &TExternalStrategyDetails{
		TotalDebt:            strategy.LastTotalDebt,
		TotalLoss:            strategy.LastTotalLoss,
		TotalGain:            strategy.LastTotalGain,
		RateLimit:            strategy.LastRateLimit,
		MinDebtPerHarvest:    strategy.LastMinDebtPerHarvest,
		MaxDebtPerHarvest:    strategy.LastMaxDebtPerHarvest,
		EstimatedTotalAssets: strategy.LastEstimatedTotalAssets,
		CreditAvailable:      strategy.LastCreditAvailable,
		DebtOutstanding:      strategy.LastDebtOutstanding,
		ExpectedReturn:       strategy.LastExpectedReturn,
		DelegatedAssets:      strategy.LastDelegatedAssets,
		DelegatedValue:       delegatedValue,
		Protocols:            strategy.Protocols,
		PerformanceFee:       strategy.LastPerformanceFee.Uint64(),
		LastReport:           strategy.LastReport.Uint64(),
		Activation:           strategy.TimeActivated.Uint64(),
		KeepCRV:              (strategy.KeepCRV.Uint64() + strategy.KeepCRVPercent.Uint64()),
		KeepCVX:              strategy.KeepCVX.Uint64(),
		DebtRatio:            strategy.LastDebtRatio.Uint64(),
		DebtLimit:            strategy.LastDebtLimit.Uint64(),
		DoHealthCheck:        strategy.DoHealthCheck,
		InQueue:              strategy.IsInQueue,
		EmergencyExit:        strategy.EmergencyExit,
		IsActive:             strategy.IsActive,
	}
	return v
}
func (v TStrategy) ShouldBeIncluded(condition string) bool {
	if condition == `all` {
		return true
	} else if condition == `absolute` &&
		v.Details.InQueue &&
		v.Details.IsActive &&
		v.Details.TotalDebt.Gt(bigNumber.Zero) {
		return true
	} else if condition == `inQueue` && v.Details.InQueue {
		return true
	} else if condition == `debtLimit` && v.Details.DebtLimit > 0 {
		return true
	} else if condition == `debtRatio` && (v.Details.DebtRatio > 0 || v.Details.DebtLimit > 0) {
		return true
	}
	return false
}
