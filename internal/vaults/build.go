package vaults

import (
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/common/addresses"
	"github.com/yearn/ydaemon/common/bigNumber"
	"github.com/yearn/ydaemon/common/helpers"
	"github.com/yearn/ydaemon/common/logs"
	"github.com/yearn/ydaemon/internal/meta"
	"github.com/yearn/ydaemon/internal/models"
	"github.com/yearn/ydaemon/internal/strategies"
)

func BuildNames(t models.TVault, metaVaultName string) models.TVault {
	name := strings.Replace(t.Name, "\"", "", -1)
	displayName := t.Name
	formatedName := t.Token.Name

	// If the meta file has a display name, use it
	if metaVaultName != "" {
		displayName = metaVaultName
	}
	// If the formated name is missing yVault suffix, add it
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
	if strings.HasSuffix(displayName, " Auto-Compounding") {
		displayName = strings.Replace(displayName, " Auto-Compounding", "", -1)
	}

	t.Name = strings.Replace(name, "-f-f", "-f", -1)
	t.DisplayName = strings.Replace(displayName, "-f-f", "-f", -1)
	t.FormatedName = strings.Replace(formatedName, "-f-f", "-f", -1)
	return t
}

func BuildSymbol(t models.TVault, metaVaultSymbol string) models.TVault {
	symbol := strings.Replace(t.Symbol, "\"", "", -1)
	formatedSymbol := t.Token.Symbol
	displaySymbol := metaVaultSymbol

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

	t.Symbol = strings.Replace(symbol, "-f-f", "-f", -1)
	t.DisplaySymbol = strings.Replace(displaySymbol, "-f-f", "-f", -1)
	t.FormatedSymbol = strings.Replace(formatedSymbol, "-f-f", "-f", -1)
	return t
}

func BuildMigration(t models.TVault) models.TMigration {
	migration := models.TMigration{
		Available: false,
		Address:   t.Address,
	}

	if vaultFromMeta, ok := meta.GetMetaVault(t.ChainID, t.Address); ok {
		migrationAddress := t.Address
		migrationContract := common.Address{}
		migrationAvailable := vaultFromMeta.MigrationAvailable
		if vaultFromMeta.MigrationAvailable {
			migrationAddress = vaultFromMeta.MigrationTargetVault
			migrationContract = vaultFromMeta.MigrationContract
		}
		migration = models.TMigration{
			Available: migrationAvailable,
			Address:   migrationAddress,
			Contract:  migrationContract,
		}
	}
	return migration
}

func BuildAPY(t models.TVault, aggregatedVault *models.TAggregatedVault, hasLegacyAPY bool) models.TAPY {
	apy := models.TAPY{}
	vaultFromMeta, okMeta := meta.GetMetaVault(t.ChainID, t.Address)

	if hasLegacyAPY {
		apy = models.TAPY{
			Type:              aggregatedVault.LegacyAPY.Type,
			GrossAPR:          aggregatedVault.LegacyAPY.GrossAPR,
			NetAPY:            aggregatedVault.LegacyAPY.NetAPY,
			StakingRewardsAPR: aggregatedVault.LegacyAPY.StakingRewardsAPR,
			Points: models.TAPYPoints{
				WeekAgo:   aggregatedVault.LegacyAPY.Points.WeekAgo,
				MonthAgo:  aggregatedVault.LegacyAPY.Points.MonthAgo,
				Inception: aggregatedVault.LegacyAPY.Points.Inception,
			},
			Composite: models.TAPYComposite{
				Boost:      aggregatedVault.LegacyAPY.Composite.Boost,
				PoolAPY:    aggregatedVault.LegacyAPY.Composite.PoolAPY,
				BoostedAPR: aggregatedVault.LegacyAPY.Composite.BoostedAPR,
				BaseAPR:    aggregatedVault.LegacyAPY.Composite.BaseAPR,
				CvxAPR:     aggregatedVault.LegacyAPY.Composite.CvxAPR,
				RewardsAPR: aggregatedVault.LegacyAPY.Composite.RewardsAPR,
			},
			Fees: models.TAPYFees{
				Performance: aggregatedVault.LegacyAPY.Fees.Performance,
				Management:  aggregatedVault.LegacyAPY.Fees.Management,
				Withdrawal:  aggregatedVault.LegacyAPY.Fees.Withdrawal,
				KeepCRV:     aggregatedVault.LegacyAPY.Fees.KeepCRV,
				CvxKeepCRV:  aggregatedVault.LegacyAPY.Fees.CvxKeepCRV,
			},
			Error: aggregatedVault.LegacyAPY.Error,
		}
		if okMeta && vaultFromMeta.APYTypeOverride != `` {
			apy.Type = vaultFromMeta.APYTypeOverride
		}
	} else if okMeta && vaultFromMeta.APYTypeOverride != `` {
		logs.Error(`Missing APY vault data for chainID: ` + strconv.FormatUint(t.ChainID, 10) + ` and address: ` + t.Address.Hex())
		apy.Type = vaultFromMeta.APYTypeOverride
		apy.Error = `Missing APY vault data`
	}
	return apy
}

func BuildTVL(t models.TVault) models.TTVL {
	humanizedPrice, fHumanizedPrice := getHumanizedTokenPrice(t.ChainID, t.Token.Address)
	fHumanizedTVLPrice := getHumanizedValue(t.TotalAssets, int(t.Decimals), humanizedPrice)
	strategiesList := strategies.ListStrategiesForVault(t.ChainID, t.Address)
	delegatedTokenAsBN := bigNumber.NewInt(0)
	fDelegatedValue := 0.0

	for _, strat := range strategiesList {
		delegatedValue := getHumanizedValue(strat.DelegatedAssets, int(t.Decimals), humanizedPrice)
		delegatedTokenAsBN = delegatedTokenAsBN.Add(delegatedTokenAsBN, strat.DelegatedAssets)
		fDelegatedValue += delegatedValue
	}

	tvl := models.TTVL{
		TotalAssets:          t.TotalAssets,
		TotalDelegatedAssets: delegatedTokenAsBN,
		TVL:                  fHumanizedTVLPrice - fDelegatedValue,
		TVLDeposited:         fHumanizedTVLPrice,
		TVLDelegated:         fDelegatedValue,
		Price:                fHumanizedPrice,
	}
	return tvl
}

func BuildCategory(t models.TVault) string {
	category := `Volatile`
	baseForStableCurrencies := []string{`USD`, `EUR`, `AUD`, `CHF`, `KRW`, `GBP`, `JPY`}
	baseForCurve := []string{`curve`, `crv`}
	baseForBalancer := []string{`balancer`, `bal`}
	baseForVelodrom := []string{`velodrome`, `velo`}
	allNames := []string{
		strings.ToLower(t.FormatedName),
		strings.ToLower(t.Name),
		strings.ToLower(t.DisplayName),
	}

	if vaultFromMeta, ok := meta.GetMetaVault(t.ChainID, t.Address); ok {
		//Using meta classification to set the category
		if vaultFromMeta.Classification.Stability == `Volatile` {
			category = `Volatile`
		} else {
			if helpers.Contains(baseForStableCurrencies, vaultFromMeta.Classification.StableBaseAsset) {
				category = `Stablecoin`
			} else {
				category = `Volatile`
			}
		}
		if helpers.Intersects(allNames, baseForCurve) {
			category = `Curve`
		}
		if helpers.Intersects(allNames, baseForBalancer) {
			category = `Balancer`
		}
		if helpers.Intersects(allNames, baseForVelodrom) {
			category = `Velodrome`
		}
	} else {
		//No meta, back to custom classification
		baseForBitcoin := []string{`btc`, `bitcoin`}
		baseForEth := []string{`eth`, `ethereum`}
		baseForStableCoins := []string{`dai`, `rai`, `mim`, `dola`}

		for _, stable := range baseForStableCurrencies {
			baseForStableCoins = append(baseForStableCoins, strings.ToLower(stable))
		}

		if helpers.Intersects(allNames, baseForBitcoin) {
			category = `Volatile`
		}
		if helpers.Intersects(allNames, baseForEth) {
			category = `Volatile`
		}
		if helpers.Intersects(allNames, baseForStableCoins) {
			category = `Stablecoin`
		}
		if helpers.Intersects(allNames, baseForCurve) {
			category = `Curve`
		}
		if helpers.Intersects(allNames, baseForBalancer) {
			category = `Balancer`
		}
		if helpers.Intersects(allNames, baseForVelodrom) {
			category = `Velodrome`
		}
	}
	return category
}

func BuildStaking(t models.TVault) models.TStaking {
	staking := models.TStaking{
		Available: false,
	}
	stakingData := strategies.GetStakingData(t.ChainID, t.Address)
	if !addresses.Equals(stakingData.Address, common.Address{}) {
		staking = stakingData
	}
	return staking
}
