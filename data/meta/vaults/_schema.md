# Vaults
The Vaults JSON file contains all the information about the vaults. As other JSON files we use, it is composed of a `lastUpdate` field, a `version` field and a `shouldRefresh` field.
- `lastUpdate`: The last time the file was updated
- `version`: The version of the file (major, minor, patch).
	- `major`: When at least one vault is removed since the last version
	- `minor`: When at least one vault is added since the last version
	- `patch`: When at least one vault is updated since the last version
- `shouldRefresh`: If the file should be refreshed or not. This is used by the bot to know if it should update the file or not, no matter if the other conditions are met.

Each vault contains theses elements:

### The Vault Information
These elements are fetched automatically by the normal process of yDaemon with different update frequencies. No manual intervention is required and any manual change will be overwritten.  
Please do not change these fields manually.
| Field | Type | Description | Update frequency | Automatic update | 
| --- | --- | --- | --- | --- |
| `address` | `address` | Address of the vault | N/A | ✅ |
| `token` | `address` | Address of the underlying token | N/A | ✅ |
| `chainID` | `int` | The chainID of the vault | N/A | ✅ |
| `type` | `string` | Type of the vault (`Standard`, `Experimental`, `Automated`) | N/A | ✅ |
| `kind` | `string` | Kind of the vault (`Legacy`, `Multi` or `Single` Strategy) | N/A | ✅ |
| `version` | `string` | The version of the vault | N/A | ✅ |
| `activation` | `int` | When the vault was activated | N/A | ✅ |
| `endorsed` | `bool` | If the vault is endorsed by Yearn | 24h | ✅ |
| `performanceFee` | `int` | The performance fee of the vault (0 to 10000) | 1h | ✅ |
| `managementFee` | `int` | The management fee of the vault (0 to 10000) | 1h | ✅ |
| `emergencyShutdown` | `bool` | If the vault is in emergency shutdown | 1h | ✅ |
| `lastActiveStrategies` | `address[]` | The list of currently active strategies of the vault | 15m | ✅ |
| `lastPricePerShare` | `string` | The price per share of the vault | 15m | ✅ |
| `lastTotalAssets` | `string` | The total assets of the vault | 15m | ✅ |
| `metadata` | `object` | The metadata of the vault (see after) | N/A | ❌ |

### The Metadata
The metadata contains a bunch of extra information about the vault that may be useful but cannot be fetched automatically. It is composed of the following fields:
| Field | Type | Description | Automatic update |
| --- | --- | --- | --- |
| `isRetired` | `bool` | If the vault is retired. Will be excluding from APR calculation or vault fetching | ❌ |
| `isHidden` | `bool` | If we should hide the vaults from the standard results. Indicates the vault should be hidden from the UI | ❌ |
| `isAggregator` | `bool` | If the vault should be treated as an aggregator vault (aka multi-strategy) even if he only has one strategy | ❌ |
| `isBoosted` | `bool` | If the vault benefits from a boosting, whatever it is (ex: Curve Boost) | ❌ |
| `isAutomated` | `bool` | If the vault is automated | ✅ | --> DO IT
| `isPool` | `bool` | If the vault is using a pool token as want | ❌ |
| `migration` | `object` | Indicates the vault migration data and availability | ❌ |
| `stability` | `object` | Indicates the stability of the vault | ❌ |
| `displayName` | `string` | The name of the vault to use rather than the onchain name | ❌ |
| `displaySymbol` | `string` | The symbol of the vault to use rather than the onchain symbol | ❌ |
| `description` | `string` | The description of the vault | ❌ |
| `uiNotice` | `string` | A notice to display in the UI | ❌ |
| `riskLevel` | `int` | The risk level of the vault (1 to 5, -1 for not set) | ❌ |
| `protocols` | `string[]` | The protocols used by the vault. The first one is used to defined the main APR method | ❌ |

#### The migration object
| Field | Type | Description | Automatic update |
| --- | --- | --- | --- |
| `available` | `bool` | If the vault is available for migration | ❌ |
| `target` | `address` | The target address of the migration (New vault) | ❌ |
| `contract` | `address` | The contract address of the migration | ❌ |

#### The stability object
| Field | Type | Description | Automatic update |
| --- | --- | --- | --- |
| `stability` | `string` | The stability of the vault (`Stable`, `Volatile`, `Unknown`) | ❌ |
| `stableBaseAsset` | `string` | The base asset of the vault if it is stable (ex: `USD`) | ❌ |


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
			"type": "Automated Yearn Vault",
			"kind": "",
			"version": "0.4.5",
			"activation": 16373355,
			"chainID": 1,
			"endorsed": true,
			"performanceFee": 1000,
			"managementFee": 0,
			"emergencyShutdown": false,
			"lastActiveStrategies": [
				"0xa48e7a7e205147b6b0dd053ad8401fb36c1c2d3a",
				"0x2427e9e6c6ec0fe3baf9b42516d61cee7a2ef769",
				"0x9c924ee29070964cd7afde32cd6fdca620511747"
			],
			"lastPricePerShare": "1007258645643430377",
			"lastTotalAssets": "31130821900768326453019",
			"metadata": {
				"isRetired": false,
				"isHidden": false,
				"isAggregator": false,
				"isBoosted": false,
				"isAutomated": true,
				"isPool": true,
				"migration": {
					"available": false,
					"target": "0x00e8eb340f8af587eea6200d2081e31dc87285ac",
					"contract": "0x0000000000000000000000000000000000000000"
				},
				"stability": {
					"stability": "Stable",
					"stableBaseAsset": "USD"
				},
				"displayName": "",
				"displaySymbol": "",
				"description": "",
				"uiNotice": "",
				"riskLevel": -1,
				"protocols": [
					"Curve"
				]
			}
		},
	// Some other vaults
}
```
