package apr

import (
	"context"
	"errors"
	"strconv"
	"strings"

	"github.com/machinebox/graphql"
	"github.com/yearn/ydaemon/common/bigNumber"
	"github.com/yearn/ydaemon/common/env"
	"github.com/yearn/ydaemon/common/helpers"
	"github.com/yearn/ydaemon/common/logs"
	"github.com/yearn/ydaemon/internal/models"
)

/**************************************************************************************************
** createGraphQLClient returns a standard GraphQL client for the given endpoint.
** This follows the same pattern as in route.reports.all.go for consistency.
**
** @param endpoint The GraphQL endpoint URL to connect to
** @return *graphql.Client A configured GraphQL client instance
**************************************************************************************************/
var createGraphQLClient = func(endpoint string) *graphql.Client {
	return graphql.NewClient(endpoint)
}

/**************************************************************************************************
** runGraphQLRequest executes a GraphQL request and populates the response.
** This follows the same pattern as in route.reports.all.go for consistency.
**
** @param ctx The context for the request
** @param client The GraphQL client to use
** @param req The GraphQL request to execute
** @param resp The destination for the response
** @return error An error if the request fails, or nil on success
**************************************************************************************************/
var runGraphQLRequest = func(ctx context.Context, client *graphql.Client, req *graphql.Request, resp interface{}) error {
	return client.Run(ctx, req, resp)
}

/**************************************************************************************************
** graphQLRequestForLatestReport creates a GraphQL request to fetch only the latest report
** for a specific strategy.
**
** @param strategyAddress The Ethereum address of the strategy to fetch the report for
** @return *graphql.Request A ready-to-use GraphQL request for the latest report
**************************************************************************************************/
func graphQLRequestForLatestReport(strategyAddress string) *graphql.Request {
	// We only need the first (latest) report
	return graphql.NewRequest(`{
		strategy(id: "` + strings.ToLower(strategyAddress) + `") {
			` + helpers.GetStrategyReports(1) + `
		}
	}`)
}

/**************************************************************************************************
** GetCurrentStrategyAPR retrieves the latest APR for a specific strategy from the subgraph.
** It returns the APR value as a float64, or 0 if the APR cannot be determined.
**
** @param chainID The blockchain network ID
** @param strategyAddress The strategy address to check
** @return float64 The latest APR value or 0 if not available
** @return error An error if the request fails
**************************************************************************************************/
func GetCurrentStrategyAPR(chainID uint64, strategyAddress string) (*bigNumber.Float, error) {
	// Get chain configuration
	chain, ok := env.GetChain(chainID)
	if !ok {
		return bigNumber.NewFloat(0), errors.New("chain not found")
	}

	// Ensure subgraph endpoint is available
	graphQLEndpoint := chain.SubgraphURI
	if graphQLEndpoint == "" {
		logs.Error("No graph endpoint for chainID", chainID)
		return bigNumber.NewFloat(0), errors.New("no graph endpoint for chainID")
	}

	// Create GraphQL client and request
	client := createGraphQLClient(graphQLEndpoint)
	request := graphQLRequestForLatestReport(strategyAddress)

	// Execute the request
	var responseRaw models.TReportsFromGraph
	if err := runGraphQLRequest(context.Background(), client, request, &responseRaw); err != nil {
		logs.Error(err)
		return bigNumber.NewFloat(0), err
	}

	// Process the response - extract the APR from the latest report
	if len(responseRaw.Strategy.Reports) == 0 || len(responseRaw.Strategy.Reports[0].Results) == 0 {
		// No reports or results available
		return bigNumber.NewFloat(0), errors.New("no reports or results available")
	}

	// Get the APR from the latest report's first result
	latestReport := responseRaw.Strategy.Reports[0]
	if len(latestReport.Results) == 0 {
		return bigNumber.NewFloat(0), errors.New("no results available")
	}

	apr, err := strconv.ParseFloat(latestReport.Results[0].Apr, 64)
	if err != nil {
		logs.Error("Error parsing APR value:", err)
		return bigNumber.NewFloat(0), err
	}

	return bigNumber.NewFloat(apr), nil
}
