package vaults

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/yearn/ydaemon/common/env"
)

/**************************************************************************************************
** setupBlacklistedVaultsTest creates a test environment for testing the blacklisted vaults endpoint.
**
** This function sets up a Gin test context with optional query parameters for chain ID.
** It configures a mock response writer to capture the endpoint's output.
**
** @param t *testing.T - The testing object
** @param chainID string - Optional chain ID to use in the test
** @return *gin.Context - The configured Gin context
** @return *httptest.ResponseRecorder - The HTTP response recorder
**************************************************************************************************/
func setupBlacklistedVaultsTest(t *testing.T, chainID string) (*gin.Context, *httptest.ResponseRecorder) {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	// Set up the request with query parameters if provided
	req, _ := http.NewRequest(http.MethodGet, "/", nil)

	if chainID != "" {
		req.URL.RawQuery = "chainID=" + chainID
	}

	c.Request = req
	return c, w
}

/**************************************************************************************************
** TestGetBlacklistedVaults_AllChains tests the GetBlacklistedVaults function with no chain ID.
**
** This test verifies that the function correctly returns blacklisted vaults across all chains
** when no specific chain ID is provided or when chainID=0.
**
** @param t *testing.T - The testing object
**************************************************************************************************/
func TestGetBlacklistedVaults_AllChains(t *testing.T) {
	// Setup
	c, w := setupBlacklistedVaultsTest(t, "")
	controller := Controller{}

	// Execute
	controller.GetBlacklistedVaults(c)

	// Assert
	assert.Equal(t, http.StatusOK, w.Code)

	// Parse response body
	var blacklistedVaults []common.Address
	err := json.Unmarshal(w.Body.Bytes(), &blacklistedVaults)
	assert.NoError(t, err)

	// Verify we have blacklisted vaults from all chains combined
	expectedVaultCount := 0
	for _, chain := range env.GetChains() {
		expectedVaultCount += len(chain.BlacklistedVaults)
	}

	// The response should include blacklisted vaults from all chains
	assert.LessOrEqual(t, expectedVaultCount, len(blacklistedVaults),
		"Expected at least %d blacklisted vaults across all chains", expectedVaultCount)
}

/**************************************************************************************************
** TestGetBlacklistedVaults_SpecificChain tests the GetBlacklistedVaults function with a valid chain ID.
**
** This test verifies that the function correctly returns blacklisted vaults for a specific chain
** when a valid chain ID is provided.
**
** @param t *testing.T - The testing object
**************************************************************************************************/
func TestGetBlacklistedVaults_SpecificChain(t *testing.T) {
	// Use Ethereum mainnet for testing
	chainID := "1"

	// Setup
	c, w := setupBlacklistedVaultsTest(t, chainID)
	controller := Controller{}

	// Execute
	controller.GetBlacklistedVaults(c)

	// Assert
	assert.Equal(t, http.StatusOK, w.Code)

	// Parse response body
	var blacklistedVaults []common.Address
	err := json.Unmarshal(w.Body.Bytes(), &blacklistedVaults)
	assert.NoError(t, err)

	// Get the expected blacklisted vaults for this chain
	chain, ok := env.GetChain(1)
	if !ok {
		t.Skip("Ethereum mainnet chain configuration not found, skipping test")
	}

	// The response should include exactly the blacklisted vaults for the specified chain
	assert.Equal(t, len(chain.BlacklistedVaults), len(blacklistedVaults),
		"Expected %d blacklisted vaults for chain ID %s", len(chain.BlacklistedVaults), chainID)
}

/**************************************************************************************************
** TestGetBlacklistedVaults_InvalidChain tests the GetBlacklistedVaults function with an invalid chain ID.
**
** This test verifies that the function correctly returns a 404 Not Found response when
** an invalid chain ID is provided.
**
** @param t *testing.T - The testing object
**************************************************************************************************/
func TestGetBlacklistedVaults_InvalidChain(t *testing.T) {
	// Setup with an invalid chain ID
	c, w := setupBlacklistedVaultsTest(t, "99999")
	controller := Controller{}

	// Execute
	controller.GetBlacklistedVaults(c)

	// Assert
	assert.Equal(t, http.StatusNotFound, w.Code)

	// Verify error message
	var response map[string]string
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "chain not found", response["error"])
}
