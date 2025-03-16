package vaults

import (
	"net/http"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"github.com/yearn/ydaemon/common/helpers"
	"github.com/yearn/ydaemon/internal/models"
	"github.com/yearn/ydaemon/internal/storage"
)

/**************************************************************************************************
** GetVault retrieves detailed information about a specific vault.
**
** This endpoint returns comprehensive data about a single vault identified by its chain ID
** and address. It includes complete vault information, associated strategies, and related
** data like APR, TVL, and token details.
**
** The function handles the following:
** 1. Validating the chain ID and address parameters
** 2. Retrieving the vault from storage
** 3. Converting the internal vault structure to the external format
** 4. Filtering strategies based on the strategiesCondition parameter
** 5. Including simplified version if the vault is also a strategy
**
** Endpoint: GET /vaults/:chainID/:address
**
** @param c *gin.Context - The Gin context containing the HTTP request
** @return void - Response is sent directly via Gin
**************************************************************************************************/
func (y Controller) GetVault(c *gin.Context) {
	chainID, ok := helpers.AssertChainID(c.Param("chainID"))
	if !ok {
		c.String(http.StatusBadRequest, "invalid chainID")
		return
	}
	address, ok := helpers.AssertAddress(c.Param("address"), chainID)
	if !ok {
		c.String(http.StatusBadRequest, "invalid address")
		return
	}

	strategiesCondition := selectStrategiesCondition(getQuery(c, "strategiesCondition"))
	currentVault, ok := storage.GetVault(chainID, address)
	if !ok {
		c.String(http.StatusBadRequest, "invalid vault")
		return
	}

	newVault, err := NewVault().AssignTVault(currentVault)
	if err != nil {
		c.String(http.StatusBadRequest, "vault not found")
		return
	}
	vaultStrategies, _ := storage.ListStrategiesForVault(chainID, currentVault.Address)
	newVault.Strategies = []TStrategy{}
	for _, strategy := range vaultStrategies {
		strategyWithDetails := NewStrategy().AssignTStrategy(strategy)
		if !strategyWithDetails.ShouldBeIncluded(strategiesCondition) {
			continue
		}
		newVault.Strategies = append(newVault.Strategies, strategyWithDetails)
	}

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
