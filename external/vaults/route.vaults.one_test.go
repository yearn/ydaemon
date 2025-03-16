package vaults

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

/**************************************************************************************************
** TestGetVault verifies that the GetVault handler correctly validates input parameters
** and returns appropriate responses.
**************************************************************************************************/
func TestGetVault(t *testing.T) {
	// Setup
	gin.SetMode(gin.TestMode)
	router := gin.New()
	controller := Controller{}

	// Register the handler
	router.GET("/vaults/:chainID/:address", controller.GetVault)

	// Test cases
	testCases := []struct {
		name           string
		chainID        string
		address        string
		queryParams    string
		expectedStatus int
	}{
		{
			name:           "Invalid chain ID",
			chainID:        "invalid",
			address:        "0x0bc529c00C6401aEF6D220BE8C6Ea1667F6Ad93e", // YFI token address as example
			queryParams:    "",
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "Invalid address",
			chainID:        "1", // Ethereum
			address:        "invalid-address",
			queryParams:    "",
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "Valid parameters but non-existent vault",
			chainID:        "1",                                          // Ethereum
			address:        "0x0000000000000000000000000000000000000001", // Non-existent vault address
			queryParams:    "",
			expectedStatus: http.StatusBadRequest, // The endpoint returns 400 for non-existent vaults
		},
		{
			name:           "Valid parameters with strategies condition",
			chainID:        "1",                                          // Ethereum
			address:        "0x0bc529c00C6401aEF6D220BE8C6Ea1667F6Ad93e", // Example address
			queryParams:    "?strategiesCondition=all",
			expectedStatus: http.StatusBadRequest, // We expect this to fail since we're not mocking storage
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Create a test request
			url := "/vaults/" + tc.chainID + "/" + tc.address
			if tc.queryParams != "" {
				url += tc.queryParams
			}

			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", url, nil)

			// Serve the request
			router.ServeHTTP(w, req)

			// Check the response code
			assert.Equal(t, tc.expectedStatus, w.Code, "Should return expected status code")
		})
	}
}

/**************************************************************************************************
** TestSelectStrategiesCondition verifies the selectStrategiesCondition helper function
** correctly selects strategy condition based on input.
**************************************************************************************************/
func TestSelectStrategiesConditionFromOne(t *testing.T) {
	testCases := []struct {
		name           string
		input          string
		expectedOutput string
	}{
		{
			name:           "Empty input",
			input:          "",
			expectedOutput: "debtRatio", // default value
		},
		{
			name:           "Valid input - all",
			input:          "all",
			expectedOutput: "all",
		},
		{
			name:           "Valid input - inQueue",
			input:          "inQueue",
			expectedOutput: "inQueue",
		},
		{
			name:           "Valid input - debtRatio",
			input:          "debtRatio",
			expectedOutput: "debtRatio",
		},
		{
			name:           "Valid input - absolute",
			input:          "absolute",
			expectedOutput: "absolute",
		},
		{
			name:           "Invalid input",
			input:          "invalid",
			expectedOutput: "debtRatio", // should default to debtRatio
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := selectStrategiesCondition(tc.input)
			assert.Equal(t, tc.expectedOutput, result, "Should select the correct strategy condition")
		})
	}
}
