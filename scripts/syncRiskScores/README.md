# Scripts Documentation

This directory contains utility scripts for maintaining and synchronizing data in the ydaemon project.

## sync-risk-scores.js

A Node.js script that synchronizes vault risk information from ydaemon's vault metadata to the [risk-score repository](https://github.com/yearn/risk-score).

### Purpose

This script ensures that vault risk data (riskLevel and riskScore objects) stored in ydaemon's vault metadata files are properly synchronized to the risk-score repository structure. It creates individual JSON files for each vault that has valid risk data.

**IMPORTANT**: This script NEVER overwrites existing files in the risk-score repository. Existing files are considered authoritative and take higher priority than ydaemon data. The script only creates NEW files for vaults that don't already exist in the risk-score repository.

### Features

- **Multi-chain support**: Processes vaults across all supported chains (1, 10, 100, 137, 146, 250, 42161, 747474, 8453)
- **Intelligent filtering**: Only processes vaults that have complete and valid risk data
- **File preservation**: Never overwrites existing files in risk-score repository
- **Chain directory management**: Automatically creates chain-specific directories in risk-score repo
- **Comprehensive logging**: Provides detailed logs of operations performed
- **Dry-run capability**: Test what would happen without making actual changes
- **Error handling**: Gracefully handles missing files, invalid data, and other error conditions

### Prerequisites

1. The [risk-score repository](https://github.com/yearn/risk-score) must be cloned in the same parent directory as ydaemon:
   ```
   parent-directory/
   ├── ydaemon/
   └── risk-score/
   ```

2. Node.js must be installed on the system

### Usage

```bash
# Basic usage - sync all vault risk data
node scripts/syncRiskScores/syncRiskScores.js

# Dry run - see what would happen without making changes
node scripts/syncRiskScores/syncRiskScores.js --dry-run

# Verbose output - see detailed logs of operations
node scripts/syncRiskScores/syncRiskScores.js --verbose

# Combine flags
node scripts/syncRiskScores/syncRiskScores.js --dry-run --verbose
```

### Output Structure

The script creates files in the risk-score repository following this structure:

```
risk-score/
└── vaults/
    ├── 1/           # Ethereum mainnet
    ├── 137/         # Polygon
    ├── 42161/       # Arbitrum
    ├── 8453/        # Base
    └── ...          # Other chains
        └── [vaultAddress].json
```

Each vault file contains:

```json
{
    "riskLevel": 2,
    "riskScore": {
        "review": 2,
        "testing": 1,
        "complexity": 1,
        "riskExposure": 3,
        "protocolIntegration": 1,
        "centralizationRisk": 1,
        "externalProtocolAudit": 2,
        "externalProtocolCentralisation": 3,
        "externalProtocolTvl": 2,
        "externalProtocolLongevity": 1,
        "externalProtocolType": 4,
        "comment": ""
    }
}
```

### Validation Rules

The script only processes vaults that meet all these criteria:

1. **riskLevel**: Must be a number between 1-5
2. **riskScore**: Must be an object containing all required fields
3. **All risk score fields**: Must be numbers between 1-5:
   - review, testing, complexity, riskExposure
   - protocolIntegration, centralizationRisk
   - externalProtocolAudit, externalProtocolCentralisation
   - externalProtocolTvl, externalProtocolLongevity, externalProtocolType
4. **comment**: Can be any string (including empty)

### Example Output

```
[2025-08-20T15:51:16.597Z] [INFO] Starting vault risk score synchronization
[2025-08-20T15:51:16.597Z] [INFO] Found 9 vault files to process
[2025-08-20T15:51:16.597Z] [INFO] Created risk score for vault 0x06ed7c67755344548fafe1822bee365c4208a57f on chain 137
[2025-08-20T15:51:16.597Z] [DEBUG] Skipping vault 0x985cc9c306bfe075f7c67ec275fb0b80f0b21976 - file already exists in risk-score repo (preserving existing data)

=== SYNCHRONIZATION SUMMARY ===
Total vaults processed: 1138
Risk score files created: 98
Vaults skipped (no risk data or existing files preserved): 1040
Errors encountered: 0

Synchronization completed successfully
```

### Error Handling

The script handles various error conditions:

- **Missing risk-score repository**: Exits with clear error message
- **Invalid vault data**: Skips vaults with incomplete or invalid risk data
- **File system errors**: Logs errors and continues processing other vaults
- **JSON parsing errors**: Logs errors and continues with next file

### Integration

This script is designed to be run:

- **Manually** when vault risk data is updated
- **As part of CI/CD** to ensure risk-score repository stays synchronized
- **On a schedule** to catch any missed updates

### Troubleshooting

**"Risk score repository not found"**
- Ensure risk-score repository is cloned in the same parent directory as ydaemon

**"No vault files found to process"**
- Check that vault JSON files exist in `data/meta/vaults/`
- Ensure files follow naming pattern `[chainId].json`

**"Skipping vault [address] - file already exists in risk-score repo"**
- This is normal - the script preserves existing files and never overwrites them
