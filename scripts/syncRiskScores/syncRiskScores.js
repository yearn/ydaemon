#!/usr/bin/env node

/************************************************************************************************
* Vault Risk Score Synchronization Script
* 
* This script synchronizes vault risk information from ydaemon's vault metadata to the 
* risk-score repository. It processes all vaults across different chains and creates new
* risk score files in the risk-score repository structure.
*
* IMPORTANT: This script NEVER overwrites existing files in the risk-score repository.
* Existing files are considered authoritative and take higher priority than ydaemon data.
*
* Features:
* - Processes all chain vault files (1.json, 10.json, 137.json, etc.)
* - Creates chain-specific directories in risk-score repo if they don't exist
* - Only syncs vaults that have risk data (riskLevel and riskScore)
* - PRESERVES all existing risk score files (never overwrites)
* - Only creates NEW files for vaults not yet present in risk-score repo
* - Provides detailed logging of operations performed
*
* Usage: node scripts/syncRiskScores/syncRiskScores.js [--dry-run] [--verbose]
************************************************************************************************/

const fs = require('fs');
const path = require('path');

// Configuration
const YDAEMON_ROOT = __dirname + '/../..';
const VAULTS_DIR = path.join(YDAEMON_ROOT, 'data', 'meta', 'vaults');
const RISK_SCORE_ROOT = path.join(YDAEMON_ROOT, '..', 'risk-score');
const RISK_SCORE_STRATEGY_DIR = path.join(RISK_SCORE_ROOT, 'strategy');

// Command line arguments
const args = process.argv.slice(2);
const isDryRun = args.includes('--dry-run');
const isVerbose = args.includes('--verbose') || isDryRun;

/**
 * Logs messages based on verbosity settings
 */
function log(message, level = 'info') {
    if (isVerbose || level === 'error') {
        const timestamp = new Date().toISOString();
        console.log(`[${timestamp}] [${level.toUpperCase()}] ${message}`);
    }
}

/**
 * Checks if a vault has valid risk data that should be synced
 */
function hasValidRiskData(vault) {
    const metadata = vault.metadata;
    if (!metadata) return false;
    
    // Check if riskLevel exists and is a valid number (1-5)
    if (typeof metadata.riskLevel !== 'number' || 
        metadata.riskLevel < 1 || 
        metadata.riskLevel > 5) {
        return false;
    }
    
    // Check if riskScore exists and has the expected structure
    if (!metadata.riskScore || typeof metadata.riskScore !== 'object') {
        return false;
    }
    
    // Verify riskScore has required fields
    const requiredFields = [
        'review', 'testing', 'complexity', 'riskExposure', 
        'protocolIntegration', 'centralizationRisk', 'externalProtocolAudit',
        'externalProtocolCentralisation', 'externalProtocolTvl', 
        'externalProtocolLongevity', 'externalProtocolType'
    ];
    
    for (const field of requiredFields) {
        if (typeof metadata.riskScore[field] !== 'number' ||
            metadata.riskScore[field] < 1 ||
            metadata.riskScore[field] > 5) {
            return false;
        }
    }
    
    return true;
}

/**
 * Creates the risk score object from vault metadata
 */
function createRiskScoreData(vault) {
    const metadata = vault.metadata;
    
    return {
        riskLevel: metadata.riskLevel,
        riskScore: {
            review: metadata.riskScore.review,
            testing: metadata.riskScore.testing,
            complexity: metadata.riskScore.complexity,
            riskExposure: metadata.riskScore.riskExposure,
            protocolIntegration: metadata.riskScore.protocolIntegration,
            centralizationRisk: metadata.riskScore.centralizationRisk,
            externalProtocolAudit: metadata.riskScore.externalProtocolAudit,
            externalProtocolCentralisation: metadata.riskScore.externalProtocolCentralisation,
            externalProtocolTvl: metadata.riskScore.externalProtocolTvl,
            externalProtocolLongevity: metadata.riskScore.externalProtocolLongevity,
            externalProtocolType: metadata.riskScore.externalProtocolType,
            comment: metadata.riskScore.comment || ""
        }
    };
}

/**
 * Ensures a directory exists, creating it if necessary
 */
function ensureDirectoryExists(dirPath) {
    if (!fs.existsSync(dirPath)) {
        if (!isDryRun) {
            fs.mkdirSync(dirPath, { recursive: true });
        }
        log(`Created directory: ${dirPath}`);
        return true;
    }
    return false;
}

/**
 * Writes risk score data to file
 */
function writeRiskScoreFile(filePath, data) {
    const jsonContent = JSON.stringify(data, null, 4) + '\n';
    
    if (!isDryRun) {
        fs.writeFileSync(filePath, jsonContent, 'utf8');
    }
    
    return jsonContent;
}

/**
 * Checks if two risk score objects are different (deprecated - no longer used)
 * Kept for backward compatibility with any external usage
 */
function isRiskDataDifferent(existing, newData) {
    const existingStr = JSON.stringify(existing, null, 4);
    const newStr = JSON.stringify(newData, null, 4);
    return existingStr !== newStr;
}

/**
 * Processes a single vault file for a specific chain
 */
function processVaultFile(chainId, filePath) {
    log(`Processing chain ${chainId}: ${filePath}`);
    
    let vaultData;
    try {
        const fileContent = fs.readFileSync(filePath, 'utf8');
        vaultData = JSON.parse(fileContent);
    } catch (error) {
        log(`Error reading vault file ${filePath}: ${error.message}`, 'error');
        return { processed: 0, created: 0, skipped: 0, errors: 1 };
    }
    
    if (!vaultData.vaults) {
        log(`No vaults found in ${filePath}`, 'error');
        return { processed: 0, created: 0, skipped: 0, errors: 1 };
    }
    
    const stats = { processed: 0, created: 0, skipped: 0, errors: 0 };
    const chainDir = path.join(RISK_SCORE_STRATEGY_DIR, chainId.toString());
    let chainDirCreated = false;
    
    // Process each vault
    for (const [vaultAddress, vault] of Object.entries(vaultData.vaults)) {
        stats.processed++;
        
        if (!hasValidRiskData(vault)) {
            log(`Skipping vault ${vaultAddress} - no valid risk data`, 'debug');
            stats.skipped++;
            continue;
        }
        
        const riskScoreData = createRiskScoreData(vault);
        const riskScoreFilePath = path.join(chainDir, `${vaultAddress.toLowerCase()}.json`);
        
        try {
            // Check if file already exists - if it does, skip it (risk-score repo takes priority)
            if (fs.existsSync(riskScoreFilePath)) {
                log(`Skipping vault ${vaultAddress} - file already exists in risk-score repo (preserving existing data)`, 'debug');
                stats.skipped++;
                continue;
            }
            
            // Ensure chain directory exists only when we need to create a file
            if (!chainDirCreated) {
                ensureDirectoryExists(chainDir);
                chainDirCreated = true;
            }
            
            // Only create new files, never overwrite existing ones
            writeRiskScoreFile(riskScoreFilePath, riskScoreData);
            log(`Created risk score for vault ${vaultAddress} on chain ${chainId}`);
            stats.created++;
            
        } catch (error) {
            log(`Error processing vault ${vaultAddress}: ${error.message}`, 'error');
            stats.errors++;
        }
    }
    
    return stats;
}

/**
 * Main function that orchestrates the synchronization process
 */
function main() {
    log(`Starting vault risk score synchronization${isDryRun ? ' (DRY RUN)' : ''}`);
    log(`ydaemon root: ${YDAEMON_ROOT}`);
    log(`Risk score repo: ${RISK_SCORE_ROOT}`);
    
    // Check if risk-score repository exists
    if (!fs.existsSync(RISK_SCORE_ROOT)) {
        log(`Risk score repository not found at: ${RISK_SCORE_ROOT}`, 'error');
        log('Please ensure the risk-score repository is cloned in the same parent directory as ydaemon', 'error');
        process.exit(1);
    }
    
    // Ensure strategy directory exists
    ensureDirectoryExists(RISK_SCORE_STRATEGY_DIR);
    
    // Get all vault files
    const vaultFiles = fs.readdirSync(VAULTS_DIR)
        .filter(file => file.endsWith('.json') && file !== '_schema.json')
        .map(file => ({
            chainId: file.replace('.json', ''),
            path: path.join(VAULTS_DIR, file)
        }));
    
    if (vaultFiles.length === 0) {
        log('No vault files found to process', 'error');
        process.exit(1);
    }
    
    log(`Found ${vaultFiles.length} vault files to process`);
    
    // Process each vault file
    const totalStats = { processed: 0, created: 0, skipped: 0, errors: 0 };
    
    for (const { chainId, path: filePath } of vaultFiles) {
        const stats = processVaultFile(chainId, filePath);
        
        // Accumulate stats
        totalStats.processed += stats.processed;
        totalStats.created += stats.created;
        totalStats.skipped += stats.skipped;
        totalStats.errors += stats.errors;
    }
    
    // Print summary
    log('\n=== SYNCHRONIZATION SUMMARY ===');
    log(`Total vaults processed: ${totalStats.processed}`);
    log(`Risk score files created: ${totalStats.created}`);
    log(`Vaults skipped (no risk data or existing files preserved): ${totalStats.skipped}`);
    log(`Errors encountered: ${totalStats.errors}`);
    
    if (isDryRun) {
        log('\nDRY RUN: No files were actually modified');
        log('Run without --dry-run to apply changes');
    }
    
    if (totalStats.errors > 0) {
        log(`\nSynchronization completed with ${totalStats.errors} errors`, 'error');
        process.exit(1);
    } else {
        log('\nSynchronization completed successfully');
    }
}

// Run the script
if (require.main === module) {
    main();
}

module.exports = {
    hasValidRiskData,
    createRiskScoreData,
    isRiskDataDifferent
};
