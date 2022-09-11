package utils

import (
	"bytes"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/machinebox/graphql"
	"github.com/yearn/ydaemon/internal/utils/ethereum"
	"github.com/yearn/ydaemon/internal/utils/helpers"
)

//GetSupportedChains returns a list of supported chains by the API
func GetSupportedChains(c *gin.Context) {
	c.JSON(http.StatusOK, helpers.SUPPORTED_CHAIN_IDS)
}

//GetGraph returns a list of blacklisted vaults by the API
func GetGraph(c *gin.Context) {
	chainID, ok := helpers.AssertChainID(c.Param("chainID"))
	if !ok {
		c.String(http.StatusBadRequest, "invalid chainID")
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
