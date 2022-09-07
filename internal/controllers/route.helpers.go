package controllers

import (
	"bytes"
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/machinebox/graphql"
	"github.com/yearn/ydaemon/internal/ethereum"
	"github.com/yearn/ydaemon/internal/utils"
)

//GetSupportedChains returns a list of supported chains by the API
func (y controller) GetSupportedChains(c *gin.Context) {
	c.JSON(http.StatusOK, utils.SUPPORTED_CHAIN_IDS)
}

//GetBlacklistedVaults returns a list of blacklisted vaults by the API
func (y controller) GetBlacklistedVaults(c *gin.Context) {
	chainID := utils.ValueWithFallback(c.Query("chainID"), "0")
	if chainID == "0" {
		c.JSON(http.StatusOK, utils.BLACKLISTED_VAULTS)
	} else {
		chainIDAsUint, _ := strconv.ParseUint(chainID, 10, 64)
		c.JSON(http.StatusOK, utils.BLACKLISTED_VAULTS[chainIDAsUint])
	}
}

//GetGraph returns a list of blacklisted vaults by the API
func (y controller) GetGraph(c *gin.Context) {
	// Get the chainID from the URI
	chainID, err := strconv.ParseUint(c.Param("chainID"), 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{`error`: `invalid chainID`})
		return
	}

	// Read the body as string
	buf := new(bytes.Buffer)
	buf.ReadFrom(c.Request.Body)

	// Proxy the request to the subgraph
	client := graphql.NewClient(ethereum.GetGraphURI(chainID))
	request := graphql.NewRequest(buf.String())
	var response interface{}
	if err := client.Run(context.Background(), request, &response); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{`error`: `bad request`})
		return
	}
	c.JSON(http.StatusOK, response)
}
