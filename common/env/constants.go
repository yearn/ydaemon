package env

import (
	"math"
	"path"
	"path/filepath"
	"runtime"

	"github.com/ethereum/go-ethereum/common"
)

type TContractData struct {
	Address    common.Address // Address of the contract
	Block      uint64         // Block number where the contract was deployed
	Version    uint64         // Version of the contract. May be empty.
	Activation uint64         // Timestamp of the contract activation. May be empty.
}

func getCurrentPath() string {
	_, filename, _, _ := runtime.Caller(1)
	return path.Dir(filename)
}

// BASE_DATA_PATH is the base path to access the data informations
var BASE_DATA_PATH, _ = filepath.Abs(getCurrentPath() + `../../../data/`)

// var BASE_ASSET_URL = `https://raw.githubusercontent.com/yearn/yearn-assets/master/icons/multichain-tokens/`
// BASE_ASSET_URL is the base URL to access the tokens icons from the assets repo
var BASE_ASSET_URL = `https://assets.smold.app/api/token/`

// GECKO_PRICE_URL contains the URL for the CoinGecko API
var GECKO_PRICE_URL = `https://api.coingecko.com/api/v3/simple/token_price/`

// LLAMA_PRICE_URL contains the URL for the DeFiLlama pricing API
var LLAMA_PRICE_URL = `https://coins.llama.fi/prices/current/`

// API_V1_BASE_URL is the base URL to access query the legacy Yearn's api
var API_V1_BASE_URL = `https://api.yexporter.io/v1/chains/`

// SUPPORTED_CHAIN_IDS is the list of supported chain IDs
var SUPPORTED_CHAIN_IDS = []uint64{1, 10, 137, 250, 8453, 42161}

// MAX_BLOCK_RANGE is the maximum number of blocks we can query in a single request
var MAX_BLOCK_RANGE = map[uint64]uint64{
	1:     100_000_000,
	10:    100_000_000,
	137:   100_000_000,
	250:   100_000_000,
	8453:  100_000_000,
	42161: 100_000_000,
}

// MAX_BATCH_SIZE is the maximum size request for a multicall
var MAX_BATCH_SIZE = map[uint64]int{
	1:     math.MaxInt64,
	10:    math.MaxInt64,
	137:   math.MaxInt64,
	250:   math.MaxInt64,
	8453:  math.MaxInt64,
	42161: math.MaxInt64,
}

// RPC_ENDPOINTS contains the node endpoints to connect the blockchains
// Can be overwritten by env variables
var RPC_ENDPOINTS = map[uint64]string{
	1:     `https://eth.public-rpc.com`,
	10:    `https://mainnet.optimism.io`,
	56:    `https://bsc.rpc.blxrbdn.com`,
	137:   `https://polygon.llamarpc.com`,
	250:   `https://rpc.ftm.tools`,
	8453:  `https://developer-access-mainnet.base.org`,
	42161: `https://arbitrum.public-rpc.com`,
}

// THEGRAPH_ENDPOINTS contains the URI of the GraphQL provider to use
var THEGRAPH_ENDPOINTS = map[uint64]string{
	1:  `https://api.thegraph.com/subgraphs/name/rareweasel/yearn-vaults-v2-subgraph-mainnet`,
	10: `https://api.thegraph.com/subgraphs/name/yearn/yearn-vaults-v2-optimism`,
	// 137:   //TODO: ADD,
	250:   `https://api.thegraph.com/subgraphs/name/yearn/yearn-vaults-v2-fantom`,
	8453:  `https://api.thegraph.com/subgraphs/name/rareweasel/yearn-vaults-v2-subgraph-base`,
	42161: `https://api.thegraph.com/subgraphs/name/yearn/yearn-vaults-v2-arbitrum`,
}

// YEARN_REGISTRIES is the list of registries used by Yearn across all the supported chains, with the version and the activation block
var YEARN_REGISTRIES = map[uint64][]TContractData{
	1: {
		{Address: common.HexToAddress("0xe15461b18ee31b7379019dc523231c57d1cbc18c"), Version: 1, Block: 11563389},
		{Address: common.HexToAddress("0x50c1a2eA0a861A967D9d0FFE2AE4012c2E053804"), Version: 2, Block: 12045555},
		{Address: common.HexToAddress("0xaF1f5e1c19cB68B30aAD73846eFfDf78a5863319"), Version: 3, Block: 16215519},
	},
	10: {
		{Address: common.HexToAddress("0x81291ceb9bB265185A9D07b91B5b50Df94f005BF"), Version: 3, Block: 22450349},
		{Address: common.HexToAddress("0x79286Dd38C9017E5423073bAc11F53357Fc5C128"), Version: 3, Block: 22451152},
	},
	137: {
		{Address: common.HexToAddress("0xFBB087B456a656Ab815EB2D0f3f21Aa409Cec33F"), Version: 4, Block: 47462833},
	},
	250: {
		{Address: common.HexToAddress("0x727fe1759430df13655ddb0731dE0D0FDE929b04"), Version: 2, Block: 18455565},
	},
	8453: {
		{Address: common.HexToAddress("0xF3885eDe00171997BFadAa98E01E167B53a78Ec5"), Version: 3, Block: 3263730},
	},
	42161: {
		{Address: common.HexToAddress("0x3199437193625DCcD6F9C9e98BDf93582200Eb1f"), Version: 2, Block: 4841854},
	},
}

// LENS_ADDRESSES contains the address of the Lens oracle for a specific chainID
var LENS_ADDRESSES = map[uint64]common.Address{
	1:  common.HexToAddress(`0x83d95e0D5f402511dB06817Aff3f9eA88224B030`),
	10: common.HexToAddress(`0xB082d9f4734c535D9d80536F7E87a6f4F471bF65`),
	// 137:   //TODO: ADD THIS
	250:   common.HexToAddress(`0x57AA88A0810dfe3f9b71a9b179Dd8bF5F956C46A`),
	8453:  common.HexToAddress(`0xE0F3D78DB7bC111996864A32d22AB0F59Ca5Fa86`),
	42161: common.HexToAddress(`0x043518AB266485dC085a1DB095B8d9C2Fc78E9b9`),
}

// MULTICALL_ADDRESSES contains the address of the multicall2 contract for a specific chainID
var MULTICALL_ADDRESSES = map[uint64]common.Address{
	1:     common.HexToAddress(`0x5ba1e12693dc8f9c48aad8770482f4739beed696`),
	10:    common.HexToAddress(`0xca11bde05977b3631167028862be2a173976ca11`),
	56:    common.HexToAddress(`0xca11bde05977b3631167028862be2a173976ca11`),
	137:   common.HexToAddress(`0xca11bde05977b3631167028862be2a173976ca11`),
	250:   common.HexToAddress(`0x470ADB45f5a9ac3550bcFFaD9D990Bf7e2e941c9`),
	8453:  common.HexToAddress(`0xca11bde05977b3631167028862be2a173976ca11`),
	42161: common.HexToAddress(`0x842eC2c7D803033Edf55E478F461FC547Bc54EB2`),
}

// STACKING_REWARD_ADDRESSES contains the address of the stacking reward contract
var STACKING_REWARD_ADDRESSES = map[uint64]TContractData{
	1: {},
	10: {
		Address: common.HexToAddress(`0x8ed9f6343f057870f1def47aae7cd88dfaa049a8`),
		Block:   uint64(85969070),
	},
	137:   {},
	250:   {},
	8453:  {},
	42161: {},
}

// YBRIBE_V3_ADDRESSES contains the address of the yBribe contract
var YBRIBE_V3_ADDRESSES = map[uint64]TContractData{
	1: {
		Address: common.HexToAddress(`0x03dFdBcD4056E2F92251c7B07423E1a33a7D3F6d`),
		Block:   uint64(15878262),
	},
	10:    {},
	137:   {},
	250:   {},
	8453:  {},
	42161: {},
}

// PARTNER_TRACKERS_ADDRESSES contains the address of the yBribe contract
var PARTNER_TRACKERS_ADDRESSES = map[uint64]TContractData{
	1: {
		Address: common.HexToAddress(`0x8ee392a4787397126C163Cb9844d7c447da419D8`),
		Block:   uint64(14166636),
	},
	10: {
		Address: common.HexToAddress(`0x7E08735690028cdF3D81e7165493F1C34065AbA2`),
		Block:   uint64(29675215),
	},
	137: {
		//TODO: ADD PARTNER_TRACKERS_ADDRESSES FOR POLYGON
	},
	250: {
		Address: common.HexToAddress(`0x086865B2983320b36C42E48086DaDc786c9Ac73B`),
		Block:   uint64(40499061),
	},
	8453: {
		//TODO: ADD PARTNER_TRACKERS_ADDRESSES FOR BASE
	},
	42161: {
		Address: common.HexToAddress(`0x0e5b46E4b2a05fd53F5a4cD974eb98a9a613bcb7`),
		Block:   uint64(30385403),
	},
}

var KnownRefferers = map[uint64]map[common.Address]string{
	1: {
		common.HexToAddress(`0x558247e365be655f9144e1a0140D793984372Ef3`): `Ledger`,
		common.HexToAddress(`0xFEB4acf3df3cDEA7399794D0869ef76A6EfAff52`): `Yearn.finance`,
	},
	10:    {},
	137:   {},
	250:   {},
	8453:  {},
	42161: {},
}
