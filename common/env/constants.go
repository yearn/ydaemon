package env

import (
	"path"
	"path/filepath"
	"runtime"

	"github.com/yearn/ydaemon/common/types/common"
)

func getCurrentPath() string {
	_, filename, _, _ := runtime.Caller(1)
	return path.Dir(filename)
}

// GITHUB_ICON_BASE_URL is the base URL to access the tokens icons from the yearn-assets repo
var GITHUB_ICON_BASE_URL = `https://raw.githubusercontent.com/yearn/yearn-assets/master/icons/multichain-tokens/`

// BASE_DATA_PATH is the base path to access the data informations
var BASE_DATA_PATH, _ = filepath.Abs(getCurrentPath() + `../../../data/`)

// API_V1_BASE_URL is the base URL to access query the legacy Yearn's api
var API_V1_BASE_URL = `https://api.yearn.finance/v1/chains/`

// SUPPORTED_CHAIN_IDS is the list of supported chain IDs
var SUPPORTED_CHAIN_IDS = []uint64{1, 10, 250, 42161}

// BLACKLISTED_VAULTS contains the vault we should not work with
var BLACKLISTED_VAULTS = map[uint64][]common.Address{
	1: {
		common.HexToAddress("0x662fBF2c1E4b04342EeBA6371ec1C7420042B86F"), // Test deployment - Nothing
		common.HexToAddress("0x9C13e225AE007731caA49Fd17A41379ab1a489F4"), // Test deployment - Nothing
		common.HexToAddress("0xBF7AA989192b020a8d3e1C65a558e123834325cA"), // HBTC yVault - Not Yearn - PPS 0
		common.HexToAddress("0x4Fdd1B06eF986238446be0F3EA163C1b6Fe28cC1"), // GUSD yVault - Not Yearn - PPS 100
		common.HexToAddress("0x8a0889d47f9Aa0Fac1cC718ba34E26b867437880"), // Old st-yCRV
		common.HexToAddress("0x61f46C65E403429266e8b569F23f70dD75d9BeE7"), // Old lp-yCRV

	},
	10: {
		common.HexToAddress("0x6884bd538Db61A626Da0a05E10807BFC5Aea2b32"), // Test deployment - Nothing
		common.HexToAddress("0xDB8bBF2b0e28721F9BAc603e687E39bcF52201f8"), // Test deployment - Nothing
		common.HexToAddress("0xed5D83bB6Af23bcb05C144DC816f45A389d622a0"), // Test deployment - Nothing
	},
	250: {
		common.HexToAddress("0x03B82e4070cA32FF63A03F2EcfC16c0165689a9d"), // Test deployment - AVAX
	},
	42161: {
		common.HexToAddress("0x5796698A29F3626c9FE13C4d3d3dEE987c84EBB3"), // Test deployment - Nothing
		common.HexToAddress("0x976a1C749cd8153909e0B04EebE931eF8957b15b"), // Test deployment - PHPTest
		common.HexToAddress("0xFa247d0D55a324ca19985577a2cDcFC383D87953"), // Test deployment - PHP
	},
}

// CURVE_FACTORY_URI contains the URI of the Curve Factory to use
var CURVE_FACTORY_URI = map[uint64][]string{
	1: {
		`https://api.curve.fi/api/getPools/ethereum/factory`,
		`https://api.curve.fi/api/getPools/ethereum/factory-crypto`,
	},
	10: {
		`https://api.curve.fi/api/getPools/optimism/factory`,
	},
	250: {
		`https://api.curve.fi/api/getPools/fantom/factory`,
	},
	42161: {
		`https://api.curve.fi/api/getPools/arbitrum/factory`,
	},
}

// RPC_ENDPOINTS contains the node endpoints to connect the blockchains
// Can be overwritten by env variables
var RPC_ENDPOINTS = map[uint64]string{
	1:     `https://eth.public-rpc.com`,
	10:    `https://mainnet.optimism.io`,
	250:   `https://rpc.ftm.tools`,
	42161: `https://arbitrum.public-rpc.com`,
}

// THEGRAPH_ENDPOINTS contains the URI of the GraphQL provider to use
var THEGRAPH_ENDPOINTS = map[uint64]string{
	1:     `https://api.thegraph.com/subgraphs/name/rareweasel/yearn-vaults-v2-subgraph-mainnet`,
	10:    `https://api.thegraph.com/subgraphs/name/yearn/yearn-vaults-v2-optimism`,
	250:   `https://api.thegraph.com/subgraphs/name/bsamuels453/yearn-fantom-validation-grafted`,
	42161: `https://api.thegraph.com/subgraphs/name/yearn/yearn-vaults-v2-arbitrum`,
}

// LENS_ADDRESSES contains the address of the Lens oracle for a specific chainID
var LENS_ADDRESSES = map[uint64]common.Address{
	1:     common.HexToAddress(`0x83d95e0D5f402511dB06817Aff3f9eA88224B030`),
	10:    common.HexToAddress(`0xB082d9f4734c535D9d80536F7E87a6f4F471bF65`),
	250:   common.HexToAddress(`0x57AA88A0810dfe3f9b71a9b179Dd8bF5F956C46A`),
	42161: common.HexToAddress(`0x043518AB266485dC085a1DB095B8d9C2Fc78E9b9`),
}

// MULTICALL_ADDRESSES contains the address of the multicall2 contract for a specific chainID
var MULTICALL_ADDRESSES = map[uint64]common.Address{
	1:     common.HexToAddress(`0x5ba1e12693dc8f9c48aad8770482f4739beed696`),
	10:    common.HexToAddress(`0xca11bde05977b3631167028862be2a173976ca11`),
	250:   common.HexToAddress(`0x470ADB45f5a9ac3550bcFFaD9D990Bf7e2e941c9`),
	42161: common.HexToAddress(`0x842eC2c7D803033Edf55E478F461FC547Bc54EB2`),
}

// YBRIBE_V3_ADDRESSES contains the address of the yBribe contract
var YBRIBE_V3_ADDRESSES = map[uint64]common.Address{
	1:     common.HexToAddress(`0x03dFdBcD4056E2F92251c7B07423E1a33a7D3F6d`),
	10:    {},
	250:   {},
	42161: {},
}

// STRATEGIES_CONDITIONS contains the possible conditions to determine which strategies should
// be used in the return value.
// If the strategy is `absolute`, an active strategy will be isActive, inQueue and with a debt > 0
// If the strategy is `inQueue`, we will check if the vault has the strategy in it's withdrawal queue
// If the strategy is `debtLimit`, we will check if the vault has allocated a debtLimit to the strategy
var STRATEGIES_CONDITIONS = []string{`inQueue`, `debtLimit`, `absolute`}
