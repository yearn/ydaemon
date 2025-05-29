package env

import (
	"math"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/internal/models"
)

var POLYGON = TChain{
	ID:              137,
	RpcURI:          `https://polygon.llamarpc.com`,
	SubgraphURI:     ``, //TODO: not deployed
	EtherscanURI:    `https://api.etherscan.io/v2/api`,
	MaxBlockRange:   100_000_000,
	MaxBatchSize:    math.MaxInt64,
	AvgBlocksPerDay: 40_000,
	CanUseWebsocket: true,
	LensContract:    TContractData{}, //TODO: not deployed
	MulticallContract: TContractData{
		Address: common.HexToAddress(`0xca11bde05977b3631167028862be2a173976ca11`),
		Block:   25770160,
	},
	APROracleContract: TContractData{
		Address: common.HexToAddress(`0x1981AD9F44F2EA9aDd2dC4AD7D075c102C70aF92`),
		Block:   52516525,
	},
	Coin: models.TERC20Token{
		Address:                   DEFAULT_COIN_ADDRESS,
		UnderlyingTokensAddresses: []common.Address{},
		Type:                      models.TokenTypeNative,
		Name:                      `Matic`,
		Symbol:                    `Matic`,
		DisplayName:               `Matic`,
		DisplaySymbol:             `Matic`,
		Description:               `Matic is the Value Layer of the Internet`,
		Icon:                      BASE_ASSET_URL + strconv.FormatUint(137, 10) + `/` + DEFAULT_COIN_ADDRESS.Hex() + `/logo-128.png`,
		Decimals:                  18,
		ChainID:                   137,
	},
	Registries: []TContractData{
		{
			Address: common.HexToAddress("0xfF5e3A7C4cBfA9Dd361385c24C3a0A4eE63CE500"),
			Version: 4,
			Block:   49_100_596,
			Label:   `YEARN`,
		},
		{
			Address: common.HexToAddress("0xff31A1B020c868F6eA3f61Eb953344920EeCA3af"),
			Version: 4,
			Block:   52_488_140,
			Label:   `YEARN`,
		},
		{
			Address: common.HexToAddress("0xd5967178702250d9f0eac34258ebba99b9a28ed0"),
			Version: 6,
			Block:   51_550_217,
			Label:   `YEARN`,
		},
		{
			Address: common.HexToAddress("0xE9E8C89c8Fc7E8b8F23425688eb68987231178e5"),
			Version: 5,
			Block:   48_907_735,
			Label:   `JUICED`,
		},
		{
			Address: common.HexToAddress("0x444045c5C13C246e117eD36437303cac8E250aB0"),
			Version: 5,
			Block:   54_308_118,
			Tag:     `STEALTH`,
			Label:   `PUBLIC_ERC4626`,
		},
		{
			Address: common.HexToAddress("0x770D0d1Fb036483Ed4AbB6d53c1C88fb277D812F"),
			Version: 5,
			Block:   63721181,
			Tag:     `STEALTH`,
			Label:   `PUBLIC_ERC4626`,
		},
		{
			Address: common.HexToAddress("0xd40ecF29e001c76Dcc4cC0D9cd50520CE845B038"),
			Version: 4,
			Block:   64224620,
			Label:   `YEARN`,
		},
	},
	StakingRewardRegistry: []TContractData{
		// {
		// 	Address: common.HexToAddress(`0x869D79ab412eB74f73b2eB983BA1b265AA1c1af1`),
		// 	Block:   19265999,
		// 	Tag:     `JUICED`,
		// },
	},
	ExtraVaults: []models.TVaultsFromRegistry{},
	BlacklistedVaults: []common.Address{
		common.HexToAddress(`0x0128E1D15ED9f0c8572967825Ef46309BDA39836`), //Test USDC
	},
	ExtraStakingContracts: []TExtraStakingContracts{
		{
			VaultAddress:   common.HexToAddress(`0xF54a15F6da443041Bb075959EA66EE47655DDFcA`),
			StakingAddress: common.HexToAddress(`0x602920E7e0a335137E02DF139CdF8D1381DAdBfD`),
			Tag:            `JUICED`,
		},
	},
	ExtraTokens: []common.Address{
		common.HexToAddress("0xA63b19647787Da652D0826424460D1BBf43Bf9c6"), // wAjna
	},
	IgnoredTokens: []common.Address{},
	Curve: TChainCurve{
		RegistryAddress: common.HexToAddress(`0x0000000022d53366457f9d5e68ec105046fc4383`),
		FactoryAddress:  common.Address{},
		PoolsURIs: []string{
			`https://api.curve.finance/api/getPools/all/polygon`,
		},
		GaugesURI: `https://api.curve.finance/api/getAllGauges?blockchainId=polygon`,
	},
	ExtraURI: TChainExtraURI{
		GammaMerklURI: `https://api.angle.money/v2/merkl?chainIds%5B%5D=137`,
		GammaHypervisorURI: []string{
			`https://wire2.gamma.xyz/polygon/hypervisors/allData`,
			`https://wire2.gamma.xyz/quickswap/polygon/hypervisors/allData`,
		},
		PendleCoreURI: ``,
	},
}
