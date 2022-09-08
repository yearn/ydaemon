package strategies

import (
	"context"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/machinebox/graphql"
	"github.com/yearn/ydaemon/internal/utils/ethereum"
	"github.com/yearn/ydaemon/internal/utils/helpers"
	"github.com/yearn/ydaemon/internal/utils/logs"
	"github.com/yearn/ydaemon/internal/utils/models"
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
	chainID, ok := helpers.AssertChainID(c, c.Param("chainID"))
	if !ok {
		return
	}
	address, ok := helpers.AssertAddress(c, c.Param("address"), chainID)
	if !ok {
		return
	}

	client := graphql.NewClient(ethereum.GetGraphURI(chainID))
	request := graphQLRequestForReports(address.String(), c)
	var response models.TReportsFromGraph
	if err := client.Run(context.Background(), request, &response); err != nil {
		logs.Error(err)
		c.String(http.StatusInternalServerError, "Impossible to fetch subgraph")
		return
	}

	c.JSON(http.StatusOK, response.Strategy.Reports)
}
