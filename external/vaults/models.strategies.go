package vaults

import (
	"strconv"

	"github.com/yearn/ydaemon/common/bigNumber"
	"github.com/yearn/ydaemon/common/helpers"
	"github.com/yearn/ydaemon/common/types/common"
	"github.com/yearn/ydaemon/internal/prices"
	"github.com/yearn/ydaemon/internal/strategies"
	"github.com/yearn/ydaemon/internal/vaults"
)

// TStrategyDetails contains the details about a strategy.
type TStrategyDetails struct {
	Keeper                  common.Address `json:"keeper"`
	Strategist              common.Address `json:"strategist"`
	Rewards                 common.Address `json:"rewards"`
	HealthCheck             common.Address `json:"healthCheck"`
	TotalDebt               *bigNumber.Int `json:"totalDebt"`
	TotalLoss               *bigNumber.Int `json:"totalLoss"`
	TotalGain               *bigNumber.Int `json:"totalGain"`
	RateLimit               *bigNumber.Int `json:"rateLimit,omitempty"`         // Only < 0.3.2
	MinDebtPerHarvest       *bigNumber.Int `json:"minDebtPerHarvest,omitempty"` // Only >= 0.3.2
	MaxDebtPerHarvest       *bigNumber.Int `json:"maxDebtPerHarvest,omitempty"` // Only >= 0.3.2
	EstimatedTotalAssets    *bigNumber.Int `json:"estimatedTotalAssets"`
	CreditAvailable         *bigNumber.Int `json:"creditAvailable"`
	DebtOutstanding         *bigNumber.Int `json:"debtOutstanding"`
	ExpectedReturn          *bigNumber.Int `json:"expectedReturn"`
	DelegatedAssets         *bigNumber.Int `json:"delegatedAssets"`
	DelegatedValue          string         `json:"delegatedValue"`
	Version                 string         `json:"version"`
	Protocols               []string       `json:"protocols"`
	APR                     float64        `json:"apr"`
	PerformanceFee          uint64         `json:"performanceFee"`
	LastReport              uint64         `json:"lastReport"`
	Activation              uint64         `json:"activation"`
	KeepCRV                 uint64         `json:"keepCRV"`
	DebtRatio               uint64         `json:"debtRatio,omitempty"` // Only > 0.2.2
	DebtLimit               uint64         `json:"debtLimit"`
	WithdrawalQueuePosition int64          `json:"withdrawalQueuePosition"`
	DoHealthCheck           bool           `json:"doHealthCheck"`
	InQueue                 bool           `json:"inQueue"`
	EmergencyExit           bool           `json:"emergencyExit"`
	IsActive                bool           `json:"isActive"`
}

// TStrategy contains all the information useful about the strategies currently active in this vault.
type TStrategy struct {
	Address     common.Address              `json:"address"`
	Name        string                      `json:"name"`
	Description string                      `json:"description"`
	Details     *TStrategyDetails           `json:"details,omitempty"`
	Risk        *TExternalStrategyRiskScore `json:"risk,omitempty"`
}

func buildDelegated(delegatedBalanceToken *bigNumber.Int, decimals int, humanizedPrice *bigNumber.Float) float64 {
	_, delegatedTVL := helpers.FormatAmount(delegatedBalanceToken.String(), decimals)
	fHumanizedTVLPrice, _ := bigNumber.NewFloat().Mul(delegatedTVL, humanizedPrice).Float64()
	return fHumanizedTVLPrice
}

func NewStrategy() *TStrategy {
	return &TStrategy{}
}
func (v *TStrategy) AssignTStrategy(strategy *strategies.TStrategy) *TStrategy {
	delegatedValue := `0`
	if vault, ok := vaults.FindVault(strategy.ChainID, common.FromAddress(strategy.VaultAddress)); ok {
		if tokenPrice, ok := prices.FindPrice(strategy.ChainID, common.FromAddress(vault.Token.Address)); ok {
			// tokenPrice
			_, humanizedTokenPrice := helpers.FormatAmount(tokenPrice.String(), 6)
			delegatedValue = strconv.FormatFloat(
				buildDelegated(
					strategy.DelegatedAssets,
					int(vault.Decimals),
					humanizedTokenPrice,
				), 'f', -1, 64,
			)
		}
	}

	v.Address = common.FromAddress(strategy.Address)
	v.Name = strategy.Name
	v.Description = strategy.Description
	v.Details = &TStrategyDetails{
		Keeper:               common.FromAddress(strategy.KeeperAddress),      //keeper
		Strategist:           common.FromAddress(strategy.StrategistAddress),  //strategist
		Rewards:              common.FromAddress(strategy.RewardsAddress),     //rewards
		HealthCheck:          common.FromAddress(strategy.HealthCheckAddress), //healthCheck
		TotalDebt:            strategy.TotalDebt,
		TotalLoss:            strategy.TotalLoss,
		TotalGain:            strategy.TotalGain,
		RateLimit:            strategy.RateLimit,
		MinDebtPerHarvest:    strategy.MinDebtPerHarvest,
		MaxDebtPerHarvest:    strategy.MaxDebtPerHarvest,
		EstimatedTotalAssets: strategy.EstimatedTotalAssets,
		CreditAvailable:      strategy.CreditAvailable,
		DebtOutstanding:      strategy.DebtOutstanding,
		ExpectedReturn:       strategy.ExpectedReturn,
		DelegatedAssets:      strategy.DelegatedAssets,
		DelegatedValue:       delegatedValue,
		Protocols:            strategy.Protocols,
		Version:              strategy.APIVersion,
		// APR:                     strategy.APR, //NOT AVAILABLE
		PerformanceFee:          strategy.PerformanceFee.Uint64(),
		LastReport:              strategy.LastReport.Uint64(),
		Activation:              strategy.Activation.Uint64(),
		KeepCRV:                 strategy.KeepCRV.Uint64(),
		DebtRatio:               strategy.DebtRatio.Uint64(),
		DebtLimit:               strategy.DebtLimit.Uint64(),
		WithdrawalQueuePosition: strategy.WithdrawalQueuePosition.Int64(),
		DoHealthCheck:           strategy.DoHealthCheck,
		InQueue:                 strategy.IsInQueue,
		EmergencyExit:           strategy.EmergencyExit,
		IsActive:                strategy.IsActive,
	}
	return v
}
