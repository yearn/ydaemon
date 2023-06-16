package apy

import (
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/common/bigNumber"
	"github.com/yearn/ydaemon/common/ethereum"
	"github.com/yearn/ydaemon/common/helpers"
	"github.com/yearn/ydaemon/internal/models"
	"github.com/yearn/ydaemon/internal/multicalls"
)

/**************************************************************************************************
** Check if the strategies related to that vault are a curve strategy. If the vault has more than
** 2, it can't be.
** TLDR; check if name contains curve or convex
**************************************************************************************************/
func isCurveVault(strategies []*models.TStrategy) bool {
	if len(strategies) > 2 {
		return false
	}
	isCurveOrConvexStrategy := false
	for _, strategy := range strategies {
		strategyName := strings.ToLower(strategy.Name)
		if strings.Contains(strategyName, `curve`) || strings.Contains(strategyName, `convex`) {
			isCurveOrConvexStrategy = true
			break
		}
	}
	return isCurveOrConvexStrategy
}

/**************************************************************************************************
** We need to know how is performing the pool in term of APY. We could calculate the whole what is
** the price now, vs what it was one week ago, and cale it to one year etc. but the weekly apy
** is already available in the Curve API.
** So just use it.
**************************************************************************************************/
func getPoolWeeklyAPY(subgraphItem models.CurveSubgraphData) *bigNumber.Float {
	poolAPY := bigNumber.NewFloat(0).Div(
		bigNumber.NewFloat(0).SetFloat64(subgraphItem.LatestWeeklyApy),
		bigNumber.NewFloat(100),
	)
	return poolAPY
}

/**************************************************************************************************
** We need to know how is performing the pool in term of APY. We could calculate the whole what is
** the price now, vs what it was one week ago, and cale it to one year etc. but the daily apy
** is already available in the Curve API.
** So just use it. It's mainly used for the convex type of curve APR.
**************************************************************************************************/
func getPoolDailyAPY(subgraphItem models.CurveSubgraphData) *bigNumber.Float {
	poolAPY := bigNumber.NewFloat(0).Div(
		bigNumber.NewFloat(0).SetFloat64(subgraphItem.LatestDailyApy),
		bigNumber.NewFloat(100),
	)
	return poolAPY
}

/**************************************************************************************************
** The poolPrice is returned by the curve API. It's the virtualPrice of the gauge endpoint.
** As it's resolved as an unknown type (either string or int), we need to use this hack to make
** sure we are using the right value.
** Returns the poolPrice as a bigNumber.Float in base 2 (humanized)
**************************************************************************************************/
func getPoolPrice(gauge models.CurveGauge) *bigNumber.Float {
	virtualPrice := bigNumber.NewInt(0)
	if gauge.SwapData.VirtualPrice != nil {
		switch gauge.SwapData.VirtualPrice.(type) {
		case string:
			virtualPrice = bigNumber.NewInt(0).SetString(gauge.SwapData.VirtualPrice.(string))
		case uint64:
			virtualPrice = bigNumber.NewInt(0).SetUint64(gauge.SwapData.VirtualPrice.(uint64))
		}
	}
	return helpers.ToNormalizedAmount(virtualPrice, 18)
}

/**************************************************************************************************
** Retrieve the current gauge boost for a given address. Boost is a multiplier applied to the
** rewards received by the gauge and is between 1 and 2.5x.
**************************************************************************************************/
func getCurveBoost(chainID uint64, voter common.Address, gauge common.Address) *bigNumber.Float {
	calls := []ethereum.Call{}
	calls = append(calls, multicalls.GetCurveWorkingBalance(gauge.Hex(), gauge, voter))
	calls = append(calls, multicalls.GetCurveBalanceOf(gauge.Hex(), gauge, voter))
	response := multicalls.Perform(chainID, calls, nil)
	workingBalance := helpers.DecodeBigInt(response[gauge.Hex()+`working_balances`])
	balanceOf := helpers.DecodeBigInt(response[gauge.Hex()+`balanceOf`])

	if balanceOf.Lte(bigNumber.NewInt(0)) {
		if chainID == 1 {
			return bigNumber.NewFloat(0).SetFloat64(2.5)
		}
		return bigNumber.NewFloat(0).SetFloat64(1)
	}

	boost := bigNumber.NewFloat(0).Div(
		helpers.ToNormalizedAmount(workingBalance, 18),
		bigNumber.NewFloat(0).Mul(
			bigNumber.NewFloat(0).SetFloat64(0.4),
			helpers.ToNormalizedAmount(balanceOf, 18),
		),
	)
	return boost
}

/**************************************************************************************************
** Calculate the rewards APR. Each gauge can get bribes, which are extra rewards on top of the
** normal rewards.
** An unknown amount of unknown tokens can be distributed to the gauge, but we still want to
** calculate the APR for them.
** The list of rewards and their APR is available in the Curve API.
**************************************************************************************************/
func getRewardsAPR(chainID uint64, pool models.CurvePool) *bigNumber.Float {
	totalRewardAPR := bigNumber.NewFloat(0)
	if len(pool.GaugeRewards) == 0 {
		return totalRewardAPR
	}

	for _, reward := range pool.GaugeRewards {
		rewardAPR := bigNumber.NewFloat(0).Div(
			bigNumber.NewFloat(0).SetFloat64(reward.APY),
			bigNumber.NewFloat(100),
		)
		totalRewardAPR = bigNumber.NewFloat(0).Add(totalRewardAPR, rewardAPR)
	}
	return totalRewardAPR
}

/**************************************************************************************************
** Determine the keepCRV value for the vault. Somehow we only check for the first strategy keepCRV
** value.
** If no keepCRV function is available, we can try the keepCrvPercent function.
**************************************************************************************************/
func determineCurveKeepCRV(strategy *models.TStrategy) *bigNumber.Float {
	keepValue := bigNumber.NewInt(0).Add(strategy.KeepCRV, strategy.KeepCRVPercent)
	keepCrv := helpers.ToNormalizedAmount(keepValue, 4)
	return keepCrv
}
