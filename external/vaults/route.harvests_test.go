package vaults

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/machinebox/graphql"
	"github.com/stretchr/testify/assert"
)

/**************************************************************************************************
** setupHarvestsTest creates a test environment for testing the harvests endpoints.
**
** This function sets up a Gin test context with predefined path parameters for chain ID and
** vault addresses, which are required for the harvests endpoint. It also configures a mock
** response writer to capture the endpoint's output.
**
** @param t *testing.T - The testing object
** @param chainID string - The chain ID to use in the test
** @param addresses string - Comma-separated vault addresses
** @return *gin.Context - The configured Gin context
** @return *httptest.ResponseRecorder - The HTTP response recorder
**************************************************************************************************/
func setupHarvestsTest(t *testing.T, chainID, addresses string) (*gin.Context, *httptest.ResponseRecorder) {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	// Set up the request with path parameters
	req, _ := http.NewRequest(http.MethodGet, "/", nil)
	c.Request = req

	// Set path parameters
	c.Params = gin.Params{
		{Key: "chainID", Value: chainID},
		{Key: "addresses", Value: addresses},
	}

	return c, w
}

/**************************************************************************************************
** TestGraphQLHarvestRequestForOneVault tests the graphQLHarvestRequestForOneVault function.
**
** This test verifies that the function correctly constructs a GraphQL query based on
** provided vault addresses. It checks that the request is properly created.
**
** @param t *testing.T - The testing object
**************************************************************************************************/
func TestGraphQLHarvestRequestForOneVault(t *testing.T) {
	// Setup
	c, _ := setupHarvestsTest(t, "1", "0x1234,0x5678")
	vaultAddresses := []string{"0x1234", "0x5678"}

	// Execute
	request := graphQLHarvestRequestForOneVault(vaultAddresses, c)

	// Assert
	assert.NotNil(t, request)
	assert.IsType(t, &graphql.Request{}, request)
}

/**************************************************************************************************
** TestGetHarvestsForVault_InvalidChainID tests that the GetHarvestsForVault method returns an
** error when an invalid chain ID is provided.
**
** @param t *testing.T - The testing object
**************************************************************************************************/
func TestGetHarvestsForVault_InvalidChainID(t *testing.T) {
	// Setup
	c, w := setupHarvestsTest(t, "invalid", "0x1234")
	controller := Controller{}

	// Execute
	controller.GetHarvestsForVault(c)

	// Assert
	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, "invalid chainID", w.Body.String())
}
