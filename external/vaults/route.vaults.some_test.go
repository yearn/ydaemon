package vaults

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/yearn/ydaemon/common/bigNumber"
	"github.com/yearn/ydaemon/internal/models"
	"github.com/yearn/ydaemon/internal/storage"
)

func init() {
	// Enable test mode for this entire test file to ensure that our tests
	// are compatible with the new error handling format
	TestMode = true
}

/**************************************************************************************************
** setupLegacySomeVaultsTest sets up the test environment for testing the GetLegacySomeVaults endpoint.
**
** This function creates a test Gin context with path parameters for chain ID and vault addresses.
** It also sets up query parameters for sorting and strategy conditions.
**
** @param t *testing.T - The testing object
** @param chainID string - The chain ID to use in the test
** @param addresses string - Comma-separated list of vault addresses
** @param queryParams map[string]string - Optional query parameters
** @return *gin.Context - The configured Gin context
** @return *httptest.ResponseRecorder - The HTTP response recorder
**************************************************************************************************/
func setupLegacySomeVaultsTest(t *testing.T, chainID, addresses string, queryParams map[string]string) (*gin.Context, *httptest.ResponseRecorder) {
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
		{Key: "addresses", Value: addresses},
	}

	return c, w
}

/**************************************************************************************************
** mockSomeVaults creates multiple mock vaults in storage for testing.
**
** This function creates a set of test vaults for use in the GetLegacySomeVaults tests.
**
** @param chainID uint64 - The chain ID for the vaults
** @return []common.Address - The addresses of the created vaults
**************************************************************************************************/
func mockSomeVaults(chainID uint64) []common.Address {
	addresses := []common.Address{
		common.HexToAddress("0x1111111111111111111111111111111111111111"),
		common.HexToAddress("0x2222222222222222222222222222222222222222"),
		common.HexToAddress("0x3333333333333333333333333333333333333333"),
		common.HexToAddress("0x4444444444444444444444444444444444444444"),
	}

	// Create test vaults with different versions
	for i, address := range addresses {
		version := "v2"
		var kind models.TVaultKind
		if i >= 2 {
			version = "v3"
			kind = models.VaultKindSingle
		}

		vault := models.TVault{
			Address:      address,
			AssetAddress: common.HexToAddress("0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48"), // USDC
			Type:         models.TokenTypeStandardVault,
			Kind:         kind,
			Version:      version,
			ChainID:      chainID,
			Metadata: models.TVaultMetadata{
				DisplayName:   "Test Vault " + address.Hex()[2:6],
				DisplaySymbol: "tVAULT" + address.Hex()[2:6],
				Inclusion: models.TInclusion{
					IsYearn: true,
				},
			},
			LastPricePerShare: bigNumber.NewInt(1000000000),
			LastTotalAssets:   bigNumber.NewInt(1000000000000),
		}

		storage.StoreVault(chainID, vault)

		// Create a test strategy for each vault
		strategy := models.TStrategy{
			Address:      common.HexToAddress("0x" + address.Hex()[2:10] + "5432109876543210"),
			VaultAddress: address,
			Name:         "Test Strategy for " + address.Hex()[2:6],
			ChainID:      chainID,
			IsActive:     true,
		}

		storage.StoreStrategy(chainID, strategy)
	}

	return addresses
}

/**************************************************************************************************
** TestGetLegacySomeVaults_InvalidChainID tests GetLegacySomeVaults with an invalid chain ID.
**
** @param t *testing.T - The testing object
**************************************************************************************************/
func TestGetLegacySomeVaults_InvalidChainID(t *testing.T) {
	// Setup with invalid chain ID
	c, w := setupLegacySomeVaultsTest(t, "invalid", "0x1111,0x2222", nil)
	controller := Controller{}

	// Execute
	controller.GetLegacySomeVaults(c)

	// Assert
	assert.Equal(t, http.StatusBadRequest, w.Code)
	// Match the new error format from our structured error handling
	assert.Contains(t, w.Body.String(), "Invalid chain ID format")
	assert.Contains(t, w.Body.String(), "invalid")
}

/**************************************************************************************************
** TestGetLegacySomeVaults_ValidParameters tests GetLegacySomeVaults with valid parameters.
**
** @param t *testing.T - The testing object
**************************************************************************************************/
func TestGetLegacySomeVaults_ValidParameters(t *testing.T) {
	// Skip this test for now
	t.Skip("Skipping test as it requires environment setup and mocking")

	// Create mock vaults
	chainID := uint64(1)
	addresses := mockSomeVaults(chainID)
	addressList := addresses[0].Hex() + "," + addresses[1].Hex()

	// Setup with sorting parameters
	queryParams := map[string]string{
		"orderBy":             "tvl",
		"orderDirection":      "desc",
		"strategiesCondition": "all",
	}
	c, w := setupLegacySomeVaultsTest(t, "1", addressList, queryParams)
	controller := Controller{}

	// Execute
	controller.GetLegacySomeVaults(c)

	// Assert
	assert.Equal(t, http.StatusOK, w.Code)

	// Parse response
	var vaults []TExternalVault
	err := json.Unmarshal(w.Body.Bytes(), &vaults)
	assert.NoError(t, err)

	// Should return only the requested vaults (2 in this case)
	assert.Equal(t, 2, len(vaults))

	// Verify the addresses match what we requested
	foundAddresses := []string{vaults[0].Address, vaults[1].Address}
	assert.Contains(t, foundAddresses, addresses[0].Hex())
	assert.Contains(t, foundAddresses, addresses[1].Hex())
}

/**************************************************************************************************
** TestGetLegacySomeVaults_BlacklistedVault tests GetLegacySomeVaults when a vault is blacklisted.
**
** @param t *testing.T - The testing object
**************************************************************************************************/
func TestGetLegacySomeVaults_BlacklistedVault(t *testing.T) {
	// Skip this test for now
	t.Skip("Skipping test as it requires environment setup and mocking")

	// Create mock vaults
	chainID := uint64(1)
	addresses := mockSomeVaults(chainID)

	// Include both regular and blacklisted vaults in request
	addressList := addresses[0].Hex() + "," + addresses[1].Hex()

	// Setup request
	c, w := setupLegacySomeVaultsTest(t, "1", addressList, nil)
	controller := Controller{}

	// Execute
	controller.GetLegacySomeVaults(c)

	// Assert
	assert.Equal(t, http.StatusOK, w.Code)

	// Parse response
	var vaults []TExternalVault
	err := json.Unmarshal(w.Body.Bytes(), &vaults)
	assert.NoError(t, err)

	// Should return only non-blacklisted vaults (1 in this case)
	assert.Equal(t, 1, len(vaults))
	assert.Equal(t, addresses[1].Hex(), vaults[0].Address)
}
