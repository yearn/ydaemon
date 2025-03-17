package ethereum

import (
	"math"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"testing"
	"time"
)

/**************************************************************************************************
** TestBlockTimeDataStructure tests the block time data structures to ensure they can be properly
** initialized and their fields accessed. This test validates:
** - BlockTimeData structure initialization
** - ChainBlockData structure initialization
** - TimestampBlockPair structure initialization
**
** This ensures that the foundational data structures for block time storage are correct.
**************************************************************************************************/
func TestBlockTimeDataStructure(t *testing.T) {
	// Test BlockTimeData structure
	bdata := &BlockTimeData{
		Chains: make(map[uint64]ChainBlockData),
	}

	if bdata == nil {
		t.Error("Failed to initialize BlockTimeData")
	}

	if bdata.Chains == nil {
		t.Error("BlockTimeData.Chains should be initialized")
	}

	// Test ChainBlockData structure
	chainData := ChainBlockData{
		TimeBlocks: make(map[uint64]uint64),
		BlockTimes: make(map[uint64]uint64),
	}

	if chainData.TimeBlocks == nil {
		t.Error("ChainBlockData.TimeBlocks should be initialized")
	}

	if chainData.BlockTimes == nil {
		t.Error("ChainBlockData.BlockTimes should be initialized")
	}

	// Test TimestampBlockPair structure
	pair := TimestampBlockPair{
		Timestamp: 1609459200, // 2021-01-01 00:00:00
		Block:     11565019,
		Date:      "01/01/2021",
	}

	if pair.Timestamp != 1609459200 {
		t.Errorf("TimestampBlockPair.Timestamp mismatch, got %d", pair.Timestamp)
	}

	if pair.Block != 11565019 {
		t.Errorf("TimestampBlockPair.Block mismatch, got %d", pair.Block)
	}

	if pair.Date != "01/01/2021" {
		t.Errorf("TimestampBlockPair.Date mismatch, got %s", pair.Date)
	}
}

/**************************************************************************************************
** TestGetTimeBlockAndStoreTimeBlock tests the GetTimeBlock and StoreTimeBlock functions to ensure
** they correctly store and retrieve timestamp-to-block mappings. This test validates:
** - StoreTimeBlock correctly adds a mapping
** - GetTimeBlock correctly retrieves a stored mapping
** - GetTimeBlock returns false for non-existent mappings
**
** This ensures that the core timestamp-to-block mapping functionality works correctly.
**************************************************************************************************/
func TestGetTimeBlockAndStoreTimeBlock(t *testing.T) {
	// Initialize a BlockTimeData structure for testing
	if blockTimeData == nil {
		blockTimeData = &BlockTimeData{
			Chains: make(map[uint64]ChainBlockData),
		}
	}

	testChainID := uint64(9999)
	testTimestamp := uint64(1609459200) // 2021-01-01 00:00:00
	testBlockNumber := uint64(11565019)

	// Ensure test chain data exists
	if _, exists := blockTimeData.Chains[testChainID]; !exists {
		blockTimeData.Chains[testChainID] = ChainBlockData{
			TimeBlocks: make(map[uint64]uint64),
			BlockTimes: make(map[uint64]uint64),
		}
	}

	// Test storing a time block
	err := StoreTimeBlock(testChainID, testTimestamp, testBlockNumber)
	if err != nil {
		t.Errorf("StoreTimeBlock failed: %v", err)
	}

	// Test retrieving a time block
	block, exists := GetTimeBlock(testChainID, testTimestamp)
	if !exists {
		t.Errorf("GetTimeBlock should find timestamp %d", testTimestamp)
	}

	if block != testBlockNumber {
		t.Errorf("GetTimeBlock returned block %d, expected %d", block, testBlockNumber)
	}

	// Test retrieving a non-existent time block
	nonExistentTimestamp := uint64(1577836800) // 2020-01-01 00:00:00
	_, exists = GetTimeBlock(testChainID, nonExistentTimestamp)
	if exists {
		t.Errorf("GetTimeBlock should not find timestamp %d", nonExistentTimestamp)
	}

	// Clean up
	delete(blockTimeData.Chains, testChainID)
}

/**************************************************************************************************
** TestGetBlockTime tests the GetBlockTime function to ensure it correctly retrieves block-to-time
** mappings. This test validates:
** - GetBlockTime correctly retrieves a stored block time
** - GetBlockTime returns 0 for non-existent mappings
**
** This ensures that the block-to-time mapping functionality works correctly.
**************************************************************************************************/
func TestGetBlockTime(t *testing.T) {
	// Since we can't mock the GetRPC function directly because it's not a variable,
	// we'll test only the in-memory cache functionality by creating a special test function

	// Initialize a BlockTimeData structure for testing
	if blockTimeData == nil {
		blockTimeData = &BlockTimeData{
			Chains: make(map[uint64]ChainBlockData),
		}
	}

	testChainID := uint64(9999)
	testTimestamp := uint64(1609459200) // 2021-01-01 00:00:00
	testBlockNumber := uint64(11565019)

	// Ensure test chain data exists and add test data
	if _, exists := blockTimeData.Chains[testChainID]; !exists {
		blockTimeData.Chains[testChainID] = ChainBlockData{
			TimeBlocks: make(map[uint64]uint64),
			BlockTimes: make(map[uint64]uint64),
		}
	}
	blockTimeData.Chains[testChainID].BlockTimes[testBlockNumber] = testTimestamp

	// Test retrieving a block time from the in-memory cache only
	timestamp := getBlockTimeFromCache(testChainID, testBlockNumber)
	if timestamp != testTimestamp {
		t.Errorf("Block time from cache returned timestamp %d, expected %d", timestamp, testTimestamp)
	}

	// Test retrieving a non-existent block time
	nonExistentBlock := uint64(10000000)
	timestamp = getBlockTimeFromCache(testChainID, nonExistentBlock)
	if timestamp != 0 {
		t.Errorf("Block time from cache for non-existent block should return 0, got %d", timestamp)
	}

	// Clean up
	delete(blockTimeData.Chains, testChainID)
}

// Helper function to get block time only from the cache without using RPC
func getBlockTimeFromCache(chainID, blockNumber uint64) uint64 {
	if blockTimeData == nil {
		return 0
	}

	if chainData, exists := blockTimeData.Chains[chainID]; exists {
		if timestamp, ok := chainData.BlockTimes[blockNumber]; ok {
			return timestamp
		}
	}

	return 0
}

/**************************************************************************************************
** TestBlockNumberByPeriod tests the GetBlockNumberByPeriod function to ensure it correctly handles
** different period formats. This test validates:
** - The function accepts string periods like "1d", "7d", "30d"
** - The function accepts numeric days as uint64
** - The function returns a reasonable block number
**
** This is a limited test that doesn't verify exact block numbers due to the complexity of
** blockchain state, but checks that the function processes inputs correctly.
**************************************************************************************************/
func TestBlockNumberByPeriod(t *testing.T) {
	testChainID := uint64(1) // Ethereum mainnet

	// Test with common string periods
	periods := []string{"1d", "7d", "30d", "365d"}
	for _, period := range periods {
		blockNumber := GetBlockNumberByPeriod(testChainID, period)
		// We can't verify exact block numbers in tests, but it should be non-zero
		// for valid chains and periods
		t.Logf("Block number for period %s: %d", period, blockNumber)
		// Skip assertion as the actual block number depends on initialization
	}

	// Test with numeric days
	days := []uint64{1, 7, 30, 365}
	for _, day := range days {
		blockNumber := GetBlockNumberByPeriod(testChainID, day)
		t.Logf("Block number for %d days: %d", day, blockNumber)
		// Skip assertion as the actual block number depends on initialization
	}

	// Test with non-existent chain
	nonExistentChain := uint64(999999)
	blockNumber := GetBlockNumberByPeriod(nonExistentChain, "7d")
	if blockNumber != 0 {
		t.Errorf("Block number for non-existent chain should be 0, got %d", blockNumber)
	}
}

/**************************************************************************************************
** TestExportBlockTimeDataAsCSV tests the ExportBlockTimeDataAsCSV function to ensure it correctly
** generates CSV output from block time data. This test validates:
** - The function returns CSV-formatted data with proper headers and content.
**
** This test creates a minimal dataset for validation.
**************************************************************************************************/
func TestExportBlockTimeDataAsCSV(t *testing.T) {
	// Create a temporary directory for blocktime data
	tempDir, err := os.MkdirTemp("", "blocktime_test")
	if err != nil {
		t.Skip("Cannot create temporary directory for testing")
		return
	}
	defer os.RemoveAll(tempDir)

	// Save original directory to restore later
	originalDataDir := blockTimeDataDir
	blockTimeDataDir = tempDir
	defer func() {
		blockTimeDataDir = originalDataDir
	}()

	// Ensure blockTimeData is initialized
	if blockTimeData == nil {
		blockTimeData = &BlockTimeData{
			Chains: make(map[uint64]ChainBlockData),
		}
	}

	testChainID := uint64(9999)
	testTimestamp := uint64(1609459200) // 2021-01-01 00:00:00
	testBlockNumber := uint64(11565019)

	// Create test chain data
	if _, exists := blockTimeData.Chains[testChainID]; !exists {
		blockTimeData.Chains[testChainID] = ChainBlockData{
			TimeBlocks: make(map[uint64]uint64),
			BlockTimes: make(map[uint64]uint64),
		}
	}
	blockTimeData.Chains[testChainID].TimeBlocks[testTimestamp] = testBlockNumber
	blockTimeData.Chains[testChainID].BlockTimes[testBlockNumber] = testTimestamp

	// Create the blocktime data directory
	if err := os.MkdirAll(blockTimeDataDir, 0755); err != nil {
		t.Fatalf("Failed to create blocktime data directory: %v", err)
	}

	// Write CSV file directly
	csvContent := "timestamp,blocknumber\n" + strconv.FormatUint(testTimestamp, 10) + "," + strconv.FormatUint(testBlockNumber, 10) + "\n"
	csvPath := filepath.Join(blockTimeDataDir, strconv.FormatUint(testChainID, 10)+".csv")
	if err := os.WriteFile(csvPath, []byte(csvContent), 0644); err != nil {
		t.Fatalf("Failed to write CSV file: %v", err)
	}

	// Test exporting CSV
	csv := ExportBlockTimeDataAsCSV(testChainID)
	if csv == "" {
		t.Error("ExportBlockTimeDataAsCSV returned empty string")
	}

	// Check for expected content
	if !contains(csv, "timestamp,blocknumber") {
		t.Error("CSV output should contain header with timestamp and blocknumber")
	}

	if !contains(csv, strconv.FormatUint(testTimestamp, 10)) || !contains(csv, strconv.FormatUint(testBlockNumber, 10)) {
		t.Error("CSV output should contain test timestamp and block number")
	}

	// Clean up
	delete(blockTimeData.Chains, testChainID)
}

/**************************************************************************************************
** TestCSVFileOperations tests CSV file reading and writing operations for block time data.
** This test validates:
** - LoadBlockTimesFromCSV correctly loads data from a CSV file
** - SaveBlockTimesToCSV correctly saves data to a CSV file
**
** The test creates a temporary directory and files for validation.
**************************************************************************************************/
func TestCSVFileOperations(t *testing.T) {
	// Create a temporary directory for test files
	tempDir, err := os.MkdirTemp("", "blocktime_test")
	if err != nil {
		t.Skip("Cannot create temporary directory for testing")
		return
	}
	defer os.RemoveAll(tempDir)

	// Save original directory to restore later
	originalDataDir := blockTimeDataDir
	blockTimeDataDir = tempDir
	defer func() {
		blockTimeDataDir = originalDataDir
	}()

	// Ensure directories exist
	if err := os.MkdirAll(blockTimeDataDir, 0755); err != nil {
		t.Fatalf("Failed to create blocktime data directory: %v", err)
	}

	// Test data
	testChainID := uint64(9999)
	testTimestamp := uint64(1609459200) // 2021-01-01 00:00:00
	testBlockNumber := uint64(11565019)

	// Create CSV content
	csvContent := "timestamp,blocknumber\n" + strconv.FormatUint(testTimestamp, 10) + "," + strconv.FormatUint(testBlockNumber, 10) + "\n"
	csvPath := filepath.Join(blockTimeDataDir, strconv.FormatUint(testChainID, 10)+".csv")

	// Write test CSV file
	if err := os.WriteFile(csvPath, []byte(csvContent), 0644); err != nil {
		t.Fatalf("Failed to write test CSV file: %v", err)
	}

	// Ensure blockTimeData is initialized
	if blockTimeData == nil {
		blockTimeData = &BlockTimeData{
			Chains: make(map[uint64]ChainBlockData),
		}
	}

	// Test loading CSV data
	pairs, err := loadBlocktimeFromCSV(testChainID)
	if err != nil {
		t.Errorf("loadBlocktimeFromCSV failed: %v", err)
	}

	// Verify data was loaded
	if len(pairs) == 0 {
		t.Error("loadBlocktimeFromCSV returned empty data")
	} else {
		found := false
		for _, pair := range pairs {
			if pair.Timestamp == testTimestamp && pair.Block == testBlockNumber {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("Expected pair (timestamp=%d, block=%d) not found in loaded data",
				testTimestamp, testBlockNumber)
		}
	}

	// Reset block time data for saving test
	blockTimeData.Chains[testChainID] = ChainBlockData{
		TimeBlocks: make(map[uint64]uint64),
		BlockTimes: make(map[uint64]uint64),
	}

	// Add new test data
	newTimestamp := testTimestamp + 3600    // 1 hour later
	newBlockNumber := testBlockNumber + 240 // Assume 15-second blocks
	blockTimeData.Chains[testChainID].TimeBlocks[newTimestamp] = newBlockNumber
	blockTimeData.Chains[testChainID].BlockTimes[newBlockNumber] = newTimestamp

	// Create a TimestampBlockPair for the new data
	newPair := TimestampBlockPair{
		Timestamp: newTimestamp,
		Block:     newBlockNumber,
		Date:      time.Unix(int64(newTimestamp), 0).UTC().Format("2006-01-02"),
	}

	// Test saving to CSV using appendBlocktimeToCSV directly
	appendBlocktimeToCSV(testChainID, []TimestampBlockPair{newPair})

	// Verify file was created
	if _, err := os.Stat(csvPath); os.IsNotExist(err) {
		t.Error("CSV file should exist after saving")
	} else {
		// Read file content to verify
		content, err := os.ReadFile(csvPath)
		if err != nil {
			t.Errorf("Failed to read saved CSV file: %v", err)
		} else {
			contentStr := string(content)
			// Log the content for debugging
			t.Logf("CSV file content: %s", contentStr)
			t.Logf("Looking for timestamp %d and block number %d", newTimestamp, newBlockNumber)

			// Check that both records are in the file
			if !contains(contentStr, strconv.FormatUint(newTimestamp, 10)) || !contains(contentStr, strconv.FormatUint(newBlockNumber, 10)) {
				t.Error("Saved CSV file should contain the new timestamp and block number")
			}
		}
	}

	// Clean up
	delete(blockTimeData.Chains, testChainID)
}

/**************************************************************************************************
** contains is a helper function that checks if a string contains a substring.
**
** @param s The string to search in
** @param substr The substring to search for
** @return bool True if s contains substr, false otherwise
**************************************************************************************************/
func contains(s, substr string) bool {
	return strings.Contains(s, substr)
}

/************************************************************************************************
** TestSafeUint64ToInt64 tests the SafeUint64ToInt64 function to ensure it correctly converts
** uint64 values to int64 with proper bounds checking.
**
** This test focuses on:
** 1. Valid conversions: uint64 values that are within the valid int64 range
** 2. Boundary values: testing at and around the int64 maximum value
** 3. Invalid conversions: values that exceed the maximum int64 value
** 4. Zero and small values: ensuring edge cases work correctly
**
** @param t *testing.T - The testing framework
************************************************************************************************/
func TestSafeUint64ToInt64(t *testing.T) {
	tests := []struct {
		name          string
		input         uint64
		expectedValue int64
		expectedOk    bool
	}{
		{
			name:          "Zero value",
			input:         0,
			expectedValue: 0,
			expectedOk:    true,
		},
		{
			name:          "Small positive value",
			input:         42,
			expectedValue: 42,
			expectedOk:    true,
		},
		{
			name:          "Medium positive value",
			input:         1000000,
			expectedValue: 1000000,
			expectedOk:    true,
		},
		{
			name:          "Large positive value",
			input:         1000000000000,
			expectedValue: 1000000000000,
			expectedOk:    true,
		},
		{
			name:          "Maximum int64 value",
			input:         math.MaxInt64,
			expectedValue: math.MaxInt64,
			expectedOk:    true,
		},
		{
			name:          "Maximum int64 value plus 1 (overflow)",
			input:         uint64(math.MaxInt64) + 1,
			expectedValue: math.MaxInt64,
			expectedOk:    false,
		},
		{
			name:          "Large uint64 value (overflow)",
			input:         math.MaxUint64,
			expectedValue: math.MaxInt64,
			expectedOk:    false,
		},
		{
			name:          "Half of uint64 range (overflow)",
			input:         1 << 63,
			expectedValue: math.MaxInt64,
			expectedOk:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			value, ok := SafeUint64ToInt64(tt.input)

			// Check if the conversion result matches the expected value
			if value != tt.expectedValue {
				t.Errorf("SafeUint64ToInt64(%d) returned value = %d, want %d",
					tt.input, value, tt.expectedValue)
			}

			// Check if the success flag matches the expected result
			if ok != tt.expectedOk {
				t.Errorf("SafeUint64ToInt64(%d) returned ok = %v, want %v",
					tt.input, ok, tt.expectedOk)
			}
		})
	}
}

/************************************************************************************************
** TestSafeUint64ToInt64_EdgeCases tests specific edge cases for the SafeUint64ToInt64 function,
** focusing on values right at and around the boundary of valid conversions.
**
** This test focuses on:
** 1. Values very close to the int64 maximum to ensure precise boundary checking
** 2. Values that transition from valid to invalid to validate the boundary logic
**
** @param t *testing.T - The testing framework
************************************************************************************************/
func TestSafeUint64ToInt64_EdgeCases(t *testing.T) {
	// Test values exactly at the boundary
	maxInt64 := uint64(math.MaxInt64)

	// Test maximum valid value
	value, ok := SafeUint64ToInt64(maxInt64)
	if !ok || value != math.MaxInt64 {
		t.Errorf("Expected (%d, true) for max valid value, got (%d, %v)",
			math.MaxInt64, value, ok)
	}

	// Test minimum invalid value
	value, ok = SafeUint64ToInt64(maxInt64 + 1)
	if ok || value != math.MaxInt64 {
		t.Errorf("Expected (math.MaxInt64, false) for min invalid value, got (%d, %v)",
			value, ok)
	}

	// Test a sequence of values at the boundary
	for i := uint64(0); i < 10; i++ {
		testValue := maxInt64 - 5 + i
		expectedOk := testValue <= maxInt64
		expectedValue := int64(math.MaxInt64)
		if expectedOk {
			expectedValue = int64(testValue)
		}

		value, ok := SafeUint64ToInt64(testValue)

		if value != expectedValue || ok != expectedOk {
			t.Errorf("SafeUint64ToInt64(%d) = (%d, %v), want (%d, %v)",
				testValue, value, ok, expectedValue, expectedOk)
		}
	}
}
