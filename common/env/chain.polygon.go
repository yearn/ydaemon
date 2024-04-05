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
	EtherscanURI:    `https://api.polygonscan.com/api`,
	MaxBlockRange:   100_000_000,
	MaxBatchSize:    math.MaxInt64,
	AvgBlocksPerDay: 40_000,
	LensContract:    TContractData{}, //TODO: not deployed
	MulticallContract: TContractData{
		Address: common.HexToAddress(`0xca11bde05977b3631167028862be2a173976ca11`),
		Block:   25770160,
	},
	APROracleContract: TContractData{
		Address: common.HexToAddress(`0x27aD2fFc74F74Ed27e1C0A19F1858dD0963277aE`),
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
		},
		{
			Address: common.HexToAddress("0xff31A1B020c868F6eA3f61Eb953344920EeCA3af"),
			Version: 4,
			Block:   52_488_140,
		},
		{
			Address: common.HexToAddress("0xd5967178702250d9f0eac34258ebba99b9a28ed0"),
			Version: 6,
			Block:   51_550_217,
		},
	},
	ExtraVaults: []models.TVaultsFromRegistry{},
	BlacklistedVaults: []common.Address{
		common.HexToAddress(`0x0128E1D15ED9f0c8572967825Ef46309BDA39836`), //Test USDC
	},
	ExtraTokens:   []common.Address{},
	IgnoredTokens: []common.Address{},
	Curve: TChainCurve{
		RegistryAddress: common.HexToAddress(`0x0000000022d53366457f9d5e68ec105046fc4383`),
		FactoryAddress:  common.Address{},
		PoolsURIs: []string{
			`https://api.curve.fi/api/getPools/all/polygon`,
		},
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
