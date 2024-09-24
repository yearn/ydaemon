package apr

import (
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
	"github.com/yearn/ydaemon/internal/multicalls"
	"github.com/yearn/ydaemon/internal/storage"
)

var VELO_STAKING_POOLS_REGISTRY = common.HexToAddress(`0x41c914ee0c7e1a5edcd0295623e6dc557b5abf3c`)

/**************************************************************************************************
** Check if the vault is a velodrom vault. In order to check this, we need to check if the token
** has a staking pool in velodrome
**************************************************************************************************/
func isVeloVault(chainID uint64, vault models.TVault) (common.Address, bool) {
	if chainID != 10 {
		return common.Address{}, false
	}

	veloVoter, err := contracts.NewYVelodromeVoterRegistry(VELO_STAKING_POOLS_REGISTRY, ethereum.GetRPC(chainID))
	if err != nil {
		logs.Error(err)
		return common.Address{}, false
	}
	gaugeAddressForVoter, err := veloVoter.Gauges(nil, vault.AssetAddress)
	if err != nil {
		logs.Error(err)
		return common.Address{}, false
	}

	return gaugeAddressForVoter, gaugeAddressForVoter != common.Address{}
}

func calculateVeloLikeStrategyAPY(
	vault models.TVault,
	strategy models.TStrategy,
	veloStakingPoolAddress common.Address,
) TStrategyAPY {
	/**********************************************************************************************
	** First we will need a few things from the staking contract. We will use a multicall to
	** retrieve the following:
	** - The periodFinish
	** - The rewardRate
	** - The totalSupply
	**********************************************************************************************/
	calls := []ethereum.Call{}
	calls = append(calls, multicalls.GetPeriodFinish(veloStakingPoolAddress.Hex(), veloStakingPoolAddress))
	calls = append(calls, multicalls.GetRewardRate(veloStakingPoolAddress.Hex(), veloStakingPoolAddress))
	calls = append(calls, multicalls.GetTotalSupply(veloStakingPoolAddress.Hex(), veloStakingPoolAddress))
	calls = append(calls, multicalls.GetRewardToken(veloStakingPoolAddress.Hex(), veloStakingPoolAddress))
	calls = append(calls, multicalls.GetDecimals(veloStakingPoolAddress.Hex(), veloStakingPoolAddress))
	calls = append(calls, multicalls.GetStategyLocalKeepVelo(strategy.Address.Hex(), strategy.Address))
	response := multicalls.Perform(vault.ChainID, calls, nil)
	periodFinish := helpers.DecodeBigInt(response[veloStakingPoolAddress.Hex()+`periodFinish`])
	rewardRateRaw := helpers.DecodeBigInt(response[veloStakingPoolAddress.Hex()+`rewardRate`])
	totalSupplyRaw := helpers.DecodeBigInt(response[veloStakingPoolAddress.Hex()+`totalSupply`])
	rewardTokenRaw := helpers.DecodeAddress(response[veloStakingPoolAddress.Hex()+`rewardToken`])
	localKeepVeloRaw := helpers.DecodeBigInt(response[strategy.Address.Hex()+`localKeepVELO`])
	localKeepVelo := helpers.ToNormalizedAmount(localKeepVeloRaw, 4)

	/**********************************************************************************************
	** If periodFinish is before now, aka rewards are over, we can stop here
	**********************************************************************************************/
	now := time.Now().Unix()
	if periodFinish.Int64() < now {
		return TStrategyAPY{
			Type: `v2:velo_unpopular`,
			Composite: TCompositeData{
				KeepVelo: localKeepVelo,
			},
		}
	}

	/**********************************************************************************************
	** If the total supply is 0, we can stop here, aka nothing is staked, so no rewards
	**********************************************************************************************/
	if totalSupplyRaw.IsZero() {
		return TStrategyAPY{
			Type: `v2:velo_unpopular`,
			Composite: TCompositeData{
				KeepVelo: localKeepVelo,
			},
		}
	}

	/**********************************************************************************************
	** We need to retrieve a bunch to be able to proceed. They are already in the vault object
	** - the performanceFee for that vault
	** - the managementFee for that vault
	**********************************************************************************************/
	debtRatio := helpers.ToNormalizedAmount(strategy.LastDebtRatio, 4)
	vaultPerformanceFee := helpers.ToNormalizedAmount(bigNumber.NewInt(int64(vault.PerformanceFee)), 4)
	vaultManagementFee := helpers.ToNormalizedAmount(bigNumber.NewInt(int64(vault.ManagementFee)), 4)
	oneMinusKeepVelo := bigNumber.NewFloat(0).Sub(bigNumber.NewFloat(1), localKeepVelo)
	oneMinusPerfFee := bigNumber.NewFloat(0).Sub(bigNumber.NewFloat(1), vaultPerformanceFee)
	rewardRate := helpers.ToNormalizedAmount(rewardRateRaw, 18)
	totalSupply := helpers.ToNormalizedAmount(totalSupplyRaw, 18)
	secondsPerYear := bigNumber.NewFloat(31_556_952)

	/**********************************************************************************************
	** If the reward rate is 0, we can stop here, aka no rewards
	**********************************************************************************************/
	if rewardRate.IsZero() || oneMinusKeepVelo.IsZero() {
		return TStrategyAPY{
			Type: `v2:velo_unpopular`,
			Composite: TCompositeData{
				KeepVelo: localKeepVelo,
			},
		}
	}

	/**********************************************************************************************
	** If that's good, we will need the price of the vault token and the price of the rewards token
	** to compute the APR.
	**********************************************************************************************/
	poolPrice := bigNumber.NewFloat(0)
	if tokenPrice, ok := storage.GetPrice(vault.ChainID, vault.AssetAddress); ok {
		poolPrice = tokenPrice.HumanizedPrice
	}

	rewardsPrice := bigNumber.NewFloat(0)
	if tokenPrice, ok := storage.GetPrice(vault.ChainID, rewardTokenRaw); ok {
		rewardsPrice = tokenPrice.HumanizedPrice
	}

	/**********************************************************************************************
	** And now we can compute the APR
	**********************************************************************************************/
	rewardRate = bigNumber.NewFloat(0).Mul(rewardRate, oneMinusKeepVelo) // rewardRate = rewardRate * (1 - keep)
	grossAPRTop := bigNumber.NewFloat(0).Mul(rewardRate, rewardsPrice)   // rewardRate * token_price
	grossAPRTop = bigNumber.NewFloat(0).Mul(grossAPRTop, secondsPerYear) // rewardRate * token_price * SECONDS_PER_YEAR
	grossAPRBottom := bigNumber.NewFloat(0).Mul(poolPrice, totalSupply)  // pool_price * totalSupply
	grossAPR := bigNumber.NewFloat(0).Div(grossAPRTop, grossAPRBottom)   // (rewardRate * token_price * SECONDS_PER_YEAR) / (pool_price * totalSupply)

	/**********************************************************************************************
	** Calculate the strategy Net APR:
	** Take the gross APR and remove the performance fee and the management fee
	**********************************************************************************************/
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

	apyStruct := TStrategyAPY{
		Type:      "v2:velo",
		DebtRatio: debtRatio,
		NetAPY:    bigNumber.NewFloat(0).Mul(netAPY, debtRatio),
		Composite: TCompositeData{
			KeepVelo: localKeepVelo,
		},
	}
	return apyStruct
}

/**************************************************************************************************
** If the vault is a velo vault or a fork of it, we can calculate the forward APR  using always the
** same base formula
**************************************************************************************************/
func computeVeloLikeForwardAPY(
	vault models.TVault,
	allStrategiesForVault map[string]models.TStrategy,
	veloStakingPoolAddress common.Address,
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
	for _, strategy := range allStrategiesForVault {
		if strategy.LastDebtRatio == nil || strategy.LastDebtRatio.IsZero() {
			if os.Getenv("ENVIRONMENT") == "dev" {
				logs.Info("Skipping strategy " + strategy.Address.Hex() + " for vault " + vault.Address.Hex() + " because debt ratio is zero")
			}
			continue
		}

		strategyAPY := calculateVeloLikeStrategyAPY(vault, strategy, veloStakingPoolAddress)
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
