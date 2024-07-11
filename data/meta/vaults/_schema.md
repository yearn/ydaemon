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
| `type` | `string` | Type of the vault, used for some internal computations (`Yearn Vault`, `Standard`, `Experimental Yearn Vault`, `Experimental`, `Automated Yearn Vault`, `Automated`) | N/A | ✅ |
| `kind` | `string` | Kind of the vault (`Legacy`, `Multi` or `Single` Strategy). Legacy works for V2 vaults while a V3 vault is either `Multi` or `Single`. This is set on deployment Log based on the `VaultType` log. | N/A | ✅ |
| `version` | `string` | The version of the vault | N/A | ✅ |
| `activation` | `int` | When the vault was activated | N/A | ✅ |
| `endorsed` | `bool` | If the vault is endorsed by Yearn. DEPRECATED | 24h | ✅ |
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
| `isHidden` | `bool` | If we should hide the vaults from the standard results. Indicates the vault should be hidden from the default API response | ❌ |
| `isAggregator` | `bool` | If the vault should be treated as an aggregator vault (aka multi-strategy) even if he only has one strategy | ❌ |
| `isBoosted` | `bool` | If the vault benefits from a boosting, whatever it is (ex: Curve Boost) | ❌ |
| `isAutomated` | `bool` | If the vault is automated | ✅ |
| `isHighlighted` | `bool` | If the vault is set as highlighted (specific section on UI, higher ordering score) | ❌ |
| `isPool` | `bool` | If the vault is using a pool token as want | ❌ |
| `migration` | `object` | Indicates the vault migration data and availability | ❌ |
| `stability` | `object` | Indicates the stability of the vault | ❌ |
| `category` | `string` | Indicates the category if the vault. This is set to `auto` by default (aka yDaemon will try to guess) but can be fixed by setting anything else. It's used as a case sensitive filter on the FE | ❌ |
| `displayName` | `string` | The name of the vault to use rather than the onchain name | ❌ |
| `displaySymbol` | `string` | The symbol of the vault to use rather than the onchain symbol | ❌ |
| `description` | `string` | The description of the vault | ❌ |
| `sourceURI` | `string` | A URI linked to this vault (ex: the curve deposit page) | ❌ |
| `uiNotice` | `string` | A notice to display in the UI | ❌ |
| `riskLevel` | `int` | The risk level of the vault (0 to 5, -1 for not set) | ❌ |
| `protocols` | `string[]` | The protocols used by the vault. The first one is used to defined the main APR method | ❌ |
| `inclusion` | `object` | Which project should include this vault. It's auto-set the first time an not updated after | ❌ |

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

#### The inclusion object
| Field | Type | Description | Automatic update |
| --- | --- | --- | --- |
| `isSet` | `bool` | Indicates if the inclusion has been set. It so, this will not be auto updated | ❌ |
| `isYearn` | `bool` | Indicates if the vault should be displayer on yearn.fi. True or false based on the detected registry. | ❌ |
| `isYearnJuiced` | `bool` | Indicates if the vault should be displayer on juiced.app. True or false based on the detected registry. | ❌ |
| `isGimme` | `bool` | Indicates if the vault should be displayer on gimme. Default to false. | ❌ |
| `isPublicERC4626` | `bool` | Indicates if the registry for that vault is a public ERC4626. True or false based on the detected registry. | ❌ |
| `isYearnXPoolTogether` | `bool` | Indicates if the vault should be displayer on YearnXPoolTogether UI. Default to false. | ❌ |

/!\ The first time the vault inclusion is set, if `inclusion.isPublicERC4626 == true`, the vault will be set as `endorced = false`, `isHidden = true` and `IsHighlighted = false`.


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
	"0x00cb87656196dd835b9e4d67018ae0477a1de8c1": {
			"address": "0x00cb87656196dd835b9e4d67018ae0477a1de8c1",
			"token": "0xf1376bcef0f78459c0ed0ba5ddce976f1ddf51f4",
			"registry": "0xff31a1b020c868f6ea3f61eb953344920eeca3af",
			"type": "Yearn Vault",
			"kind": "Multi Strategy",
			"version": "3.0.0",
			"activation": 20182127,
			"chainID": 1,
			"endorsed": true,
			"performanceFee": 0,
			"managementFee": 0,
			"emergencyShutdown": false,
			"lastActiveStrategies": null,
			"lastPricePerShare": null,
			"lastTotalAssets": null,
			"metadata": {
				"isRetired": false,
				"isHidden": false,
				"isAggregator": false,
				"isBoosted": false,
				"isAutomated": false,
				"isHighlighted": true,
				"isPool": false,
				"shouldUseV2APR": true,
				"migration": {
					"available": false,
					"target": "0x0000000000000000000000000000000000000000",
					"contract": "0x0000000000000000000000000000000000000000"
				},
				"stability": {
					"stability": ""
				},
				"category": "Pendle Autorollover",
				"displayName": "yPT-uniETH Yearn Auto-Rolling Pendle PT",
				"displaySymbol": "",
				"description": "This vault invests into [Pendle PT Markets](https://app.pendle.finance/trade/markets) and automatically rolls them gas-free into the next maturity upon expiry. \u003cbr/\u003e\u003cbr/\u003eRead the [Pendle Docs](https://docs.pendle.finance/) to learn about the associated risks.",
				"sourceURI": "",
				"uiNotice": "",
				"riskLevel": 1,
				"protocols": null,
				"inclusion": {
					"isSet": true,
					"isYearn": true,
					"isYearnJuiced": false,
					"isGimme": false,
					"isPublicERC4626": false
				}
			}
		},
	// Some other vaults
}
```
