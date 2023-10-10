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
	"github.com/yearn/ydaemon/internal/storage"
	"github.com/yearn/ydaemon/internal/strategies"
)

func BuildNames(t models.TVault, vaultName string) (string, string, string) {
	name := ``
	vaultToken, ok := storage.GetERC20(t.ChainID, t.Address)
	if ok {
		name = strings.Replace(vaultToken.Name, "\"", "", -1)
	}

	displayName := name
	formatedName := name
	underlyingToken, ok := storage.GetERC20(t.ChainID, t.AssetAddress)
	if ok {
		formatedName = underlyingToken.Name
	}

	// If the meta file has a display name, use it
	if vaultName != "" {
		displayName = vaultName
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

	name = strings.Replace(name, "-f-f", "-f", -1)
	displayName = strings.Replace(displayName, "-f-f", "-f", -1)
	formatedName = strings.Replace(formatedName, "-f-f", "-f", -1)
	return name, displayName, formatedName
}

func BuildSymbol(t models.TVault, vaultSymbol string) (string, string, string) {
	symbol := ``
	vaultToken, ok := storage.GetERC20(t.ChainID, t.Address)
	if ok {
		symbol = strings.Replace(vaultToken.Symbol, "\"", "", -1)
	}

	displaySymbol := vaultSymbol
	formatedSymbol := vaultSymbol
	underlyingToken, ok := storage.GetERC20(t.ChainID, t.AssetAddress)
	if ok {
		formatedSymbol = underlyingToken.Symbol
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
	symbol = strings.Replace(symbol, "-f-f", "-f", -1)
	displaySymbol = strings.Replace(displaySymbol, "-f-f", "-f", -1)
	formatedSymbol = strings.Replace(formatedSymbol, "-f-f", "-f", -1)
	return symbol, displaySymbol, formatedSymbol
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
	vaultFromMeta, okMeta := meta.GetMetaVault(t.ChainID, t.Address)
	apy := models.TAPY{}
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
				KeepVelo:    aggregatedVault.LegacyAPY.Fees.KeepVelo,
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
	vaultToken, ok := storage.GetERC20(t.ChainID, t.Address)
	if !ok {
		return models.TTVL{}
	}
	humanizedPrice, fHumanizedPrice := getHumanizedTokenPrice(t.ChainID, t.AssetAddress)

	/**********************************************************************************************
	** Notice: The vault was implicated in a hack, so the price is now effectively 0 as the pool
	** is frozen.
	**********************************************************************************************/
	if addresses.Equals(t.Address, `0x718AbE90777F5B778B52D553a5aBaa148DD0dc5D`) {
		humanizedPrice = bigNumber.NewFloat(0)
		fHumanizedPrice = 0.0
	}

	fHumanizedTVLPrice := getHumanizedValue(t.LastTotalAssets, int(vaultToken.Decimals), humanizedPrice)
	strategiesList := strategies.ListStrategiesForVault(t.ChainID, t.Address)
	delegatedTokenAsBN := bigNumber.NewInt(0)
	fDelegatedValue := 0.0

	for _, strat := range strategiesList {
		delegatedValue := getHumanizedValue(strat.DelegatedAssets, int(vaultToken.Decimals), humanizedPrice)
		delegatedTokenAsBN = delegatedTokenAsBN.Add(delegatedTokenAsBN, strat.DelegatedAssets)
		fDelegatedValue += delegatedValue
	}

	tvl := models.TTVL{
		TotalAssets:          t.LastTotalAssets,
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
	baseForVelodrome := []string{`velodrome`, `velo`}
	baseForAerodrome := []string{`aerodrome`, `aero`}
	name, displayName, formatedName := BuildNames(t, ``)
	allNames := []string{
		strings.ToLower(name),
		strings.ToLower(displayName),
		strings.ToLower(formatedName),
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
		if helpers.Intersects(allNames, baseForVelodrome) {
			category = `Velodrome`
		}
		if helpers.Intersects(allNames, baseForAerodrome) {
			category = `Aerodrome`
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
		if helpers.Intersects(allNames, baseForVelodrome) {
			category = `Velodrome`
		}
		if helpers.Intersects(allNames, baseForAerodrome) {
			category = `Aerodrome`
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
