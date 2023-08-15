package apy

import (
	"time"

	"github.com/yearn/ydaemon/common/bigNumber"
	"github.com/yearn/ydaemon/common/ethereum"
	"github.com/yearn/ydaemon/common/helpers"
	"github.com/yearn/ydaemon/common/logs"
	"github.com/yearn/ydaemon/common/store"
	"github.com/yearn/ydaemon/internal/models"
	"github.com/yearn/ydaemon/internal/multicalls"
)

func computeStakingRewardsAPR(chainID uint64, vault models.TVault) {
	/**********************************************************************************************
	** First thing to do is to check if the vault has a staking contract associated with it.
	** We can retrieve that from the store.
	**********************************************************************************************/
	stakingContract, ok := store.GetStakingPoolForVault(chainID, vault.Address)
	if !ok {
		return
	}

	/**********************************************************************************************
	** Once we got it we will need a few things from the staking contract. We will use a multicall
	** to retrieve the following:
	** - The periodFinish
	** - The rewardRate
	** - The totalSupply
	**********************************************************************************************/
	calls := []ethereum.Call{}
	calls = append(calls, multicalls.GetPeriodFinish(stakingContract.StackingPoolAddress.Hex(), stakingContract.StackingPoolAddress))
	calls = append(calls, multicalls.GetRewardRate(stakingContract.StackingPoolAddress.Hex(), stakingContract.StackingPoolAddress))
	calls = append(calls, multicalls.GetTotalSupply(stakingContract.StackingPoolAddress.Hex(), stakingContract.StackingPoolAddress))
	calls = append(calls, multicalls.GetRewardsToken(stakingContract.StackingPoolAddress.Hex(), stakingContract.StackingPoolAddress))
	response := multicalls.Perform(chainID, calls, nil)
	periodFinish := helpers.DecodeBigInt(response[stakingContract.StackingPoolAddress.Hex()+`periodFinish`])
	rewardRateRaw := helpers.DecodeBigInt(response[stakingContract.StackingPoolAddress.Hex()+`rewardRate`])
	totalSupplyRaw := helpers.DecodeBigInt(response[stakingContract.StackingPoolAddress.Hex()+`totalSupply`])
	rewardsToken := helpers.DecodeAddress(response[stakingContract.StackingPoolAddress.Hex()+`rewardsToken`])

	/**********************************************************************************************
	** If periodFinish is before now, aka rewards are over, we can stop here
	**********************************************************************************************/
	now := time.Now().Unix()
	if periodFinish.Int64() < now {
		return
	}

	/**********************************************************************************************
	** If the total supply is 0, we can stop here, aka nothing is staked, so no rewards
	**********************************************************************************************/
	if totalSupplyRaw.IsZero() {
		return
	}

	/**********************************************************************************************
	** If that's good, we will need the price of the vault token and the price of the rewards token
	** to compute the APR.
	**********************************************************************************************/
	logs.Info(`CHECKING PRICE FOR VAULT TOKEN`)
	vaultPrice := getTokenPrice(chainID, vault.Address)
	logs.Info(`CHECKING PRICE FOR REWARD TOKEN`)
	rewardsPrice := getTokenPrice(chainID, rewardsToken)
	logs.Pretty(vaultPrice, rewardsPrice)

	/**********************************************************************************************
	** Then, we need to scale the decimals of the rewardRate and the totalSupply to match the
	** decimals of the vault.
	**********************************************************************************************/
	rewardRate := helpers.ToNormalizedAmount(rewardRateRaw, vault.Decimals)
	totalSupply := helpers.ToNormalizedAmount(totalSupplyRaw, vault.Decimals)
	perStakingTokenRate := bigNumber.NewFloat(0).Div(rewardRate, totalSupply)
	secondsPerYear := bigNumber.NewFloat(31_556_952)

	emissionsAPR := bigNumber.NewFloat(0).Mul(secondsPerYear, perStakingTokenRate)
	emissionsAPR = bigNumber.NewFloat(0).Mul(emissionsAPR, rewardsPrice)
	emissionsAPR = bigNumber.NewFloat(0).Div(emissionsAPR, vaultPrice)

	// rewards_vault_apy = (await rewards_vault.apy(samples)).net_apy
	// emissions_apr = SECONDS_PER_YEAR * per_staking_token_rate * rewards_asset_price / vault_price
	// return emissions_apr * (1 + rewards_vault_apy)

	logs.Pretty(
		`Staking contract found for `+vault.Address.Hex()+`: `+stakingContract.StackingPoolAddress.Hex(),
		periodFinish,
		rewardRate,
		totalSupply,
		emissionsAPR,
	)

}
