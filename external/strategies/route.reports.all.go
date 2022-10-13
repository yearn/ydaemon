package strategies

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

func graphQLRequestForReports(strategyAddress string, c *gin.Context) *graphql.Request {
	return graphql.NewRequest(`{
		strategy(id: "` + strings.ToLower(strategyAddress) + `") {
			` + helpers.GetStrategyReports() + `
		}
	}`)
}

//GetReports will, for a given strategy on a given chainID, return a list of reports
func (y Controller) GetReports(c *gin.Context) {
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
	request := graphQLRequestForReports(address.String(), c)
	var response models.TReportsFromGraph
	if err := client.Run(context.Background(), request, &response); err != nil {
		logs.Error(err)
		c.String(http.StatusInternalServerError, "Impossible to fetch subgraph")
		return
	}

	c.JSON(http.StatusOK, response.Strategy.Reports)
}
