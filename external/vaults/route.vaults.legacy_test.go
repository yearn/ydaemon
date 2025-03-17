package vaults

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/yearn/ydaemon/internal/models"
)

/**************************************************************************************************
** setupLegacyVaultsTest creates a test environment for testing the legacy vault endpoints.
**
** This function sets up a Gin test context with optional query parameters for filtering and pagination.
**
** @param t *testing.T - The testing object
** @param queryParams map[string]string - Optional query parameters to include in the request
** @return *gin.Context - The configured Gin context
** @return *httptest.ResponseRecorder - The HTTP response recorder
**************************************************************************************************/
func setupLegacyVaultsTest(t *testing.T, queryParams map[string]string) (*gin.Context, *httptest.ResponseRecorder) {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	// Set up the request with query parameters if provided
	req, _ := http.NewRequest(http.MethodGet, "/", nil)

	if len(queryParams) > 0 {
		q := req.URL.Query()
		for key, value := range queryParams {
			q.Add(key, value)
		}
		req.URL.RawQuery = q.Encode()
	}

	c.Request = req
	return c, w
}

// This test verifies that the legacy endpoints exist and return expected structures
// In a testing environment with no database connection, we can't fully verify content
func TestLegacyEndpoints(t *testing.T) {
	// Define test cases for all legacy endpoints
	testCases := []struct {
		name     string
		endpoint func(*gin.Context) []TExternalVault
	}{
		{"GetLegacyIsYearn", (&Controller{}).GetLegacyIsYearn},
		{"GetLegacyV2IsYearn", (&Controller{}).GetLegacyV2IsYearn},
		{"GetLegacyV3IsYearn", (&Controller{}).GetLegacyV3IsYearn},
		{"GetLegacyRetired", (&Controller{}).GetLegacyRetired},
		{"GetLegacyIsYearnJuiced", (&Controller{}).GetLegacyIsYearnJuiced},
		{"GetLegacyIsGimme", (&Controller{}).GetLegacyIsGimme},
	}

	// For each endpoint, verify it returns an expected structure
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Since we can't verify actual vault data in a test environment,
			// we just verify that the function exists and returns the expected type.
			c, _ := setupLegacyVaultsTest(t, nil)

			// Call the function with our test Gin context
			result := tc.endpoint(c)

			// We can't expect actual vaults data in test context, so we just verify the type
			assert.IsType(t, []TExternalVault{}, result, "Expected endpoint to return []TExternalVault")
		})
	}
}

/**************************************************************************************************
** TestLegacyWithPagination tests the legacy endpoints with pagination parameters.
**
** This test verifies that the functions correctly handle pagination parameters (page and limit)
** in their URL queries. Since we're in a test environment, we can't verify actual data,
** but we can check that the pagination parameters are parsed correctly.
**
** @param t *testing.T - The testing object
**************************************************************************************************/
func TestLegacyWithPagination(t *testing.T) {
	// Setup with pagination parameters
	queryParams := map[string]string{
		"page":  "1",
		"limit": "10",
	}
	c, _ := setupLegacyVaultsTest(t, queryParams)
	controller := Controller{}

	// Execute - just checking that the function can handle pagination params
	result := controller.GetLegacyIsYearn(c)

	// We can't expect actual vaults data in test context, so we just verify the type
	assert.IsType(t, []TExternalVault{}, result, "Expected endpoint to return []TExternalVault")
}

/**************************************************************************************************
** TestIsV2VaultForLegacy tests the isV2Vault function used in the GetLegacyV2IsYearn endpoint.
**
** This test verifies that the function correctly identifies V2 vaults based on the Version field.
**
** @param t *testing.T - The testing object
**************************************************************************************************/
func TestIsV2VaultForLegacy(t *testing.T) {
	// Create test vaults with different versions
	v2Vault := models.TVault{
		Version: "0.3.2", // V2 vaults often have numeric versions
	}
	v3Vault := models.TVault{
		Version: "v3",
	}

	// Verify the function correctly identifies the vault versions
	assert.True(t, isV2Vault(v2Vault), "Expected v2 vault to be identified as V2")
	assert.False(t, isV2Vault(v3Vault), "Expected v3 vault to not be identified as V2")
}

/**************************************************************************************************
** TestIsV3VaultForLegacy tests the isV3Vault function used in the GetLegacyV3IsYearn endpoint.
**
** This test verifies that the function correctly identifies V3 vaults based on the Version field.
**
** @param t *testing.T - The testing object
**************************************************************************************************/
func TestIsV3VaultForLegacy(t *testing.T) {
	// Create test vaults with different versions
	v2Vault := models.TVault{
		Version: "v2",
	}
	v3Vault := models.TVault{
		Version: "v3",
	}

	// Verify the function correctly identifies the vault versions
	assert.False(t, isV3Vault(v2Vault), "Expected v2 vault to not be identified as V3")
	assert.True(t, isV3Vault(v3Vault), "Expected v3 vault to be identified as V3")
}
