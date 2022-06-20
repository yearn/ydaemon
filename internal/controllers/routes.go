package controllers

import (
	"context"
	"net/http"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"github.com/machinebox/graphql"
	"github.com/majorfi/ydaemon/internal/ethereum"
	"github.com/majorfi/ydaemon/internal/logs"
	"github.com/majorfi/ydaemon/internal/models"
	"github.com/majorfi/ydaemon/internal/store"
)

type controller struct{}

//DoSomething will do something
func (y controller) DoSomething(c *gin.Context) {
	// get the chainID from the URI
	chainID, err := strconv.ParseUint(c.Param("chainID"), 10, 64)
	if err != nil {
		chainID = 1
		c.String(http.StatusBadRequest, "Invalid chainID")
		return
	}

	client := graphql.NewClient(ethereum.GetGraphURI(chainID))
	request := graphql.NewRequest(`
        {
			vaults(first: 1000) {
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
        }
    `)
	var response models.TGraphQueryResponseForVaults
	if err := client.Run(context.Background(), request, &response); err != nil {
		logs.Error(err)
		c.String(http.StatusBadRequest, "Invalid chainID")
		return
	}

	data := []*models.TVault{}
	for _, vaultFromGraph := range response.Vaults {
		vaultFromMeta := store.VaultsFromMeta[chainID][common.HexToAddress(vaultFromGraph.Id).String()]
		tokenFromMeta := store.TokensFromMeta[chainID][common.HexToAddress(vaultFromGraph.Token.Id).String()]
		apyFromAPIV1 := store.VaultsFromAPIV1[chainID][common.HexToAddress(vaultFromGraph.Id).String()]
		strategiesFromMeta := store.StrategiesFromMeta[chainID]
		pricesForChainID := store.TokenPrices[chainID]

		data = append(data, prepareVaultSchema(
			vaultFromGraph,
			vaultFromMeta,
			tokenFromMeta,
			strategiesFromMeta,
			apyFromAPIV1,
			pricesForChainID,
		))
	}
	c.JSON(http.StatusOK, data)
}
