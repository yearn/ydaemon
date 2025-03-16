package vaults

import (
	"context"
	"net/http"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"github.com/machinebox/graphql"
	"github.com/yearn/ydaemon/common/bigNumber"
	"github.com/yearn/ydaemon/common/env"
	"github.com/yearn/ydaemon/common/helpers"
	"github.com/yearn/ydaemon/common/logs"
	"github.com/yearn/ydaemon/common/sort"
	"github.com/yearn/ydaemon/internal/models"
)

/**************************************************************************************************
** graphQLHarvestRequestForOneVault creates a GraphQL query to fetch harvest events for vaults.
**
** This helper function builds a GraphQL query that retrieves harvest events for a specific set
** of vault addresses. Harvest events represent when strategies report profits or losses back
** to their associated vaults.
**
** The query limits results to the most recent 1000 harvests and includes detailed information
** about each harvest event, including vault and strategy details, transaction information,
** timestamps, and profit/loss figures.
**
** @param vaultAddresses []string - Array of vault addresses to fetch harvests for
** @param c *gin.Context - The Gin context (for potential parameter access)
** @return *graphql.Request - The constructed GraphQL request object ready to be executed
**************************************************************************************************/
func graphQLHarvestRequestForOneVault(vaultAddresses []string, c *gin.Context) *graphql.Request {
	return graphql.NewRequest(`{
		harvests(first: 1000, orderBy: timestamp, orderDirection: desc, where: {vault_in: ["` + strings.Join(vaultAddresses, `", "`) + `"]})
		` + helpers.GetHarvestsForVaults() + `
    }`)
}

/**************************************************************************************************
** GetHarvestsForVault retrieves harvest events for specified vaults.
**
** This endpoint returns a history of harvest events for the specified vaults on a particular
** chain. Harvests represent profit-reporting events where strategies report gains or losses
** back to their vaults. This information is valuable for analyzing vault performance over time.
**
** The endpoint supports the following parameters:
** - chainID: The chain ID from the URL path parameter
** - addresses: Comma-separated list of vault addresses from the URL path parameter
** - orderBy: Field to sort results by (default: 'timestamp')
** - orderDirection: Sort direction, 'asc' or 'desc' (default: 'desc')
**
** The function processes data through the following steps:
** 1. Validates the chain ID and retrieves sorting parameters
** 2. Ensures the chain has a valid subgraph URI for querying
** 3. Builds and executes a GraphQL query to fetch harvest events
** 4. Processes each harvest event, calculating USD values for profits/losses
** 5. Filters out harvests with zero profit and loss
** 6. Sorts the results according to the specified order parameters
**
** Endpoint: GET /vaults/:chainID/harvests/:addresses
**
** @param c *gin.Context - The Gin context containing the HTTP request
** @return void - Response is sent directly via Gin with the list of harvest events
**************************************************************************************************/
func (y Controller) GetHarvestsForVault(c *gin.Context) {
	chainID, ok := helpers.AssertChainID(c.Param(`chainID`))
	if !ok {
		c.String(http.StatusBadRequest, `invalid chainID`)
		return
	}
	addressesStr := strings.Split(strings.ToLower(c.Param(`addresses`)), `,`)
	orderBy := helpers.SafeString(getQuery(c, `orderBy`), `timestamp`)
	orderDirection := helpers.SafeString(getQuery(c, `orderDirection`), `desc`)

	chain, ok := env.GetChain(chainID)
	if !ok {
		return
	}
	graphQLEndpoint := chain.SubgraphURI
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
		tokenPriceBigFloat, _ := buildTokenPrice(chainID, common.HexToAddress(harvest.Vault.Token.Id))
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
