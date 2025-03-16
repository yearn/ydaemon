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
	orderBy := helpers.SafeString(getQuery(c, `orderBy`), `featuringScore`)
	orderDir := helpers.SafeString(getQuery(c, `orderDirection`), `asc`)
	stratCon := selectStrategiesCondition(getQuery(c, `strategiesCondition`))

	// Validate chain ID
	chainID, ok := helpers.AssertChainID(c.Param(`chainID`))
	if !ok {
		c.String(http.StatusBadRequest, `invalid chainID`)
		return
	}

	// Validate address parameter
	addressesParam := c.Param(`addresses`)
	if addressesParam == "" {
		c.String(http.StatusBadRequest, `addresses parameter cannot be empty`)
		return
	}

	// Process addresses
	addressesStr := strings.Split(strings.ToLower(addressesParam), `,`)
	if len(addressesStr) == 0 {
		c.String(http.StatusBadRequest, `at least one address must be provided`)
		return
	}

	// Validate each address format (basic check)
	for i, addr := range addressesStr {
		if !strings.HasPrefix(addr, "0x") || len(addr) != 42 {
			c.String(http.StatusBadRequest, fmt.Sprintf(`invalid address format at position %d: %s`, i, addr))
			return
		}
	}

	// Fetch vaults and prepare response
	data := []TExternalVault{}
	allVaults, err := storage.ListVaults(chainID)
	if err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("failed to retrieve vaults: %v", err))
		return
	}

	// Process matching vaults
	matchCount := 0
	for _, currentVault := range allVaults {
		vaultAddressLower := strings.ToLower(currentVault.Address.Hex())
		if !helpers.Contains(addressesStr, vaultAddressLower) {
			continue
		}
		matchCount++

		vaultAddress := currentVault.Address
		chain, ok := env.GetChain(chainID)
		if !ok {
			continue
		}
		if helpers.Contains(chain.BlacklistedVaults, vaultAddress) {
			continue
		}

		newVault, err := NewVault().AssignTVault(currentVault)
		if err != nil {
			// Log error but continue with other vaults
			continue
		}

		vaultStrategies, _ := storage.ListStrategiesForVault(chainID, vaultAddress)

		newVault.Strategies = []TStrategy{}
		for _, strategy := range vaultStrategies {
			strategyWithDetails := NewStrategy().AssignTStrategy(strategy)
			if !strategyWithDetails.ShouldBeIncluded(stratCon) {
				continue
			}

			newVault.Strategies = append(newVault.Strategies, strategyWithDetails)
		}

		data = append(data, newVault)
	}

	// Provide informative response if no vaults matched
	if len(data) == 0 && matchCount == 0 {
		c.JSON(http.StatusOK, []TExternalVault{}) // Return empty array, not an error
		return
	}

	//Sort by details.order by default
	sort.SortBy(orderBy, orderDir, data)
	c.JSON(http.StatusOK, data)
}
