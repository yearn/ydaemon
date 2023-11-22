package env

import (
	"math"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/internal/models"
)

var BASE = TChain{
	ID:              8453,
	RpcURI:          `https://developer-access-mainnet.base.org`,
	SubgraphURI:     `https://api.thegraph.com/subgraphs/name/rareweasel/yearn-vaults-v2-subgraph-base`,
	EtherscanURI:    `https://api.basescan.org/api`,
	MaxBlockRange:   100_000_000,
	MaxBatchSize:    math.MaxInt64,
	AvgBlocksPerDay: 43_200,
	LensContract: TContractData{
		Address: common.HexToAddress(`0xE0F3D78DB7bC111996864A32d22AB0F59Ca5Fa86`),
		Block:   3318817,
	},
	MulticallContract: TContractData{
		Address: common.HexToAddress(`0xca11bde05977b3631167028862be2a173976ca11`),
		Block:   5022,
	},
	Coin: models.TERC20Token{
		Address:                   DEFAULT_COIN_ADDRESS,
		UnderlyingTokensAddresses: []common.Address{},
		Type:                      models.TokenTypeNative,
		Name:                      `Ether`,
		Symbol:                    `ETH`,
		DisplayName:               `Ether`,
		DisplaySymbol:             `ETH`,
		Description:               `Base is a Layer 2 scaling solution based on Optimism.`,
		Icon:                      BASE_ASSET_URL + strconv.FormatUint(8453, 10) + `/` + DEFAULT_COIN_ADDRESS.Hex() + `/logo-128.png`,
		Decimals:                  18,
		ChainID:                   8453,
	},
	Registries: []TContractData{
		{
			Address: common.HexToAddress("0xF3885eDe00171997BFadAa98E01E167B53a78Ec5"),
			Version: 3,
			Block:   3263730,
		},
	},
	ExtraVaults:       []models.TVaultsFromRegistry{},
	BlacklistedVaults: []common.Address{},
	ExtraTokens:       []common.Address{},
	IgnoredTokens:     []common.Address{},
	Curve: TChainCurve{
		RegistryAddress: common.Address{},
		FactoryAddress:  common.Address{},
		FactoryURIs: []string{
			`https://api.curve.fi/api/getPools/base/factory`,
		},
		PoolsURIs: []string{
			`https://api.curve.fi/api/getPools/base/main`,
			`https://api.curve.fi/api/getPools/base/crypto`,
			`https://api.curve.fi/api/getPools/base/factory`,
		},
	},
}
