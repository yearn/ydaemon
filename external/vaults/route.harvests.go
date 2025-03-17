package vaults

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/machinebox/graphql"
	"github.com/yearn/ydaemon/common/bigNumber"
	"github.com/yearn/ydaemon/common/env"
	"github.com/yearn/ydaemon/common/helpers"
	"github.com/yearn/ydaemon/common/sort"
	"github.com/yearn/ydaemon/internal/models"
)

/**************************************************************************************************
** graphQLHarvestRequestForOneVault creates a GraphQL query to fetch harvest events for vaults.
**
** This helper function builds a GraphQL query that retrieves harvest events for a specific set
** of vault addresses. Harvest events represent when strategies report profits or losses back
** to their associated vaults.
**
** The query limits results to the most recent 1000 harvests and includes detailed information
** about each harvest event, including vault and strategy details, transaction information,
** timestamps, and profit/loss figures.
**
** @param vaultAddresses []string - Array of vault addresses to fetch harvests for
** @param orderBy string - Field to sort results by
** @param orderDirection string - Sort direction, 'asc' or 'desc'
** @param limit int - Maximum number of records to return (defaults to 1000 if 0 or negative)
** @return *graphql.Request - The constructed GraphQL request object ready to be executed
**************************************************************************************************/
func graphQLHarvestRequestForOneVault(vaultAddresses []string, orderBy, orderDirection string, limit int) *graphql.Request {
	// Validate and set default limit
	if limit <= 0 {
		limit = 1000 // Default limit
	} else if limit > 5000 {
		limit = 5000 // Cap maximum to avoid oversized requests
	}

	// Validate orderBy field
	validOrderFields := map[string]bool{
		"timestamp":   true,
		"profit":      true,
		"loss":        true,
		"profitValue": true,
		"lossValue":   true,
	}
	if !validOrderFields[orderBy] {
		orderBy = "timestamp" // Default to timestamp if invalid
	}

	// Validate orderDirection
	if orderDirection != "asc" && orderDirection != "desc" {
		orderDirection = "desc" // Default to desc if invalid
	}

	// Join addresses with proper formatting and escaping
	var sanitizedAddresses []string
	for _, addr := range vaultAddresses {
		// Basic validation for address format
		if strings.HasPrefix(addr, "0x") && len(addr) == 42 {
			sanitizedAddresses = append(sanitizedAddresses, addr)
		}
	}

	// Construct and return the query
	return graphql.NewRequest(`{
		harvests(first: ` + fmt.Sprint(limit) + `, orderBy: ` + orderBy + `, orderDirection: ` + orderDirection + `, where: {vault_in: ["` + strings.Join(sanitizedAddresses, `", "`) + `"]})
		` + helpers.GetHarvestsForVaults() + `
    }`)
}

/**************************************************************************************************
** GetHarvestsForVault retrieves harvest events for specified vaults.
**
** This endpoint returns a history of harvest events for the specified vaults on a particular
** chain. Harvests represent profit-reporting events where strategies report gains or losses
** back to their vaults. This information is valuable for analyzing vault performance over time.
**
** The endpoint supports the following parameters:
** - chainID: The chain ID from the URL path parameter
** - addresses: Comma-separated list of vault addresses from the URL path parameter
** - orderBy: Field to sort results by (default: 'timestamp')
** - orderDirection: Sort direction, 'asc' or 'desc' (default: 'desc')
** - limit: Maximum number of results to return (default: 1000)
**
** The function processes data through the following steps:
** 1. Validates the chain ID and retrieves sorting parameters
** 2. Ensures the chain has a valid subgraph URI for querying
** 3. Builds and executes a GraphQL query to fetch harvest events
** 4. Processes each harvest event, calculating USD values for profits/losses
** 5. Filters out harvests with zero profit and loss
** 6. Sorts the results according to the specified order parameters
**
** Endpoint: GET /vaults/:chainID/harvests/:addresses
**
** @param c *gin.Context - The Gin context containing the HTTP request
** @return void - Response is sent directly via Gin with the list of harvest events
**************************************************************************************************/
func (y Controller) GetHarvestsForVault(c *gin.Context) {
	// Create a timeout context for the request to prevent hanging operations
	ctx, cancel := context.WithTimeout(c.Request.Context(), 15*time.Second)
	defer cancel()

	// Validate chain ID using the utility function
	chainID, ok := validateChainID(c, "chainID")
	if !ok {
		// validateChainID already sets the response, so we just return
		return
	}

	// Validate addresses parameter
	addressesParam := c.Param("addresses")
	if addressesParam == "" {
		handleError(c, fmt.Errorf("addresses parameter cannot be empty"),
			http.StatusBadRequest, "Missing required parameter", "GetHarvestsForVault")
		return
	}

	// Process addresses
	addressesStr := strings.Split(strings.ToLower(addressesParam), ",")
	if len(addressesStr) == 0 {
		handleError(c, fmt.Errorf("at least one address must be provided"),
			http.StatusBadRequest, "Invalid parameter value", "GetHarvestsForVault")
		return
	}

	// Validate each address format (basic check)
	for i, addr := range addressesStr {
		if !strings.HasPrefix(addr, "0x") || len(addr) != 42 {
			handleError(c, fmt.Errorf("invalid address format at position %d: %s", i, addr),
				http.StatusBadRequest, "Invalid address format", "GetHarvestsForVault")
			return
		}
	}

	// Get and validate query parameters with defaults
	orderBy := validateStringChoiceQuery(c, "orderBy", "timestamp",
		[]string{"timestamp", "profit", "loss", "profitValue", "lossValue"}, "GetHarvestsForVault")

	orderDirection := validateStringChoiceQuery(c, "orderDirection", "desc",
		[]string{"asc", "desc"}, "GetHarvestsForVault")

	limit := int(validateNumericQuery(c, "limit", 1000, 1, 5000, "GetHarvestsForVault"))

	// Get chain configuration
	chain, ok := env.GetChain(chainID)
	if !ok {
		handleError(c, fmt.Errorf("chain configuration not found for chainID %d", chainID),
			http.StatusInternalServerError, "Internal configuration error", "GetHarvestsForVault")
		return
	}

	// Validate subgraph endpoint availability
	graphQLEndpoint := chain.SubgraphURI
	if graphQLEndpoint == "" {
		handleError(c, fmt.Errorf("no graph endpoint configured for chainID %d", chainID),
			http.StatusInternalServerError, "Subgraph not available", "GetHarvestsForVault")
		return
	}

	// Create GraphQL request
	client := graphql.NewClient(graphQLEndpoint)
	request := graphQLHarvestRequestForOneVault(addressesStr, orderBy, orderDirection, limit)

	// Execute GraphQL request with timeout context
	var response models.TGraphQLHarvestRequestForOneVault
	if err := client.Run(ctx, request, &response); err != nil {
		// Check if this is a context timeout
		if ctx.Err() == context.DeadlineExceeded {
			handleError(c, fmt.Errorf("GraphQL request timed out after 15 seconds: %w", err),
				http.StatusGatewayTimeout, "Request to subgraph timed out", "GetHarvestsForVault")
			return
		}

		// Otherwise, it's another kind of GraphQL error
		handleError(c, fmt.Errorf("failed to execute GraphQL request: %w", err),
			http.StatusInternalServerError, "Failed to fetch data from subgraph", "GetHarvestsForVault")
		return
	}

	// For each harvest from the subgraph, compute our TExternalVaultHarvest structure
	harvests := make([]TExternalVaultHarvest, 0, len(response.Harvests))
	harvestsFromSubgraph := response.Harvests

	for _, harvest := range harvestsFromSubgraph {
		// Check for context timeout in long loops
		select {
		case <-ctx.Done():
			handleError(c, fmt.Errorf("request timed out while processing harvest data"),
				http.StatusGatewayTimeout, "Request processing timed out", "GetHarvestsForVault")
			return
		default:
			// Continue processing
		}

		// Process the harvest record
		// Use the available fields from the response and construct a valid TExternalVaultHarvest
		// Note: Some fields like BlockNumber and APR might not be available in the GraphQL response

		// Use profit for profitValue and convert to float
		profitsBN := bigNumber.NewFloat().SetString(harvest.Profit)
		// Skip records with invalid data
		if profitsBN == nil {
			c.Error(fmt.Errorf("skipping harvest record with invalid profit data for vault %s",
				harvest.Vault.Id))
			continue
		}

		profitsFloat, _ := profitsBN.Float64()
		lossFloat := 0.0
		if harvest.Loss != "" {
			lossValue := bigNumber.NewFloat().SetString(harvest.Loss)
			if lossValue != nil {
				lossFloat, _ = lossValue.Float64()
			}
		}

		// Only include records with non-zero profit or loss
		if profitsFloat == 0.0 && lossFloat == 0.0 {
			continue
		}

		harvestNewFormat := TExternalVaultHarvest{
			Timestamp:       harvest.Timestamp,
			Profit:          harvest.Profit,
			Loss:            harvest.Loss,
			TxHash:          harvest.Transaction.Hash,
			ProfitValue:     profitsFloat,
			LossValue:       lossFloat,
			VaultAddress:    harvest.Vault.Id,
			StrategyAddress: harvest.Strategy.Id,
		}

		harvests = append(harvests, harvestNewFormat)
	}

	// Sort the results
	if len(harvests) > 0 {
		sort.SortBy(orderBy, orderDirection, harvests)
	}

	c.JSON(http.StatusOK, harvests)
}
