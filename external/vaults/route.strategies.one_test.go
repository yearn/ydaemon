package vaults

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/yearn/ydaemon/common/bigNumber"
	"github.com/yearn/ydaemon/common/helpers"
	"github.com/yearn/ydaemon/internal/models"
	"github.com/yearn/ydaemon/internal/storage"
)

/**************************************************************************************************
** setupStrategyTest creates a test environment for testing the strategy endpoints.
**
** This function sets up a Gin test context with predefined path parameters for chain ID and
** strategy address, which are required for the strategy endpoint. It also configures a mock
** response writer to capture the endpoint's output.
**
** @param t *testing.T - The testing object
** @param chainID string - The chain ID to use in the test
** @param address string - The strategy address to use in the test
** @return *gin.Context - The configured Gin context
** @return *httptest.ResponseRecorder - The HTTP response recorder
**************************************************************************************************/
func setupStrategyTest(t *testing.T, chainID, address string) (*gin.Context, *httptest.ResponseRecorder) {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	// Set up the request with path parameters
	req, _ := http.NewRequest(http.MethodGet, "/", nil)
	c.Request = req

	// Set path parameters
	c.Params = gin.Params{
		{Key: "chainID", Value: chainID},
		{Key: "address", Value: address},
	}

	return c, w
}

/**************************************************************************************************
** mockStrategy creates a mock strategy in the storage for testing.
**
** @param chainID uint64 - The chain ID for the mock strategy
** @param address string - The address for the mock strategy
** @return models.TStrategy - The created mock strategy
**************************************************************************************************/
func mockStrategy(chainID uint64, address string) models.TStrategy {
	stratAddr := common.HexToAddress(address)
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
	}

	// Store the strategy in storage
	storage.StoreStrategy(chainID, strategy)

	return strategy
}

/**************************************************************************************************
** TestGetStrategy tests the GetStrategy method of the Controller struct.
**
** This test verifies that the endpoint correctly handles various input scenarios, including:
** - Invalid chain ID
** - Invalid address format
** - Non-existent strategy
** - Valid strategy retrieval
**
** @param t *testing.T - The testing object
**************************************************************************************************/
func TestGetStrategy(t *testing.T) {
	tests := []struct {
		name             string
		chainID          string
		address          string
		setupMock        bool
		expectedStatus   int
		validateResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name:           "Invalid chain ID",
			chainID:        "invalid",
			address:        "0x1234567890123456789012345678901234567890",
			setupMock:      false,
			expectedStatus: http.StatusBadRequest,
			validateResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				assert.Equal(t, "invalid chainID: invalid", recorder.Body.String())
			},
		},
		{
			name:           "Invalid address",
			chainID:        "1",
			address:        "invalid",
			setupMock:      false,
			expectedStatus: http.StatusBadRequest,
			validateResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				assert.Equal(t, "invalid address: invalid", recorder.Body.String())
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Setup
			c, w := setupStrategyTest(t, tc.chainID, tc.address)

			if tc.setupMock {
				chainID, _ := helpers.AssertChainID(tc.chainID)
				mockStrategy(chainID, tc.address)
			}

			controller := Controller{}

			// Execute
			controller.GetStrategy(c)

			// Assert
			assert.Equal(t, tc.expectedStatus, w.Code)
			tc.validateResponse(t, w)
		})
	}
}
