package env

import (
	"math"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/internal/models"
)

var GNOSIS = TChain{
	ID:              100,
	RpcURI:          `https://rpc.gnosis.gateway.fm`,
	SubgraphURI:     ``,
	EtherscanURI:    `https://gnosisscan.io/api`,
	MaxBlockRange:   9_000,
	MaxBatchSize:    math.MaxInt64,
	AvgBlocksPerDay: 16_000,
	CanUseWebsocket: false,
	LensContract:    TContractData{},
	MulticallContract: TContractData{
		Address: common.HexToAddress(`0xca11bde05977b3631167028862be2a173976ca11`),
		Block:   821923,
	},
	PartnerContract:   TContractData{},
	APROracleContract: TContractData{},
	Coin: models.TERC20Token{
		Address:                   DEFAULT_COIN_ADDRESS,
		UnderlyingTokensAddresses: []common.Address{},
		Type:                      models.TokenTypeNative,
		Name:                      `DAI`,
		Symbol:                    `DAI`,
		DisplayName:               `DAI`,
		DisplaySymbol:             `DAI`,
		Description:               `DAI`,
		Icon:                      BASE_ASSET_URL + strconv.FormatUint(10, 10) + `/` + DEFAULT_COIN_ADDRESS.Hex() + `/logo-128.png`,
		Decimals:                  18,
		ChainID:                   10,
	},
	Registries: []TContractData{
		{
			Address: common.HexToAddress("0x444045c5C13C246e117eD36437303cac8E250aB0"),
			Version: 5,
			Block:   32_784_248,
			Tag:     `STEALTH`,
			Label:   `PUBLIC_ERC4626`,
		},
	},
	StakingRewardRegistry: []TContractData{
		{
			Address: common.HexToAddress(`0x776779de45D947B30A5D929C305Dc35eeE1F23be`),
			Block:   33343074,
			Tag:     `JUICED`,
		},
	},
	ExtraStakingContracts: []TExtraStakingContracts{
		{
			VaultAddress:   common.HexToAddress(`0x39b68451f05aaa020611cf887a7338f0991ffd60`),
			StakingAddress: common.HexToAddress(`0xd4263aBDdD2afdaAE0A0a69Eb09Deb8000dd642e`),
			Tag:            `JUICED`,
		},
	},
	ExtraVaults:       []models.TVaultsFromRegistry{},
	BlacklistedVaults: []common.Address{},
	ExtraTokens: []common.Address{
		common.HexToAddress("0x67Ee2155601e168F7777F169Cd74f3E22BB5E0cE"), // wAjna
	},
	IgnoredTokens: []common.Address{},
	Curve: TChainCurve{
		RegistryAddress: common.HexToAddress(`0x0000000022d53366457f9d5e68ec105046fc4383`),
		FactoryAddress:  common.Address{},
		PoolsURIs: []string{
			`https://api.curve.fi/api/getPools/all/xdai`,
		},
		GaugesURI: `https://api.curve.fi/api/getAllGauges?blockchainId=xdai`,
	},
	ExtraURI: TChainExtraURI{},
}
