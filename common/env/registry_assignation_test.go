package env

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
)

/**************************************************************************************************
** TestIsRegistryFromYearnCore tests the IsRegistryFromYearnCore function with various inputs.
** This test validates:
** - The function correctly identifies Yearn Core registries
** - The function correctly rejects non-Yearn Core registries
** - The function handles unsupported chain IDs gracefully
**
** Yearn Core registries are the standard Yearn vaults that should be displayed on the Yearn.fi
** website, so correctly identifying them is crucial.
**************************************************************************************************/
func TestIsRegistryFromYearnCore(t *testing.T) {
	// Create a test setup with a temporary registry entry
	originalEthereumRegistries := make([]TContractData, len(CHAINS[1].Registries))
	copy(originalEthereumRegistries, CHAINS[1].Registries)

	// Add test registries to Ethereum chain
	testYearnRegistry := TContractData{
		Address: common.HexToAddress("0x1234567890123456789012345678901234567890"),
		Block:   15000000,
		Label:   "YEARN", // This is what the function checks for
	}

	testNonYearnRegistry := TContractData{
		Address: common.HexToAddress("0x0987654321098765432109876543210987654321"),
		Block:   15000000,
		Label:   "NON_YEARN",
	}

	// Add test registries to the chain configuration
	chainCopy := CHAINS[1]
	chainCopy.Registries = append([]TContractData{testYearnRegistry, testNonYearnRegistry}, chainCopy.Registries...)
	CHAINS[1] = chainCopy

	// Test cases
	testCases := []struct {
		name           string
		chainID        uint64
		registryAddr   common.Address
		expectedResult bool
	}{
		{
			name:           "Yearn Core registry",
			chainID:        1, // Ethereum
			registryAddr:   testYearnRegistry.Address,
			expectedResult: true,
		},
		{
			name:           "Non-Yearn Core registry",
			chainID:        1, // Ethereum
			registryAddr:   testNonYearnRegistry.Address,
			expectedResult: false,
		},
		{
			name:           "Non-existent registry",
			chainID:        1, // Ethereum
			registryAddr:   common.HexToAddress("0xdeadbeefdeadbeefdeadbeefdeadbeefdeadbeef"),
			expectedResult: false,
		},
		{
			name:           "Unsupported chain",
			chainID:        999999, // Non-existent chain
			registryAddr:   testYearnRegistry.Address,
			expectedResult: false,
		},
	}

	// Run test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := IsRegistryFromYearnCore(tc.chainID, tc.registryAddr)
			if result != tc.expectedResult {
				t.Errorf("IsRegistryFromYearnCore(%d, %s) = %v, expected %v",
					tc.chainID, tc.registryAddr.Hex(), result, tc.expectedResult)
			}
		})
	}

	// Restore original registries
	chainCopy = CHAINS[1]
	chainCopy.Registries = originalEthereumRegistries
	CHAINS[1] = chainCopy
}

/**************************************************************************************************
** TestIsRegistryFromJuiced tests the IsRegistryFromJuiced function with various inputs.
** This test validates:
** - The function correctly identifies Juiced registries
** - The function correctly rejects non-Juiced registries
** - The function handles unsupported chain IDs gracefully
**
** Juiced registries are for specialized vaults that should be displayed on the Juiced.app
** website, so correctly identifying them is important for proper categorization.
**************************************************************************************************/
func TestIsRegistryFromJuiced(t *testing.T) {
	// Create a test setup with a temporary registry entry
	originalEthereumRegistries := make([]TContractData, len(CHAINS[1].Registries))
	copy(originalEthereumRegistries, CHAINS[1].Registries)

	// Add test registries to Ethereum chain
	testJuicedRegistry := TContractData{
		Address: common.HexToAddress("0x1234567890123456789012345678901234567891"),
		Block:   15000000,
		Label:   "JUICED", // This is what the function checks for
	}

	testNonJuicedRegistry := TContractData{
		Address: common.HexToAddress("0x0987654321098765432109876543210987654322"),
		Block:   15000000,
		Label:   "NON_JUICED",
	}

	// Add test registries to the chain configuration
	chainCopy := CHAINS[1]
	chainCopy.Registries = append([]TContractData{testJuicedRegistry, testNonJuicedRegistry}, chainCopy.Registries...)
	CHAINS[1] = chainCopy

	// Test cases
	testCases := []struct {
		name           string
		chainID        uint64
		registryAddr   common.Address
		expectedResult bool
	}{
		{
			name:           "Juiced registry",
			chainID:        1, // Ethereum
			registryAddr:   testJuicedRegistry.Address,
			expectedResult: true,
		},
		{
			name:           "Non-Juiced registry",
			chainID:        1, // Ethereum
			registryAddr:   testNonJuicedRegistry.Address,
			expectedResult: false,
		},
		{
			name:           "Non-existent registry",
			chainID:        1, // Ethereum
			registryAddr:   common.HexToAddress("0xdeadbeefdeadbeefdeadbeefdeadbeefdeadbeef"),
			expectedResult: false,
		},
		{
			name:           "Unsupported chain",
			chainID:        999999, // Non-existent chain
			registryAddr:   testJuicedRegistry.Address,
			expectedResult: false,
		},
	}

	// Run test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := IsRegistryFromJuiced(tc.chainID, tc.registryAddr)
			if result != tc.expectedResult {
				t.Errorf("IsRegistryFromJuiced(%d, %s) = %v, expected %v",
					tc.chainID, tc.registryAddr.Hex(), result, tc.expectedResult)
			}
		})
	}

	// Restore original registries
	chainCopy = CHAINS[1]
	chainCopy.Registries = originalEthereumRegistries
	CHAINS[1] = chainCopy
}

/**************************************************************************************************
** TestIsRegistryFromPublicERC4626 tests the IsRegistryFromPublicERC4626 function with various inputs.
** This test validates:
** - The function correctly identifies PUBLIC_ERC4626 registries
** - The function correctly rejects non-PUBLIC_ERC4626 registries
** - The function handles unsupported chain IDs gracefully
**
** PUBLIC_ERC4626 registries represent standard ERC4626 vaults that should not be displayed
** by default on the Yearn.fi website.
**************************************************************************************************/
func TestIsRegistryFromPublicERC4626(t *testing.T) {
	// Create a test setup with a temporary registry entry
	originalEthereumRegistries := make([]TContractData, len(CHAINS[1].Registries))
	copy(originalEthereumRegistries, CHAINS[1].Registries)

	// Add test registries to Ethereum chain
	testPublicERC4626Registry := TContractData{
		Address: common.HexToAddress("0x1234567890123456789012345678901234567892"),
		Block:   15000000,
		Label:   "PUBLIC_ERC4626", // This is what the function checks for
	}

	testNonPublicERC4626Registry := TContractData{
		Address: common.HexToAddress("0x0987654321098765432109876543210987654323"),
		Block:   15000000,
		Label:   "NON_PUBLIC_ERC4626",
	}

	// Add test registries to the chain configuration
	chainCopy := CHAINS[1]
	chainCopy.Registries = append([]TContractData{testPublicERC4626Registry, testNonPublicERC4626Registry}, chainCopy.Registries...)
	CHAINS[1] = chainCopy

	// Test cases
	testCases := []struct {
		name           string
		chainID        uint64
		registryAddr   common.Address
		expectedResult bool
	}{
		{
			name:           "PUBLIC_ERC4626 registry",
			chainID:        1, // Ethereum
			registryAddr:   testPublicERC4626Registry.Address,
			expectedResult: true,
		},
		{
			name:           "Non-PUBLIC_ERC4626 registry",
			chainID:        1, // Ethereum
			registryAddr:   testNonPublicERC4626Registry.Address,
			expectedResult: false,
		},
		{
			name:           "Non-existent registry",
			chainID:        1, // Ethereum
			registryAddr:   common.HexToAddress("0xdeadbeefdeadbeefdeadbeefdeadbeefdeadbeef"),
			expectedResult: false,
		},
		{
			name:           "Unsupported chain",
			chainID:        999999, // Non-existent chain
			registryAddr:   testPublicERC4626Registry.Address,
			expectedResult: false,
		},
	}

	// Run test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := IsRegistryFromPublicERC4626(tc.chainID, tc.registryAddr)
			if result != tc.expectedResult {
				t.Errorf("IsRegistryFromPublicERC4626(%d, %s) = %v, expected %v",
					tc.chainID, tc.registryAddr.Hex(), result, tc.expectedResult)
			}
		})
	}

	// Restore original registries
	chainCopy = CHAINS[1]
	chainCopy.Registries = originalEthereumRegistries
	CHAINS[1] = chainCopy
}

/**************************************************************************************************
** TestIsRegistryFromPoolTogether tests the IsRegistryFromPoolTogether function with various inputs.
** This test validates:
** - The function correctly identifies PoolTogether registries on Optimism (chain ID 10)
** - The function correctly rejects non-PoolTogether registries
** - The function correctly rejects registries on chains other than Optimism
**
** PoolTogether registries are for specialized Yearn X PoolTogether vaults that should not be
** displayed by default on the Yearn.fi website.
**************************************************************************************************/
func TestIsRegistryFromPoolTogether(t *testing.T) {
	// The specific address in the function: 0x0c379e9b71ba7079084ada0d1c1aeb85d24dfd39
	poolTogetherRegistryAddr := common.HexToAddress("0x0c379e9b71ba7079084ada0d1c1aeb85d24dfd39")

	// Test cases
	testCases := []struct {
		name           string
		chainID        uint64
		registryAddr   common.Address
		expectedResult bool
	}{
		{
			name:           "PoolTogether registry on Optimism",
			chainID:        10, // Optimism
			registryAddr:   poolTogetherRegistryAddr,
			expectedResult: true,
		},
		{
			name:           "Different address on Optimism",
			chainID:        10, // Optimism
			registryAddr:   common.HexToAddress("0x1234567890123456789012345678901234567890"),
			expectedResult: false,
		},
		{
			name:           "PoolTogether address on Ethereum",
			chainID:        1, // Ethereum
			registryAddr:   poolTogetherRegistryAddr,
			expectedResult: false, // Function specifically checks for chain ID 10
		},
		{
			name:           "PoolTogether address on unsupported chain",
			chainID:        999999, // Non-existent chain
			registryAddr:   poolTogetherRegistryAddr,
			expectedResult: false,
		},
	}

	// Run test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := IsRegistryFromPoolTogether(tc.chainID, tc.registryAddr)
			if result != tc.expectedResult {
				t.Errorf("IsRegistryFromPoolTogether(%d, %s) = %v, expected %v",
					tc.chainID, tc.registryAddr.Hex(), result, tc.expectedResult)
			}
		})
	}
}

/**************************************************************************************************
** TestIsRegistryDisabled tests the IsRegistryDisabled function with various inputs.
** This test validates:
** - The function correctly identifies disabled registries (with Tag "DISABLED")
** - The function correctly identifies non-disabled registries
** - The function handles unsupported chain IDs gracefully
**
** Identifying disabled registries is important to prevent displaying deprecated or
** non-functioning vaults to users.
**************************************************************************************************/
func TestIsRegistryDisabled(t *testing.T) {
	// Create a test setup with a temporary registry entry
	originalEthereumRegistries := make([]TContractData, len(CHAINS[1].Registries))
	copy(originalEthereumRegistries, CHAINS[1].Registries)

	// Add test registries to Ethereum chain
	testDisabledRegistry := TContractData{
		Address: common.HexToAddress("0x1234567890123456789012345678901234567893"),
		Block:   15000000,
		Label:   "ANY_LABEL",
		Tag:     "DISABLED", // This is what the function checks for
	}

	testEnabledRegistry := TContractData{
		Address: common.HexToAddress("0x0987654321098765432109876543210987654324"),
		Block:   15000000,
		Label:   "ANY_LABEL",
		Tag:     "ENABLED",
	}

	// Add test registries to the chain configuration
	chainCopy := CHAINS[1]
	chainCopy.Registries = append([]TContractData{testDisabledRegistry, testEnabledRegistry}, chainCopy.Registries...)
	CHAINS[1] = chainCopy

	// Test cases
	testCases := []struct {
		name           string
		chainID        uint64
		registryAddr   common.Address
		expectedResult bool
	}{
		{
			name:           "Disabled registry",
			chainID:        1, // Ethereum
			registryAddr:   testDisabledRegistry.Address,
			expectedResult: true,
		},
		{
			name:           "Enabled registry",
			chainID:        1, // Ethereum
			registryAddr:   testEnabledRegistry.Address,
			expectedResult: false,
		},
		{
			name:           "Non-existent registry",
			chainID:        1, // Ethereum
			registryAddr:   common.HexToAddress("0xdeadbeefdeadbeefdeadbeefdeadbeefdeadbeef"),
			expectedResult: false,
		},
		{
			name:           "Unsupported chain",
			chainID:        999999, // Non-existent chain
			registryAddr:   testDisabledRegistry.Address,
			expectedResult: false,
		},
	}

	// Run test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := IsRegistryDisabled(tc.chainID, tc.registryAddr)
			if result != tc.expectedResult {
				t.Errorf("IsRegistryDisabled(%d, %s) = %v, expected %v",
					tc.chainID, tc.registryAddr.Hex(), result, tc.expectedResult)
			}
		})
	}

	// Restore original registries
	chainCopy = CHAINS[1]
	chainCopy.Registries = originalEthereumRegistries
	CHAINS[1] = chainCopy
}
