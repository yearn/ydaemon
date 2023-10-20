# Strategies
The Strategies JSON file contains all the information about the strategies. As other JSON files we use, it is composed of a `lastUpdate` field, a `version` field and a `shouldRefresh` field.
- `lastUpdate`: The last time the file was updated
- `version`: The version of the file (major, minor, patch).
	- `major`: When at least one strategy is removed since the last version
	- `minor`: When at least one strategy is added since the last version
	- `patch`: When at least one strategy is updated since the last version
- `shouldRefresh`: If the file should be refreshed or not. This is used by the bot to know if it should update the file or not, no matter if the other conditions are met.

Each strategies contains theses elements:

#### The immutable elements

We are not expecting them to change. They are fetched every 24 hours.

- `address`: Address of the strategy
- `vaultAddress`: Address of the vault
- `name`: Name of the strategy
- `vaultVersion`: The version of the vault
- `activation`: When the strategy was activated
- `chainID`: The chainID of the strategy
- `doHealthCheck`: If the strategy is doing health check
- `isActive`: If the strategy is active
- `isInQueue`: If the strategy is in the vault queue

#### The semi-mutable elements

We are expecting them to change in some rare situations, but we still want this to update more often than the immutable elements. They are fetched every 1 hour.

- `keepCRV`: The ratio of CRV to keep in the strategy
- `keepCRVPercent`: The ratio of CRV to keep in the strategy in percent
- `keepCVX`: The ratio of CVX to keep in the strategy

#### The mutable elements

We are expecting them to change often. They are fetched every 15 minutes.

- `lastTotalDebt`: The last known total debt
- `lastTotalLoss`: The last known total loss
- `lastTotalGain`: The last known total gain
- `lastPerformanceFee`: The last known performance fee
- `lastReport`: The last time the strategy was reported
- `lastDebtRatio`: The last known debt ratio. Only > 0.2.2
- `lastEstimatedTotalAssets`: The last known estimated total assets

#### The manual elements

For some specific use cases, we want to be able to manually update some elements. This is a manual process, not automated and a reboot of the bot is required to update them.

- `isRetired`: If the strategy is retired and will be ignored
- `displayName`: An override of the name to be displayed
- `description`: A description for this strategy
- `protocols`: The protocols used by this strategy

## Example

```json
"lastUpdate": "2023-10-20T14:43:22.410975+02:00",
"version": {
	"major": 0,
	"minor": 0,
	"patch": 0
},
"shouldRefresh": false,
"strategies": {
	"0x000007252ab8d9120005d30aa15567ea8de9a110": {
		"address": "0x000007252ab8d9120005d30aa15567ea8de9a110",
		"vaultAddress": "0x490bd0886f221a5f79713d3e84404355a9293c50",
		"name": "StrategyConvexibCHF",
		"vaultVersion": "0.4.3",
		"activation": 14139749,
		"chainID": 1,
		"doHealthCheck": true,
		"isActive": false,
		"isInQueue": false,
		"keepCRV": "1000",
		"keepCRVPercent": null,
		"keepCVX": null,
		"lastTotalDebt": "0",
		"lastTotalLoss": "0",
		"lastTotalGain": "118169170152051977696595",
		"lastPerformanceFee": "0",
		"lastReport": "1686670799",
		"lastDebtRatio": "0",
		"lastEstimatedTotalAssets": "0",
		"isRetired": true,
		"displayName": "",
		"description": "",
		"protocols": null
	}
}
```
