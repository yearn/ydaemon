# The vault structure is composed in 4 parts:

#### The immutable elements

We are not expecting them to change. They are fetched every 24 hours.

- `address`: Address of the vault
- `assetAddress`: Address of the underlying token
- `managementnt`: Address of the management contract
- `governance`: Address of the governance contract
- `guardian`: Address of the guardian contract
- `rewards`: Address of the rewards contract
- `version`: The version of the vault
- `activation`: When the vault was activated
- `chainID`: The chainID of the vault
- `endorsed`: If the vault is endorsed by Yearn

#### The semi-mutable elements

We are expecting them to change in some rare situations, but we still want this to update more often than the immutable elements. They are fetched every 1 hour.

- `performanceFee`: The performance fee of the vault
- `managementFee`: The management fee of the vault
- `emergencyShutdown`: If the vault is in emergency shutdown

#### The mutable elements

We are expecting them to change often. They are fetched every 15 minutes.

- `lastActiveStrategies`: The list of currently active strategies of the vault
- `lastPricePerShare`: The current price per share of the vault
- `lastTotalAssets`: The current total assets of the vault
- `lastUpdate`: The last time the update was performed

#### The manual elements

For some specific use cases, we want to be able to manually update some elements. This is a manual process, not automated and a reboot of the bot is required to update them.

- `isRetired`: If the vault is retired
- `isHidden`: If we should hide the vaults from the standard results
- `migration`: Indicates the vault migration data and availability
- `classification`: Give some more details to be able to classify the vaults in some UIs.
