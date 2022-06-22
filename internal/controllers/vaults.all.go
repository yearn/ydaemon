package controllers

import (
	"math"
	"math/big"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/majorfi/ydaemon/internal/models"
	"github.com/majorfi/ydaemon/internal/utils"
)

func getVaultName(
	vaultFromGraph models.TVaultFromGraph,
	vaultFromMeta models.TVaultFromMeta,
	tokenFromMeta models.TTokenFromMeta,
) (name string, displayName string, formatedName string) {
	name = vaultFromGraph.ShareToken.Name
	displayName = vaultFromMeta.DisplayName
	formatedName = vaultFromGraph.Token.Name

	// Remove all extra "
	name = strings.Replace(name, "\"", "", -1)
	displayName = strings.Replace(displayName, "\"", "", -1)

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

	return name, displayName, formatedName
}

func getVaultSymbol(
	vaultFromGraph models.TVaultFromGraph,
	vaultFromMeta models.TVaultFromMeta,
	shareTokenFromMeta models.TTokenFromMeta,
) (symbol string, displaySymbol string, formatedSymbol string) {
	symbol = vaultFromGraph.ShareToken.Symbol
	displaySymbol = shareTokenFromMeta.Symbol
	formatedSymbol = vaultFromGraph.Token.Symbol

	// Remove all extra "
	symbol = strings.Replace(symbol, "\"", "", -1)
	displaySymbol = strings.Replace(displaySymbol, "\"", "", -1)

	//If the formated symbol is missing yv prefix, add it
	if !strings.HasPrefix(formatedSymbol, "yv") {
		formatedSymbol = "yv" + formatedSymbol
	}
	// If a display name exist, use it for the formating.
	if displaySymbol != "" && !strings.HasPrefix(displaySymbol, "yv") {
		formatedSymbol = "yv" + displaySymbol
	}
	// If the symbol is empty, use the displaySymbol instead
	if symbol == "" {
		symbol = displaySymbol
	}
	// If the symbol is still empty, use the formated symbol instead
	if symbol == "" {
		symbol = displaySymbol
	}

	return symbol, displaySymbol, formatedSymbol
}

func prepareStrategies(
	strategiesCondition string,
	vaultFromGraph models.TVaultFromGraph,
	strategiesFromMeta map[string]models.TStrategyFromMeta,
) []models.TStrategy {
	strategies := []models.TStrategy{}
	for _, strategy := range vaultFromGraph.Strategies {
		if strategiesCondition == `inQueue` {
			if strategy.InQueue {
				strategies = append(strategies, models.TStrategy{
					Address:     common.HexToAddress(strategy.Address).String(),
					Name:        strategy.Name,
					Description: strategiesFromMeta[common.HexToAddress(strategy.Address).String()].Description,
				})
			}
		}
		if strategiesCondition == `debtLimit` {
			debtLimit, err := strconv.ParseUint(strategy.DebtLimit, 10, 64)
			if debtLimit > 0 && err == nil {
				strategies = append(strategies, models.TStrategy{
					Address:     common.HexToAddress(strategy.Address).String(),
					Name:        strategy.Name,
					Description: strategiesFromMeta[common.HexToAddress(strategy.Address).String()].Description,
				})
			}
		}
	}
	return strategies
}

func prepareVaultSchema(
	chainID uint64,
	strategiesCondition string,
	vaultFromGraph models.TVaultFromGraph,
	vaultFromMeta models.TVaultFromMeta,
	shareTokenFromMeta models.TTokenFromMeta,
	tokenFromMeta models.TTokenFromMeta,
	strategiesFromMeta map[string]models.TStrategyFromMeta,
	apyFromAPIV1 models.TAPIV1Vault,
	prices map[common.Address]uint64,
) *models.TVault {
	chainIDAsString := strconv.FormatUint(chainID, 10)

	// Get the last time the vault was updated. On error, time set to 0.
	updated, err := strconv.ParseUint(vaultFromGraph.LatestUpdate.Timestamp, 10, 64)
	if err != nil {
		updated = 0
	}

	// Get the activation date for this vault. On error, set activation to 0.
	activation, err := strconv.ParseUint(vaultFromGraph.Activation, 10, 64)
	if err != nil {
		activation = 0
	}

	// Get the displayName for the underlying token. If the data is not set, add a fallback to the token Name.
	tokenDisplayName := tokenFromMeta.Name
	if tokenDisplayName == "" || tokenDisplayName == "\"\"" {
		tokenDisplayName = vaultFromGraph.Token.Name
	}

	// Get the symbol for the underlying token. If the data is not set, add a fallback to the actual symbol.
	tokenDisplaySymbol := tokenFromMeta.Symbol
	if tokenDisplaySymbol == "" || tokenDisplaySymbol == "\"\"" {
		tokenDisplaySymbol = vaultFromGraph.Token.Symbol
	}

	vaultName, vaultDisplayName, vaultFormatedName := getVaultName(vaultFromGraph, vaultFromMeta, tokenFromMeta)
	vaultSymbol, vaultDisplaySymbol, vaultFormatedSymbol := getVaultSymbol(vaultFromGraph, vaultFromMeta, shareTokenFromMeta)

	// Get the migration informations for this vault. By default, migrationAvailable is false and
	// the migrationAddress matches the vault address. If a migration is available, point this last
	// one to MigrationTargetVault.
	migrationAddress := common.HexToAddress(vaultFromGraph.Id).String()
	migrationAvailable := vaultFromMeta.MigrationAvailable
	// migrationAvailable":true
	if vaultFromMeta.MigrationAvailable {
		migrationAddress = common.HexToAddress(vaultFromMeta.MigrationTargetVault).String()
	}

	// Get the price of the underlying asset. This is tricky because of the decimals. The prices are fetched
	// using the lens oracle daemon, stored in the TokenPrices map, and based on the USDC token, aka with 6
	// decimals. We first need to parse the BigInt Price to a float64, then divide by 10^6 to get the price
	// in an human readable USDC format.
	fPrice := new(big.Float)
	fPrice.SetUint64(prices[common.HexToAddress(vaultFromGraph.Token.Id)])
	humanizedPrice := new(big.Float).Quo(fPrice, big.NewFloat(math.Pow10(int(6))))
	fHumanizedPrice, _ := humanizedPrice.Float64()

	// Get the total assets locked in this vault. This is tricky because of the decimals. The total asset value
	// is a string which will be formated as a float64 and scaled with the underlying token decimals. With that
	// we should have a human readable Total Asset value, and we should be able to get the Total Value Locked
	// in the vault thanks to the above humanizedPrice value.
	fTotalAssets := new(big.Float)
	fTotalAssets.SetString(vaultFromGraph.BalanceTokens)
	humanizedTVL := new(big.Float).Quo(fTotalAssets, big.NewFloat(math.Pow10(int(vaultFromGraph.Token.Decimals))))
	fHumanizedTVLPrice, _ := big.NewFloat(0).Mul(humanizedTVL, humanizedPrice).Float64()

	// We can now extract the strategies from the graph based on the selected filter condition.
	strategies := prepareStrategies(strategiesCondition, vaultFromGraph, strategiesFromMeta)

	vault := &models.TVault{
		Inception:      activation,
		Address:        common.HexToAddress(vaultFromGraph.Id).String(),
		Symbol:         vaultSymbol,
		DisplaySymbol:  vaultDisplaySymbol,
		FormatedSymbol: vaultFormatedSymbol,
		Name:           vaultName,
		DisplayName:    vaultDisplayName,
		FormatedName:   vaultFormatedName,
		Icon:           utils.GITHUB_ICON_BASE_URL + chainIDAsString + `/` + common.HexToAddress(vaultFromGraph.Id).String() + `/logo-128.png`,
		Token: models.TToken{
			Address:     common.HexToAddress(vaultFromGraph.Token.Id).String(),
			Name:        vaultFromGraph.Token.Name,
			DisplayName: tokenDisplayName,
			Symbol:      tokenDisplaySymbol,
			Description: tokenFromMeta.Description,
			Decimals:    vaultFromGraph.Token.Decimals,
			Icon:        utils.GITHUB_ICON_BASE_URL + chainIDAsString + `/` + common.HexToAddress(vaultFromGraph.Token.Id).String() + `/logo-128.png`,
		},
		TVL: models.TTVL{
			TotalAssets: vaultFromGraph.BalanceTokens,
			TVL:         fHumanizedTVLPrice,
			Price:       fHumanizedPrice,
		},
		APY: models.TAPY{
			Type:     apyFromAPIV1.APY.Type,
			GrossAPR: apyFromAPIV1.APY.GrossAPR,
			NetAPY:   apyFromAPIV1.APY.NetAPY,
			Points: models.TAPYPoints{
				WeekAgo:   apyFromAPIV1.APY.Points.WeekAgo,
				MonthAgo:  apyFromAPIV1.APY.Points.MonthAgo,
				Inception: apyFromAPIV1.APY.Points.Inception,
			},
			Composite: models.TAPYComposite{
				Boost:      apyFromAPIV1.APY.Composite.Boost,
				PoolAPY:    apyFromAPIV1.APY.Composite.PoolAPY,
				BoostedAPR: apyFromAPIV1.APY.Composite.BoostedAPR,
				BaseAPR:    apyFromAPIV1.APY.Composite.BaseAPR,
				CvxAPR:     apyFromAPIV1.APY.Composite.CvxAPR,
				RewardsAPR: apyFromAPIV1.APY.Composite.RewardsAPR,
			},
			Fees: models.TAPYFees{
				Performance: float64(vaultFromGraph.PerformanceFeeBps) / 10000,
				Management:  float64(vaultFromGraph.ManagementFeeBps) / 10000,
				Withdrawal:  apyFromAPIV1.APY.Fees.Withdrawal,
				KeepCRV:     apyFromAPIV1.APY.Fees.KeepCRV,
				CvxKeepCRV:  apyFromAPIV1.APY.Fees.CvxKeepCRV,
			},
		},
		Strategies: strategies,
		Endorsed:   vaultFromGraph.Classification == "Endorsed",
		Version:    vaultFromGraph.ApiVersion,
		Decimals:   vaultFromGraph.ShareToken.Decimals,
		Type:       "v2", //No v1 in the subgraph
		// Emergency_shutdown: ,
		Updated: updated / 1000,
		Migration: models.TMigration{
			Available: migrationAvailable,
			Address:   migrationAddress,
		},
	}

	return vault
}
