package apy

import (
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
)

var AERO_STAKING_POOLS_REGISTRY = common.HexToAddress(`0x16613524e02ad97eDfeF371bC883F2F5d6C480A5`)

/**************************************************************************************************
** Check if the vault is a aerodrome vault. In order to check this, we need to check if the token
** has a staking pool in aerodrome
**************************************************************************************************/
func isAeroVault(chainID uint64, vault models.TVault) (common.Address, bool) {
	if chainID != 8453 {
		return common.Address{}, false
	}

	aeroVoter, err := contracts.NewAerodromeVoterRegistryCaller(AERO_STAKING_POOLS_REGISTRY, ethereum.GetRPC(chainID))
	if err != nil {
		logs.Error(err)
		return common.Address{}, false
	}
	gaugeAddressForVoter, err := aeroVoter.Gauges(nil, vault.Token.Address)
	if err != nil {
		logs.Error(err)
		return common.Address{}, false
	}

	return gaugeAddressForVoter, gaugeAddressForVoter != common.Address{}
}

func calculateAeroLikeStrategyAPR(
	vault models.TVault,
	strategy *models.TStrategy,
	aeroStakingPoolAddress common.Address,
) TStrategyAPR {
	/**********************************************************************************************
	** First we will need a few things from the staking contract. We will use a multicall to
	** retrieve the following:
	** - The periodFinish
	** - The rewardRate
	** - The totalSupply
	**********************************************************************************************/
	calls := []ethereum.Call{}
	calls = append(calls, multicalls.GetPeriodFinish(aeroStakingPoolAddress.Hex(), aeroStakingPoolAddress))
	calls = append(calls, multicalls.GetRewardRate(aeroStakingPoolAddress.Hex(), aeroStakingPoolAddress))
	calls = append(calls, multicalls.GetTotalSupply(aeroStakingPoolAddress.Hex(), aeroStakingPoolAddress))
	calls = append(calls, multicalls.GetRewardToken(aeroStakingPoolAddress.Hex(), aeroStakingPoolAddress))
	calls = append(calls, multicalls.GetDecimals(aeroStakingPoolAddress.Hex(), aeroStakingPoolAddress))
	calls = append(calls, multicalls.GetStategyLocalKeepVelo(strategy.Address.Hex(), strategy.Address))
	response := multicalls.Perform(vault.ChainID, calls, nil)
	periodFinish := helpers.DecodeBigInt(response[aeroStakingPoolAddress.Hex()+`periodFinish`])
	rewardRateRaw := helpers.DecodeBigInt(response[aeroStakingPoolAddress.Hex()+`rewardRate`])
	totalSupplyRaw := helpers.DecodeBigInt(response[aeroStakingPoolAddress.Hex()+`totalSupply`])
	rewardTokenRaw := helpers.DecodeAddress(response[aeroStakingPoolAddress.Hex()+`rewardToken`])
	localKeepAeroRaw := helpers.DecodeBigInt(response[strategy.Address.Hex()+`localKeepAERO`])

	/**********************************************************************************************
	** If periodFinish is before now, aka rewards are over, we can stop here
	**********************************************************************************************/
	now := time.Now().Unix()
	if periodFinish.Int64() < now {
		return TStrategyAPR{
			Type: `v2:aero_unpopular`,
		}
	}

	/**********************************************************************************************
	** If the total supply is 0, we can stop here, aka nothing is staked, so no rewards
	**********************************************************************************************/
	if totalSupplyRaw.IsZero() {
		return TStrategyAPR{
			Type: `v2:aero_unpopular`,
		}
	}

	/**********************************************************************************************
	** We need to retrieve a bunch to be able to proceed. They are already in the vault object
	** - the performanceFee for that vault
	** - the managementFee for that vault
	**********************************************************************************************/
	debtRatio := helpers.ToNormalizedAmount(strategy.DebtRatio, 4)
	vaultPerformanceFee := helpers.ToNormalizedAmount(bigNumber.NewInt(int64(vault.PerformanceFee)), 4)
	vaultManagementFee := helpers.ToNormalizedAmount(bigNumber.NewInt(int64(vault.ManagementFee)), 4)
	localKeepAero := helpers.ToNormalizedAmount(localKeepAeroRaw, 4)
	oneMinusKeepAero := bigNumber.NewFloat(0).Sub(bigNumber.NewFloat(1), localKeepAero)
	oneMinusPerfFee := bigNumber.NewFloat(0).Sub(bigNumber.NewFloat(1), vaultPerformanceFee)
	rewardRate := helpers.ToNormalizedAmount(rewardRateRaw, 18)
	totalSupply := helpers.ToNormalizedAmount(totalSupplyRaw, 18)
	secondsPerYear := bigNumber.NewFloat(31_556_952)

	/**********************************************************************************************
	** If the reward rate is 0, we can stop here, aka no rewards
	**********************************************************************************************/
	if rewardRate.IsZero() || oneMinusKeepAero.IsZero() {
		return TStrategyAPR{
			Type: `v2:aero_unpopular`,
		}
	}

	/**********************************************************************************************
	** If that's good, we will need the price of the vault token and the price of the rewards token
	** to compute the APR.
	**********************************************************************************************/
	poolPrice := getTokenPrice(vault.ChainID, vault.Token.Address)
	rewardsPrice := getTokenPrice(vault.ChainID, rewardTokenRaw)

	/**********************************************************************************************
	** And now we can compute the APR
	**********************************************************************************************/
	rewardRate = bigNumber.NewFloat(0).Mul(rewardRate, oneMinusKeepAero) // rewardRate = rewardRate * (1 - keep)
	grossAPRTop := bigNumber.NewFloat(0).Mul(rewardRate, rewardsPrice)   // rewardRate * token_price
	grossAPRTop = bigNumber.NewFloat(0).Mul(grossAPRTop, secondsPerYear) // rewardRate * token_price * SECONDS_PER_YEAR
	grossAPRBottom := bigNumber.NewFloat(0).Mul(poolPrice, totalSupply)  // pool_price * totalSupply
	grossAPR := bigNumber.NewFloat(0).Div(grossAPRTop, grossAPRBottom)   // (rewardRate * token_price * SECONDS_PER_YEAR) / (pool_price * totalSupply)

	/**********************************************************************************************
	** Calculate the strategy Net APR:
	** Take the gross APR and remove the performance fee and the management fee
	**********************************************************************************************/
	netAPR := bigNumber.NewFloat(0).Mul(grossAPR, oneMinusPerfFee) // grossAPR * (1 - perfFee)
	netAPR = bigNumber.NewFloat(0).Sub(netAPR, vaultManagementFee) // (grossAPR * (1 - perfFee)) - managementFee

	/**********************************************************************************************
	** Calculate the strategy Net APY:
	** Take the net APR and compound it
	**********************************************************************************************/
	netAPY := bigNumber.NewFloat(0).Div(netAPR, bigNumber.NewFloat(365)) // netAPR / 365
	netAPY = bigNumber.NewFloat(0).Add(netAPY, bigNumber.NewFloat(1))    // 1 + (netAPR / 365)
	netAPY = bigNumber.NewFloat(0).Pow(netAPY, 365)                      // (1 + (netAPR / 365)) ^ 365
	netAPY = bigNumber.NewFloat(0).Sub(netAPY, bigNumber.NewFloat(1))    // ((1 + (netAPR / 365)) ^ 365) - 1

	apyStruct := TStrategyAPR{
		Type:      "v2:aero",
		DebtRatio: debtRatio,
		GrossAPR:  bigNumber.NewFloat(0).Mul(grossAPR, debtRatio),
		NetAPY:    bigNumber.NewFloat(0).Mul(netAPR, debtRatio),
	}
	return apyStruct
}

/**************************************************************************************************
** If the vault is a aero vault or a fork of it, we can calculate the forward APR  using always the
** same base formula
**************************************************************************************************/
func computeAeroLikeForwardAPR(
	vault models.TVault,
	allStrategiesForVault []*models.TStrategy,
	aeroStakingPoolAddress common.Address,
) TForwardAPR {
	TypeOf := ``
	GrossAPR := bigNumber.NewFloat(0)
	NetAPY := bigNumber.NewFloat(0)
	Boost := bigNumber.NewFloat(0)
	PoolAPY := bigNumber.NewFloat(0)
	BoostedAPR := bigNumber.NewFloat(0)
	BaseAPR := bigNumber.NewFloat(0)
	CvxAPR := bigNumber.NewFloat(0)
	RewardsAPR := bigNumber.NewFloat(0)
	for _, strategy := range allStrategiesForVault {
		if strategy.DebtRatio == nil || strategy.DebtRatio.IsZero() {
			logs.Info("Skipping strategy " + strategy.Address.Hex() + " for vault " + vault.Address.Hex() + " because debt ratio is zero")
			continue
		}
		if strategy.TotalDebt == nil || strategy.TotalDebt.IsZero() {
			logs.Info("Skipping strategy " + strategy.Address.Hex() + " for vault " + vault.Address.Hex() + " because total debt is zero")
			continue
		}

		strategyAPR := calculateAeroLikeStrategyAPR(vault, strategy, aeroStakingPoolAddress)
		TypeOf += strings.TrimSpace(` ` + strategyAPR.Type)
		GrossAPR = bigNumber.NewFloat(0).Add(GrossAPR, strategyAPR.GrossAPR)
		NetAPY = bigNumber.NewFloat(0).Add(NetAPY, strategyAPR.NetAPY)
		Boost = bigNumber.NewFloat(0).Add(Boost, strategyAPR.Composite.Boost)
		PoolAPY = bigNumber.NewFloat(0).Add(PoolAPY, strategyAPR.Composite.PoolAPY)
		BoostedAPR = bigNumber.NewFloat(0).Add(BoostedAPR, strategyAPR.Composite.BoostedAPR)
		BaseAPR = bigNumber.NewFloat(0).Add(BaseAPR, strategyAPR.Composite.BaseAPR)
		CvxAPR = bigNumber.NewFloat(0).Add(CvxAPR, strategyAPR.Composite.CvxAPR)
		RewardsAPR = bigNumber.NewFloat(0).Add(RewardsAPR, strategyAPR.Composite.RewardsAPR)
	}

	return TForwardAPR{
		Type:     strings.TrimSpace(TypeOf),
		GrossAPR: GrossAPR,
		NetAPY:   NetAPY,
		Composite: TAPIV1Composite{
			Boost:      Boost,
			PoolAPY:    PoolAPY,
			BoostedAPR: BoostedAPR,
			BaseAPR:    BaseAPR,
			CvxAPR:     CvxAPR,
			RewardsAPR: RewardsAPR,
		},
	}
}
