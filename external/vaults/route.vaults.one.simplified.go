package vaults

import (
	"fmt"
	"net/http"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"github.com/yearn/ydaemon/common/env"
	"github.com/yearn/ydaemon/common/helpers"
	"github.com/yearn/ydaemon/internal/models"
	"github.com/yearn/ydaemon/internal/storage"
)

/**************************************************************************************************
** GetSimplifiedVault retrieves a lightweight version of a specific vault's details.
**
** This endpoint provides a simplified representation of a single vault, optimized for clients
** that need essential information without the full detail load. It's particularly useful for
** client applications focusing on performance or bandwidth constraints.
**
** The endpoint handles the following:
** 1. Validating the chain ID and address parameters
** 2. Retrieving the vault from storage and converting to external format
** 3. Including only relevant strategies based on the strategiesCondition parameter
** 4. Applying the featuring score calculation for sorting relevance
** 5. Handling special cases where the vault is also registered as a strategy
** 6. Returning a simplified representation with essential vault information
**
** Endpoint: GET /vaults/:chainID/:address/simplified
**
** @param c *gin.Context - The Gin context containing the HTTP request
** @return void - Response is sent directly via Gin with the simplified vault representation
**************************************************************************************************/
func (y Controller) GetSimplifiedVault(c *gin.Context) {
	/** 🔵 - Yearn *************************************************************************************
	** chainID: The chain IDs for which the vaults are to be returned. It is obtained from the
	** 'chainIDs' query parameter in the request.
	**************************************************************************************************/
	// Validate chainID parameter
	chainIDParam := c.Param("chainID")
	if chainIDParam == "" {
		c.String(http.StatusBadRequest, "chainID parameter is required")
		return
	}

	chainID, ok := helpers.AssertChainID(chainIDParam)
	if !ok {
		c.String(http.StatusBadRequest, "invalid chainID")
		return
	}

	// Verify chain is supported
	if !helpers.Contains(env.SUPPORTED_CHAIN_IDS, chainID) {
		c.String(http.StatusBadRequest, fmt.Sprintf("chain %d is not supported", chainID))
		return
	}

	/** 🔵 - Yearn *************************************************************************************
	** address: The address of the vault to be returned. It is obtained from the 'address' query
	** parameter in the request.
	**************************************************************************************************/
	// Validate address parameter
	addressParam := c.Param("address")
	if addressParam == "" {
		c.String(http.StatusBadRequest, "address parameter is required")
		return
	}

	address, ok := helpers.AssertAddress(addressParam, chainID)
	if !ok {
		c.String(http.StatusBadRequest, "invalid address")
		return
	}

	/** 🔵 - Yearn *************************************************************************************
	** strategiesCondition: A string that determines the condition for selecting strategies. It is
	** obtained from the 'strategiesCondition' query parameter in the request.
	**************************************************************************************************/
	strategiesCondition := selectStrategiesCondition(getQuery(c, `strategiesCondition`))

	/** 🔵 - Yearn *************************************************************************************
	** The following block of code will store the final vault to be returned in the response, which will
	** receive a bunch of mutation to be transformed to a simplified version of the vault.
	**************************************************************************************************/
	currentVault, ok := storage.GetVault(chainID, address)
	if !ok {
		c.String(http.StatusBadRequest, "vault not found")
		return
	}

	newVault, err := NewVault().AssignTVault(currentVault)
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("failed to process vault data: %v", err))
		return
	}

	APRAsFloat := 0.0
	if newVault.APR.NetAPR != nil {
		// APR.NetAPR.Float64() returns (float64, big.Accuracy), not an error
		// big.Accuracy is an int8 type that indicates precision, not an error
		APRAsFloat, _ = newVault.APR.NetAPR.Float64()
	}
	newVault.FeaturingScore = newVault.TVL.TVL * APRAsFloat
	if newVault.Details.IsHighlighted {
		newVault.FeaturingScore = newVault.FeaturingScore * 1e18
	}

	/** 🔵 - Yearn *************************************************************************************
	** For the vault, it performs the following operations:
	**
	** 1. It calls the 'ListStrategiesForVault' function to get a list of all strategies for the
	**    current vault. The function takes two parameters: the chain ID of the vault and the address
	**    of the vault.
	**
	** 2. It initializes an empty slice of pointers to 'TStrategy' objects for the current vault.
	**
	** 3. It then enters a nested loop that iterates over each strategy in 'vaultStrategies'. For each
	**    strategy, it performs the following operations:
	**
	**    a. It initializes a pointer to a 'TStrategy' object.
	**
	**    b. It calls the 'NewStrategy' function to create a new 'TStrategy' object and assigns the
	**       current strategy to it using the 'AssignTStrategy' method.
	**
	**    c. It checks if the strategy should be included based on the 'strategiesCondition'. If not,
	**       it skips the current iteration of the loop.
	**
	** 4. It appends the current strategies to the vault
	**************************************************************************************************/
	vaultStrategies, _ := storage.ListStrategiesForVault(currentVault.ChainID, common.HexToAddress(newVault.Address))

	newVault.Strategies = []TStrategy{}
	for _, strategy := range vaultStrategies {
		var externalStrategy TStrategy
		strategyWithDetails := NewStrategy().AssignTStrategy(strategy)
		if !strategyWithDetails.ShouldBeIncluded(strategiesCondition) {
			continue
		}

		externalStrategy = strategyWithDetails
		newVault.Strategies = append(newVault.Strategies, externalStrategy)
	}

	// Special handling for vaults that are also registered as strategies
	if vaultAsStrategy, ok := storage.GuessStrategy(newVault.ChainID, common.HexToAddress(newVault.Address)); ok {
		simplified := toSimplifiedVersion(newVault, vaultAsStrategy)
		simplified.Description = newVault.Description
		if simplified.Description == "" {
			simplified.Description = vaultAsStrategy.Description
		}
		c.JSON(http.StatusOK, simplified)
		return
	}

	simplified := toSimplifiedVersion(newVault, models.TStrategy{})
	simplified.Description = newVault.Description

	c.JSON(http.StatusOK, simplified)
}
