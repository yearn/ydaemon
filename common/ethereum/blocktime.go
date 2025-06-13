package ethereum

import (
	"bufio"
	"context"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"math/big"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common/math"
	"github.com/yearn/ydaemon/common/env"
	"github.com/yearn/ydaemon/common/helpers"
	"github.com/yearn/ydaemon/common/logs"
)

var (
	/**************************************************************************************************
	** blockTimeMap maintains historical block numbers for specific time periods (7, 14, 30 days ago)
	** for each chain ID. This map is populated by InitBlockTimestamp and used by GetBlockNumberXDaysAgo.
	**
	** The structure is: map[chainID]map[daysPeriod]blockNumber
	**************************************************************************************************/
	blockTimeMap   = make(map[uint64]map[uint64]uint64)
	blockTimeMutex sync.RWMutex
)

/**************************************************************************************************
** BlockTimeData represents the structure of our block time data storage.
** It contains a mapping of chain IDs to their respective block time data.
**************************************************************************************************/
type BlockTimeData struct {
	Chains map[uint64]ChainBlockData
}

/**************************************************************************************************
** ChainBlockData stores block data for a specific chain.
** It maintains bidirectional mappings between timestamps and block numbers.
**************************************************************************************************/
type ChainBlockData struct {
	TimeBlocks map[uint64]uint64 // timestamp -> block number
	BlockTimes map[uint64]uint64 // block number -> timestamp
}

/**************************************************************************************************
** TimestampBlockPair represents a simple timestamp to block mapping.
** It's used for processing and storing block time data, especially when reading from
** or writing to CSV files.
**************************************************************************************************/
type TimestampBlockPair struct {
	Timestamp uint64
	Block     uint64
	Date      string // Optional human-readable date
}

var (
	/**************************************************************************************************
	** In-memory representation of the block time data and the directory where CSV files are stored.
	** blockTimeData holds all block time mappings loaded from CSV files.
	** blockTimeDataDir defines the directory path where CSV files are stored, one per chain.
	**************************************************************************************************/
	blockTimeData    *BlockTimeData
	blockTimeDataDir string = "data/blocktime"
)

/************************************************************************************************
** SafeUint64ToInt64 safely converts a uint64 to an int64 with bounds checking.
**
** This helper function prevents overflow errors when converting from unsigned to signed integers.
** It ensures that the uint64 value does not exceed the maximum value that can be represented
** by an int64 (math.MaxInt64).
**
** @param value uint64 - The unsigned 64-bit integer to convert
** @return int64 - The converted signed 64-bit integer (or math.MaxInt64 if conversion would overflow)
** @return bool - True if conversion was successful, false if it would overflow
************************************************************************************************/
func SafeUint64ToInt64(value uint64) (int64, bool) {
	// Check if the uint64 value exceeds the maximum int64 value
	if value > math.MaxInt64 {
		return math.MaxInt64, false
	}

	// Safe to convert
	return int64(value), true
}

/**************************************************************************************************
** InitBlockTimestamp initializes block timestamp mappings for a specific chain ID. It fetches
** block numbers for timestamps from 7 days ago, 14 days ago, and 30 days ago using the chain's
** Etherscan-compatible API.
**
** The results are stored in blockTimeMap for later retrieval via GetBlockNumberXDaysAgo.
** If the API call fails, it falls back to an estimate based on the chain's average blocks per day.
**
** @param chainID The chain ID to initialize block timestamps for
**************************************************************************************************/
func InitBlockTimestamp(chainID uint64) {
	type TScanResult struct {
		Status  string `json:"status"`
		Message string `json:"message"`
		Result  string `json:"result"`
	}
	var err error
	now := time.Now()
	noonUTC := time.Date(now.Year(), now.Month(), now.Day(), 12, 0, 0, 0, time.UTC)
	lastWeekTimestamp := noonUTC.AddDate(0, 0, -7).Unix()
	lastTwoWeekTimestamp := noonUTC.AddDate(0, 0, -14).Unix()
	lastMonthTimestamp := noonUTC.AddDate(0, -1, 0).Unix()

	chain, ok := env.GetChain(chainID)
	if !ok {
		return
	}

	APIKey := os.Getenv("SCAN_API_KEY")
	lastWeekBlock := helpers.FetchJSON[TScanResult](chain.EtherscanURI + `?chainid=` + strconv.FormatUint(chainID, 10) + `&module=block&action=getblocknobytime&timestamp=` + strconv.FormatInt(lastWeekTimestamp, 10) + `&closest=before&apikey=` + APIKey)
	lastTwoWeekBlock := helpers.FetchJSON[TScanResult](chain.EtherscanURI + `?chainid=` + strconv.FormatUint(chainID, 10) + `&module=block&action=getblocknobytime&timestamp=` + strconv.FormatInt(lastTwoWeekTimestamp, 10) + `&closest=before&apikey=` + APIKey)
	lastMonthBlock := helpers.FetchJSON[TScanResult](chain.EtherscanURI + `?chainid=` + strconv.FormatUint(chainID, 10) + `&module=block&action=getblocknobytime&timestamp=` + strconv.FormatInt(lastMonthTimestamp, 10) + `&closest=before&apikey=` + APIKey)

	blockTimeMutex.Lock()
	defer blockTimeMutex.Unlock()

	if blockTimeMap[chainID] == nil {
		blockTimeMap[chainID] = make(map[uint64]uint64)
	}
	if lastWeekBlock.Status == "1" {
		blockTimeMap[chainID][7], err = strconv.ParseUint(lastWeekBlock.Result, 10, 64)
		if err != nil {
			blockTimeMap[chainID][7] = uint64(chain.AvgBlocksPerDay * 7)
		}
	} else {
		blockTimeMap[chainID][7] = uint64(chain.AvgBlocksPerDay * 7)
	}

	if lastTwoWeekBlock.Status == "1" {
		blockTimeMap[chainID][14], err = strconv.ParseUint(lastTwoWeekBlock.Result, 10, 64)
		if err != nil {
			blockTimeMap[chainID][14] = uint64(chain.AvgBlocksPerDay * 14)
		}
	} else {
		blockTimeMap[chainID][14] = uint64(chain.AvgBlocksPerDay * 14)
	}

	if lastMonthBlock.Status == "1" {
		blockTimeMap[chainID][30], err = strconv.ParseUint(lastMonthBlock.Result, 10, 64)
		if err != nil {
			blockTimeMap[chainID][30] = uint64(chain.AvgBlocksPerDay * 30)
		}
	} else {
		blockTimeMap[chainID][30] = uint64(chain.AvgBlocksPerDay * 30)
	}
}

/**************************************************************************************************
** GetBlockNumberXDaysAgo retrieves a block number from X days ago for a given chain.
** It uses the data initialized by InitBlockTimestamp.
**
** This function only supports specific intervals: 7, 14, and 30 days.
** For different intervals or a more flexible approach, consider using GetBlockNumberByPeriod.
**
** @param chainID The chain ID to get the block number for
** @param days The number of days to look back (must be 7, 14, or 30)
** @return uint64 The block number from X days ago, or 0 if not available or not supported
**************************************************************************************************/
func GetBlockNumberXDaysAgo(chainID uint64, days uint64) uint64 {
	blockTimeMutex.RLock()
	defer blockTimeMutex.RUnlock()

	if (days != 7 && days != 14 && days != 30) || blockTimeMap[chainID] == nil {
		return 0
	}
	return blockTimeMap[chainID][days]
}

/**************************************************************************************************
** InitBlockTimeData initializes the block time data storage system by loading from CSV files.
** If the CSV files don't exist, it creates empty data structures.
**
** This function is called during application startup to ensure block time data is available for
** all operations.
**************************************************************************************************/
func InitBlockTimeData() {
	blocktimeLog("==== STARTING BLOCKTIME DATA INITIALIZATION ====")

	// Ensure the data directory exists
	if _, err := os.Stat(blockTimeDataDir); os.IsNotExist(err) {
		blocktimeLog(fmt.Sprintf("Creating blocktime data directory: %s", blockTimeDataDir))
		if err := os.MkdirAll(blockTimeDataDir, 0755); err != nil {
			blocktimeError(fmt.Sprintf("Failed to create data directory: %v", err))
			return
		}
	}

	// Initialize the in-memory data structure
	blockTimeData = &BlockTimeData{
		Chains: make(map[uint64]ChainBlockData),
	}

	// Get a list of all CSV files in the directory
	files, err := filepath.Glob(filepath.Join(blockTimeDataDir, "*.csv"))
	if err != nil {
		logs.Warning("Failed to list blocktime CSV files:", err)
		return
	}

	blocktimeLog(fmt.Sprintf("Found %d CSV blocktime files in %s", len(files), blockTimeDataDir))

	if len(files) == 0 {
		blocktimeLog("No blocktime CSV files found, starting with empty data")
		return
	}

	// Process each CSV file (one per chain)
	for _, file := range files {
		// Extract chainID from filename (e.g., "1.csv" -> 1)
		baseFilename := filepath.Base(file)
		chainIDStr := strings.TrimSuffix(baseFilename, filepath.Ext(baseFilename))
		chainID, err := strconv.ParseUint(chainIDStr, 10, 64)
		if err != nil {
			logs.Warning(fmt.Sprintf("Invalid chainID in filename: %s", baseFilename))
			continue
		}

		blocktimeLog(fmt.Sprintf("Chain %d - Processing blocktime file: %s", chainID, baseFilename))

		// Load data from this CSV file
		pairs, err := loadBlocktimeFromCSV(chainID)
		if err != nil {
			logs.Warning(fmt.Sprintf("Chain %d - Failed to load blocktime data: %v", chainID, err))
			continue
		}

		// Create chain data structure if it doesn't exist
		if _, exists := blockTimeData.Chains[chainID]; !exists {
			blockTimeData.Chains[chainID] = ChainBlockData{
				TimeBlocks: make(map[uint64]uint64),
				BlockTimes: make(map[uint64]uint64),
			}
		}

		// Add the data to our in-memory structure
		for _, pair := range pairs {
			blockTimeData.Chains[chainID].TimeBlocks[pair.Timestamp] = pair.Block
			blockTimeData.Chains[chainID].BlockTimes[pair.Block] = pair.Timestamp
		}

		blocktimeSuccess(fmt.Sprintf("Chain %d - Loaded %d blocktime records", chainID, len(pairs)))

		// Check if we need to update with newer data
		updateBlocktimeUntilToday(chainID, pairs)
	}

	blocktimeLog("==== BLOCKTIME DATA INITIALIZATION COMPLETE ====")
}

/**************************************************************************************************
** loadBlocktimeFromCSV loads block time data from a CSV file for a specific chain.
**
** @param chainID The chain ID to load data for
** @return []TimestampBlockPair Array of timestamp-block pairs
** @return error An error if loading fails, or nil on success
**************************************************************************************************/
func loadBlocktimeFromCSV(chainID uint64) ([]TimestampBlockPair, error) {
	filePath := filepath.Join(blockTimeDataDir, fmt.Sprintf("%d.csv", chainID))

	blocktimeLog(fmt.Sprintf("Chain %d - Loading blocktime data from %s", chainID, filePath))

	// Check if file exists
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		blocktimeWarning(fmt.Sprintf("Chain %d - Blocktime file not found: %s", chainID, filePath))
		return nil, fmt.Errorf("blocktime file not found for chain %d", chainID)
	}

	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		blocktimeError(fmt.Sprintf("Chain %d - Failed to open blocktime file: %v", chainID, err))
		return nil, err
	}
	defer file.Close()

	// Parse CSV
	reader := csv.NewReader(file)

	// Read header
	header, err := reader.Read()
	if err != nil {
		blocktimeError(fmt.Sprintf("Chain %d - Failed to read CSV header: %v", chainID, err))
		return nil, err
	}

	// Verify header format
	if len(header) < 2 || header[0] != "timestamp" || header[1] != "blocknumber" {
		blocktimeError(fmt.Sprintf("Chain %d - Invalid CSV format in %s: expected timestamp,blocknumber headers", chainID, filePath))
		return nil, fmt.Errorf("invalid CSV format: expected timestamp,blocknumber headers")
	}

	pairs := make([]TimestampBlockPair, 0)
	lineNum := 1 // Skip header

	// Read and parse each line
	for {
		lineNum++
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			blocktimeWarning(fmt.Sprintf("Chain %d - Error reading line %d in %s: %v", chainID, lineNum, filePath, err))
			continue
		}

		if len(record) < 2 {
			blocktimeWarning(fmt.Sprintf("Chain %d - Invalid format at line %d in %s: too few columns", chainID, lineNum, filePath))
			continue
		}

		// Parse timestamp and block number
		timestamp, err := strconv.ParseUint(record[0], 10, 64)
		if err != nil {
			blocktimeWarning(fmt.Sprintf("Chain %d - Invalid timestamp at line %d in %s: %v", chainID, lineNum, filePath, err))
			continue
		}

		block, err := strconv.ParseUint(record[1], 10, 64)
		if err != nil {
			blocktimeWarning(fmt.Sprintf("Chain %d - Invalid block number at line %d in %s: %v", chainID, lineNum, filePath, err))
			continue
		}

		// Get date if available
		date := ""
		if len(record) > 2 {
			date = record[2]
		}

		pairs = append(pairs, TimestampBlockPair{
			Timestamp: timestamp,
			Block:     block,
			Date:      date,
		})
	}

	// Sort by timestamp (oldest first) for consistency
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].Timestamp < pairs[j].Timestamp
	})

	// Log the date range of the data
	if len(pairs) > 0 {
		firstTimestamp, ok := SafeUint64ToInt64(pairs[0].Timestamp)
		if !ok {
			blocktimeWarning(fmt.Sprintf("Chain %d - Timestamp overflow for first record: %d",
				chainID, pairs[0].Timestamp))
			firstTimestamp = 0
		}

		lastTimestamp, ok := SafeUint64ToInt64(pairs[len(pairs)-1].Timestamp)
		if !ok {
			blocktimeWarning(fmt.Sprintf("Chain %d - Timestamp overflow for last record: %d",
				chainID, pairs[len(pairs)-1].Timestamp))
			lastTimestamp = 0
		}

		firstDate := time.Unix(firstTimestamp, 0).UTC().Format("2006-01-02")
		lastDate := time.Unix(lastTimestamp, 0).UTC().Format("2006-01-02")
		blocktimeLog(fmt.Sprintf("Chain %d - Loaded %d records spanning from %s to %s",
			chainID, len(pairs), firstDate, lastDate))
	} else {
		blocktimeLog(fmt.Sprintf("Chain %d - Loaded 0 records from %s", chainID, filePath))
	}

	return pairs, nil
}

/**************************************************************************************************
** updateBlocktimeUntilToday checks the last timestamp in the provided data and fetches
** new timestamps up to today at noon UTC. New data is stored in the CSV file.
**
** @param chainID The chain ID to update data for
** @param existingPairs Existing timestamp-block pairs loaded from CSV
**************************************************************************************************/
func updateBlocktimeUntilToday(chainID uint64, existingPairs []TimestampBlockPair) {
	blocktimeLog(fmt.Sprintf("Chain %d - Checking if blocktime data needs to be updated...", chainID))

	if len(existingPairs) == 0 {
		blocktimeLog(fmt.Sprintf("Chain %d - No existing blocktime data, will populate from scratch", chainID))
		fetchBlocktimeForDateRange(chainID, nil, nil)
		return
	}

	// Find the last timestamp in the data
	lastPair := existingPairs[len(existingPairs)-1]
	lastTimestamp, ok := SafeUint64ToInt64(lastPair.Timestamp)
	if !ok {
		blocktimeWarning(fmt.Sprintf("Chain %d - Timestamp overflow for last record: %d",
			chainID, lastPair.Timestamp))
		return // Skip update if timestamp is invalid
	}
	lastTime := time.Unix(lastTimestamp, 0).UTC()

	blocktimeLog(fmt.Sprintf("Chain %d - Last blocktime record is from %s (timestamp %d, block %d)",
		chainID, lastTime.Format("2006-01-02 15:04:05"), lastPair.Timestamp, lastPair.Block))

	// Get today's date at noon UTC
	now := time.Now().UTC()
	today := time.Date(now.Year(), now.Month(), now.Day(), 12, 0, 0, 0, time.UTC)

	// If last timestamp is from today or future, nothing to update
	if lastTime.After(today) || lastTime.Equal(today) {
		blocktimeLog(fmt.Sprintf("Chain %d - Blocktime data is already up to date (last record: %s)",
			chainID, lastTime.Format("2006-01-02 15:04:05")))
		return
	}

	// Calculate the next day after the last timestamp
	nextDay := time.Date(lastTime.Year(), lastTime.Month(), lastTime.Day(), 12, 0, 0, 0, time.UTC).AddDate(0, 0, 1)

	blocktimeLog(fmt.Sprintf("Chain %d - Blocktime data needs updating from %s to %s (%d days)",
		chainID, nextDay.Format("2006-01-02"), today.Format("2006-01-02"),
		int(today.Sub(nextDay).Hours()/24)+1))

	// Fetch data for the missing date range
	fetchBlocktimeForDateRange(chainID, &nextDay, &today)
}

/**************************************************************************************************
** fetchBlocktimeForDateRange fetches block times for a range of dates at noon UTC.
** If startDate is nil, it starts from 365 days ago.
** If endDate is nil, it ends at today.
**
** @param chainID The chain ID to fetch data for
** @param startDate The start date (or nil for 365 days ago)
** @param endDate The end date (or nil for today)
**************************************************************************************************/
func fetchBlocktimeForDateRange(chainID uint64, startDate, endDate *time.Time) {
	now := time.Now().UTC()
	today := time.Date(now.Year(), now.Month(), now.Day(), 12, 0, 0, 0, time.UTC)

	// Default: start from 365 days ago if not specified
	start := today.AddDate(-1, 0, 0)
	if startDate != nil {
		start = *startDate
	}

	// Default: end at today if not specified
	end := today
	if endDate != nil {
		end = *endDate
	}

	chain, ok := env.GetChain(chainID)
	if !ok {
		blocktimeWarning(fmt.Sprintf("Chain %d - Chain not found in configuration, cannot fetch blocktime data", chainID))
		return
	}

	blocktimeLog(fmt.Sprintf("Chain %d - Starting blocktime data fetch from %s to %s",
		chainID, start.Format("2006-01-02"), end.Format("2006-01-02")))

	// Calculate number of days
	days := int(end.Sub(start).Hours()/24) + 1

	// Prepare a slice to store new pairs
	newPairs := make([]TimestampBlockPair, 0, days)

	// Get EtherScan API key
	APIKey := os.Getenv("SCAN_API_KEY_FOR_" + strconv.FormatUint(chainID, 10))
	if APIKey == "" {
		blocktimeWarning(fmt.Sprintf("Chain %d - No Etherscan API key found in environment, API rate limits may apply", chainID))
	}

	fetchCount := 0
	skipCount := 0
	errorCount := 0

	// For each day in the range
	for d := 0; d < days; d++ {
		currentDate := start.AddDate(0, 0, d)
		timestamp := uint64(currentDate.Unix())

		// Skip if we already have this timestamp
		if _, exists := GetTimeBlock(chainID, timestamp); exists {
			skipCount++
			if skipCount%10 == 0 || skipCount == 1 {
				blocktimeLog(fmt.Sprintf("Chain %d - Skipped %d timestamps that already exist in data", chainID, skipCount))
			}
			continue
		}

		blocktimeLog(fmt.Sprintf("Chain %d - Fetching block at %s (timestamp %d)",
			chainID, currentDate.Format("2006-01-02 15:04:05"), timestamp))

		// Fetch block number from chain's explorer API
		blockNumber, err := fetchBlockNumberFromAPI(chain, timestamp, APIKey)
		if err != nil {
			errorCount++
			blocktimeWarning(fmt.Sprintf("Chain %d - Failed to fetch block for %s: %v",
				chainID, currentDate.Format("2006-01-02 15:04:05"), err))
			continue
		}

		fetchCount++
		blocktimeSuccess(fmt.Sprintf("Chain %d - ✓ Found block %d for %s (timestamp %d)",
			chainID, blockNumber, currentDate.Format("2006-01-02 15:04:05"), timestamp))

		// Format date string
		dateStr := currentDate.Format("02/01/2006")

		// Add to new pairs
		newPairs = append(newPairs, TimestampBlockPair{
			Timestamp: timestamp,
			Block:     blockNumber,
			Date:      dateStr,
		})

		// Store in memory
		blockTimeData.Chains[chainID].TimeBlocks[timestamp] = blockNumber
		blockTimeData.Chains[chainID].BlockTimes[blockNumber] = timestamp

		// Add small delay to avoid rate limiting
		time.Sleep(200 * time.Millisecond)
	}

	if len(newPairs) > 0 {
		blocktimeSuccess(fmt.Sprintf("Chain %d - Completed blocktime fetch: %d new blocks, %d skipped, %d errors",
			chainID, fetchCount, skipCount, errorCount))

		// Append to CSV file
		appendBlocktimeToCSV(chainID, newPairs)
	} else {
		blocktimeLog(fmt.Sprintf("Chain %d - No new blocktime records needed (%d skipped, %d errors)",
			chainID, skipCount, errorCount))
	}
}

/**************************************************************************************************
** fetchBlockNumberFromAPI fetches a block number for a given timestamp from the appropriate API.
**
** @param chain The chain configuration
** @param timestamp The Unix timestamp to get the block for
** @param apiKey The API key for the chain's explorer
** @return uint64 The block number
** @return error An error if fetching fails
**************************************************************************************************/
func fetchBlockNumberFromAPI(chain env.TChain, timestamp uint64, apiKey string) (uint64, error) {
	type TScanResult struct {
		Status  string `json:"status"`
		Message string `json:"message"`
		Result  string `json:"result"`
	}

	// Convert timestamp to date string for logging
	timestampInt64, ok := SafeUint64ToInt64(timestamp)
	if !ok {
		blocktimeWarning(fmt.Sprintf("Chain %d - Timestamp overflow: %d", chain.ID, timestamp))
		return 0, fmt.Errorf("timestamp value too large: %d", timestamp)
	}
	dateStr := time.Unix(timestampInt64, 0).UTC().Format("2006-01-02 15:04:05")
	blocktimeLog(fmt.Sprintf("Chain %d - Attempting to fetch block for %s", chain.ID, dateStr))

	// Try to use DefiLlama API first for supported chains
	chainName := chainIDToName(chain.ID)
	if chainName != "Unknown" {
		defillamaURI := fmt.Sprintf("https://coins.llama.fi/block/%s/%d", chainName, timestamp)
		blocktimeLog(fmt.Sprintf("Chain %d - Trying DeFiLlama API: %s", chain.ID, defillamaURI))

		resp, err := http.Get(defillamaURI)
		if err != nil {
			blocktimeWarning(fmt.Sprintf("Chain %d - DeFiLlama API request failed: %v", chain.ID, err))
		} else if resp.StatusCode == 200 {
			defer resp.Body.Close()
			body, err := io.ReadAll(resp.Body)
			if err == nil {
				var data TLlamaBlock
				if err := json.Unmarshal(body, &data); err != nil {
					blocktimeWarning(fmt.Sprintf("Chain %d - Failed to parse DeFiLlama response: %v", chain.ID, err))
				} else if data.Height > 0 {
					blocktimeSuccess(fmt.Sprintf("Chain %d - DeFiLlama API success: block %d at %s",
						chain.ID, data.Height, dateStr))
					return data.Height, nil
				} else {
					blocktimeWarning(fmt.Sprintf("Chain %d - DeFiLlama returned invalid block height: %d", chain.ID, data.Height))
				}
			}
		} else {
			blocktimeWarning(fmt.Sprintf("Chain %d - DeFiLlama API returned status %d", chain.ID, resp.StatusCode))
		}
		if resp != nil && resp.Body != nil {
			resp.Body.Close()
		}
	}

	// Fallback to Etherscan API
	if chain.EtherscanURI != "" {
		apiURL := fmt.Sprintf("%s?chainid=%d&module=block&action=getblocknobytime&timestamp=%d&closest=before&apikey=%s",
			chain.EtherscanURI, chain.ID, timestamp, apiKey)

		blocktimeLog(fmt.Sprintf("Chain %d - Trying Etherscan compatible API: %s",
			chain.ID, chain.EtherscanURI))

		result := helpers.FetchJSON[TScanResult](apiURL)
		if result.Status == "1" {
			blockNumber, err := strconv.ParseUint(result.Result, 10, 64)
			if err == nil {
				blocktimeSuccess(fmt.Sprintf("Chain %d - Etherscan API success: block %d at %s",
					chain.ID, blockNumber, dateStr))
				return blockNumber, nil
			}
			blocktimeWarning(fmt.Sprintf("Chain %d - Failed to parse block number from Etherscan: %v", chain.ID, err))
			return 0, err
		}
		blocktimeWarning(fmt.Sprintf("Chain %d - Etherscan API error: status=%s, message=%s",
			chain.ID, result.Status, result.Message))
		return 0, fmt.Errorf("API returned status: %s, message: %s", result.Status, result.Message)
	}

	// Last resort: estimate based on average blocks per day
	if chain.AvgBlocksPerDay > 0 {
		blocktimeLog(fmt.Sprintf("Chain %d - Attempting to estimate block using average blocks per day: %d",
			chain.ID, chain.AvgBlocksPerDay))

		now := time.Now().UTC()
		timestampInt64, ok := SafeUint64ToInt64(timestamp)
		if !ok {
			blocktimeWarning(fmt.Sprintf("Chain %d - Timestamp overflow: %d", chain.ID, timestamp))
			return 0, fmt.Errorf("timestamp value too large: %d", timestamp)
		}
		daysDiff := int(now.Sub(time.Unix(timestampInt64, 0).UTC()).Hours() / 24)
		if daysDiff >= 0 {
			// Get current block number
			latestBlock, err := RPC[chain.ID].BlockNumber(context.Background())
			if err == nil {
				// Estimate block number
				estimatedBlock := latestBlock - uint64(chain.AvgBlocksPerDay*daysDiff)
				blocktimeWarning(fmt.Sprintf("Chain %d - Using estimated block %d for %s (current: %d, days: %d)",
					chain.ID, estimatedBlock, dateStr, latestBlock, daysDiff))
				return estimatedBlock, nil
			}
			blocktimeWarning(fmt.Sprintf("Chain %d - Failed to get current block number: %v", chain.ID, err))
		} else {
			blocktimeWarning(fmt.Sprintf("Chain %d - Timestamp %d is in the future, cannot estimate block", chain.ID, timestamp))
		}
	}

	blocktimeError(fmt.Sprintf("Chain %d - All methods failed to fetch block for timestamp %d (%s)", chain.ID, timestamp, dateStr))
	return 0, fmt.Errorf("could not fetch or estimate block number")
}

/**************************************************************************************************
** appendBlocktimeToCSV appends new timestamp-block pairs to the CSV file for a specific chain.
** If the file doesn't exist, it creates a new one.
**
** @param chainID The chain ID to append data for
** @param pairs The timestamp-block pairs to append
**************************************************************************************************/
func appendBlocktimeToCSV(chainID uint64, pairs []TimestampBlockPair) {
	filePath := filepath.Join(blockTimeDataDir, fmt.Sprintf("%d.csv", chainID))
	blocktimeLog(fmt.Sprintf("Chain %d - Saving %d new blocktime records to %s", chainID, len(pairs), filePath))

	// Check if file exists
	fileExists := true
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		fileExists = false
		blocktimeLog(fmt.Sprintf("Chain %d - Creating new CSV file: %s", chainID, filePath))
	} else {
		blocktimeLog(fmt.Sprintf("Chain %d - Appending to existing CSV file: %s", chainID, filePath))
	}

	// Open file in append mode, or create if it doesn't exist
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		blocktimeWarning(fmt.Sprintf("Chain %d - Failed to open blocktime CSV file: %v", chainID, err))
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)

	// Write header if new file
	if !fileExists {
		if _, err := writer.WriteString("timestamp,blocknumber,date\n"); err != nil {
			blocktimeWarning(fmt.Sprintf("Chain %d - Failed to write CSV header: %v", chainID, err))
			return
		}
		blocktimeLog(fmt.Sprintf("Chain %d - Wrote CSV header to new file", chainID))
	}

	// Write each pair
	writtenPairs := 0
	for _, pair := range pairs {
		timestampInt64, ok := SafeUint64ToInt64(pair.Timestamp)
		if !ok {
			blocktimeWarning(fmt.Sprintf("Chain %d - Skipping record with timestamp overflow: %d",
				chainID, pair.Timestamp))
			continue
		}

		// If date is not already set, format it
		pairDate := pair.Date
		if pairDate == "" {
			pairDate = time.Unix(timestampInt64, 0).UTC().Format("02/01/2006")
		}

		line := fmt.Sprintf("%d,%d,%s\n", pair.Timestamp, pair.Block, pairDate)
		if _, err := writer.WriteString(line); err != nil {
			blocktimeWarning(fmt.Sprintf("Chain %d - Failed to write CSV line for timestamp %d: %v",
				chainID, pair.Timestamp, err))
			continue
		}
		writtenPairs++
	}

	// Flush the buffer to ensure all data is written to disk
	if err := writer.Flush(); err != nil {
		blocktimeWarning(fmt.Sprintf("Chain %d - Failed to flush CSV data: %v", chainID, err))
		return
	}

	blocktimeSuccess(fmt.Sprintf("Chain %d - Successfully wrote %d/%d records to CSV file",
		chainID, writtenPairs, len(pairs)))
}

/**************************************************************************************************
** GetTimeBlock retrieves a block number for a given chain ID and timestamp from the internal
** storage. It returns the block number and a boolean indicating whether the data was found.
**
** @param chainID The chain ID to get the block for
** @param timestamp The Unix timestamp to find the corresponding block for
** @return uint64 The block number, or 0 if not found
** @return bool True if found, false otherwise
**************************************************************************************************/
func GetTimeBlock(chainID uint64, timestamp uint64) (uint64, bool) {
	blockTimeMutex.RLock()
	defer blockTimeMutex.RUnlock()

	chainData, exists := blockTimeData.Chains[chainID]
	if !exists {
		return 0, false
	}

	blockNum, exists := chainData.TimeBlocks[timestamp]
	return blockNum, exists
}

/**************************************************************************************************
** StoreTimeBlock saves a mapping from timestamp to block number for a specific chain in the
** internal storage system and persists it to the CSV file.
**
** @param chainID The chain ID to store the block for
** @param timestamp The Unix timestamp
** @param blockNumber The corresponding block number
** @return error An error if saving fails, or nil on success
**************************************************************************************************/
func StoreTimeBlock(chainID uint64, timestamp uint64, blockNumber uint64) error {
	blockTimeMutex.Lock()
	defer blockTimeMutex.Unlock()

	// Initialize chain data if it doesn't exist
	chainData, exists := blockTimeData.Chains[chainID]
	if !exists {
		chainData = ChainBlockData{
			TimeBlocks: make(map[uint64]uint64),
			BlockTimes: make(map[uint64]uint64),
		}
		blockTimeData.Chains[chainID] = chainData
	}

	// Check if we already have this data
	if existingBlock, exists := chainData.TimeBlocks[timestamp]; exists && existingBlock == blockNumber {
		return nil // Data already exists and matches
	}

	// Store in memory
	blockTimeData.Chains[chainID].TimeBlocks[timestamp] = blockNumber
	blockTimeData.Chains[chainID].BlockTimes[blockNumber] = timestamp

	// Format date string
	date := time.Unix(int64(timestamp), 0).UTC().Format("02/01/2006")

	// Append to CSV file
	newPair := TimestampBlockPair{
		Timestamp: timestamp,
		Block:     blockNumber,
		Date:      date,
	}

	appendBlocktimeToCSV(chainID, []TimestampBlockPair{newPair})

	return nil
}

/**************************************************************************************************
** GetBlockTime will try, for a specific blockNumber on a specific chain, to find its execution
** timestamp. This timestamp should be available in the internal storage. If it is not, it will
** fetch it from the blockchain and store it in the internal storage.
**************************************************************************************************/
func GetBlockTime(chainID uint64, blockNumber uint64) (blockTime uint64) {
	blockTimeMutex.RLock()
	chainData, exists := blockTimeData.Chains[chainID]
	if exists {
		timestamp, found := chainData.BlockTimes[blockNumber]
		if found {
			blockTimeMutex.RUnlock()
			return timestamp
		}
	}
	blockTimeMutex.RUnlock()

	// Not found in our data, fetch from blockchain
	client := RPC[chainID]
	block, err := client.HeaderByNumber(context.Background(), big.NewInt(int64(blockNumber)))
	if err != nil {
		blocktimeWarning(fmt.Sprintf("impossible to retrieve block %s on chain %s",
			strconv.FormatUint(blockNumber, 10), strconv.FormatUint(chainID, 10)))
		return 0
	}

	// Store the retrieved data
	timestamp := block.Time
	_ = StoreTimeBlock(chainID, timestamp, blockNumber)

	return timestamp
}

/**************************************************************************************************
** TLlamaBlock represents the response structure from DeFiLlama API.
** It contains the block height and timestamp data returned by the API.
**************************************************************************************************/
type TLlamaBlock struct {
	Height    uint64 `json:"height"`
	Timestamp uint64 `json:"timestamp"`
}

/**************************************************************************************************
** chainIDToName maps a blockchain's chain ID to the name expected by the DeFiLlama API.
** This function supports converting common chain IDs to their respective network names.
**
** @param chainID The chain ID to convert to a name
** @return string The corresponding network name for DeFiLlama API, or "Unknown" if not supported
**************************************************************************************************/
func chainIDToName(chainID uint64) string {
	switch chainID {
	case 1:
		return "ethereum"
	case 10:
		return "optimism"
	case 100:
		return "xdai"
	case 137:
		return "polygon"
	case 250:
		return "fantom"
	case 8453:
		return "base"
	case 42161:
		return "arbitrum"
	case 43114:
		return "avalanche"
	default:
		return "Unknown"
	}
}

var (
	/**************************************************************************************************
	** Map to store historical block numbers by chainID and time period.
	** This cache prevents redundant API calls by storing previously fetched historical blocks.
	** The structure is: map[chainID]map[periodString]blockNumber where periodString is "24h", "7d", etc.
	**************************************************************************************************/
	historicalBlockMap   = make(map[uint64]map[string]uint64)
	historicalBlockMutex sync.RWMutex
)

/**************************************************************************************************
** GetHistoricalBlockNumbers fetches block numbers for specific time periods (24h ago, 7 days ago,
** 30 days ago, and 365 days ago) for a given chain.
**
** The function first tries to use the internal storage (CSV-based data), and only fetches
** from external APIs if needed.
**
** @param chainID The chain ID to fetch historical block numbers for
** @return map[string]uint64 A map with keys "24h", "7d", "30d", "365d" and corresponding block numbers
**************************************************************************************************/
func GetHistoricalBlockNumbers(chainID uint64) map[string]uint64 {
	blocktimeLog(fmt.Sprintf("Chain %d - Retrieving historical block numbers for time periods", chainID))

	// Check if we already have the data for this chain
	historicalBlockMutex.RLock()
	if blockData, exists := historicalBlockMap[chainID]; exists && len(blockData) == 4 {
		blocktimeLog(fmt.Sprintf("Chain %d - Using cached historical block data", chainID))
		for period, block := range blockData {
			blocktimeLog(fmt.Sprintf("Chain %d - Period %s: block %d", chainID, period, block))
		}
		historicalBlockMutex.RUnlock()
		return blockData
	}
	historicalBlockMutex.RUnlock()

	// Calculate timestamps for each period
	now := time.Now()
	noonUTC := time.Date(now.Year(), now.Month(), now.Day(), 12, 0, 0, 0, time.UTC)

	timestamps := map[string]int64{
		"now":  noonUTC.Unix(),
		"24h":  noonUTC.AddDate(0, 0, -1).Unix(),
		"7d":   noonUTC.AddDate(0, 0, -7).Unix(),
		"30d":  noonUTC.AddDate(0, -1, 0).Unix(),
		"365d": noonUTC.AddDate(-1, 0, 0).Unix(),
	}

	blocktimeLog(fmt.Sprintf("Chain %d - Fetching blocks for time periods: %+v", chainID, timestamps))

	// Lock for writing
	historicalBlockMutex.Lock()
	defer historicalBlockMutex.Unlock()

	// Initialize the map for this chain if it doesn't exist
	if historicalBlockMap[chainID] == nil {
		historicalBlockMap[chainID] = make(map[string]uint64)
	}

	// For each period, try to get the block number
	for period, timestamp := range timestamps {
		blocktimeLog(fmt.Sprintf("Chain %d - Looking up block for period %s (timestamp %d, %s)",
			chainID, period, timestamp, time.Unix(timestamp, 0).UTC().Format("2006-01-02 15:04:05")))

		// Check if we already have the data in the internal storage (CSV)
		blockNumber, ok := GetTimeBlock(chainID, uint64(timestamp))
		if ok {
			blocktimeSuccess(fmt.Sprintf("Chain %d - Found block %d for period %s in CSV data",
				chainID, blockNumber, period))
			historicalBlockMap[chainID][period] = blockNumber
			continue
		}

		blocktimeLog(fmt.Sprintf("Chain %d - Block not found for period %s, fetching from API...", chainID, period))

		// If not found, fetch and store it
		chain, ok := env.GetChain(chainID)
		if !ok {
			blocktimeWarning(fmt.Sprintf("Chain %d - Chain not found in configuration", chainID))
			continue
		}

		APIKey := os.Getenv("SCAN_API_KEY_FOR_" + strconv.FormatUint(chainID, 10))
		blockNumber, err := fetchBlockNumberFromAPI(chain, uint64(timestamp), APIKey)
		if err != nil {
			blocktimeWarning(fmt.Sprintf("Chain %d - Failed to fetch historical block for period %s: %v",
				chainID, period, err))
			continue
		}

		blocktimeSuccess(fmt.Sprintf("Chain %d - Successfully fetched block %d for period %s",
			chainID, blockNumber, period))

		// Store in internal storage and in memory
		_ = StoreTimeBlock(chainID, uint64(timestamp), blockNumber)
		historicalBlockMap[chainID][period] = blockNumber
	}

	blocktimeSuccess(fmt.Sprintf("Chain %d - Historical block numbers fetched: %+v",
		chainID, historicalBlockMap[chainID]))
	return historicalBlockMap[chainID]
}

/**************************************************************************************************
** GetBlockNumberByPeriod retrieves a block number for a specific time period for the given chain.
** It uses the cached data if available, otherwise it calls GetHistoricalBlockNumbers to fetch
** the data.
**
** @param chainID The chain ID to get the block number for
** @param periodOrDays The time period ("now", "24h", "7d", "30d", "365d") or a numeric days value (0, 1, 7, 30, 365)
**               When using a numeric value: 0 = now, 1=24h, 7=7d, 30=30d, 365=365d
** @return uint64 The block number for the specified period, or 0 if not available
**************************************************************************************************/
func GetBlockNumberByPeriod(chainID uint64, periodOrDays interface{}) uint64 {
	var period string

	// Log stack trace to find caller
	if callerInfo := getCallerInfo(); callerInfo != "" {
		blocktimeLog(fmt.Sprintf("Chain %d - GetBlockNumberByPeriod called from: %s", chainID, callerInfo))
	}

	// Handle different parameter types
	switch value := periodOrDays.(type) {
	case string:
		// If it's already a string, use it directly
		period = value
		blocktimeLog(fmt.Sprintf("Chain %d - GetBlockNumberByPeriod received string period: '%s'", chainID, value))
	case int:
		// If it's a numeric integer, convert to appropriate period string
		blocktimeLog(fmt.Sprintf("Chain %d - GetBlockNumberByPeriod received int days: %d", chainID, value))
		period = daysToPeriod(uint64(value))
		blocktimeLog(fmt.Sprintf("Chain %d - Converted %d days to period: '%s'", chainID, value, period))
	case uint64:
		// If it's an unsigned integer, convert to appropriate period string
		blocktimeLog(fmt.Sprintf("Chain %d - GetBlockNumberByPeriod received uint64 days: %d", chainID, value))
		period = daysToPeriod(value)
		blocktimeLog(fmt.Sprintf("Chain %d - Converted %d days to period: '%s'", chainID, value, period))
	default:
		blocktimeWarning(fmt.Sprintf("Chain %d - Unsupported period type in GetBlockNumberByPeriod: %T with value %v",
			chainID, periodOrDays, periodOrDays))
		return 0
	}

	// Validate the period format
	if period != "now" && period != "24h" && period != "7d" && period != "30d" && period != "365d" {
		blocktimeWarning(fmt.Sprintf("Chain %d - Unsupported period in GetBlockNumberByPeriod: '%s' (original input: %v of type %T)",
			chainID, period, periodOrDays, periodOrDays))
		return 0
	}

	// Check if we have data and fetch if needed
	historicalBlockMutex.RLock()
	blockData, exists := historicalBlockMap[chainID]
	historicalBlockMutex.RUnlock()

	if !exists || blockData[period] == 0 {
		blocktimeLog(fmt.Sprintf("Chain %d - No cached data found for period '%s', fetching historical blocks", chainID, period))
		blockData = GetHistoricalBlockNumbers(chainID)
	} else {
		blocktimeLog(fmt.Sprintf("Chain %d - Using cached block number %d for period '%s'", chainID, blockData[period], period))
	}

	if blockNumber, ok := blockData[period]; ok {
		return blockNumber
	}

	blocktimeWarning(fmt.Sprintf("Chain %d - No block number found for period '%s' after fetching historical data", chainID, period))
	return 0
}

/**************************************************************************************************
** daysToPeriod converts a numeric days value to the corresponding period string format.
** This helper function is used by GetBlockNumberByPeriod to standardize period representations.
**
** @param days The number of days (1, 7, 30, or 365)
** @return string The corresponding period string ("24h", "7d", "30d", "365d") or "invalid"
**************************************************************************************************/
func daysToPeriod(days uint64) string {
	switch days {
	case 0:
		return "now"
	case 1:
		return "24h"
	case 7:
		return "7d"
	case 30:
		return "30d"
	case 365:
		return "365d"
	default:
		blocktimeWarning(fmt.Sprintf("Invalid days value in daysToperiod: %d (supported values are 1, 7, 30, 365)", days))
		return "invalid"
	}
}

/**************************************************************************************************
** ExportBlockTimeDataAsCSV exports all the block time data for a specific chain as a CSV string.
** Since we're already using CSV as the storage format, this essentially reads and returns the
** contents of the relevant CSV file.
**
** @param chainID The chain ID to export data for
** @return string The CSV data as a string, or an empty string if no data exists
**************************************************************************************************/
func ExportBlockTimeDataAsCSV(chainID uint64) string {
	filePath := filepath.Join(blockTimeDataDir, fmt.Sprintf("%d.csv", chainID))

	// Check if file exists
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		blocktimeWarning(fmt.Sprintf("Chain %d - CSV file does not exist", chainID))
		return ""
	}

	// Read file contents
	data, err := os.ReadFile(filePath)
	if err != nil {
		blocktimeWarning(fmt.Sprintf("Chain %d - Failed to read CSV file: %v", chainID, err))
		return ""
	}

	return string(data)
}

/**************************************************************************************************
** SaveBlockTimeDataToCSV saves block time data to a CSV file for a specific chain.
** Since we're already using CSV as the primary storage format, this is a no-op function
** that's kept for backward compatibility.
**
** @param chainID The chain ID to save data for
** @return error Always returns nil (success) as the data is already in CSV format
**************************************************************************************************/
func SaveBlockTimeDataToCSV(chainID uint64) error {
	// No-op function as the data is already stored in CSV format
	// Keep for backward compatibility
	return nil
}
