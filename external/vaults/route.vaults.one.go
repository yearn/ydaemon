package vaults

import (
	"context"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/machinebox/graphql"
	"github.com/yearn/ydaemon/common/env"
	"github.com/yearn/ydaemon/common/helpers"
	"github.com/yearn/ydaemon/common/logs"
	"github.com/yearn/ydaemon/common/models"
)

func graphQLRequestForOneVault(vaultAddress string, c *gin.Context) *graphql.Request {
	withDetails := helpers.SafeString(c.Query("strategiesDetails"), "noDetails") == "withDetails"

	return graphql.NewRequest(`{
		vault(id: "` + strings.ToLower(vaultAddress) + `") {
			` + helpers.GetGraphRequestVault() + `
			` + helpers.GetGraphRequestStrategies(40, withDetails) + `
		}
	}`)
}

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

	graphQLEndpoint, ok := env.THEGRAPH_ENDPOINTS[chainID]
	if !ok {
		logs.Error("No graph endpoint for chainID", chainID)
		c.String(http.StatusInternalServerError, "Impossible to fetch subgraph")
		return
	}

	client := graphql.NewClient(graphQLEndpoint)
	request := graphQLRequestForOneVault(address.String(), c)
	var response models.TGraphQueryResponseForVault
	if err := client.Run(context.Background(), request, &response); err != nil {
		logs.Error(err)
		c.String(http.StatusInternalServerError, "Impossible to fetch subgraph")
		return
	}

	strategiesCondition := selectStrategiesCondition(c.Query("strategiesCondition"))
	withStrategiesDetails := c.Query("strategiesDetails") == "withDetails"
	withStrategiesRisk := c.Query("strategiesRisk") == "withRisk"
	vaultFromGraph := response.Vault

	data := prepareVaultSchema(
		chainID,
		strategiesCondition,
		withStrategiesRisk,
		withStrategiesDetails,
		vaultFromGraph,
	)
	c.JSON(http.StatusOK, data)
}
