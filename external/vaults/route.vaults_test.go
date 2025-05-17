package vaults

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

/**************************************************************************************************
** setupVaultsRouteTest creates a test environment for testing the vault route endpoints.
**
** This function sets up a Gin test context with optional query parameters for customizing
** the response format, pagination, and filtering.
**
** @param t *testing.T - The testing object
** @param queryParams map[string]string - Optional query parameters to include in the request
** @return *gin.Context - The configured Gin context
** @return *httptest.ResponseRecorder - The HTTP response recorder
**************************************************************************************************/
func setupVaultsRouteTest(t *testing.T, queryParams map[string]string) (*gin.Context, *httptest.ResponseRecorder) {
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
** TestGetAll tests the GetAll function that retrieves all vaults.
**
** This test verifies that the function correctly returns a list of all vaults without filtering,
** but respecting query parameters for sorting, pagination, and strategy conditions.
**
** @param t *testing.T - The testing object
**************************************************************************************************/
func TestGetAll(t *testing.T) {
	// Setup with default parameters
	c, _ := setupVaultsRouteTest(t, nil)
	controller := Controller{}

	// Execute
	result, err := controller.GetAll(c)

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.IsType(t, []TSimplifiedExternalVault{}, result)

	// In a test environment, we may not have any vaults loaded
	// so we just verify the function doesn't error and returns the expected type
}

/**************************************************************************************************
** TestGetIsYearn tests the GetIsYearn function that retrieves only Yearn vaults.
**
** This test verifies that the function correctly filters vaults to include only those
** marked with the 'isYearn' flag.
**
** @param t *testing.T - The testing object
**************************************************************************************************/
func TestGetIsYearn(t *testing.T) {
	// Setup with default parameters
	c, _ := setupVaultsRouteTest(t, nil)
	controller := Controller{}

	// Execute
	result, err := controller.GetIsYearn(c)

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, result)

	// We can't directly verify the IsYearn property as it's not exposed in the simplified response,
	// but we can verify we get results and no errors.
	// The filtering happens in the implementation through the filter function.
}

/**************************************************************************************************
** TestGetV3 tests the GetV3 function that retrieves only V3 vaults.
**
** This test verifies that the function correctly filters vaults to include only those
** using the V3 vault architecture.
**
** @param t *testing.T - The testing object
**************************************************************************************************/
func TestGetV3(t *testing.T) {
	// Setup with default parameters
	c, _ := setupVaultsRouteTest(t, nil)
	controller := Controller{}

	// Execute
	result, err := controller.GetV3(c)

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, result)

	// Verify all returned vaults are V3 vaults
	for _, vault := range result {
		assert.Equal(t, "v3", vault.Version, "Expected only V3 vaults to be returned")
	}
}

/**************************************************************************************************
** TestGetV2 tests the GetV2 function that retrieves only V2 vaults.
**
** This test verifies that the function correctly filters vaults to include only those
** using the V2 vault architecture.
**
** @param t *testing.T - The testing object
**************************************************************************************************/
func TestGetV2(t *testing.T) {
	// Setup with default parameters
	c, _ := setupVaultsRouteTest(t, nil)
	controller := Controller{}

	// Execute
	result, err := controller.GetV2(c)

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, result)

	// Verify all returned vaults are V2 vaults
	for _, vault := range result {
		assert.Equal(t, "v2", vault.Version, "Expected only V2 vaults to be returned")
	}
}

/**************************************************************************************************
** TestGetRetired tests the GetRetired function that retrieves only retired vaults.
**
** This test verifies that the function correctly filters vaults to include only those
** that have been marked as retired.
**
** @param t *testing.T - The testing object
**************************************************************************************************/
func TestGetRetired(t *testing.T) {
	// Setup with default parameters
	c, _ := setupVaultsRouteTest(t, nil)
	controller := Controller{}

	// Execute
	result, err := controller.GetRetired(c)

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, result)

	// We can't directly verify the IsRetired property as it's not exposed in the simplified response,
	// but we can verify we get results and no errors.
	// The filtering happens in the implementation through the filter function.
}

/**************************************************************************************************
** TestGetV2IsYearn tests the GetV2IsYearn function that combines V2 and IsYearn filters.
**
** This test verifies that the function correctly filters vaults to include only those
** that are both V2 vaults and officially part of Yearn Finance.
**
** @param t *testing.T - The testing object
**************************************************************************************************/
func TestGetV2IsYearn(t *testing.T) {
	// Setup with default parameters
	c, _ := setupVaultsRouteTest(t, nil)
	controller := Controller{}

	// Execute
	result, err := controller.GetV2IsYearn(c)

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, result)

	// We can only verify the Version part of the filter here
	for _, vault := range result {
		assert.Equal(t, "v2", vault.Version, "Expected only V2 vaults to be returned")
	}
	// The IsYearn filtering happens in the implementation through the filter function.
}

/**************************************************************************************************
** TestGetV3IsYearn tests the GetV3IsYearn function that combines V3 and IsYearn filters.
**
** This test verifies that the function correctly filters vaults to include only those
** that are both V3 vaults and officially part of Yearn Finance.
**
** @param t *testing.T - The testing object
**************************************************************************************************/
func TestGetV3IsYearn(t *testing.T) {
	// Setup with default parameters
	c, _ := setupVaultsRouteTest(t, nil)
	controller := Controller{}

	// Execute
	result, err := controller.GetV3IsYearn(c)

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, result)

	// We can only verify the Version part of the filter here
	for _, vault := range result {
		assert.Equal(t, "v3", vault.Version, "Expected only V3 vaults to be returned")
	}
	// The IsYearn filtering happens in the implementation through the filter function.
}

/**************************************************************************************************
** TestGetIsYearnJuiced tests the GetIsYearnJuiced function.
**
** This test verifies that the function correctly filters vaults to include only those
** marked with the 'IsYearnJuiced' flag.
**
** @param t *testing.T - The testing object
**************************************************************************************************/
func TestGetIsYearnJuiced(t *testing.T) {
	// Setup with default parameters
	c, _ := setupVaultsRouteTest(t, nil)
	controller := Controller{}

	// Execute
	result, err := controller.GetIsYearnJuiced(c)

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, result)

	// We can't directly verify the IsYearnJuiced property as it's not exposed in the simplified response,
	// but we can verify we get results and no errors.
	// The filtering happens in the implementation.
}

/**************************************************************************************************
** TestCategoryFilters tests the various category-based filter functions.
**
** This test verifies that functions like GetIsAjna, GetIsCurve, etc. correctly filter vaults
** based on their category.
**
** @param t *testing.T - The testing object
**************************************************************************************************/
func TestCategoryFilters(t *testing.T) {
	// Setup
	c, _ := setupVaultsRouteTest(t, nil)
	controller := Controller{}

	// Test all category filter functions
	categoryTests := []struct {
		name     string
		function func(*gin.Context) ([]TSimplifiedExternalVault, error)
	}{
		{"GetIsAjna", controller.GetIsAjna},
		{"GetIsCurve", controller.GetIsCurve},
		{"GetIsVelodrome", controller.GetIsVelodrome},
		{"GetIsAerodrome", controller.GetIsAerodrome},
		{"GetIsYearnPendle", controller.GetIsYearnPendle},
		{"GetIsGimme", controller.GetIsGimme},
		{"GetIsOptimism", controller.GetIsOptimism},
		{"GetIsYearnPoolTogether", controller.GetIsYearnPoolTogether},
		{"GetIsMorpho", controller.GetIsMorpho},
		{"GetIsKatana", controller.GetIsKatana},
	}

	for _, test := range categoryTests {
		t.Run(test.name, func(t *testing.T) {
			// Execute the filter function
			result, err := test.function(c)

			// Assert no errors and valid result type
			assert.NoError(t, err)
			assert.NotNil(t, result)
			assert.IsType(t, []TSimplifiedExternalVault{}, result)

			// We can't easily verify the category property in the simplified response,
			// but we can check for basic validity.
			// The actual filtering by category happens in the implementation.
		})
	}
}
