package vaults

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/yearn/ydaemon/internal/utils/env"
	"github.com/yearn/ydaemon/internal/utils/helpers"
)

//GetBlacklistedVaults returns a list of blacklisted vaults by the API
func (y Controller) GetBlacklistedVaults(c *gin.Context) {
	chainID := helpers.SafeString(c.Query("chainID"), "0")
	if chainID == "0" {
		c.JSON(http.StatusOK, env.BLACKLISTED_VAULTS)
	} else {
		chainIDAsUint, _ := strconv.ParseUint(chainID, 10, 64)
		c.JSON(http.StatusOK, env.BLACKLISTED_VAULTS[chainIDAsUint])
	}
}
