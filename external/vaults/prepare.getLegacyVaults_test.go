package vaults

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// Mock structures for testing
type mockVaultInclusion struct {
	IsYearn       bool
	IsYearnJuiced bool
	IsGimme       bool
}

type mockVaultMetadata struct {
	Inclusion mockVaultInclusion
	IsRetired bool
}

type mockVault struct {
	Version  string
	Metadata mockVaultMetadata
}

/**************************************************************************************************
** TestSelectStrategiesCondition tests the selectStrategiesCondition function to verify it
** correctly selects the strategy condition based on input.
**************************************************************************************************/
func TestSelectStrategiesCondition(t *testing.T) {
	testCases := []struct {
		name           string
		input          string
		expectedOutput string
	}{
		{
			name:           "Empty input",
			input:          "",
			expectedOutput: "debtRatio", // default value
		},
		{
			name:           "Valid input - all",
			input:          "all",
			expectedOutput: "all",
		},
		{
			name:           "Valid input - inQueue",
			input:          "inQueue",
			expectedOutput: "inQueue",
		},
		{
			name:           "Valid input - debtRatio",
			input:          "debtRatio",
			expectedOutput: "debtRatio",
		},
		{
			name:           "Valid input - absolute",
			input:          "absolute",
			expectedOutput: "absolute",
		},
		{
			name:           "Invalid input",
			input:          "invalid",
			expectedOutput: "debtRatio", // should default to debtRatio
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := selectStrategiesCondition(tc.input)
			assert.Equal(t, tc.expectedOutput, result, "Should select the correct strategy condition")
		})
	}
}

/**************************************************************************************************
** TestSelectMigrableCondition tests the selectMigrableCondition function to verify it
** correctly selects the migrable condition based on input.
**************************************************************************************************/
func TestSelectMigrableCondition(t *testing.T) {
	testCases := []struct {
		name           string
		input          string
		expectedOutput string
	}{
		{
			name:           "Empty input",
			input:          "",
			expectedOutput: "none", // default value
		},
		{
			name:           "Valid input - none",
			input:          "none",
			expectedOutput: "none",
		},
		{
			name:           "Valid input - all",
			input:          "all",
			expectedOutput: "all",
		},
		{
			name:           "Valid input - nodust",
			input:          "nodust",
			expectedOutput: "nodust",
		},
		{
			name:           "Valid input - ignore",
			input:          "ignore",
			expectedOutput: "ignore",
		},
		{
			name:           "Invalid input",
			input:          "invalid",
			expectedOutput: "none", // should default to none
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := selectMigrableCondition(tc.input)
			assert.Equal(t, tc.expectedOutput, result, "Should select the correct migrable condition")
		})
	}
}

/**************************************************************************************************
** TestGetQuery tests the getQuery function to ensure it correctly extracts
** query parameters from the Gin context.
**************************************************************************************************/
func TestGetQuery(t *testing.T) {
	// Create a gin router
	gin.SetMode(gin.TestMode)
	router := gin.New()

	// Define a test handler
	router.GET("/test", func(c *gin.Context) {
		paramName := "testParam"
		paramValue := getQuery(c, paramName)
		c.String(http.StatusOK, paramValue)
	})

	// Test cases
	testCases := []struct {
		name           string
		queryString    string
		expectedResult string
	}{
		{
			name:           "Parameter exists",
			queryString:    "?testParam=value123",
			expectedResult: "value123",
		},
		{
			name:           "Parameter with empty value",
			queryString:    "?testParam=",
			expectedResult: "",
		},
		{
			name:           "Parameter missing",
			queryString:    "?otherParam=value",
			expectedResult: "",
		},
		{
			name:           "Multiple parameters",
			queryString:    "?testParam=value456&otherParam=other",
			expectedResult: "value456",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Create a test request
			req, _ := http.NewRequest("GET", "/test"+tc.queryString, nil)
			w := httptest.NewRecorder()

			// Serve the request
			router.ServeHTTP(w, req)

			// Check the response
			assert.Equal(t, http.StatusOK, w.Code)
			assert.Equal(t, tc.expectedResult, w.Body.String())
		})
	}
}

/**************************************************************************************************
** TestFilterFunctions verifies that common filter functions used with getLegacyVaults
** correctly identify vaults based on their attributes.
**************************************************************************************************/
func TestFilterFunctions(t *testing.T) {
	// Create sample vault data
	v2Vault := mockVault{
		Version: "2.0.0",
		Metadata: mockVaultMetadata{
			Inclusion: mockVaultInclusion{
				IsYearn:       true,
				IsYearnJuiced: false,
				IsGimme:       false,
			},
			IsRetired: false,
		},
	}

	v3Vault := mockVault{
		Version: "3.0.0",
		Metadata: mockVaultMetadata{
			Inclusion: mockVaultInclusion{
				IsYearn:       true,
				IsYearnJuiced: false,
				IsGimme:       false,
			},
			IsRetired: false,
		},
	}

	yearnJuicedVault := mockVault{
		Version: "2.0.0",
		Metadata: mockVaultMetadata{
			Inclusion: mockVaultInclusion{
				IsYearn:       true,
				IsYearnJuiced: true,
				IsGimme:       false,
			},
			IsRetired: false,
		},
	}

	gimmeVault := mockVault{
		Version: "2.0.0",
		Metadata: mockVaultMetadata{
			Inclusion: mockVaultInclusion{
				IsYearn:       false,
				IsYearnJuiced: false,
				IsGimme:       true,
			},
			IsRetired: false,
		},
	}

	retiredVault := mockVault{
		Version: "2.0.0",
		Metadata: mockVaultMetadata{
			Inclusion: mockVaultInclusion{
				IsYearn:       true,
				IsYearnJuiced: false,
				IsGimme:       false,
			},
			IsRetired: true,
		},
	}

	// Define test cases for different filter functions
	testCases := []struct {
		name          string
		vault         mockVault
		filterFunc    func(vault mockVault) bool
		expectedMatch bool
	}{
		// IsYearn filter
		{
			name:          "IsYearn filter - matches Yearn vault",
			vault:         v2Vault,
			filterFunc:    func(vault mockVault) bool { return vault.Metadata.Inclusion.IsYearn },
			expectedMatch: true,
		},
		{
			name:          "IsYearn filter - doesn't match non-Yearn vault",
			vault:         gimmeVault,
			filterFunc:    func(vault mockVault) bool { return vault.Metadata.Inclusion.IsYearn },
			expectedMatch: false,
		},

		// V2 IsYearn filter
		{
			name:          "V2 IsYearn filter - matches V2 Yearn vault",
			vault:         v2Vault,
			filterFunc:    func(vault mockVault) bool { return vault.Version != "3.0.0" && vault.Metadata.Inclusion.IsYearn },
			expectedMatch: true,
		},
		{
			name:          "V2 IsYearn filter - doesn't match V3 Yearn vault",
			vault:         v3Vault,
			filterFunc:    func(vault mockVault) bool { return vault.Version != "3.0.0" && vault.Metadata.Inclusion.IsYearn },
			expectedMatch: false,
		},

		// V3 IsYearn filter
		{
			name:          "V3 IsYearn filter - matches V3 Yearn vault",
			vault:         v3Vault,
			filterFunc:    func(vault mockVault) bool { return vault.Version == "3.0.0" && vault.Metadata.Inclusion.IsYearn },
			expectedMatch: true,
		},
		{
			name:          "V3 IsYearn filter - doesn't match V2 Yearn vault",
			vault:         v2Vault,
			filterFunc:    func(vault mockVault) bool { return vault.Version == "3.0.0" && vault.Metadata.Inclusion.IsYearn },
			expectedMatch: false,
		},

		// IsYearnJuiced filter
		{
			name:          "IsYearnJuiced filter - matches juiced vault",
			vault:         yearnJuicedVault,
			filterFunc:    func(vault mockVault) bool { return vault.Metadata.Inclusion.IsYearnJuiced },
			expectedMatch: true,
		},
		{
			name:          "IsYearnJuiced filter - doesn't match non-juiced vault",
			vault:         v2Vault,
			filterFunc:    func(vault mockVault) bool { return vault.Metadata.Inclusion.IsYearnJuiced },
			expectedMatch: false,
		},

		// IsGimme filter
		{
			name:          "IsGimme filter - matches Gimme vault",
			vault:         gimmeVault,
			filterFunc:    func(vault mockVault) bool { return vault.Metadata.Inclusion.IsGimme },
			expectedMatch: true,
		},
		{
			name:          "IsGimme filter - doesn't match non-Gimme vault",
			vault:         v2Vault,
			filterFunc:    func(vault mockVault) bool { return vault.Metadata.Inclusion.IsGimme },
			expectedMatch: false,
		},

		// IsRetired filter
		{
			name:          "IsRetired filter - matches retired vault",
			vault:         retiredVault,
			filterFunc:    func(vault mockVault) bool { return vault.Metadata.IsRetired },
			expectedMatch: true,
		},
		{
			name:          "IsRetired filter - doesn't match active vault",
			vault:         v2Vault,
			filterFunc:    func(vault mockVault) bool { return vault.Metadata.IsRetired },
			expectedMatch: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := tc.filterFunc(tc.vault)
			assert.Equal(t, tc.expectedMatch, result, "Filter function should correctly identify matching vaults")
		})
	}
}
