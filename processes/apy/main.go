package apy

import (
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/common/bigNumber"
	"github.com/yearn/ydaemon/common/contracts"
	"github.com/yearn/ydaemon/common/ethereum"
	"github.com/yearn/ydaemon/common/helpers"
	"github.com/yearn/ydaemon/common/logs"
	"github.com/yearn/ydaemon/common/store"
	"github.com/yearn/ydaemon/internal/meta"
	"github.com/yearn/ydaemon/internal/models"
	"github.com/yearn/ydaemon/internal/strategies"
	"github.com/yearn/ydaemon/internal/vaults"
	"github.com/yearn/ydaemon/processes/vaultPricePerShare"
)

var ZERO = bigNumber.NewFloat(0)
var ONE = bigNumber.NewFloat(1)
var WEI = bigNumber.NewFloat(1e18)

var YEARN_VOTER_ADDRESS = common.HexToAddress(`0xF147b8125d2ef93FB6965Db97D6746952a133934`)
var CONVEX_VOTER_ADDRESS = common.HexToAddress(`0x989AEb4d175e16225E39E87d0D97A3360524AD80`)
var CVX_BOOSTER_ADDRESS = common.HexToAddress(`0xF403C135812408BFbE8713b5A23a04b3D48AAE31`)
var CRV_TOKEN_ADDRESS = common.HexToAddress(`0xD533a949740bb3306d119CC777fa900bA034cd52`)
var CVX_TOKEN_ADDRESS = common.HexToAddress(`0x4e3FBD56CD56c3e72c1403e103b45Db9da5B9D2B`)

var COMPUTED_APR = make(map[uint64]map[common.Address]TAPIV1APY)

func calculateGaugeBaseAPR(
	gauge models.CurveGauge,
	crvTokenPrice *bigNumber.Float,
	poolPrice *bigNumber.Float,
	baseAssetPrice *bigNumber.Float,
) *bigNumber.Float {
	inflationRate := helpers.ToNormalizedAmount(bigNumber.NewInt(0).SetString(gauge.GaugeController.InflationRate), 18)
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

	return baseAPR
}

func calculateStrategyAPR(
	vault models.TVault,
	strategy *models.TStrategy,
	gauge models.CurveGauge,
	pool models.CurvePool,
	fraxPool TFraxPool,
	subgraphItem models.CurveSubgraphData,
) TStrategyAPR {
	/**********************************************************************************************
	** First thing is to retrieve the data we need from curve API. This includes the pool, the
	** gauges and some APY data from their subgraph.
	**********************************************************************************************/
	chainID := vault.ChainID

	if subgraphItem.LatestWeeklyApy == 0 {
		logs.Error(`No APY data for vault ` + vault.Address.Hex())
		// logs.Pretty(subgraphItem)
		return TStrategyAPR{}
	}

	/**********************************************************************************************
	** We will need a bunch of prices to calculate the APY.
	** - We get the base asset price from the gauge LpTokenPrice field, which is in base 2.
	** - We get the CRV price from our price package, which is in base 6 but converted in 2.
	** - We get the pool price from the curve API, which is in base 18 but converted in 2.
	**********************************************************************************************/
	baseAssetPrice := bigNumber.NewFloat(0).SetFloat64(gauge.LpTokenPrice)
	crvPrice := getTokenPrice(chainID, CRV_TOKEN_ADDRESS)
	poolPrice := getPoolPrice(gauge)

	/**********************************************************************************************
	** With theses info and the gauges data, we cal calculate both the base APR and the curve
	** boost for the Yearn Voter.
	**********************************************************************************************/
	baseAPR := calculateGaugeBaseAPR(gauge, crvPrice, poolPrice, baseAssetPrice)

	/**********************************************************************************************
	** Once it's done, we need to calculate the rewards APR. Each gauge can get bribes, which are
	** extra rewards on top of the normal rewards.
	** An unknown amount of unknown tokens can be distributed to the gauge, but we still want to
	** calculate the APR for them.
	**********************************************************************************************/
	rewardAPR := getRewardsAPR(chainID, pool)

	/**********************************************************************************************
	** We need to know how is performing the pool in term of APY. We could calculate the whole what
	** is the price now, vs what it was one week ago, and cale it to one year etc. but the weekly
	** apy is already available in the Curve API.
	** So just use it.
	**********************************************************************************************/
	poolWeeklyAPY := getPoolWeeklyAPY(subgraphItem)
	poolDailyAPY := getPoolDailyAPY(subgraphItem)

	/**********************************************************************************************
	** Now, we need to know if we are in a standard curve strat, or a convex strat. If it's a
	** convex strat, we need to calculate the convex APY instead of the standard one.
	** If it's a Frax strategy, we need to calculate the Frax APY instead of the standard one.
	**********************************************************************************************/
	if isFraxStrategy(strategy) {
		return calculateFraxForwardAPR(
			TCalculateFraxAPYDataStruct{
				vault:          vault,
				gaugeAddress:   common.HexToAddress(gauge.Gauge),
				strategy:       strategy,
				baseAssetPrice: baseAssetPrice,
				poolPrice:      poolPrice,
				baseAPR:        baseAPR,
				rewardAPR:      rewardAPR,
				poolDailyAPY:   poolDailyAPY,
			},
			fraxPool,
		)
	}
	if isConvexStrategy(strategy) {
		return calculateConvexForwardAPR(
			TCalculateConvexAPYDataStruct{
				vault:          vault,
				gaugeAddress:   common.HexToAddress(gauge.Gauge),
				strategy:       strategy,
				baseAssetPrice: baseAssetPrice,
				poolPrice:      poolPrice,
				baseAPR:        baseAPR,
				rewardAPR:      rewardAPR,
				poolDailyAPY:   poolDailyAPY,
			},
		)
	}

	return calculateCurveForwardAPR(
		TCalculateCurveAPYDataStruct{
			vault:        vault,
			gaugeAddress: common.HexToAddress(gauge.Gauge),
			strategy:     strategy,
			baseAPR:      baseAPR,
			rewardAPR:    rewardAPR,
			poolAPY:      poolWeeklyAPY,
		},
	)
}

func ComputeChainAPR(chainID uint64) {
	if chainID != 1 {
		return
	}
	allVaults := vaults.ListVaults(chainID)
	gauges := retrieveCurveGauges(chainID)
	pools := retrieveCurveGetPools(chainID)
	subgraphData := retrieveCurveSubgraphData(chainID)
	fraxPools := retrieveFraxPools(chainID)

	if COMPUTED_APR[chainID] == nil {
		COMPUTED_APR[chainID] = make(map[common.Address]TAPIV1APY)
	}

	for _, vault := range allVaults {
		vaultFromMeta, ok := meta.GetMetaVault(chainID, vault.Address)
		if ok {
			if vaultFromMeta.Retired {
				continue
			}
		}

		allStrategiesForVault := strategies.ListStrategiesForVault(chainID, vault.Address)
		ppsPerTime, _ := store.ListPricePerShare(chainID, vault.Address)
		ppsInception := bigNumber.NewFloat(1)
		ppsToday := helpers.GetToday(ppsPerTime, vault.Decimals)
		if ppsToday == nil || ppsToday.IsZero() {
			vaultContract, err := contracts.NewYearnVaultCaller(vault.Address, ethereum.GetRPC(chainID))
			if err != nil {
				continue
			}
			pps, _ := vaultContract.PricePerShare(nil)
			ppsToday = helpers.ToNormalizedAmount(bigNumber.SetInt(pps), vault.Decimals)
		}
		ppsWeekAgo := helpers.GetLastWeek(ppsPerTime, vault.Decimals)
		ppsMonthAgo := helpers.GetLastMonth(ppsPerTime, vault.Decimals)

		vaultPerformanceFee := helpers.ToNormalizedAmount(bigNumber.NewInt(int64(vault.PerformanceFee)), 4)
		vaultManagementFee := helpers.ToNormalizedAmount(bigNumber.NewInt(int64(vault.ManagementFee)), 4)
		oneMinusPerfFee := bigNumber.NewFloat(0).Sub(bigNumber.NewFloat(1), vaultPerformanceFee)
		grossAPR := helpers.GetAPR(ppsToday, ppsMonthAgo, bigNumber.NewFloat(30))
		netAPR := bigNumber.NewFloat(0).Mul(grossAPR, oneMinusPerfFee)
		netAPR = bigNumber.NewFloat(0).Sub(netAPR, vaultManagementFee)

		vaultAPY := TAPIV1APY{
			Type:     "v2:averaged",
			GrossAPR: grossAPR,
			NetAPY:   netAPR,
			Fees: TAPIV1Fees{
				Performance: vaultPerformanceFee,
				Management:  vaultManagementFee,
			},
			Points: TAPIV1Points{
				WeekAgo:   helpers.GetAPR(ppsToday, ppsWeekAgo, bigNumber.NewFloat(7)),
				MonthAgo:  helpers.GetAPR(ppsToday, ppsMonthAgo, bigNumber.NewFloat(30)),
				Inception: helpers.GetAPR(ppsToday, ppsInception, bigNumber.NewFloat(365)),
			},
		}
		if !isCurveVault(allStrategiesForVault) {
			COMPUTED_APR[chainID][vault.Address] = vaultAPY
			continue
		}
		gauge := findGaugeForVault(vault.Token.Address, gauges)
		pool := findPoolForVault(vault.Token.Address, pools)
		fraxPool := findFraxPoolForVault(vault.Token.Address, fraxPools)
		subgraphItem := findSubgraphItemForVault(common.HexToAddress(gauge.Swap), subgraphData)

		for _, strategy := range allStrategiesForVault {
			if strategy.DebtRatio == nil || strategy.DebtRatio.IsZero() {
				logs.Info("Skipping strategy " + strategy.Address.Hex() + " for vault " + vault.Address.Hex() + " because debt ratio is zero")
				continue
			}
			strategyAPR := calculateStrategyAPR(
				vault,
				strategy,
				gauge,
				pool,
				fraxPool,
				subgraphItem,
			)
			vaultAPY.Type += strings.TrimSpace(` ` + strategyAPR.Type)
			vaultAPY.GrossAPR = bigNumber.NewFloat(0).Add(vaultAPY.GrossAPR, strategyAPR.GrossAPR)
			vaultAPY.NetAPY = bigNumber.NewFloat(0).Add(vaultAPY.NetAPY, strategyAPR.NetAPY)
			vaultAPY.Composite.Boost = bigNumber.NewFloat(0).Add(vaultAPY.Composite.Boost, strategyAPR.Composite.Boost)
			vaultAPY.Composite.PoolAPY = bigNumber.NewFloat(0).Add(vaultAPY.Composite.PoolAPY, strategyAPR.Composite.PoolAPY)
			vaultAPY.Composite.BoostedAPR = bigNumber.NewFloat(0).Add(vaultAPY.Composite.BoostedAPR, strategyAPR.Composite.BoostedAPR)
			vaultAPY.Composite.BaseAPR = bigNumber.NewFloat(0).Add(vaultAPY.Composite.BaseAPR, strategyAPR.Composite.BaseAPR)
			vaultAPY.Composite.CvxAPR = bigNumber.NewFloat(0).Add(vaultAPY.Composite.CvxAPR, strategyAPR.Composite.CvxAPR)
			vaultAPY.Composite.RewardsAPR = bigNumber.NewFloat(0).Add(vaultAPY.Composite.RewardsAPR, strategyAPR.Composite.RewardsAPR)
		}
		COMPUTED_APR[chainID][vault.Address] = vaultAPY

		// aggregatedVault, okLegacyAPI := externalVaults.GetAggregatedVault(chainID, vault.Address.Hex())
		// legacyAPY := vaults.BuildAPY(vault, aggregatedVault, okLegacyAPI)
		// spew.Printf("\n%s (%s) - (%s)\n", vault.Name, vault.Address.Hex(), vaultAPY.Type)
		// checkDiff("GrossAPR      ", legacyAPY.GrossAPR, vaultAPY.GrossAPR)
		// checkDiff("NetAPY        ", legacyAPY.NetAPY, vaultAPY.NetAPY)
		// checkDiff("Mngt Fee      ", legacyAPY.Fees.Management, vaultAPY.Fees.Management)
		// checkDiff("Perf Fee      ", legacyAPY.Fees.Performance, vaultAPY.Fees.Performance)
		// checkDiff("PPS Week      ", legacyAPY.Points.WeekAgo, vaultAPY.Points.WeekAgo)
		// checkDiff("PPS Month     ", legacyAPY.Points.MonthAgo, vaultAPY.Points.MonthAgo)
		// checkDiff("PPS Inception ", legacyAPY.Points.Inception, vaultAPY.Points.Inception)
		// checkDiff("BaseAPR       ", legacyAPY.Composite.BaseAPR, vaultAPY.Composite.BaseAPR)
		// checkDiff("Boost         ", legacyAPY.Composite.Boost, vaultAPY.Composite.Boost)
		// checkDiff("BoostedAPR    ", legacyAPY.Composite.BoostedAPR, vaultAPY.Composite.BoostedAPR)
		// checkDiff("PoolAPY       ", legacyAPY.Composite.PoolAPY, vaultAPY.Composite.PoolAPY)
		// checkDiff("CvxAPR        ", legacyAPY.Composite.CvxAPR, vaultAPY.Composite.CvxAPR)
		// checkDiff("RewardsAPR    ", legacyAPY.Composite.RewardsAPR, vaultAPY.Composite.RewardsAPR)
		// spew.Printf("\n")
	}
	logs.Success(`Computing APR done`)
}

func Run(chainID uint64) {
	initYearnEcosystem(chainID)
	vaultPricePerShare.Run(chainID)
	ComputeChainAPR(chainID)
}
