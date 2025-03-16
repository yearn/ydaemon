# Vaults Package

## Overview

The `vaults` package is a core component of yDaemon that provides API endpoints and data structures for interacting with Yearn Finance vaults across multiple blockchain networks. This package serves as the primary interface for clients to access comprehensive vault information, including vault details, APR/APY data, strategies, harvest data, TVL metrics, and user-specific earnings calculations.

The package follows consistent patterns for error handling, parameter validation, and response formatting to ensure a reliable API experience. It processes raw blockchain data into standardized formats suitable for frontend applications, data analytics, and integration with third-party services.

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

- `utils.go`: Helper functions for vault filtering, parameter validation, and error handling
- `validation.go`: Functions for validating inputs and performing version checks
- `test_helpers.go`: Testing utilities and mock functions

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

- **`TStrategy`**: Detailed information about vault strategies, including:
  - Address and name
  - Description and purpose
  - Performance data (total debt, gains, losses)
  - Historical metrics (last report timestamp, performance fee)
  - Operational status (debt ratio, queue position)

- **`TExternalStrategyDetails`**: Financial and operational metrics for a strategy:
  - Total debt allocated from the vault
  - Cumulative gains and losses
  - Performance fee percentage
  - Timestamp of last report
  - Debt ratio allocation (for v0.2.2+ vaults)

### Performance Metrics

- **`TExternalVaultAPR`**: Comprehensive yield data including:
  - Net APR (primary yield metric)
  - Fee information (management, performance, etc.)
  - Historical performance points (day, week, month ago)
  - Forward-looking projections (expected future yields)
  - Component yields (base, boosted, rewards)
  - Special rewards (staking, Gamma, etc.)

- **`TExternalVaultHarvest`**: Strategy harvest event data:
  - Transaction details (hash, timestamp)
  - Profit and loss amounts (both raw and USD value)
  - Calculated historical APR based on harvest
  - Associated vault and strategy information

- **`TEarned`**: User-specific earnings calculated with FIFO methodology:
  - Realized gains (from completed deposit/withdrawal pairs)
  - Unrealized gains (from active deposits based on current share price)
  - Token amount and USD value of earnings
  - Deposit timestamps and duration information

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

- `GET /vaults/blacklisted`: Retrieve blacklisted vaults (vaults excluded from standard listings due to deprecation, security concerns, or other issues)
  - Optional `chainID` parameter to filter by chain
  - Returns vaults with detailed metadata explaining why they're blacklisted

- `GET /vaults/tvl`: Get Total Value Locked information across all vaults
  - Returns aggregate TVL plus breakdown by chain
  - Includes vault count and USD value per chain
  - Optional `chainIDs` parameter to filter by specific chains

- `GET /vaults/:chainID/:addresses`: Retrieve specific vaults by their addresses
  - Comma-separated list of vault addresses
  - Returns full vault details for the specified addresses
  - Supports all standard query parameters for sorting and filtering

- `GET /chains/:chainID/vaults/harvests/:address`: Get harvest events for a specific vault
  - Returns chronological list of strategy harvests
  - Parameters:
    - `limit`: Maximum number of harvests to return (default: 20)
    - `startTimestamp`: Filter harvests after this timestamp
    - `endTimestamp`: Filter harvests before this timestamp
    - `orderBy`: Sort field, options: `timestamp`, `profit`, `loss` (default: `timestamp`)
    - `orderDirection`: Sort order, `asc` or `desc` (default: `desc`)

- `GET /chains/:chainID/vaults/earned/:address`: Calculate user earnings across all vaults on a chain
- `GET /chains/:chainID/vaults/earned/:address/:vaults`: Calculate user earnings for specific vaults
- `GET /vaults/earned/:address`: Calculate user earnings across all chains and vaults
  - All earned endpoints use the same FIFO calculation methodology
  - Return both realized and unrealized gains
  - Include original deposit timestamps and durations

- `GET /chains/:chainID/strategies`: Get all strategies for a chain
  - Parameters:
    - `strategiesCondition`: Filter by strategy status (default: `debtRatio`)
    - `orderBy`: Sort field, options: `address`, `tvl`, `apy` (default: `address`)
    - `orderDirection`: Sort order, `asc` or `desc` (default: `asc`)

- `GET /chains/:chainID/strategies/:address`: Get details for a specific strategy
  - Returns comprehensive details including all financial metrics
  - Contains both current and historical performance data

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

### Sorting Parameters
- `orderBy`: Field to sort results by (default: `featuringScore`)
  - Common values: `featuringScore`, `name`, `symbol`, `tvl`, `apyTypeOverall`
  - Strategy endpoints: `address`, `totalDebt`, `debtRatio`
  - Harvest endpoints: `timestamp`, `profit`, `loss`
- `orderDirection`: Sort direction, `asc` or `desc` (default: `asc`)

### Filtering Parameters
- `strategiesCondition`: Filter for strategies, with the following values:
  - `debtRatio`: Include strategies with non-zero debt ratio allocation (default)
  - `inQueue`: Include only strategies in the vault's withdrawal queue
  - `debtLimit`: Include strategies with a non-zero debt limit
  - `absolute`: Include only strategies with actual debt > 0
  - `all`: Include all strategies regardless of status
- `hideAlways`: Whether to hide retired and hidden vaults (default: `false`)
- `migrable`: Filter for vaults with migration options:
  - `none`: Exclude migrable vaults if they are hidden/retired (default)
  - `all`: Include all vaults with migration available
  - `nodust`: Include only migrable vaults with TVL > $100
  - `ignore`: Exclude all vaults with migration available

### Pagination Parameters
- `page`: Page number starting from 1 (default: 1)
- `limit`: Number of results per page (default: 200, max: 1000)

### Chain Selection
- `chainIDs`: Comma-separated list of chain IDs to include 
  - Examples: `1` (Ethereum), `10` (Optimism), `137` (Polygon), etc.
  - Defaults to all supported chains if not specified

### Time Range Parameters (Harvest Endpoints)
- `startTimestamp`: Filter results after this Unix timestamp
- `endTimestamp`: Filter results before this Unix timestamp

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
   - The simplified model is typically 50-70% smaller than the full model
   - Most listing views can use the simplified model effectively

2. **Apply Proper Filtering**: Leverage query parameters to filter results on the server side rather than client side.
   - Use `strategiesCondition` to limit strategy data transfer
   - Use `hideAlways=true` to exclude retired and hidden vaults
   - Combine `chainIDs` with specific filter endpoints for targeted results

3. **Consider Version-Specific Endpoints**: Use `GetV2` or `GetV3` endpoints when targeting specific vault architectures.
   - V3 vaults have different fields and capabilities than V2
   - Version-specific endpoints ensure consistent data structure

4. **Cache Responses When Appropriate**: Many vault endpoints return substantial data that doesn't change frequently.
   - TVL and APR data typically update every 15-30 minutes
   - Full vault details rarely change except for performance metrics
   - Strategy allocation data changes infrequently (typically on harvests)

5. **Handle Pagination**: For endpoints that return large datasets, implement proper pagination using `page` and `limit` parameters.
   - Large chains may have 50+ vaults and 200+ strategies
   - Calculate total pages using the data length and limit
   - Always use pagination for strategies and harvests endpoints

6. **Chain ID Handling**: Always validate chain IDs and provide appropriate fallbacks when querying multi-chain data.
   - Check for supported chains using the environment variables
   - Handle scenarios where a chain might be temporarily unavailable
   - Provide clear error messages for unsupported chain IDs

7. **Error Handling**: All endpoints follow a consistent error response format that should be handled appropriately.
   - HTTP status codes indicate the category of error
   - Response bodies contain detailed error messages
   - JSON responses include both user-friendly messages and technical details

8. **Performance Considerations**:
   - For user interfaces showing multiple vaults, use the `/vaults/:chainID/:addresses` endpoint
   - When fetching harvest data, use appropriate time filters to limit result size
   - For dashboard views, consider using the TVL endpoint for aggregate metrics

## GraphQL Integration

Several endpoints (particularly for harvests and earnings) utilize GraphQL for historical data retrieval. These queries interact with Yearn's subgraphs to fetch time-series data like:

- Harvest events
- User deposit/withdrawal history for FIFO calculations
- Historical price per share values

## FIFO Calculation Methodology

The earnings endpoints implement a First-In-First-Out (FIFO) methodology to calculate user gains, which provides the most accurate representation of realized and unrealized yields over time:

1. **Data Collection**: All user deposits and withdrawals are retrieved from the subgraph
   - Transaction timestamps, amounts, and share prices are recorded
   - Share prices are retrieved for each transaction time

2. **Chronological Processing**: Deposits are processed in order of occurrence
   - Each deposit creates a "position" with a specific share price and amount
   - When a withdrawal occurs, it's matched against the oldest available deposit

3. **Realized Gain Calculation**:
   - When a position is partially or fully closed (withdrawal matched against deposit):
   - `realized_gain = (withdrawal_share_price - deposit_share_price) * withdrawn_shares`
   - This represents actual profit/loss for completed transactions

4. **Unrealized Gain Calculation**:
   - For positions that remain open (deposits without matching withdrawals):
   - `unrealized_gain = (current_share_price - deposit_share_price) * remaining_shares`
   - This represents potential profit/loss based on current market conditions

5. **Result Compilation**:
   - Results include both token amounts and USD-equivalent values
   - Duration information shows how long each position has been active
   - Historical deposit information is preserved for audit purposes

This methodology ensures accurate accounting of gains across multiple deposits and withdrawals over time, and properly distinguishes between completed transactions and open positions.
