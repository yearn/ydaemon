package vaults

import (
	"context"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/machinebox/graphql"
	"github.com/yearn/ydaemon/internal/utils/env"
	"github.com/yearn/ydaemon/internal/utils/helpers"
	"github.com/yearn/ydaemon/internal/utils/logs"
	"github.com/yearn/ydaemon/internal/utils/sort"
	"github.com/yearn/ydaemon/internal/utils/types/bigNumber"
	"github.com/yearn/ydaemon/internal/utils/types/common"
)

type TGraphQLHarvestRequestForOneVault struct {
	Harvests []struct {
		Transaction struct {
			Hash string `json:"hash"`
		}
		Vault struct {
			Id    string `json:"id"`
			Token struct {
				Id       string `json:"id"`
				Decimals int    `json:"decimals"`
			}
		}
		Timestamp string `json:"timestamp"`
		Profit    string `json:"profit"`
		Loss      string `json:"loss"`
	} `json:"harvests"`
}

type TVaultHarvest struct {
	VaultAddress string  `json:"vaultAddress"`
	TxHash       string  `json:"txHash"`
	Timestamp    string  `json:"timestamp"`
	Profit       string  `json:"profit"`
	ProfitValue  float64 `json:"profitValue"`
	Loss         string  `json:"loss"`
	LossValue    float64 `json:"lossValue"`
}

func graphQLHarvestRequestForOneVault(vaultAddresses []string, c *gin.Context) *graphql.Request {
	return graphql.NewRequest(`{
		harvests(where: {vault_in: ["` + strings.Join(vaultAddresses, `", "`) + `"]}) {
			vault {
				id
				token {
					id
					decimals
				}
			}
			transaction {
				hash
			}
			timestamp
			profit
			loss
		}
	}`)
}

//GetHarvestForVault will, for a given chainID, return a list of all vaults
func (y Controller) GetHarvestForVault(c *gin.Context) {
	chainID, ok := helpers.AssertChainID(c.Param("chainID"))
	if !ok {
		c.String(http.StatusBadRequest, "invalid chainID")
		return
	}
	addressesStr := strings.Split(strings.ToLower(c.Param("addresses")), ",")

	graphQLEndpoint, ok := env.THEGRAPH_ENDPOINTS[chainID]
	if !ok {
		logs.Error("No graph endpoint for chainID", chainID)
		c.String(http.StatusInternalServerError, "Impossible to fetch subgraph")
		return
	}

	client := graphql.NewClient(graphQLEndpoint)
	request := graphQLHarvestRequestForOneVault(addressesStr, c)
	var response TGraphQLHarvestRequestForOneVault
	if err := client.Run(context.Background(), request, &response); err != nil {
		logs.Error(err)
		c.String(http.StatusInternalServerError, "Impossible to fetch subgraph")
		return
	}

	harvests := []TVaultHarvest{}
	harvestsFromSubgraph := response.Harvests
	for _, harvest := range harvestsFromSubgraph {
		tokenPriceBigFloat, _ := buildTokenPrice(chainID, common.HexToAddress(harvest.Vault.Token.Id))
		decimals := harvest.Vault.Token.Decimals
		if bigNumber.NewInt().SetString(harvest.Profit).IsZero() && bigNumber.NewInt().SetString(harvest.Loss).IsZero() {
			continue
		}
		harvests = append(harvests, TVaultHarvest{
			VaultAddress: harvest.Vault.Id,
			TxHash:       harvest.Transaction.Hash,
			Timestamp:    harvest.Timestamp,
			Profit:       harvest.Profit,
			ProfitValue:  buildTVL(bigNumber.NewInt().SetString(harvest.Profit), decimals, tokenPriceBigFloat),
			Loss:         harvest.Loss,
			LossValue:    buildTVL(bigNumber.NewInt().SetString(harvest.Loss), decimals, tokenPriceBigFloat),
		})
	}

	var sortedData []interface{} = make([]interface{}, len(harvests))
	for i, d := range harvests {
		sortedData[i] = d
	}
	sort.SortBy(sortedData, `Timestamp`, `desc`)

	c.JSON(http.StatusOK, sortedData)
}
