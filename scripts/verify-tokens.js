#!/usr/bin/env bun

import {readdir, readFile} from 'fs/promises';
import {join} from 'path';
import Ajv from 'ajv';

const ajv = new Ajv();

// Define the schema for token metadata
const tokenSchema = {
	type: 'object',
	properties: {
		address: {
			type: 'string',
			pattern: '^0x[a-fA-F0-9]{40}$'
		},
		underlyingTokensAddresses: {
			type: ['array', 'null'],
			items: {
				type: 'string',
				pattern: '^0x[a-fA-F0-9]{40}$'
			}
		},
		type: {type: 'string'},
		name: {type: 'string'},
		symbol: {type: 'string'},
		displayName: {type: 'string'},
		displaySymbol: {type: 'string'},
		description: {type: 'string'},
		category: {type: 'string'},
		icon: {type: 'string'},
		decimals: {type: 'number'},
		chainID: {type: 'number'}
	},
	required: [
		'address',
		'underlyingTokensAddresses',
		'type',
		'name',
		'symbol',
		'displayName',
		'displaySymbol',
		'description',
		'category',
		'icon',
		'decimals',
		'chainID'
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
		tokens: {
			type: 'object',
			patternProperties: {
				'^0x[a-fA-F0-9]{40}$': tokenSchema
			},
			additionalProperties: false
		}
	},
	required: ['lastUpdate', 'version', 'shouldRefresh', 'tokens']
};

const validate = ajv.compile(tokenSchema);
const validateMeta = ajv.compile(fullSchema);

async function verifyTokenFiles() {
	const tokensDir = join(process.cwd(), 'data', 'meta', 'tokens');
	const files = await readdir(tokensDir);
	const jsonFiles = files.filter(file => file.endsWith('.json'));

	let hasError = false;

	for (const file of jsonFiles) {
		const filePath = join(tokensDir, file);
		const content = await readFile(filePath, 'utf-8');

		try {
			const data = JSON.parse(content);

			// Validate each token
			for (const [address, tokens] of Object.entries(data.tokens)) {
				const isValid = validate(tokens);

				if (!isValid) {
					console.error(`Validation failed for token ${address} in file ${file}:`);
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
		console.log('All token files are valid.');
	}
}

verifyTokenFiles().catch(error => {
	console.error('An unexpected error occurred:', error);
	process.exit(1);
});
