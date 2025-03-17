# Ethereum Package

> The ethereum package is the cornerstone of blockchain interactions within yDaemon, providing
> robust interfaces for communication with Ethereum and EVM-compatible networks. This package
> manages connections to RPC endpoints, handles block time data, processes contract interactions,
> and offers utilities for retrieving historical data across multiple chains.
> This documentation serves as a guide for developers working with the ethereum package.

## Core Components

### Connection Management

The package manages connections to blockchain networks through various methods:

- **RPC Connections**: Maintained in `RPC` map, initialized during application startup
- **WebSocket Connections**: Maintained in `WS` map, created on-demand with fallback mechanisms
- **Connection Configuration**: URIs are retrieved from environment variables with chain-specific overrides

```go
// Get an RPC client for a specific chain
client := ethereum.GetRPC(chainID)

// Get a WebSocket client with automatic protocol conversion
wsClient, err := ethereum.GetWSClient(chainID, true)
```

### Block Time Functionality

A key feature of this package is the extensive block time management system:

- **Block-to-Time Mapping**: Bidirectional mapping between timestamps and block numbers
- **CSV-Based Storage**: Persistent storage of block time data in CSV format
- **Automatic Updates**: Fetches missing block time data on startup
- **Historical Block Retrieval**: Functions to get blocks from specific time periods (24h, 7d, 30d, 365d ago)
- **Multi-Source Fetching**: Uses DefiLlama API with fallback to Etherscan-compatible APIs

```go
// Get a block number from 7 days ago
blockNumber := ethereum.GetBlockNumberByPeriod(chainID, "7d")

// Get timestamp for a specific block
timestamp := ethereum.GetBlockTime(chainID, blockNumber)
```

### Yearn Vault Interactions

The package provides specialized functions for interacting with Yearn vaults:

- **Price Per Share (PPS) Retrieval**: Functions to fetch current and historical PPS values
- **Historical Performance Data**: Retrieval of PPS from 7 days and 30 days ago for performance calculations

```go
// Get current PPS for a vault
ppsToday := ethereum.FetchPPSToday(chainID, vaultAddress, decimals)

// Get PPS from 7 days ago
ppsLastWeek := ethereum.FetchPPSLastWeek(chainID, vaultAddress, decimals)
```

### Multicall Support

Optimizes blockchain calls by batching multiple read operations into a single request:

- **Chain-Specific Clients**: Maintained in `MulticallClientForChainID` map
- **Configurability**: Supports custom RPC endpoints for multicall operations
- **Efficiency**: Reduces RPC call count and network overhead

### Verbose Logging System

The package includes a sophisticated logging system, especially for block time operations:

- **Verbosity Control**: Toggle verbose logging via environment variables or programmatically
- **Log Levels**: Different functions for info, warning, success, and error logs
- **Contextual Information**: Logs include chain IDs and operation details for easier debugging

## Key Files

- **utils.go**: Core utilities for RPC and WebSocket connections
- **blocktime.go**: Block time data management and historical block retrieval
- **blocktime.logger.go**: Logging utilities for block time operations
- **pps.go**: Price per share functions for Yearn vaults
- **initializer.go**: Application startup initialization
- **multicall.go**: Batch contract call functionality
- **stream.go**: WebSocket event streaming

## Initialization Process

The package follows a specific initialization sequence:

1. Establish RPC connections for all supported chains
2. Create multicall clients for each chain
3. Initialize block time data storage from CSV files
4. Update block time data for any missing periods
5. Initialize block timestamp mappings for specific periods

This happens automatically when the application starts via the `Initialize()` function.

## Best Practices

- **Chain ID Validation**: Always check that a chain is supported before operations
- **Error Handling**: Handle RPC errors gracefully, as network connectivity issues are common
- **Performance Considerations**: Use multicall for batching read operations
- **Logging Control**: Use verbose logging sparingly in production environments
- **Historical Data**: Prefer `GetBlockNumberByPeriod` over `GetBlockNumberXDaysAgo` for flexibility

## Common Pitfalls

- **WebSocket Limitations**: Not all chains support WebSocket connections
- **API Rate Limits**: Be aware of rate limits when frequently querying block time data
- **Missing Block Data**: Some historical blocks might not be available on certain chains
- **RPC Reliability**: Always implement retries and fallbacks for RPC operations

## Example Usage

### Retrieving Historical Performance

```go
// Calculate APY for a vault over 7 days
ppsToday := ethereum.FetchPPSToday(chainID, vaultAddress, decimals)
ppsLastWeek := ethereum.FetchPPSLastWeek(chainID, vaultAddress, decimals)

if !ppsLastWeek.IsZero() {
    percentageChange := ppsToday.Sub(ppsLastWeek).Div(ppsLastWeek).Mul(big.NewFloat(100))
    weeklyAPY := percentageChange.Mul(big.NewFloat(52)) // Annualize weekly growth
}
```

### Working with Block Time Data

```go
// Get a block from a specific timestamp
timestamp := time.Date(2023, 1, 1, 12, 0, 0, 0, time.UTC).Unix()
blockNumber, exists := ethereum.GetTimeBlock(chainID, uint64(timestamp))

// Get the timestamp for a specific block
blockTimestamp := ethereum.GetBlockTime(chainID, blockNumber)
```
