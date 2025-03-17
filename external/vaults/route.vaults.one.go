package vaults

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"github.com/yearn/ydaemon/internal/models"
	"github.com/yearn/ydaemon/internal/storage"
)

/************************************************************************************************
** GetVault retrieves detailed information about a single vault by its chain ID and address.
**
** This endpoint returns comprehensive details about a vault, including:
** - Basic vault information (address, name, symbol, description)
** - Token information (underlying assets)
** - APR/APY data
** - TVL (Total Value Locked) metrics
** - Migration status if applicable
** - Detailed risk metrics
** - Associated strategies based on the specified condition
**
** The endpoint accepts the following parameters:
** - chainID: The ID of the chain the vault is deployed on (path parameter)
** - address: The address of the vault to retrieve (path parameter)
** - strategiesCondition: Filter condition for strategies to include (query parameter)
**   Valid values: "all", "inQueue", "debtRatio", "absolute" (default: "debtRatio")
**
** Example request:
**   GET /vaults/1/0x12345...6789?strategiesCondition=all
**
** @route GET /vaults/:chainID/:address
** @param chainID - The chain ID as a URL parameter
** @param address - The vault address as a URL parameter
** @param strategiesCondition - Optional query parameter to filter strategies
** @return TSimplifiedExternalVault - The vault details with filtered strategies
************************************************************************************************/
func (y Controller) GetVault(c *gin.Context) {
	// Create a timeout context for the request to prevent hanging operations
	ctx, cancel := context.WithTimeout(c.Request.Context(), 10*time.Second)
	defer cancel()

	// Validate chainID parameter using the utility function
	chainID, ok := validateChainID(c, "chainID")
	if !ok {
		// validateChainID already sets the response, so we just return
		return
	}

	// Validate address parameter using the utility function
	address, ok := validateAddress(c, "address", chainID)
	if !ok {
		// validateAddress already sets the response, so we just return
		return
	}

	// Validate and process strategiesCondition
	strategiesCondition := validateStrategyCondition(c, "strategiesCondition")

	// Get vault from storage
	currentVault, ok := storage.GetVault(chainID, address)
	if !ok {
		handleError(c, fmt.Errorf("vault not found: %s on chain %d", address.Hex(), chainID),
			http.StatusNotFound, "Vault not found", "GetVault")
		return
	}

	// Verify context is still valid before proceeding
	select {
	case <-ctx.Done():
		handleError(c, fmt.Errorf("request timed out while retrieving vault data"),
			http.StatusGatewayTimeout, "Request processing timed out", "GetVault")
		return
	default:
		// Continue processing
	}

	// Convert to external vault format
	newVault, err := CreateExternalVault(currentVault)
	if err != nil {
		handleError(c, fmt.Errorf("failed to process vault data for vault %s on chain %d: %w",
			address.Hex(), chainID, err),
			http.StatusInternalServerError, "Error processing vault data", "GetVault")
		return
	}

	// Get and filter strategies for the vault
	strategies, success := ProcessStrategiesForVault(
		ctx,
		c,
		chainID,
		currentVault.Address,
		strategiesCondition,
		"GetVault",
	)

	if !success {
		// ProcessStrategiesForVault handles error response
		return
	}

	newVault.Strategies = strategies

	// Verify context is still valid before proceeding to response
	select {
	case <-ctx.Done():
		handleError(c, fmt.Errorf("request timed out before sending response"),
			http.StatusGatewayTimeout, "Request processing timed out", "GetVault")
		return
	default:
		// Continue to response
	}

	// Special handling for vaults that are also strategies
	if vaultAsStrategy, ok := storage.GuessStrategy(newVault.ChainID, common.HexToAddress(newVault.Address)); ok {
		simplified := toSimplifiedVersion(newVault, vaultAsStrategy)
		simplified.Description = newVault.Description
		if simplified.Description == "" {
			simplified.Description = vaultAsStrategy.Description
		}
		c.JSON(http.StatusOK, simplified)
		return
	}

	// Standard vault response
	simplified := toSimplifiedVersion(newVault, models.TStrategy{})
	simplified.Description = newVault.Description
	c.JSON(http.StatusOK, simplified)
}

// isValidStrategiesCondition checks if the strategies condition parameter is valid
func isValidStrategiesCondition(condition string) bool {
	validConditions := map[string]bool{
		"all":       true,
		"inQueue":   true,
		"debtRatio": true,
		"absolute":  true,
	}
	return validConditions[condition]
}
