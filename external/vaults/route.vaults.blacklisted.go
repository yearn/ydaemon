package vaults

import (
	"net/http"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"github.com/yearn/ydaemon/common/env"
	"github.com/yearn/ydaemon/common/helpers"
)

// GetBlacklistedVaults returns a list of blacklisted vaults by the API
func (y Controller) GetBlacklistedVaults(c *gin.Context) {
	chainID := helpers.SafeString(getQuery(c, "chainID"), "0")
	if chainID == "0" {
		blacklistedVaults := []common.Address{}
		for _, chain := range env.GetChains() {
			blacklistedVaults = append(blacklistedVaults, chain.BlacklistedVaults...)
		}
		c.JSON(http.StatusOK, blacklistedVaults)
	} else {
		chainIDAsUint, _ := strconv.ParseUint(chainID, 10, 64)
		chain, ok := env.GetChain(chainIDAsUint)
		if !ok {
			c.JSON(http.StatusNotFound, gin.H{"error": "chain not found"})
			return
		}
		c.JSON(http.StatusOK, chain.BlacklistedVaults)
	}
}
