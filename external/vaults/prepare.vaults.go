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
func buildTokenPrice(chainID uint64, tokenAddress common.MixedcaseAddress) (*bigNumber.Float, float64) {
	price, ok := storage.GetPrice(chainID, tokenAddress.Address())
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
	stakingData, hasStakingPool := storage.GetStakingPoolForVault(vault.ChainID, common.HexToAddress(vault.Address))
	if hasStakingPool {
		staking = TStakingData{
			Address:   stakingData.StackingPoolAddress.Hex(),
			Available: hasStakingPool,
			Source:    `OP Boost`,
		}
	}

	gaugeData, hasVeYFIGauge := storage.GetGaugeForVault(vault.ChainID, common.HexToAddress(vault.Address))
	if !staking.Available && hasVeYFIGauge {
		staking = TStakingData{
			Address:   gaugeData.Hex(),
			Available: hasVeYFIGauge,
			Source:    `VeYFI`,
		}
	}

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
		Category:       vault.Category,
		Decimals:       vault.Decimals,
		ChainID:        vault.ChainID,
		APR:            vault.APR,
		Migration:      vault.Migration,
		Retired:        vault.Details.IsRetired,
		Boosted:        vault.Details.IsBoosted,
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
		Strategies: vault.Strategies,
		Staking:    staking,
		Info:       vault.Info,
	}
	return simplifiedVault
}
