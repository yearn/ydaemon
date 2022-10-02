package utils

import (
	"bytes"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/machinebox/graphql"
	"github.com/yearn/ydaemon/internal/utils/env"
	"github.com/yearn/ydaemon/internal/utils/helpers"
)

//GetSupportedChains returns a list of supported chains by the API
func GetSupportedChains(c *gin.Context) {
	c.JSON(http.StatusOK, env.SUPPORTED_CHAIN_IDS)
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
	graphQLEndpoint, ok := env.THEGRAPH_ENDPOINTS[chainID]
	if !ok {
		c.String(http.StatusInternalServerError, "Impossible to fetch subgraph")
		return
	}

	client := graphql.NewClient(graphQLEndpoint)
	request := graphql.NewRequest(buf.String())
	var response interface{}
	if err := client.Run(context.Background(), request, &response); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{`error`: `bad request`})
		return
	}
	c.JSON(http.StatusOK, response)
}
