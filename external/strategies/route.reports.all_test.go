package strategies

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/machinebox/graphql"
	"github.com/stretchr/testify/assert"
	"github.com/yearn/ydaemon/common/bigNumber"
	"github.com/yearn/ydaemon/common/env"
	"github.com/yearn/ydaemon/internal/models"
)

// mockGraphQLClient is a mock implementation for testing
type mockGraphQLClient struct {
	mockResponse      models.TReportsFromGraph
	shouldReturnError bool
}

// mockResponse generates test data for strategy reports
func generateMockResponse() models.TReportsFromGraph {
	return models.TReportsFromGraph{
		Strategy: struct {
			Reports []struct {
				ID        string         `json:"id"`
				DebtAdded *bigNumber.Int `json:"debtAdded"`
				DebtLimit *bigNumber.Int `json:"debtLimit"`
				TotalDebt *bigNumber.Int `json:"totalDebt"`
				Gain      *bigNumber.Int `json:"gain"`
				TotalGain *bigNumber.Int `json:"totalGain"`
				Loss      *bigNumber.Int `json:"loss"`
				TotalLoss *bigNumber.Int `json:"totalLoss"`
				DebtPaid  *bigNumber.Int `json:"debtPaid"`
				Timestamp string         `json:"timestamp"`
				Results   []struct {
					Duration   string `json:"duration"`
					DurationPr string `json:"durationPR"`
					Apr        string `json:"APR"`
				} `json:"results"`
			} `json:"reports"`
		}{
			Reports: []struct {
				ID        string         `json:"id"`
				DebtAdded *bigNumber.Int `json:"debtAdded"`
				DebtLimit *bigNumber.Int `json:"debtLimit"`
				TotalDebt *bigNumber.Int `json:"totalDebt"`
				Gain      *bigNumber.Int `json:"gain"`
				TotalGain *bigNumber.Int `json:"totalGain"`
				Loss      *bigNumber.Int `json:"loss"`
				TotalLoss *bigNumber.Int `json:"totalLoss"`
				DebtPaid  *bigNumber.Int `json:"debtPaid"`
				Timestamp string         `json:"timestamp"`
				Results   []struct {
					Duration   string `json:"duration"`
					DurationPr string `json:"durationPR"`
					Apr        string `json:"APR"`
				} `json:"results"`
			}{
				{
					ID:        "report-1",
					DebtAdded: bigNumber.NewInt().SetString("1000000000000000000"),
					DebtLimit: bigNumber.NewInt().SetString("10000000000000000000"),
					TotalDebt: bigNumber.NewInt().SetString("5000000000000000000"),
					Gain:      bigNumber.NewInt().SetString("500000000000000000"),
					TotalGain: bigNumber.NewInt().SetString("2000000000000000000"),
					Loss:      bigNumber.NewInt().SetString("100000000000000000"),
					TotalLoss: bigNumber.NewInt().SetString("300000000000000000"),
					DebtPaid:  bigNumber.NewInt().SetString("800000000000000000"),
					Timestamp: "1630000000",
					Results: []struct {
						Duration   string `json:"duration"`
						DurationPr string `json:"durationPR"`
						Apr        string `json:"APR"`
					}{
						{
							Duration:   "604800",
							DurationPr: "86400",
							Apr:        "0.05",
						},
					},
				},
				{
					ID:        "report-2",
					DebtAdded: bigNumber.NewInt().SetString("2000000000000000000"),
					DebtLimit: bigNumber.NewInt().SetString("10000000000000000000"),
					TotalDebt: bigNumber.NewInt().SetString("7000000000000000000"),
					Gain:      bigNumber.NewInt().SetString("700000000000000000"),
					TotalGain: bigNumber.NewInt().SetString("2700000000000000000"),
					Loss:      bigNumber.NewInt().SetString("50000000000000000"),
					TotalLoss: bigNumber.NewInt().SetString("350000000000000000"),
					DebtPaid:  bigNumber.NewInt().SetString("500000000000000000"),
					Timestamp: "1631000000",
					Results: []struct {
						Duration   string `json:"duration"`
						DurationPr string `json:"durationPR"`
						Apr        string `json:"APR"`
					}{
						{
							Duration:   "604800",
							DurationPr: "86400",
							Apr:        "0.06",
						},
					},
				},
			},
		},
	}
}

/**************************************************************************************************
** TestGetReports tests the GetReports handler which retrieves reports for a specific strategy on
** a specific chain. This test validates:
** - The handler correctly validates the chain ID and strategy address
** - The handler correctly processes GraphQL responses
** - Error conditions are properly handled
**
** We use a custom handler that mimics the behavior of the real one but uses mock data.
**************************************************************************************************/
func TestGetReports(t *testing.T) {
	// Setup test environment
	gin.SetMode(gin.TestMode)

	// Set up a temporary environment with mock chains
	setupTestEnvironment := func() {
		// Store original chains and restore them after the test
		originalChains := env.CHAINS
		env.CHAINS = map[uint64]env.TChain{
			1: {
				ID:          1,
				SubgraphURI: "https://api.thegraph.com/subgraphs/name/yearn/yearn-vaults-v2-ethereum",
			},
			999: {
				ID:          999,
				SubgraphURI: "", // Empty subgraph URI for testing error case
			},
		}

		// Cleanup function
		t.Cleanup(func() {
			env.CHAINS = originalChains
		})
	}

	// Initialize the test environment
	setupTestEnvironment()

	// Create test cases
	tests := []struct {
		name           string
		chainID        string
		address        string
		expectedStatus int
		setupMock      func()
		validateJSON   func(t *testing.T, body []byte)
	}{
		{
			name:           "Valid request",
			chainID:        "1",
			address:        "0x16388463d60FFE0661Cf7F1f31a7D658aC790ff7",
			expectedStatus: http.StatusOK,
			setupMock: func() {
				// Store original functions and restore them after the test
				originalClientCreator := createGraphQLClient
				originalRequestRunner := runGraphQLRequest

				// Mock the GraphQL client creation
				createGraphQLClient = func(endpoint string) *graphql.Client {
					// Return a dummy client, it won't be used directly
					return graphql.NewClient(endpoint)
				}

				// Mock the GraphQL request execution
				runGraphQLRequest = func(ctx context.Context, client *graphql.Client, req *graphql.Request, resp interface{}) error {
					// Cast the response to the expected type
					reportsResp, ok := resp.(*models.TReportsFromGraph)
					if !ok {
						return fmt.Errorf("unexpected response type")
					}

					// Populate with mock data
					mockData := generateMockResponse()
					*reportsResp = mockData
					return nil
				}

				// Cleanup function
				t.Cleanup(func() {
					createGraphQLClient = originalClientCreator
					runGraphQLRequest = originalRequestRunner
				})
			},
			validateJSON: func(t *testing.T, body []byte) {
				var response []models.TReport
				err := json.Unmarshal(body, &response)
				assert.NoError(t, err, "Failed to unmarshal response")

				// Verify length
				assert.Len(t, response, 2, "Response should contain 2 reports")

				// Only validate if there are items in the response
				if len(response) >= 2 {
					// Verify specific values from the first report
					assert.Equal(t, "report-1", response[0].ID, "First report should have ID 'report-1'")
					assert.Equal(t, uint64(1630000000), response[0].Timestamp, "Timestamp should be correctly parsed")

					// Verify results were correctly parsed
					assert.GreaterOrEqual(t, len(response[0].Results), 1, "First report should have at least 1 result")
					if len(response[0].Results) > 0 {
						assert.Equal(t, 0.05, response[0].Results[0].APR, "APR should be correctly parsed")
						assert.Equal(t, uint64(604800), response[0].Results[0].Duration, "Duration should be correctly parsed")
					}
				}
			},
		},
		{
			name:           "Invalid chain ID",
			chainID:        "invalid",
			address:        "0x16388463d60FFE0661Cf7F1f31a7D658aC790ff7",
			expectedStatus: http.StatusBadRequest,
			setupMock:      func() {},
		},
		{
			name:           "Invalid address",
			chainID:        "1",
			address:        "invalid",
			expectedStatus: http.StatusBadRequest,
			setupMock:      func() {},
		},
		{
			name:           "GraphQL error",
			chainID:        "1",
			address:        "0x16388463d60FFE0661Cf7F1f31a7D658aC790ff7",
			expectedStatus: http.StatusInternalServerError,
			setupMock: func() {
				// Store original functions and restore them after the test
				originalClientCreator := createGraphQLClient
				originalRequestRunner := runGraphQLRequest

				// Mock the GraphQL client creation
				createGraphQLClient = func(endpoint string) *graphql.Client {
					// Return a dummy client, it won't be used directly
					return graphql.NewClient(endpoint)
				}

				// Mock the GraphQL request execution with an error
				runGraphQLRequest = func(ctx context.Context, client *graphql.Client, req *graphql.Request, resp interface{}) error {
					return fmt.Errorf("simulated GraphQL error")
				}

				// Cleanup function
				t.Cleanup(func() {
					createGraphQLClient = originalClientCreator
					runGraphQLRequest = originalRequestRunner
				})
			},
		},
		{
			name:           "Missing subgraph URI",
			chainID:        "999", // Using our test chain with no subgraph URI
			address:        "0x16388463d60FFE0661Cf7F1f31a7D658aC790ff7",
			expectedStatus: http.StatusInternalServerError,
			setupMock:      func() {},
		},
	}

	// For each test case
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Setup mocks
			tc.setupMock()

			// Create a new test router
			router := gin.New()

			// Register the handler
			controller := Controller{}
			router.GET("/strategies/:chainID/:address/reports", controller.GetReports)

			// Create request
			req, _ := http.NewRequest("GET", "/strategies/"+tc.chainID+"/"+tc.address+"/reports", nil)
			w := httptest.NewRecorder()

			// Perform the request
			router.ServeHTTP(w, req)

			// Check status code
			assert.Equal(t, tc.expectedStatus, w.Code, "Status code should match expected")

			// Validate JSON response if needed
			if tc.validateJSON != nil && w.Code == http.StatusOK {
				tc.validateJSON(t, w.Body.Bytes())
			}
		})
	}
}
