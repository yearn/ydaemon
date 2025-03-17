package vaults

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	"github.com/yearn/ydaemon/common/bigNumber"
	"github.com/yearn/ydaemon/internal/models"
)

/**************************************************************************************************
** TestCreateExternalStrategy tests the CreateExternalStrategy function to verify it properly converts
** an internal strategy model to the external TStrategy format.
**************************************************************************************************/
func TestCreateExternalStrategy(t *testing.T) {
	// Create test data
	address := common.HexToAddress("0x1234567890123456789012345678901234567890")
	totalDebt := bigNumber.NewInt(10000)
	totalLoss := bigNumber.NewInt(100)
	totalGain := bigNumber.NewInt(500)
	performanceFee := bigNumber.NewInt(2000) // 20%
	lastReport := bigNumber.NewInt(1234567890)
	debtRatio := bigNumber.NewInt(10000) // 100%

	internalStrategy := models.TStrategy{
		Address:            address,
		Name:               "Test Strategy",
		DisplayName:        "Test Strategy Display",
		Description:        "Test Strategy Description",
		LastTotalDebt:      totalDebt,
		LastTotalLoss:      totalLoss,
		LastTotalGain:      totalGain,
		LastPerformanceFee: performanceFee,
		LastReport:         lastReport,
		LastDebtRatio:      debtRatio,
		IsInQueue:          true,
	}

	// Test with display name present
	t.Run("With display name", func(t *testing.T) {
		strategy := CreateExternalStrategy(internalStrategy)

		// Verify the result
		assert.Equal(t, address.Hex(), strategy.Address, "Address should match")
		assert.Equal(t, internalStrategy.DisplayName, strategy.Name, "Name should match DisplayName")
		assert.Equal(t, internalStrategy.Description, strategy.Description, "Description should match")
		assert.NotNil(t, strategy.Details, "Details should not be nil")
		assert.True(t, strategy.Details.TotalDebt.Eq(totalDebt), "TotalDebt should match")
		assert.True(t, strategy.Details.TotalLoss.Eq(totalLoss), "TotalLoss should match")
		assert.True(t, strategy.Details.TotalGain.Eq(totalGain), "TotalGain should match")
		assert.Equal(t, performanceFee.Uint64(), strategy.Details.PerformanceFee, "PerformanceFee should match")
		assert.Equal(t, lastReport.Uint64(), strategy.Details.LastReport, "LastReport should match")
		assert.Equal(t, debtRatio.Uint64(), strategy.Details.DebtRatio, "DebtRatio should match")
		assert.Equal(t, internalStrategy.IsInQueue, strategy.Details.InQueue, "InQueue should match")
	})

	// Test with no display name
	t.Run("Without display name", func(t *testing.T) {
		internalStrategyNoDisplay := internalStrategy
		internalStrategyNoDisplay.DisplayName = ""

		strategy := CreateExternalStrategy(internalStrategyNoDisplay)

		// Verify the result
		assert.Equal(t, internalStrategyNoDisplay.Name, strategy.Name, "Name should fall back to Name when DisplayName is empty")
	})
}

/**************************************************************************************************
** TestStrategyBeShouldIncluded tests the ShouldBeIncluded method to verify it correctly
** determines whether a strategy should be included based on different conditions.
**************************************************************************************************/
func TestStrategyShouldBeIncluded(t *testing.T) {
	// Create various test cases
	testCases := []struct {
		name           string
		strategy       TStrategy
		condition      string
		expectedResult bool
	}{
		{
			name: "All condition",
			strategy: TStrategy{
				Details: &TExternalStrategyDetails{
					TotalDebt: bigNumber.NewInt(0),
					DebtRatio: 0,
					InQueue:   false,
				},
			},
			condition:      "all",
			expectedResult: true,
		},
		{
			name: "Absolute condition with debt",
			strategy: TStrategy{
				Details: &TExternalStrategyDetails{
					TotalDebt: bigNumber.NewInt(100),
					DebtRatio: 0,
					InQueue:   false,
				},
			},
			condition:      "absolute",
			expectedResult: true,
		},
		{
			name: "Absolute condition without debt",
			strategy: TStrategy{
				Details: &TExternalStrategyDetails{
					TotalDebt: bigNumber.NewInt(0),
					DebtRatio: 0,
					InQueue:   false,
				},
			},
			condition:      "absolute",
			expectedResult: false,
		},
		{
			name: "InQueue condition with strategy in queue",
			strategy: TStrategy{
				Details: &TExternalStrategyDetails{
					TotalDebt: bigNumber.NewInt(0),
					DebtRatio: 0,
					InQueue:   true,
				},
			},
			condition:      "inQueue",
			expectedResult: true,
		},
		{
			name: "InQueue condition with strategy not in queue",
			strategy: TStrategy{
				Details: &TExternalStrategyDetails{
					TotalDebt: bigNumber.NewInt(0),
					DebtRatio: 0,
					InQueue:   false,
				},
			},
			condition:      "inQueue",
			expectedResult: false,
		},
		{
			name: "DebtRatio condition with positive debt ratio",
			strategy: TStrategy{
				Details: &TExternalStrategyDetails{
					TotalDebt: bigNumber.NewInt(0),
					DebtRatio: 5000,
					InQueue:   false,
				},
			},
			condition:      "debtRatio",
			expectedResult: true,
		},
		{
			name: "DebtRatio condition with zero debt ratio",
			strategy: TStrategy{
				Details: &TExternalStrategyDetails{
					TotalDebt: bigNumber.NewInt(0),
					DebtRatio: 0,
					InQueue:   false,
				},
			},
			condition:      "debtRatio",
			expectedResult: false,
		},
		{
			name: "Unknown condition",
			strategy: TStrategy{
				Details: &TExternalStrategyDetails{
					TotalDebt: bigNumber.NewInt(100),
					DebtRatio: 5000,
					InQueue:   true,
				},
			},
			condition:      "unknown",
			expectedResult: false,
		},
	}

	// Run test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := tc.strategy.ShouldBeIncluded(tc.condition)
			assert.Equal(t, tc.expectedResult, result, "ShouldBeIncluded result should match expected value")
		})
	}
}
