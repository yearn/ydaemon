# Vaults Package

## Overview

The `vaults` package is a core component of yDaemon that provides API endpoints and data structures for interacting with Yearn Finance vaults across multiple blockchain networks. This package serves as the primary interface for clients to access comprehensive vault information, including vault details, APR/APY data, strategies, harvest data, TVL metrics, and user-specific earnings calculations.

## Package Structure

The vaults package is organized into several logical components:

### Data Models

- `models.go`: Core data structures for vault representation
- `models.strategies.go`: Strategy-related data structures and methods

### Data Preparation

- `prepare.getVaults.go`: Functions for retrieving and filtering vaults
- `prepare.getLegacyVaults.go`: Specialized functions for handling legacy vault formats
- `prepare.vaultObject.go`: Utility functions for processing vault objects and transforming data

### API Endpoints

- `route.vaults.go`: Main vault endpoint handlers
- `route.vaults.one.go`: Handlers for retrieving individual vault details
- `route.vaults.one.simplified.go`: Simplified representation of individual vaults
- `route.vaults.legacy.go`: Legacy format endpoints for backwards compatibility
- `route.vaults.blacklisted.go`: Endpoints for retrieving blacklisted vaults
- `route.vaults.some.go`: Endpoints for retrieving specific subsets of vaults
- `route.vaults.earned.go`: Earnings calculation endpoints with FIFO methodology
- `route.vaults.tvl.go`: Total Value Locked calculation endpoints
- `route.vaults.custom.go`: Specialized endpoints for integration with Rotki and other platforms
- `route.harvests.go`: Endpoints for retrieving harvest event data
- `route.strategies.one.go` and `route.strategies.all.go`: Strategy-related endpoints

### Utilities

- `utils.go`: Helper functions for vault filtering and identification
- `controller_test.go` and `getVaults_test.go`: Test files for API endpoints and vault retrieval

## Key Data Models

### Vault Representations

The package provides multiple vault representations to serve different use cases:

1. **`TExternalVault`**: The comprehensive vault model with complete details including:

    - Basic identifiers (address, name, symbol)
    - Token information (underlying asset details)
    - APR/APY data with historical performance
    - TVL metrics and price information
    - Associated strategies
    - Migration status
    - Risk assessment data

2. **`TSimplifiedExternalVault`**: A lightweight representation optimized for frontend display with:
    - Essential identifiers and metadata
    - Minimal token information
    - Current APR and TVL
    - Basic risk information

### Strategy Representations

- **`TStrategy`**: Detailed information about vault strategies
- **`TExternalStrategyDetails`**: External representation of strategy data

### Performance Metrics

- **`TExternalVaultAPR`**: APR/APY data with historical metrics
- **`TEarned`**: User-specific earnings data with realized and unrealized gains

## API Endpoints

### Primary Vault Endpoints

- `GET /vaults`: Retrieve all vaults with filtering options
- `GET /vaults/:address`: Get details for a specific vault
- `GET /vaults/simplified/:address`: Get simplified representation of a specific vault
- `GET /vaults/v2`: Filter vaults to V2 architecture only
- `GET /vaults/v3`: Filter vaults to V3 architecture only
- `GET /vaults/yearn`: Get only official Yearn vaults
- `GET /vaults/retired`: Get only retired vaults

### Specialized Endpoints

- `GET /vaults/blacklisted`: Retrieve blacklisted vaults
- `GET /vaults/tvl`: Get TVL information across all vaults
- `GET /chains/:chainID/vaults/earned/:address`: Calculate user earnings across all vaults on a chain
- `GET /chains/:chainID/vaults/earned/:address/:vaults`: Calculate user earnings for specific vaults
- `GET /vaults/earned/:address`: Calculate user earnings across all chains and vaults
- `GET /chains/:chainID/vaults/harvests/:address`: Get harvest events for a specific vault
- `GET /chains/:chainID/strategies`: Get all strategies for a chain
- `GET /chains/:chainID/strategies/:address`: Get details for a specific strategy

### Legacy Format Endpoints

- `GET /vaults/legacy/yearn`: Get all Yearn vaults in legacy format
- `GET /vaults/legacy/v2/yearn`: Get V2 Yearn vaults in legacy format
- `GET /vaults/legacy/v3/yearn`: Get V3 Yearn vaults in legacy format
- `GET /vaults/legacy/retired`: Get retired vaults in legacy format
- `GET /vaults/legacy/juiced`: Get juiced vaults in legacy format
- `GET /vaults/legacy/gimme`: Get Gimme-branded vaults in legacy format

### Custom Integration Endpoints

- `GET /vaults/custom/rotki`: Get vaults in Rotki-compatible format
- `GET /vaults/custom/rotki/count`: Get vault count for Rotki integration

## Query Parameters

Most vault endpoints support these common query parameters:

- `orderBy`: Field to sort results by (default: `featuringScore`)
- `orderDirection`: Sort direction, `asc` or `desc` (default: `asc`)
- `strategiesCondition`: Filter for strategies, values: `inQueue`, `debtLimit`, `debtRatio`, `absolute`, `all` (default: `debtRatio`)
- `hideAlways`: Whether to hide certain vaults (default: `false`)
- `migrable`: Filter for migrable vaults, values: `none`, `all`, `nodust`, `ignore` (default: `none`)
- `page` & `limit`: For pagination control (default: page 1, limit 200)
- `chainIDs`: Comma-separated list of chain IDs to include (defaults to all supported chains)

## Usage Examples

### Retrieving All Vaults

```go
// Initialize the controller
controller := vaults.Controller{}

// Use in a Gin route
router.GET("/vaults", func(c *gin.Context) {
    vaults := controller.GetAll(c)
    c.JSON(http.StatusOK, vaults)
})
```

### Retrieving User Earnings

```go
// Initialize the controller
controller := vaults.Controller{}

// Use in a Gin route to get earnings for a specific user
router.GET("/chains/:chainID/vaults/earned/:address", func(c *gin.Context) {
    controller.GetEarnedPerUser(c)
})
```

### Fetching Vault TVL

```go
// Initialize the controller
controller := vaults.Controller{}

// Use in a Gin route to get TVL information
router.GET("/vaults/tvl", func(c *gin.Context) {
    controller.GetVaultsTVL(c)
})
```

## Best Practices

1. **Use Simplified Models When Possible**: Use `TSimplifiedExternalVault` when full details aren't needed to reduce payload size.

2. **Apply Proper Filtering**: Leverage query parameters to filter results on the server side rather than client side.

3. **Consider Version-Specific Endpoints**: Use `GetV2` or `GetV3` endpoints when targeting specific vault architectures.

4. **Cache Responses When Appropriate**: Many vault endpoints return substantial data that doesn't change frequently.

5. **Handle Pagination**: For endpoints that return large datasets, implement proper pagination using `page` and `limit` parameters.

6. **Chain ID Handling**: Always validate chain IDs and provide appropriate fallbacks when querying multi-chain data.

7. **Error Handling**: All endpoints have standardized error responses that should be handled appropriately.

## GraphQL Integration

Several endpoints (particularly for harvests and earnings) utilize GraphQL for historical data retrieval. These queries interact with Yearn's subgraphs to fetch time-series data like:

- Harvest events
- User deposit/withdrawal history for FIFO calculations
- Historical price per share values

## FIFO Calculation Methodology

The earnings endpoints implement a First-In-First-Out (FIFO) methodology to calculate user gains:

1. All user deposits and withdrawals are retrieved from the subgraph
2. Deposits are processed in chronological order and paired with withdrawals
3. Price per share differences between deposit and withdrawal times are used to calculate realized gains
4. Remaining active deposits are evaluated against current price per share for unrealized gains
5. Results include both token-denominated and USD-equivalent values
