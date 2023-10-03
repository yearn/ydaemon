package env

import (
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/internal/models"
)

var DEFAULT_COIN_ADDRESS = common.HexToAddress(`0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE`)

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
	137: {},
	250: {
		common.HexToAddress("0x03B82e4070cA32FF63A03F2EcfC16c0165689a9d"), // Test deployment - AVAX
	},
	8453: {},
	42161: {
		common.HexToAddress("0x5796698A29F3626c9FE13C4d3d3dEE987c84EBB3"), // Test deployment - Nothing
		common.HexToAddress("0x976a1C749cd8153909e0B04EebE931eF8957b15b"), // Test deployment - PHPTest
		common.HexToAddress("0xFa247d0D55a324ca19985577a2cDcFC383D87953"), // Test deployment - PHP
	},
}

// EXTRA_VAULTS is a list of vaults that are not registered in the registries, but are still used by Yearn
var EXTRA_VAULTS = map[uint64][]models.TVaultsFromRegistry{
	1:   {},
	10:  {},
	137: {},
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
	8453:  {},
	42161: {},
}

// EXTRA_TOKENS contans some extra token we may want to check for various reasons
var EXTRA_TOKENS = map[uint64][]common.Address{
	1: {
		common.HexToAddress("0x34fe2a45D8df28459d7705F37eD13d7aE4382009"), // yvWBTC
		common.HexToAddress("0xD533a949740bb3306d119CC777fa900bA034cd52"), // CRV - used by yBribe UI
		common.HexToAddress("0x090185f2135308BaD17527004364eBcC2D37e5F6"), // Spell - used by yBribe UI
		common.HexToAddress("0xCdF7028ceAB81fA0C6971208e83fa7872994beE5"), // TNT - used by yBribe UI
		common.HexToAddress("0xba100000625a3754423978a60c9317c58a424e3D"), // BAL - used by yBAL UI
		common.HexToAddress("0x5c6Ee304399DBdB9C8Ef030aB642B10820DB8F56"), // BALWETH - used by yBAL UI
		common.HexToAddress(`0x30D20208d987713f46DFD34EF128Bb16C404D10f`), // Stader
		common.HexToAddress(`0xa2E3356610840701BDf5611a53974510Ae27E2e1`), // wBETH
		common.HexToAddress(`0xac3E018457B222d93114458476f3E3416Abbe38F`), // Staked Frax Ether
		common.HexToAddress(`0xf951E335afb289353dc249e82926178EaC7DEd78`), // Swell Network Ether
		common.HexToAddress(`0x7f39C581F595B53c5cb19bD0b3f8dA6c935E2Ca0`), // Wrapped liquid staked Ether 2.0
		common.HexToAddress(`0xA35b1B31Ce002FBF2058D22F30f95D405200A15b`), // Stader ETHx
		common.HexToAddress(`0xBe9895146f7AF43049ca1c1AE358B0541Ea49704`), // Coinbase Wrapped Staked ETH

	},
	10: {
		common.HexToAddress("0x4200000000000000000000000000000000000042"), // Opt
	},
	137:  {},
	250:  {},
	8453: {},
	42161: {
		common.HexToAddress(`0x82e3A8F066a6989666b031d916c43672085b1582`), // YFI
		common.HexToAddress(`0x11cDb42B0EB46D95f990BeDD4695A6e3fA034978`), // CRV
	},
}

// IGNORED_TOKENS contans some tokens we may want to ignore for various reasons
var IGNORED_TOKENS = map[uint64][]common.Address{
	1: {
		common.HexToAddress(`0x7AB4a7BE740131BdE216521B54ADddD672F44A05`), // nothing
		common.HexToAddress(`0x61f46C65E403429266e8b569F23f70dD75d9BeE7`), // Old lp-yCRV
		common.HexToAddress(`0x8a0889d47f9Aa0Fac1cC718ba34E26b867437880`), // Old st-yCRV
		common.HexToAddress(`0x4c1317326fD8EFDeBdBE5e1cd052010D97723bd6`), // Another old st-yCRV
		common.HexToAddress(`0x2E919d27D515868f3D5Bc9110fa738f9449FC6ad`), // Old yvCurve-yveCRV pool
		common.HexToAddress(`0x7E46fd8a30869aa9ed55af031067Df666EfE87da`), // Old yvecrv-f
		common.HexToAddress(`0x5904BAcE7a9cCab585242e9d22f67C9f2F1BF7E2`), // nothing
		common.HexToAddress(`0x0309A528bBa0394dC4A2Ce59123C52E317A54604`), // Old yCRV-f
		common.HexToAddress(`0xBF7AA989192b020a8d3e1C65a558e123834325cA`), // Irrelevant HBTC yVault
		common.HexToAddress(`0xe92AE2cF5b373c1713eB5855D4D3aF81D8a8aCAE`), // Curve Stax Frax/Temple xLP + LP yVault - Unlisted
		common.HexToAddress(`0x3883f5e181fccaF8410FA61e12b59BAd963fb645`), // Theta: Old Token
		common.HexToAddress(`0xC4C319E2D4d66CcA4464C0c2B32c9Bd23ebe784e`), // Hacked alETH
		common.HexToAddress(`0x718AbE90777F5B778B52D553a5aBaa148DD0dc5D`), // Hacked vault
		common.HexToAddress("0x662fBF2c1E4b04342EeBA6371ec1C7420042B86F"), // Test deployment - Nothing
		common.HexToAddress("0x9C13e225AE007731caA49Fd17A41379ab1a489F4"), // Test deployment - Nothing
		common.HexToAddress("0x4Fdd1B06eF986238446be0F3EA163C1b6Fe28cC1"), // GUSD yVault - Not Yearn - PPS 100
	},
	10: {
		common.HexToAddress("0x6884bd538Db61A626Da0a05E10807BFC5Aea2b32"), // Test deployment - Nothing
		common.HexToAddress("0xDB8bBF2b0e28721F9BAc603e687E39bcF52201f8"), // Test deployment - Nothing
		common.HexToAddress("0xed5D83bB6Af23bcb05C144DC816f45A389d622a0"), // Test deployment - Nothing
	},
	137: {},
	250: {
		common.HexToAddress("0x03B82e4070cA32FF63A03F2EcfC16c0165689a9d"), // Test deployment - AVAX
	},
	8453: {},
	42161: {
		common.HexToAddress("0x5796698A29F3626c9FE13C4d3d3dEE987c84EBB3"), // Test deployment - Nothing
		common.HexToAddress("0x976a1C749cd8153909e0B04EebE931eF8957b15b"), // Test deployment - PHPTest
		common.HexToAddress("0xFa247d0D55a324ca19985577a2cDcFC383D87953"), // Test deployment - PHP
	},
}

// CHAIN_COIN is a list of vaults that are not registered in the registries, but are still used by Yearn
var CHAIN_COIN = map[uint64]models.TERC20Token{
	1: {
		Address:                   DEFAULT_COIN_ADDRESS,
		UnderlyingTokensAddresses: []common.Address{},
		Type:                      models.TokenTypeNative,
		Name:                      `Ether`,
		Symbol:                    `ETH`,
		DisplayName:               `Ether`,
		DisplaySymbol:             `ETH`,
		Description:               `Ether is the native cryptocurrency for the Ethereum blockchain.`,
		Icon:                      BASE_ASSET_URL + strconv.FormatUint(1, 10) + `/` + DEFAULT_COIN_ADDRESS.Hex() + `/logo-128.png`,
		Decimals:                  18,
		ChainID:                   1,
	},
	10: {
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
	137: {
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
	250: {
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
	8453: {
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
	42161: {
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
}
