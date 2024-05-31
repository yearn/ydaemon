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
func getConvexRewardAPR(
	chainID uint64,
	strategy models.TStrategy,
	baseAssetPrice *bigNumber.Float,
	poolPrice *bigNumber.Float,
) *bigNumber.Float {
	client := ethereum.GetRPC(chainID)
	convexStrategyContract, _ := contracts.NewConvexBaseStrategy(strategy.Address, client)
	cvxBoosterContract, _ := contracts.NewCVXBooster(storage.CVX_BOOSTER_ADDRESS[chainID], client)
	rewardPID, err := convexStrategyContract.Pid(nil)
	if err != nil {
		rewardPID, err = convexStrategyContract.ID(nil)
		if err != nil {
			rewardPID, err = convexStrategyContract.FraxPid(nil)
			if err != nil {
				if os.Getenv("ENVIRONMENT") == "dev" {
					logs.Error(`Unable to get reward PID for convex strategy ` + strategy.Address.Hex())
				}
				return storage.ZERO
			}
		}
	}
	rewardContract, err := cvxBoosterContract.PoolInfo(nil, rewardPID)
	if err != nil {
		logs.Error(err)
		return storage.ZERO
	}
	crvRewardContract, _ := contracts.NewCrvRewards(rewardContract.CrvRewards, client)
	rewardsLength, _ := crvRewardContract.ExtraRewardsLength(nil)

	now := time.Now().Unix()
	totalRewardsAPR := bigNumber.NewFloat(0)
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
	return totalRewardsAPR
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
func getCVXPoolAPR(
	chainID uint64,
	strategyAddress common.Address,
	virtualPoolPrice *bigNumber.Float,
) (*bigNumber.Float, *bigNumber.Float) {
	crvAPR := bigNumber.NewFloat(0)
	cvxAPR := bigNumber.NewFloat(0)

	/**********************************************************************************************
	** First thing to do to be able to calculate the APR is to retrieve the `crvRewards` contract
	** for this given convex strategy. It's a multiple step with way too many contracts involved,
	** but in the end we should be able to query the `rewardRate` for it.
	***********************************************************************************************/
	client := ethereum.GetRPC(chainID)
	convexStrategyContract, err := contracts.NewConvexBaseStrategy(strategyAddress, client)
	if err != nil {
		return crvAPR, cvxAPR
	}

	/**********************************************************************************************
	** We need to know the PID of the pool. Based on the contract, it can be PID or ID, so we need
	** to try both, and one call can revert.
	***********************************************************************************************/
	rewardPID, err := convexStrategyContract.Pid(nil)
	if err != nil {
		rewardPID, err = convexStrategyContract.ID(nil)
		if err != nil {
			rewardPID, err = convexStrategyContract.FraxPid(nil)
			if err != nil {
				if os.Getenv("ENVIRONMENT") == "dev" {
					logs.Error(`Unable to get reward PID for convex strategy ` + strategyAddress.Hex())
				}
				return crvAPR, cvxAPR
			}
		}
	}

	/**********************************************************************************************
	** Once we got the PID, we can query the convexBooster contract to get the `poolInfo` for this
	** and retrieve the `crvRewards` contract
	***********************************************************************************************/
	cvxBoosterContract, err := contracts.NewCVXBooster(storage.CVX_BOOSTER_ADDRESS[chainID], client)
	if err != nil {
		return crvAPR, cvxAPR
	}
	poolInfo, err := cvxBoosterContract.PoolInfo(nil, rewardPID)
	if err != nil {
		return crvAPR, cvxAPR
	}

	/**********************************************************************************************
	** Once we got the poolInfo, we can init a new contract connector, which would be a
	** `BaseRewardPool`, and we should be able to query the rewardRate for it.
	***********************************************************************************************/
	rewardContract, err := contracts.NewCrvRewards(poolInfo.CrvRewards, client)
	if err != nil {
		return crvAPR, cvxAPR
	}
	rateResult, err1 := rewardContract.RewardRate(nil)
	supplyResult, err2 := rewardContract.TotalSupply(nil)
	if err1 != nil || err2 != nil {
		return crvAPR, cvxAPR
	}

	/**********************************************************************************************
	** Then we should be able to calculate the cvxAPR just like it's done on the CVX subgraph
	***********************************************************************************************/
	rate := helpers.ToNormalizedAmount(bigNumber.NewInt(0).Set(rateResult), 18)
	supply := helpers.ToNormalizedAmount(bigNumber.NewInt(0).Set(supplyResult), 18)
	crvPerUnderlying := bigNumber.NewFloat(0)
	virtualSupply := bigNumber.NewFloat(0).Mul(supply, virtualPoolPrice)

	if virtualSupply.Gt(storage.ZERO) {
		crvPerUnderlying = bigNumber.NewFloat(0).Div(rate, virtualSupply)
	}
	crvPerYear := bigNumber.NewFloat(0).Mul(crvPerUnderlying, bigNumber.NewFloat(31536000))
	cvxPerYear := getCVXForCRV(chainID, crvPerYear)

	crvPrice := bigNumber.NewFloat(0)
	if tokenPrice, ok := storage.GetPrice(chainID, storage.CRV_TOKEN_ADDRESS[chainID]); ok {
		crvPrice = tokenPrice.HumanizedPrice
	}
	cvxPrice := bigNumber.NewFloat(0)
	if tokenPrice, ok := storage.GetPrice(chainID, storage.CVX_TOKEN_ADDRESS[chainID]); ok {
		cvxPrice = tokenPrice.HumanizedPrice
	}
	crvAPR = bigNumber.NewFloat(0).Mul(crvPerYear, crvPrice)
	cvxAPR = bigNumber.NewFloat(0).Mul(cvxPerYear, cvxPrice)
	crvAPR = bigNumber.NewFloat(0).Div(crvAPR, bigNumber.NewFloat(100))
	cvxAPR = bigNumber.NewFloat(0).Div(cvxAPR, bigNumber.NewFloat(100))

	return crvAPR, cvxAPR
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
