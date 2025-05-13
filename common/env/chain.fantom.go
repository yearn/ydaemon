package env

import (
	"math"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/internal/models"
)

var FANTOM = TChain{
	ID:              250,
	RpcURI:          `https://rpc.ftm.tools`,
	SubgraphURI:     ``,
	EtherscanURI:    `https://api.ftmscan.com/api`,
	MaxBlockRange:   100_000_000,
	MaxBatchSize:    math.MaxInt64,
	AvgBlocksPerDay: 45_000,
	CanUseWebsocket: true,
	LensContract: TContractData{
		Address: common.HexToAddress(`0x57AA88A0810dfe3f9b71a9b179Dd8bF5F956C46A`),
		Block:   17091856,
	},
	MulticallContract: TContractData{
		Address: common.HexToAddress(`0x470ADB45f5a9ac3550bcFFaD9D990Bf7e2e941c9`),
		Block:   19328162,
	},
	PartnerContract: TContractData{
		Address: common.HexToAddress(`0x086865B2983320b36C42E48086DaDc786c9Ac73B`),
		Block:   40499061,
	},
	Coin: models.TERC20Token{
		Address:                   DEFAULT_COIN_ADDRESS,
		UnderlyingTokensAddresses: []common.Address{},
		Type:                      models.TokenTypeNative,
		Name:                      `Fantom`,
		Symbol:                    `FTM`,
		DisplayName:               `Fantom`,
		DisplaySymbol:             `FTM`,
		Description:               `Fantom is a Layer 2 scaling solution for Ethereum.`,
		Icon:                      BASE_ASSET_URL + strconv.FormatUint(250, 10) + `/` + DEFAULT_COIN_ADDRESS.Hex() + `/logo-128.png`,
		Decimals:                  18,
		ChainID:                   250,
	},
	Registries: []TContractData{
		{
			Address: common.HexToAddress("0x727fe1759430df13655ddb0731dE0D0FDE929b04"),
			Version: 2,
			Block:   18455565,
			Tag:     `DISABLED`,
			Label:   `YEARN`,
		},
	},
	ExtraVaults: []models.TVaultsFromRegistry{
		{
			//yvMIM, alone in it's own registry, not work registering and listening to it
			ChainID:         250,
			Address:         common.HexToAddress(`0x0A0b23D9786963DE69CB2447dC125c49929419d8`),
			RegistryAddress: common.HexToAddress(`0x265F7b1413F6B06654746cf2485082182389A5d0`),
			TokenAddress:    common.HexToAddress(`0x82f0b8b456c1a451378467398982d4834b6829c1`),
			APIVersion:      `0.4.3`,
			BlockNumber:     18302860,
			Type:            models.TokenTypeStandardVault,
		},
	},
	BlacklistedVaults: []common.Address{
		common.HexToAddress("0x03B82e4070cA32FF63A03F2EcfC16c0165689a9d"), // Test deployment - AVAX
	},
	ExtraTokens: []common.Address{},
	IgnoredTokens: []common.Address{
		common.HexToAddress("0x03B82e4070cA32FF63A03F2EcfC16c0165689a9d"), // Test deployment - AVAX
	},
	Curve: TChainCurve{
		RegistryAddress: common.HexToAddress(`0x0000000022d53366457f9d5e68ec105046fc4383`),
		FactoryAddress:  common.Address{},
		PoolsURIs: []string{
			`https://api.curve.finance/api/getPools/all/fantom`,
		},
		GaugesURI: `https://api.curve.finance/api/getAllGauges?blockchainId=fantom`,
	},
	ExtraURI: TChainExtraURI{
		PendleCoreURI: ``,
	},
}
