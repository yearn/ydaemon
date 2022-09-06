package vaultsController

import (
	"context"
	"net/http"
	"sort"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"github.com/machinebox/graphql"
	"github.com/yearn/ydaemon/internal/ethereum"
	"github.com/yearn/ydaemon/internal/logs"
	"github.com/yearn/ydaemon/internal/models"
	"github.com/yearn/ydaemon/internal/store"
	"github.com/yearn/ydaemon/internal/utils"
)

func graphQLRequestForAllVaults(c *gin.Context) *graphql.Request {
	skip := utils.ValueWithFallback(c.Query("skip"), "0")
	first := utils.ValueWithFallback(c.Query("first"), "1000")
	orderDirection := utils.ValueWithFallback(c.Query("orderDirection"), "desc")
	orderBy := utils.ValueWithFallback(c.Query("orderBy"), "activation")
	withDetails := utils.ValueWithFallback(c.Query("strategiesDetails"), "noDetails") == "withDetails"
	onlyEndorsed := utils.ValueWithFallback(c.Query("classification"), "endorsed") == "endorsed"

	if onlyEndorsed {
		return graphql.NewRequest(`{
			vaults(where: {classification: Endorsed}, skip: ` + skip + `, first: ` + first + `, orderBy: ` + orderBy + `, orderDirection: ` + orderDirection + `) {
				` + utils.GetGraphRequestVault() + `
				` + utils.GetGraphRequestStrategies(40, withDetails) + `
			}
		}`)
	}
	return graphql.NewRequest(`{
		vaults(skip: ` + skip + `, first: ` + first + `, orderBy: ` + orderBy + `, orderDirection: ` + orderDirection + `) {
			` + utils.GetGraphRequestVault() + `
			` + utils.GetGraphRequestStrategies(40, withDetails) + `
		}
    }`)
}

//GetAllVaults will, for a given chainID, return a list of all vaults
func (y Controller) GetAllVaults(c *gin.Context) {
	// get the chainID from the URI
	chainID, err := strconv.ParseUint(c.Param("chainID"), 10, 64)
	if err != nil {
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

	data := []*models.TVault{}
	for _, vaultFromGraph := range response.Vaults {
		vaultAddress := common.HexToAddress(vaultFromGraph.Id)
		if utils.ContainsAddress(utils.BLACKLISTED_VAULTS[chainID], vaultAddress) {
			continue
		}
		vaultFromMeta, ok := store.VaultsFromMeta[chainID][vaultAddress]
		if ok && vaultFromMeta.HideAlways {
			continue
		}

		strategiesCondition := selectStrategiesCondition(c.Query("strategiesCondition"))
		withStrategiesDetails := c.Query("strategiesDetails") == "withDetails"
		withStrategiesRisk := c.Query("strategiesRisk") == "withRisk"

		data = append(data, prepareVaultSchema(
			chainID,
			strategiesCondition,
			withStrategiesRisk,
			withStrategiesDetails,
			vaultFromGraph,
		))
	}

	sort.Slice(data, func(i, j int) bool {
		return data[i].Details.Order < data[j].Details.Order
	})

	c.JSON(http.StatusOK, data)
}
