package env

import (
	"math"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/internal/models"
)

var ARBITRUM = TChain{
	ID:              42161,
	RpcURI:          `https://arbitrum.public-rpc.com`,
	SubgraphURI:     `https://api.thegraph.com/subgraphs/name/yearn/yearn-vaults-v2-arbitrum`,
	EtherscanURI:    `https://api.arbiscan.io/api`,
	MaxBlockRange:   100_000_000,
	MaxBatchSize:    math.MaxInt64,
	AvgBlocksPerDay: 320_000,
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
	APROracleContract: TContractData{
		Address: common.HexToAddress(`0x27aD2fFc74F74Ed27e1C0A19F1858dD0963277aE`),
		Block:   19070394,
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
			Block:   4_841_854,
			Tag:     `DISABLED`,
		},
		{
			Address: common.HexToAddress("0xff31A1B020c868F6eA3f61Eb953344920EeCA3af"),
			Version: 4,
			Block:   171_850_013,
		},
	},
	ExtraVaults: []models.TVaultsFromRegistry{
		{
			//yvMIM, alone in it's own registry, not work registering and listening to it
			ChainID:         42161,
			Address:         common.HexToAddress(`0x074943fEfE3391D033A15557dfa1b6f246Ce5fD0`),
			RegistryAddress: common.HexToAddress(`0xff31A1B020c868F6eA3f61Eb953344920EeCA3af`),
			TokenAddress:    common.HexToAddress(`0xE11f9786B06438456b044B3E21712228ADcAA0D1`),
			APIVersion:      `3.0.2`,
			BlockNumber:     195564702,
			Type:            models.TokenTypeAutomatedVault,
		},
	},
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
		PoolsURIs: []string{
			`https://api.curve.fi/api/getPools/all/arbitrum`,
		},
	},
	ExtraURI: TChainExtraURI{
		GammaMerklURI: `https://api.angle.money/v2/merkl?chainIds%5B%5D=42161`,
		GammaHypervisorURI: []string{
			`https://wire2.gamma.xyz/arbitrum/hypervisors/allData`,
		},
		PendleCoreURI: `https://api-v2.pendle.finance/core/v1/42161`,
	},
}
