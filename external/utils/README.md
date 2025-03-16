# Utils Package

## Overview

The `utils` package provides utility API endpoints for the yDaemon service that offer information about the system configuration, supported chains, and other helper functionality. These endpoints serve as supporting utilities for client applications that need to discover system capabilities or retrieve configuration information without diving into specific domain contexts like tokens or pricing.

## API Endpoints

### Get Supported Chains

```
GET /utils/supported-chains
```

Returns an array of blockchain network IDs that are supported by the yDaemon system. This endpoint is useful for clients that need to discover which chains are available for data retrieval.

#### Response Format

```json
[
	1, // Ethereum Mainnet
	10, // Optimism
	42161, // Arbitrum
	100, // Gnosis
	137, // Polygon
	250, // Fantom
	8453 // Base
]
```

#### Example Usage

```bash
curl -X GET "https://api.example.com/utils/supported-chains"
```

## Integration with Other Packages

The utils package depends on:

- **env**: For accessing system configuration such as supported chains
- **gin**: For HTTP routing and response handling

It serves as a gateway for clients to discover system capabilities that are applicable across all other packages.

## Best Practices

1. **Cache Response Data**: The supported chains list rarely changes, so client-side caching is recommended
2. **Use for Discovery**: Integrate this endpoint into application boot sequences to discover available chains
3. **Error Handling**: Although errors are unlikely for these endpoints, clients should still implement appropriate error handling

## Example Integration

This example shows how a client application might use the supported chains endpoint to dynamically build UI elements:

```javascript
// Fetch the supported chains
fetch('https://api.example.com/utils/supported-chains')
	.then(response => response.json())
	.then(chains => {
		// Build a chain selector in the UI
		const selector = document.getElementById('chain-selector');

		// Chain ID to name mapping
		const chainNames = {
			1: 'Ethereum',
			10: 'Optimism',
			42161: 'Arbitrum',
			100: 'Gnosis',
			137: 'Polygon',
			250: 'Fantom',
			8453: 'Base'
		};

		// Create options for each supported chain
		chains.forEach(chainID => {
			const option = document.createElement('option');
			option.value = chainID;
			option.textContent = chainNames[chainID] || `Chain ${chainID}`;
			selector.appendChild(option);
		});
	})
	.catch(error => console.error('Failed to fetch supported chains:', error));
```

## Response Formats

All endpoints in the utils package return JSON data with appropriate HTTP status codes:

| Endpoint                    | Success Status | Error Status | Response Type     |
| --------------------------- | -------------- | ------------ | ----------------- |
| GET /utils/supported-chains | 200 OK         | N/A          | Array of integers |

## Future Expansion

The utils package is designed to be expanded with additional utility endpoints as needed, such as:

- System version information
- API rate limits and quotas
- Health check endpoints
- Documentation links

These endpoints would follow the same pattern of providing system-wide information that doesn't fit into other domain-specific packages.
