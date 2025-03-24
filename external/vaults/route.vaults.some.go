package vaults

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/yearn/ydaemon/common/env"
	"github.com/yearn/ydaemon/common/helpers"
	"github.com/yearn/ydaemon/common/sort"
	"github.com/yearn/ydaemon/internal/storage"
)

/**************************************************************************************************
** GetLegacySomeVaults retrieves a specific set of vaults by their addresses.
**
** This endpoint returns detailed information about a specified list of vaults on a particular
** chain. It's useful when a client needs data for a specific subset of vaults rather than all
** vaults or vaults matching a category filter.
**
** The endpoint supports the following parameters:
** - chainID: The chain ID from the URL path parameter
** - addresses: Comma-separated list of vault addresses from the URL path parameter
** - orderBy: Field to sort results by (default: 'featuringScore')
** - orderDirection: Sort direction, 'asc' or 'desc' (default: 'asc')
** - strategiesCondition: Filter for strategies (default: 'debtRatio')
**
** The function processes data through the following steps:
** 1. Validates parameters and retrieves sorting preferences
** 2. Parses the comma-separated addresses list
** 3. Retrieves all vaults for the specified chain
** 4. Filters to include only the requested vault addresses
** 5. Filters out blacklisted vaults
** 6. Processes each vault into the external format with its strategies
** 7. Sorts the results according to the specified order parameters
**
** Endpoint: GET /vaults/:chainID/:addresses
**
** @param c *gin.Context - The Gin context containing the HTTP request
** @return void - Response is sent directly via Gin with the requested vaults
**************************************************************************************************/
func (y Controller) GetLegacySomeVaults(c *gin.Context) {
	orderBy := helpers.SafeString(getQueryParam(c, `orderBy`), `featuringScore`)
	orderDir := helpers.SafeString(getQueryParam(c, `orderDirection`), `asc`)
	stratCon := validateStrategyCondition(c, "strategiesCondition")

	// Validate chain ID using the utility function
	chainID, ok := validateChainID(c, `chainID`)
	if !ok {
		return
	}

	// Validate address parameter
	addressesParam := c.Param(`addresses`)
	if addressesParam == "" {
		handleError(c, fmt.Errorf("addresses parameter cannot be empty"),
			http.StatusBadRequest, "Missing required parameter", "GetLegacySomeVaults")
		return
	}

	// Process addresses
	addressesStr := strings.Split(strings.ToLower(addressesParam), `,`)
	if len(addressesStr) == 0 {
		handleError(c, fmt.Errorf("at least one address must be provided"),
			http.StatusBadRequest, "Invalid parameter value", "GetLegacySomeVaults")
		return
	}

	// Validate each address format (basic check)
	for i, addr := range addressesStr {
		if !strings.HasPrefix(addr, "0x") || len(addr) != 42 {
			handleError(c, fmt.Errorf("invalid address format at position %d: %s", i, addr),
				http.StatusBadRequest, "Invalid address format", "GetLegacySomeVaults")
			return
		}
	}

	// Get chain configuration
	chain, ok := env.GetChain(chainID)
	if !ok {
		handleError(c, fmt.Errorf("chain configuration not found for chainID %d", chainID),
			http.StatusInternalServerError, "Internal configuration error", "GetLegacySomeVaults")
		return
	}

	// Prepare response data
	data := []TExternalVault{}
	vaultsMap, _ := storage.ListVaults(chainID)

	// Only process the requested vaults instead of iterating through all vaults
	for _, addressStr := range addressesStr {
		// Convert string address to common.Address
		address, ok := helpers.AssertAddress(addressStr, chainID)
		if !ok {
			// Skip invalid addresses (we've already validated format, but address might still be invalid)
			continue
		}

		// Check if vault exists in storage
		currentVault, exists := vaultsMap[address]
		if !exists {
			// Vault not found, skip
			continue
		}

		// Check if vault is blacklisted
		if helpers.Contains(chain.BlacklistedVaults, address) {
			continue
		}

		// Process the vault
		newVault, err := CreateExternalVault(currentVault)
		if err != nil {
			// Log error but continue with other vaults
			handleError(c, fmt.Errorf("failed to process vault %s: %w", address.Hex(), err),
				http.StatusInternalServerError, "Error processing vault", "GetLegacySomeVaults")
			continue
		}

		// Get and filter strategies
		vaultStrategies, _ := storage.ListStrategiesForVault(chainID, address)
		newVault.Strategies = []TExternalStrategy{}
		for _, strategy := range vaultStrategies {
			strategyWithDetails := CreateExternalStrategy(strategy)
			if !strategyWithDetails.ShouldBeIncluded(stratCon) {
				continue
			}

			newVault.Strategies = append(newVault.Strategies, strategyWithDetails)
		}

		data = append(data, newVault)
	}

	// Sort the results
	sort.SortBy(orderBy, orderDir, data)
	c.JSON(http.StatusOK, data)
}
