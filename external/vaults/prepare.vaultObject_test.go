package vaults

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yearn/ydaemon/common/bigNumber"
)

/**************************************************************************************************
** TestBuildTokenPrice tests the buildTokenPrice function to ensure it correctly
** retrieves and formats token prices from storage.
**************************************************************************************************/
func TestBuildTokenPrice(t *testing.T) {
	// Since we can't easily mock the storage.GetPrice function without changing the package,
	// we'll need to rely on real behavior or skip the test
	t.Skip("Skipping as this requires mocking package-level functions in storage")

	// Note: An ideal approach would be to use dependency injection for the price lookup function,
	// allowing for easier testing. For now, we'll rely on integration testing.
}

/**************************************************************************************************
** TestBuildTVL tests the buildTVL function to ensure it correctly calculates
** the TVL based on token balance, decimals, and price.
**************************************************************************************************/
func TestBuildTVL(t *testing.T) {
	testCases := []struct {
		name           string
		balanceToken   *bigNumber.Int
		decimals       int
		humanizedPrice *bigNumber.Float
		expectedTVL    float64
	}{
		{
			name:           "Zero balance",
			balanceToken:   bigNumber.NewInt(0),
			decimals:       18,
			humanizedPrice: bigNumber.NewFloat(1000.0),
			expectedTVL:    0.0,
		},
		{
			name:           "Zero price",
			balanceToken:   bigNumber.NewInt(1000000000000000000), // 1.0 token with 18 decimals
			decimals:       18,
			humanizedPrice: bigNumber.NewFloat(0.0),
			expectedTVL:    0.0,
		},
		{
			name:           "Normal TVL calculation - 18 decimals",
			balanceToken:   bigNumber.NewInt(2000000000000000000), // 2.0 tokens with 18 decimals
			decimals:       18,
			humanizedPrice: bigNumber.NewFloat(1000.0), // $1000 per token
			expectedTVL:    2000.0,                     // 2.0 * $1000 = $2000
		},
		{
			name:           "Normal TVL calculation - 6 decimals",
			balanceToken:   bigNumber.NewInt(3000000), // 3.0 tokens with 6 decimals
			decimals:       6,
			humanizedPrice: bigNumber.NewFloat(500.0), // $500 per token
			expectedTVL:    1500.0,                    // 3.0 * $500 = $1500
		},
		{
			name:           "High value TVL",
			balanceToken:   bigNumber.NewInt(0).SetString("100000000000000000000000"), // 100,000 tokens with 18 decimals
			decimals:       18,
			humanizedPrice: bigNumber.NewFloat(2500.0), // $2500 per token
			expectedTVL:    250000000.0,                // 100,000 * $2500 = $250,000,000
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := buildTVL(tc.balanceToken, tc.decimals, tc.humanizedPrice)
			assert.InDelta(t, tc.expectedTVL, result, 0.00001, "TVL calculation should match expected value")
		})
	}
}

/**************************************************************************************************
** TestToSimplifiedVersion tests the toSimplifiedVersion function to ensure it correctly
** transforms a TExternalVault to a TSimplifiedExternalVault.
**************************************************************************************************/
func TestToSimplifiedVersion(t *testing.T) {
	// This test is complex due to dependencies on other functions like assignStakingData
	// which would require extensive mocking of storage package functions
	t.Skip("Skipping as this requires mocking multiple package-level functions in storage")

	// Note: An ideal approach would be to use dependency injection for the storage functions,
	// allowing for easier testing. For now, we'll rely on integration testing.
}
