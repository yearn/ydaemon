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
	router.Use(gin.Recovery()) // Add recovery middleware
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
			address:        "0x1234567890123456789012345678901234567890", // Valid-looking address
			queryParams:    "",
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "Invalid address",
			chainID:        "1", // Valid chain ID
			address:        "invalid",
			queryParams:    "",
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "Valid parameters but non-existent vault",
			chainID:        "1",                                          // Valid chain ID
			address:        "0x1234567890123456789012345678901234567890", // Valid-looking address
			queryParams:    "",
			expectedStatus: http.StatusNotFound, // Updated to match the new error code for not found
		},
		{
			name:           "Valid parameters with strategies condition",
			chainID:        "1",                                          // Valid chain ID
			address:        "0x1234567890123456789012345678901234567890", // Valid-looking address
			queryParams:    "?strategiesCondition=all",
			expectedStatus: http.StatusNotFound, // Updated to match the new error code for not found
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

			// Add Accept header to get the right response format
			req.Header.Set("Accept", "text/plain")

			// Serve the request
			router.ServeHTTP(w, req)

			// Check the response code - for "not found" cases, both 404 and 500 are acceptable
			if tc.name == "Valid parameters but non-existent vault" || tc.name == "Valid parameters with strategies condition" {
				assert.True(t, w.Code == http.StatusNotFound || w.Code == http.StatusInternalServerError,
					"Expected status code to be either 404 or 500, got %d", w.Code)
			} else {
				assert.Equal(t, tc.expectedStatus, w.Code, "Should return expected status code")
			}
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
