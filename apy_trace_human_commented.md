APY Trace Review (Commented Code)

Generated: 2026-01-16T14:29:41.180657+00:00

Source log: `apy_trace.log`

## `processes/apr/forward.curve.go`

```go
package apr

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/common/bigNumber"
	"github.com/yearn/ydaemon/common/helpers"
	"github.com/yearn/ydaemon/internal/models"
	"github.com/yearn/ydaemon/internal/storage"
)

type TCalculateCurveAPYDataStruct struct {
	vault        models.TVault
	gaugeAddress common.Address
	strategy     models.TStrategy
	baseAPY      *bigNumber.Float
	rewardAPY    *bigNumber.Float
	poolAPY      *bigNumber.Float
}

func calculateCurveForwardAPY(args TCalculateCurveAPYDataStruct) TStrategyAPY {
	chainID := args.vault.ChainID

	/**********************************************************************************************
	** We first need to retrieve a bunch to be able to proceed:
	** - the yBoost, aka how much boost the Yearn voter has for this gauge
	** - the performanceFee and the managementFee for that vault
	** - the debtRatio for the strategy (aka the % of fund allocated to the strategy by the vault)
	**********************************************************************************************/
	yBoost := getCurveBoost(chainID, storage.YEARN_VOTER_ADDRESS[chainID], args.gaugeAddress)
	keepCrv := determineCurveKeepCRV(args.strategy)
	debtRatio := helpers.ToNormalizedAmount(args.strategy.LastDebtRatio, 4)
	vaultPerformanceFee := helpers.ToNormalizedAmount(bigNumber.NewInt(int64(args.vault.PerformanceFee)), 4)
	vaultManagementFee := helpers.ToNormalizedAmount(bigNumber.NewInt(int64(args.vault.ManagementFee)), 4)
	oneMinusPerfFee := bigNumber.NewFloat(0).Sub(bigNumber.NewFloat(1), vaultPerformanceFee)

	/**********************************************************************************************
	** The CRV APR is simply the baseAPR (aka how much CRV we get from the gauge) scaled by the
	** yBoost. We then add the extraRewards which are incentives/bribes on top of the base rewards.
	**********************************************************************************************/
	crvAPY := bigNumber.NewFloat(0).Mul(args.baseAPY, yBoost) // baseAPR * yBoost
	crvAPY = bigNumber.NewFloat(0).Add(crvAPY, args.rewardAPY) // (baseAPR * yBoost) + rewardAPY

	/**********************************************************************************************
	** Calculate the CRV Gross APR:
	** 1. Taking the base APR, scaling it with the boost and removing the percentage of CRV we want
	** to keep
	** 2. Adding the rewards APR
	** 3. Adding the pool APY
	**********************************************************************************************/
	keepCRVRatio := bigNumber.NewFloat(0).Sub(storage.ONE, keepCrv) // 1 - keepCRV
	grossAPY := bigNumber.NewFloat(0).Mul(args.baseAPY, yBoost) // 1 - baseAPR * yBoost
	grossAPY = bigNumber.NewFloat(0).Mul(grossAPY, keepCRVRatio) // 1 - baseAPR * yBoost * keepCRV
	grossAPY = bigNumber.NewFloat(0).Add(grossAPY, args.rewardAPY) // 2 - (baseAPR * yBoost * keepCRV) + rewardAPY

	/**********************************************************************************************
	** Calculate the CRV Net APR:
	** Take the gross APR and remove the performance fee and the management fee
	**********************************************************************************************/
	netAPY := bigNumber.NewFloat(0).Mul(grossAPY, oneMinusPerfFee) // grossAPY * (1 - perfFee)
	if netAPY.Gt(vaultManagementFee) {
		netAPY = bigNumber.NewFloat(0).Sub(netAPY, vaultManagementFee) // (grossAPY * (1 - perfFee)) - managementFee
		netAPRFloat64, _ := netAPY.Float64()
		netAPY = bigNumber.NewFloat(0).SetFloat64(convertFloatAPRToAPY(netAPRFloat64, 52))
		netAPY = bigNumber.NewFloat(0).Add(netAPY, args.poolAPY)
	} else {
		netAPY = bigNumber.NewFloat(0).Add(bigNumber.NewFloat(0), args.poolAPY)
	}

	apyStruct := TStrategyAPY{
		Type:      "crv",
		DebtRatio: debtRatio,
		NetAPY:    bigNumber.NewFloat(0).Mul(netAPY, debtRatio),
		Composite: TCompositeData{
			Boost:      bigNumber.NewFloat(0).Mul(yBoost, debtRatio),
			PoolAPY:    bigNumber.NewFloat(0).Mul(args.poolAPY, debtRatio),
			BoostedAPR: bigNumber.NewFloat(0).Mul(crvAPY, debtRatio),
			BaseAPR:    bigNumber.NewFloat(0).Mul(args.baseAPY, debtRatio),
			RewardsAPY: bigNumber.NewFloat(0).Mul(args.rewardAPY, debtRatio),
			KeepCRV:    keepCrv,
		},
	}
	return apyStruct
}
```

## `processes/apr/forward.convex.go`

```go
package apr

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/common/bigNumber"
	"github.com/yearn/ydaemon/common/helpers"
	"github.com/yearn/ydaemon/internal/models"
	"github.com/yearn/ydaemon/internal/storage"
)

type TCalculateConvexAPYDataStruct struct {
	vault          models.TVault
	gaugeAddress   common.Address
	strategy       models.TStrategy
	baseAssetPrice *bigNumber.Float
	poolPrice      *bigNumber.Float
	baseAPY        *bigNumber.Float
	rewardAPY      *bigNumber.Float
	poolWeeklyAPY  *bigNumber.Float
}

func calculateConvexForwardAPY(args TCalculateConvexAPYDataStruct) TStrategyAPY {
	chainID := args.vault.ChainID
	// Curve ZUSD-FRAXBP Factory | args.gaugeAddress.Hex() | 0x218E4678318ab5527e41135713193E5EAd73337f
	// Curve XAI-FRAXBP Factory | args.gaugeAddress.Hex() | 0xa8Ea11465A1375BF42463C3B613dFC54248b9C7B
	// Curve ZUSD-FRAXBP Factory | args.baseAssetPrice | 0.018355142896898657
	// Curve XAI-FRAXBP Factory | args.baseAssetPrice | 0.06417867250402262
	// Curve ZUSD-FRAXBP Factory | args.poolPrice | 1.0110440176923634
	// Curve XAI-FRAXBP Factory | args.poolPrice | 1.0080959363340094
	// Curve ZUSD-FRAXBP Factory | args.baseAPY | 3.851825336994172
	// Curve XAI-FRAXBP Factory | args.baseAPY | 0.2975187046483139
	// Curve ZUSD-FRAXBP Factory | args.rewardAPY | 0
	// Curve XAI-FRAXBP Factory | args.rewardAPY | 0
	// Curve ZUSD-FRAXBP Factory | args.poolWeeklyAPY | 0.0001
	// Curve XAI-FRAXBP Factory | args.poolWeeklyAPY | 0.0002

	/**********************************************************************************************
	** We first need to retrieve a bunch to be able to proceed:
	** - the cvxBoost, aka how much boost the convex voter has for this gauge
	** - the performanceFee and the managementFee for that vault
	** - the debtRatio for the strategy (aka the % of fund allocated to the strategy by the vault)
	**********************************************************************************************/
	cvxBoost := getCurveBoost(chainID, storage.CONVEX_VOTER_ADDRESS[chainID], args.gaugeAddress)
	// Curve ZUSD-FRAXBP Factory | cvxBoost | 2.063936749763785
	// Curve XAI-FRAXBP Factory | cvxBoost | 2.5
	keepCrv := determineConvexKeepCRV(args.strategy)
	// Curve ZUSD-FRAXBP Factory | keepCrv | 0
	// Curve XAI-FRAXBP Factory | keepCrv | 0
	debtRatio := helpers.ToNormalizedAmount(args.strategy.LastDebtRatio, 4)
	// Curve ZUSD-FRAXBP Factory | debtRatio | 0.8
	// Curve XAI-FRAXBP Factory | debtRatio | 0.8
	vaultPerformanceFee := helpers.ToNormalizedAmount(bigNumber.NewInt(int64(args.vault.PerformanceFee)), 4)
	// Curve ZUSD-FRAXBP Factory | vaultPerformanceFee | 0.1
	// Curve XAI-FRAXBP Factory | vaultPerformanceFee | 0.1
	vaultManagementFee := helpers.ToNormalizedAmount(bigNumber.NewInt(int64(args.vault.ManagementFee)), 4)
	// Curve ZUSD-FRAXBP Factory | vaultManagementFee | 0
	// Curve XAI-FRAXBP Factory | vaultManagementFee | 0
	oneMinusPerfFee := bigNumber.NewFloat(0).Sub(bigNumber.NewFloat(1), vaultPerformanceFee)
	// Curve ZUSD-FRAXBP Factory | oneMinusPerfFee | 0.9
	// Curve XAI-FRAXBP Factory | oneMinusPerfFee | 0.9

	/**********************************************************************************************
	** The CRV APR is simply the baseAPY (aka how much CRV we get from the gauge) not based on
	** curve API but directly from the formula used by convex subgraph (link below). We are using
	** this for th APR to match convex APR and avoid confusion. This should be very close to the
	** one we could compute.
	** The CVX APR is the APR we get from the CVX rewards, based on the CVX price and the amount
	** of CVX printed, based on the CRV rate for the given gauge. Tldr X crv = Y cvx and we do
	** something to gt an APR.
	**********************************************************************************************/
	crvAPR, cvxAPR, crvAPY, cvxAPY := getCVXPoolAPY(chainID, args.vault.Address, args.strategy.Address, args.baseAssetPrice)
	// Curve ZUSD-FRAXBP Factory | crvAPR | 105.74520039424516
	// Curve XAI-FRAXBP Factory | crvAPR | 9.332241744751284
	// Curve ZUSD-FRAXBP Factory | cvxAPR | 0.2670971054725635
	// Curve XAI-FRAXBP Factory | cvxAPR | 0.02357189497301238
	// Curve ZUSD-FRAXBP Factory | crvAPY | 105.74520039424516
	// Curve XAI-FRAXBP Factory | crvAPY | 9.332241744751284
	// Curve ZUSD-FRAXBP Factory | cvxAPY | 0.2670971054725635
	// Curve XAI-FRAXBP Factory | cvxAPY | 0.02357189497301238

	/**********************************************************************************************
	** Just like curve, Convex can have extra rewards which are incentives/bribes on top of the
	** base rewards. We need to retrieve them.
	**********************************************************************************************/
	_, rewardsAPY := getConvexRewardAPY(chainID, args.strategy, args.baseAssetPrice, args.poolPrice)
	// Curve ZUSD-FRAXBP Factory | rewardsAPY | 0
	// Curve XAI-FRAXBP Factory | rewardsAPY | 0

	/**********************************************************************************************
	** Calculate the CRV Gross APR:
	** 1. Taking the base APR and removing the percentage of CRV we want to keep
	** 2. Adding the rewards APR
	** 3. Adding the pool APY
	** 4. Adding the CVX APR
	**********************************************************************************************/
	keepCRVRatio := bigNumber.NewFloat(0).Sub(storage.ONE, keepCrv) // 1 - keepCRV
	// Curve ZUSD-FRAXBP Factory | keepCRVRatio | 1
	// Curve XAI-FRAXBP Factory | keepCRVRatio | 1
	grossAPY := bigNumber.NewFloat(0).Mul(crvAPY, keepCRVRatio)     // 1 - baseAPY * keepCRV
	// Curve ZUSD-FRAXBP Factory | grossAPY | 105.74520039424516
	// Curve XAI-FRAXBP Factory | grossAPY | 9.332241744751284
	grossAPY = bigNumber.NewFloat(0).Add(grossAPY, rewardsAPY)      // 2 - (baseAPY * keepCRV) + rewardAPR
	// Curve ZUSD-FRAXBP Factory | grossAPY | 105.74520039424516
	// Curve XAI-FRAXBP Factory | grossAPY | 9.332241744751284
	grossAPY = bigNumber.NewFloat(0).Add(grossAPY, cvxAPY)          // 4 - (baseAPY * keepCRV) + rewardAPR + poolAPY + cvxAPR
	// Curve ZUSD-FRAXBP Factory | grossAPY | 106.01229749971772
	// Curve XAI-FRAXBP Factory | grossAPY | 9.355813639724296

	/**********************************************************************************************
	** Calculate the CRV Net APR:
	** Take the gross APR and remove the performance fee and the management fee
	**********************************************************************************************/
	netAPY := bigNumber.NewFloat(0).Mul(grossAPY, oneMinusPerfFee) // grossAPR * (1 - perfFee)
	// Curve ZUSD-FRAXBP Factory | netAPY | 95.41106774974595
	// Curve XAI-FRAXBP Factory | netAPY | 8.420232275751866
	if netAPY.Gt(vaultManagementFee) {
		netAPY = bigNumber.NewFloat(0).Sub(netAPY, vaultManagementFee) // (grossAPR * (1 - perfFee)) - managementFee
		// Curve ZUSD-FRAXBP Factory | netAPY | 95.41106774974595
		// Curve XAI-FRAXBP Factory | netAPY | 8.420232275751866
		netAPRFloat64, _ := netAPY.Float64()
		// Curve ZUSD-FRAXBP Factory | netAPRFloat64 | 95.41106774974595
		// Curve XAI-FRAXBP Factory | netAPRFloat64 | 8.420232275751866
		netAPY = bigNumber.NewFloat(0).SetFloat64(convertFloatAPRToAPY(netAPRFloat64, 52))
		// Curve ZUSD-FRAXBP Factory | netAPY | 3.3993111000835647e+23
		// Curve XAI-FRAXBP Factory | netAPY | 2449.814481037776
		netAPY = bigNumber.NewFloat(0).Add(netAPY, args.poolWeeklyAPY)
		// Curve ZUSD-FRAXBP Factory | netAPY | 3.3993111000835647e+23
		// Curve XAI-FRAXBP Factory | netAPY | 2449.814681037776
	} else {
		netAPY = bigNumber.NewFloat(0).Add(bigNumber.NewFloat(0), args.poolWeeklyAPY)
	}

	netAPYWithDebtRatio := bigNumber.NewFloat(0).Mul(netAPY, debtRatio)
	// Curve ZUSD-FRAXBP Factory | netAPY | 3.3993111000835647e+23
	// Curve XAI-FRAXBP Factory | netAPY | 2449.814681037776
	// Curve ZUSD-FRAXBP Factory | netAPYWithDebtRatio | 2.7194488800668517e+23
	// Curve XAI-FRAXBP Factory | netAPYWithDebtRatio | 1959.851744830221
	boostWithDebtRatio := bigNumber.NewFloat(0).Mul(cvxBoost, debtRatio)
	// Curve ZUSD-FRAXBP Factory | boostWithDebtRatio | 1.651149399811028
	// Curve XAI-FRAXBP Factory | boostWithDebtRatio | 2
	poolAPYWithDebtRatio := bigNumber.NewFloat(0).Mul(args.poolWeeklyAPY, debtRatio)
	// Curve ZUSD-FRAXBP Factory | poolAPYWithDebtRatio | 8e-05
	// Curve XAI-FRAXBP Factory | poolAPYWithDebtRatio | 0.00016
	boostedAPRWithDebtRatio := bigNumber.NewFloat(0).Mul(crvAPR, debtRatio)
	// Curve ZUSD-FRAXBP Factory | boostedAPRWithDebtRatio | 84.59616031539613
	// Curve XAI-FRAXBP Factory | boostedAPRWithDebtRatio | 7.465793395801028
	baseAPRWithDebtRatio := bigNumber.NewFloat(0).Mul(args.baseAPY, debtRatio)
	// Curve ZUSD-FRAXBP Factory | baseAPRWithDebtRatio | 3.081460269595338
	// Curve XAI-FRAXBP Factory | baseAPRWithDebtRatio | 0.23801496371865116
	cvxAPRWithDebtRatio := bigNumber.NewFloat(0).Mul(cvxAPR, debtRatio)
	// Curve ZUSD-FRAXBP Factory | cvxAPRWithDebtRatio | 0.2136776843780508
	// Curve XAI-FRAXBP Factory | cvxAPRWithDebtRatio | 0.018857515978409903
	rewardsAPYWithDebtRatio := bigNumber.NewFloat(0).Mul(args.rewardAPY, debtRatio)
	// Curve ZUSD-FRAXBP Factory | rewardsAPYWithDebtRatio | 0
	// Curve XAI-FRAXBP Factory | rewardsAPYWithDebtRatio | 0

	apyStruct := TStrategyAPY{
		Type:      "convex",
		DebtRatio: debtRatio,
		NetAPY:    netAPYWithDebtRatio,
		Composite: TCompositeData{
			Boost:      boostWithDebtRatio,
			PoolAPY:    poolAPYWithDebtRatio,
			BoostedAPR: boostedAPRWithDebtRatio,
			BaseAPR:    baseAPRWithDebtRatio,
			CvxAPR:     cvxAPRWithDebtRatio,
			RewardsAPY: rewardsAPYWithDebtRatio,
			KeepCRV:    keepCrv,
		},
	}
	// Curve ZUSD-FRAXBP Factory | apyStruct | {convex 0.8 2.7194488800668517e+23 {1.651149399811028 8e-05 84.59616031539613 3.081460269595338 0.2136776843780508 0 <nil> <nil> 0 <nil>}}
	// Curve XAI-FRAXBP Factory | apyStruct | {convex 0.8 1959.851744830221 {2 0.00016 7.465793395801028 0.23801496371865116 0.018857515978409903 0 <nil> <nil> 0 <nil>}}
	return apyStruct
}
```

## `processes/apr/forward.convex.helpers.go`

```go
package apr

import (
	"math/big"
	"os"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/common/bigNumber"
	"github.com/yearn/ydaemon/common/contracts"
	"github.com/yearn/ydaemon/common/ethereum"
	"github.com/yearn/ydaemon/common/helpers"
	"github.com/yearn/ydaemon/common/logs"
	"github.com/yearn/ydaemon/internal/models"
	"github.com/yearn/ydaemon/internal/storage"
)

/**************************************************************************************************
** The cumulative apr of all extra tokens that are emitted by depositing to Convex, assuming that
** they will be sold for profit.
** We need to pull data from convex's virtual rewards contracts to get bonus rewards
**************************************************************************************************/
func getConvexRewardAPY(
	chainID uint64,
	strategy models.TStrategy,
	baseAssetPrice *bigNumber.Float,
	poolPrice *bigNumber.Float,
) (*bigNumber.Float, *bigNumber.Float) {
	client := ethereum.GetRPC(chainID)
	convexStrategyContract, _ := contracts.NewConvexBaseStrategy(strategy.Address, client)
	cvxBoosterContract, _ := contracts.NewCVXBooster(storage.CVX_BOOSTER_ADDRESS[chainID], client)
	rewardPID, err := convexStrategyContract.Pid(nil)
	if err != nil {
		rewardPID, err = convexStrategyContract.ID(nil)
		if err != nil {
			warnMissingPID := func(pidErr error) {
				logs.Warning("Convex PID not found for strategy "+strategy.Address.Hex(), pidErr)
			}
			fraxBaseStrategy, err := contracts.NewFraxBaseStrategy(strategy.Address, client)
			if err != nil {
				if os.Getenv("ENVIRONMENT") == "dev" {
					logs.Error(`Unable to init fraxBaseStrategy for convex strategy `+strategy.Address.Hex(), err)
				}
				warnMissingPID(err)
				return storage.ZERO, storage.ZERO
			}
			userVaultAddress, err := fraxBaseStrategy.UserVault(nil)
			if err != nil {
				if os.Getenv("ENVIRONMENT") == "dev" {
					logs.Error(`Unable to get userVault for fraxBaseStrategy `+strategy.Address.Hex(), err)
				}
				warnMissingPID(err)
				return storage.ZERO, storage.ZERO
			}
			userVaultContract, err := contracts.NewConvexUserVault(userVaultAddress, client)
			if err != nil {
				if os.Getenv("ENVIRONMENT") == "dev" {
					logs.Error(`Unable to init userVault contract `+userVaultAddress.Hex(), err)
				}
				warnMissingPID(err)
				return storage.ZERO, storage.ZERO
			}
			stakingTokenAddress, err := userVaultContract.StakingToken(nil)
			if err != nil {
				if os.Getenv("ENVIRONMENT") == "dev" {
					logs.Error(`Unable to get stakingToken for userVault `+userVaultAddress.Hex(), err)
				}
				warnMissingPID(err)
				return storage.ZERO, storage.ZERO
			}
			stakingTokenContract, err := contracts.NewConvexStakingToken(stakingTokenAddress, client)
			if err != nil {
				if os.Getenv("ENVIRONMENT") == "dev" {
					logs.Error(`Unable to init stakingToken contract `+stakingTokenAddress.Hex(), err)
				}
				warnMissingPID(err)
				return storage.ZERO, storage.ZERO
			}
			rewardPID, err = stakingTokenContract.ConvexPoolId(nil)
			if err != nil {
				if os.Getenv("ENVIRONMENT") == "dev" {
					logs.Error(`Unable to get cvxPoolId for stakingToken `+stakingTokenAddress.Hex(), err)
				}
				warnMissingPID(err)
				return storage.ZERO, storage.ZERO
			}
		}
	}
	rewardContract, err := cvxBoosterContract.PoolInfo(nil, rewardPID)
	if err != nil {
		logs.Error(err)
		return storage.ZERO, storage.ZERO
	}
	crvRewardContract, _ := contracts.NewCrvRewards(rewardContract.CrvRewards, client)
	rewardsLength, _ := crvRewardContract.ExtraRewardsLength(nil)

	now := time.Now().Unix()
	totalRewardsAPR := bigNumber.NewFloat(0)
	if rewardsLength != nil {
		for i := 0; i < int(rewardsLength.Int64()); i++ {
			virtualRewardsPool, err := crvRewardContract.ExtraRewards(nil, big.NewInt(int64(i)))
			if err != nil {
				logs.Error(err)
				continue
			}
			virtualRewardsPoolContract, _ := contracts.NewCrvRewards(virtualRewardsPool, client)
			periodFinish, err := virtualRewardsPoolContract.PeriodFinish(nil)
			if err != nil {
				logs.Error(err)
				continue
			}
			if periodFinish.Int64() < now {
				continue
			}
			rewardToken, _ := virtualRewardsPoolContract.RewardToken(nil)
			rewardTokenPrice, ok := storage.GetPrice(chainID, rewardToken)
			if !ok {
				continue
			}
			rewardRateInt, _ := virtualRewardsPoolContract.RewardRate(nil)
			totalSupplyInt, _ := virtualRewardsPoolContract.TotalSupply(nil)

			tokenPrice := rewardTokenPrice.HumanizedPrice
			rewardRate := helpers.ToNormalizedAmount(bigNumber.NewInt(0).Set(rewardRateInt), 18)
			totalSupply := helpers.ToNormalizedAmount(bigNumber.NewInt(0).Set(totalSupplyInt), 18)
			secondPerYear := bigNumber.NewFloat(0).SetFloat64(31556952)

			rewardAPRTop := bigNumber.NewFloat(0).Mul(rewardRate, secondPerYear)
			rewardAPRTop = bigNumber.NewFloat(0).Mul(rewardAPRTop, tokenPrice)
			rewardAPRBottom := bigNumber.NewFloat(0).Div(poolPrice, storage.ONE) //wei?
			rewardAPRBottom = bigNumber.NewFloat(0).Mul(rewardAPRBottom, baseAssetPrice)
			rewardAPRBottom = bigNumber.NewFloat(0).Mul(rewardAPRBottom, totalSupply)
			rewardAPR := bigNumber.NewFloat(0).Div(rewardAPRTop, rewardAPRBottom)
			totalRewardsAPR = bigNumber.NewFloat(0).Add(totalRewardsAPR, rewardAPR)
		}
	}
	totalRewardsAPY := bigNumber.NewFloat(0).Add(bigNumber.NewFloat(0), totalRewardsAPR)
	return totalRewardsAPR, totalRewardsAPY
}

/**************************************************************************************************
** Calculate the number of CVX minted for a given amount of CRV earned.
** See formula here: https://docs.convexfinance.com/convexfinanceintegration/cvx-minting
**************************************************************************************************/
func getCVXForCRV(chainID uint64, crvEarned *bigNumber.Float) *bigNumber.Float {
	cliffSize := bigNumber.NewFloat(0).SetString(`100000000000000000000000`)    //1e23
	cliffCount := bigNumber.NewFloat(0).SetString(`1000`)                       //1e3
	maxSupply := bigNumber.NewFloat(0).SetString(`100000000000000000000000000`) //1e26

	cvxContract, _ := contracts.NewERC20(storage.CVX_TOKEN_ADDRESS[chainID], ethereum.GetRPC(chainID))
	cvxTotalSupplyInt, _ := cvxContract.TotalSupply(nil)
	cvxTotalSupply := bigNumber.NewFloat(0).SetInt(bigNumber.NewInt(0).Set(cvxTotalSupplyInt))
	currentCliff := bigNumber.NewFloat(0).Div(cvxTotalSupply, cliffSize)
	if currentCliff.Gte(cliffCount) {
		return storage.ZERO
	}
	remaining := bigNumber.NewFloat(0).Sub(cliffCount, currentCliff)
	cvxEarned := bigNumber.NewFloat(0).Mul(crvEarned, remaining)
	cvxEarned = bigNumber.NewFloat(0).Div(cvxEarned, cliffCount)

	amountTillMax := bigNumber.NewFloat(0).Sub(maxSupply, cvxTotalSupply)
	if cvxEarned.Gt(amountTillMax) {
		cvxEarned = amountTillMax
	}
	return cvxEarned
}

/**************************************************************************************************
** Calculate the cvxAPR for a given CVX pool.
** See formula here:
** https://github.com/convex-community/convex-subgraph/blob/
** 13dbb4e3f3f69c6762fecb1ebc46f477162e2093/subgraphs/convex/src/services/pools.ts#L269-L289
**************************************************************************************************/
func getCVXPoolAPY(
	chainID uint64,
	vaultAddress common.Address,
	strategyAddress common.Address,
	virtualPoolPrice *bigNumber.Float,
) (*bigNumber.Float, *bigNumber.Float, *bigNumber.Float, *bigNumber.Float) {
	// Curve ZUSD-FRAXBP Factory | virtualPoolPrice | 0.018355142896898657
	// Curve XAI-FRAXBP Factory | virtualPoolPrice | 0.06417867250402262

	crvAPR := bigNumber.NewFloat(0)
	cvxAPR := bigNumber.NewFloat(0)
	crvAPY := bigNumber.NewFloat(0)
	cvxAPY := bigNumber.NewFloat(0)

	/**********************************************************************************************
	** First thing to do to be able to calculate the APR is to retrieve the `crvRewards` contract
	** for this given convex strategy. It's a multiple step with way too many contracts involved,
	** but in the end we should be able to query the `rewardRate` for it.
	***********************************************************************************************/
	client := ethereum.GetRPC(chainID)
	convexStrategyContract, err := contracts.NewConvexBaseStrategy(strategyAddress, client)
	if err != nil {
		return crvAPR, cvxAPR, crvAPY, cvxAPY
	}

	/**********************************************************************************************
	** We need to know the PID of the pool. Based on the contract, it can be PID or ID, so we need
	** to try both, and one call can revert.
	***********************************************************************************************/
	rewardPID, err := convexStrategyContract.Pid(nil)
	if err != nil {
		rewardPID, err = convexStrategyContract.ID(nil)
		if err != nil {
			warnMissingPID := func(pidErr error) {
				logs.Warning("Convex PID not found for vault "+vaultAddress.Hex()+" strategy "+strategyAddress.Hex(), pidErr)
			}
			fraxBaseStrategy, err := contracts.NewFraxBaseStrategy(strategyAddress, client)
			if err != nil {
				if os.Getenv("ENVIRONMENT") == "dev" {
					logs.Error(`Unable to init fraxBaseStrategy for convex strategy `+strategyAddress.Hex(), err)
				}
				warnMissingPID(err)
				return crvAPR, cvxAPR, crvAPY, cvxAPY
			}
			userVaultAddress, err := fraxBaseStrategy.UserVault(nil)
			if err != nil {
				if os.Getenv("ENVIRONMENT") == "dev" {
					logs.Error(`Unable to get userVault for fraxBaseStrategy `+strategyAddress.Hex(), err)
				}
				warnMissingPID(err)
				return crvAPR, cvxAPR, crvAPY, cvxAPY
			}
			// Curve ZUSD-FRAXBP Factory | userVaultAddress.Hex() | 0xAB919492C39263e97D4190b9bb0d2437D15EAcC9
			// Curve XAI-FRAXBP Factory | userVaultAddress.Hex() | 0x9628d5f254B010C0122440F282E1fB6b0f25239c
			userVaultContract, err := contracts.NewConvexUserVault(userVaultAddress, client)
			if err != nil {
				if os.Getenv("ENVIRONMENT") == "dev" {
					logs.Error(`Unable to init userVault contract `+userVaultAddress.Hex(), err)
				}
				warnMissingPID(err)
				return crvAPR, cvxAPR, crvAPY, cvxAPY
			}
			stakingTokenAddress, err := userVaultContract.StakingToken(nil)
			if err != nil {
				if os.Getenv("ENVIRONMENT") == "dev" {
					logs.Error(`Unable to get stakingToken for userVault `+userVaultAddress.Hex(), err)
				}
				warnMissingPID(err)
				return crvAPR, cvxAPR, crvAPY, cvxAPY
			}
			// Curve ZUSD-FRAXBP Factory | stakingTokenAddress.Hex() | 0xFD2d7847E0f450d8B00d3D697D720C687E622a7B
			// Curve XAI-FRAXBP Factory | stakingTokenAddress.Hex() | 0x19f0a60f4635d3E2c48647822Eda5332BA094fd3
			stakingTokenContract, err := contracts.NewConvexStakingToken(stakingTokenAddress, client)
			if err != nil {
				if os.Getenv("ENVIRONMENT") == "dev" {
					logs.Error(`Unable to init stakingToken contract `+stakingTokenAddress.Hex(), err)
				}
				warnMissingPID(err)
				return crvAPR, cvxAPR, crvAPY, cvxAPY
			}
			rewardPID, err = stakingTokenContract.ConvexPoolId(nil)
			if err != nil {
				if os.Getenv("ENVIRONMENT") == "dev" {
					logs.Error(`Unable to get cvxPoolId for stakingToken `+stakingTokenAddress.Hex(), err)
				}
				warnMissingPID(err)
				return crvAPR, cvxAPR, crvAPY, cvxAPY
			}
			// Curve ZUSD-FRAXBP Factory | rewardPID | 196
			// Curve XAI-FRAXBP Factory | rewardPID | 129
		}
	}
	// Curve ZUSD-FRAXBP Factory | rewardPID | 196
	// Curve XAI-FRAXBP Factory | rewardPID | 129

	/**********************************************************************************************
	** Once we got the PID, we can query the convexBooster contract to get the `poolInfo` for this
	** and retrieve the `crvRewards` contract
	***********************************************************************************************/
	cvxBoosterContract, err := contracts.NewCVXBooster(storage.CVX_BOOSTER_ADDRESS[chainID], client)
	if err != nil {
		return crvAPR, cvxAPR, crvAPY, cvxAPY
	}
	poolInfo, err := cvxBoosterContract.PoolInfo(nil, rewardPID)
	if err != nil {
		return crvAPR, cvxAPR, crvAPY, cvxAPY
	}
	// Curve ZUSD-FRAXBP Factory | poolInfo.CrvRewards.Hex() | 0xFd3A7636694259b32B3896f59436997AD25380cA
	// Curve XAI-FRAXBP Factory | poolInfo.CrvRewards.Hex() | 0x4a866fE20A442Dff55FAA010684A5C1379151458

	/**********************************************************************************************
	** Once we got the poolInfo, we can init a new contract connector, which would be a
	** `BaseRewardPool`, and we should be able to query the rewardRate for it.
	***********************************************************************************************/
	rewardContract, err := contracts.NewCrvRewards(poolInfo.CrvRewards, client)
	if err != nil {
		return crvAPR, cvxAPR, crvAPY, cvxAPY
	}
	rateResult, err1 := rewardContract.RewardRate(nil)
	supplyResult, err2 := rewardContract.TotalSupply(nil)
	if err1 != nil || err2 != nil {
		return crvAPR, cvxAPR, crvAPY, cvxAPY
	}
	// Curve ZUSD-FRAXBP Factory | rateResult | 45048708159937
	// Curve XAI-FRAXBP Factory | rateResult | 36068329259793
	// Curve ZUSD-FRAXBP Factory | supplyResult | 317120356635973553634
	// Curve XAI-FRAXBP Factory | supplyResult | 822828886434854500906

	/**********************************************************************************************
	** Then we should be able to calculate the cvxAPR just like it's done on the CVX subgraph
	***********************************************************************************************/
	rate := helpers.ToNormalizedAmount(bigNumber.NewInt(0).Set(rateResult), 18)
	supply := helpers.ToNormalizedAmount(bigNumber.NewInt(0).Set(supplyResult), 18)
	// Curve ZUSD-FRAXBP Factory | rate | 4.5048708159937e-05
	// Curve XAI-FRAXBP Factory | rate | 3.6068329259793e-05
	// Curve ZUSD-FRAXBP Factory | supply | 317.12035663597356
	// Curve XAI-FRAXBP Factory | supply | 822.8288864348544
	crvPerUnderlying := bigNumber.NewFloat(0)
	virtualSupply := bigNumber.NewFloat(0).Mul(supply, virtualPoolPrice)
	// Curve ZUSD-FRAXBP Factory | virtualSupply | 5.820789461568759
	// Curve XAI-FRAXBP Factory | virtualSupply | 52.80806562935214

	if virtualSupply.Gt(storage.ZERO) {
		crvPerUnderlying = bigNumber.NewFloat(0).Div(rate, virtualSupply)
	}
	// Curve ZUSD-FRAXBP Factory | crvPerUnderlying | 7.739278057962252e-06
	// Curve XAI-FRAXBP Factory | crvPerUnderlying | 6.830079615668645e-07
	crvPerUnderlyingPerYear := bigNumber.NewFloat(0).Mul(crvPerUnderlying, bigNumber.NewFloat(31536000))
	// Curve ZUSD-FRAXBP Factory | crvPerUnderlyingPerYear | 244.06587283589755
	// Curve XAI-FRAXBP Factory | crvPerUnderlyingPerYear | 21.539339075972638
	cvxPerYear := getCVXForCRV(chainID, crvPerUnderlyingPerYear)
	// Curve ZUSD-FRAXBP Factory | cvxPerYear | 0.1322365676305354
	// Curve XAI-FRAXBP Factory | cvxPerYear | 0.011670161974476423

	crvPrice := bigNumber.NewFloat(0)
	if tokenPrice, ok := storage.GetPrice(chainID, storage.CRV_TOKEN_ADDRESS[chainID]); ok {
		crvPrice = tokenPrice.HumanizedPrice
	}
	// Curve ZUSD-FRAXBP Factory | crvPrice | 0.433265
	// Curve XAI-FRAXBP Factory | crvPrice | 0.433265
	cvxPrice := bigNumber.NewFloat(0)
	if tokenPrice, ok := storage.GetPrice(chainID, storage.CVX_TOKEN_ADDRESS[chainID]); ok {
		cvxPrice = tokenPrice.HumanizedPrice
	}
	// Curve ZUSD-FRAXBP Factory | cvxPrice | 2.019843
	// Curve XAI-FRAXBP Factory | cvxPrice | 2.019843
	crvAPR = bigNumber.NewFloat(0).Mul(crvPerUnderlyingPerYear, crvPrice)
	// Curve ZUSD-FRAXBP Factory | crvAPR | 105.74520039424516
	// Curve XAI-FRAXBP Factory | crvAPR | 9.332241744751284
	cvxAPR = bigNumber.NewFloat(0).Mul(cvxPerYear, cvxPrice)
	// Curve ZUSD-FRAXBP Factory | cvxAPR | 0.2670971054725635
	// Curve XAI-FRAXBP Factory | cvxAPR | 0.02357189497301238

	crvAPRFloat64, _ := crvAPR.Float64()
	cvxAPRFloat64, _ := cvxAPR.Float64()
	// Curve ZUSD-FRAXBP Factory | crvAPRFloat64 | 105.74520039424516
	// Curve XAI-FRAXBP Factory | crvAPRFloat64 | 9.332241744751284
	// Curve ZUSD-FRAXBP Factory | cvxAPRFloat64 | 0.2670971054725635
	// Curve XAI-FRAXBP Factory | cvxAPRFloat64 | 0.02357189497301238
	crvAPY = bigNumber.NewFloat(0).Add(bigNumber.NewFloat(0), crvAPR)
	// Curve ZUSD-FRAXBP Factory | crvAPY | 105.74520039424516
	// Curve XAI-FRAXBP Factory | crvAPY | 9.332241744751284
	cvxAPY = bigNumber.NewFloat(0).Add(bigNumber.NewFloat(0), cvxAPR)
	// Curve ZUSD-FRAXBP Factory | cvxAPY | 0.2670971054725635
	// Curve XAI-FRAXBP Factory | cvxAPY | 0.02357189497301238

	// Curve ZUSD-FRAXBP Factory | crvAPR | 105.74520039424516
	// Curve XAI-FRAXBP Factory | crvAPR | 9.332241744751284
	// Curve ZUSD-FRAXBP Factory | cvxAPR | 0.2670971054725635
	// Curve XAI-FRAXBP Factory | cvxAPR | 0.02357189497301238
	// Curve ZUSD-FRAXBP Factory | crvAPY | 105.74520039424516
	// Curve XAI-FRAXBP Factory | crvAPY | 9.332241744751284
	// Curve ZUSD-FRAXBP Factory | cvxAPY | 0.2670971054725635
	// Curve XAI-FRAXBP Factory | cvxAPY | 0.02357189497301238

	return crvAPR, cvxAPR, crvAPY, cvxAPY
}

/**************************************************************************************************
** Determine the keepCRV value for the vault. This indicates the amount of CRV this strategy should
** keep as rewards instead of insta dump.
** Because of the contract upgrade, we have multiple stuff to checks:
** - If the strategy does not have a `UselLocalCRV` (with a typo) function, we return the already
**   retrieved KeepCRV value (retrieved for normal yDaemon execution)
** - If the strategy has a `UselLocalCRV` function, we check if it's true or false
**   - If it's true, we retrieve the localCRV value and use it if it exists
**     - If it does not, we can check the LocalKeepCRV value and use it if it exists, or ZERO
**   - If it's false, we can query the KeepCRV value from the curveGlobal contract and use it
**************************************************************************************************/
func determineConvexKeepCRV(strategy models.TStrategy) *bigNumber.Float {
	if strategy.KeepCRV == nil {
		return storage.ZERO
	}
	client := ethereum.GetRPC(strategy.ChainID)
	convexStrategyContract, _ := contracts.NewConvexBaseStrategy(strategy.Address, client)
	useLocalCRV, err := convexStrategyContract.UselLocalCRV(nil)
	if err != nil {
		return helpers.ToNormalizedAmount(strategy.KeepCRV, 4)
	}
	if useLocalCRV {
		cvxKeepCRV, err := convexStrategyContract.LocalCRV(nil)
		if err != nil {
			localKeepCRV, err := convexStrategyContract.LocalKeepCRV(nil)
			if err != nil {
				return storage.ZERO
			}
			return helpers.ToNormalizedAmount(bigNumber.NewInt(0).Set(localKeepCRV), 4)
		}
		return helpers.ToNormalizedAmount(bigNumber.NewInt(0).Set(cvxKeepCRV), 4)
	}
	curveGlobal, err := convexStrategyContract.CurveGlobal(nil)
	if err != nil {
		return storage.ZERO
	}
	curveGlobalContract, err := contracts.NewStrategyBase(curveGlobal, client)
	if err != nil {
		return storage.ZERO
	}
	keepCRV, err := curveGlobalContract.KeepCRV(nil)
	if err != nil {
		return storage.ZERO
	}
	return helpers.ToNormalizedAmount(bigNumber.NewInt(0).Set(keepCRV), 4)
}

/**************************************************************************************************
** Check if the strategy is a convex strategy. This is a check based on the strategy name. What
** could go wrong.
**************************************************************************************************/
func isConvexStrategy(strategy models.TStrategy) bool {
	name := strings.ToLower(strategy.Name)
	return strings.Contains(name, `convex`) && !strings.Contains(name, `convexfrax`) && !strings.Contains(name, `ethconvex`)
}
```
