package env

import (
	"os"
	"testing"
)

/**************************************************************************************************
** TestEnvironmentInitialization tests that the environment package initializes correctly with
** all the expected chains and their configurations. This test validates:
** - All expected chains are loaded into the CHAINS map
** - SUPPORTED_CHAIN_IDS contains the correct chain IDs
** - Core chains like Ethereum, Optimism, and Arbitrum are properly configured
**
** This ensures that the basic chain initialization works correctly, which is fundamental
** for the entire application.
**************************************************************************************************/
func TestEnvironmentInitialization(t *testing.T) {
	// Verify that we have the expected number of chains
	if len(CHAINS) == 0 {
		t.Error("CHAINS map is empty, initialization may have failed")
	}

	// Verify that we have the expected chains
	expectedChains := []uint64{1, 10, 100, 137, 250, 8453, 42161}
	for _, expectedChainID := range expectedChains {
		if _, exists := CHAINS[expectedChainID]; !exists {
			t.Errorf("Expected chain ID %d is missing from CHAINS map", expectedChainID)
		}
	}

	// Verify SUPPORTED_CHAIN_IDS matches CHAINS keys
	if len(SUPPORTED_CHAIN_IDS) != len(CHAINS) {
		t.Errorf("SUPPORTED_CHAIN_IDS length (%d) doesn't match CHAINS length (%d)",
			len(SUPPORTED_CHAIN_IDS), len(CHAINS))
	}

	// Check for specific chains
	ethereum, exists := CHAINS[1]
	if !exists {
		t.Error("Ethereum chain (ID 1) not found in CHAINS map")
	} else if ethereum.ID != 1 {
		t.Errorf("Ethereum chain ID mismatch, got %d, expected 1", ethereum.ID)
	}

	optimism, exists := CHAINS[10]
	if !exists {
		t.Error("Optimism chain (ID 10) not found in CHAINS map")
	} else if optimism.ID != 10 {
		t.Errorf("Optimism chain ID mismatch, got %d, expected 10", optimism.ID)
	}

	arbitrum, exists := CHAINS[42161]
	if !exists {
		t.Error("Arbitrum chain (ID 42161) not found in CHAINS map")
	} else if arbitrum.ID != 42161 {
		t.Errorf("Arbitrum chain ID mismatch, got %d, expected 42161", arbitrum.ID)
	}
}

/**************************************************************************************************
** TestGetChains tests the GetChains function to ensure it returns the complete CHAINS map.
** This test validates:
** - The function returns a map with the same length as the global CHAINS map
** - The returned map contains the same chains as the global CHAINS map
**
** This ensures that the GetChains accessor function correctly provides access to chain data.
**************************************************************************************************/
func TestGetChains(t *testing.T) {
	chains := GetChains()

	// Verify the chains map has the expected length
	if len(chains) != len(CHAINS) {
		t.Errorf("GetChains returned %d chains, expected %d", len(chains), len(CHAINS))
	}

	// Verify the chains map contains the expected chains
	for chainID, chain := range CHAINS {
		if returnedChain, exists := chains[chainID]; !exists {
			t.Errorf("Chain ID %d missing from GetChains result", chainID)
		} else if returnedChain.ID != chain.ID {
			t.Errorf("Chain ID mismatch for chain %d, got %d", chainID, returnedChain.ID)
		}
	}
}

/**************************************************************************************************
** TestGetChain tests the GetChain function for both valid and invalid chain IDs.
** This test validates:
** - The function returns the correct chain for a valid chain ID
** - The function indicates non-existence for an invalid chain ID
**
** This ensures that the GetChain function correctly handles both normal and edge cases.
**************************************************************************************************/
func TestGetChain(t *testing.T) {
	// Test with valid chain ID (Ethereum)
	ethereum, exists := GetChain(1)
	if !exists {
		t.Error("GetChain reported Ethereum (ID 1) as non-existent")
	} else if ethereum.ID != 1 {
		t.Errorf("GetChain returned chain with ID %d, expected 1", ethereum.ID)
	}

	// Test with invalid chain ID
	_, exists = GetChain(999999)
	if exists {
		t.Error("GetChain reported non-existent chain ID 999999 as existent")
	}
}

/**************************************************************************************************
** TestSetEnvOverrides tests that environment variables can override default chain configurations.
** This test validates:
** - Setting an RPC URI environment variable updates the chain's RPC URI
** - Setting CoinGecko API keys environment variable updates the global CG_DEMO_KEYS slice
**
** This ensures that the SetEnv function correctly applies environment variable overrides.
**************************************************************************************************/
func TestSetEnvOverrides(t *testing.T) {
	// Save original values to restore later
	originalEthereumRPC := CHAINS[1].RpcURI
	originalCGKeys := make([]string, len(CG_DEMO_KEYS))
	copy(originalCGKeys, CG_DEMO_KEYS)

	// Set test environment variables
	os.Setenv("RPC_URI_FOR_1", "https://test-ethereum-rpc.com")
	os.Setenv("CG_DEMO_KEYS", "testkey1,testkey2,testkey3")

	// Clear current values
	CG_DEMO_KEYS = []string{}

	// Call SetEnv to apply the environment variables
	SetEnv()

	// Test RPC URI override
	if CHAINS[1].RpcURI != "https://test-ethereum-rpc.com" {
		t.Errorf("RPC URI override failed, got %s, expected https://test-ethereum-rpc.com",
			CHAINS[1].RpcURI)
	}

	// Test CoinGecko keys override
	expectedKeys := []string{"testkey1", "testkey2", "testkey3"}
	if len(CG_DEMO_KEYS) != len(expectedKeys) {
		t.Errorf("CG_DEMO_KEYS length mismatch, got %d, expected %d",
			len(CG_DEMO_KEYS), len(expectedKeys))
	} else {
		for i, key := range expectedKeys {
			if i < len(CG_DEMO_KEYS) && CG_DEMO_KEYS[i] != key {
				t.Errorf("CG_DEMO_KEYS[%d] = %s, expected %s", i, CG_DEMO_KEYS[i], key)
			}
		}
	}

	// Restore original values
	chain := CHAINS[1]
	chain.RpcURI = originalEthereumRPC
	CHAINS[1] = chain
	CG_DEMO_KEYS = originalCGKeys

	// Clean up environment variables
	os.Unsetenv("RPC_URI_FOR_1")
	os.Unsetenv("CG_DEMO_KEYS")
}
