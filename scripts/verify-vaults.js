#!/usr/bin/env bun

import {readdir, readFile} from 'fs/promises';
import {join} from 'path';
import Ajv from 'ajv';

const ajv = new Ajv();

// Define the schema for vault metadata
const vaultSchema = {
	type: 'object',
	properties: {
		address: {
			type: 'string',
			pattern: '^0x[a-fA-F0-9]{40}$'
		},
		token: {
			type: 'string',
			pattern: '^0x[a-fA-F0-9]{40}$'
		},
		registry: {
			type: 'string',
			pattern: '^0x[a-fA-F0-9]{40}$'
		},
		accountant: {
			type: 'string',
			pattern: '^0x[a-fA-F0-9]{40}$'
		},
		type: {type: 'string'},
		kind: {type: 'string'},
		version: {type: 'string'},
		activation: {type: 'number'},
		chainID: {type: 'number'},
		endorsed: {type: 'boolean'},
		performanceFee: {type: 'number'},
		managementFee: {type: 'number'},
		emergencyShutdown: {type: 'boolean'},
		lastActiveStrategies: {
			type: ['array', 'null'],
			items: {
				type: 'string',
				pattern: '^0x[a-fA-F0-9]{40}$'
			}
		},
		lastPricePerShare: {type: ['string', 'null']},
		lastTotalAssets: {type: ['string', 'null']},
		metadata: {
			type: 'object',
			properties: {
				isRetired: {type: 'boolean'},
				isHidden: {type: 'boolean'},
				isAggregator: {type: 'boolean'},
				isBoosted: {type: 'boolean'},
				isAutomated: {type: 'boolean'},
				isHighlighted: {type: 'boolean'},
				isPool: {type: 'boolean'},
				shouldUseV2APR: {type: 'boolean'},
				migration: {
					type: 'object',
					properties: {
						available: {type: 'boolean'},
						target: {
							type: 'string',
							pattern: '^0x[a-fA-F0-9]{40}$'
						},
						contract: {
							type: 'string',
							pattern: '^0x[a-fA-F0-9]{40}$'
						}
					},
					required: ['available', 'target', 'contract']
				},
				stability: {
					type: 'object',
					properties: {
						stability: {type: 'string'}
					},
					required: ['stability']
				},
				category: {type: 'string'},
				displayName: {type: 'string'},
				displaySymbol: {type: 'string'},
				description: {type: 'string'},
				sourceURI: {type: 'string'},
				uiNotice: {type: 'string'},
				protocols: {type: ['null', 'array']},
				inclusion: {
					type: 'object',
					properties: {
						isSet: {type: 'boolean'},
						isYearn: {type: 'boolean'},
						isYearnJuiced: {type: 'boolean'},
						isGimme: {type: 'boolean'},
						isPoolTogether: {type: 'boolean'},
						isMorpho: {type: 'boolean'},
						isPublicERC4626: {type: 'boolean'}
					},
					required: ['isSet', 'isYearn', 'isYearnJuiced', 'isGimme', 'isPoolTogether', 'isPublicERC4626']
				},
				riskLevel: {type: 'number'},
				riskScore: {
					type: 'object',
					properties: {
						review: {type: 'number'},
						testing: {type: 'number'},
						complexity: {type: 'number'},
						riskExposure: {type: 'number'},
						protocolIntegration: {type: 'number'},
						centralizationRisk: {type: 'number'},
						externalProtocolAudit: {type: 'number'},
						externalProtocolCentralisation: {type: 'number'},
						externalProtocolTvl: {type: 'number'},
						externalProtocolLongevity: {type: 'number'},
						externalProtocolType: {type: 'number'},
						comment: {type: 'string'}
					},
					required: [
						'review',
						'testing',
						'complexity',
						'riskExposure',
						'protocolIntegration',
						'centralizationRisk',
						'externalProtocolAudit',
						'externalProtocolCentralisation',
						'externalProtocolTvl',
						'externalProtocolLongevity',
						'externalProtocolType',
						'comment'
					]
				}
			},
			required: [
				'isRetired',
				'isHidden',
				'isAggregator',
				'isBoosted',
				'isAutomated',
				'isHighlighted',
				'isPool',
				'shouldUseV2APR',
				'migration',
				'stability',
				'category',
				'displayName',
				'displaySymbol',
				'description',
				'sourceURI',
				'uiNotice',
				'protocols',
				'inclusion',
				'riskLevel',
				'riskScore'
			]
		}
	},
	required: [
		'address',
		'token',
		'registry',
		'type',
		'kind',
		'version',
		'activation',
		'chainID',
		'endorsed',
		'performanceFee',
		'managementFee',
		'emergencyShutdown',
		'lastActiveStrategies',
		'lastPricePerShare',
		'lastTotalAssets',
		'metadata'
	]
};

const fullSchema = {
	type: 'object',
	properties: {
		lastUpdate: {type: 'string'},
		version: {
			type: 'object',
			properties: {
				major: {type: 'number'},
				minor: {type: 'number'},
				patch: {type: 'number'}
			},
			required: ['major', 'minor', 'patch']
		},
		shouldRefresh: {type: 'boolean'},
		vaults: {
			type: 'object',
			patternProperties: {
				'^0x[a-fA-F0-9]{40}$': vaultSchema
			},
			additionalProperties: false
		}
	},
	required: ['lastUpdate', 'version', 'shouldRefresh', 'vaults']
};

const validate = ajv.compile(vaultSchema);
const validateMeta = ajv.compile(fullSchema);

async function verifyVaultFiles() {
	const vaultsDir = join(process.cwd(), 'data', 'meta', 'vaults');
	const files = await readdir(vaultsDir);
	const jsonFiles = files.filter(file => file.endsWith('.json'));

	let hasError = false;

	for (const file of jsonFiles) {
		const filePath = join(vaultsDir, file);
		const content = await readFile(filePath, 'utf-8');

		try {
			const data = JSON.parse(content);

			// Validate each vault
			for (const [address, vault] of Object.entries(data.vaults)) {
				const isValid = validate(vault);

				if (!isValid) {
					console.error(`Validation failed for vault ${address} in file ${file}:`);
					console.error(validate.errors);
					hasError = true;
				}
			}

			// Validate meta properties
			const isMetaValid = validateMeta(data);
			if (!isMetaValid) {
				console.error(`Meta validation failed for file ${file}:`);
				console.error(validateMeta.errors);
				hasError = true;
				continue; // Skip vault validation if meta validation fails
			}
		} catch (error) {
			console.error(`Error parsing ${file}:`, error.message);
			hasError = true;
		}
	}

	if (hasError) {
		console.error('Validation failed. Please fix the errors before pushing.');
		process.exit(1);
	} else {
		console.log('All vault files are valid.');
	}
}

verifyVaultFiles().catch(error => {
	console.error('An unexpected error occurred:', error);
	process.exit(1);
});
