package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/majorfi/ydaemon/internal/utils"
)

//GetSupportedChains returns a list of supported chains by the API
func (y controller) GetSupportedChains(c *gin.Context) {
	c.JSON(http.StatusOK, []uint64{1, 250, 42161})
}

//GetBlacklistedVaults returns a list of blacklisted vaults by the API
func (y controller) GetBlacklistedVaults(c *gin.Context) {
	chainID := queryWithFallback(c.Query("chainID"), "0")
	if chainID == "0" {
		c.JSON(http.StatusOK, utils.BLACKLISTED_VAULTS)
	} else {
		chainIDAsUint, _ := strconv.ParseUint(chainID, 10, 64)
		c.JSON(http.StatusOK, utils.BLACKLISTED_VAULTS[chainIDAsUint])
	}
}
