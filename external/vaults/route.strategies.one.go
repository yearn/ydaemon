package vaults

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yearn/ydaemon/common/helpers"
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

	strategy, ok := storage.GuessStrategy(chainID, address)
	if !ok {
		c.String(http.StatusBadRequest, "invalid strategy")
		return
	}
	newStrategy := NewStrategy().AssignTStrategy(strategy)
	c.JSON(http.StatusOK, newStrategy)
}
