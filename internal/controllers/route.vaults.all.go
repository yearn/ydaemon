package controllers

import (
	"context"
	"net/http"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"github.com/machinebox/graphql"
	"github.com/majorfi/ydaemon/internal/ethereum"
	"github.com/majorfi/ydaemon/internal/models"
	"github.com/majorfi/ydaemon/internal/store"
	"github.com/majorfi/ydaemon/internal/utils"
)

func graphQLRequestForAllVaults(c *gin.Context) *graphql.Request {
	skip := queryWithFallback(c.Query("skip"), "0")
	first := queryWithFallback(c.Query("first"), "1000")
	orderDirection := queryWithFallback(c.Query("orderDirection"), "desc")
	orderBy := queryWithFallback(c.Query("orderBy"), "activation")

	return graphql.NewRequest(`{
		vaults(skip: ` + skip + `, first: ` + first + `, orderBy: ` + orderBy + `, orderDirection: ` + orderDirection + `) {
			id
			activation
			apiVersion
			classification
			managementFeeBps
			performanceFeeBps
			balanceTokens
			latestUpdate {
				timestamp
			}
			shareToken {
				name
				symbol
				id
				decimals
			}
			token {
				name
				symbol
				id
				decimals
			}
			strategies(first: 40) {
				address
				name
				inQueue
				debtLimit
			}
		}
    }`)
}

//GetAllVaults will, for a given chainID, return a list of all vaults
func (y controller) GetAllVaults(c *gin.Context) {
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
		c.String(http.StatusBadRequest, "invalid chainID")
		return
	}

	data := []*models.TVault{}
	for _, vaultFromGraph := range response.Vaults {
		vaultAddress := common.HexToAddress(vaultFromGraph.Id)
		if utils.ContainsAddress(utils.BLACKLISTED_VAULTS[chainID], vaultAddress) {
			continue
		}

		strategiesCondition := selectStrategiesCondition(c.Query("strategiesCondition"))
		vaultFromMeta := store.VaultsFromMeta[chainID][vaultAddress.String()]
		tokenFromMeta := store.TokensFromMeta[chainID][common.HexToAddress(vaultFromGraph.Token.Id).String()]
		shareTokenFromMeta := store.TokensFromMeta[chainID][vaultAddress.String()]
		apyFromAPIV1 := store.VaultsFromAPIV1[chainID][vaultAddress.String()]
		strategiesFromMeta := store.StrategiesFromMeta[chainID]
		pricesForChainID := store.TokenPrices[chainID]

		data = append(data, prepareVaultSchema(
			chainID,
			strategiesCondition,
			vaultFromGraph,
			vaultFromMeta,
			shareTokenFromMeta,
			tokenFromMeta,
			strategiesFromMeta,
			apyFromAPIV1,
			pricesForChainID,
		))
	}
	c.JSON(http.StatusOK, data)
}
