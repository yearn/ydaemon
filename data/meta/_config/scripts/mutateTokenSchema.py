# run with `brownie run scripts/mutateTokenSchema.py main [selectedChainID] --network [matchingNetworkName]`
# Ethereum: `brownie run scripts/mutateTokenSchema.py main 1 --network mainnet`
# Fantom:   `brownie run scripts/mutateTokenSchema.py main 250 --network ftm-main`
# Arbitrum: `brownie run scripts/mutateTokenSchema.py main 42161 --network arbitrum-main`

import os
import json
from brownie import (Contract)
from collections import OrderedDict

ERC20ABIMIN = [{"constant":True,"inputs":[],"name":"symbol","outputs":[{"internalType":"string","name":"","type":"string"}],"payable":False,"stateMutability":"view","type":"function"},{"constant":True,"inputs":[],"name":"name","outputs":[{"internalType":"string","name":"","type":"string"}],"payable":False,"stateMutability":"view","type":"function"}]

tokensFolder = "./data/tokens/"
# NOTE: Clone https://github.com/yearn/yearn-assets/blob/master/icons/aliases.json in ./scripts/tokens.json first
tokensSource = "./scripts/tokens.json"
folders = [tokensFolder]

def executeMigration(chainID):
	for folder in folders:
		for dirpath, _, filenames in os.walk(folder):
			for fileName in filenames:
				if (fileName.endswith('json')):
					filePath = os.path.join(dirpath, fileName)
					if not filePath.startswith('./data/tokens/'+chainID):
						continue

					tokenSourcePath = os.path.join(tokensSource)
					with open(filePath, 'r') as f:
						data = OrderedDict(json.load(f))
					with open(tokenSourcePath, 'r') as f:
						source = json.load(f)

					for token in source:
						inSourceAddress = token['address'].lower()
						inDestAddress = fileName.lower().removesuffix('.json')
						if (inSourceAddress == inDestAddress):
							if ('tokenNameOverride' in data):
								data['name'] = data['tokenNameOverride']
							else:
								data['name'] = token['name']

							if ('tokenSymbolOverride' in data):
								data['symbol'] = data['tokenSymbolOverride']
							else:
								data['symbol'] = token['symbol']
						elif inDestAddress == '0x0000000000000000000000000000000000000000' and chainID == '250':
							data['name'] = 'Fantom'
							data['symbol'] = 'FTM'
						else:
							contract = Contract.from_abi(inDestAddress, inDestAddress, ERC20ABIMIN)
							if ('tokenNameOverride' in data):
								data['name'] = data['tokenNameOverride']
							else:
								data['name'] = contract.name()

							if ('tokenSymbolOverride' in data):
								data['symbol'] = data['tokenSymbolOverride']
							else:
								data['symbol'] = contract.symbol()

						ord_list = ['$schema', 'name', 'symbol', 'description', 'categories', 'website', 'localization']

						res = dict()
						for key in ord_list:
							res[key] = data[key]

					with open(filePath, 'w') as f:
						json.dump(res, f, indent=2, ensure_ascii=False, sort_keys=False)
	print("The Schema has been updated")

def main(chainID):
	executeMigration(chainID)  
