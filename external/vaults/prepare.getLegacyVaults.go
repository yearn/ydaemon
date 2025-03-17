package vaults

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yearn/ydaemon/common/env"
	"github.com/yearn/ydaemon/common/helpers"
	"github.com/yearn/ydaemon/common/sort"
	"github.com/yearn/ydaemon/internal/models"
	"github.com/yearn/ydaemon/internal/storage"
)

/**************************************************************************************************
** getLegacyVaults retrieves a filtered list of vaults in the legacy TExternalVault format.
**
** This function is the core implementation for all legacy vault endpoints. It retrieves vaults
** from storage, applies filtering based on the provided filter function and query parameters,
** and returns them in the original TExternalVault format (as opposed to the newer
** TSimplifiedExternalVault format).
**
** The function supports comprehensive filtering and sorting capabilities through query parameters:
** - page & limit: For pagination control (default: page 1, limit 200)
** - hideAlways: Whether to hide vaults marked as hidden or retired (default: false)
** - orderBy: Field to sort results by (default: 'featuringScore')
** - orderDirection: Sort direction, 'asc' or 'desc' (default: 'asc')
** - strategiesCondition: Filter for strategies (default: 'debtRatio')
** - migrable: Filter for migrable vaults, values: 'none', 'all', 'nodust', 'ignore' (default: 'none')
**
** The function follows these processing steps:
** 1. Extracts and validates query parameters for pagination, filtering, and sorting
** 2. Retrieves all vaults for the specified chain
** 3. Applies the provided filter function to include only matching vaults
** 4. Filters out blacklisted vaults
** 5. Applies additional filters based on migration status and hide settings
** 6. Retrieves and filters strategies for each vault
** 7. Sorts and paginates the final results
**
** @param c *gin.Context - The Gin context containing the HTTP request
** @param filterFunc func(vault models.TVault) bool - Function to filter vaults
** @return []TExternalVault - The filtered, sorted, and paginated list of vaults
**************************************************************************************************/
func getLegacyVaults(
	c *gin.Context,
	filterFunc func(vault models.TVault) bool,
) []TExternalVault {
	/** 🔵 - Yearn *************************************************************************************
	** Validate and process pagination parameters
	**
	** page: A uint64 value that represents the page number for pagination. It is obtained from the
	** 'page' query parameter in the request. If the parameter is not provided, it defaults to 1.
	**
	** limit: A uint64 value that represents the number of vaults to be returned per page. It is
	** obtained from the 'limit' query parameter in the request. If the parameter is not provided,
	** it defaults to 200.
	**************************************************************************************************/
	page := validateNumericQuery(c, "page", 1, 1, 1000, "GetLegacyVaults")
	limit := validateNumericQuery(c, "limit", 200, 1, 1000, "GetLegacyVaults")

	/** 🔵 - Yearn *************************************************************************************
	** This function takes in a context from the gin framework. The context contains information
	** about the environment and request. From the context, the function extracts the following
	** parameters:
	**
	** hideAlways: A boolean value that determines whether to hide certain vaults. It is obtained
	** from the 'hideAlways' query parameter in the request. If the parameter is not provided,
	** it defaults to 'false'.
	**
	** orderBy: A string that determines the order in which the vaults are returned. It is obtained
	** from the 'orderBy' query parameter in the request. If the parameter is not provided,
	** it defaults to 'details.order'.
	**************************************************************************************************/
	hideAlways := validateBooleanQuery(c, "hideAlways", false, "GetLegacyVaults")

	// Define valid order fields
	validOrderFields := []string{
		"featuringScore", "name", "symbol", "decimals", "type", "tvl",
		"apr.net_apy", "apr.composite.boost", "apr.points.week_ago", "apr.type", "details.order",
	}
	orderBy := validateStringChoiceQuery(c, "orderBy", "featuringScore", validOrderFields, "GetLegacyVaults")

	// Validate order direction
	validDirections := []string{"asc", "desc"}
	orderDirection := validateStringChoiceQuery(c, "orderDirection", "asc", validDirections, "GetLegacyVaults")

	/** 🔵 - Yearn *************************************************************************************
	** strategiesCondition: A string that determines the condition for selecting strategies. It is
	** obtained from the 'strategiesCondition' query parameter in the request.
	**
	** migrable: A string that determines the condition for selecting migrable vaults. It is
	** obtained from the 'migrable' query parameter in the request.
	**
	** chainID: A string that represents the chain ID of the vaults to be returned. It is obtained
	** from the 'chainID' path parameter in the request.
	**************************************************************************************************/
	strategiesCondition := validateStrategyCondition(c, "strategiesCondition")
	migrable := validateMigrableCondition(c, "migrable")

	// Validate chain ID using the utility function
	chainID, ok := validateChainID(c, `chainID`)
	if !ok {
		return nil
	}

	if migrable != `none` && hideAlways {
		handleError(c, fmt.Errorf("migrable and hideAlways cannot be true at the same time"),
			http.StatusBadRequest, "Invalid parameter combination", "getLegacyVaults")
		return nil
	}

	// Get chain configuration early to avoid repeated lookups
	chain, ok := env.GetChain(chainID)
	if !ok {
		handleError(c, fmt.Errorf("chain configuration not found for chainID %d", chainID),
			http.StatusInternalServerError, "Internal configuration error", "getLegacyVaults")
		return nil
	}

	// Get all vaults and pre-estimate the capacity
	allVaults, _ := storage.ListVaults(chainID)
	estimatedCapacity := 0
	for _, v := range allVaults {
		if filterFunc(v) && !helpers.Contains(chain.BlacklistedVaults, v.Address) {
			estimatedCapacity++
		}
	}

	// Pre-allocate the slice to avoid reallocations
	data := make([]TExternalVault, 0, estimatedCapacity)

	// Process vaults
	for _, currentVault := range allVaults {
		if !filterFunc(currentVault) {
			continue
		}

		vaultAddress := currentVault.Address
		if helpers.Contains(chain.BlacklistedVaults, vaultAddress) {
			continue
		}

		newVault, err := CreateExternalVault(currentVault)
		if err != nil {
			continue
		}
		if migrable == `none` && (newVault.Details.IsHidden || newVault.Details.IsRetired) && hideAlways {
			continue
		} else if migrable == `nodust` && (newVault.TVL.TVL < 100 || !newVault.Migration.Available) {
			continue
		} else if migrable == `all` && !newVault.Migration.Available {
			continue
		} else if migrable == `ignore` && (newVault.Migration.Available || newVault.Details.IsHidden || newVault.Details.IsRetired) {
			continue
		}

		vaultStrategies, _ := storage.ListStrategiesForVault(chainID, vaultAddress)
		newVault.Strategies = []TStrategy{}
		for _, strategy := range vaultStrategies {
			strategyWithDetails := CreateExternalStrategy(strategy)
			if !strategyWithDetails.ShouldBeIncluded(strategiesCondition) {
				continue
			}

			newVault.Strategies = append(newVault.Strategies, strategyWithDetails)
		}

		data = append(data, newVault)
	}

	//Sort by details.order by default
	sort.SortBy(orderBy, orderDirection, data)
	start := (page - 1) * limit
	end := page * limit
	if end > uint64(len(data)) {
		end = uint64(len(data))
	}
	data = data[start:end]

	return data
}
