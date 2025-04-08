package vaults

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/yearn/ydaemon/common/env"
	"github.com/yearn/ydaemon/common/helpers"
	"github.com/yearn/ydaemon/common/logs"
	"github.com/yearn/ydaemon/common/sort"
	"github.com/yearn/ydaemon/internal/models"
	"github.com/yearn/ydaemon/internal/storage"
)

/**************************************************************************************************
** getVaults retrieves a filtered list of vaults based on query parameters and a filter function.
**
** This is the core function that powers most vault-related API endpoints. It retrieves vaults
** from storage, applies filtering based on the provided filter function and query parameters,
** and returns a list of simplified external vault structures.
**
** The function handles various query parameters:
** - orderBy: Field to sort the results by (default: 'featuringScore')
** - orderDirection: Sort direction, 'asc' or 'desc' (default: 'asc')
** - strategiesCondition: Condition for including strategies in results (default: 'debtRatio')
** - hideAlways: Whether to hide certain vaults (default: false)
** - migrable: Condition for including migrable vaults (default: 'none')
** - page/limit: Pagination controls (defaults: page 1, limit 200)
** - chainIDs: Comma-separated list of chain IDs to include (default: all supported chains)
**
** @param c *gin.Context - The Gin context containing the HTTP request
** @param filterFunc func(vault models.TVault) bool - Function that determines if a vault should be included
** @return []TSimplifiedExternalVault - The filtered and sorted list of vaults
** @return error - Any error encountered during processing
**************************************************************************************************/
func getVaults(
	c *gin.Context,
	filterFunc func(vault models.TVault) bool,
) ([]TSimplifiedExternalVault, error) {
	/** 🔵 - Yearn *************************************************************************************
	** This function takes in a context from the gin framework. The context contains information
	** about the environment and request. From the context, the function extracts the following
	** parameters:
	**
	** orderBy: A string that determines the order in which the vaults are returned. It is obtained
	** from the 'orderBy' query parameter in the request. If the parameter is not provided, it
	** defaults to 'featuringScore'.
	**
	** orderDirection: A string that determines the direction in which the vaults are ordered. It is
	** obtained from the 'orderDirection' query parameter in the request. If the parameter is not
	** provided, it defaults to 'asc'.
	**
	** hideAlways: A boolean value that determines whether to hide certain vaults. It is obtained
	** from the 'hideAlways' query parameter in the request. If the parameter is not provided,
	** it defaults to 'false'.
	**
	** strategiesCondition: A string that determines the condition for selecting strategies. It is
	** obtained from the 'strategiesCondition' query parameter in the request.
	**************************************************************************************************/
	orderBy := helpers.SafeString(getQueryParam(c, `orderBy`), `featuringScore`)
	orderDirection := helpers.SafeString(getQueryParam(c, `orderDirection`), `asc`)
	hideAlways := helpers.StringToBool(getQueryParam(c, `hideAlways`))
	stratCon := validateStrategyCondition(c, "strategiesCondition")

	/** 🔵 - Yearn *************************************************************************************
	** migrable: A string that determines the condition for selecting migrable vaults. It is
	** obtained from the 'migrable' query parameter in the request.
	**************************************************************************************************/
	migrable := validateMigrableCondition(c, `migrable`)
	if migrable != `none` && hideAlways {
		handleError(c, fmt.Errorf("migrable and hideAlways cannot be true at the same time"),
			http.StatusBadRequest, "Invalid parameter combination", "GetVaults")
		return []TSimplifiedExternalVault{}, fmt.Errorf("migrable and hideAlways cannot be true at the same time")
	}

	/** 🔵 - Yearn *************************************************************************************
	** page: A uint64 value that represents the page number for pagination. It is obtained from the
	** 'page' query parameter in the request. If the parameter is not provided, it defaults to 1.
	**
	** limit: A uint64 value that represents the number of vaults to be returned per page. It is
	** obtained from the 'limit' query parameter in the request. If the parameter is not provided,
	** it defaults to 200.
	**************************************************************************************************/
	page := validateNumericQuery(c, "page", DEFAULT_PAGE_NUMBER, DEFAULT_PAGE_NUMBER, MAX_PAGE_LIMIT, "GetVaults")
	limit := validateNumericQuery(c, "limit", DEFAULT_PAGE_LIMIT, DEFAULT_PAGE_NUMBER, MAX_PAGE_LIMIT, "GetVaults")

	/** 🔵 - Yearn *************************************************************************************
	** chainsStr: A string that represents the chain IDs for which the vaults are to be returned. It is
	** obtained from the 'chainIDs' query parameter in the request. The string is split by commas to
	** obtain an array of chain IDs.
	**
	** chains: An array of uint64 values that represents the chain IDs for which the vaults are to be
	** returned. If the 'chains' query parameter is not provided or is empty, this array defaults to
	** all supported chain IDs.
	**
	** The 'chains' array is populated by iterating over the 'chainsStr' array and converting each
	** chain ID to a uint64 value. If the conversion is not successful, the chain ID is ignored.
	**
	** The 'chains' array is used to filter the vaults that are returned in the response.
	**************************************************************************************************/
	chainsStr := strings.Split(getQueryParam(c, `chainIDs`), `,`)
	chains := []uint64{}
	if len(chainsStr) == 0 || (len(chainsStr) == 1 && chainsStr[0] == ``) {
		chains = env.SUPPORTED_CHAIN_IDS
	} else {
		for _, chainStr := range chainsStr {
			chain, ok := helpers.AssertChainID(chainStr)
			if !ok {
				continue
			}
			chains = append(chains, chain)
		}
	}

	/** 🔵 - Yearn *************************************************************************************
	** The following code processes vaults across all specified chains and applies filtering.
	** It retrieves vaults for each chain, applies the filter function, and processes valid vaults
	** into a simplified format. It also applies additional filters based on parameters like
	** hideAlways and migrable.
	**
	** Optimizations include:
	** 1. Early validation of chain configuration
	** 2. Pre-filtering vaults before expensive operations
	** 3. Skipping unnecessary strategy lookups when possible
	** 4. Using a capacity hint for the allVaults slice to reduce reallocations
	**************************************************************************************************/
	// Provide capacity hint based on typical vault counts to reduce reallocations
	estimatedVaultCount := len(chains) * 300 // Rough estimate of 300 vaults per chain
	allVaults := make([]TSimplifiedExternalVault, 0, estimatedVaultCount)

	for _, chainID := range chains {
		// Get chain configuration early to validate
		chain, ok := env.GetChain(chainID)
		if !ok {
			logs.Error(fmt.Errorf("chain configuration not found for chainID %d", chainID),
				http.StatusInternalServerError, "Internal configuration error", "GetVaults")
			continue
		}

		// Retrieve vaults for this chain
		_, vaultsSlice := storage.ListVaults(chainID)
		if len(vaultsSlice) == 0 {
			// Skip empty chains
			continue
		}

		// Process each vault with optimized filtering
		for _, currentVault := range vaultsSlice {
			// Apply early filters to avoid unnecessary processing
			if helpers.Contains(chain.BlacklistedVaults, currentVault.Address) {
				continue
			}

			// Apply custom filter function
			if !filterFunc(currentVault) {
				continue
			}

			// Skip retired vaults when hideAlways is true
			if migrable == `none` && currentVault.Metadata.IsRetired && hideAlways {
				continue
			}

			// Process vault data - use the optimized CreateExternalVault function
			newVault, err := CreateExternalVault(currentVault)
			if err != nil {
				logs.Error(fmt.Errorf("failed to process vault %s on chain %d: %s",
					currentVault.Address.Hex(), chainID, err), http.StatusInternalServerError,
					"Error processing vault data", "GetVaults")
				continue
			}

			// Calculate APR and featuring score
			APRAsFloat := 0.0
			if newVault.APR.NetAPR != nil {
				APRAsFloat, _ = newVault.APR.NetAPR.Float64()
			}

			newVault.FeaturingScore = newVault.TVL.TVL * APRAsFloat
			if newVault.Details.IsHighlighted {
				newVault.FeaturingScore = newVault.FeaturingScore * HIGHLIGHTING_MULTIPLIER
			}

			// Apply migrable-specific filters
			if !currentVault.Metadata.IsRetired {
				if migrable == MIGRABLE_CONDITION_NO_DUST && (newVault.TVL.TVL < MIN_DUST_TVL || !newVault.Migration.Available) {
					continue
				} else if migrable == MIGRABLE_CONDITION_ALL && !newVault.Migration.Available {
					continue
				} else if migrable == MIGRABLE_CONDITION_IGNORE && (newVault.Migration.Available || newVault.Details.IsHidden) {
					continue
				}
			}

			vaultStrategies, _ := storage.ListStrategiesForVault(chainID, currentVault.Address)
			newVault.Strategies = []TExternalStrategy{}
			for _, strategy := range vaultStrategies {
				strategyWithDetails := CreateExternalStrategy(strategy)
				if !strategyWithDetails.ShouldBeIncluded(stratCon) {
					continue
				}

				newVault.Strategies = append(newVault.Strategies, strategyWithDetails)
			}

			// Convert directly to simplified format
			simplified := toSimplifiedVersion(newVault, models.TStrategy{})
			simplified.Description = newVault.Description
			allVaults = append(allVaults, simplified)
		}
	}

	/** 🔵 - Yearn *************************************************************************************
	** Finally, all vaults in 'allVaults' are sorted based on 'orderBy' and 'orderDirection'.
	**
	** The sorted vaults are then paginated based on 'page' and 'limit', and the paginated data is
	** returned in the response.
	**************************************************************************************************/
	sort.SortBy(orderBy, orderDirection, allVaults)
	start := (page - 1) * limit
	end := page * limit
	if start > uint64(len(allVaults)) {
		// If the start index is beyond the total number of vaults, return an empty array
		return []TSimplifiedExternalVault{}, nil
	}
	if end > uint64(len(allVaults)) {
		end = uint64(len(allVaults))
	}
	data := allVaults[start:end]

	return data, nil
}

/**************************************************************************************************
** strategyDataIsRequired determines if strategy data is needed based on migrable condition.
**
** This helper function checks if the vault strategies need to be loaded based on the migrable
** parameter. This helps avoid unnecessary strategy lookups for certain filter conditions.
**
** @param migrable string - The migrable condition being used
** @return bool - True if strategy data is required for the given condition
**************************************************************************************************/
func strategyDataIsRequired(migrable string) bool {
	// For most migrable conditions, we need strategy data
	return migrable != "none" && migrable != "ignore"
}

/**************************************************************************************************
** Helper functions for vault version checks were consolidated in validation.go.
**
** To maintain compatibility with existing code, these forwarding functions are kept
** but delegate to the centralized implementation.
**************************************************************************************************/

/**************************************************************************************************
** isV3Vault determines if a vault is a v3 vault based on its version and kind.
**
** @param vault models.TVault - The vault to check
** @return bool - True if the vault is a v3 vault, false otherwise
**************************************************************************************************/
func isV3Vault(vault models.TVault) bool {
	return VaultVersionChecks.IsV3(vault)
}

/**************************************************************************************************
** isV2Vault determines if a vault is a v2 vault based on its version.
**
** @param vault models.TVault - The vault to check
** @return bool - True if the vault is a v2 vault, false otherwise
**************************************************************************************************/
func isV2Vault(vault models.TVault) bool {
	return VaultVersionChecks.IsV2(vault)
}
