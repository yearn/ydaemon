package apr

import (
	"time"

	"github.com/yearn/ydaemon/common/bigNumber"
	"github.com/yearn/ydaemon/common/ethereum"
	"github.com/yearn/ydaemon/common/helpers"
	"github.com/yearn/ydaemon/internal/models"
	"github.com/yearn/ydaemon/internal/multicalls"
	"github.com/yearn/ydaemon/internal/storage"
)

func computeJuicedStakingRewardsAPY(chainID uint64, vault models.TVault) (*bigNumber.Float, *bigNumber.Float, bool) {
	/**********************************************************************************************
	** First thing to do is to check if the vault has a staking contract associated with it.
	** We can retrieve that from the store.
	**********************************************************************************************/
	stakingElement, ok := storage.GetJuicedStakingDataForVault(chainID, vault.Address)
	if !ok {
		return bigNumber.NewFloat(0), bigNumber.NewFloat(0), false
	}

	cumulatedStakingRewardAPR := bigNumber.NewFloat(0)
	cumulatedStakingRewardAPY := bigNumber.NewFloat(0)

	/**********************************************************************************************
	** We will need the price of the vault token to compute the APR
	**********************************************************************************************/
	vaultPrice := bigNumber.NewFloat(0)
	if tokenPrice, ok := storage.GetPrice(vault.ChainID, vault.Address); ok {
		vaultPrice = tokenPrice.HumanizedPrice
	}

	/**********************************************************************************************
	** Once we got it we will need a few things from the staking contract. We will use a multicall
	** to retrieve the following:
	** - The totalSupply
	** - The vault decimals
	**********************************************************************************************/
	calls := []ethereum.Call{}
	calls = append(calls, multicalls.GetTotalSupply(stakingElement.StakingAddress.Hex(), stakingElement.StakingAddress))
	calls = append(calls, multicalls.GetDecimals(vault.Address.Hex(), vault.Address))
	response := multicalls.Perform(chainID, calls, nil)
	rawTotalSupply := helpers.DecodeBigInt(response[stakingElement.StakingAddress.Hex()+`totalSupply`])
	rawVaultDecimals := helpers.DecodeUint64(response[vault.Address.Hex()+`decimals`])

	/**********************************************************************************************
	** Once it's done, we can loop through the reward tokens and compute the APR for each of them
	** and cumulate them.
	**********************************************************************************************/
	for _, rewardToken := range stakingElement.RewardTokens {
		rewardRate := rewardToken.Rate
		rewardDuration := rewardToken.Duration
		rewardPeriodFinish := rewardToken.PeriodFinish
		rewardDecimals := rewardToken.Decimals

		/**********************************************************************************************
		** If periodFinish is before now, aka rewards are over, we can stop here
		**********************************************************************************************/
		now := time.Now().Unix()
		if rewardPeriodFinish < uint64(now) {
			storage.AssignJuicedStakingRewardAPY(chainID, vault.Address, rewardToken.Address, bigNumber.NewFloat(0))
			continue
		}

		rewardPerWeek := bigNumber.NewFloat(0).Mul(
			bigNumber.NewFloat(0).SetUint64(rewardRate),
			bigNumber.NewFloat(0).SetUint64(rewardDuration),
		)
		normalizedTotalSupply := helpers.ToNormalizedAmount(rawTotalSupply, rawVaultDecimals)

		/**********************************************************************************************
		** Retrieve the price of the reward token from our storage
		**********************************************************************************************/
		rewardsPrice := bigNumber.NewFloat(0)
		if tokenPrice, ok := storage.GetPrice(vault.ChainID, rewardToken.Address); ok {
			rewardsPrice = tokenPrice.HumanizedPrice
		}
		secondsPerYear := bigNumber.NewFloat(31_556_952)

		/**********************************************************************************************
		** Start computing the APR. The formula is quite simple:
		** (rewardPerWeek * rewardsPrice) / (vaultPrice * normalizedTotalSupply) / decimals
		** However, we want to scale it to a yearly reward. This means we will need to reduce the APR
		** to a second reward and then multiply it by the number of seconds in a year.
		**********************************************************************************************/
		decimalsScale := bigNumber.NewFloat(0).SetInt(bigNumber.NewInt(10).Exp(bigNumber.NewInt(10), bigNumber.NewInt(int64(rewardDecimals)), nil))
		stakingRewardAPR := bigNumber.NewFloat(0).Mul(rewardPerWeek, rewardsPrice)          // (rewardPerWeek * rewardsPrice)
		vaultPricePerSupply := bigNumber.NewFloat(0).Mul(vaultPrice, normalizedTotalSupply) // (vaultPrice * normalizedTotalSupply)
		stakingRewardAPR = bigNumber.NewFloat(0).Div(stakingRewardAPR, vaultPricePerSupply) // (rewardPerWeek * rewardsPrice) / (vaultPrice * normalizedTotalSupply)
		stakingRewardAPR = bigNumber.NewFloat(0).Div(stakingRewardAPR, decimalsScale)       // (rewardPerWeek * rewardsPrice) / (vaultPrice * normalizedTotalSupply) / decimals

		// The stakingRewardAPR at this point is based on a RewardDuration. We want to annualize it based on secondsPerYear
		rewardDurationAsFloat := bigNumber.NewFloat(0).SetUint64(rewardDuration)
		stakingRewardAPR = bigNumber.NewFloat(0).Div(stakingRewardAPR, rewardDurationAsFloat)
		stakingRewardAPR = bigNumber.NewFloat(0).Mul(stakingRewardAPR, secondsPerYear)
		stakingRewardAPY := convertAPRToAPY(stakingRewardAPR, bigNumber.NewFloat(365.0/15.0))

		cumulatedStakingRewardAPR = bigNumber.NewFloat(0).Add(cumulatedStakingRewardAPR, stakingRewardAPR)
		cumulatedStakingRewardAPY = bigNumber.NewFloat(0).Add(cumulatedStakingRewardAPY, stakingRewardAPY)
		storage.AssignJuicedStakingRewardAPY(chainID, vault.Address, rewardToken.Address, stakingRewardAPY)
	}

	/**********************************************************************************************
	** Finally, we can compute the APR
	**********************************************************************************************/
	return cumulatedStakingRewardAPR, cumulatedStakingRewardAPY, true
}
