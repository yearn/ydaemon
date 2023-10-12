# The Strategies structure is composed in 5 parts:

#### The immutable elements

We are not expecting them to change. They are fetched every 24 hours.

- `address`: Address of the strategy
- `vaultAddress`: Address of the vault
- `keeperAddress`: Address of the keeper
- `strategistAddress`: Address of the strategist
- `rewardsAddress`: Address of the rewards
- `healthCheckAddress`: Address of the health check contract
- `name`: Name of the strategy
- `vaultVersion`: The version of the vault
- `activation`: When the strategy was activated
- `chainID`: The chainID of the strategy
- `endorsed`: If the vault is endorsed by Yearn
- `doHealthCheck`: If the strategy is doing health check
- `isActive`: If the strategy is active
- `isInQueue`: If the strategy is in the vault queue
- `withdrawalQueuePosition`: The position of the strategy in the withdrawal queue

#### The semi-mutable elements

We are expecting them to change in some rare situations, but we still want this to update more often than the immutable elements. They are fetched every 1 hour.

- `keepCRV`: The ratio of CRV to keep in the strategy
- `keepCRVPercent`: The ratio of CRV to keep in the strategy in percent
- `keepCVX`: The ratio of CVX to keep in the strategy
- `emergencyExit`: If the strategy is in emergency exit

#### The mutable elements

We are expecting them to change often. They are fetched every 15 minutes.

- `lastCreditAvailable`: The last known credit available
- `lastDebtOutstanding`: The last known debt outstanding
- `lastExpectedReturn`: The last known expected return
- `lastReport`: The last time the strategy was reported
- `lastTotalDebt`: The last known total debt
- `lastTotalGain`: The last known total gain
- `lastTotalLoss`: The last known total loss
- `lastEstimatedTotalAssets`: The last known estimated total assets
- `lastDelegatedAssets`: The last known delegated assets
- `lastPerformanceFee`: The last known performance fee
- `lastDebtRatio`: The last known debt ratio. Only > 0.2.2
- `lastDebtLimit`: The last known debt limit. Only = 0.2.2
- `lastRateLimit`: The last known rate limit. Only < 0.3.2
- `lastMinDebtPerHarvest`: The last known min debt per harvest. Only >= 0.3.2
- `lastMaxDebtPerHarvest`: The last known max debt per harvest. Only >= 0.3.2
- `lastUpdate`: The last time the update was performed

#### The manual elements

For some specific use cases, we want to be able to manually update some elements. This is a manual process, not automated and a reboot of the bot is required to update them.

- `isRetired`: If the strategy is retired and will be ignored
- `displayName`: An override of the name to be displayed
- `description`: A description for this strategy
- `protocols`: The protocols used by this strategy

#### The controls elements

We may want to trigger refresh or some extra action. Theses elements can be added to the JSON and will trigger a specific action. They will be removed before being stored again in the JSON.

- `shouldRefresh`: Indicate that we should bypass checks and do a full refresh of the strategy

## Example

```json
"0x000007252ab8d9120005d30aa15567ea8de9a110": {
	"address": "0x000007252ab8d9120005d30aa15567ea8de9a110",
	"vaultAddress": "0x490bd0886f221a5f79713d3e84404355a9293c50",
	"keeperAddress": "0x736d7e3c5a6cb2ce3b764300140abf476f6cfccf",
	"strategistAddress": "0xc6387e937bcef8de3334f80edc623275d42457ff",
	"rewardsAddress": "0xc491599b9a20c3a2f0a85697ee6d9434efa9f503",
	"healthCheckAddress": "0xddcea799ff1699e98edf118e0629a974df7df012",
	"name": "StrategyConvexibCHF",
	"vaultVersion": "0.4.3",
	"activation": 14139749,
	"chainID": 1,
	"doHealthCheck": true,
	"isActive": false,
	"isInQueue": false,
	"withdrawalQueuePosition": "-1",
	"keepCRV": "1000",
	"keepCRVPercent": null,
	"keepCVX": null,
	"emergencyExit": false,
	"lastCreditAvailable": "0",
	"lastDebtOutstanding": "0",
	"lastExpectedReturn": "0",
	"lastReport": "1686670799",
	"lastTotalDebt": "0",
	"lastTotalGain": "118169170152051977696595",
	"lastTotalLoss": "0",
	"lastEstimatedTotalAssets": "0",
	"lastDelegatedAssets": "0",
	"lastPerformanceFee": "0",
	"lastDebtRatio": "0",
	"lastMinDebtPerHarvest": "0",
	"lastMaxDebtPerHarvest": "115792089237316195423570985008687907853269984665640564039457584007913129639935",
	"lastUpdate": "2023-10-12T11:21:02.265295+02:00",
	"isRetired": true,
	"displayName": "Convex Reinvest",
	"description": "Supplies {{token}} to [Convex Finance](https://www.convexfinance.com/stake) boosted by Convex's veCRV to earn CRV and CVX (and any other available tokens). Earned tokens are harvested, sold for more {{token}} which is deposited back into the strategy.",
	"protocols": [
		"Convex Finance",
		"Curve Finance"
	]
}
```
