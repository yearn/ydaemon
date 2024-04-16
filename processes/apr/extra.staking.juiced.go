package apr

import (
	"math/big"
	"strconv"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/common/bigNumber"
	"github.com/yearn/ydaemon/common/ethereum"
	"github.com/yearn/ydaemon/common/helpers"
	"github.com/yearn/ydaemon/internal/models"
	"github.com/yearn/ydaemon/internal/multicalls"
	"github.com/yearn/ydaemon/internal/storage"
)

func computeJuicedStakingRewardsAPR(chainID uint64, vault models.TVault) (*bigNumber.Float, bool) {
	/**********************************************************************************************
	** First thing to do is to check if the vault has a staking contract associated with it.
	** We can retrieve that from the store.
	**********************************************************************************************/
	stakingContract, ok := storage.GetJuicedStakingForVault(chainID, vault.Address)
	if !ok {
		return bigNumber.NewFloat(0), false
	}

	cumulatedStakingRewardAPR := bigNumber.NewFloat(0)

	/**********************************************************************************************
	** We will need the price of the vault token to compute the APR
	**********************************************************************************************/
	vaultPrice := bigNumber.NewFloat(0)
	if tokenPrice, ok := storage.GetPrice(vault.ChainID, vault.Address); ok {
		vaultPrice = tokenPrice.HumanizedPrice
	}

	/**********************************************************************************************
	** We have no way to know how many reward tokens there are, so we will loop through the first
	** 10 reward tokens and check if they exist. If they do, we will compute the APR for them.
	** First step is then to get the rewards tokens for these indexes and remove the empty ones.
	** We will also need to get the totalSupply of the staking contract. As we need it only
	** once, it's more efficient to get it here.
	**********************************************************************************************/
	calls := []ethereum.Call{}
	rewardTokens := []common.Address{}
	calls = append(calls, multicalls.GetTotalSupply(stakingContract.Hex(), stakingContract))
	calls = append(calls, multicalls.GetDecimals(vault.Address.Hex(), vault.Address))
	for i := 0; i < 10; i++ {
		calls = append(calls, multicalls.GetRewardTokens(strconv.Itoa(i), stakingContract, int64(i)))
	}

	response := multicalls.Perform(chainID, calls, nil)
	rawTotalSupply := helpers.DecodeBigInt(response[stakingContract.Hex()+`totalSupply`])
	rawVaultDecimals := helpers.DecodeUint64(response[vault.Address.Hex()+`decimals`])
	for i := 0; i < 10; i++ {
		rewardToken := helpers.DecodeAddress(response[strconv.Itoa(i)+`rewardTokens`])
		if (rewardToken != common.Address{}) {
			rewardTokens = append(rewardTokens, rewardToken)
		}
	}

	/**********************************************************************************************
	** Once we have the indexes of the reward tokens, we can do another multicall to get, for each
	** of them, the reward data, the decimals of the reward token
	**********************************************************************************************/
	rewardTokensCalls := []ethereum.Call{}
	for _, rewardToken := range rewardTokens {
		rewardTokensCalls = append(rewardTokensCalls, multicalls.GetRewardData(rewardToken.Hex(), stakingContract, rewardToken))
		rewardTokensCalls = append(rewardTokensCalls, multicalls.GetDecimals(rewardToken.Hex(), rewardToken))
	}
	rewardTokensResponse := multicalls.Perform(chainID, rewardTokensCalls, nil)

	/**********************************************************************************************
	** Once it's done, we can loop through the reward tokens and compute the APR for each of them
	** and cumulate them.
	**********************************************************************************************/
	for _, rewardToken := range rewardTokens {
		rewardDataKey := rewardTokensResponse[rewardToken.Hex()+`rewardData`]
		rewardRate := bigNumber.NewInt(0)
		rewardDuration := bigNumber.NewInt(0)
		rewardPeriodFinish := bigNumber.NewInt(0)
		if len(rewardDataKey) > 0 {
			rewardDuration = bigNumber.SetInt(rewardDataKey[1].(*big.Int))
			rewardPeriodFinish = bigNumber.SetInt(rewardDataKey[2].(*big.Int))
			rewardRate = bigNumber.SetInt(rewardDataKey[3].(*big.Int))
		}
		rewardDecimals := helpers.DecodeUint64(rewardTokensResponse[rewardToken.Hex()+`decimals`])

		/**********************************************************************************************
		** If periodFinish is before now, aka rewards are over, we can stop here
		**********************************************************************************************/
		now := time.Now().Unix()
		if rewardPeriodFinish.Int64() < now {
			return bigNumber.NewFloat(0), false
		}

		rewardPerWeek := bigNumber.NewFloat(0).Mul(
			bigNumber.NewFloat(0).SetInt(rewardRate),
			bigNumber.NewFloat(0).SetInt(rewardDuration),
		)
		normalizedTotalSupply := helpers.ToNormalizedAmount(rawTotalSupply, rawVaultDecimals)

		/**********************************************************************************************
		** Retrieve the price of the reward token from our storage
		**********************************************************************************************/
		rewardsPrice := bigNumber.NewFloat(0)
		if tokenPrice, ok := storage.GetPrice(vault.ChainID, rewardToken); ok {
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
		rewardDurationAsFloat := bigNumber.NewFloat(0).SetInt(rewardDuration)
		stakingRewardAPR = bigNumber.NewFloat(0).Div(stakingRewardAPR, rewardDurationAsFloat)
		stakingRewardAPR = bigNumber.NewFloat(0).Mul(stakingRewardAPR, secondsPerYear)
		cumulatedStakingRewardAPR = bigNumber.NewFloat(0).Add(cumulatedStakingRewardAPR, stakingRewardAPR)
	}

	/**********************************************************************************************
	** Finally, we can compute the APR
	**********************************************************************************************/
	return cumulatedStakingRewardAPR, true
}
