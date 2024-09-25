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
func isPendleVault(chainID uint64, vault models.TVault) bool {
	if _, ok := storage.GetCachedPendleMarkets(chainID); !ok {
		storage.RetrievePendleMarkets(chainID)
	}
	pendleMarkets, ok := storage.GetCachedPendleMarkets(chainID)
	if !ok {
		return false
	}

	_, ok = pendleMarkets[vault.AssetAddress.Hex()]
	return ok
}

/**************************************************************************************************
** For a given gamma strategy, we will calculate the APR based on the cached data we have.
**************************************************************************************************/
func calculatePendleStrategyAPY(
	vault models.TVault,
	strategy models.TStrategy,
) TStrategyAPY {
	if _, ok := storage.GetCachedPendleMarkets(vault.ChainID); !ok {
		storage.RefreshPendleMarkets(vault.ChainID)
	}

	debtRatio := helpers.ToNormalizedAmount(strategy.LastDebtRatio, 4)
	vaultPerformanceFee := helpers.ToNormalizedAmount(bigNumber.NewInt(int64(vault.PerformanceFee)), 4)
	vaultManagementFee := helpers.ToNormalizedAmount(bigNumber.NewInt(int64(vault.ManagementFee)), 4)
	oneMinusPerfFee := bigNumber.NewFloat(0).Sub(bigNumber.NewFloat(1), vaultPerformanceFee)

	pendleMarkets, ok := storage.GetCachedPendleMarkets(vault.ChainID)
	if !ok {
		return TStrategyAPY{
			Type:      `pendle`,
			DebtRatio: debtRatio,
			NetAPY:    bigNumber.NewFloat(0),
			Composite: TCompositeData{},
		}
	}
	data, ok := pendleMarkets[vault.AssetAddress.Hex()]
	if !ok {
		return TStrategyAPY{
			Type:      `pendle`,
			DebtRatio: debtRatio,
			NetAPY:    bigNumber.NewFloat(0),
			Composite: TCompositeData{},
		}
	}

	grossAPY := bigNumber.NewFloat(data.AggretatedAPY)             // Using aggregatedApy
	netAPY := bigNumber.NewFloat(0).Mul(grossAPY, oneMinusPerfFee) // grossAPY * (1 - perfFee)
	if netAPY.Gt(vaultManagementFee) {                             // Management fee can never induce a negative APR
		netAPY = bigNumber.NewFloat(0).Sub(netAPY, vaultManagementFee) // (grossAPY * (1 - perfFee)) - managementFee
	} else {
		netAPY = bigNumber.NewFloat(0)
	}

	return TStrategyAPY{
		Type:      `pendle`,
		DebtRatio: debtRatio,
		NetAPY:    netAPY, //Actually APY
		Composite: TCompositeData{},
	}
}

/**************************************************************************************************
** For a given gamma vault, we will calculate the APR based on the cached data we have.
**************************************************************************************************/
func computePendleForwardAPY(
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
		strategyAPY := calculatePendleStrategyAPY(vault, vaultAsStrategy)
		TypeOf = strategyAPY.Type
		netAPY = strategyAPY.NetAPY
		boost = strategyAPY.Composite.Boost
		poolAPY = strategyAPY.Composite.PoolAPY
		boostedAPR = strategyAPY.Composite.BoostedAPR
		baseAPR = strategyAPY.Composite.BaseAPR
		cvxAPR = strategyAPY.Composite.CvxAPR
		rewardsAPY = strategyAPY.Composite.RewardsAPY
		keepCRV = strategyAPY.Composite.KeepCRV
		keepVelo = strategyAPY.Composite.KeepVelo
	} else {
		for _, strategy := range allStrategiesForVault {
			if strategy.LastDebtRatio == nil || strategy.LastDebtRatio.IsZero() {
				if os.Getenv("ENVIRONMENT") == "dev" {
					logs.Info("Skipping strategy " + strategy.Address.Hex() + " for vault " + vault.Address.Hex() + " because debt ratio is zero")
				}
				continue
			}

			strategyAPY := calculatePendleStrategyAPY(vault, strategy)
			TypeOf += strings.TrimSpace(` ` + strategyAPY.Type)
			netAPY = bigNumber.NewFloat(0).Add(netAPY, strategyAPY.NetAPY)
			boost = bigNumber.NewFloat(0).Add(boost, strategyAPY.Composite.Boost)
			poolAPY = bigNumber.NewFloat(0).Add(poolAPY, strategyAPY.Composite.PoolAPY)
			boostedAPR = bigNumber.NewFloat(0).Add(boostedAPR, strategyAPY.Composite.BoostedAPR)
			baseAPR = bigNumber.NewFloat(0).Add(baseAPR, strategyAPY.Composite.BaseAPR)
			cvxAPR = bigNumber.NewFloat(0).Add(cvxAPR, strategyAPY.Composite.CvxAPR)
			rewardsAPY = bigNumber.NewFloat(0).Add(rewardsAPY, strategyAPY.Composite.RewardsAPY)
			keepCRV = bigNumber.NewFloat(0).Add(keepCRV, strategyAPY.Composite.KeepCRV)
			keepVelo = bigNumber.NewFloat(0).Add(keepVelo, strategyAPY.Composite.KeepVelo)
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
