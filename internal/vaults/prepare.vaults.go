package vaults

import (
	"math"
	"strconv"
	"strings"

	"github.com/yearn/ydaemon/internal/meta"
	"github.com/yearn/ydaemon/internal/prices"
	"github.com/yearn/ydaemon/internal/strategies"
	"github.com/yearn/ydaemon/internal/utils/env"
	"github.com/yearn/ydaemon/internal/utils/helpers"
	"github.com/yearn/ydaemon/internal/utils/models"
	"github.com/yearn/ydaemon/internal/utils/types/bigNumber"
	"github.com/yearn/ydaemon/internal/utils/types/common"
)

func buildVaultName(
	chainID uint64,
	vaultAddress common.Address,
	vaultName string,
	metaVaultName string,
	tokenName string,
) (name string, displayName string, formatedName string) {
	name = strings.Replace(vaultName, "\"", "", -1)
	formatedName = tokenName
	if metaVaultName != "" {
		displayName = metaVaultName
	} else {
		vaultFromMeta, ok := meta.Store.VaultsFromMeta[chainID][vaultAddress]
		if ok {
			displayName = strings.Replace(vaultFromMeta.DisplayName, "\"", "", -1)
		}
	}

	//If the formated name is missing yVault suffix, add it
	if !strings.HasSuffix(formatedName, "yVault") {
		formatedName = formatedName + " yVault"
	}
	// If a display name exist, use it for the formating.
	if displayName != "" && !strings.HasSuffix(displayName, "yVault") {
		formatedName = displayName + " yVault"
	}
	// If the name is empty, use the displayName instead
	if name == "" {
		name = displayName
	}
	// If the name is still empty, use the formated name instead
	if name == "" {
		name = formatedName
	}
	if displayName == "" && strings.HasSuffix(name, " Auto-Compounding yVault") {
		displayName = strings.Replace(name, " Auto-Compounding yVault", "", -1)
	}

	return name, displayName, formatedName
}

func buildVaultSymbol(
	chainID uint64,
	tokenAddress common.Address,
	vaultSymbol string,
	tokenSymbol string,
) (symbol string, displaySymbol string, formatedSymbol string) {
	symbol = strings.Replace(vaultSymbol, "\"", "", -1)
	formatedSymbol = tokenSymbol
	shareTokenFromMeta, ok := meta.Store.TokensFromMeta[chainID][tokenAddress]
	if ok {
		displaySymbol = strings.Replace(shareTokenFromMeta.Symbol, "\"", "", -1)
	}

	//If the formated symbol is missing yv prefix, add it
	if !strings.HasPrefix(formatedSymbol, "yv") {
		formatedSymbol = "yv" + formatedSymbol
	}
	// If a display name exist, use it for the formating.
	if displaySymbol != "" && !strings.HasPrefix(displaySymbol, "yv") {
		formatedSymbol = "yv" + displaySymbol
	}
	symbol = helpers.SafeString(symbol, displaySymbol)
	symbol = helpers.SafeString(symbol, formatedSymbol)
	displaySymbol = helpers.SafeString(displaySymbol, symbol)

	return symbol, displaySymbol, formatedSymbol
}

// Get the price of the underlying asset. This is tricky because of the decimals. The prices are fetched
// using the lens oracle daemon, stored in the TokenPrices map, and based on the USDC token, aka with 6
// decimals. We first need to parse the Int Price to a float64, then divide by 10^6 to get the price
// in an human readable USDC format.
func buildTokenPrice(chainID uint64, tokenAddress common.Address) (*bigNumber.Float, float64) {
	prices := prices.Store.TokenPrices[chainID]
	fPrice := new(bigNumber.Float)
	price, ok := prices[tokenAddress]
	if ok {
		fPrice.SetInt(price)
		humanizedPrice := new(bigNumber.Float).Quo(fPrice, bigNumber.NewFloat(math.Pow10(int(6))))
		fHumanizedPrice, _ := humanizedPrice.Float64()
		return humanizedPrice, fHumanizedPrice
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

// From the legacy API, build the schema for the APY, models.TAPY, used to get the details and the
// breakdown of the vault.
func buildAPY(
	vaultAddress common.Address,
	chainID uint64,
	PerformanceFeeBps uint64,
	ManagementFeeBps uint64,
	APYTypeOverride string,
) TAPY {
	apy := TAPY{}
	apyFromAPIV1, ok := Store.VaultsFromAPIV1[chainID][vaultAddress]

	if ok {
		apy = TAPY{
			Type:     apyFromAPIV1.APY.Type,
			GrossAPR: apyFromAPIV1.APY.GrossAPR,
			NetAPY:   apyFromAPIV1.APY.NetAPY,
			Points: TAPYPoints{
				WeekAgo:   apyFromAPIV1.APY.Points.WeekAgo,
				MonthAgo:  apyFromAPIV1.APY.Points.MonthAgo,
				Inception: apyFromAPIV1.APY.Points.Inception,
			},
			Composite: TAPYComposite{
				Boost:      apyFromAPIV1.APY.Composite.Boost,
				PoolAPY:    apyFromAPIV1.APY.Composite.PoolAPY,
				BoostedAPR: apyFromAPIV1.APY.Composite.BoostedAPR,
				BaseAPR:    apyFromAPIV1.APY.Composite.BaseAPR,
				CvxAPR:     apyFromAPIV1.APY.Composite.CvxAPR,
				RewardsAPR: apyFromAPIV1.APY.Composite.RewardsAPR,
			},
			Fees: TAPYFees{
				Performance: float64(PerformanceFeeBps) / 10000,
				Management:  float64(ManagementFeeBps) / 10000,
				Withdrawal:  apyFromAPIV1.APY.Fees.Withdrawal,
				KeepCRV:     apyFromAPIV1.APY.Fees.KeepCRV,
				CvxKeepCRV:  apyFromAPIV1.APY.Fees.CvxKeepCRV,
			},
		}
	}
	if APYTypeOverride != "" {
		apy.Type = APYTypeOverride
	}
	return apy
}

// Get the migration informations for this vault. By default, migrationAvailable is false and
// the migrationAddress matches the vault address. If a migration is available, point this last
// one to MigrationTargetVault.
func buildMigration(chainID uint64, vaultAddress common.Address) TMigration {
	migration := TMigration{}
	vaultFromMeta, ok := meta.Store.VaultsFromMeta[chainID][vaultAddress]

	if ok {
		migrationAddress := vaultAddress
		migrationAvailable := vaultFromMeta.MigrationAvailable
		if vaultFromMeta.MigrationAvailable {
			migrationAddress = vaultFromMeta.MigrationTargetVault
		}

		migration = TMigration{
			Available: migrationAvailable,
			Address:   migrationAddress,
		}
	}
	return migration
}

func prepareVaultSchema(
	chainID uint64,
	strategiesCondition string,
	withStrategiesRisk bool,
	withStrategiesDetails bool,
	vaultFromGraph models.TVaultFromGraph,
) *TVault {
	chainIDAsString := strconv.FormatUint(chainID, 10)
	vaultAddress := vaultFromGraph.Id
	tokenAddress := vaultFromGraph.Token.Id
	tokenFromMeta := meta.Store.TokensFromMeta[chainID][tokenAddress]
	updated := helpers.FormatUint64(vaultFromGraph.LatestUpdate.Timestamp, 0)
	activation := helpers.FormatUint64(vaultFromGraph.Activation, 0)
	tokenDisplayName := helpers.SafeString(tokenFromMeta.Name, vaultFromGraph.Token.Name)
	tokenDisplaySymbol := helpers.SafeString(tokenFromMeta.Symbol, vaultFromGraph.Token.Symbol)
	vaultFromMeta, ok := meta.Store.VaultsFromMeta[chainID][vaultAddress]
	if !ok {
		// If the vault file is missing, we set the default values for its fields
		vaultFromMeta = meta.TVaultFromMeta{
			Order:               1000000000,
			HideAlways:          false,
			DepositsDisabled:    false,
			WithdrawalsDisabled: false,
			MigrationAvailable:  false,
			AllowZapIn:          true,
			AllowZapOut:         true,
			Retired:             false,
		}
	}

	vaultName, vaultDisplayName, vaultFormatedName := buildVaultName(
		chainID,
		vaultAddress,
		vaultFromGraph.ShareToken.Name,
		vaultFromMeta.DisplayName,
		vaultFromGraph.Token.Name,
	)
	vaultSymbol, vaultDisplaySymbol, vaultFormatedSymbol := buildVaultSymbol(
		chainID,
		vaultFromGraph.ShareToken.Id,
		vaultFromGraph.ShareToken.Symbol,
		vaultFromGraph.Token.Symbol,
	)
	humanizedPrice, fHumanizedPrice := buildTokenPrice(
		chainID,
		tokenAddress,
	)

	strategies := strategies.BuildStrategies(
		chainID,
		withStrategiesDetails,
		withStrategiesRisk,
		strategiesCondition,
		humanizedPrice,
		vaultFromGraph,
	)

	fHumanizedTVLPrice := buildTVL(
		vaultFromGraph.BalanceTokens,
		int(vaultFromGraph.Token.Decimals),
		humanizedPrice,
	)
	delegatedTokenAsBN := bigNumber.NewInt(0)
	fDelegatedValue := 0.0

	for _, strat := range strategies {
		stratDelegatedValueAsFloat, err := strconv.ParseFloat(strat.DelegatedValue, 64)
		if err == nil {
			delegatedTokenAsBN = delegatedTokenAsBN.Add(delegatedTokenAsBN, strat.DelegatedAssets)
			fDelegatedValue += stratDelegatedValueAsFloat
		}
	}

	vault := &TVault{
		Inception:      activation,
		Address:        vaultAddress,
		Symbol:         vaultSymbol,
		DisplaySymbol:  vaultDisplaySymbol,
		FormatedSymbol: vaultFormatedSymbol,
		Name:           vaultName,
		DisplayName:    vaultDisplayName,
		FormatedName:   vaultFormatedName,
		Icon:           env.GITHUB_ICON_BASE_URL + chainIDAsString + `/` + vaultAddress.String() + `/logo-128.png`,
		Token: TToken{
			Address:     vaultFromGraph.Token.Id,
			Name:        vaultFromGraph.Token.Name,
			DisplayName: tokenDisplayName,
			Symbol:      tokenDisplaySymbol,
			Description: tokenFromMeta.Description,
			Decimals:    vaultFromGraph.Token.Decimals,
			Icon:        env.GITHUB_ICON_BASE_URL + chainIDAsString + `/` + vaultFromGraph.Token.Id.String() + `/logo-128.png`,
		},
		TVL: TTVL{
			TotalAssets:          vaultFromGraph.BalanceTokens,
			TotalDelegatedAssets: delegatedTokenAsBN,
			TVL:                  fHumanizedTVLPrice - fDelegatedValue,
			TVLDeposited:         fHumanizedTVLPrice,
			TVLDelegated:         fDelegatedValue,
			Price:                fHumanizedPrice,
		},
		Details: &TVaultDetails{
			Management:            vaultFromGraph.Management,
			Governance:            vaultFromGraph.Governance,
			Guardian:              vaultFromGraph.Guardian,
			Rewards:               vaultFromGraph.Rewards,
			DepositLimit:          vaultFromGraph.DepositLimit,
			Comment:               vaultFromMeta.Comment,
			AvailableDepositLimit: vaultFromGraph.AvailableDepositLimit,
			Order:                 vaultFromMeta.Order,
			PerformanceFee:        vaultFromGraph.PerformanceFeeBps,
			ManagementFee:         vaultFromGraph.ManagementFeeBps,
			DepositsDisabled:      vaultFromMeta.DepositsDisabled,
			WithdrawalsDisabled:   vaultFromMeta.WithdrawalsDisabled,
			AllowZapIn:            vaultFromMeta.AllowZapIn,
			AllowZapOut:           vaultFromMeta.AllowZapOut,
			Retired:               vaultFromMeta.Retired,
			APYTypeOverride:       vaultFromMeta.APYTypeOverride,
			APYOverride:           vaultFromMeta.APYOverride,
		},
		APY: buildAPY(
			vaultAddress,
			chainID,
			vaultFromGraph.PerformanceFeeBps,
			vaultFromGraph.ManagementFeeBps,
			vaultFromMeta.APYTypeOverride,
		),
		Strategies: strategies,
		Endorsed:   vaultFromGraph.Classification == "Endorsed",
		Version:    vaultFromGraph.ApiVersion,
		Decimals:   vaultFromGraph.ShareToken.Decimals,
		Type:       "v2", //No v1 in the subgraph
		// Emergency_shutdown: ,
		Updated:   updated / 1000,
		Migration: buildMigration(chainID, vaultAddress),
	}

	return vault
}
