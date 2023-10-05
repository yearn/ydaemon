package env

import (
	"math"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/internal/models"
)

var ARBITRUM = TChain{
	ID:            42161,
	RpcURI:        `https://arbitrum.public-rpc.com`,
	SubgraphURI:   `https://api.thegraph.com/subgraphs/name/yearn/yearn-vaults-v2-arbitrum`,
	MaxBlockRange: 100_000_000,
	MaxBatchSize:  math.MaxInt64,
	LensContract: TContractData{
		Address: common.HexToAddress(`0x043518AB266485dC085a1DB095B8d9C2Fc78E9b9`),
		Block:   2396321,
	},
	MulticallContract: TContractData{
		Address: common.HexToAddress(`0x842eC2c7D803033Edf55E478F461FC547Bc54EB2`),
		Block:   821923,
	},
	PartnerContract: TContractData{
		Address: common.HexToAddress(`0x0e5b46E4b2a05fd53F5a4cD974eb98a9a613bcb7`),
		Block:   30385403,
	},
	Coin: models.TERC20Token{
		Address:                   DEFAULT_COIN_ADDRESS,
		UnderlyingTokensAddresses: []common.Address{},
		Type:                      models.TokenTypeNative,
		Name:                      `Arbitrum`,
		Symbol:                    `ARB`,
		DisplayName:               `Arbitrum`,
		DisplaySymbol:             `ARB`,
		Description:               `Arbitrum is a Layer 2 scaling solution for Ethereum.`,
		Icon:                      BASE_ASSET_URL + strconv.FormatUint(42161, 10) + `/` + DEFAULT_COIN_ADDRESS.Hex() + `/logo-128.png`,
		Decimals:                  18,
		ChainID:                   42161,
	},
	Registries: []TContractData{
		{
			Address: common.HexToAddress("0x3199437193625DCcD6F9C9e98BDf93582200Eb1f"),
			Version: 2,
			Block:   4841854,
		},
	},
	ExtraVaults: []models.TVaultsFromRegistry{},
	BlacklistedVaults: []common.Address{
		common.HexToAddress("0x5796698A29F3626c9FE13C4d3d3dEE987c84EBB3"), // Test deployment - Nothing
		common.HexToAddress("0x976a1C749cd8153909e0B04EebE931eF8957b15b"), // Test deployment - PHPTest
		common.HexToAddress("0xFa247d0D55a324ca19985577a2cDcFC383D87953"), // Test deployment - PHP
	},
	ExtraTokens: []common.Address{
		common.HexToAddress(`0x82e3A8F066a6989666b031d916c43672085b1582`), // YFI
		common.HexToAddress(`0x11cDb42B0EB46D95f990BeDD4695A6e3fA034978`), // CRV
	},
	IgnoredTokens: []common.Address{
		common.HexToAddress("0x5796698A29F3626c9FE13C4d3d3dEE987c84EBB3"), // Test deployment - Nothing
		common.HexToAddress("0x976a1C749cd8153909e0B04EebE931eF8957b15b"), // Test deployment - PHPTest
		common.HexToAddress("0xFa247d0D55a324ca19985577a2cDcFC383D87953"), // Test deployment - PHP
	},
	Curve: TChainCurve{
		RegistryAddress: common.HexToAddress(`0x0000000022d53366457f9d5e68ec105046fc4383`),
		FactoryAddress:  common.Address{},
		FactoryURIs: []string{
			`https://api.curve.fi/api/getPools/arbitrum/factory`,
		},
		PoolsURIs: []string{
			`https://api.curve.fi/api/getPools/arbitrum/main`,
			`https://api.curve.fi/api/getPools/arbitrum/crypto`,
			`https://api.curve.fi/api/getPools/arbitrum/factory`,
		},
	},
}
