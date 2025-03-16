package vaults

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"github.com/yearn/ydaemon/common/env"
	"github.com/yearn/ydaemon/common/helpers"
)

/**************************************************************************************************
** GetBlacklistedVaults retrieves the list of vaults excluded from API results.
**
** This endpoint returns the addresses of vaults that are blacklisted across all chains or for
** a specific chain. Blacklisted vaults are excluded from regular API responses for various
** reasons such as being deprecated, vulnerable, or not meeting current standards.
**
** The endpoint supports an optional chainID query parameter:
** - If chainID is omitted or "0", returns blacklisted vaults across all supported chains
** - If chainID is provided, returns blacklisted vaults for only that specific chain
**
** Endpoint: GET /vaults/blacklisted?chainID=:chainID
**
** @param c *gin.Context - The Gin context containing the HTTP request
** @return void - Response is sent directly via Gin with the list of blacklisted vault addresses
**************************************************************************************************/
func (y Controller) GetBlacklistedVaults(c *gin.Context) {
	// Create a timeout context for the request to prevent hanging operations
	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	// Get and validate the chainID parameter
	chainIDParam := getQuery(c, "chainID")

	// Verify the param exists and isn't empty
	if chainIDParam == "" {
		chainIDParam = "0" // Default to all chains
	}

	// Validate that the chainID is a valid number format
	chainID := chainIDParam
	if chainID != "0" {
		_, err := strconv.ParseUint(chainID, 10, 64)
		if err != nil {
			HandleError(c, fmt.Errorf("invalid chainID format: %s - %w", chainID, err),
				http.StatusBadRequest, "Invalid chainID parameter", "GetBlacklistedVaults")
			return
		}
	}

	// All chains case
	if chainID == "0" {
		blacklistedVaults := []common.Address{}
		chains := env.GetChains()

		// Check that we have at least one chain configured
		if len(chains) == 0 {
			HandleError(c, fmt.Errorf("no chains configured in the environment"),
				http.StatusInternalServerError, "Configuration error", "GetBlacklistedVaults")
			return
		}

		// Pre-allocate the slice with estimated capacity to avoid reallocations
		estimatedCapacity := len(chains) * 5 // assuming average 5 blacklisted vaults per chain
		blacklistedVaults = make([]common.Address, 0, estimatedCapacity)

		for _, chain := range chains {
			// Check for context timeout
			select {
			case <-ctx.Done():
				HandleError(c, fmt.Errorf("operation timed out while processing blacklisted vaults"),
					http.StatusGatewayTimeout, "Request processing timed out", "GetBlacklistedVaults")
				return
			default:
				// Continue processing
			}

			blacklistedVaults = append(blacklistedVaults, chain.BlacklistedVaults...)
		}

		// Check for context timeout before sending response
		select {
		case <-ctx.Done():
			HandleError(c, fmt.Errorf("operation timed out before sending response"),
				http.StatusGatewayTimeout, "Request processing timed out", "GetBlacklistedVaults")
			return
		default:
			// Continue to response
		}

		c.JSON(http.StatusOK, blacklistedVaults)
		return
	}

	// Parse and validate the chainID
	chainIDAsUint, err := strconv.ParseUint(chainID, 10, 64)
	if err != nil {
		HandleError(c, fmt.Errorf("invalid chainID format: %s - %w", chainID, err),
			http.StatusBadRequest, "Invalid chainID parameter", "GetBlacklistedVaults")
		return
	}

	// Verify chain is supported
	if !helpers.Contains(env.SUPPORTED_CHAIN_IDS, chainIDAsUint) {
		HandleError(c, fmt.Errorf("chain %d is not supported", chainIDAsUint),
			http.StatusBadRequest, "Unsupported chain", "GetBlacklistedVaults")
		return
	}

	// Get chain configuration
	chain, ok := env.GetChain(chainIDAsUint)
	if !ok {
		HandleError(c, fmt.Errorf("chain configuration not found for chainID %d", chainIDAsUint),
			http.StatusInternalServerError, "Chain configuration error", "GetBlacklistedVaults")
		return
	}

	// Check for context timeout before sending response
	select {
	case <-ctx.Done():
		HandleError(c, fmt.Errorf("operation timed out before sending response"),
			http.StatusGatewayTimeout, "Request processing timed out", "GetBlacklistedVaults")
		return
	default:
		// Continue to response
	}

	// Return the blacklisted vaults for the specified chain
	c.JSON(http.StatusOK, chain.BlacklistedVaults)
}
