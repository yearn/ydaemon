package fetcher

import (
	"strings"

	"github.com/yearn/ydaemon/common/addresses"
	"github.com/yearn/ydaemon/common/helpers"
	"github.com/yearn/ydaemon/internal/models"
	"github.com/yearn/ydaemon/internal/storage"
)

/**************************************************************************************************
** BuildVaultNames processes vault token information to generate three different name formats.
**
** This function handles the complex naming conventions for Yearn vaults by creating three
** distinct name formats:
**
** 1. name: The primary vault name (e.g., "USDC yVault")
**    - Uses the vault token name if available
**    - Falls back to displayName or formatedName if needed
**
** 2. displayName: A user-friendly name for UI display (e.g., "USDC")
**    - Uses provided vaultName parameter if available
**    - Removes "Auto-Compounding" suffixes for cleaner display
**    - Optimized for user readability without technical suffixes
**
** 3. formatedName: A standardized name with yVault suffix (e.g., "USDC yVault")
**    - Uses underlying token name with consistent "yVault" suffix
**    - Ensures all vault names follow the same pattern for standardization
**
** The function also performs cleanup operations like removing duplicate "-f" sequences
** and handling empty or missing names through fallback logic.
**
** @param t models.TVault - The vault model containing vault and token information
** @param vaultName string - Optional explicit vault name to use (often from metadata)
** @return string - The standard vault name
** @return string - The display-friendly vault name
** @return string - The formatted vault name with standardized suffix
**************************************************************************************************/
func BuildVaultNames(t models.TVault, vaultName string) (string, string, string) {
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

/**************************************************************************************************
** BuildVaultSymbol processes vault token information to generate three different symbol formats.
**
** This function creates standardized symbol formats for Yearn vaults:
**
** 1. symbol: The primary vault symbol (e.g., "yvUSDC")
**    - Uses the vault token symbol if available
**    - Falls back to displaySymbol or formatedSymbol if needed
**
** 2. displaySymbol: A cleaner symbol for UI display
**    - Uses the provided vaultSymbol parameter if available
**    - Maintains consistent format for user readability
**
** 3. formatedSymbol: A standardized symbol with "yv" prefix (e.g., "yvUSDC")
**    - Based on the underlying token symbol
**    - Ensures consistent "yv" prefix for all vault symbols
**
** The function also performs cleanup operations like removing duplicate "-f" sequences
** and handles empty or missing symbols through fallback logic using the helpers.SafeString
** function to choose the first non-empty value.
**
** @param t models.TVault - The vault model containing vault and token information
** @param vaultSymbol string - Optional explicit vault symbol to use (often from metadata)
** @return string - The standard vault symbol
** @return string - The display-friendly vault symbol
** @return string - The formatted vault symbol with standardized prefix
**************************************************************************************************/
func BuildVaultSymbol(t models.TVault, vaultSymbol string) (string, string, string) {
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

/**************************************************************************************************
** BuildVaultTVL calculates the Total Value Locked (TVL) for a vault.
**
** This function computes the financial metrics for a vault's TVL by:
**
** 1. Retrieving the vault token information from storage
** 2. Getting the humanized price of the underlying asset
** 3. Converting the raw total assets to a humanized USD value
**
** The function handles a special case for a specific vault address that was involved
** in a security incident, forcing its price to zero as the pool is frozen.
**
** The TVL structure returned contains:
** - TotalAssets: The raw amount of assets in the vault (in token base units)
** - TVL: The total value locked in USD
** - Price: The price of the underlying token in USD
**
** @param t models.TVault - The vault to calculate TVL for
** @return models.TTVL - A structure containing the TVL and related financial metrics
**************************************************************************************************/
func BuildVaultTVL(t models.TVault) models.TTVL {
	_, ok := storage.GetERC20(t.ChainID, t.Address)
	if !ok {
		return models.TTVL{}
	}
	_, fHumanizedPrice := getHumanizedTokenPrice(t.ChainID, t.AssetAddress)

	/**********************************************************************************************
	** Notice: The vault was implicated in a hack, so the price is now effectively 0 as the pool
	** is frozen.
	**********************************************************************************************/
	if addresses.Equals(t.Address, `0x718AbE90777F5B778B52D553a5aBaa148DD0dc5D`) {
		fHumanizedPrice = 0.0
	}


	kongTVL, ok := storage.GetKongTVL(t.ChainID, t.Address)
	if !ok {
		kongTVL = 0
	}

	tvl := models.TTVL{
		TotalAssets: t.LastTotalAssets,
		TVL:		float64(kongTVL),
		Price:       fHumanizedPrice,
	}
	return tvl
}

/**************************************************************************************************
** BuildVaultCategory determines the appropriate category for a vault based on multiple factors.
**
** This function implements a sophisticated categorization system for Yearn vaults by analyzing:
**
** 1. Explicit metadata category (highest priority if available)
** 2. Protocol-specific name patterns (Curve, Balancer, Velodrome, etc.)
** 3. Stability characteristics from metadata
** 4. Strategy name patterns (e.g., Prisma strategies)
** 5. Asset name patterns (BTC, ETH, stablecoins)
**
** The function uses a multi-layered approach that first checks for explicit category assignments,
** then searches for protocol indicators in the vault and token names, then considers stability
** metrics, and finally examines strategy-specific indicators.
**
** Categories include:
** - Protocol-specific: "Curve", "Balancer", "Velodrome", "Aerodrome", "Gamma", "Pendle", "Prisma"
** - Asset-specific: "Stablecoin", "Volatile"
**
** @param t models.TVault - The vault model to categorize
** @param strategies map[string]models.TStrategy - Map of strategies associated with the vault
** @return string - The determined category for the vault
**************************************************************************************************/
func BuildVaultCategory(t models.TVault, strategies map[string]models.TStrategy) string {
	category := ``
	baseForStableCurrencies := []string{`USD`, `EUR`, `AUD`, `CHF`, `KRW`, `GBP`, `JPY`}
	baseForCurve := []string{`curve`, `crv`}
	baseForBalancer := []string{`balancer`, `bal`}
	baseForVelodrome := []string{`velodrome`, `velo`}
	baseForAerodrome := []string{`aerodrome`, `aero`}
	baseForBitcoin := []string{`btc`, `bitcoin`}
	baseForEth := []string{`eth`, `ethereum`}
	baseForStableCoins := []string{`dai`, `rai`, `mim`, `dola`}
	baseForPrisma := []string{`prisma`}
	baseForGamma := []string{`gamma`}
	baseForPendle := []string{`pendle`}
	name, displayName, formatedName := BuildVaultNames(t, ``)
	symbol, displaySymbol, formatedSymbol := BuildVaultSymbol(t, ``)
	allNames := []string{
		strings.ToLower(name),
		strings.ToLower(displayName),
		strings.ToLower(formatedName),
		strings.ToLower(symbol),
		strings.ToLower(displaySymbol),
		strings.ToLower(formatedSymbol),
	}

	if t.Metadata.Category != `` && t.Metadata.Category != models.VaultCategoryAutomatic {
		return string(t.Metadata.Category)
	}

	//Using meta stability to set the category
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
	if helpers.Intersects(allNames, baseForGamma) {
		category = `Gamma`
	}
	if helpers.Intersects(allNames, baseForPendle) {
		category = `Pendle`
	}

	if len(strategies) > 0 {
		for _, strategy := range strategies {
			allStratNames := []string{
				strings.ToLower(strategy.Name),
				strings.ToLower(strategy.DisplayName),
			}
			if helpers.Intersects(allStratNames, baseForPrisma) {
				category = `Prisma`
				break
			}
		}
	}
	if t.Metadata.Stability.Stability == models.VaultStabilityVolatile {
		category = `Volatile`
	} else {
		if helpers.Contains(baseForStableCurrencies, t.Metadata.Stability.StableBaseAsset) {
			category = `Stablecoin`
		}
	}

	if category == `` {
		for _, stable := range baseForStableCurrencies {
			baseForStableCoins = append(baseForStableCoins, strings.ToLower(stable))
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
		if helpers.Intersects(allNames, baseForGamma) {
			category = `Gamma`
		}
		if helpers.Intersects(allNames, baseForPendle) {
			category = `Pendle`
		}
	}

	if category == `` {
		if helpers.Intersects(allNames, baseForBitcoin) || helpers.Intersects(allNames, baseForEth) {
			category = `Volatile`
		} else if helpers.Intersects(allNames, baseForStableCoins) {
			category = `Stablecoin`
		} else {
			category = `Volatile`
		}
	}
	return category
}
