package vaults

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yearn/ydaemon/common/helpers"
	"github.com/yearn/ydaemon/internal/vaults"
)

//GetVault will, for a given chainID, return a list of all vaults
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

	// strategiesCondition := selectStrategiesCondition(c.Query("strategiesCondition"))
	// withStrategiesDetails := c.Query("strategiesDetails") == "withDetails"
	// withStrategiesRisk := c.Query("strategiesRisk") == "withRisk"
	currentVault, ok := vaults.FindVault(chainID, address)
	if !ok {
		c.String(http.StatusBadRequest, "invalid vault")
		return
	}
	c.JSON(http.StatusOK, currentVault)
}
