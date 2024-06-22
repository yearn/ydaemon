package vaults

import (
	"context"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/machinebox/graphql"
	"github.com/yearn/ydaemon/common/addresses"
	"github.com/yearn/ydaemon/common/bigNumber"
	"github.com/yearn/ydaemon/common/env"
	"github.com/yearn/ydaemon/common/helpers"
	"github.com/yearn/ydaemon/common/logs"
	"github.com/yearn/ydaemon/common/sort"
	"github.com/yearn/ydaemon/internal/models"
)

func graphQLHarvestRequestForOneVault(vaultAddresses []string, c *gin.Context) *graphql.Request {
	return graphql.NewRequest(`{
		harvests(first: 1000, orderBy: timestamp, orderDirection: desc, where: {vault_in: ["` + strings.Join(vaultAddresses, `", "`) + `"]})
		` + helpers.GetHarvestsForVaults() + `
    }`)
}

// GetHarvestsForVault will, for a given chainID, return a list of all vaults
func (y Controller) GetHarvestsForVault(c *gin.Context) {
	chainID, ok := helpers.AssertChainID(c.Param(`chainID`))
	if !ok {
		c.String(http.StatusBadRequest, `invalid chainID`)
		return
	}
	addressesStr := strings.Split(strings.ToLower(c.Param(`addresses`)), `,`)
	orderBy := helpers.SafeString(getQuery(c, `orderBy`), `timestamp`)
	orderDirection := helpers.SafeString(getQuery(c, `orderDirection`), `desc`)

	graphQLEndpoint := env.CHAINS[chainID].SubgraphURI
	if graphQLEndpoint == `` {
		logs.Error(`No graph endpoint for chainID`, chainID)
		c.String(http.StatusInternalServerError, `impossible to fetch subgraph`)
		return
	}

	client := graphql.NewClient(graphQLEndpoint)
	request := graphQLHarvestRequestForOneVault(addressesStr, c)
	var response models.TGraphQLHarvestRequestForOneVault
	if err := client.Run(context.Background(), request, &response); err != nil {
		logs.Error(err)
		c.String(http.StatusInternalServerError, `invalid graphQL response`)
		return
	}

	// For each harvest from the subgraph, compute our TExternalVaultHarvest structure
	harvests := []TExternalVaultHarvest{}
	harvestsFromSubgraph := response.Harvests
	for _, harvest := range harvestsFromSubgraph {
		tokenPriceBigFloat, _ := buildTokenPrice(
			chainID,
			addresses.ToMixedcase(harvest.Vault.Token.Id),
		)
		decimals := harvest.Vault.Token.Decimals

		if bigNumber.NewInt().SetString(harvest.Profit).IsZero() && bigNumber.NewInt().SetString(harvest.Loss).IsZero() {
			continue
		}
		harvests = append(harvests, TExternalVaultHarvest{
			VaultAddress:    harvest.Vault.Id,
			StrategyAddress: harvest.Strategy.Id,
			TxHash:          harvest.Transaction.Hash,
			Timestamp:       harvest.Timestamp,
			Profit:          harvest.Profit,
			ProfitValue:     buildTVL(bigNumber.NewInt().SetString(harvest.Profit), decimals, tokenPriceBigFloat),
			Loss:            harvest.Loss,
			LossValue:       buildTVL(bigNumber.NewInt().SetString(harvest.Loss), decimals, tokenPriceBigFloat),
		})
	}

	// Sorting the data by timestamp
	sort.SortBy(orderBy, orderDirection, harvests)
	c.JSON(http.StatusOK, harvests)
}
