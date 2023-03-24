package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yearn/ydaemon/common/env"
	"github.com/yearn/ydaemon/common/helpers"
	"github.com/yearn/ydaemon/internal/harvests"
)

// GetSupportedChains returns a list of supported chains by the API
func GetSupportedChains(c *gin.Context) {
	c.JSON(http.StatusOK, env.SUPPORTED_CHAIN_IDS)
}

// GetHarvests returns a list of harvests
func GetHarvests(c *gin.Context) {
	chainID, ok := helpers.AssertChainID(c.Param("chainID"))
	if !ok {
		c.String(http.StatusBadRequest, "invalid chainID")
		return
	}
	vaultAddress, ok := helpers.AssertAddress(c.Param("address"), chainID)
	if !ok {
		c.String(http.StatusBadRequest, "invalid address")
		return
	}

	allHarvest := harvests.Get(vaultAddress)
	c.JSON(http.StatusOK, allHarvest)
}
