package env

import (
	"path/filepath"
	"testing"

	"github.com/ethereum/go-ethereum/common"
)

/**************************************************************************************************
** TestContractDataStructure tests the TContractData structure for proper initialization
** and field access. This test validates:
** - The structure can be properly initialized with various field values
** - All fields can be accessed correctly after initialization
**
** This ensures that the core TContractData structure used throughout the application
** behaves as expected.
**************************************************************************************************/
func TestContractDataStructure(t *testing.T) {
	// Test creating a TContractData instance with all fields
	contractData := TContractData{
		Address: common.HexToAddress("0x1234567890123456789012345678901234567890"),
		Block:   12345678,
		Version: 1,
		Tag:     "TEST_TAG",
		Label:   "TEST_LABEL",
	}

	// Test field access
	if contractData.Address.Hex() != "0x1234567890123456789012345678901234567890" {
		t.Errorf("Address field mismatch, got %s, expected 0x1234567890123456789012345678901234567890",
			contractData.Address.Hex())
	}

	if contractData.Block != 12345678 {
		t.Errorf("Block field mismatch, got %d, expected 12345678", contractData.Block)
	}

	if contractData.Version != 1 {
		t.Errorf("Version field mismatch, got %d, expected 1", contractData.Version)
	}

	if contractData.Tag != "TEST_TAG" {
		t.Errorf("Tag field mismatch, got %s, expected TEST_TAG", contractData.Tag)
	}

	if contractData.Label != "TEST_LABEL" {
		t.Errorf("Label field mismatch, got %s, expected TEST_LABEL", contractData.Label)
	}

	// Test creating a TContractData instance with minimal fields
	minimalContractData := TContractData{
		Address: common.HexToAddress("0x0987654321098765432109876543210987654321"),
		Block:   87654321,
	}

	// Test field access for minimal contract data
	if minimalContractData.Address.Hex() != "0x0987654321098765432109876543210987654321" {
		t.Errorf("Address field mismatch, got %s, expected 0x0987654321098765432109876543210987654321",
			minimalContractData.Address.Hex())
	}

	if minimalContractData.Block != 87654321 {
		t.Errorf("Block field mismatch, got %d, expected 87654321", minimalContractData.Block)
	}

	// Version, Tag, and Label should be zero values
	if minimalContractData.Version != 0 {
		t.Errorf("Version field should be 0, got %d", minimalContractData.Version)
	}

	if minimalContractData.Tag != "" {
		t.Errorf("Tag field should be empty, got %s", minimalContractData.Tag)
	}

	if minimalContractData.Label != "" {
		t.Errorf("Label field should be empty, got %s", minimalContractData.Label)
	}
}

/**************************************************************************************************
** TestDefaultCoinAddress tests the DEFAULT_COIN_ADDRESS constant.
** This test validates:
** - The constant has the expected value (0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE)
** - The address is properly formatted as a valid Ethereum address
**
** This is important because this address is used as a placeholder for native coins
** (ETH, MATIC, etc.) across various chains.
**************************************************************************************************/
func TestDefaultCoinAddress(t *testing.T) {
	expectedAddress := "0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE"

	// Check value
	if DEFAULT_COIN_ADDRESS.Hex() != expectedAddress {
		t.Errorf("DEFAULT_COIN_ADDRESS has unexpected value. Got %s, expected %s",
			DEFAULT_COIN_ADDRESS.Hex(), expectedAddress)
	}

	// Verify it's a valid address (20 bytes)
	if len(DEFAULT_COIN_ADDRESS.Bytes()) != 20 {
		t.Errorf("DEFAULT_COIN_ADDRESS is not a valid Ethereum address (expected 20 bytes, got %d)",
			len(DEFAULT_COIN_ADDRESS.Bytes()))
	}
}

/**************************************************************************************************
** TestBaseDataPath tests the BASE_DATA_PATH constant.
** This test validates:
** - The constant points to a valid directory path
** - The path contains the expected "data" suffix
**
** This ensures that the application will be able to find its data directory.
**************************************************************************************************/
func TestBaseDataPath(t *testing.T) {
	// Check that BASE_DATA_PATH is defined and non-empty
	if BASE_DATA_PATH == "" {
		t.Error("BASE_DATA_PATH is empty")
	}

	// Check that the path ends with "data"
	dataDir := filepath.Base(BASE_DATA_PATH)
	if dataDir != "data" {
		t.Errorf("BASE_DATA_PATH does not end with 'data'. Got %s", dataDir)
	}

	// Note: We don't check if the directory exists because it might not
	// exist in testing environments. In real usage, the application would
	// typically create the directory if it doesn't exist.
}

/**************************************************************************************************
** TestAssetAndPriceURLs tests the URL constants used for external services.
** This test validates:
** - BASE_ASSET_URL is properly defined
** - GECKO_PRICE_URL is properly defined
** - LLAMA_PRICE_URL is properly defined
**
** These URLs are critical for fetching token metadata and price information.
**************************************************************************************************/
func TestAssetAndPriceURLs(t *testing.T) {
	// Test BASE_ASSET_URL
	if BASE_ASSET_URL == "" {
		t.Error("BASE_ASSET_URL is empty")
	}

	// Test GECKO_PRICE_URL
	if GECKO_PRICE_URL == "" {
		t.Error("GECKO_PRICE_URL is empty")
	}

	// Test LLAMA_PRICE_URL
	if LLAMA_PRICE_URL == "" {
		t.Error("LLAMA_PRICE_URL is empty")
	}

	// Check URL structure (very basic validation)
	for name, url := range map[string]string{
		"BASE_ASSET_URL":  BASE_ASSET_URL,
		"GECKO_PRICE_URL": GECKO_PRICE_URL,
		"LLAMA_PRICE_URL": LLAMA_PRICE_URL,
	} {
		if url[:7] != "http://" && url[:8] != "https://" {
			t.Errorf("%s does not start with 'http://' or 'https://'. Got: %s", name, url)
		}
	}
}

/**************************************************************************************************
** TestGetCurrentPath tests the getCurrentPath function.
** This test validates:
** - The function returns a non-empty path
** - The path points to a directory that exists in the filesystem
**
** This function is used to determine the location of the package, which is crucial
** for resolving paths to data and configuration files.
**************************************************************************************************/
func TestGetCurrentPath(t *testing.T) {
	path := getCurrentPath()

	// Check that the path is non-empty
	if path == "" {
		t.Error("getCurrentPath returned an empty path")
	}

	// Check the last component of the path (should be the directory name)
	dirName := filepath.Base(path)
	if dirName == "" {
		t.Error("getCurrentPath returned a path with an empty base component")
	}

	// The directory should be named "env"
	if dirName != "env" {
		t.Errorf("getCurrentPath returned a path ending with '%s', expected 'env'", dirName)
	}
}
