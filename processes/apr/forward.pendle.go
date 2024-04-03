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
func calculatePendleStrategyAPR(
	vault models.TVault,
	strategy models.TStrategy,
) TStrategyAPR {
	if _, ok := storage.GetCachedPendleMarkets(vault.ChainID); !ok {
		storage.RefreshPendleMarkets(vault.ChainID)
	}

	debtRatio := helpers.ToNormalizedAmount(strategy.LastDebtRatio, 4)
	vaultPerformanceFee := helpers.ToNormalizedAmount(bigNumber.NewInt(int64(vault.PerformanceFee)), 4)
	vaultManagementFee := helpers.ToNormalizedAmount(bigNumber.NewInt(int64(vault.ManagementFee)), 4)
	oneMinusPerfFee := bigNumber.NewFloat(0).Sub(bigNumber.NewFloat(1), vaultPerformanceFee)

	pendleMarkets, ok := storage.GetCachedPendleMarkets(vault.ChainID)
	if !ok {
		return TStrategyAPR{
			Type:      `pendle`,
			DebtRatio: debtRatio,
			NetAPR:    bigNumber.NewFloat(0),
			Composite: TCompositeData{},
		}
	}
	data, ok := pendleMarkets[vault.AssetAddress.Hex()]
	if !ok {
		return TStrategyAPR{
			Type:      `pendle`,
			DebtRatio: debtRatio,
			NetAPR:    bigNumber.NewFloat(0),
			Composite: TCompositeData{},
		}
	}

	grossAPR := bigNumber.NewFloat(data.MaxBoostedAPY)             // Using maxBoostedAPY as Yearn is maxBoosted
	netAPR := bigNumber.NewFloat(0).Mul(grossAPR, oneMinusPerfFee) // grossAPR * (1 - perfFee)
	if netAPR.Gt(vaultManagementFee) {                             // Management fee can never induce a negative APR
		netAPR = bigNumber.NewFloat(0).Sub(netAPR, vaultManagementFee) // (grossAPR * (1 - perfFee)) - managementFee
	} else {
		netAPR = bigNumber.NewFloat(0)
	}

	return TStrategyAPR{
		Type:      `pendle`,
		DebtRatio: debtRatio,
		NetAPR:    netAPR,
		Composite: TCompositeData{},
	}
}

/**************************************************************************************************
** For a given gamma vault, we will calculate the APR based on the cached data we have.
**************************************************************************************************/
func computePendleForwardAPR(
	vault models.TVault,
	allStrategiesForVault map[string]models.TStrategy,
) TForwardAPR {
	TypeOf := ``
	NetAPR := bigNumber.NewFloat(0)
	Boost := bigNumber.NewFloat(0)
	PoolAPY := bigNumber.NewFloat(0)
	BoostedAPR := bigNumber.NewFloat(0)
	BaseAPR := bigNumber.NewFloat(0)
	CvxAPR := bigNumber.NewFloat(0)
	RewardsAPR := bigNumber.NewFloat(0)
	KeepCRV := bigNumber.NewFloat(0)
	KeepVelo := bigNumber.NewFloat(0)

	if len(allStrategiesForVault) == 0 {
		vaultAsStrategy := models.TStrategy{
			LastDebtRatio: bigNumber.NewUint64(10000),
		}
		strategyAPR := calculatePendleStrategyAPR(vault, vaultAsStrategy)
		TypeOf = strategyAPR.Type
		NetAPR = strategyAPR.NetAPR
		Boost = strategyAPR.Composite.Boost
		PoolAPY = strategyAPR.Composite.PoolAPY
		BoostedAPR = strategyAPR.Composite.BoostedAPR
		BaseAPR = strategyAPR.Composite.BaseAPR
		CvxAPR = strategyAPR.Composite.CvxAPR
		RewardsAPR = strategyAPR.Composite.RewardsAPR
		KeepCRV = strategyAPR.Composite.KeepCRV
		KeepVelo = strategyAPR.Composite.KeepVelo
	} else {
		for _, strategy := range allStrategiesForVault {
			if strategy.LastDebtRatio == nil || strategy.LastDebtRatio.IsZero() {
				if os.Getenv("ENVIRONMENT") == "dev" {
					logs.Info("Skipping strategy " + strategy.Address.Hex() + " for vault " + vault.Address.Hex() + " because debt ratio is zero")
				}
				continue
			}

			strategyAPR := calculatePendleStrategyAPR(vault, strategy)
			TypeOf += strings.TrimSpace(` ` + strategyAPR.Type)
			NetAPR = bigNumber.NewFloat(0).Add(NetAPR, strategyAPR.NetAPR)
			Boost = bigNumber.NewFloat(0).Add(Boost, strategyAPR.Composite.Boost)
			PoolAPY = bigNumber.NewFloat(0).Add(PoolAPY, strategyAPR.Composite.PoolAPY)
			BoostedAPR = bigNumber.NewFloat(0).Add(BoostedAPR, strategyAPR.Composite.BoostedAPR)
			BaseAPR = bigNumber.NewFloat(0).Add(BaseAPR, strategyAPR.Composite.BaseAPR)
			CvxAPR = bigNumber.NewFloat(0).Add(CvxAPR, strategyAPR.Composite.CvxAPR)
			RewardsAPR = bigNumber.NewFloat(0).Add(RewardsAPR, strategyAPR.Composite.RewardsAPR)
			KeepCRV = bigNumber.NewFloat(0).Add(KeepCRV, strategyAPR.Composite.KeepCRV)
			KeepVelo = bigNumber.NewFloat(0).Add(KeepVelo, strategyAPR.Composite.KeepVelo)
		}
	}

	return TForwardAPR{
		Type:   strings.TrimSpace(TypeOf),
		NetAPR: NetAPR,
		Composite: TCompositeData{
			Boost:      Boost,
			PoolAPY:    PoolAPY,
			BoostedAPR: BoostedAPR,
			BaseAPR:    BaseAPR,
			CvxAPR:     CvxAPR,
			RewardsAPR: RewardsAPR,
			KeepCRV:    KeepCRV,
			KeepVelo:   KeepVelo,
		},
	}
}
