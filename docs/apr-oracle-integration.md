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
