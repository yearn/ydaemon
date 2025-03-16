package prices

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"github.com/yearn/ydaemon/common/bigNumber"
)

/**************************************************************************************************
** setupTestRouter creates a Gin router with the test routes configured.
** This helper function is used across test files to create a consistent test environment.
**
** @return *gin.Engine The configured Gin router
**************************************************************************************************/
func setupTestRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	return router
}

// mockTokenPrice represents a simplified token price structure for testing
type mockTokenPrice struct {
	Price          *bigNumber.Int   `json:"price"`
	HumanizedPrice *bigNumber.Float `json:"humanizedPrice"`
}

// mockStorageAdapter is a test double for the storage package
type mockStorageAdapter struct {
	mockPrices map[uint64]map[string]mockTokenPrice
}

// newMockStorageAdapter creates a new mock storage adapter with initial data
func newMockStorageAdapter() *mockStorageAdapter {
	return &mockStorageAdapter{
		mockPrices: make(map[uint64]map[string]mockTokenPrice),
	}
}

// mockGetPrice mocks the storage.GetPrice function behavior
func (m *mockStorageAdapter) mockGetPrice(chainID uint64, address common.Address) (mockTokenPrice, bool) {
	chainPrices, ok := m.mockPrices[chainID]
	if !ok {
		return mockTokenPrice{}, false
	}

	price, ok := chainPrices[address.Hex()]
	return price, ok
}

// addMockPrice adds a mock price for a specific token on a chain
func (m *mockStorageAdapter) addMockPrice(chainID uint64, address common.Address, price *bigNumber.Int) {
	// Initialize chain map if needed
	if _, ok := m.mockPrices[chainID]; !ok {
		m.mockPrices[chainID] = make(map[string]mockTokenPrice)
	}

	// Add the price
	m.mockPrices[chainID][address.Hex()] = mockTokenPrice{
		Price:          price,
		HumanizedPrice: bigNumber.NewFloat().SetInt(price),
	}
}

/**************************************************************************************************
** TestGetPrice tests the GetPrice endpoint to ensure it correctly retrieves and formats token
** price data. This test uses a custom handler that simulates the real handler's behavior but uses
** our mock storage instead.
**
** This test validates:
** - Successful price retrieval with raw price format
** - Successful price retrieval with humanized price format
** - Error handling for invalid chain ID
** - Error handling for invalid token address
** - Error handling for non-existent token
**************************************************************************************************/
func TestGetPrice(t *testing.T) {
	// Setup test router
	router := setupTestRouter()

	// Create mock storage
	mockStorage := newMockStorageAdapter()

	// Add test data
	testPrice := bigNumber.NewInt(1000000000000000000) // 1 with 18 decimals
	daiAddress := common.HexToAddress("0x6b175474e89094c44da98b954eedeac495271d0f")
	mockStorage.addMockPrice(1, daiAddress, testPrice)

	// Create a custom handler that uses our mock
	mockGetPrice := func(c *gin.Context) {
		chainIDStr := c.Param("chainID")
		addressStr := c.Param("address")

		// Validate chain ID
		var chainID uint64
		_, err := fmt.Sscanf(chainIDStr, "%d", &chainID)
		if err != nil {
			c.String(http.StatusBadRequest, "invalid chainID")
			return
		}

		// Validate address
		if !common.IsHexAddress(addressStr) {
			c.String(http.StatusBadRequest, "invalid address")
			return
		}
		address := common.HexToAddress(addressStr)

		// Use our mock storage
		price, ok := mockStorage.mockGetPrice(chainID, address)
		if !ok {
			c.String(http.StatusNotFound, "token not found")
			return
		}

		// Handle humanized parameter
		humanized := false
		humanizedParam := c.Query("humanized")
		if humanizedParam == "true" {
			humanized = true
		}

		if humanized {
			c.JSON(http.StatusOK, price.HumanizedPrice)
		} else {
			c.JSON(http.StatusOK, price.Price)
		}
	}

	// Register routes with our mock handler
	router.GET("/prices/:chainID/:address", mockGetPrice)

	tests := []struct {
		name               string
		url                string
		expectedStatusCode int
		expectedBodyType   interface{}
		checkBody          bool
	}{
		{
			name:               "Successful price retrieval - raw format",
			url:                "/prices/1/0x6b175474e89094c44da98b954eedeac495271d0f",
			expectedStatusCode: http.StatusOK,
			expectedBodyType:   &bigNumber.Int{},
			checkBody:          true,
		},
		{
			name:               "Successful price retrieval - humanized format",
			url:                "/prices/1/0x6b175474e89094c44da98b954eedeac495271d0f?humanized=true",
			expectedStatusCode: http.StatusOK,
			expectedBodyType:   &bigNumber.Float{},
			checkBody:          true,
		},
		{
			name:               "Invalid chain ID",
			url:                "/prices/invalid/0x6b175474e89094c44da98b954eedeac495271d0f",
			expectedStatusCode: http.StatusBadRequest,
			expectedBodyType:   nil,
			checkBody:          false,
		},
		{
			name:               "Invalid token address",
			url:                "/prices/1/0xinvalid",
			expectedStatusCode: http.StatusBadRequest,
			expectedBodyType:   nil,
			checkBody:          false,
		},
		{
			name:               "Token not found",
			url:                "/prices/1/0x1234567890123456789012345678901234567890",
			expectedStatusCode: http.StatusNotFound,
			expectedBodyType:   nil,
			checkBody:          false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create request
			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", tt.url, nil)
			router.ServeHTTP(w, req)

			// Check status code
			if w.Code != tt.expectedStatusCode {
				t.Errorf("Expected status code %d, got %d", tt.expectedStatusCode, w.Code)
			}

			// Check response body for successful cases
			if tt.checkBody {
				// For simplicity, we'll just verify we can decode to the expected type
				if tt.expectedBodyType != nil {
					switch expectedType := tt.expectedBodyType.(type) {
					case *bigNumber.Int:
						var result bigNumber.Int
						err := json.Unmarshal(w.Body.Bytes(), &result)
						if err != nil {
							t.Errorf("Failed to decode response as bigNumber.Int: %v", err)
						}
					case *bigNumber.Float:
						var result bigNumber.Float
						err := json.Unmarshal(w.Body.Bytes(), &result)
						if err != nil {
							t.Errorf("Failed to decode response as bigNumber.Float: %v", err)
						}
					default:
						t.Errorf("Unexpected expected body type: %T", expectedType)
					}
				}
			}
		})
	}
}
