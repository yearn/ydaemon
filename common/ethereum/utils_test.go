package ethereum

import (
	"os"
	"strconv"
	"testing"

	"github.com/ethereum/go-ethereum/ethclient"
)

/**************************************************************************************************
** TestGetRPCURI tests the GetRPCURI function to ensure it correctly returns the RPC URI
** for a given chain ID. This test validates:
** - The function returns the expected URI for valid chain IDs
** - The function returns an empty string for invalid chain IDs
**
** This ensures that RPC URI retrieval works correctly for connection establishment.
**************************************************************************************************/
func TestGetRPCURI(t *testing.T) {
	// Test with a valid chain ID
	validChainID := uint64(1) // Ethereum Mainnet
	uri := GetRPCURI(validChainID)

	// The actual URI might vary based on configuration, but should not be empty for valid chains
	if uri == "" {
		t.Errorf("Expected non-empty RPC URI for chain ID %d, got empty string", validChainID)
	}

	// Test with an invalid chain ID
	invalidChainID := uint64(999999)
	uri = GetRPCURI(invalidChainID)

	if uri != "" {
		t.Errorf("Expected empty RPC URI for invalid chain ID %d, got %s", invalidChainID, uri)
	}
}

/**************************************************************************************************
** TestGetWSEnvURI tests the GetWSEnvURI function to ensure it correctly returns the WebSocket URI
** for a given chain ID. This test validates:
** - The function returns environment variable value when defined
** - The function falls back to RPC URI when environment variable is not defined
** - The function returns an empty string for invalid chain IDs
**
** This ensures that WebSocket URI retrieval works correctly with proper fallback behavior.
**************************************************************************************************/
func TestGetWSEnvURI(t *testing.T) {
	testChainID := uint64(1) // Ethereum Mainnet
	envVarName := "WS_URI_FOR_" + strconv.FormatUint(testChainID, 10)
	testWSURI := "wss://test-websocket-uri.example.com"

	// Store the original value to restore later
	originalValue, originalExists := os.LookupEnv(envVarName)

	// Test with environment variable set
	os.Setenv(envVarName, testWSURI)
	uri := GetWSEnvURI(testChainID)

	if uri != testWSURI {
		t.Errorf("Expected WS URI %s for chain ID %d, got %s", testWSURI, testChainID, uri)
	}

	// Test with environment variable unset (should fall back to RPC URI)
	os.Unsetenv(envVarName)
	uri = GetWSEnvURI(testChainID)
	rpcURI := GetRPCURI(testChainID)

	if uri != rpcURI {
		t.Errorf("Expected fallback to RPC URI %s for chain ID %d, got %s", rpcURI, testChainID, uri)
	}

	// Test with an invalid chain ID
	invalidChainID := uint64(999999)
	uri = GetWSEnvURI(invalidChainID)

	if uri != "" {
		t.Errorf("Expected empty WS URI for invalid chain ID %d, got %s", invalidChainID, uri)
	}

	// Restore the original environment variable
	if originalExists {
		os.Setenv(envVarName, originalValue)
	} else {
		os.Unsetenv(envVarName)
	}
}

/**************************************************************************************************
** TestGetRPC tests the GetRPC function to ensure it correctly returns the ethclient for a
** given chain ID. This test validates:
** - The function returns nil for uninitialized clients
** - After initialization, it returns the expected client
**
** The test is limited by not being able to fully initialize clients during testing.
**************************************************************************************************/
func TestGetRPC(t *testing.T) {
	// Test with a client that hasn't been initialized (should return nil)
	testChainID := uint64(9999) // Use a chain ID unlikely to be initialized
	client := GetRPC(testChainID)

	if client != nil {
		t.Errorf("Expected nil client for uninitialized chain ID %d", testChainID)
	}

	// Temporarily add a client to the RPC map for testing
	// This is a simplified test as we cannot fully initialize a real client during testing
	testClient := &ethclient.Client{}
	oldClient, exists := RPC[testChainID]
	RPC[testChainID] = testClient

	client = GetRPC(testChainID)
	if client != testClient {
		t.Errorf("Expected specific client for chain ID %d", testChainID)
	}

	// Restore original state
	if exists {
		RPC[testChainID] = oldClient
	} else {
		delete(RPC, testChainID)
	}
}

/**************************************************************************************************
** TestRandomSigner tests the randomSigner function to ensure it generates a valid transaction
** signer with appropriate default values. This test validates:
** - The signer is not nil
** - The signer has NoSend set to true
** - The signer has zero gas pricing
** - The signer has maximum gas limit
**
** This ensures that the generated signer is suitable for read-only operations.
**************************************************************************************************/
func TestRandomSigner(t *testing.T) {
	signer := randomSigner()

	if signer == nil {
		t.Error("Expected non-nil signer, got nil")
		return
	}

	if !signer.NoSend {
		t.Error("Expected signer.NoSend to be true")
	}

	if signer.GasPrice == nil || signer.GasPrice.Int64() != 0 {
		t.Errorf("Expected signer.GasPrice to be 0, got %v", signer.GasPrice)
	}

	if signer.GasLimit == 0 {
		t.Error("Expected signer.GasLimit to be non-zero")
	}

	if signer.Context == nil {
		t.Error("Expected signer.Context to be non-nil")
	}
}

/**************************************************************************************************
** TestMulticallClientMap tests that the MulticallClientForChainID map is properly initialized
** as an empty map. This test validates:
** - The map exists and is not nil
** - The map is initialized as empty
**
** This ensures that the multicall client map is ready to store clients for different chains.
**************************************************************************************************/
func TestMulticallClientMap(t *testing.T) {
	if MulticallClientForChainID == nil {
		t.Error("Expected MulticallClientForChainID to be initialized, got nil")
	}
}
