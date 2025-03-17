# Prices Package

> The prices package provides API route handlers for accessing token price data across different blockchain networks. It serves as a critical component of yDaemon, offering consistent interfaces for accessing both raw and humanized price information for crypto assets through a RESTful API.

## Core Components

The package is organized into several route handlers, each designed for specific price retrieval scenarios:

### Price Retrieval by Scope

- **All Prices**: Fetch prices for all tokens across all chains or for a specific chain
- **Specific Tokens**: Retrieve prices for a selected subset of tokens
- **Single Token**: Get price data for an individual token on a specific chain
- **Detailed Price Information**: Access comprehensive price data including metadata

### Response Format Options

All endpoints support two response formats, controlled via the `humanized` query parameter:

- **Raw Prices**: Integer values representing token prices with full precision (default)
- **Humanized Prices**: Float values formatted for human readability

## API Endpoints

### Chain-Wide Price Endpoints

```
GET /prices/all
```

Returns prices for all tokens across all supported chains.

**Example Response (Raw):**

```json
{
	"1": {
		"0x6b175474e89094c44da98b954eedeac495271d0f": "1000000000000000000",
		"0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2": "1500000000000000000"
	},
	"42161": {
		"0xDA10009cBd5D07dd0CeCc66161FC93D7c9000da1": "1000000000000000000",
		"0x82aF49447D8a07e3bd95BD0d56f35241523fBab1": "1500000000000000000"
	}
}
```

**Example Response (Humanized):**

```json
{
	"1": {
		"0x6b175474e89094c44da98b954eedeac495271d0f": 1.0,
		"0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2": 1.5
	},
	"42161": {
		"0xDA10009cBd5D07dd0CeCc66161FC93D7c9000da1": 1.0,
		"0x82aF49447D8a07e3bd95BD0d56f35241523fBab1": 1.5
	}
}
```

```
GET /prices/:chainID
```

Returns all token prices for a specific blockchain network.

**Example Response (Raw):**

```json
{
	"0x6b175474e89094c44da98b954eedeac495271d0f": "1000000000000000000",
	"0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2": "1500000000000000000"
}
```

**Example Response (Humanized):**

```json
{
	"0x6b175474e89094c44da98b954eedeac495271d0f": 1.0,
	"0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2": 1.5
}
```

```
GET /prices/:chainID/all
```

Returns detailed price information for all tokens on a specific chain.

**Example Response:**

```json
{
	"0x6b175474e89094c44da98b954eedeac495271d0f": {
		"price": "1000000000000000000",
		"humanizedPrice": 1.0
	},
	"0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2": {
		"price": "1500000000000000000",
		"humanizedPrice": 1.5
	}
}
```

### Specific Token Endpoints

```
GET /prices/:chainID/:address
```

Returns the price for a single token on a specific chain.

**Example Response (Raw):**

```json
"1000000000000000000"
```

**Example Response (Humanized):**

```json
1.0
```

```
GET /prices/:chainID/some/:addresses
```

Returns prices for multiple tokens on the same chain. The `addresses` parameter should be a comma-separated list.

**Example Response (Raw):**

```json
{
	"0x6b175474e89094c44da98b954eedeac495271d0f": "1000000000000000000",
	"0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2": "1500000000000000000"
}
```

**Example Response (Humanized):**

```json
{
	"0x6b175474e89094c44da98b954eedeac495271d0f": 1.0,
	"0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2": 1.5
}
```

```
GET /prices/some/:addresses
```

Returns prices for multiple tokens across different chains. The `addresses` parameter should be a comma-separated list in the format `chainID:address`.

**Example Response (Raw):**

```json
{
	"1": {
		"0x6b175474e89094c44da98b954eedeac495271d0f": "1000000000000000000"
	},
	"42161": {
		"0xDA10009cBd5D07dd0CeCc66161FC93D7c9000da1": "1000000000000000000"
	}
}
```

**Example Response (Humanized):**

```json
{
	"1": {
		"0x6b175474e89094c44da98b954eedeac495271d0f": 1.0
	},
	"42161": {
		"0xDA10009cBd5D07dd0CeCc66161FC93D7c9000da1": 1.0
	}
}
```

```
POST /prices/some
```

Similar to the GET endpoint, but accepts a list of token addresses in the request body.

**Example Request:**

```json
{
	"addresses": "1:0x6b175474e89094c44da98b954eedeac495271d0f,42161:0xDA10009cBd5D07dd0CeCc66161FC93D7c9000da1"
}
```

**Example Response (Raw):**

```json
{
	"1": {
		"0x6b175474e89094c44da98b954eedeac495271d0f": "1000000000000000000"
	},
	"42161": {
		"0xDA10009cBd5D07dd0CeCc66161FC93D7c9000da1": "1000000000000000000"
	}
}
```

## Error Responses

All endpoints implement consistent error handling patterns:

### Invalid Chain ID

```json
{
	"error": "invalid chainID"
}
```

### Invalid Token Address

```json
{
	"error": "invalid address"
}
```

### Token Not Found

```json
{
	"error": "token not found"
}
```

### Invalid Request Body

```json
{
	"error": "invalid JSON: syntax error at line 1, character 5"
}
```

### Too Many Addresses Requested

```json
{
	"error": "Too many addresses requested",
	"limit": 100
}
```

## Usage Examples

### Fetching All Prices

```bash
# Get all prices across all chains in raw format
curl https://api.example.com/prices/all

# Get all prices on Ethereum mainnet in humanized format
curl https://api.example.com/prices/1?humanized=true
```

### Fetching Specific Token Prices

```bash
# Get DAI price on Ethereum mainnet
curl https://api.example.com/prices/1/0x6b175474e89094c44da98b954eedeac495271d0f

# Get prices for multiple tokens on Ethereum mainnet
curl https://api.example.com/prices/1/some/0x6b175474e89094c44da98b954eedeac495271d0f,0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2

# Get prices for tokens across different chains
curl https://api.example.com/prices/some/1:0x6b175474e89094c44da98b954eedeac495271d0f,10:0xDA10009cBd5D07dd0CeCc66161FC93D7c9000da1
```

### Using POST for Large Token Lists

```bash
curl -X POST https://api.example.com/prices/some \
  -H "Content-Type: application/json" \
  -d '{"addresses": "1:0x6b175474e89094c44da98b954eedeac495271d0f,10:0xDA10009cBd5D07dd0CeCc66161FC93D7c9000da1"}'
```

## Integration with Other Packages

This package primarily integrates with:

- **storage**: Retrieves price data from the internal storage system
- **bigNumber**: Handles precise numeric representations for token prices
- **helpers**: Provides utility functions for parameter validation and formatting

## Best Practices

1. **Use Humanized Format for UI Applications**: Set `humanized=true` when displaying prices to users
2. **Batch Requests for Multiple Tokens**: Use the `/some` endpoints instead of multiple single token requests
3. **Use POST for Large Token Lists**: When fetching prices for many tokens, use the POST endpoint
4. **Handle Missing Prices Gracefully**: Some tokens may not have price data available
5. **Consider Rate Limiting**: Implement appropriate rate limiting for high-volume requests
6. **Cache Responses**: For frequently accessed prices, implement client-side caching with appropriate TTL
7. **Handle Errors Properly**: Check for specific error status codes and messages
8. **Validate Addresses Client-Side**: Pre-validate addresses to avoid unnecessary API calls

## Improvements

### 1. Silent Failure for Invalid Addresses in Batch Requests

**File:** `route.prices.some.go`
**Issue:** When processing a batch of addresses, invalid ones are silently skipped without any indication to the client.

**Impact:** Clients may not realize that some of their requested prices were not included in the response.

**Recommendation:** Include metadata in the response to indicate which addresses were invalid or not found:

```go
type response struct {
    Prices map[string]*bigNumber.Int `json:"prices"`
    Errors map[string]string         `json:"errors,omitempty"`
}
```

### 2. Lack of Response Pagination

**Issue:** Endpoints that return all prices don't support pagination.

**Impact:** For chains with many tokens, the response size could be very large, leading to performance issues.

**Recommendation:** Add pagination support:

```go
// Query parameters: ?page=1&limit=100
page := helpers.SafeStringToUint64(getQuery(c, "page"), 1)
limit := helpers.SafeStringToUint64(getQuery(c, "limit"), 100)
// Implement pagination logic
```

### 3. Inconsistent Response Structure

**Issue:** Different endpoints return different structures (maps vs. arrays, nested vs. flat).

**Impact:** Makes client integration more complex.

**Recommendation:** Standardize on a consistent response format across all endpoints.

### 4. No Input Length Validation

**Issue:** No limits on the number of addresses that can be requested in a single call.

**Impact:** Could enable denial-of-service attacks by requesting too many prices at once.

**Recommendation:** Add reasonable limits to batch size:

```go
const MaxBatchSize = 100

addressesStr := strings.Split(c.Param("addresses"), ",")
if len(addressesStr) > MaxBatchSize {
    c.JSON(http.StatusBadRequest, gin.H{"error": "Too many addresses requested"})
    return
}
```

### 5. No Rate Limiting

**Issue:** No built-in rate limiting for API endpoints.

**Recommendation:** Implement rate limiting middleware to prevent abuse.
