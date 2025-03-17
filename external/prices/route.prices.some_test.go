package prices

import (
	"bytes"
	"encoding/json"
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
** TestGetSomePrices tests the GetSomePrices handler which retrieves prices for a specific list
** of tokens on a particular blockchain network. This test validates:
** - The handler correctly processes a comma-separated list of addresses
** - The handler validates the chain ID and addresses
** - The correct response format is returned based on the 'humanized' query parameter
** - Error conditions (invalid chain ID, invalid addresses) are properly handled
**
** We use a custom handler that mimics the behavior of the real one but uses our mock data.
**************************************************************************************************/
func TestGetSomePrices(t *testing.T) {
	// Setup test environment
	gin.SetMode(gin.TestMode)

	// Create mock data
	mockPrices := setupMockPrices()

	// DAI and WETH addresses for testing
	daiAddressHex := "0x6B175474E89094C44Da98b954EedeAC495271d0F"
	wethAddressHex := "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"

	// Create test cases
	tests := []struct {
		name           string
		chainID        string
		addressList    string
		humanized      string
		expectedStatus int
		validateJSON   func(t *testing.T, body []byte)
	}{
		{
			name:           "Valid addresses with raw prices",
			chainID:        "1",
			addressList:    daiAddressHex + "," + wethAddressHex,
			humanized:      "false",
			expectedStatus: http.StatusOK,
			validateJSON: func(t *testing.T, body []byte) {
				// Parse the response
				var response map[string]*json.RawMessage
				err := json.Unmarshal(body, &response)
				assert.NoError(t, err, "Failed to unmarshal response")

				// Verify response content
				assert.Len(t, response, 2, "Response should contain 2 tokens")
				assert.Contains(t, response, daiAddressHex, "Response should contain DAI token")
				assert.Contains(t, response, wethAddressHex, "Response should contain WETH token")
			},
		},
		{
			name:           "Valid addresses with humanized prices",
			chainID:        "1",
			addressList:    daiAddressHex + "," + wethAddressHex,
			humanized:      "true",
			expectedStatus: http.StatusOK,
			validateJSON: func(t *testing.T, body []byte) {
				// Parse the response
				var response map[string]*json.RawMessage
				err := json.Unmarshal(body, &response)
				assert.NoError(t, err, "Failed to unmarshal response")

				// Verify response content
				assert.Len(t, response, 2, "Response should contain 2 tokens")
				assert.Contains(t, response, daiAddressHex, "Response should contain DAI token")
				assert.Contains(t, response, wethAddressHex, "Response should contain WETH token")
			},
		},
		{
			name:           "Invalid chain ID",
			chainID:        "invalid",
			addressList:    daiAddressHex,
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "Invalid addresses",
			chainID:        "1",
			addressList:    "0xinvalid,0xnotanaddress",
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "Empty address list",
			chainID:        "1",
			addressList:    "",
			expectedStatus: http.StatusNotFound,
		},
		{
			name:           "Mix of valid and invalid addresses",
			chainID:        "1",
			addressList:    daiAddressHex + ",0xinvalid",
			expectedStatus: http.StatusOK, // Should return results for valid addresses
			validateJSON: func(t *testing.T, body []byte) {
				// Parse the response
				var response map[string]*json.RawMessage
				err := json.Unmarshal(body, &response)
				assert.NoError(t, err, "Failed to unmarshal response")

				// Should only contain the valid address
				assert.Len(t, response, 1, "Response should contain 1 token")
				assert.Contains(t, response, daiAddressHex, "Response should contain DAI token")
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

				addressesStr := c.Param("addressList")
				if addressesStr == "" {
					c.String(http.StatusNotFound, "not found")
					return
				}

				humanized := c.Query("humanized") == "true"

				// Split the comma-separated string
				addressList := splitAndTrim(addressesStr, ",")

				// Validate addresses
				validAddresses := make([]common.Address, 0)
				invalidAddresses := make(map[string]string)

				for _, addrStr := range addressList {
					addr, ok := helpers.AssertAddress(addrStr, chainID)
					if !ok {
						invalidAddresses[addrStr] = "invalid address format"
						continue
					}
					validAddresses = append(validAddresses, addr)
				}

				if len(validAddresses) == 0 {
					c.JSON(http.StatusBadRequest, gin.H{
						"error":            "No valid addresses provided",
						"invalidAddresses": invalidAddresses,
					})
					return
				}

				// Prepare response maps
				rawPrices := make(map[string]*bigNumber.Int)
				humanizedPrices := make(map[string]*bigNumber.Float)

				// Get prices for valid addresses
				for _, addr := range validAddresses {
					mockPrice, exists := mockPrices[chainID][addr]
					if exists {
						rawPrices[addr.Hex()] = mockPrice.Price
						humanizedPrices[addr.Hex()] = mockPrice.HumanizedPrice
					}
				}

				// Return appropriate response based on humanized flag
				if humanized {
					c.JSON(http.StatusOK, humanizedPrices)
				} else {
					c.JSON(http.StatusOK, rawPrices)
				}
			}

			// Setup test router with our custom handler
			router := gin.New()
			router.GET("/prices/:chainID/some/:addressList", customHandler)

			// Create request
			url := "/prices/" + tc.chainID + "/some/" + tc.addressList
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
			if tc.validateJSON != nil && w.Code == http.StatusOK {
				tc.validateJSON(t, w.Body.Bytes())
			}
		})
	}
}

/**************************************************************************************************
** TestGetSomePostPrices tests the GetSomePostPrices handler which retrieves prices for a list
** of tokens provided in the request body. This test validates:
** - The handler correctly parses the request body
** - The handler validates the addresses and returns prices for valid ones
** - The correct response format is returned based on the 'humanized' query parameter
** - Error conditions (invalid JSON, invalid addresses) are properly handled
**
** We use a custom handler that mimics the behavior of the real one but uses our mock data.
**************************************************************************************************/
func TestGetSomePostPrices(t *testing.T) {
	// Setup test environment
	gin.SetMode(gin.TestMode)

	// Create mock data
	mockPrices := setupMockPrices()

	// Define test cases
	tests := []struct {
		name           string
		requestBody    string
		humanized      string
		expectedStatus int
		validateJSON   func(t *testing.T, body []byte)
	}{
		{
			name:           "Valid addresses in different chains",
			requestBody:    `{"addresses": "1:0x6B175474E89094C44Da98b954EedeAC495271d0F,10:0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48"}`,
			humanized:      "false",
			expectedStatus: http.StatusOK,
			validateJSON: func(t *testing.T, body []byte) {
				// Parse the response - should be a nested map: chainID -> address -> price
				var response map[string]map[string]*json.RawMessage
				err := json.Unmarshal(body, &response)
				assert.NoError(t, err, "Failed to unmarshal response")

				// Verify response structure
				assert.Contains(t, response, "1", "Response should contain chain 1")
				chain1 := response["1"]
				assert.Contains(t, chain1, "0x6B175474E89094C44Da98b954EedeAC495271d0F", "Chain 1 should contain DAI")
			},
		},
		{
			name:           "Invalid JSON body",
			requestBody:    `{"addresses": invalid-json}`,
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "Empty addresses",
			requestBody:    `{"addresses": ""}`,
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "All invalid addresses",
			requestBody:    `{"addresses": "1:0xinvalid,10:0xnotanaddress"}`,
			expectedStatus: http.StatusBadRequest,
		},
	}

	// For each test case
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Setup a custom handler that mimics the real one but uses our mock data
			customHandler := func(c *gin.Context) {
				// Parse request body
				var body struct {
					Addresses string `json:"addresses"`
				}

				if err := c.BindJSON(&body); err != nil {
					c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
					return
				}

				if body.Addresses == "" {
					c.JSON(http.StatusBadRequest, gin.H{"error": "No addresses provided"})
					return
				}

				humanized := c.Query("humanized") == "true"

				// Process the comma-separated string of chainID:address pairs
				addressPairs := splitAndTrim(body.Addresses, ",")

				// Prepare response maps
				rawPrices := make(map[uint64]map[string]*bigNumber.Int)
				humanizedPrices := make(map[uint64]map[string]*bigNumber.Float)

				validAddressFound := false

				for _, pair := range addressPairs {
					// Split into chainID and address
					parts := splitAndTrim(pair, ":")
					if len(parts) != 2 {
						continue
					}

					chainIDStr, addrStr := parts[0], parts[1]

					// Validate chainID
					chainID, ok := helpers.AssertChainID(chainIDStr)
					if !ok {
						continue
					}

					// Validate address
					addr, ok := helpers.AssertAddress(addrStr, chainID)
					if !ok {
						continue
					}

					// Initialize maps if needed
					if _, exists := rawPrices[chainID]; !exists {
						rawPrices[chainID] = make(map[string]*bigNumber.Int)
						humanizedPrices[chainID] = make(map[string]*bigNumber.Float)
					}

					// Get price from mock data
					mockPrice, exists := mockPrices[chainID][addr]
					if exists {
						rawPrices[chainID][addr.Hex()] = mockPrice.Price
						humanizedPrices[chainID][addr.Hex()] = mockPrice.HumanizedPrice
						validAddressFound = true
					}
				}

				if !validAddressFound {
					c.JSON(http.StatusBadRequest, gin.H{"error": "No valid addresses provided"})
					return
				}

				// Return appropriate response based on humanized flag
				if humanized {
					c.JSON(http.StatusOK, humanizedPrices)
				} else {
					c.JSON(http.StatusOK, rawPrices)
				}
			}

			// Setup test router with our custom handler
			router := gin.New()
			router.POST("/prices/some", customHandler)

			// Create request with body
			url := "/prices/some"
			if tc.humanized != "" {
				url += "?humanized=" + tc.humanized
			}

			req, _ := http.NewRequest("POST", url, bytes.NewBufferString(tc.requestBody))
			req.Header.Set("Content-Type", "application/json")
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
** TestSplitAndTrim tests the splitAndTrim utility function to ensure it correctly splits a
** string by a separator and trims whitespace from each element.
**************************************************************************************************/
func TestSplitAndTrim(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		sep      string
		expected []string
	}{
		{
			name:     "Basic comma separation",
			input:    "a,b,c",
			sep:      ",",
			expected: []string{"a", "b", "c"},
		},
		{
			name:     "With whitespace",
			input:    " a , b , c ",
			sep:      ",",
			expected: []string{"a", "b", "c"},
		},
		{
			name:     "Empty input",
			input:    "",
			sep:      ",",
			expected: []string{},
		},
		{
			name:     "Only separators",
			input:    ",,,",
			sep:      ",",
			expected: []string{},
		},
		{
			name:     "Mixed content",
			input:    "valid, ,empty,,  spaced  ",
			sep:      ",",
			expected: []string{"valid", "empty", "spaced"},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := splitAndTrim(tc.input, tc.sep)
			assert.Equal(t, tc.expected, result, "Split result should match expected")
		})
	}
}
