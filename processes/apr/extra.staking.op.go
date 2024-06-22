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

func computeOPBoostStakingRewardsAPR(chainID uint64, vault models.TVault) (*bigNumber.Float, bool) {
	/**********************************************************************************************
	** First thing to do is to check if the vault has a staking contract associated with it.
	** We can retrieve that from the store.
	**********************************************************************************************/
	stakingContract, ok := storage.GetOPStakingForVault(chainID, vault.Address)
	if !ok {
		return bigNumber.NewFloat(0), false
	}
	if len(stakingContract.RewardTokens) == 0 {
		return bigNumber.NewFloat(0), false
	}

	/**********************************************************************************************
	** Once we got it we will need a few things from the staking contract. We will use a multicall
	** to retrieve the following:
	** - The totalSupply
	** - The vault decimals
	**********************************************************************************************/
	calls := []ethereum.Call{}
	calls = append(calls, multicalls.GetTotalSupply(stakingContract.StakingAddress.Hex(), stakingContract.StakingAddress))
	response := multicalls.Perform(chainID, calls, nil)
	rawTotalSupply := helpers.DecodeBigInt(response[stakingContract.StakingAddress.Hex()+`totalSupply`])

	rewardPeriodFinish := stakingContract.RewardTokens[0].PeriodFinish
	rewardRate := stakingContract.RewardTokens[0].Rate
	rewardsToken := stakingContract.RewardTokens[0].Address

	/**********************************************************************************************
	** If periodFinish is before now, aka rewards are over, we can stop here
	**********************************************************************************************/
	now := time.Now().Unix()
	if rewardPeriodFinish < uint64(now) {
		storage.AssignOPStakingRewardAPR(chainID, vault.Address, rewardsToken, bigNumber.NewFloat(0))
		return bigNumber.NewFloat(0), false
	}

	/**********************************************************************************************
	** If the total supply is 0, we can stop here, aka nothing is staked, so no rewards
	**********************************************************************************************/
	if rawTotalSupply.IsZero() {
		storage.AssignOPStakingRewardAPR(chainID, vault.Address, rewardsToken, bigNumber.NewFloat(0))
		return bigNumber.NewFloat(0), false
	}

	/**********************************************************************************************
	** If that's good, we will need the price of the vault token and the price of the rewards token
	** to compute the APR.
	**********************************************************************************************/
	vaultPrice := bigNumber.NewFloat(0)
	if tokenPrice, ok := storage.GetPrice(vault.ChainID, vault.Address); ok {
		vaultPrice = tokenPrice.HumanizedPrice
	}

	rewardsPrice := bigNumber.NewFloat(0)
	if tokenPrice, ok := storage.GetPrice(vault.ChainID, rewardsToken); ok {
		rewardsPrice = tokenPrice.HumanizedPrice
	}

	/**********************************************************************************************
	** Then, we need to scale the decimals of the rewardRate and the totalSupply to match the
	** decimals of the vault.
	**********************************************************************************************/
	rewardRateFloat := bigNumber.NewFloat(0).SetUint64(rewardRate)
	totalSupplyFloat := bigNumber.NewFloat(0).SetInt(rawTotalSupply)
	perStakingTokenRate := bigNumber.NewFloat(0).Div(rewardRateFloat, totalSupplyFloat)
	secondsPerYear := bigNumber.NewFloat(31_556_952)

	/**********************************************************************************************
	** Finally, we can compute the APR
	**********************************************************************************************/
	stakingRewardAPR := bigNumber.NewFloat(0).Mul(secondsPerYear, perStakingTokenRate)
	stakingRewardAPR = bigNumber.NewFloat(0).Mul(stakingRewardAPR, rewardsPrice)
	stakingRewardAPR = bigNumber.NewFloat(0).Div(stakingRewardAPR, vaultPrice)

	storage.AssignOPStakingRewardAPR(chainID, vault.Address, rewardsToken, stakingRewardAPR)
	return stakingRewardAPR, true
}
