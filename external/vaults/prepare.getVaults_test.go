package vaults

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yearn/ydaemon/internal/models"
)

/**************************************************************************************************
** TestIsV3VaultDetection tests the isV3Vault function to verify it correctly identifies V3 vaults
** based on version and kind.
**************************************************************************************************/
func TestIsV3VaultDetection(t *testing.T) {
	testCases := []struct {
		name         string
		vault        models.TVault
		expectedIsV3 bool
	}{
		{
			name: "V3 vault by version",
			vault: models.TVault{
				Version: "3.0.0",
				Kind:    "", // Not a V3-specific kind
			},
			expectedIsV3: true,
		},
		{
			name: "V3 vault by kind - Multiple",
			vault: models.TVault{
				Version: "2.0.0",                  // Version says V2
				Kind:    models.VaultKindMultiple, // But kind is V3-specific
			},
			expectedIsV3: true,
		},
		{
			name: "V3 vault by kind - Single",
			vault: models.TVault{
				Version: "2.0.0",                // Version says V2
				Kind:    models.VaultKindSingle, // But kind is V3-specific
			},
			expectedIsV3: true,
		},
		{
			name: "Not V3 vault",
			vault: models.TVault{
				Version: "2.0.0",
				Kind:    "",
			},
			expectedIsV3: false,
		},
		{
			name: "V3 subversion",
			vault: models.TVault{
				Version: "3.1.2",
				Kind:    "",
			},
			expectedIsV3: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := isV3Vault(tc.vault)
			assert.Equal(t, tc.expectedIsV3, result, "isV3Vault result should match expected value")
		})
	}
}

/**************************************************************************************************
** TestIsV2VaultDetection tests the isV2Vault function to verify it correctly identifies V2 vaults
** based on version.
**************************************************************************************************/
func TestIsV2VaultDetection(t *testing.T) {
	testCases := []struct {
		name         string
		vault        models.TVault
		expectedIsV2 bool
	}{
		{
			name: "V2 vault by version",
			vault: models.TVault{
				Version: "2.0.0",
				Kind:    "",
			},
			expectedIsV2: true,
		},
		{
			name: "V2 subversion",
			vault: models.TVault{
				Version: "2.4.3",
				Kind:    "",
			},
			expectedIsV2: true,
		},
		{
			name: "Not V2 vault (V3)",
			vault: models.TVault{
				Version: "3.0.0",
				Kind:    "",
			},
			expectedIsV2: false,
		},
		{
			name: "V1 vault treated as V2",
			vault: models.TVault{
				Version: "1.0.0",
				Kind:    "",
			},
			expectedIsV2: false, // V1 should NOT be treated as V2 for this function
		},
		{
			name: "Empty version treated as V2",
			vault: models.TVault{
				Version: "",
				Kind:    "",
			},
			expectedIsV2: false, // Empty version should NOT be treated as V2 for this function
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := isV2Vault(tc.vault)
			assert.Equal(t, tc.expectedIsV2, result, "isV2Vault result should match expected value")
		})
	}
}
