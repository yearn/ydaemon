package apr

import (
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/common/bigNumber"
	"github.com/yearn/ydaemon/common/ethereum"
	"github.com/yearn/ydaemon/common/helpers"
	"github.com/yearn/ydaemon/internal/models"
	"github.com/yearn/ydaemon/internal/multicalls"
	"github.com/yearn/ydaemon/internal/storage"
)

/*
*************************************************************************************************
** Check if all the strategies related to that vault are curve strategies. If any strategies are
not curve strategies, this returns false
** TLDR; check if name contains curve or convex
*************************************************************************************************
*/
func isCurveVault(strategies map[string]models.TStrategy) bool {
	hasActive := false

	for _, strategy := range strategies {
		// skip any strategy that has zero debt or isn’t active
		if strategy.LastDebtRatio == nil ||
			strategy.LastDebtRatio.IsZero() {
			continue
		}
		hasActive = true

		strategyName := strings.ToLower(strategy.Name)
		if !(strings.Contains(strategyName, `curve`) ||
			strings.Contains(strategyName, `convex`)) ||
			strings.Contains(strategyName, `ajna-`) ||
			!strategy.LastDebtRatio.Gt(bigNumber.NewInt(0)) {
			return false // Return false if any active strategy does not meet the criteria
		}
	}
	// if we never saw an active strategy, it’s not a Curve vault
	if !hasActive {
		return false
	}

	return true
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
func getRewardsAPY(chainID uint64, pool models.CurvePool) *bigNumber.Float {
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
func determineCurveKeepCRV(strategy models.TStrategy) *bigNumber.Float {
	keepValue := bigNumber.NewInt(0).Add(strategy.KeepCRV, strategy.KeepCRVPercent)
	keepCrv := helpers.ToNormalizedAmount(keepValue, 4)
	return keepCrv
}

/**********************************************************************************************
** A Curve gauge has a baseAPR which can be scaled by the gauge boost. We can use that function
** to calculate the baseAPR for a given gauge.
**********************************************************************************************/
func calculateGaugeBaseAPR(
	gauge models.CurveGauge,
	crvTokenPrice *bigNumber.Float,
	poolPrice *bigNumber.Float,
	baseAssetPrice *bigNumber.Float,
) (*bigNumber.Float, *bigNumber.Float) {
	inflationRate := bigNumber.NewFloat(0)
	switch gauge.GaugeController.InflationRate.(type) {
	case string:
		inflationRate = helpers.ToNormalizedAmount(bigNumber.NewInt(0).SetString(gauge.GaugeController.InflationRate.(string)), 18)
	case float64:
		inflationRate = bigNumber.NewFloat(0).SetFloat64(gauge.GaugeController.InflationRate.(float64))
	}

	gaugeWeight := helpers.ToNormalizedAmount(bigNumber.NewInt(0).SetString(gauge.GaugeController.GaugeRelativeWeight), 18)
	secondPerYear := bigNumber.NewFloat(0).SetFloat64(31556952)
	workingSupply := helpers.ToNormalizedAmount(bigNumber.NewInt(0).SetString(gauge.GaugeData.WorkingSupply), 18)
	perMaxBoost := bigNumber.NewFloat(0).SetFloat64(0.4)
	crvPrice := bigNumber.NewFloat(0).Clone(crvTokenPrice)

	// baseAPR = (inflationRate * gaugeWeight * (secondPerYear / workingSupply) * (perMaxBoost / poolPrice) * crvPrice) / baseAssetPrice
	baseAPR := bigNumber.NewFloat(0).Mul(inflationRate, gaugeWeight)
	baseAPR = bigNumber.NewFloat(0).Mul(baseAPR, bigNumber.NewFloat(0).Div(secondPerYear, workingSupply))
	baseAPR = bigNumber.NewFloat(0).Mul(baseAPR, bigNumber.NewFloat(0).Div(perMaxBoost, poolPrice))
	baseAPR = bigNumber.NewFloat(0).Mul(baseAPR, crvPrice)
	baseAPR = bigNumber.NewFloat(0).Div(baseAPR, baseAssetPrice)
	baseAPY := bigNumber.NewFloat(0).Add(bigNumber.NewFloat(0), baseAPR)

	return baseAPR, baseAPY
}

/**********************************************************************************************
** For a given strategy, we need to calculate the APR based on which curveLike strategy it is,
** aka if it's vanilla Curve, Convex or Convex and Frax.
**********************************************************************************************/
func calculateCurveLikeStrategyAPR(
	vault models.TVault,
	strategy models.TStrategy,
	gauge models.CurveGauge,
	pool models.CurvePool,
	fraxPool TFraxPool,
	subgraphItem models.CurveSubgraphData,
) TStrategyAPY {
	/**********************************************************************************************
	** First thing is to retrieve the data we need from curve API. This includes the pool, the
	** gauges and some APY data from their subgraph.
	**********************************************************************************************/
	chainID := vault.ChainID

	// if subgraphItem.LatestWeeklyApy == 0 {
	// logs.Warning(`No APY data for vault ` + vault.Address.Hex())
	// }

	/**********************************************************************************************
	** We will need a bunch of prices to calculate the APY.
	** - We get the base asset price from the gauge LpTokenPrice field, which is in base 2.
	** - We get the CRV price from our price package, which is in base 6 but converted in 2.
	** - We get the pool price from the curve API, which is in base 18 but converted in 2.
	**********************************************************************************************/
	baseAssetPrice := bigNumber.NewFloat(0).SetFloat64(gauge.LpTokenPrice)
	crvPrice := bigNumber.NewFloat(0)
	if tokenPrice, ok := storage.GetPrice(chainID, storage.CRV_TOKEN_ADDRESS[chainID]); ok {
		crvPrice = tokenPrice.HumanizedPrice
	}
	poolPrice := getPoolPrice(gauge)

	/**********************************************************************************************
	** With theses info and the gauges data, we cal calculate both the base APR and the curve
	** boost for the Yearn Voter.
	**********************************************************************************************/
	_, baseAPY := calculateGaugeBaseAPR(gauge, crvPrice, poolPrice, baseAssetPrice)

	/**********************************************************************************************
	** Once it's done, we need to calculate the rewards APR. Each gauge can get bribes, which are
	** extra rewards on top of the normal rewards.
	** An unknown amount of unknown tokens can be distributed to the gauge, but we still want to
	** calculate the APR for them.
	**********************************************************************************************/
	rewardAPY := getRewardsAPY(chainID, pool)

	/**********************************************************************************************
	** We need to know how is performing the pool in term of APY. We could calculate the whole what
	** is the price now, vs what it was one week ago, and cale it to one year etc. but the weekly
	** apy is already available in the Curve API.
	** So just use it.
	**********************************************************************************************/
	poolWeeklyAPY := getPoolWeeklyAPY(subgraphItem)

	/**********************************************************************************************
	** Now, we need to know if we are in a standard curve strat, or a convex strat. If it's a
	** convex strat, we need to calculate the convex APY instead of the standard one.
	** If it's a Frax strategy, we need to calculate the Frax APY instead of the standard one.
	** If it's a Prisma strategy, we need to calculate the Prisma APY instead of the previous ones.
	**********************************************************************************************/
	if isPrismaStrategy(strategy) {
		return calculatePrismaForwardAPR(
			TCalculatePrismaAPYDataStruct{
				vault:          vault,
				gaugeAddress:   common.HexToAddress(gauge.Gauge),
				strategy:       strategy,
				baseAssetPrice: baseAssetPrice,
				poolPrice:      poolPrice,
				baseAPY:        baseAPY,
				rewardAPY:      rewardAPY,
				poolWeeklyAPY:  poolWeeklyAPY,
			},
		)
	}
	if isFraxStrategy(strategy) {
		return calculateFraxForwardAPR(
			TCalculateFraxAPYDataStruct{
				vault:          vault,
				gaugeAddress:   common.HexToAddress(gauge.Gauge),
				strategy:       strategy,
				baseAssetPrice: baseAssetPrice,
				poolPrice:      poolPrice,
				baseAPY:        baseAPY,
				rewardAPY:      rewardAPY,
				poolWeeklyAPY:  poolWeeklyAPY,
			},
			fraxPool,
		)
	}
	if isConvexStrategy(strategy) {
		return calculateConvexForwardAPY(
			TCalculateConvexAPYDataStruct{
				vault:          vault,
				gaugeAddress:   common.HexToAddress(gauge.Gauge),
				strategy:       strategy,
				baseAssetPrice: baseAssetPrice,
				poolPrice:      poolPrice,
				baseAPY:        baseAPY,
				rewardAPY:      rewardAPY,
				poolWeeklyAPY:  poolWeeklyAPY,
			},
		)
	}

	return calculateCurveForwardAPY(
		TCalculateCurveAPYDataStruct{
			vault:        vault,
			gaugeAddress: common.HexToAddress(gauge.Gauge),
			strategy:     strategy,
			baseAPY:      baseAPY,
			rewardAPY:    rewardAPY,
			poolAPY:      poolWeeklyAPY,
		},
	)
}

/**************************************************************************************************
** If the vault is a curve vault, a convex vault or a frax vault, we can calculate the forward APR
** for it, using always the same base formula, which require fetching some elements like:
** - The gauge
** - The pool
** - The subgraph data
** - The frax pool
**************************************************************************************************/
func computeCurveLikeForwardAPY(
	vault models.TVault,
	allStrategiesForVault map[string]models.TStrategy,
	gauges []models.CurveGauge,
	pools []models.CurvePool,
	subgraphData []models.CurveSubgraphData,
	fraxPools []TFraxPool,
) TForwardAPY {
	gauge := findGaugeForVault(vault.AssetAddress, gauges)
	pool := findPoolForVault(vault.AssetAddress, pools)
	fraxPool := findFraxPoolForVault(vault.AssetAddress, fraxPools)
	subgraphItem := findSubgraphItemForVault(common.HexToAddress(gauge.Swap), subgraphData)

	/**********************************************************************************************
	** If we can't resolve the gauge or pool for this vault, bail out so we don't clobber any
	** previously computed forward APY (e.g. oracle based values).
	**********************************************************************************************/
	if gauge.Gauge == "" || pool.Address == "" {
		return models.TForwardAPY{}
	}

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
	for _, strategy := range allStrategiesForVault {
		if strategy.LastDebtRatio == nil || strategy.LastDebtRatio.IsZero() {
			continue
		}

		strategyAPR := calculateCurveLikeStrategyAPR(
			vault,
			strategy,
			gauge,
			pool,
			fraxPool,
			subgraphItem,
		)
		TypeOf += strings.TrimSpace(` ` + strategyAPR.Type)
		netAPY = bigNumber.NewFloat(0).Add(netAPY, strategyAPR.NetAPY)
		boost = bigNumber.NewFloat(0).Add(boost, strategyAPR.Composite.Boost)
		poolAPY = bigNumber.NewFloat(0).Add(poolAPY, strategyAPR.Composite.PoolAPY)
		boostedAPR = bigNumber.NewFloat(0).Add(boostedAPR, strategyAPR.Composite.BoostedAPR)
		baseAPR = bigNumber.NewFloat(0).Add(baseAPR, strategyAPR.Composite.BaseAPR)
		cvxAPR = bigNumber.NewFloat(0).Add(cvxAPR, strategyAPR.Composite.CvxAPR)
		rewardsAPY = bigNumber.NewFloat(0).Add(rewardsAPY, strategyAPR.Composite.RewardsAPY)
		keepCRV = bigNumber.NewFloat(0).Add(keepCRV, strategyAPR.Composite.KeepCRV)
		keepVelo = bigNumber.NewFloat(0).Add(keepVelo, strategyAPR.Composite.KeepVelo)
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
