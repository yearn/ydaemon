# Environment Package

/**********************\*\*\*\***********************\*\***********************\*\*\*\***********************
** The env package is a fundamental component of yDaemon that manages all environment-specific
** configurations, particularly related to blockchain networks. It serves as the central
** repository for chain-specific settings, contract addresses, supported features, and external
** service endpoints. This package establishes the foundation for blockchain interactions across
** the application.
**
** This documentation serves as a comprehensive guide for developers working with the env package.
********************\*\*\*\***********************\*\*\*\***********************\*\*\*\***********************/

## Core Components

### Chain Configurations

The package defines detailed configurations for all supported blockchain networks:

- **Chain Definitions**: Separate files for each supported chain (Ethereum, Optimism, Polygon, etc.)
- **Network Parameters**: RPC endpoints, Etherscan APIs, block ranges, and other network-specific settings
- **Contract Addresses**: Addresses for key contracts like multicall, lens, registries, and more
- **Performance Metrics**: Values like average blocks per day to estimate historical blocks
- **Feature Support**: Flags for features like WebSocket support and specialized endpoints

```go
// Access configuration for a specific chain
chain, ok := env.GetChain(1) // 1 is Ethereum's chain ID
if ok {
    rpcEndpoint := chain.RpcURI
    // Use other chain properties...
}
```

### Core Data Structures

The package defines key data structures for blockchain interactions:

- **TChain**: The primary structure containing all chain-specific configuration
- **TContractData**: Structure for contract addresses with deployment blocks and metadata
- **TChainCurve**: Curve-specific configuration for each chain
- **TExtraStakingContracts**: Information about additional staking contracts
- **TChainExtraURI**: URLs for external services like Gamma Merkl, Hypervisor, and Pendle

### Environment Variable Management

Handles loading and applying environment variables to override default configurations:

- **Environment Loading**: Automatic loading from .env files
- **RPC Configuration**: Override RPC endpoints through environment variables
- **Dynamic Configuration**: Updates chain configurations based on environment settings

```go
// Environment variables are formatted as:
// RPC_URI_FOR_1=https://custom-ethereum-rpc.com (for chain ID 1)
// SUBGRAPGH_FOR_1=https://api.thegraph.com/... (for Ethereum subgraph)
```

### Registry Classification

Functions to identify and categorize different types of Yearn protocol registries:

- **Core Registries**: Standard Yearn vaults displayed on Yearn.fi
- **Specialized Registries**: Registries for projects like Juiced, PoolTogether, etc.
- **ERC4626 Registries**: Public ERC4626-compatible vault registries
- **Registry Status**: Functions to check if registries are disabled or deprecated

```go
// Check if a registry belongs to the core Yearn protocol
isCore := env.IsRegistryFromYearnCore(chainID, registryAddress)

// Check if a registry is for specialized projects
isJuiced := env.IsRegistryFromJuiced(chainID, registryAddress)
```

### External Service Integration

Manages configuration for external API services:

- **Asset URLs**: Endpoints for token icons and metadata
- **Price Feeds**: URLs for CoinGecko and DeFiLlama price APIs
- **API Keys**: Management of CoinGecko API keys for rate limit distribution

## Initialization Process

The package follows a specific initialization sequence:

1. Load default settings for all supported chains
2. Override with environment variables from .env files
3. Register all chains in the global CHAINS map
4. Populate the list of supported chain IDs
5. Apply any network-specific configurations

This happens automatically at application startup through Go's init() function.

## How to Add a New Chain

Adding support for a new blockchain network to yDaemon requires several steps to ensure proper integration. Follow this guide to add a new chain:

### 1. Create a New Chain Configuration File

Create a new file named `chain.[networkname].go` in the `common/env` directory. Use the following template:

```go
package env

import (
	"math"

	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/internal/models"
)

var NETWORKNAME = TChain{
	ID:              [CHAIN_ID],  // The network's unique chain ID
	RpcURI:          `[DEFAULT_RPC_ENDPOINT]`, // Public RPC endpoint
	EtherscanURI:    `[EXPLORER_API_URL]`, // Chain explorer API URL
	MaxBlockRange:   [BLOCK_RANGE],  // Maximum block range for queries (e.g., 10000)
	MaxBatchSize:    math.MaxInt64,  // Or appropriate batch size limit
	AvgBlocksPerDay: [BLOCKS_PER_DAY], // Average blocks produced per day
	CanUseWebsocket: [true/false],    // Whether WebSocket connections are supported

	// Multicall contract - required for efficient blockchain queries
	MulticallContract: TContractData{
		Address: common.HexToAddress(`[MULTICALL_ADDRESS]`),
		Block:   [DEPLOYMENT_BLOCK], // Block where the contract was deployed
	},

	// Native coin configuration
	Coin: models.TERC20Token{
		Address:                   DEFAULT_COIN_ADDRESS,
		UnderlyingTokensAddresses: []common.Address{},
		Type:                      models.TokenTypeNative,
		Symbol:                    `[NATIVE_SYMBOL]`, // E.g., ETH, MATIC, etc.
		Name:                      `[NATIVE_NAME]`,   // E.g., Ethereum, Polygon, etc.
		Decimals:                  18, // Usually 18, adjust if different
		Icon:                      `[COIN_ICON_URL]`, // Optional
	},

	// Registry contracts for Yearn vaults
	Registries: []TContractData{
		{
			Address: common.HexToAddress(`[REGISTRY_ADDRESS]`),
			Block:   [DEPLOYMENT_BLOCK],
			Label:   "YEARN", // Use YEARN for core registries
		},
		// Add more registries if needed
	},

	// Optional: Curve configuration if the chain has Curve protocol
	Curve: TChainCurve{
		RegistryAddress: common.HexToAddress(`[CURVE_REGISTRY]`),
		FactoryAddress:  common.HexToAddress(`[CURVE_FACTORY]`),
		PoolsURIs:       []string{`[CURVE_POOLS_API]`},
		GaugesURI:       `[CURVE_GAUGES_API]`,
	},
}
```

### 2. Gather Required Information

Before implementing the chain configuration, gather all necessary information:

- **Chain ID**: The official chain ID number (e.g., 1 for Ethereum, 137 for Polygon)
- **RPC Endpoint**: A reliable public RPC endpoint for the chain
- **Explorer API**: URL for the chain's explorer API (Etherscan-compatible)
- **Block Parameters**: Average blocks per day, max block range for queries
- **Contract Addresses**: Multicall, registries, and other required contracts
- **Native Token**: Symbol, name, and decimals of the native currency

### 3. Research Contract Deployments

For Yearn to work properly, identify deployment addresses and blocks for:

- **Multicall Contract**: Essential for batched calls (often at `0x5ba1e12693dc8f9c48aad8770482f4739beed696`)
- **Yearn Registry**: If Yearn vaults exist on the chain
- **Lens Contract**: If available, for vault data aggregation
- **Partner Contracts**: If any partnerships exist on this chain

### 4. Update the Initialization Code

Modify the `init()` function in `environment.go` to include your new chain:

```go
func init() {
	godotenv.Load(`.env`)
	// Set subgraph URIs from environment variables if needed
	NETWORKNAME.SubgraphURI = os.Getenv("SUBGRAPGH_FOR_[CHAIN_ID]")

	// Add to CHAINS map
	CHAINS[1] = ETHEREUM
	// ... existing chains ...
	CHAINS[[CHAIN_ID]] = NETWORKNAME  // Add your new chain here

	SetEnv()

	// SUPPORTED_CHAIN_IDS will be populated automatically
}
```

### 5. Test the Integration

After adding the new chain:

1. Create environment variables for custom RPC endpoints:

    ```
    RPC_URI_FOR_[CHAIN_ID]=https://your-custom-rpc-endpoint.com
    ```

2. Test basic connectivity to ensure the RPC endpoint works:

    ```go
    chain, ok := env.GetChain([CHAIN_ID])
    if !ok {
        // Chain not found
    }
    client := ethereum.GetRPC([CHAIN_ID])
    blockNumber, err := client.BlockNumber(context.Background())
    ```

3. Verify registry scanning works if applicable

### 6. Chain-Specific Customizations

Some chains may require additional customizations:

- **DefiLlama Integration**: Update the `chainIDToName` function in `blocktime.go` to map your chain ID to the name used by DeFiLlama
- **Gas Price Strategies**: If the chain has unique gas pricing, configure appropriate settings
- **Rate Limiting**: Consider API rate limits for the chain's explorer and adjust accordingly

### 7. Documentation Updates

Update relevant documentation to include the new chain:

- Add to chain list in appropriate README files
- Document any chain-specific quirks or requirements
- Update environment variable examples to include the new chain

## Best Practices

- **Chain Support Checking**: Always check that a chain is supported before operations
- **Environment Variables**: Use environment variables for deployment-specific configurations
- **Registry Classification**: Use the provided functions to categorize registries correctly
- **Contract Addresses**: Reference contract addresses from this package rather than hardcoding

## Usage Examples

### Working with Chains

```go
// Get all supported chains
chains := env.GetChains()

// Loop through all supported chain IDs
for _, chainID := range env.SUPPORTED_CHAIN_IDS {
    chain, _ := env.GetChain(chainID)
    // Use chain configuration...
}
```

### Registry Classification

```go
// Determine what type of registry we're working with
if env.IsRegistryFromYearnCore(chainID, registryAddress) {
    // Handle core Yearn vaults
} else if env.IsRegistryFromJuiced(chainID, registryAddress) {
    // Handle Juiced vaults
} else if env.IsRegistryFromPublicERC4626(chainID, registryAddress) {
    // Handle public ERC4626 vaults
}

// Skip disabled registries
if env.IsRegistryDisabled(chainID, registryAddress) {
    // Skip this registry
}
```

### Working with Contract Addresses

```go
// Get the multicall contract address for a specific chain
chain, ok := env.GetChain(chainID)
if ok {
    multicallAddress := chain.MulticallContract.Address
    multicallDeployBlock := chain.MulticallContract.Block
    // Use contract details...
}
```

## Package Organization

- **chains.go**: Defines core data structures and accessor functions
- **constants.go**: Global constants used throughout the application
- **environment.go**: Environment variable handling and initialization
- **registryAssignation.go**: Registry classification functions
- **chain.[network].go**: Network-specific configurations
    - chain.ethereum.go
    - chain.optimism.go
    - chain.polygon.go
    - chain.arbitrum.go
    - chain.base.go
    - chain.fantom.go
    - chain.gnosis.go
