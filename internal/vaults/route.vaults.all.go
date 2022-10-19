package vaults

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/machinebox/graphql"
	"github.com/yearn/ydaemon/internal/meta"
	"github.com/yearn/ydaemon/internal/utils/env"
	"github.com/yearn/ydaemon/internal/utils/helpers"
	"github.com/yearn/ydaemon/internal/utils/logs"
	"github.com/yearn/ydaemon/internal/utils/models"
	"github.com/yearn/ydaemon/internal/utils/sort"
)

func graphQLRequestForAllVaults(c *gin.Context) *graphql.Request {
	withDetails := helpers.SafeString(c.Query("strategiesDetails"), "noDetails") == "withDetails"
	onlyEndorsed := helpers.SafeString(c.Query("classification"), "endorsed") == "endorsed"

	if onlyEndorsed {
		return graphql.NewRequest(`{
			vaults(where: {classification: Endorsed}, first: 1000) {
				` + helpers.GetGraphRequestVault() + `
				` + helpers.GetGraphRequestStrategies(40, withDetails) + `
			}
		}`)
	}
	return graphql.NewRequest(`{
		vaults(first: 1000) {
			` + helpers.GetGraphRequestVault() + `
			` + helpers.GetGraphRequestStrategies(40, withDetails) + `
		}
    }`)
}

//GetAllVaults will, for a given chainID, return a list of all vaults
func (y Controller) GetAllVaults(c *gin.Context) {
	hideAlways := helpers.StringToBool(helpers.SafeString(c.Query("hideAlways"), "false"))
	orderBy := helpers.SafeString(c.Query("orderBy"), "details.order")
	orderDirection := helpers.SafeString(c.Query("orderDirection"), "asc")
	strategiesCondition := selectStrategiesCondition(c.Query("strategiesCondition"))
	withStrategiesDetails := c.Query("strategiesDetails") == "withDetails"
	withStrategiesRisk := c.Query("strategiesRisk") == "withRisk"
	chainID, ok := helpers.AssertChainID(c.Param("chainID"))
	if !ok {
		c.String(http.StatusBadRequest, "invalid chainID")
		return
	}

	graphQLEndpoint, ok := env.THEGRAPH_ENDPOINTS[chainID]
	if !ok {
		logs.Error("No graph endpoint for chainID", chainID)
		c.String(http.StatusInternalServerError, "impossible to fetch subgraph")
		return
	}

	client := graphql.NewClient(graphQLEndpoint)
	request := graphQLRequestForAllVaults(c)
	var response models.TGraphQueryResponseForVaults
	if err := client.Run(context.Background(), request, &response); err != nil {
		logs.Error(err)
		c.String(http.StatusInternalServerError, "invalid graphQL response")
		return
	}

	data := []TVault{}
	for _, vaultFromGraph := range response.Vaults {
		vaultAddress := vaultFromGraph.Id
		if helpers.ContainsAddress(env.BLACKLISTED_VAULTS[chainID], vaultAddress) {
			continue
		}
		vaultFromMeta, ok := meta.Store.VaultsFromMeta[chainID][vaultAddress]
		if ok && vaultFromMeta.HideAlways && hideAlways {
			continue
		}

		data = append(data, *prepareVaultSchema(
			chainID,
			strategiesCondition,
			withStrategiesRisk,
			withStrategiesDetails,
			vaultFromGraph,
		))
	}

	// Preparing the sort. This specifics steps are needed to avoid a panic
	var sortedData []interface{} = make([]interface{}, len(data))
	for i, d := range data {
		sortedData[i] = d
	}
	sort.SortBy(sortedData, orderBy, orderDirection) //Sort by details.order by default

	c.JSON(http.StatusOK, sortedData)
}
