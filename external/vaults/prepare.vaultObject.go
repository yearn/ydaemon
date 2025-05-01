package vaults

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/common/bigNumber"
	"github.com/yearn/ydaemon/common/helpers"
	"github.com/yearn/ydaemon/internal/models"
	"github.com/yearn/ydaemon/internal/storage"
	"github.com/yearn/ydaemon/processes/risks"
)

/**************************************************************************************************
** buildTokenPrice retrieves and formats the price of a token.
**
** This function handles the complexity of decimals. The prices are fetched using the lens oracle
** daemon, stored in the TokenPrices map, and based on the USDC token (6 decimals). It parses the
** Int Price to a float64, then divides by 10^6 to get the price in a human-readable USDC format.
**
** @param chainID uint64 - The chain ID where the token exists
** @param tokenAddress common.Address - The address of the token
** @return *bigNumber.Float - The humanized price as a bigNumber.Float
** @return float64 - The humanized price as a float64
**************************************************************************************************/
func buildTokenPrice(chainID uint64, tokenAddress common.Address) (*bigNumber.Float, float64) {
	price, ok := storage.GetPrice(chainID, tokenAddress)
	if ok {
		fPrice, _ := price.HumanizedPrice.Float64()
		return price.HumanizedPrice, fPrice
	}
	return bigNumber.NewFloat(), 0.0
}

/**************************************************************************************************
** buildTVL calculates the Total Value Locked (TVL) for a vault.
**
** This function handles the complexity of decimals. The total asset value is formatted as a float64
** and scaled with the token decimals to produce a human-readable Total Asset value. The TVL is then
** calculated by multiplying the humanized balance with the token price.
**
** @param balanceToken *bigNumber.Int - The token balance
** @param decimals int - The number of decimals for the token
** @param humanizedPrice *bigNumber.Float - The humanized price of the token
** @return float64 - The calculated TVL in USD
**************************************************************************************************/
func buildTVL(balanceToken *bigNumber.Int, decimals int, humanizedPrice *bigNumber.Float) float64 {
	_, humanizedTVL := helpers.FormatAmount(balanceToken.String(), decimals)
	fHumanizedTVLPrice, _ := bigNumber.NewFloat().Mul(humanizedTVL, humanizedPrice).Float64()
	return fHumanizedTVLPrice
}

/**************************************************************************************************
** assignStakingRewards processes and formats the staking rewards for a vault.
**
** This function calculates reward rates, duration-scaled rewards, and normalizes values to produce
** a consistent representation of staking rewards that can be presented to users.
**
** @param chainID uint64 - The chain ID where the staking contract exists
** @param stakingData storage.TStakingData - The raw staking data from storage
** @param source string - The source of the staking rewards
** @return TStakingData - The processed staking data with formatted rewards
**************************************************************************************************/
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
			FinishedAt: reward.PeriodFinish,
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

/**************************************************************************************************
** assignStakingData retrieves and formats staking data for a vault.
**
** This function attempts to find staking contracts associated with a vault by checking various
** sources like Convex, Velodrome, Lido, and standard Yearn staking. When a staking contract is
** found, the rewards are processed using assignStakingRewards to create a standardized format.
**
** If no staking contract is found, a default TStakingData with Available=false is returned.
**
** @param chainID uint64 - The chain ID where the vault exists
** @param vaultAddress common.Address - The address of the vault
** @return TStakingData - The processed staking data or a default unavailable staking data
**************************************************************************************************/
func assignStakingData(chainID uint64, vaultAddress common.Address) TStakingData {
	staking := TStakingData{
		Address:   ``,
		Available: false,
	}

	/**********************************************************************************************
	** Check, retrieve and assign the staking pool for the vault. The staking pool here can only
	** be an op boost staking pool.
	**********************************************************************************************/
	opStakingData, hasStakingPool := storage.GetOPStakingForVault(chainID, vaultAddress)
	if !staking.Available && hasStakingPool {
		staking = assignStakingRewards(chainID, opStakingData, `OP Boost`)
	}

	veYFIStakingData, hasVeYFIGauge := storage.GetVeYFIStakingForVault(chainID, vaultAddress)
	if !staking.Available && hasVeYFIGauge {
		staking = assignStakingRewards(chainID, veYFIStakingData, `VeYFI`)
	}

	juicedStakingData, hasJuicedGauge := storage.GetJuicedStakingDataForVault(chainID, vaultAddress)
	if !staking.Available && hasJuicedGauge {
		staking = assignStakingRewards(chainID, juicedStakingData, `Juiced`)
	}

	v3StakingData, hasV3Gauge := storage.GetV3StakingDataForVault(chainID, vaultAddress)
	if !staking.Available && hasV3Gauge {
		staking = assignStakingRewards(chainID, v3StakingData, `V3 Staking`)
	}
	return staking
}

/************************************************************************************************
** toSimplifiedVersion converts a full vault object into a simplified representation.
**
** This function optimizes memory usage by directly mapping required fields instead of creating
** intermediate objects.
**
** @param vault TExternalVault - The complete vault data
** @param vaultAsStrategy models.TStrategy - Optional strategy data if the vault is also a strategy
** @return TSimplifiedExternalVault - The simplified vault representation
************************************************************************************************/
func toSimplifiedVersion(
	vault TExternalVault,
	vaultAsStrategy models.TStrategy,
) TSimplifiedExternalVault {
	// Determine the best name to use without creating intermediate strings
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

	// Determine the best symbol to use without creating intermediate strings
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

	// Get token info using the helper function
	tokenInfo := getSimplifiedTokenInfo(vault.Token)

	// Handle risk information
	// Create a copy of the info struct to avoid mutating the original
	info := vault.Info
	info.IsRetired = vault.Details.IsRetired
	info.IsBoosted = vault.Details.IsBoosted
	info.IsHighlighted = vault.Details.IsHighlighted

	// Try to use cached risk score if available
	cachedRiskScore, err := risks.GetCachedRiskScore(vault.ChainID, common.HexToAddress(vault.Address))
	if err == nil {
		info.RiskLevel = cachedRiskScore.RiskLevel
		info.RiskScore = [11]int8{
			cachedRiskScore.RiskScore.Review,
			cachedRiskScore.RiskScore.Testing,
			cachedRiskScore.RiskScore.Complexity,
			cachedRiskScore.RiskScore.RiskExposure,
			cachedRiskScore.RiskScore.ProtocolIntegration,
			cachedRiskScore.RiskScore.CentralizationRisk,
			cachedRiskScore.RiskScore.ExternalProtocolAudit,
			cachedRiskScore.RiskScore.ExternalProtocolCentralisation,
			cachedRiskScore.RiskScore.ExternalProtocolTvl,
			cachedRiskScore.RiskScore.ExternalProtocolLongevity,
			cachedRiskScore.RiskScore.ExternalProtocolType,
		}
		info.RiskScoreComment = cachedRiskScore.RiskScore.Comment
	}

	// Create the simplified vault directly without intermediate objects
	return TSimplifiedExternalVault{
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
		Token:          tokenInfo,
		TVL: TSimplifiedExternalVaultTVL{
			TotalAssets: vault.TVL.TotalAssets,
			TVL:         vault.TVL.TVL,
			Price:       vault.TVL.Price,
		},
		Strategies:    vault.Strategies,
		Staking:       assignStakingData(vault.ChainID, common.HexToAddress(vault.Address)),
		Info:          info,
		PricePerShare: vault.PricePerShare,
	}
}

/**************************************************************************************************
** getSimplifiedTokenInfo creates a simplified token representation for API responses.
**
** This function extracts only the essential token information needed for simplified views,
** prioritizing display names and symbols when available. It provides fallbacks to ensure
** that even with missing data, a valid and useful token representation is returned.
**
** @param token TExternalERC20Token - The full token data structure
** @return TSimplifiedExternalERC20Token - The simplified token representation
**************************************************************************************************/
func getSimplifiedTokenInfo(token TExternalERC20Token) TSimplifiedExternalERC20Token {
	// Determine token name with priority for display name
	tokenName := token.DisplayName
	if tokenName == "" {
		tokenName = token.Name
	}
	if tokenName == "" {
		tokenName = `Unknown`
	}

	// Determine token symbol with priority for display symbol
	tokenSymbol := token.DisplaySymbol
	if tokenSymbol == "" {
		tokenSymbol = token.Symbol
	}
	if tokenSymbol == "" {
		tokenSymbol = `Unknown`
	}

	return TSimplifiedExternalERC20Token{
		Address:     token.Address,
		Name:        tokenName,
		Symbol:      tokenSymbol,
		Description: token.Description,
		Decimals:    token.Decimals,
	}
}
