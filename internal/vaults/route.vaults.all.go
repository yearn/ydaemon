package vaults

import (
	"context"
	"net/http"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"github.com/machinebox/graphql"
	"github.com/yearn/ydaemon/internal/meta"
	"github.com/yearn/ydaemon/internal/sort"
	"github.com/yearn/ydaemon/internal/utils/ethereum"
	"github.com/yearn/ydaemon/internal/utils/helpers"
	"github.com/yearn/ydaemon/internal/utils/logs"
	"github.com/yearn/ydaemon/internal/utils/models"
)

func graphQLRequestForAllVaults(c *gin.Context) *graphql.Request {
	withDetails := helpers.ValueWithFallback(c.Query("strategiesDetails"), "noDetails") == "withDetails"
	onlyEndorsed := helpers.ValueWithFallback(c.Query("classification"), "endorsed") == "endorsed"

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
	chainID, ok := helpers.AssertChainID(c.Param("chainID"))
	if !ok {
		c.String(http.StatusBadRequest, "invalid chainID")
		return
	}

	client := graphql.NewClient(ethereum.GetGraphURI(chainID))
	request := graphQLRequestForAllVaults(c)
	var response models.TGraphQueryResponseForVaults
	if err := client.Run(context.Background(), request, &response); err != nil {
		logs.Error(err)
		c.String(http.StatusBadRequest, "invalid graphQL response")
		return
	}

	data := []TVault{}
	for _, vaultFromGraph := range response.Vaults {
		vaultAddress := common.HexToAddress(vaultFromGraph.Id)
		if helpers.ContainsAddress(helpers.BLACKLISTED_VAULTS[chainID], vaultAddress) {
			continue
		}
		vaultFromMeta, ok := meta.Store.VaultsFromMeta[chainID][vaultAddress]
		if ok && vaultFromMeta.HideAlways {
			continue
		}

		strategiesCondition := selectStrategiesCondition(c.Query("strategiesCondition"))
		withStrategiesDetails := c.Query("strategiesDetails") == "withDetails"
		withStrategiesRisk := c.Query("strategiesRisk") == "withRisk"

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
	orderBy := helpers.ValueWithFallback(c.Query("orderBy"), "details.order")
	orderDirection := helpers.ValueWithFallback(c.Query("orderDirection"), "asc")
	sort.SortBy(orderBy, orderDirection, sortedData) //Sort by details.order by default

	c.JSON(http.StatusOK, sortedData)
}
