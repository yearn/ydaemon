package vaults

import (
	"fmt"
	"math"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/yearn/ydaemon/common/env"
	"github.com/yearn/ydaemon/common/helpers"
	"github.com/yearn/ydaemon/internal/fetcher"
	"github.com/yearn/ydaemon/internal/storage"
)

/**************************************************************************************************
** computeChainTVL calculates the total value locked (TVL) across all Yearn vaults on a chain.
**
** This helper function iterates through all vaults on the specified chain, filters for only
** Yearn vaults (excluding blacklisted ones), and sums their individual TVL values to produce
** a total TVL figure for the chain.
**
** The function performs the following steps:
** 1. Retrieves all vaults for the specified chain from storage
** 2. Filters out non-Yearn vaults using the IsYearn inclusion flag
** 3. Filters out blacklisted vaults defined in the chain configuration
** 4. For each valid vault, calculates its TVL using the BuildVaultTVL function
** 5. Accumulates the individual vault TVLs into a total chain TVL
**
** @param chainID uint64 - The ID of the blockchain to calculate TVL for
** @param c *gin.Context - The Gin context (used for potential error reporting)
** @return float64 - The calculated total TVL for the chain in USD
**************************************************************************************************/
func computeChainTVL(chainID uint64, c *gin.Context) float64 {
	tvl := 0.0

	// Verify chain is supported
	if !helpers.Contains(env.SUPPORTED_CHAIN_IDS, chainID) {
		HandleError(c, fmt.Errorf("chain %d is not supported", chainID), http.StatusBadRequest,
			fmt.Sprintf("Chain %d is not supported", chainID), "ComputeChainTVL")
		return tvl
	}

	// Get chain configuration
	chain, ok := env.GetChain(chainID)
	if !ok {
		HandleError(c, fmt.Errorf("chain configuration not found for chainID %d", chainID), http.StatusInternalServerError,
			"Internal configuration error", "ComputeChainTVL")
		return tvl
	}

	// Retrieve all vaults for the chain
	_, vaultsList := storage.ListVaults(chainID)

	// Process each vault and calculate total TVL
	for _, currentVault := range vaultsList {
		// Skip non-Yearn vaults
		if !currentVault.Metadata.Inclusion.IsYearn {
			continue
		}

		// Skip blacklisted vaults
		if helpers.Contains(chain.BlacklistedVaults, currentVault.Address) {
			continue
		}

		// Calculate TVL for the vault
		vaultTVL := fetcher.BuildVaultTVL(currentVault)

		// Add to total (handle potential NaN/Inf values)
		if !math.IsNaN(vaultTVL.TVL) && !math.IsInf(vaultTVL.TVL, 0) {
			tvl += vaultTVL.TVL
		} else {
			HandleError(c, fmt.Errorf("invalid TVL value for vault %s: %v", currentVault.Address.Hex(), vaultTVL.TVL),
				http.StatusInternalServerError, "Invalid TVL calculation result", "ComputeChainTVL")
		}
	}

	return tvl
}

/**************************************************************************************************
** GetAllVaultsTVL retrieves the TVL across all supported chains.
**
** This endpoint provides a comprehensive view of the total value locked (TVL) across all
** blockchains supported by Yearn. It calculates TVL for each chain concurrently using goroutines
** and returns both the per-chain breakdown and the overall total.
**
** The function processes data through the following steps:
** 1. Initializes a wait group to coordinate concurrent operations
** 2. Creates a mutex to protect shared access to the TVL map and total variable
** 3. For each supported chain, spawns a goroutine to calculate its TVL
** 4. Each goroutine safely adds its result to the chain-specific TVL map and increments the total
** 5. Waits for all calculations to complete
** 6. Returns a JSON response with both the total TVL and per-chain breakdowns
**
** Endpoint: GET /vaults/tvl
**
** @param c *gin.Context - The Gin context containing the HTTP request
** @return void - Response is sent directly via Gin with the TVL data
**************************************************************************************************/
func (y Controller) GetAllVaultsTVL(c *gin.Context) {
	var wg sync.WaitGroup
	var mutex sync.Mutex
	total := 0.0
	tvl := make(map[uint64]float64)

	for chainID := range env.GetChains() {
		wg.Add(1)
		go func(chainID uint64) {
			defer wg.Done()

			// Calculate chain TVL
			chainTVL := computeChainTVL(chainID, c)

			// Safely update shared data structures
			mutex.Lock()
			tvl[chainID] = chainTVL
			total += chainTVL
			mutex.Unlock()
		}(uint64(chainID))
	}

	wg.Wait()

	c.JSON(http.StatusOK, gin.H{
		"total":  total,
		"chains": tvl,
	})
}

/**************************************************************************************************
** GetVaultsTVL retrieves the TVL for a specific chain.
**
** This endpoint calculates and returns the total value locked (TVL) across all Yearn vaults
** on the specified blockchain network. This provides insight into the total assets managed by
** Yearn on that particular chain.
**
** The function processes data through the following steps:
** 1. Validates the chain ID parameter from the request
** 2. Calls computeChainTVL to calculate the total TVL for the chain
** 3. Returns the TVL as a JSON response
**
** Endpoint: GET /vaults/:chainID/tvl
**
** @param c *gin.Context - The Gin context containing the HTTP request
** @return void - Response is sent directly via Gin with the chain's TVL value
**************************************************************************************************/
func (y Controller) GetVaultsTVL(c *gin.Context) {
	// Validate chain ID using the utility function
	chainID, ok := ValidateChainID(c, "chainID")
	if !ok {
		return
	}
	c.JSON(http.StatusOK, computeChainTVL(chainID, c))
}
