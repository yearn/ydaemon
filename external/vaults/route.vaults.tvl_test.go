package vaults

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

/**************************************************************************************************
** TestGetAllVaultsTVL verifies that the GetAllVaultsTVL handler correctly
** aggregates TVL across all supported chains and returns a properly formatted response.
**************************************************************************************************/
func TestGetAllVaultsTVL(t *testing.T) {
	// Setup
	gin.SetMode(gin.TestMode)
	router := gin.New()
	controller := Controller{}

	// Register the handler
	router.GET("/vaults/tvl", controller.GetAllVaultsTVL)

	// Create a test request
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/vaults/tvl", nil)

	// Serve the request
	router.ServeHTTP(w, req)

	// This test verifies the handler doesn't crash and returns JSON
	// A more complete test would mock external dependencies

	// Check the response code
	assert.Equal(t, http.StatusOK, w.Code, "Should return 200 OK")

	// Check the response is valid JSON with expected structure
	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.Nil(t, err, "Response should be valid JSON")
	assert.Contains(t, response, "total", "Response should contain total TVL")
	assert.Contains(t, response, "chains", "Response should contain chains breakdown")
}

/**************************************************************************************************
** TestGetVaultsTVL verifies that the GetVaultsTVL handler correctly validates
** chain ID input and returns appropriate responses.
**************************************************************************************************/
func TestGetVaultsTVL(t *testing.T) {
	// Setup
	gin.SetMode(gin.TestMode)
	router := gin.New()
	controller := Controller{}

	// Register the handler
	router.GET("/vaults/:chainID/tvl", controller.GetVaultsTVL)

	// Test cases
	testCases := []struct {
		name           string
		chainID        string
		expectedStatus int
	}{
		{
			name:           "Valid chain ID",
			chainID:        "1", // Ethereum
			expectedStatus: http.StatusOK,
		},
		{
			name:           "Invalid chain ID - non-numeric",
			chainID:        "invalid",
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "Invalid chain ID - negative",
			chainID:        "-1",
			expectedStatus: http.StatusBadRequest,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Create a test request
			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", "/vaults/"+tc.chainID+"/tvl", nil)

			// Serve the request
			router.ServeHTTP(w, req)

			// Check the response code
			assert.Equal(t, tc.expectedStatus, w.Code, "Should return expected status code")

			if tc.expectedStatus == http.StatusOK {
				// If successful, check that the response is a number (TVL)
				var tvl float64
				err := json.Unmarshal(w.Body.Bytes(), &tvl)
				assert.Nil(t, err, "Response should be a valid JSON number")
			}
		})
	}
}
