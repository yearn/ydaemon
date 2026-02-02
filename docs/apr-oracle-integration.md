# APR Oracle Integration (Kong)

As of version `yDaemon v3.0.0+`, the integration with the on-chain APR Oracle contract for V3 vaults has been replaced with a direct integration with Kong's GraphQL API.

## Overview

Previously, `yDaemon` calculated forward-looking APRs for V3 vaults by making RPC calls to a specific Oracle smart contract. This dependency has been removed in favor of fetching pre-calculated oracle data directly from Yearn's data aggregator, **Kong**.

## Data Source

The `performance.oracle` field from Kong's GraphQL API is now the single source of truth for:
-   **Oracle APR**: The annualized rate based on recent performance.
-   **Oracle APY**: The annualized yield (compounded).

## Implementation Details

-   **Internal Logic**: The `forward.v3.go` process queries the internal storage (populated by the Kong indexer) instead of computing values on-chain.
-   **Strategies**: Strategy-level APRs (`NetAPR`) for V3 also utilize this data source via `GetCurrentStrategyAPR` or the vault-level fallback, ensuring consistency.
-   **Fallback**: If Kong data is unavailable or null, the system gracefully defaults to zero values, preventing API crashes.

## API Response

The external API response structure for vaults remains unchanged to ensure backward compatibility. The `apr.forwardAPR` object is populated with the data retrieved from Kong.

```json
"forwardAPR": {
  "type": "v3:onchainOracle",
  "netAPR": 0.045, // Sourced from Kong
  "composite": {
    "v3OracleCurrentAPR": 0.045,
    "v3OracleStratRatioAPR": 0.045
  }
}
```

---

## V2 Vault Estimated APR Integration (Kong)

As of the implementation of [issue #560](https://github.com/yearn/ydaemon/issues/560), V2 vaults with Curve, Convex, Velodrome, and Aerodrome strategies now prioritize pre-calculated estimated APRs from Kong over local RPC-based calculations.

### Overview

Previously, `yDaemon` computed forward-looking APRs for V2 Curve-like vaults (Curve, Convex, Velodrome, Aerodrome) by making multiple on-chain RPC calls to calculate:
- Gauge boost values
- Pool APY from Curve subgraph
- Reward rates and token prices
- Various fee components

This approach was:
- **RPC-intensive**: Required multiple multicalls per vault
- **Computation-heavy**: Performed complex calculations locally
- **Slower**: Dependent on RPC response times

### New Data Source

The `performance.estimated` field from Kong's GraphQL API is now the **primary** source for V2 forward APR data:
- **Estimated APR**: Pre-computed annualized rate
- **Estimated APY**: Pre-computed annualized yield (compounded)
- **Type**: Strategy type identifier (e.g., `"crv"`, `"v2:velo"`)
- **Components**: Detailed breakdown including:
  - `boost`: Gauge boost multiplier
  - `poolAPY`: Pool swap fee APY
  - `boostedAPR`: Boosted reward APR
  - `baseAPR`: Base reward APR
  - `rewardsAPR`: Additional rewards APR
  - `rewardsAPY`: Additional rewards APY (compounded)
  - `cvxAPR`: Convex-specific APR
  - `keepCRV`: Percentage of CRV kept vs. swapped
  - `keepVelo`: Percentage of VELO kept vs. swapped

### Supported Chains

The estimated APR integration applies to V2 vaults on:
- **Ethereum (chainId: 1)**: Curve, Convex
- **Optimism (chainId: 10)**: Velodrome
- **Fantom (chainId: 250)**: Curve
- **Arbitrum (chainId: 42161)**: Curve
- **Base (chainId: 8453)**: Aerodrome

### Implementation Details

#### Priority System
1. **Primary**: Check for Kong `performance.estimated` data
2. **Fallback**: If Kong data unavailable, perform local RPC-based calculation
3. **Graceful degradation**: System continues to function even if Kong is temporarily unavailable

#### Internal Logic
- **Kong Indexer** (`internal/indexer/indexer.kong.go`): Fetches and stores estimated APR data from Kong GraphQL API
- **Storage Accessor** (`internal/storage/elem.vaults.go`): Provides `GetKongEstimatedAPY()` function to retrieve cached Kong data
- **APR Calculation** (`processes/apr/main.go`): Checks Kong data first before falling back to local computation
- **Helper Function** (`processes/apr/forward.curve.helpers.go`): `convertKongEstimatedAprToForwardAPY()` converts Kong format to internal `TForwardAPY` structure

#### Benefits
- **Reduced RPC Calls**: Eliminates multiple multicalls for gauge data, prices, and boost calculations
- **Faster Response**: Pre-computed data from Kong cache vs. real-time RPC queries
- **Consistency**: Kong serves as single source of truth for APR calculations
- **Maintained Components**: All APR component breakdowns preserved in API responses

### API Response

The external API response structure remains **fully backward compatible**. The `apr.forwardAPR` object structure is unchanged:

```json
"forwardAPR": {
  "type": "crv",                    // From Kong estimated.type
  "netAPY": 0.156,                  // From Kong estimated.apy
  "composite": {
    "boost": 2.5,                   // From Kong estimated.components.boost
    "poolAPY": 0.023,               // From Kong estimated.components.poolAPY
    "boostedAPR": 0.187,            // From Kong estimated.components.boostedAPR
    "baseAPR": 0.075,               // From Kong estimated.components.baseAPR
    "cvxAPR": 0.045,                // From Kong estimated.components.cvxAPR (Convex only)
    "rewardsAPY": 0.034,            // From Kong estimated.components.rewardsAPY
    "keepCRV": 0.1                  // From Kong estimated.components.keepCRV
  }
}
```

### Migration Notes

- **No Breaking Changes**: API consumers experience no differences in response structure
- **Automatic Fallback**: If Kong estimated data is unavailable, yDaemon seamlessly falls back to local computation
- **Strategy Support**: Works for all Curve, Convex, Velodrome, and Aerodrome V2 strategies

