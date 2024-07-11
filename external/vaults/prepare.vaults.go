package vaults

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/common/bigNumber"
	"github.com/yearn/ydaemon/common/helpers"
	"github.com/yearn/ydaemon/internal/models"
	"github.com/yearn/ydaemon/internal/storage"
)

// Get the price of the underlying asset. This is tricky because of the decimals. The prices are fetched
// using the lens oracle daemon, stored in the TokenPrices map, and based on the USDC token, aka with 6
// decimals. We first need to parse the Int Price to a float64, then divide by 10^6 to get the price
// in an human readable USDC format.
func buildTokenPrice(chainID uint64, tokenAddress common.Address) (*bigNumber.Float, float64) {
	price, ok := storage.GetPrice(chainID, tokenAddress)
	if ok {
		fPrice, _ := price.HumanizedPrice.Float64()
		return price.HumanizedPrice, fPrice
	}
	return bigNumber.NewFloat(), 0.0
}

// Get the total assets locked in this vault. This is tricky because of the decimals. The total asset value
// is a string which will be formated as a float64 and scaled with the underlying token decimals. With that
// we should have a human readable Total Asset value, and we should be able to get the Total Value Locked
// in the vault thanks to the above humanizedPrice value.
func buildTVL(balanceToken *bigNumber.Int, decimals int, humanizedPrice *bigNumber.Float) float64 {
	_, humanizedTVL := helpers.FormatAmount(balanceToken.String(), decimals)
	fHumanizedTVLPrice, _ := bigNumber.NewFloat().Mul(humanizedTVL, humanizedPrice).Float64()
	return fHumanizedTVLPrice
}

func assignStakingRewards(chainID uint64, stakingData storage.TStakingData, source string) TStakingData {
	rewards := []TStakingRewardsData{}
	for _, reward := range stakingData.RewardTokens {
		normalizedRewardRate := helpers.ToNormalizedAmount(bigNumber.NewFloat(0).SetUint64(reward.Rate).Int(), reward.Decimals)
		rewardPerDuration := bigNumber.NewFloat(0).Mul(normalizedRewardRate, bigNumber.NewFloat(0).SetUint64(reward.Duration))
		durationScaledToWeek := bigNumber.NewFloat(0).Div(bigNumber.NewFloat(0).SetUint64(reward.Duration), bigNumber.NewFloat(0).SetUint64(604800))
		rewardsPerWeek := bigNumber.NewFloat(0).Div(rewardPerDuration, durationScaledToWeek)
		_, tokenPrice := buildTokenPrice(chainID, reward.Address)

		if reward.IsFinished {
			rewardsPerWeek = bigNumber.NewFloat()
		}
		rewards = append(rewards, TStakingRewardsData{
			Address:    reward.Address.Hex(),
			Name:       reward.Name,
			Symbol:     reward.Symbol,
			Decimals:   reward.Decimals,
			IsFinished: reward.IsFinished,
			APR:        reward.APR,
			PerWeek:    rewardsPerWeek,
			Price:      tokenPrice,
		})
	}
	staking := TStakingData{
		Address:   stakingData.StakingAddress.Hex(),
		Available: true,
		Source:    source,
		Rewards:   rewards,
	}
	return staking
}

func toSimplifiedVersion(
	vault TExternalVault,
	vaultAsStrategy models.TStrategy,
) TSimplifiedExternalVault {
	vaultName := vault.DisplayName
	if (vaultName == "" && vaultAsStrategy.Address != common.Address{}) {
		vaultName = vaultAsStrategy.DisplayName
	}
	if vaultName == "" {
		vaultName = vault.Name
	}
	if vaultName == "" {
		vaultName = vault.FormatedName
	}
	if vaultName == "" {
		vaultName = `Unknown`
	}

	vaultSymbol := vault.DisplaySymbol
	if vaultSymbol == "" {
		vaultSymbol = vault.Symbol
	}
	if vaultSymbol == "" {
		vaultSymbol = vault.FormatedSymbol
	}
	if vaultSymbol == "" {
		vaultSymbol = `Unknown`
	}

	tokenName := vault.Token.DisplayName
	if tokenName == "" {
		tokenName = vault.Token.Name
	}
	if tokenName == "" {
		tokenName = `Unknown`
	}

	tokenSymbol := vault.Token.DisplaySymbol
	if tokenSymbol == "" {
		tokenSymbol = vault.Token.Symbol
	}
	if tokenSymbol == "" {
		tokenSymbol = `Unknown`
	}

	staking := TStakingData{
		Address:   ``,
		Available: false,
	}

	/**********************************************************************************************
	** Check, retrieve and assign the staking pool for the vault. The staking pool here can only
	** be an op boost staking pool.
	**********************************************************************************************/
	opStakingData, hasStakingPool := storage.GetOPStakingForVault(vault.ChainID, common.HexToAddress(vault.Address))
	if !staking.Available && hasStakingPool {
		staking = assignStakingRewards(vault.ChainID, opStakingData, `OP Boost`)
	}

	veYFIStakingData, hasVeYFIGauge := storage.GetVeYFIStakingForVault(vault.ChainID, common.HexToAddress(vault.Address))
	if !staking.Available && hasVeYFIGauge {
		staking = assignStakingRewards(vault.ChainID, veYFIStakingData, `VeYFI`)
	}

	juicedStakingData, hasJuicedGauge := storage.GetJuicedStakingDataForVault(vault.ChainID, common.HexToAddress(vault.Address))
	if !staking.Available && hasJuicedGauge {
		staking = assignStakingRewards(vault.ChainID, juicedStakingData, `Juiced`)
	}

	v3StakingData, hasV3Gauge := storage.GetV3StakingDataForVault(vault.ChainID, common.HexToAddress(vault.Address))
	if !staking.Available && hasV3Gauge {
		staking = assignStakingRewards(vault.ChainID, v3StakingData, `V3 Staking`)
	}

	info := vault.Info
	info.IsRetired = vault.Details.IsRetired
	info.IsBoosted = vault.Details.IsBoosted
	info.IsHighlighted = vault.Details.IsHighlighted
	info.RiskLevel = vault.Info.RiskLevel

	/**********************************************************************************************
	** Create the simplified version of the vault.
	** The simplified version of the vault is a struct that contains only the necessary data
	** to be displayed in the frontend.
	**********************************************************************************************/
	simplifiedVault := TSimplifiedExternalVault{
		Address:        vault.Address,
		Type:           vault.Type,
		Kind:           vault.Kind,
		Symbol:         vault.Symbol,
		Name:           vaultName,
		Description:    vault.Description,
		Category:       vault.Category,
		Decimals:       vault.Decimals,
		ChainID:        vault.ChainID,
		APR:            vault.APR,
		Migration:      vault.Migration,
		Version:        vault.Version,
		FeaturingScore: vault.FeaturingScore,
		Token: TSimplifiedExternalERC20Token{
			Address:     vault.Token.Address,
			Name:        tokenName,
			Symbol:      tokenSymbol,
			Description: vault.Token.Description,
			Decimals:    vault.Token.Decimals,
		},
		TVL: TSimplifiedExternalVaultTVL{
			TotalAssets: vault.TVL.TotalAssets,
			TVL:         vault.TVL.TVL,
			Price:       vault.TVL.Price,
		},
		Strategies:    vault.Strategies,
		Staking:       staking,
		Info:          info,
		PricePerShare: vault.PricePerShare,
	}
	return simplifiedVault
}
