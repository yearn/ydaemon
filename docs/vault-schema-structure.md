# Vault Schema Structure

The vault schema is structured as a JSON file containing information about Yearn vaults. The file consists of the following main components:

1. **Metadata**
   - `lastUpdate`: Timestamp of the last update
   - `version`: Object containing `major`, `minor`, and `patch` version numbers
   - `shouldRefresh`: Boolean indicating if the file should be refreshed

2. **Vaults**
   A dictionary of vault objects, keyed by their addresses. Each vault object contains:

   ## Vault Information (Automatically Updated)
   - `address`: Vault contract address
   - `token`: Underlying token address
   - `registry`: Registry contract address
   - `type`: Vault type (e.g., "Yearn Vault", "Standard", "Experimental")
   - `kind`: Vault kind ("Legacy", "Multi", or "Single" Strategy)
   - `version`: Vault version
   - `activation`: Activation block number
   - `chainID`: Chain ID
   - `endorsed`: Deprecated field indicating if the vault is endorsed by Yearn
   - `performanceFee`: Performance fee (0 to 10000)
   - `managementFee`: Management fee (0 to 10000)
   - `emergencyShutdown`: Boolean indicating if the vault is in emergency shutdown
   - `lastActiveStrategies`: Array of active strategy addresses
   - `lastPricePerShare`: Latest price per share
   - `lastTotalAssets`: Latest total assets

   ## Metadata (Manually Updated)
   - `isRetired`: Boolean indicating if the vault is retired
   - `isHidden`: Boolean indicating if the vault should be hidden from standard results
   - `isAggregator`: Boolean indicating if the vault should be treated as an aggregator vault
   - `isBoosted`: Boolean indicating if the vault benefits from boosting
   - `isAutomated`: Boolean indicating if the vault is automated
   - `isHighlighted`: Boolean indicating if the vault is highlighted
   - `isPool`: Boolean indicating if the vault uses a pool token as want
   - `migration`: Object containing migration data
   - `stability`: Object indicating the stability of the vault
   - `category`: String indicating the vault category
   - `displayName`: Custom display name for the vault
   - `displaySymbol`: Custom display symbol for the vault
   - `description`: Vault description
   - `sourceURI`: URI linked to the vault
   - `uiNotice`: Notice to display in the UI
   - `riskLevel`: Integer representing the risk level (1-5)
   - `protocols`: Array of protocols used by the vault
   - `inclusion`: Object indicating which projects should include this vault
   - `riskScore`: Object containing detailed risk scores for Single Strategy Vaults

This structure allows for efficient storage and retrieval of vault information, combining both automatically updated fields and manually curated metadata to provide a comprehensive view of each vault's characteristics and status.

For more detailed information about the vault schema structure, please refer to the [_schema.md](../data/meta/vaults/_schema.md) file.

