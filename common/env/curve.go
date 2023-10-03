package env

import (
	"github.com/ethereum/go-ethereum/common"
)

// CURVE_FACTORY_URI contains the URI of the Curve Factory to use
var CURVE_FACTORY_URI = map[uint64][]string{
	1: {
		`https://api.curve.fi/api/getPools/ethereum/factory`,
		`https://api.curve.fi/api/getPools/ethereum/factory-crypto`,
	},
	10: {
		`https://api.curve.fi/api/getPools/optimism/factory`,
	},
	137: {
		`https://api.curve.fi/api/getPools/polygon/factory`,
	},
	250: {
		`https://api.curve.fi/api/getPools/fantom/factory`,
	},
	8453: {
		`https://api.curve.fi/api/getPools/base/factory`,
	},
	42161: {
		`https://api.curve.fi/api/getPools/arbitrum/factory`,
	},
}

// CURVE_POOLS_URI contains the URI of the Curve pools to use
var CURVE_POOLS_URI = map[uint64][]string{
	1: {
		`https://api.curve.fi/api/getPools/ethereum/main`,
		`https://api.curve.fi/api/getPools/ethereum/crypto`,
		`https://api.curve.fi/api/getPools/ethereum/factory`,
		`https://api.curve.fi/api/getPools/ethereum/factory-crypto`,
	},
	10: {
		`https://api.curve.fi/api/getPools/optimism/main`,
		`https://api.curve.fi/api/getPools/optimism/crypto`,
		`https://api.curve.fi/api/getPools/optimism/factory`,
	},
	137: {
		`https://api.curve.fi/api/getPools/polygon/main`,
		`https://api.curve.fi/api/getPools/polygon/crypto`,
		`https://api.curve.fi/api/getPools/polygon/factory`,
	},
	250: {
		`https://api.curve.fi/api/getPools/fantom/main`,
		`https://api.curve.fi/api/getPools/fantom/crypto`,
		`https://api.curve.fi/api/getPools/fantom/factory`,
	},
	8453: {
		`https://api.curve.fi/api/getPools/base/main`,
		`https://api.curve.fi/api/getPools/base/crypto`,
		`https://api.curve.fi/api/getPools/base/factory`,
	},
	42161: {
		`https://api.curve.fi/api/getPools/arbitrum/main`,
		`https://api.curve.fi/api/getPools/arbitrum/crypto`,
		`https://api.curve.fi/api/getPools/arbitrum/factory`,
	},
}

// CURVE_REGISTRY_ADDRESSES contains the address of the Curve Registry contract
var CURVE_REGISTRY_ADDRESSES = map[uint64]common.Address{
	1:     common.HexToAddress(`0x90E00ACe148ca3b23Ac1bC8C240C2a7Dd9c2d7f5`),
	10:    common.HexToAddress(`0x0000000022d53366457f9d5e68ec105046fc4383`),
	137:   common.HexToAddress(`0x0000000022d53366457f9d5e68ec105046fc4383`),
	250:   common.HexToAddress(`0x0000000022d53366457f9d5e68ec105046fc4383`),
	8453:  {}, //TODO: ADD CURVE_REGISTRY_ADDRESSES FOR BASE
	42161: common.HexToAddress(`0x0000000022d53366457f9d5e68ec105046fc4383`),
}

// CURVE_FACTORIES_ADDRESSES contains the address of the Curve Registry contract
var CURVE_FACTORIES_ADDRESSES = map[uint64]common.Address{
	1:     common.HexToAddress(`0xF18056Bbd320E96A48e3Fbf8bC061322531aac99`),
	10:    {},
	137:   {},
	250:   {},
	8453:  {},
	42161: {},
}
