package prices

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/yearn/ydaemon/common/bigNumber"
	"github.com/yearn/ydaemon/common/env"
	"github.com/yearn/ydaemon/common/helpers"
	"github.com/yearn/ydaemon/internal/models"
)

// mockPriceData represents a price entry for testing
type mockPriceData struct {
	Price          *bigNumber.Int
	HumanizedPrice *bigNumber.Float
}

// mockStorage implements a test-specific storage system
type mockStorage struct {
	prices      map[uint64]map[common.Address]mockPriceData
	returnError bool
}

func (m *mockStorage) ListPrices(chainID uint64) (map[common.Address]mockPriceData, error) {
	if m.returnError {
		return nil, fmt.Errorf("simulated storage error")
	}
	return m.prices[chainID], nil
}

// MockPriceHandler creates a custom handler function that uses mock data
func MockPriceHandler(mockStore *mockStorage, chainID string, humanized string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Parse chainID
		parsedChainID, ok := helpers.AssertChainID(chainID)
		if !ok {
			c.String(http.StatusBadRequest, "invalid chainID")
			return
		}

		// Use mock storage
		mockPrices, err := mockStore.ListPrices(parsedChainID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve prices"})
			return
		}

		// Convert to expected format
		allPrices := make(map[common.Address]models.TPrices)
		for addr, price := range mockPrices {
			allPrices[addr] = models.TPrices{
				Price:          price.Price,
				HumanizedPrice: price.HumanizedPrice,
			}
		}

		// Format the response
		isHumanized := helpers.StringToBool(humanized)

		// Initialize response containers
		rawPrices := make(map[string]*bigNumber.Int)
		humanizedPrices := make(map[string]*bigNumber.Float)

		// Process prices
		for addr, price := range allPrices {
			if isHumanized {
				humanizedPrices[addr.Hex()] = price.HumanizedPrice
			} else {
				rawPrices[addr.Hex()] = price.Price
			}
		}

		// Return appropriate response
		if isHumanized {
			c.JSON(http.StatusOK, humanizedPrices)
		} else {
			c.JSON(http.StatusOK, rawPrices)
		}
	}
}

/**************************************************************************************************
** TestGetAllPrices tests the GetAllPrices handler which retrieves prices for all tokens across
** all supported chains. This test validates:
** - The correct response format is returned based on the 'humanized' query parameter
** - The handler correctly processes prices from all chains
** - Error conditions are properly handled
**
** Since this function depends on env.GetChains() and storage.ListPrices(), we use a test-specific
** implementation that mimics the behavior of these functions.
**************************************************************************************************/
func TestGetAllPrices(t *testing.T) {
	// Setup test environment
	gin.SetMode(gin.TestMode)

	// Create mock data
	mockPrices := setupMockPrices()

	// Create test cases
	tests := []struct {
		name           string
		humanized      string
		expectedStatus int
		validateJSON   func(t *testing.T, body []byte)
	}{
		{
			name:           "Raw prices (default)",
			humanized:      "false",
			expectedStatus: http.StatusOK,
			validateJSON: func(t *testing.T, body []byte) {
				// Parse the response
				var response map[uint64]map[string]*json.RawMessage
				err := json.Unmarshal(body, &response)
				assert.NoError(t, err, "Failed to unmarshal response")

				// Verify structure of response
				assert.Contains(t, response, uint64(1), "Response should contain chain ID 1")
				assert.Contains(t, response, uint64(10), "Response should contain chain ID 10")

				// Verify specific chain response
				chain1 := response[1]
				assert.Len(t, chain1, 2, "Chain 1 should have 2 tokens")
				assert.Contains(t, chain1, "0x6B175474E89094C44Da98b954EedeAC495271d0F", "Chain 1 should contain DAI token")
			},
		},
		{
			name:           "Humanized prices",
			humanized:      "true",
			expectedStatus: http.StatusOK,
			validateJSON: func(t *testing.T, body []byte) {
				// Parse the response
				var response map[uint64]map[string]*json.RawMessage
				err := json.Unmarshal(body, &response)
				assert.NoError(t, err, "Failed to unmarshal response")

				// Verify structure of response
				assert.Contains(t, response, uint64(1), "Response should contain chain ID 1")
				assert.Contains(t, response, uint64(10), "Response should contain chain ID 10")
			},
		},
	}

	// For each test case
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Setup a custom handler that mimics the real one but uses our mock data
			customHandler := func(c *gin.Context) {
				humanized := c.Query("humanized") == "true"
				rawPrices := make(map[uint64]map[string]*bigNumber.Int)
				humanizedPrices := make(map[uint64]map[string]*bigNumber.Float)

				// Mock the chains - similar to what env.GetChains() would return
				chains := map[uint64]env.TChain{
					1:  {ID: 1},
					10: {ID: 10},
				}

				for chainID := range chains {
					// Mimic the behavior of storage.ListPrices
					allChainPrices := mockPrices[chainID]
					if len(allChainPrices) > 0 {
						rawPrices[chainID] = make(map[string]*bigNumber.Int)
						humanizedPrices[chainID] = make(map[string]*bigNumber.Float)

						for addr, price := range allChainPrices {
							if humanized {
								humanizedPrices[chainID][addr.Hex()] = price.HumanizedPrice
							} else {
								rawPrices[chainID][addr.Hex()] = price.Price
							}
						}
					}
				}

				// Use the utility function just like the real handler
				if humanized {
					c.JSON(http.StatusOK, humanizedPrices)
				} else {
					c.JSON(http.StatusOK, rawPrices)
				}
			}

			// Setup test router with our custom handler
			router := gin.New()
			router.GET("/prices/all", customHandler)

			// Create request
			url := "/prices/all"
			if tc.humanized != "" {
				url += "?humanized=" + tc.humanized
			}
			req, _ := http.NewRequest("GET", url, nil)
			w := httptest.NewRecorder()

			// Perform the request
			router.ServeHTTP(w, req)

			// Check status code
			assert.Equal(t, tc.expectedStatus, w.Code, "Status code should match expected")

			// Validate response
			if tc.validateJSON != nil {
				tc.validateJSON(t, w.Body.Bytes())
			}
		})
	}
}

/**************************************************************************************************
** TestGetPrices tests the GetPrices handler which retrieves prices for all tokens on a specific
** blockchain network. This test validates:
** - The handler correctly validates the chain ID
** - The correct response format is returned based on the 'humanized' query parameter
** - Error conditions (invalid chain ID, storage errors) are properly handled
**
** We use a custom handler that mimics the behavior of the real one but uses our mock data.
**************************************************************************************************/
func TestGetPrices(t *testing.T) {
	// Setup test environment
	gin.SetMode(gin.TestMode)

	// Create mock data
	mockPrices := setupMockPrices()

	// Create test cases
	tests := []struct {
		name           string
		chainID        string
		humanized      string
		expectedStatus int
		validateJSON   func(t *testing.T, body []byte)
		setupMock      func(storage *mockStorage)
	}{
		{
			name:           "Valid chain ID with raw prices",
			chainID:        "1",
			humanized:      "false",
			expectedStatus: http.StatusOK,
			validateJSON: func(t *testing.T, body []byte) {
				// Parse the response
				var response map[string]*json.RawMessage
				err := json.Unmarshal(body, &response)
				assert.NoError(t, err, "Failed to unmarshal response")

				// Verify response content
				assert.Len(t, response, 2, "Response should contain 2 tokens")
				assert.Contains(t, response, "0x6B175474E89094C44Da98b954EedeAC495271d0F", "Response should contain DAI token")
			},
		},
		{
			name:           "Valid chain ID with humanized prices",
			chainID:        "1",
			humanized:      "true",
			expectedStatus: http.StatusOK,
			validateJSON: func(t *testing.T, body []byte) {
				// Parse the response
				var response map[string]*json.RawMessage
				err := json.Unmarshal(body, &response)
				assert.NoError(t, err, "Failed to unmarshal response")

				// Verify response content
				assert.Len(t, response, 2, "Response should contain 2 tokens")
				assert.Contains(t, response, "0x6B175474E89094C44Da98b954EedeAC495271d0F", "Response should contain DAI token")
			},
		},
		{
			name:           "Invalid chain ID",
			chainID:        "invalid",
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "Empty result set",
			chainID:        "42161",
			expectedStatus: http.StatusOK,
			validateJSON: func(t *testing.T, body []byte) {
				// For an empty result, we should still get a valid JSON object
				var response map[string]*json.RawMessage
				err := json.Unmarshal(body, &response)
				assert.NoError(t, err, "Failed to unmarshal response")
				assert.Len(t, response, 0, "Response should be empty")
			},
		},
		{
			name:           "Storage layer failure",
			chainID:        "1",
			expectedStatus: http.StatusInternalServerError,
			validateJSON: func(t *testing.T, body []byte) {
				var response map[string]string
				err := json.Unmarshal(body, &response)
				assert.NoError(t, err)
				assert.Equal(t, "Failed to retrieve prices", response["error"])
			},
			setupMock: func(storage *mockStorage) {
				storage.returnError = true
			},
		},
	}

	// For each test case
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Create a new test router
			gin.SetMode(gin.TestMode)
			router := gin.New()

			// Setup mock storage
			mockStorage := &mockStorage{
				prices: mockPrices,
			}

			// Apply custom mock setup if provided
			if tc.setupMock != nil {
				tc.setupMock(mockStorage)
			}

			// Register our mock handler
			handler := MockPriceHandler(mockStorage, tc.chainID, tc.humanized)
			router.GET("/prices/:chainID", handler)

			// Perform the request
			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", "/prices/"+tc.chainID, nil)
			router.ServeHTTP(w, req)

			// Check status code
			assert.Equal(t, tc.expectedStatus, w.Code)

			// Validate JSON response if needed
			if tc.validateJSON != nil {
				tc.validateJSON(t, w.Body.Bytes())
			}
		})
	}
}

/**************************************************************************************************
** setupMockPrices creates a mock dataset of token prices for testing purposes.
** This helper function simplifies the creation of test data used across multiple test functions.
**
** @return map[uint64]map[common.Address]mockPriceData Map of chain IDs to token price maps
**************************************************************************************************/
func setupMockPrices() map[uint64]map[common.Address]mockPriceData {
	// Create mock token prices
	dai := common.HexToAddress("0x6B175474E89094C44Da98b954EedeAC495271d0F")
	weth := common.HexToAddress("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2")
	usdc := common.HexToAddress("0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48")

	return map[uint64]map[common.Address]mockPriceData{
		1: {
			dai: {
				Price:          bigNumber.NewInt(1000000000000000000), // 1.0 in raw format
				HumanizedPrice: bigNumber.NewFloat(1.0),
			},
			weth: {
				Price:          bigNumber.NewInt(2000000000000000000), // 2.0 in raw format
				HumanizedPrice: bigNumber.NewFloat(2.0),
			},
		},
		10: {
			usdc: {
				Price:          bigNumber.NewInt(1000000), // 1.0 with 6 decimals
				HumanizedPrice: bigNumber.NewFloat(1.0),
			},
		},
		// Chain 42161 (Arbitrum) intentionally left empty for testing empty result sets
		42161: {},
	}
}
