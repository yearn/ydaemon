# Vaults
The Vaults JSON file contains all the information about the vaults. As other JSON files we use, it is composed of a `lastUpdate` field, a `version` field and a `shouldRefresh` field.
- `lastUpdate`: The last time the file was updated
- `version`: The version of the file (major, minor, patch).
	- `major`: When at least one vault is removed since the last version
	- `minor`: When at least one vault is added since the last version
	- `patch`: When at least one vault is updated since the last version
- `shouldRefresh`: If the file should be refreshed or not. This is used by the bot to know if it should update the file or not, no matter if the other conditions are met.

Each vault contains theses elements:

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

#### The manual elements

For some specific use cases, we want to be able to manually update some elements. This is a manual process, not automated and a reboot of the bot is required to update them.

- `isRetired`: If the vault is retired
- `isHidden`: If we should hide the vaults from the standard results
- `isAggregator`: If the vault should be treated as an aggregator vault (aka multi-strategy) even if he only has one strategy
- `migration`: Indicates the vault migration data and availability
- `classification`: Give some more details to be able to classify the vaults in some UIs.
- `displayName`: The name of the vault to use rather than the onchain name
- `displaySymbol`: The symbol of the vault to use rather than the onchain symbol
- `description`: The description of the vault. Note: if the vault is also a strategy and this field is empty, the strategy description will be used instead.

## Example

```json
"lastUpdate": "2023-10-20T14:43:22.410975+02:00",
"version": {
	"major": 0,
	"minor": 0,
	"patch": 0
},
"shouldRefresh": false,
"vaults": {
	"0x00e8eb340f8af587eea6200d2081e31dc87285ac": {
		"address": "0x00e8eb340f8af587eea6200d2081e31dc87285ac",
		"token": "0x326290a1b0004eee78fa6ed4f1d8f4b2523ab669",
		"management": "0x16388463d60ffe0661cf7f1f31a7d658ac790ff7",
		"governance": "0xfeb4acf3df3cdea7399794d0869ef76a6efaff52",
		"guardian": "0x2c01b4ad51a67e2d8f02208f54df9ac4c0b778b6",
		"rewards": "0x93a62da5a14c80f265dabc077fcee437b1a0efde",
		"version": "0.4.5",
		"activation": 0,
		"chainID": 1,
		"endorsed": true,
		"performanceFee": 1000,
		"managementFee": 0,
		"emergencyShutdown": false,
		"lastActiveStrategies": [
			"0xa48e7a7e205147b6b0dd053ad8401fb36c1c2d3a",
			"0x2427e9e6c6ec0fe3baf9b42516d61cee7a2ef769",
			"0xd441b754114a84d274da0fdf1a34c06dbe78dd1a"
		],
		"lastPricePerShare": "1000000000000000000",
		"lastTotalAssets": "29882939331446425808937",
		"lastUpdate": "2023-10-12T11:13:18.96498+02:00",
		"isRetired": false,
		"isHidden": false,
		"migration": {
			"available": false,
			"target": "0x00e8eb340f8af587eea6200d2081e31dc87285ac",
			"contract": "0x0000000000000000000000000000000000000000"
		},
		"classification": {
			"isAutomated": true,
			"isPool": true,
			"poolProvider": "Curve",
			"stability": "Stable",
			"stableBaseAsset": "USD"
		},
		"displayName": "",
		"displaySymbol": "",
		"description": ""
	},
	// Some other vaults
}
```
