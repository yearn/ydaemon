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
	CanUseWebsocket: true,
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
			Label:   `YEARN`,
		},
		{
			Address: common.HexToAddress("0x444045c5C13C246e117eD36437303cac8E250aB0"),
			Version: 5,
			Block:   12_295_872,
			Tag:     `STEALTH`,
			Label:   `PUBLIC_ERC4626`,
		},
		{
			Address: common.HexToAddress("0x770D0d1Fb036483Ed4AbB6d53c1C88fb277D812F"),
			Version: 5,
			Block:   21802552,
			Label:   `YEARN`,
		},
		{
			Address: common.HexToAddress("0xd40ecF29e001c76Dcc4cC0D9cd50520CE845B038"),
			Version: 4,
			Block:   22344791,
			Label:   `YEARN`,
		},
	},
	APROracleContract: TContractData{
		Address: common.HexToAddress(`0x1981AD9F44F2EA9aDd2dC4AD7D075c102C70aF92`),
		Block:   63167152,
	},
	ExtraVaults:       []models.TVaultsFromRegistry{},
	BlacklistedVaults: []common.Address{},
	ExtraTokens: []common.Address{
		common.HexToAddress("0xf0f326af3b1Ed943ab95C29470730CC8Cf66ae47"), // wAjna
	},
	IgnoredTokens: []common.Address{},
	Curve: TChainCurve{
		RegistryAddress: common.Address{},
		FactoryAddress:  common.Address{},
		PoolsURIs: []string{
			`https://api.curve.fi/api/getPools/all/base`,
		},
		GaugesURI: `https://api.curve.fi/api/getAllGauges?blockchainId=base`,
	},
	ExtraURI: TChainExtraURI{
		GammaMerklURI: `https://api.angle.money/v2/merkl?chainIds%5B%5D=8453`,
		PendleCoreURI: ``,
	},
}
