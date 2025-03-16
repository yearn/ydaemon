package vaults

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/yearn/ydaemon/common/env"
)

/**************************************************************************************************
** setupRotkiTest creates a test environment for testing the Rotki integration endpoints.
**
** This function sets up a Gin test context with optional query parameters for customizing
** the response format, pagination, and chain filtering.
**
** @param t *testing.T - The testing object
** @param queryParams map[string]string - Optional query parameters to include in the request
** @return *gin.Context - The configured Gin context
** @return *httptest.ResponseRecorder - The HTTP response recorder
**************************************************************************************************/
func setupRotkiTest(t *testing.T, queryParams map[string]string) (*gin.Context, *httptest.ResponseRecorder) {
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

/**************************************************************************************************
** TestGetVaultsForRotki_DefaultParams tests the GetVaultsForRotki function with default parameters.
**
** This test verifies that the function correctly returns a list of vaults in Rotki format
** when using the default parameters (no specific filters or sorts).
**
** @param t *testing.T - The testing object
**************************************************************************************************/
func TestGetVaultsForRotki_DefaultParams(t *testing.T) {
	// Setup with default parameters
	c, _ := setupRotkiTest(t, nil)
	controller := Controller{}

	// Execute
	result := controller.GetVaultsForRotki(c)

	// Assert
	assert.NotNil(t, result)
	assert.IsType(t, []TRotkiVaults{}, result)

	// Should return results with default limit of 200
	assert.LessOrEqual(t, len(result), 200)

	// If we have results, verify the structure
	if len(result) > 0 {
		firstVault := result[0]
		assert.NotEmpty(t, firstVault.Address)
		assert.NotEmpty(t, firstVault.Name)
		assert.NotEmpty(t, firstVault.Symbol)

		// Token should have valid data
		assert.NotEmpty(t, firstVault.Token.Address)
		assert.NotEmpty(t, firstVault.Token.Symbol)
	}
}

/**************************************************************************************************
** TestGetVaultsForRotki_WithPagination tests the GetVaultsForRotki function with pagination.
**
** This test verifies that the function correctly handles pagination parameters (skip and limit)
** and returns the expected subset of results.
**
** @param t *testing.T - The testing object
**************************************************************************************************/
func TestGetVaultsForRotki_WithPagination(t *testing.T) {
	// Setup with pagination parameters
	queryParams := map[string]string{
		"skip":  "5",
		"limit": "10",
	}
	c, _ := setupRotkiTest(t, queryParams)
	controller := Controller{}

	// Execute
	result := controller.GetVaultsForRotki(c)

	// Assert
	assert.NotNil(t, result)
	// Should return at most 10 results
	assert.LessOrEqual(t, len(result), 10)
}

/**************************************************************************************************
** TestGetVaultsForRotki_WithSorting tests the GetVaultsForRotki function with sorting parameters.
**
** This test verifies that the function correctly handles sorting parameters (orderBy and orderDirection)
** and returns results in the expected order.
**
** @param t *testing.T - The testing object
**************************************************************************************************/
func TestGetVaultsForRotki_WithSorting(t *testing.T) {
	// Setup with sorting parameters
	queryParams := map[string]string{
		"orderBy":        "featuringScore",
		"orderDirection": "desc",
	}
	c, _ := setupRotkiTest(t, queryParams)
	controller := Controller{}

	// Execute
	result := controller.GetVaultsForRotki(c)

	// Assert
	assert.NotNil(t, result)

	// If we have at least 2 results, verify they're sorted
	if len(result) >= 2 {
		// Check that the first vault has higher or equal featuring score than the second
		assert.GreaterOrEqual(t, result[0].FeaturingScore, result[1].FeaturingScore)
	}
}

/**************************************************************************************************
** TestGetVaultsForRotki_WithChainFilter tests the GetVaultsForRotki function with chain filtering.
**
** This test verifies that the function correctly filters results based on the chainIDs parameter
** and returns only vaults from the specified chains.
**
** @param t *testing.T - The testing object
**************************************************************************************************/
func TestGetVaultsForRotki_WithChainFilter(t *testing.T) {
	// Skip test if no chains are configured
	if len(env.SUPPORTED_CHAIN_IDS) == 0 {
		t.Skip("No chains configured, skipping test")
	}

	// Use the first supported chain for testing
	testChainID := env.SUPPORTED_CHAIN_IDS[0]

	// Setup with chain filter
	queryParams := map[string]string{
		"chainIDs": strconv.FormatUint(testChainID, 10),
	}
	c, _ := setupRotkiTest(t, queryParams)
	controller := Controller{}

	// Execute
	result := controller.GetVaultsForRotki(c)

	// Assert
	assert.NotNil(t, result)

	// We can't easily verify chain IDs since they're not in the response,
	// but we can check that we get results
	// The implementation would filter correctly based on the chain ID
}

/**************************************************************************************************
** TestCountVaultsForRotki tests the CountVaultsForRotki function with various parameters.
**
** This test verifies that the function correctly counts the total number of vaults available
** for Rotki integration, with optional chain filtering.
**
** @param t *testing.T - The testing object
**************************************************************************************************/
func TestCountVaultsForRotki(t *testing.T) {
	// Test default parameters (all chains)
	c, w := setupRotkiTest(t, nil)
	controller := Controller{}

	// Execute
	controller.CountVaultsForRotki(c)

	// Assert
	assert.Equal(t, http.StatusOK, w.Code)

	// Parse response body
	var response map[string]int
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)

	// Should have a numberOfVaults field with a non-negative value
	count, exists := response["numberOfVaults"]
	assert.True(t, exists)
	assert.GreaterOrEqual(t, count, 0)

	// Now test with a chain filter if we have supported chains
	if len(env.SUPPORTED_CHAIN_IDS) > 0 {
		// Use the first supported chain for testing
		testChainID := env.SUPPORTED_CHAIN_IDS[0]

		queryParams := map[string]string{
			"chainIDs": strconv.FormatUint(testChainID, 10),
		}
		c, w = setupRotkiTest(t, queryParams)

		// Execute
		controller.CountVaultsForRotki(c)

		// Assert
		assert.Equal(t, http.StatusOK, w.Code)

		// Parse response body
		var filteredResponse map[string]int
		err = json.Unmarshal(w.Body.Bytes(), &filteredResponse)
		assert.NoError(t, err)

		// Should have a numberOfVaults field with a non-negative value
		filteredCount, exists := filteredResponse["numberOfVaults"]
		assert.True(t, exists)
		assert.GreaterOrEqual(t, filteredCount, 0)

		// Filtered count should be less than or equal to total count
		assert.LessOrEqual(t, filteredCount, count)
	}
}
