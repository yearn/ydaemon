package vaults

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yearn/ydaemon/common/env"
	"github.com/yearn/ydaemon/common/helpers"
	"github.com/yearn/ydaemon/common/sort"
	"github.com/yearn/ydaemon/internal/storage"
)

/**************************************************************************************************
** GetAllStrategies retrieves a complete list of strategies for a specific chain.
**
** This endpoint returns all strategies across all vaults on the specified blockchain network.
** Strategies are the core yield-generating components of Yearn vaults, and this endpoint provides
** a comprehensive view of all active strategies on a particular chain.
**
** The endpoint supports the following query parameters:
** - orderBy: Field to sort results by (default: 'address')
** - orderDirection: Sort direction, 'asc' or 'desc' (default: 'asc')
** - strategiesCondition: Filter for strategies, values: 'inQueue', 'debtLimit', 'debtRatio',
**   'absolute', 'all' (default: 'debtRatio')
**
** The function processes data through the following steps:
** 1. Validates the chain ID and retrieves sorting parameters
** 2. Retrieves all vaults for the specified chain
** 3. Filters out blacklisted vaults
** 4. For each valid vault, retrieves and processes all associated strategies
** 5. Filters strategies based on the strategiesCondition parameter
** 6. Sorts the resulting strategies according to the specified order parameters
** 7. Returns the strategies as a JSON array
**
** Endpoint: GET /strategies/:chainID
**
** @param c *gin.Context - The Gin context containing the HTTP request
** @return void - Response is sent directly via Gin with the array of strategies
**************************************************************************************************/
func (y Controller) GetAllStrategies(c *gin.Context) {
	orderBy := helpers.SafeString(getQuery(c, `orderBy`), `address`)
	orderDirection := helpers.SafeString(getQuery(c, `orderDirection`), `asc`)
	strategiesCondition := ValidateStrategyCondition(c, "strategiesCondition")

	// Validate chain ID using the utility function
	chainID, ok := ValidateChainID(c, `chainID`)
	if !ok {
		return
	}

	data := []TStrategy{}
	allVaults, _ := storage.ListVaults(chainID)
	for _, currentVault := range allVaults {
		chain, ok := env.GetChain(chainID)
		if !ok {
			continue
		}
		if helpers.Contains(chain.BlacklistedVaults, currentVault.Address) {
			continue
		}
		vaultStrategies, _ := storage.ListStrategiesForVault(chainID, currentVault.Address)
		for _, strategy := range vaultStrategies {
			strategyWithDetails := NewStrategy().AssignTStrategy(strategy)
			if !strategyWithDetails.ShouldBeIncluded(strategiesCondition) {
				continue
			}
			data = append(data, strategyWithDetails)
		}
	}

	sort.SortBy(orderBy, orderDirection, data)
	c.JSON(http.StatusOK, data)
}
