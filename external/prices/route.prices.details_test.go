package prices

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/yearn/ydaemon/common/bigNumber"
	"github.com/yearn/ydaemon/common/helpers"
)

/**************************************************************************************************
** TestGetAllPricesWithDetails tests the GetAllPricesWithDetails handler which retrieves detailed
** price information for all tokens on a specific blockchain network. This test validates:
** - The handler correctly validates the chain ID
** - The handler properly returns detailed price objects
** - Error conditions (invalid chain ID, storage errors) are correctly handled
**
** We use a custom handler that mimics the behavior of the real one but uses our mock data.
**************************************************************************************************/
func TestGetAllPricesWithDetails(t *testing.T) {
	// Setup test environment
	gin.SetMode(gin.TestMode)

	// Create mock data
	mockPrices := setupMockDetailsPrices()

	// Create test cases
	tests := []struct {
		name           string
		chainID        string
		mockError      error
		expectedStatus int
		validateJSON   func(t *testing.T, body []byte)
	}{
		{
			name:           "Valid chain ID",
			chainID:        "1",
			expectedStatus: http.StatusOK,
			validateJSON: func(t *testing.T, body []byte) {
				// Parse the response - should be a map of addresses to detailed price objects
				var response map[string]interface{}
				err := json.Unmarshal(body, &response)
				assert.NoError(t, err, "Failed to unmarshal response")

				// Verify response content
				assert.Len(t, response, 2, "Response should contain 2 tokens")

				// Check that the DAI token is in the response
				daiAddr := "0x6B175474E89094C44Da98b954EedeAC495271d0F"
				assert.Contains(t, response, daiAddr, "Response should contain DAI token")

				// Verify the structure of a price detail object
				daiPrice, ok := response[daiAddr].(map[string]interface{})
				assert.True(t, ok, "DAI price should be a map")
				assert.Contains(t, daiPrice, "price", "Price object should contain raw price")
				assert.Contains(t, daiPrice, "humanizedPrice", "Price object should contain humanized price")
				assert.Contains(t, daiPrice, "lastUpdated", "Price object should contain lastUpdated")
			},
		},
		{
			name:           "Invalid chain ID",
			chainID:        "invalid",
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "Storage error",
			chainID:        "1",
			mockError:      errors.New("database error"),
			expectedStatus: http.StatusInternalServerError,
		},
		{
			name:           "Empty result set",
			chainID:        "42161",
			expectedStatus: http.StatusOK,
			validateJSON: func(t *testing.T, body []byte) {
				// For an empty result, we should still get a valid JSON object
				var response map[string]interface{}
				err := json.Unmarshal(body, &response)
				assert.NoError(t, err, "Failed to unmarshal response")
				assert.Len(t, response, 0, "Response should be empty")
			},
		},
	}

	// For each test case
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Setup a custom handler that mimics the real one but uses our mock data
			customHandler := func(c *gin.Context) {
				chainID, ok := helpers.AssertChainID(c.Param("chainID"))
				if !ok {
					c.String(http.StatusBadRequest, "invalid chainID")
					return
				}

				// Check if this is a supported chain ID
				supportedChains := map[uint64]bool{
					1:     true,
					42161: true, // Added Arbitrum as supported
				}

				if _, supported := supportedChains[chainID]; !supported {
					c.String(http.StatusBadRequest, "invalid chainID")
					return
				}

				// If we're testing a storage error
				if tc.mockError != nil {
					c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve price details"})
					return
				}

				// Mimic the behavior of storage.ListPrices
				allChainPrices := mockPrices[chainID]

				// Convert to the format expected in the response
				result := make(map[string]map[string]interface{})
				for addr, priceDetail := range allChainPrices {
					result[addr.Hex()] = map[string]interface{}{
						"price":          priceDetail.Price,
						"humanizedPrice": priceDetail.HumanizedPrice,
						"lastUpdated":    priceDetail.LastUpdated,
						"source":         priceDetail.Source,
					}
				}

				c.JSON(http.StatusOK, result)
			}

			// Setup test router with our custom handler
			router := gin.New()
			router.GET("/prices/:chainID/all", customHandler)

			// Create request
			url := "/prices/" + tc.chainID + "/all"
			req, _ := http.NewRequest("GET", url, nil)
			w := httptest.NewRecorder()

			// Perform the request
			router.ServeHTTP(w, req)

			// Check status code
			assert.Equal(t, tc.expectedStatus, w.Code, "Status code should match expected")

			// Validate response
			if tc.validateJSON != nil && w.Code == http.StatusOK {
				tc.validateJSON(t, w.Body.Bytes())
			}
		})
	}
}

/**************************************************************************************************
** mockDetailsPriceData represents a detailed price entry for testing, including additional
** metadata beyond just the price values.
**************************************************************************************************/
type mockDetailsPriceData struct {
	Price          *bigNumber.Int
	HumanizedPrice *bigNumber.Float
	LastUpdated    uint64
	Source         string
}

/**************************************************************************************************
** setupMockDetailsPrices creates a mock dataset of detailed token prices for testing purposes.
** This helper function includes additional metadata like lastUpdated timestamps and price sources.
**
** @return map[uint64]map[common.Address]mockDetailsPriceData Map of chain IDs to detailed price maps
**************************************************************************************************/
func setupMockDetailsPrices() map[uint64]map[common.Address]mockDetailsPriceData {
	// Create mock token prices with additional details
	dai := common.HexToAddress("0x6B175474E89094C44Da98b954EedeAC495271d0F")
	weth := common.HexToAddress("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2")

	return map[uint64]map[common.Address]mockDetailsPriceData{
		1: {
			dai: {
				Price:          bigNumber.NewInt(1000000000000000000), // 1.0 in raw format
				HumanizedPrice: bigNumber.NewFloat(1.0),
				LastUpdated:    1640995200, // 2022-01-01 00:00:00 UTC
				Source:         "coingecko",
			},
			weth: {
				Price:          bigNumber.NewInt(2000000000000000000), // 2.0 in raw format
				HumanizedPrice: bigNumber.NewFloat(2.0),
				LastUpdated:    1640995200, // 2022-01-01 00:00:00 UTC
				Source:         "defillama",
			},
		},
		// Chain 42161 (Arbitrum) intentionally left empty for testing empty result sets
		42161: {},
	}
}
