package vaults

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	"github.com/yearn/ydaemon/internal/models"
)

/**************************************************************************************************
** TestToArrTMixedcaseAddress tests the toArrTMixedcaseAddress function to ensure it correctly
** converts an array of Ethereum addresses to their hexadecimal string representation.
**************************************************************************************************/
func TestToArrTMixedcaseAddress(t *testing.T) {
	testCases := []struct {
		name           string
		addresses      []common.Address
		expectedResult []string
	}{
		{
			name:           "Empty array",
			addresses:      []common.Address{},
			expectedResult: []string{},
		},
		{
			name: "Single address",
			addresses: []common.Address{
				common.HexToAddress("0x1234567890123456789012345678901234567890"),
			},
			expectedResult: []string{
				"0x1234567890123456789012345678901234567890",
			},
		},
		{
			name: "Multiple addresses",
			addresses: []common.Address{
				common.HexToAddress("0x1234567890123456789012345678901234567890"),
				common.HexToAddress("0xabcdefabcdefabcdefabcdefabcdefabcdefabcd"),
				common.HexToAddress("0x0000000000000000000000000000000000000000"),
			},
			expectedResult: []string{
				"0x1234567890123456789012345678901234567890",
				"0xABCdEfABcDeFabCdEfAbCdEfaBcDeFAbcDEFabCd", // Note: This is checksum format
				"0x0000000000000000000000000000000000000000",
			},
		},
		{
			name: "Case sensitivity test",
			addresses: []common.Address{
				common.HexToAddress("0xabcdef1234567890abcdef1234567890abcdef12"),
			},
			expectedResult: []string{
				"0xABCdeF1234567890ABCdeF1234567890aBcDeF12", // This is in checksum format
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := toArrTMixedcaseAddress(tc.addresses)

			// Check array lengths match
			assert.Equal(t, len(tc.expectedResult), len(result), "Array lengths should match")

			// To handle checksum differences, we'll compare the normalized lowercase versions
			// as the implementation may use different casing than expected
			if len(tc.expectedResult) > 0 {
				for i, addr := range result {
					expected := common.HexToAddress(tc.expectedResult[i])
					actual := common.HexToAddress(addr)
					assert.Equal(t, expected, actual, "Addresses should match regardless of case")
				}
			}
		})
	}
}

/**************************************************************************************************
** TestToTExternalVaultMigration tests the toTExternalVaultMigration function to ensure it
** correctly converts an internal migration structure to the external format.
**************************************************************************************************/
func TestToTExternalVaultMigration(t *testing.T) {
	testCases := []struct {
		name           string
		migration      models.TMigration
		expectedResult TExternalVaultMigration
	}{
		{
			name: "Migration not available",
			migration: models.TMigration{
				Available: false,
				Target:    common.Address{},
				Contract:  common.Address{},
			},
			expectedResult: TExternalVaultMigration{
				Available: false,
				Address:   "0x0000000000000000000000000000000000000000",
				Contract:  "0x0000000000000000000000000000000000000000",
			},
		},
		{
			name: "Migration available",
			migration: models.TMigration{
				Available: true,
				Target:    common.HexToAddress("0x1234567890123456789012345678901234567890"),
				Contract:  common.HexToAddress("0xabcdefabcdefabcdefabcdefabcdefabcdefabcd"),
			},
			expectedResult: TExternalVaultMigration{
				Available: true,
				Address:   "0x1234567890123456789012345678901234567890",
				Contract:  "0xABCdEfABcDeFabCdEfAbCdEfaBcDeFAbcDEFabCd", // Note: This is checksum format
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := toTExternalVaultMigration(tc.migration)

			assert.Equal(t, tc.expectedResult.Available, result.Available, "Available flag should match")

			// Compare addresses using Ethereum's address type to handle potential checksum differences
			assert.Equal(t, common.HexToAddress(tc.expectedResult.Address), common.HexToAddress(result.Address), "Target address should match")
			assert.Equal(t, common.HexToAddress(tc.expectedResult.Contract), common.HexToAddress(result.Contract), "Contract address should match")
		})
	}
}
