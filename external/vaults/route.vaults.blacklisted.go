package vaults

import (
	"net/http"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"github.com/yearn/ydaemon/common/env"
	"github.com/yearn/ydaemon/common/helpers"
)

/**************************************************************************************************
** GetBlacklistedVaults retrieves the list of vaults excluded from API results.
**
** This endpoint returns the addresses of vaults that are blacklisted across all chains or for
** a specific chain. Blacklisted vaults are excluded from regular API responses for various
** reasons such as being deprecated, vulnerable, or not meeting current standards.
**
** The endpoint supports an optional chainID query parameter:
** - If chainID is omitted or "0", returns blacklisted vaults across all supported chains
** - If chainID is provided, returns blacklisted vaults for only that specific chain
**
** Endpoint: GET /vaults/blacklisted?chainID=:chainID
**
** @param c *gin.Context - The Gin context containing the HTTP request
** @return void - Response is sent directly via Gin with the list of blacklisted vault addresses
**************************************************************************************************/
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
