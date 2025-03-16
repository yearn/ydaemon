package utils

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/yearn/ydaemon/common/env"
)

/**************************************************************************************************
** setupTestRouter creates a Gin router with the utility controller routes for testing.
**
** @return A configured Gin router for testing
**************************************************************************************************/
func setupTestRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.New()

	controller := Controller{}
	router.GET("/utils/supported-chains", controller.GetSupportedChains)
	router.GET("/utils/version", controller.GetAPIVersion)
	router.GET("/utils/health", controller.HealthCheck)

	return router
}

/**************************************************************************************************
** TestGetSupportedChains tests the GetSupportedChains endpoint to ensure it returns
** the correct list of supported chains from the environment.
**
** This test verifies:
** - The HTTP status code is 200 OK
** - The response is a valid JSON array
** - The array contains the expected chain IDs from the environment
**************************************************************************************************/
func TestGetSupportedChains(t *testing.T) {
	router := setupTestRouter()

	// Create a test request
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/utils/supported-chains", nil)
	router.ServeHTTP(w, req)

	// Verify HTTP status code
	assert.Equal(t, http.StatusOK, w.Code)

	// Parse response
	var response []uint64
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)

	// Compare with expected chains
	// Note: This test assumes env.SUPPORTED_CHAIN_IDS is correctly initialized
	assert.Equal(t, env.SUPPORTED_CHAIN_IDS, response)

	// Verify specific expected chains are included
	expectedChains := []uint64{1} // At minimum, Ethereum Mainnet should be supported
	for _, chainID := range expectedChains {
		assert.Contains(t, response, chainID)
	}
}

/**************************************************************************************************
** TestGetAPIVersion tests the GetAPIVersion endpoint to ensure it returns a valid
** version string in the expected format.
**
** This test verifies:
** - The HTTP status code is 200 OK
** - The response is a valid JSON object
** - The object contains the expected keys: "version" and "name"
** - The values are non-empty strings
**************************************************************************************************/
func TestGetAPIVersion(t *testing.T) {
	router := setupTestRouter()

	// Create a test request
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/utils/version", nil)
	router.ServeHTTP(w, req)

	// Verify HTTP status code
	assert.Equal(t, http.StatusOK, w.Code)

	// Parse response
	var response map[string]string
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)

	// Verify expected fields
	assert.Contains(t, response, "version")
	assert.Contains(t, response, "name")

	// Verify values are non-empty
	assert.NotEmpty(t, response["version"])
	assert.NotEmpty(t, response["name"])

	// Verify name is as expected
	assert.Equal(t, "yDaemon", response["name"])
}

/**************************************************************************************************
** TestHealthCheck tests the HealthCheck endpoint to ensure it returns the expected
** healthy status response.
**
** This test verifies:
** - The HTTP status code is 200 OK
** - The response is a valid JSON object
** - The object contains the expected "status" key with "healthy" value
**************************************************************************************************/
func TestHealthCheck(t *testing.T) {
	router := setupTestRouter()

	// Create a test request
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/utils/health", nil)
	router.ServeHTTP(w, req)

	// Verify HTTP status code
	assert.Equal(t, http.StatusOK, w.Code)

	// Parse response
	var response map[string]string
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)

	// Verify expected fields
	assert.Contains(t, response, "status")

	// Verify status is healthy
	assert.Equal(t, "healthy", response["status"])
}

/**************************************************************************************************
** TestControllerStructure tests that the Controller struct can be initialized successfully.
**
** While simple, this test ensures the base structure required for the package is valid.
**************************************************************************************************/
func TestControllerStructure(t *testing.T) {
	controller := Controller{}
	assert.NotNil(t, controller, "Controller should be initializable")
}
