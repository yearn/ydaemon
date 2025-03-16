package vaults

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

/**************************************************************************************************
** setupEarnedVaultsTest creates a test environment for testing the earned vaults endpoints.
**
** This function sets up a Gin test context with path parameters for chain ID and user address,
** which are required for the earned endpoints. It also configures a mock response writer to
** capture the endpoint's output.
**
** @param t *testing.T - The testing object
** @param chainID string - The chain ID to use in the test
** @param userAddress string - The user address to use in the test
** @param queryParams map[string]string - Optional query parameters to include
** @return *gin.Context - The configured Gin context
** @return *httptest.ResponseRecorder - The HTTP response recorder
**************************************************************************************************/
func setupEarnedVaultsTest(t *testing.T, chainID, userAddress string, queryParams map[string]string) (*gin.Context, *httptest.ResponseRecorder) {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	// Set up the request with path parameters
	req, _ := http.NewRequest(http.MethodGet, "/", nil)

	// Add query parameters if provided
	if len(queryParams) > 0 {
		q := req.URL.Query()
		for key, value := range queryParams {
			q.Add(key, value)
		}
		req.URL.RawQuery = q.Encode()
	}

	c.Request = req

	// Set path parameters
	c.Params = gin.Params{
		{Key: "chainID", Value: chainID},
		{Key: "address", Value: userAddress},
	}

	return c, w
}

/**************************************************************************************************
** TestGetEarnedPerVaultPerUser tests the error handling in the GetEarnedPerVaultPerUser function.
**
** This test verifies that the function correctly handles invalid input parameters.
** Note: This test only validates error handling, not the actual earned calculation which
** requires access to GraphQL endpoints.
**
** @param t *testing.T - The testing object
**************************************************************************************************/
func TestGetEarnedPerVaultPerUser(t *testing.T) {
	// Skip test for actual functionality that requires GraphQL endpoints
	t.Log("Note: Tests for actual earned calculation functionality are skipped as they require GraphQL endpoints")

	t.Run("Invalid_chain_ID", func(t *testing.T) {
		// Test with invalid chain ID
		c, w := setupEarnedVaultsTest(t, "invalid", "0x1234", nil)
		controller := Controller{}

		// Execute
		controller.GetEarnedPerVaultPerUser(c)

		// Assert we get the expected 400 Bad Request response
		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Equal(t, "invalid chainID: invalid", w.Body.String())
	})

	t.Run("Invalid_address", func(t *testing.T) {
		// Test with valid chain ID but invalid address
		c, w := setupEarnedVaultsTest(t, "1", "invalid", nil)
		controller := Controller{}
		controller.GetEarnedPerVaultPerUser(c)

		// Assert we get the expected 400 Bad Request response
		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Equal(t, "invalid address: invalid", w.Body.String())
	})
}

/**************************************************************************************************
** TestGetEarnedPerUser tests the error handling in the GetEarnedPerUser function.
**
** This test verifies that the function correctly handles invalid input parameters.
** Note: This test only validates error handling, not the actual earned calculation which
** requires access to GraphQL endpoints.
**
** @param t *testing.T - The testing object
**************************************************************************************************/
func TestGetEarnedPerUser(t *testing.T) {
	// Skip test for actual functionality that requires GraphQL endpoints
	t.Log("Note: Tests for actual earned calculation functionality are skipped as they require GraphQL endpoints")

	t.Run("Invalid_chain_ID", func(t *testing.T) {
		// Test with invalid chain ID
		c, w := setupEarnedVaultsTest(t, "invalid", "0x1234", nil)
		controller := Controller{}

		// Execute
		controller.GetEarnedPerUser(c)

		// Assert we get the expected 400 Bad Request response
		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Equal(t, "invalid chainID: invalid", w.Body.String())
	})

	t.Run("Invalid_address", func(t *testing.T) {
		// Test with valid chain ID but invalid address
		c, w := setupEarnedVaultsTest(t, "1", "invalid", nil)
		controller := Controller{}
		controller.GetEarnedPerUser(c)

		// Assert we get the expected 400 Bad Request response
		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Equal(t, "invalid address: invalid", w.Body.String())
	})
}

/**************************************************************************************************
** TestGetEarnedPerUserForAllChains tests the error handling in the GetEarnedPerUserForAllChains function.
**
** This test verifies that the function correctly handles invalid input parameters.
** Note: This test only validates error handling, not the actual earned calculation which
** requires access to GraphQL endpoints.
**
** @param t *testing.T - The testing object
**************************************************************************************************/
func TestGetEarnedPerUserForAllChains(t *testing.T) {
	// Skip test for actual functionality that requires GraphQL endpoints
	t.Log("Note: Tests for actual earned calculation functionality are skipped as they require GraphQL endpoints")

	t.Run("Invalid_address", func(t *testing.T) {
		// Test with invalid address
		c, w := setupEarnedVaultsTest(t, "", "invalid", nil)
		controller := Controller{}

		// Execute
		controller.GetEarnedPerUserForAllChains(c)

		// Assert we get the expected 400 Bad Request response
		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Equal(t, "invalid address: invalid", w.Body.String())
	})
}
