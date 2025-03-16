package strategies

import (
	"context"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/machinebox/graphql"
	"github.com/yearn/ydaemon/common/env"
	"github.com/yearn/ydaemon/common/helpers"
	"github.com/yearn/ydaemon/common/logs"
	"github.com/yearn/ydaemon/internal/models"
)

/**************************************************************************************************
** graphQLClientCreator is a function type that creates a GraphQL client for a given endpoint.
** This type definition enables dependency injection for testing by allowing the creation of
** mock clients in tests.
**
** @param endpoint The GraphQL endpoint URL to connect to
** @return *graphql.Client A configured GraphQL client instance
**************************************************************************************************/
type graphQLClientCreator func(endpoint string) *graphql.Client

/**************************************************************************************************
** createGraphQLClient is the default implementation of graphQLClientCreator that returns a
** standard GraphQL client. This variable can be replaced during testing to inject mock clients.
**
** The separation of interface (type) from implementation (variable) follows the dependency
** inversion principle, making the code more testable.
**************************************************************************************************/
var createGraphQLClient graphQLClientCreator = func(endpoint string) *graphql.Client {
	return graphql.NewClient(endpoint)
}

/**************************************************************************************************
** graphQLRunner is a function type that executes a GraphQL request and populates the response.
** This abstraction allows tests to intercept and mock the execution of GraphQL requests
** without requiring actual network calls.
**
** @param ctx The context for the request, which can carry deadlines and cancellation signals
** @param client The GraphQL client to use for the request
** @param req The GraphQL request to execute
** @param resp The destination where the response will be stored
** @return error An error if the request fails, or nil on success
**************************************************************************************************/
type graphQLRunner func(ctx context.Context, client *graphql.Client, req *graphql.Request, resp interface{}) error

/**************************************************************************************************
** runGraphQLRequest is the default implementation of graphQLRunner that executes a real GraphQL
** request. This variable can be replaced during testing to provide mock responses instead of
** making actual network calls.
**
** This approach allows for complete isolation of the handler logic during testing, ensuring
** tests are fast, reliable, and don't depend on external services.
**************************************************************************************************/
var runGraphQLRequest graphQLRunner = func(ctx context.Context, client *graphql.Client, req *graphql.Request, resp interface{}) error {
	return client.Run(ctx, req, resp)
}

/**************************************************************************************************
** graphQLRequestForReports creates a GraphQL request to fetch reports for a specific strategy.
** This function encapsulates the GraphQL query structure, making it easier to maintain and test.
**
** @param strategyAddress The Ethereum address of the strategy to fetch reports for
** @param c The Gin context containing additional request parameters
** @return *graphql.Request A ready-to-use GraphQL request
**************************************************************************************************/
func graphQLRequestForReports(strategyAddress string, c *gin.Context) *graphql.Request {
	return graphql.NewRequest(`{
		strategy(id: "` + strings.ToLower(strategyAddress) + `") {
			` + helpers.GetStrategyReports() + `
		}
	}`)
}

/**************************************************************************************************
** GetReports retrieves all reports for a specific strategy on a specific blockchain network.
** Reports contain detailed information about strategy performance including gains, losses, debt
** levels, and calculated APR.
**
** The function performs validation on:
** - Chain ID (must be a valid positive integer)
** - Strategy address (must be a valid Ethereum address for the specified chain)
** - Subgraph availability for the specified chain
**
** If the request is valid and successful, the function returns a JSON array of reports with
** detailed performance metrics. In case of validation errors or GraphQL failures, appropriate
** HTTP error responses are returned.
**
** @param c The Gin context containing request parameters
** - chainID: Path parameter specifying the blockchain network ID
** - address: Path parameter specifying the strategy address
**************************************************************************************************/
func (y Controller) GetReports(c *gin.Context) {
	// Validate chain ID
	chainID, ok := helpers.AssertChainID(c.Param("chainID"))
	if !ok {
		c.String(http.StatusBadRequest, "invalid chainID")
		return
	}

	// Validate strategy address
	address, ok := helpers.AssertAddress(c.Param("address"), chainID)
	if !ok {
		c.String(http.StatusBadRequest, "invalid address")
		return
	}

	// Get chain configuration
	chain, ok := env.GetChain(chainID)
	if !ok {
		return
	}

	// Ensure subgraph endpoint is available
	graphQLEndpoint := chain.SubgraphURI
	if graphQLEndpoint == "" {
		logs.Error("No graph endpoint for chainID", chainID)
		c.String(http.StatusInternalServerError, "impossible to fetch subgraph")
		return
	}

	// Create GraphQL client and request
	client := createGraphQLClient(graphQLEndpoint)
	request := graphQLRequestForReports(address.Hex(), c)

	// Execute the request
	var responseRaw models.TReportsFromGraph
	if err := runGraphQLRequest(context.Background(), client, request, &responseRaw); err != nil {
		logs.Error(err)
		c.String(http.StatusInternalServerError, "invalid graphQL response")
		return
	}

	// Process and format the response
	var response models.TReportsFromGraphToClient
	response.Strategy.Reports = []models.TReport{}
	for _, report := range responseRaw.Strategy.Reports {
		timestamp, err := strconv.ParseUint(report.Timestamp, 10, 64)
		if err != nil {
			timestamp = 0
		}
		newReport := models.TReport{
			ID:        report.ID,
			DebtAdded: report.DebtAdded,
			DebtLimit: report.DebtLimit,
			TotalDebt: report.TotalDebt,
			Gain:      report.Gain,
			TotalGain: report.TotalGain,
			Loss:      report.Loss,
			TotalLoss: report.TotalLoss,
			DebtPaid:  report.DebtPaid,
			Timestamp: timestamp,
			Results:   []models.TReportResult{},
		}
		for _, allResults := range report.Results {
			duration, err := strconv.ParseUint(allResults.Duration, 10, 64)
			if err != nil {
				duration = 0
			}
			durationPr, err := strconv.ParseUint(allResults.DurationPr, 10, 64)
			if err != nil {
				durationPr = 0
			}
			apr, err := strconv.ParseFloat(allResults.Apr, 64)
			if err != nil {
				apr = 0
			}

			newResult := models.TReportResult{
				Duration:   duration,
				DurationPR: durationPr,
				APR:        apr,
			}
			newReport.Results = append(newReport.Results, newResult)
		}
		response.Strategy.Reports = append(response.Strategy.Reports, newReport)
	}

	c.JSON(http.StatusOK, response.Strategy.Reports)
}
