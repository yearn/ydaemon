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
| `riskLevel` | `int` | The risk level of the vault. The value is a calculated from the sum of all risk score from the object for Single Strategy Vaults. Multi-Strategy Vault, highest `riskLevel` of all strategies is set. 1 is the most secure and 5 is the least secure.  | ❌ |
| `protocols` | `string[]` | The protocols used by the vault. The first one is used to defined the main APR method | ❌ |
| `inclusion` | `object` | Which project should include this vault. It's auto-set the first time an not updated after | ❌ |
| `riskScore` | `object` | All risk scores of the Single Strategy Vault. Multi-Strategy Vault won't have this object because its risk score is combination of multiple vaults. For risk value use `riskLevel`. (empty for Multi-Strategy Vault) | ❌ |

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

#### The risk score object

| Field | Type | Description | Automatic update |
| --- | --- | --- | --- |
| `review` | `int` | The risk review score from Yearn security. 1 is for high trust to 5 low trust | ❌ |
| `testing` | `int` | The testing coverage of the strategy being evaluated. 5 -> 80% or less; 1 -> 100% or higher | ❌ |
| `complexity` | `int` | The sLOC count of the strategy being evaluated. 5 -> 750+ sLOC; 4 -> 450-600 sLOC; 3 -> 300-450 sLOC; 2 -> 150-300 sLOC; 1 -> 0-150 sLOC | ❌ |
| `riskExposure` | `int` | This score aims to find out how much and how often a strategy can be subject to losses. 5 -> Loss of funds or non recoverable funds up to 70-100% (Example, Leveraging cross assets and got liquidated, adding liquidity to volatile pairs single sided); 4 -> Loss of funds or non recoverable funds up to 15-70% (Example, adding liquidity to single sided curve stable pools); 3 -> Loss of funds or non recoverable funds up to 10-15% (Example, Protocol specific IL exposure, very high deposit/withdrawal fees); 2 -> Loss of funds or non recoverable funds up to 0-10% (Example, deposit/withdrawal fees or anything protocol specific); 1 -> Strategy has no lossable cases, only gains, up only. | ❌ |
| `protocolIntegration` | `int` | The protocols that are integrated into the strategy that is being evaluated. 5 -> Strategy interacts with 5 external protocols; 4 -> Strategy interacts with 4 external protocols; 3 -> Strategy interacts with 3 external protocols; 2 -> Strategy interacts with 2 external protocols; 1 -> Strategy interacts with 1 external protocol | ❌ |
| `centralizationRisk` | `int` | The centralization score of the strategy that is being evaluated. 5 -> Strategy heavily relies on off-chain management, potentially exposing user funds to rug possibilities by admins; 4 -> Strategy frequently depends on off-chain management but has safeguards against rug possibilities by admins; 3 -> Strategy involves privileged roles but less frequently and with less risk of rug possibilities; 2 -> Strategy has privileged roles but they are not vital for operations and pose minimal risk of rug possibilities; 1 -> Strategy operates without dependency on any privileged roles, ensuring full permissionlessness. | ❌ |
| `externalProtocolAudit` | `int` | The public audits count of the external protocols. 5 -> No audit conducted by a trusted firm or security researcher; 4 -> Audit conducted by 1 trusted firm or security researcher conducted; 3 -> Audit conducted by 2 trusted firm or security researcher conducted; 2 -> Audit conducted by 3 trusted firm or security researcher conducted; 1 -> Audit conducted by 4 or more trusted firm or security researcher conducted | ❌ |
| `externalProtocolCentralisation` | `int` | Measurement of the centralization score of the external protocols. 5 -> Contracts owner is an EOA or a multisig with less than 4 members OR Contracts are not verified OR Contracts owner can harm our strategy completely; 4 -> Contracts owner is a multisig but the addresses are not known/hidden OR Contracts owner can harm our strategy by setting parameters in external protocol contracts up to some degree; 3 -> Contracts owner is a multisig with known people but multisig threshold is very low; 2 -> Contracts owner is a multisig with known trusted people; 1 -> Contracts owner is a multisig with known trusted people with Timelock OR Contracts are governanceless, immutable OR Contracts owner can't do any harm to our strategy by setting parameters in external protocol contracts | ❌ |
| `externalProtocolTvl` | `int` | The active TVL that the external protocol holds. 5 -> TVL of $10M or less; 4 -> TVL between $10M and $40M;3 -> TVL between $40M and $120M; 2 -> TVL; between $120M and $480M; 1 -> TVL of $480M or more | ❌ |
| `externalProtocolLongevity` | `int` | How long the external protocol contracts in scope have been deployed alive. 5 -> Less than 6 months; 4 -> Between 6 and 12 months; 3 -> Between 12 and 18 months; 2 -> Between 18 and 24 months; 1 -> 24 months or more | ❌ |
| `externalProtocolType` | `int` | This is a rough estimate of evaluating a protocol's purpose. 5 -> The main expertise of the protocol lies in off-chain operations, such as RWA protocols; 4 -> Cross-chain applications, like cross-chain bridges, cross-chain yield aggregators, and cross-chain lending/borrowing protocols; 3 -> AMM lending/borrowing protocols that are not forks of blue-chip protocols, leveraged farming protocols, as well as newly conceptualized protocols; 2 -> Slightly modified forked blue-chip protocols; 1 -> Blue-chip protocols such as AAVE, Compound, Uniswap, Curve, Convex, and Balancer. | ❌ |

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
				},
				"riskScore": {
					"review": 1,
					"testing": 1,
					"complexity": 1,
					"riskExposure": 1,
					"protocolIntegration": 1,
					"centralizationRisk": 1,
					"externalProtocolAudit": 1,
					"externalProtocolCentralisation": 1,
					"externalProtocolTvl": 1,
					"externalProtocolLongevity": 1,
					"externalProtocolType": 1
				}
			}
		},
	// Some other vaults
}
```
