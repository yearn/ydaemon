package env

import (
	"math"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/internal/models"
)

var OPTIMISM = TChain{
	ID:              10,
	RpcURI:          `https://mainnet.optimism.io`,
	SubgraphURI:     `https://api.thegraph.com/subgraphs/name/yearn/yearn-vaults-v2-optimism`,
	EtherscanURI:    `https://api-optimistic.etherscan.io/api`,
	MaxBlockRange:   100_000_000,
	MaxBatchSize:    math.MaxInt64,
	AvgBlocksPerDay: 43_200,
	LensContract: TContractData{
		Address: common.HexToAddress(`0xB082d9f4734c535D9d80536F7E87a6f4F471bF65`),
		Block:   18109291,
	},
	MulticallContract: TContractData{
		Address: common.HexToAddress(`0xca11bde05977b3631167028862be2a173976ca11`),
		Block:   4286263,
	},
	StackingRewardContract: TContractData{
		Address: common.HexToAddress(`0x8ed9f6343f057870f1def47aae7cd88dfaa049a8`),
		Block:   85969070,
	},
	PartnerContract: TContractData{
		Address: common.HexToAddress(`0x7E08735690028cdF3D81e7165493F1C34065AbA2`),
		Block:   29675215,
	},
	Coin: models.TERC20Token{
		Address:                   DEFAULT_COIN_ADDRESS,
		UnderlyingTokensAddresses: []common.Address{},
		Type:                      models.TokenTypeNative,
		Name:                      `Ether`,
		Symbol:                    `ETH`,
		DisplayName:               `Ether`,
		DisplaySymbol:             `ETH`,
		Description:               `Optimism is a Layer 2 scaling solution for Ethereum.`,
		Icon:                      BASE_ASSET_URL + strconv.FormatUint(10, 10) + `/` + DEFAULT_COIN_ADDRESS.Hex() + `/logo-128.png`,
		Decimals:                  18,
		ChainID:                   10,
	},
	Registries: []TContractData{
		{
			Address: common.HexToAddress("0x79286Dd38C9017E5423073bAc11F53357Fc5C128"),
			Version: 3,
			Block:   22451152,
		},
	},
	ExtraVaults: []models.TVaultsFromRegistry{},
	BlacklistedVaults: []common.Address{
		common.HexToAddress("0x6884bd538Db61A626Da0a05E10807BFC5Aea2b32"), // Test deployment - Nothing
		common.HexToAddress("0xDB8bBF2b0e28721F9BAc603e687E39bcF52201f8"), // Test deployment - Nothing
		common.HexToAddress("0xed5D83bB6Af23bcb05C144DC816f45A389d622a0"), // Test deployment - Nothing
		common.HexToAddress("0x5e70E0EdE43373bD9CF4Bc85199d424AcF0241EF"), // Empty sUSD
	},
	ExtraTokens: []common.Address{
		common.HexToAddress("0x4200000000000000000000000000000000000042"), // Opt
	},
	IgnoredTokens: []common.Address{
		common.HexToAddress("0x6884bd538Db61A626Da0a05E10807BFC5Aea2b32"), // Test deployment - Nothing
		common.HexToAddress("0xDB8bBF2b0e28721F9BAc603e687E39bcF52201f8"), // Test deployment - Nothing
		common.HexToAddress("0xed5D83bB6Af23bcb05C144DC816f45A389d622a0"), // Test deployment - Nothing
	},
	Curve: TChainCurve{
		RegistryAddress: common.HexToAddress(`0x0000000022d53366457f9d5e68ec105046fc4383`),
		FactoryAddress:  common.Address{},
		PoolsURIs: []string{
			`https://api.curve.fi/api/getPools/all/optimism`,
		},
		GaugesURI: `https://api.curve.fi/api/getAllGauges?blockchainId=optimism`,
	},
	ExtraURI: TChainExtraURI{
		GammaMerklURI: `https://api.angle.money/v2/merkl?chainIds%5B%5D=10`,
		GammaHypervisorURI: []string{
			`https://wire2.gamma.xyz/optimism/hypervisors/allData`,
		},
		PendleCoreURI: `https://api-v2.pendle.finance/core/v1/10`,
	},
}
