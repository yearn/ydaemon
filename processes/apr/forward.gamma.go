package apr

import (
	"os"
	"strings"

	"github.com/yearn/ydaemon/common/bigNumber"
	"github.com/yearn/ydaemon/common/helpers"
	"github.com/yearn/ydaemon/common/logs"
	"github.com/yearn/ydaemon/internal/models"
	"github.com/yearn/ydaemon/internal/storage"
)

/**************************************************************************************************
** Check if the vault is a gamma vault by checking if we have the relevant information cached.
**************************************************************************************************/
func isGammaVault(chainID uint64, vault models.TVault) bool {
	if _, ok := storage.GetCachedGammaMerkl(chainID); !ok {
		storage.RefreshGammaCalls(chainID)
	}
	if _, ok := storage.GetCachedGammaMerkl(chainID); !ok {
		return false
	}

	gammaAllData, _ := storage.GetCachedGammaAllData(chainID)
	_, ok := gammaAllData[vault.AssetAddress.Hex()]
	return ok
}

/**************************************************************************************************
** For a given gamma strategy, we will calculate the APR based on the cached data we have.
**************************************************************************************************/
func calculateGammaStrategyAPY(
	vault models.TVault,
	strategy models.TStrategy,
) (*bigNumber.Float, *bigNumber.Float) {
	if _, ok := storage.GetCachedGammaMerkl(vault.ChainID); !ok {
		storage.RefreshGammaCalls(vault.ChainID)
	}

	debtRatio := helpers.ToNormalizedAmount(strategy.LastDebtRatio, 4)
	vaultPerformanceFee := helpers.ToNormalizedAmount(bigNumber.NewInt(int64(vault.PerformanceFee)), 4)
	vaultManagementFee := helpers.ToNormalizedAmount(bigNumber.NewInt(int64(vault.ManagementFee)), 4)
	oneMinusPerfFee := bigNumber.NewFloat(0).Sub(bigNumber.NewFloat(1), vaultPerformanceFee)

	if _, ok := storage.GetCachedGammaMerkl(vault.ChainID); !ok {
		return bigNumber.NewFloat(0), bigNumber.NewFloat(0)
	}

	gammaAllData, _ := storage.GetCachedGammaAllData(vault.ChainID)
	data, ok := gammaAllData[vault.AssetAddress.Hex()]
	if !ok {
		return bigNumber.NewFloat(0), bigNumber.NewFloat(0)
	}
	grossAPR := bigNumber.NewFloat(data.Returns.Monthly.APR)
	netAPR := bigNumber.NewFloat(0).Mul(grossAPR, oneMinusPerfFee) // grossAPR * (1 - perfFee)
	if netAPR.Gt(vaultManagementFee) {
		netAPR = bigNumber.NewFloat(0).Sub(netAPR, vaultManagementFee) // (grossAPR * (1 - perfFee)) - managementFee
	} else {
		netAPR = bigNumber.NewFloat(0)
	}

	/**********************************************************************************************
	** Calculate the strategy Net APY:
	** Take the net APR and compound it every 15 days
	**********************************************************************************************/
	const daysInYear = 365
	const compoundingPeriod = 15
	const compoundingPeriodsPerYear = daysInYear / compoundingPeriod

	netAPY := bigNumber.NewFloat(0).Div(netAPR, bigNumber.NewFloat(compoundingPeriodsPerYear)) // netAPR / (365 / 15)
	netAPY = bigNumber.NewFloat(0).Add(netAPY, bigNumber.NewFloat(1))                          // 1 + (netAPR / (365 / 15))
	netAPY = bigNumber.NewFloat(0).Pow(netAPY, compoundingPeriodsPerYear)                      // (1 + (netAPR / (365 / 15))) ^ (365 / 15)
	netAPY = bigNumber.NewFloat(0).Sub(netAPY, bigNumber.NewFloat(1))                          // ((1 + (netAPR / (365 / 15))) ^ (365 / 15)) - 1

	return bigNumber.NewFloat(0).Mul(netAPR, debtRatio), bigNumber.NewFloat(0).Mul(netAPY, debtRatio)
}

/**************************************************************************************************
** For a given gamma vault, we will calculate the APR based on the cached data we have.
**************************************************************************************************/
func computeGammaForwardAPY(
	vault models.TVault,
	allStrategiesForVault map[string]models.TStrategy,
) TForwardAPY {
	TypeOf := ``
	netAPY := bigNumber.NewFloat(0)
	boost := bigNumber.NewFloat(0)
	poolAPY := bigNumber.NewFloat(0)
	boostedAPR := bigNumber.NewFloat(0)
	baseAPR := bigNumber.NewFloat(0)
	cvxAPR := bigNumber.NewFloat(0)
	rewardsAPY := bigNumber.NewFloat(0)
	keepCRV := bigNumber.NewFloat(0)
	keepVelo := bigNumber.NewFloat(0)

	if len(allStrategiesForVault) == 0 {
		vaultAsStrategy := models.TStrategy{
			LastDebtRatio: bigNumber.NewUint64(10000),
		}
		_, strategyAPY := calculateGammaStrategyAPY(vault, vaultAsStrategy)
		TypeOf = `gamma`
		netAPY = strategyAPY
	} else {
		for _, strategy := range allStrategiesForVault {
			if strategy.LastDebtRatio == nil || strategy.LastDebtRatio.IsZero() {
				if os.Getenv("ENVIRONMENT") == "dev" {
					logs.Info("Skipping strategy " + strategy.Address.Hex() + " for vault " + vault.Address.Hex() + " because debt ratio is zero")
				}
				continue
			}

			_, strategyAPY := calculateGammaStrategyAPY(vault, strategy)
			TypeOf += strings.TrimSpace(` ` + `gamma`)
			netAPY = bigNumber.NewFloat(0).Add(netAPY, strategyAPY)
		}
	}

	return TForwardAPY{
		Type:   strings.TrimSpace(TypeOf),
		NetAPY: netAPY,
		Composite: TCompositeData{
			Boost:      boost,
			PoolAPY:    poolAPY,
			BoostedAPR: boostedAPR,
			BaseAPR:    baseAPR,
			CvxAPR:     cvxAPR,
			RewardsAPY: rewardsAPY,
			KeepCRV:    keepCRV,
			KeepVelo:   keepVelo,
		},
	}
}
