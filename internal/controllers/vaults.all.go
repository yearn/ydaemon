package controllers

import (
	"math"
	"math/big"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/majorfi/ydaemon/internal/logs"
	"github.com/majorfi/ydaemon/internal/models"
	"github.com/majorfi/ydaemon/internal/store"
	"github.com/majorfi/ydaemon/internal/utils"
)

func buildVaultName(
	chainID uint64,
	vaultAddress common.Address,
	vaultName string,
	tokenName string,
) (name string, displayName string, formatedName string) {
	name = strings.Replace(vaultName, "\"", "", -1)
	formatedName = tokenName
	vaultFromMeta, ok := store.VaultsFromMeta[chainID][vaultAddress.String()]
	if ok {
		displayName = strings.Replace(vaultFromMeta.DisplayName, "\"", "", -1)
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
	shareTokenFromMeta, ok := store.TokensFromMeta[chainID][tokenAddress.String()]
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

// Get the price of the underlying asset. This is tricky because of the decimals. The prices are fetched
// using the lens oracle daemon, stored in the TokenPrices map, and based on the USDC token, aka with 6
// decimals. We first need to parse the BigInt Price to a float64, then divide by 10^6 to get the price
// in an human readable USDC format.
func buildTokenPrice(chainID uint64, tokenAddress common.Address) (*big.Float, float64) {
	prices := store.TokenPrices[chainID]
	fPrice := new(big.Float)
	price, ok := prices[tokenAddress]
	if ok {
		fPrice.SetInt(price)
		humanizedPrice := new(big.Float).Quo(fPrice, big.NewFloat(math.Pow10(int(6))))
		fHumanizedPrice, _ := humanizedPrice.Float64()
		return humanizedPrice, fHumanizedPrice
	}
	return big.NewFloat(0), 0.0
}

// Get the total assets locked in this vault. This is tricky because of the decimals. The total asset value
// is a string which will be formated as a float64 and scaled with the underlying token decimals. With that
// we should have a human readable Total Asset value, and we should be able to get the Total Value Locked
// in the vault thanks to the above humanizedPrice value.
func buildTVL(balanceToken string, decimals int, humanizedPrice *big.Float) float64 {
	fTotalAssets := new(big.Float)
	fTotalAssets.SetString(balanceToken)
	humanizedTVL := new(big.Float).Quo(fTotalAssets, big.NewFloat(math.Pow10(decimals)))
	fHumanizedTVLPrice, _ := big.NewFloat(0).Mul(humanizedTVL, humanizedPrice).Float64()
	return fHumanizedTVLPrice
}

// From the legacy API, build the schema for the APY, models.TAPY, used to get the details and the
// breakdown of the vault.
func buildAPY(chainID uint64, vaultAddress common.Address, perfFee, manaFee uint64) models.TAPY {
	apy := models.TAPY{}
	apyFromAPIV1, ok := store.VaultsFromAPIV1[chainID][vaultAddress.String()]

	if ok {
		apy = models.TAPY{
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
				Performance: float64(perfFee) / 10000,
				Management:  float64(manaFee) / 10000,
				Withdrawal:  apyFromAPIV1.APY.Fees.Withdrawal,
				KeepCRV:     apyFromAPIV1.APY.Fees.KeepCRV,
				CvxKeepCRV:  apyFromAPIV1.APY.Fees.CvxKeepCRV,
			},
		}
	}
	return apy
}

// Get the migration informations for this vault. By default, migrationAvailable is false and
// the migrationAddress matches the vault address. If a migration is available, point this last
// one to MigrationTargetVault.
func buildMigration(chainID uint64, vaultAddress common.Address) models.TMigration {
	migration := models.TMigration{}
	vaultFromMeta, ok := store.VaultsFromMeta[chainID][vaultAddress.String()]

	if ok {
		migrationAddress := vaultAddress.String()
		migrationAvailable := vaultFromMeta.MigrationAvailable
		if vaultFromMeta.MigrationAvailable {
			migrationAddress = common.HexToAddress(vaultFromMeta.MigrationTargetVault).String()
		}

		migration = models.TMigration{
			Available: migrationAvailable,
			Address:   migrationAddress,
		}
	}
	return migration
}

// We can now extract the strategies from the graph based on the selected filter condition.
func buildStrategies(
	chainID uint64,
	withStrategiesDetails bool,
	strategiesCondition string,
	vaultFromGraph models.TVaultFromGraph,
) []models.TStrategy {
	strategies := []models.TStrategy{}
	strategiesFromMeta := store.StrategiesFromMeta[chainID]
	strategiesFromMulticall := store.StrategyMultiCallData[chainID]

	for _, strategy := range vaultFromGraph.Strategies {
		multicallData, ok := strategiesFromMulticall[common.HexToAddress(strategy.Address)]
		if !ok {
			multicallData = models.TStrategyMultiCallData{}
		}

		currentStrategy := models.TStrategy{
			Address:     common.HexToAddress(strategy.Address).String(),
			Name:        strategy.Name,
			Description: strategiesFromMeta[common.HexToAddress(strategy.Address).String()].Description,
		}
		debtLimit := bValueWithFallbackUint64(multicallData.DebtLimit, 0)

		if withStrategiesDetails {
			currentStrategy.Details = &models.TStrategyDetails{}
			currentStrategy.Details.Keeper = common.HexToAddress(strategy.Keeper).String()
			currentStrategy.Details.Strategist = common.HexToAddress(strategy.Strategist).String()
			currentStrategy.Details.Rewards = common.HexToAddress(strategy.Rewards).String()
			currentStrategy.Details.HealthCheck = common.HexToAddress(strategy.HealthCheck).String()
			currentStrategy.Details.Version = strategy.ApiVersion
			currentStrategy.Details.InQueue = strategy.InQueue
			currentStrategy.Details.DoHealthCheck = strategy.DoHealthCheck
			currentStrategy.Details.EmergencyExit = strategy.EmergencyExit
			currentStrategy.Details.DebtLimit = debtLimit
			currentStrategy.Details.CreditAvailable = bValueWithFallbackString(multicallData.CreditAvailable, `0`)
			currentStrategy.Details.DebtOutstanding = bValueWithFallbackString(multicallData.DebtOutstanding, `0`)
			currentStrategy.Details.ExpectedReturn = bValueWithFallbackString(multicallData.ExpectedReturn, `0`)
			currentStrategy.Details.PerformanceFee = bValueWithFallbackUint64(multicallData.PerformanceFee, 0)
			currentStrategy.Details.Activation = bValueWithFallbackUint64(multicallData.Activation, 0)
			currentStrategy.Details.DebtRatio = bValueWithFallbackUint64(multicallData.DebtRatio, 0)
			currentStrategy.Details.RateLimit = bValueWithFallbackString(multicallData.RateLimit, `0`)
			currentStrategy.Details.MinDebtPerHarvest = bValueWithFallbackString(multicallData.MinDebtPerHarvest, `0`)
			currentStrategy.Details.MaxDebtPerHarvest = bValueWithFallbackString(multicallData.MaxDebtPerHarvest, `0`)
			currentStrategy.Details.EstimatedTotalAssets = bValueWithFallbackString(multicallData.EstimatedTotalAssets, `0`)
			currentStrategy.Details.KeepCRV = bValueWithFallbackUint64(multicallData.KeepCRV, 0)
			currentStrategy.Details.LastReport = bValueWithFallbackUint64(multicallData.LastReport, 0)
			currentStrategy.Details.TotalDebt = bValueWithFallbackString(multicallData.TotalDebt, `0`)
			currentStrategy.Details.TotalGain = bValueWithFallbackString(multicallData.TotalGain, `0`)
			currentStrategy.Details.TotalLoss = bValueWithFallbackString(multicallData.TotalLoss, `0`)
			currentStrategy.Details.APR = 0.0
			currentStrategy.Details.IsActive = multicallData.IsActive

			if len(strategy.Reports) > 0 {
				var totalAPR float64
				for _, report := range strategy.Reports {
					if len(report.Results) > 0 {
						floatAPR, err := strconv.ParseFloat(report.Results[0].APR, 64)
						if err != nil {
							logs.Warning(err)
							floatAPR = 0
						}
						totalAPR += floatAPR
					}
				}
				floatTotalAPR := (totalAPR / float64(len(strategy.Reports))) * 100
				currentStrategy.Details.APR = floatTotalAPR
			}
		}

		if strategiesCondition == `inQueue` && strategy.InQueue {
			strategies = append(strategies, currentStrategy)
		}
		if strategiesCondition == `debtLimit` && debtLimit == 0 {
			strategies = append(strategies, currentStrategy)
		}
	}
	return strategies
}

func prepareVaultSchema(
	chainID uint64,
	strategiesCondition string,
	withStrategiesDetails bool,
	vaultFromGraph models.TVaultFromGraph,
) *models.TVault {
	chainIDAsString := strconv.FormatUint(chainID, 10)
	vaultAddress := common.HexToAddress(vaultFromGraph.Id)
	tokenAddress := common.HexToAddress(vaultFromGraph.Token.Id)
	tokenFromMeta := store.TokensFromMeta[chainID][tokenAddress.String()]
	updated := strToUint(vaultFromGraph.LatestUpdate.Timestamp, 0)
	activation := strToUint(vaultFromGraph.Activation, 0)
	tokenDisplayName := valueWithFallback(tokenFromMeta.Name, vaultFromGraph.Token.Name)
	tokenDisplaySymbol := valueWithFallback(tokenFromMeta.Symbol, vaultFromGraph.Token.Symbol)

	vaultName, vaultDisplayName, vaultFormatedName := buildVaultName(
		chainID,
		vaultAddress,
		vaultFromGraph.ShareToken.Name,
		vaultFromGraph.Token.Name,
	)
	vaultSymbol, vaultDisplaySymbol, vaultFormatedSymbol := buildVaultSymbol(
		chainID,
		common.HexToAddress(vaultFromGraph.ShareToken.Id),
		vaultFromGraph.ShareToken.Symbol,
		vaultFromGraph.Token.Symbol,
	)
	humanizedPrice, fHumanizedPrice := buildTokenPrice(
		chainID,
		tokenAddress,
	)
	fHumanizedTVLPrice := buildTVL(
		vaultFromGraph.BalanceTokens,
		int(vaultFromGraph.Token.Decimals),
		humanizedPrice,
	)

	strategies := buildStrategies(
		chainID,
		withStrategiesDetails,
		strategiesCondition,
		vaultFromGraph,
	)

	vault := &models.TVault{
		Inception:      activation,
		Address:        vaultAddress.String(),
		Symbol:         vaultSymbol,
		DisplaySymbol:  vaultDisplaySymbol,
		FormatedSymbol: vaultFormatedSymbol,
		Name:           vaultName,
		DisplayName:    vaultDisplayName,
		FormatedName:   vaultFormatedName,
		Icon:           utils.GITHUB_ICON_BASE_URL + chainIDAsString + `/` + vaultAddress.String() + `/logo-128.png`,
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
		APY:        buildAPY(chainID, vaultAddress, vaultFromGraph.PerformanceFeeBps, vaultFromGraph.ManagementFeeBps),
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
