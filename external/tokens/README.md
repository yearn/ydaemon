# Tokens Package

## Overview

The `tokens` package provides API endpoints for retrieving token information across multiple blockchain networks. It offers endpoints to:

1. Fetch all tokens across all supported chains
2. Fetch all tokens for a specific chain
3. Fetch a specific token by address on a specific chain

These endpoints are essential for clients that need accurate token metadata such as symbols, names, decimals, and other descriptive information.

## API Endpoints

### 1. Get All Tokens (All Chains)

```
GET /tokens
```

Returns a nested map of all tokens across all supported blockchain networks.

#### Response Format

```json
{
	"1": {
		// Ethereum Mainnet
		"0x6b175474e89094c44da98b954eedeac495271d0f": {
			"address": "0x6b175474e89094c44da98b954eedeac495271d0f",
			"name": "Dai Stablecoin",
			"symbol": "DAI",
			"decimals": 18,
			"isVault": false,
			"display_name": "Dai",
			"display_symbol": "DAI",
			"description": "Dai is a stablecoin cryptocurrency which aims to keep its value as close to one United States dollar as possible through an automated system of smart contracts on the Ethereum blockchain",
			"category": "Stablecoin"
		}
		// Other tokens on Ethereum...
	},
	"42161": {
		// Arbitrum
		// Tokens on Arbitrum...
	}
	// Other chains...
}
```

### 2. Get All Tokens (For a Specific Chain)

```
GET /tokens/:chainID
```

Returns all tokens on the specified blockchain network.

#### Parameters

- `chainID`: The ID of the blockchain network (e.g., 1 for Ethereum Mainnet, 42161 for Arbitrum)

#### Response Format

```json
{
	"0x6b175474e89094c44da98b954eedeac495271d0f": {
		"address": "0x6b175474e89094c44da98b954eedeac495271d0f",
		"name": "Dai Stablecoin",
		"symbol": "DAI",
		"decimals": 18,
		"isVault": false,
		"display_name": "Dai",
		"display_symbol": "DAI",
		"description": "Dai is a stablecoin cryptocurrency which aims to keep its value as close to one United States dollar as possible through an automated system of smart contracts on the Ethereum blockchain",
		"category": "Stablecoin"
	}
	// Other tokens...
}
```

### 3. Get Specific Token

```
GET /tokens/:chainID/:address
```

Returns detailed information about a specific token on a specific blockchain network.

#### Parameters

- `chainID`: The ID of the blockchain network (e.g., 1 for Ethereum Mainnet, 42161 for Arbitrum)
- `address`: The Ethereum address of the token (hexadecimal format)

#### Response Format

```json
{
	"address": "0x6b175474e89094c44da98b954eedeac495271d0f",
	"name": "Dai Stablecoin",
	"symbol": "DAI",
	"decimals": 18,
	"isVault": false,
	"display_name": "Dai",
	"display_symbol": "DAI",
	"description": "Dai is a stablecoin cryptocurrency which aims to keep its value as close to one United States dollar as possible through an automated system of smart contracts on the Ethereum blockchain",
	"category": "Stablecoin",
	"underlyingTokens": []
}
```

## Error Responses

| Status Code | Description                                           |
| ----------- | ----------------------------------------------------- |
| 400         | Bad Request - Invalid chainID or address format       |
| 404         | Not Found - Token not found                           |
| 500         | Internal Server Error - Failed to retrieve token data |

## Response Structure

The token response structure contains the following fields:

| Field            | Type    | Description                                              |
| ---------------- | ------- | -------------------------------------------------------- |
| address          | string  | The Ethereum address of the token                        |
| name             | string  | The full name of the token                               |
| symbol           | string  | The token symbol (ticker)                                |
| decimals         | number  | The number of decimal places the token uses              |
| isVault          | boolean | Whether the token is a Yearn vault                       |
| display_name     | string  | Human-friendly name (may differ from official name)      |
| display_symbol   | string  | Human-friendly symbol (may differ from official symbol)  |
| description      | string  | A brief description of the token (if available)          |
| category         | string  | The token category (e.g., "Stablecoin", "Governance")    |
| underlyingTokens | array   | For complex tokens, a list of underlying token addresses |

## Integration with Other Packages

The tokens package depends on:

- **storage**: For retrieving token data from the persistent storage layer
- **helpers**: For validation and utility functions
- **models**: For token data structures

## Usage Examples

### Fetching all tokens on Ethereum

```bash
curl -X GET "https://api.example.com/tokens/1"
```

### Fetching a specific token (DAI on Ethereum)

```bash
curl -X GET "https://api.example.com/tokens/1/0x6b175474e89094c44da98b954eedeac495271d0f"
```

## Best Practices

1. **Cache Response Data**: Token metadata rarely changes, so client-side caching is recommended
2. **Bulk Requests**: When needing multiple tokens, use the chain-wide endpoint rather than multiple individual requests
3. **Input Validation**: Always validate chain IDs and token addresses before making requests
4. **Error Handling**: Be prepared to handle 404 responses for addresses that aren't recognized tokens
