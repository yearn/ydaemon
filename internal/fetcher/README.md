# Fetcher Package

## Overview

The `fetcher` package is a core component of yDaemon that handles the retrieval and processing of blockchain data across multiple networks. This package serves as the primary interface for fetching comprehensive information about vaults, strategies, tokens, and their relationships within the Yearn Finance ecosystem.

The package implements efficient data retrieval patterns using multicalls, caching mechanisms, and optimized batch processing to minimize RPC calls while maintaining data accuracy. It processes raw blockchain data into standardized formats suitable for storage and API consumption.

## Package Structure

The fetcher package is organized into several logical components:

### Core Components

- `builder.go`: Functions for constructing and formatting vault and strategy objects
- `helper.go`: Utility functions for data processing and validation
- `strategies.go`: Strategy-specific data retrieval and processing
- `tokens.go`: Token information retrieval and relationship mapping
- `vaults.go`: Vault data retrieval and processing

### Data Models

The package works with several key data structures:

- `TVault`: Core vault information including addresses, metadata, and performance metrics
- `TStrategy`: Strategy details including performance data, allocation, and status
- `TERC20Token`: Token information including metadata and relationships

## Key Functions

### Vault Management

- `RetrieveAllVaults`: Fetches comprehensive vault information across chains
- `BuildVaultNames`: Constructs standardized vault names
- `BuildVaultSymbol`: Generates consistent vault symbols
- `BuildVaultTVL`: Calculates Total Value Locked metrics
- `BuildVaultCategory`: Determines vault categorization

### Strategy Management

- `RetrieveAllStrategies`: Fetches strategy information for specified vaults
- `fetchStrategiesBasicInformations`: Retrieves basic strategy data via multicalls
- `assignStrategy`: Processes and updates strategy information
- `getStrategyAPR`: Calculates strategy APR with version-specific handling

### Token Management

- `RetrieveAllTokens`: Orchestrates comprehensive token data collection
- `findAllTokens`: Processes token information retrieval
- `loadCurvePools`: Fetches Curve liquidity pool data
- `loadGammaPools`: Retrieves Gamma concentrated liquidity pool information
- `loadPendleTokens`: Collects Pendle yield token data
- `loadVeloTokens`: Fetches Velodrome ecosystem tokens

## Performance Optimizations

The package implements several optimizations to minimize RPC calls and improve efficiency:

1. **Multicall Batching**:

    - Groups multiple RPC calls into single multicall transactions
    - Reduces network overhead and improves response times
    - Implements adaptive batch sizing for different chains

2. **Caching Mechanisms**:

    - Implements strategy caching with configurable duration
    - Uses thread-safe storage with sync.Map
    - Prevents redundant fetches within cache window

3. **Efficient Data Processing**:
    - Processes tokens in chunks to manage memory usage
    - Implements early returns for inactive strategies
    - Uses optimized data structures for relationship mapping

## Usage Examples

### Fetching Vault Data

```go
// Initialize the fetcher
vaults := fetcher.RetrieveAllVaults(chainID, vaultsList, method)
```

### Processing Strategies

```go
// Fetch and process strategies
strategies := fetcher.RetrieveAllStrategies(chainID, strategiesMap)
```

### Token Information Retrieval

```go
// Get comprehensive token data
tokens := fetcher.RetrieveAllTokens(chainID, vaults)
```

## Best Practices

1. **Chain-Specific Handling**:

    - Always validate chain ID before processing
    - Use appropriate RPC endpoints for each chain
    - Handle chain-specific limitations (e.g., Fantom's RPC limits)

2. **Error Handling**:

    - Implement proper error propagation
    - Use early returns for invalid states
    - Log errors with appropriate context

3. **Data Validation**:

    - Validate addresses before processing
    - Check for nil values in responses
    - Verify data consistency across related objects

4. **Performance Considerations**:

    - Use appropriate batch sizes for multicalls
    - Implement caching where beneficial
    - Process data in chunks for large datasets

5. **Memory Management**:
    - Pre-allocate slices and maps when possible
    - Clean up temporary data structures
    - Use appropriate data types for values

## Dependencies

The package relies on several internal and external dependencies:

- `common/addresses`: Address validation and comparison
- `common/bigNumber`: Big number arithmetic
- `common/ethereum`: Ethereum client and RPC handling
- `internal/models`: Data structures and types
- `internal/multicalls`: Multicall implementation
- `internal/storage`: Data persistence layer

## Testing

The package includes comprehensive tests for all major components:

- Vault data retrieval and processing
- Strategy information handling
- Token relationship mapping
- Performance optimizations
- Error handling scenarios

## Contributing

When contributing to the fetcher package:

1. Follow the established code style and documentation format
2. Add tests for new functionality
3. Update documentation for significant changes
4. Consider performance implications of changes
5. Handle all error cases appropriately

## License

This package is part of the yDaemon project and follows the same licensing terms.
