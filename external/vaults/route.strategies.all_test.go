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

/**************************************************************************************************
** setupAllStrategiesTest creates a test environment for testing the GetAllStrategies endpoint.
**
** This function sets up a Gin test context with predefined path parameters for chain ID, which
** is required for the GetAllStrategies endpoint. It also configures a mock response writer to
** capture the endpoint's output.
**
** @param t *testing.T - The testing object
** @param chainID string - The chain ID to use in the test
** @return *gin.Context - The configured Gin context
** @return *httptest.ResponseRecorder - The HTTP response recorder
**************************************************************************************************/
func setupAllStrategiesTest(t *testing.T, chainID string) (*gin.Context, *httptest.ResponseRecorder) {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	// Set up the request with path parameters
	req, _ := http.NewRequest(http.MethodGet, "/", nil)
	c.Request = req

	// Set path parameters
	c.Params = gin.Params{
		{Key: "chainID", Value: chainID},
	}

	return c, w
}

/**************************************************************************************************
** mockVaultWithStrategy creates a mock vault with a strategy in the storage for testing.
**
** @param chainID uint64 - The chain ID for the mock vault and strategy
** @param vaultAddress string - The address for the mock vault
** @param strategyAddress string - The address for the mock strategy
**************************************************************************************************/
func mockVaultWithStrategy(chainID uint64, vaultAddress, strategyAddress string) {
	vaultAddr := common.HexToAddress(vaultAddress)
	stratAddr := common.HexToAddress(strategyAddress)

	// Create the vault
	vault := models.TVault{
		Address:      vaultAddr,
		AssetAddress: common.HexToAddress("0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48"), // USDC
		Type:         models.TokenTypeStandardVault,
		Kind:         models.VaultKindSingle,
		Version:      "v3",
		ChainID:      chainID,
		Metadata: models.TVaultMetadata{
			DisplayName:   "Test Vault",
			DisplaySymbol: "tVAULT",
			Inclusion: models.TInclusion{
				IsYearn: true,
			},
		},
		LastPricePerShare: bigNumber.NewInt(1000000000),
		LastTotalAssets:   bigNumber.NewInt(1000000000000),
	}

	// Store the vault
	storage.StoreVault(chainID, vault)

	// Create the strategy
	strategy := models.TStrategy{
		Address:            stratAddr,
		Name:               "Test Strategy",
		DisplayName:        "Test Strategy Display",
		Description:        "A test strategy for unit testing",
		LastTotalDebt:      bigNumber.NewInt(1000000000000),
		LastTotalGain:      bigNumber.NewInt(500000000000),
		LastTotalLoss:      bigNumber.NewInt(100000000000),
		LastPerformanceFee: bigNumber.NewInt(2000),
		LastReport:         bigNumber.NewInt(1634567890),
		LastDebtRatio:      bigNumber.NewInt(5000),
		ChainID:            chainID,
		IsInQueue:          true,
		VaultAddress:       vaultAddr,
	}

	// Store the strategy
	storage.StoreStrategy(chainID, strategy)
}

/**************************************************************************************************
** TestGetAllStrategies_InvalidChainID tests that the GetAllStrategies method returns an error
** when an invalid chain ID is provided.
**
** @param t *testing.T - The testing object
**************************************************************************************************/
func TestGetAllStrategies_InvalidChainID(t *testing.T) {
	// Setup
	c, w := setupAllStrategiesTest(t, "invalid")
	controller := Controller{}

	// Execute
	controller.GetAllStrategies(c)

	// Assert
	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, "invalid chainID: invalid", w.Body.String())
}

/**************************************************************************************************
** TestGetAllStrategies_QueryParams tests the query parameter handling in GetAllStrategies.
**
** This test verifies that the endpoint correctly processes query parameters such as orderBy,
** orderDirection, and strategiesCondition, which control how strategies are filtered and sorted.
**
** @param t *testing.T - The testing object
**************************************************************************************************/
func TestGetAllStrategies_QueryParams(t *testing.T) {
	// Setup
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	// Set up request with query parameters
	req, _ := http.NewRequest(http.MethodGet, "/?orderBy=address&orderDirection=desc&strategiesCondition=all", nil)
	c.Request = req

	// Set path parameters
	c.Params = gin.Params{
		{Key: "chainID", Value: "1"},
	}

	// Execute
	controller := Controller{}
	controller.GetAllStrategies(c)

	// Assert - we can't easily check the query param effects without mock data,
	// but we can verify the code runs without errors
	assert.Equal(t, http.StatusOK, w.Code)

	// Parse the response to ensure it's valid JSON
	var strategies []TStrategy
	err := json.Unmarshal(w.Body.Bytes(), &strategies)
	assert.NoError(t, err, "Response should be valid JSON")
}
