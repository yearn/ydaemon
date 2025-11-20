package vaults

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"github.com/yearn/ydaemon/common/bigNumber"
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
	// Create a timeout context for the request to prevent hanging operations
	ctx, cancel := context.WithTimeout(c.Request.Context(), 10*time.Second)
	defer cancel()

	/** 🔵 - Yearn *************************************************************************************
	** chainID: The chain IDs for which the vaults are to be returned. It is obtained from the
	** 'chainIDs' query parameter in the request.
	**************************************************************************************************/
	// Validate chainID parameter using the utility function
	chainID, ok := validateChainID(c, "chainID")
	if !ok {
		// validateChainID already sets the response, so we just return
		return
	}

	/** 🔵 - Yearn *************************************************************************************
	** address: The address of the vault to be returned. It is obtained from the 'address' query
	** parameter in the request.
	**************************************************************************************************/
	// Validate address parameter using the utility function
	address, ok := validateAddress(c, "address", chainID)
	if !ok {
		// validateAddress already sets the response, so we just return
		return
	}

	/** 🔵 - Yearn *************************************************************************************
	** strategiesCondition: A string that determines the condition for selecting strategies. It is
	** obtained from the 'strategiesCondition' query parameter in the request.
	**************************************************************************************************/
	strategiesCondition := validateStrategyCondition(c, "strategiesCondition")

	/** 🔵 - Yearn *************************************************************************************
	** The following block of code will store the final vault to be returned in the response, which will
	** receive a bunch of mutation to be transformed to a simplified version of the vault.
	**************************************************************************************************/
	currentVault, ok := storage.GetVault(chainID, address)
	if !ok {
		handleError(c, fmt.Errorf("vault not found for chain %d and address %s", chainID, address.Hex()),
			http.StatusNotFound, "Vault not found", "GetSimplifiedVault")
		return
	}

	// Verify context is still valid before proceeding
	select {
	case <-ctx.Done():
		handleError(c, fmt.Errorf("request timed out while retrieving vault data"),
			http.StatusGatewayTimeout, "Request processing timed out", "GetSimplifiedVault")
		return
	default:
		// Continue processing
	}

	// Process the vault and return a simplified version
	newVault, err := CreateExternalVault(currentVault)
	if err != nil {
		handleError(c, fmt.Errorf("failed to process vault %s on chain %d: %w",
			address.String(), chainID, err),
			http.StatusInternalServerError, "Error processing vault data", "GetSimplifiedVault")
		return
	}

	// Calculate featuring score with appropriate error checking for nil values
	APRAsFloat := 0.0
	if newVault.APR.NetAPR != nil {
		// APR.NetAPR.Float64() returns (float64, big.Accuracy), not an error
		// big.Accuracy is an int8 type that indicates precision, not an error
		APRAsFloat, _ = newVault.APR.NetAPR.Float64()
	}

	// Check for potential arithmetic overflow due to very large values
	if newVault.TVL.TVL > 1e30 || APRAsFloat > 1e10 {
		c.Error(fmt.Errorf("potentially unsafe values for featuring score calculation: TVL=%f, APR=%f",
			newVault.TVL.TVL, APRAsFloat))
		// Continue with normal calculation but cap the values to prevent overflow
		if newVault.TVL.TVL > 1e30 {
			newVault.TVL.TVL = 1e30
		}
		if APRAsFloat > 1e10 {
			APRAsFloat = 1e10
		}
	}

	// Verify context is still valid before continuing with potential expensive calculations
	select {
	case <-ctx.Done():
		handleError(c, fmt.Errorf("request timed out during featuring score calculation"),
			http.StatusGatewayTimeout, "Request processing timed out", "GetSimplifiedVault")
		return
	default:
		// Continue processing
	}

	newVault.FeaturingScore = newVault.TVL.TVL * APRAsFloat
	if newVault.Details.IsHighlighted {
		newVault.FeaturingScore = newVault.FeaturingScore * 1e18
	}

	// Fetch and process strategies with context awareness
	vaultStrategies, _ := storage.ListStrategiesForVault(chainID, address)
	newVault.Strategies = []TExternalStrategy{}
	for _, strategy := range vaultStrategies {
		strategyWithDetails := CreateExternalStrategy(strategy)

		if !strategyWithDetails.ShouldBeIncluded(strategiesCondition) {
			continue
		}

		strategyAddress := common.HexToAddress(strategyWithDetails.Address)
		
		for _, debt := range newVault.Debts {
			if debt.Strategy == strategyAddress.Hex() {
				if debt.CurrentDebt != nil {
					strategyWithDetails.Details.TotalDebt = bigNumber.NewInt().SetString(*debt.CurrentDebt)
				} else if debt.TotalDebt != nil {
					strategyWithDetails.Details.TotalDebt = bigNumber.NewInt().SetString(*debt.TotalDebt)
				} else {
					strategyWithDetails.Details.TotalDebt = bigNumber.NewInt().SetString("0")
				}
				break
			}
		}

		newVault.Strategies = append(newVault.Strategies, strategyWithDetails)
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
