package env

import (
	"math"
	"path"
	"path/filepath"
	"runtime"

	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/internal/models"
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
var API_V1_BASE_URL = `https://api.yearn.finance/v1/chains/`

// SUPPORTED_CHAIN_IDS is the list of supported chain IDs
var SUPPORTED_CHAIN_IDS = []uint64{1, 10, 250, 42161}

// MAX_BLOCK_RANGE is the maximum number of blocks we can query in a single request
var MAX_BLOCK_RANGE = map[uint64]uint64{
	1:     100_000_000,
	10:    100_000_000,
	250:   100_000_000,
	42161: 100_000_000,
}

// MAX_BATCH_SIZE is the maximum size request for a multicall
var MAX_BATCH_SIZE = map[uint64]int{
	1:     math.MaxInt64,
	10:    1_000,
	250:   math.MaxInt64,
	42161: math.MaxInt64,
}

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
	56:    `https://bsc.rpc.blxrbdn.com`,
	137:   `https://polygon.llamarpc.com`,
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

// LLAMA_CHAIN_NAMES contains the chain identifiers for the DeFiLlama API
var LLAMA_CHAIN_NAMES = map[uint64]string{
	1:     `ethereum`,
	10:    `optimism`,
	250:   `fantom`,
	42161: `arbitrum`,
}

// GECKO_CHAIN_NAMES contains the chain identifiers for the CoinGecko API
var GECKO_CHAIN_NAMES = map[uint64]string{
	1:     `ethereum`,
	10:    `optimistic-ethereum`,
	250:   `fantom`,
	42161: `arbitrum-one`,
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
	250: {
		{Address: common.HexToAddress("0x727fe1759430df13655ddb0731dE0D0FDE929b04"), Version: 2, Block: 18455565},
	},
	42161: {
		{Address: common.HexToAddress("0x3199437193625DCcD6F9C9e98BDf93582200Eb1f"), Version: 2, Block: 4841854},
	},
}

// EXTRA_VAULTS is a list of vaults that are not registered in the registries, but are still used by Yearn
var EXTRA_VAULTS = map[uint64][]models.TVaultsFromRegistry{
	1:  {},
	10: {},
	250: {
		{
			//yvMIM, alone in it's own registry, not work registering and listening to it
			ChainID:         250,
			Address:         common.HexToAddress(`0x0A0b23D9786963DE69CB2447dC125c49929419d8`),
			RegistryAddress: common.HexToAddress(`0x265F7b1413F6B06654746cf2485082182389A5d0`),
			TokenAddress:    common.HexToAddress(`0x82f0b8b456c1a451378467398982d4834b6829c1`),
			APIVersion:      `0.4.3`,
			BlockNumber:     18309707,
			Activation:      18302860,
			ManagementFee:   200,
			BlockHash:       common.HexToHash(`0x00009ee300000d281b4c0169bb3320b32f435e3fd830fe1625adcfd4cf6410cb`),
			TxIndex:         0,
			LogIndex:        0,
			Type:            models.VaultTypeStandard,
		},
	},
	42161: {},
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
	56:    common.HexToAddress(`0xca11bde05977b3631167028862be2a173976ca11`),
	137:   common.HexToAddress(`0xca11bde05977b3631167028862be2a173976ca11`),
	250:   common.HexToAddress(`0x470ADB45f5a9ac3550bcFFaD9D990Bf7e2e941c9`),
	42161: common.HexToAddress(`0x842eC2c7D803033Edf55E478F461FC547Bc54EB2`),
}

// CURVE_REGISTRY_ADDRESSES contains the address of the Curve Registry contract
var CURVE_REGISTRY_ADDRESSES = map[uint64]common.Address{
	1:     common.HexToAddress(`0x90E00ACe148ca3b23Ac1bC8C240C2a7Dd9c2d7f5`),
	10:    common.HexToAddress(`0x0000000022d53366457f9d5e68ec105046fc4383`),
	250:   common.HexToAddress(`0x0000000022d53366457f9d5e68ec105046fc4383`),
	42161: common.HexToAddress(`0x0000000022d53366457f9d5e68ec105046fc4383`),
}

// CURVE_FACTORIES_ADDRESSES contains the address of the Curve Registry contract
var CURVE_FACTORIES_ADDRESSES = map[uint64]common.Address{
	1:     common.HexToAddress(`0xF18056Bbd320E96A48e3Fbf8bC061322531aac99`),
	10:    {},
	250:   {},
	42161: {},
}

// STACKING_REWARD_ADDRESSES contains the address of the stacking reward contract
var STACKING_REWARD_ADDRESSES = map[uint64]TContractData{
	1: {},
	10: {
		Address: common.HexToAddress(`0x8ed9f6343f057870f1def47aae7cd88dfaa049a8`),
		Block:   uint64(85969070),
	},
	250:   {},
	42161: {},
}

// YBRIBE_V3_ADDRESSES contains the address of the yBribe contract
var YBRIBE_V3_ADDRESSES = map[uint64]TContractData{
	1: {
		Address: common.HexToAddress(`0x03dFdBcD4056E2F92251c7B07423E1a33a7D3F6d`),
		Block:   uint64(15878262),
	},
	10:    {},
	250:   {},
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
	250: {
		Address: common.HexToAddress(`0x086865B2983320b36C42E48086DaDc786c9Ac73B`),
		Block:   uint64(40499061),
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
	250:   {},
	42161: {},
}
