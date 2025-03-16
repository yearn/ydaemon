package vaults

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/yearn/ydaemon/internal/storage"
)

/**************************************************************************************************
** GetStrategy retrieves detailed information about a specific vault strategy.
**
** This endpoint returns comprehensive data about a single strategy identified by its chain ID
** and address. Strategies are the core components that generate yield for Yearn vaults, and
** this endpoint provides access to their configuration and performance data.
**
** The endpoint handles the following process:
** 1. Validating the chain ID and address parameters from the request
** 2. Retrieving the strategy from storage using GuessStrategy, which can find strategies
**    even if they're not explicitly linked to a vault
** 3. Converting the internal strategy structure to the external TStrategy format
** 4. Returning the strategy data as a JSON response
**
** Endpoint: GET /strategies/:chainID/:address
**
** @param c *gin.Context - The Gin context containing the HTTP request
** @return void - Response is sent directly via Gin with the strategy data
**************************************************************************************************/
func (y Controller) GetStrategy(c *gin.Context) {
	// Create a timeout context for the request to prevent hanging operations
	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	// Validate chain ID using the utility function
	chainID, ok := ValidateChainID(c, "chainID")
	if !ok {
		// ValidateChainID already sets the response, so we just return
		return
	}

	// Validate address using the utility function
	address, ok := ValidateAddress(c, "address", chainID)
	if !ok {
		// ValidateAddress already sets the response, so we just return
		return
	}

	// Check if the context is still valid before proceeding
	select {
	case <-ctx.Done():
		HandleError(c, fmt.Errorf("operation timed out while validating parameters"),
			http.StatusGatewayTimeout, "Request processing timed out", "GetStrategy")
		return
	default:
		// Continue processing
	}

	// Look up the strategy
	strategy, ok := storage.GuessStrategy(chainID, address)
	if !ok {
		HandleError(c, fmt.Errorf("strategy not found for address %s on chain %d",
			address.String(), chainID),
			http.StatusNotFound, "Strategy not found", "GetStrategy")
		return
	}

	// Try to convert the strategy to the external format
	var newStrategy TStrategy
	var conversionErr error

	// Use a function with recover to handle potential panics during conversion
	func() {
		defer func() {
			if r := recover(); r != nil {
				conversionErr = fmt.Errorf("panic while processing strategy %s: %v",
					address.String(), r)
			}
		}()

		newStrategy = CreateExternalStrategy(strategy)

		// Additional validation on the resulted strategy
		if newStrategy.Address == "" {
			conversionErr = fmt.Errorf("strategy conversion resulted in invalid data")
		}
	}()

	// Handle conversion errors
	if conversionErr != nil {
		HandleError(c, fmt.Errorf("failed to process strategy %s on chain %d: %w",
			address.String(), chainID, conversionErr),
			http.StatusInternalServerError, "Error processing strategy data", "GetStrategy")
		return
	}

	// Check if the context is still valid before sending the response
	select {
	case <-ctx.Done():
		HandleError(c, fmt.Errorf("operation timed out while processing strategy"),
			http.StatusGatewayTimeout, "Request processing timed out", "GetStrategy")
		return
	default:
		// Continue processing
	}

	// Return the strategy
	c.JSON(http.StatusOK, newStrategy)
}
