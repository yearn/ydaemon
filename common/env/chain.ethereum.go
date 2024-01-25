package env

import (
	"math"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/internal/models"
)

var ETHEREUM = TChain{
	ID:              1,
	RpcURI:          `https://eth.public-rpc.com`,
	SubgraphURI:     `https://api.thegraph.com/subgraphs/name/rareweasel/yearn-vaults-v2-subgraph-mainnet`,
	EtherscanURI:    `https://api.etherscan.io/api`,
	MaxBlockRange:   100_000_000,
	MaxBatchSize:    math.MaxInt64,
	AvgBlocksPerDay: 7150,
	YBribeV3Contract: TContractData{
		Address: common.HexToAddress(`0x03dFdBcD4056E2F92251c7B07423E1a33a7D3F6d`),
		Block:   15878262,
	},
	LensContract: TContractData{
		Address: common.HexToAddress(`0x83d95e0D5f402511dB06817Aff3f9eA88224B030`),
		Block:   12242339,
	},
	MulticallContract: TContractData{
		Address: common.HexToAddress(`0x5ba1e12693dc8f9c48aad8770482f4739beed696`),
		Block:   12336033,
	},
	PartnerContract: TContractData{
		Address: common.HexToAddress(`0x8ee392a4787397126C163Cb9844d7c447da419D8`),
		Block:   14166636,
	},
	APROracleContract: TContractData{
		Address: common.HexToAddress(`0x27aD2fFc74F74Ed27e1C0A19F1858dD0963277aE`),
		Block:   19070394,
	},
	Coin: models.TERC20Token{
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
	Registries: []TContractData{
		{
			Address: common.HexToAddress("0xe15461b18ee31b7379019dc523231c57d1cbc18c"),
			Version: 1,
			Block:   11563389,
		},
		{
			Address: common.HexToAddress("0x50c1a2eA0a861A967D9d0FFE2AE4012c2E053804"),
			Version: 2,
			Block:   12045555,
		},
		{
			Address: common.HexToAddress("0xaF1f5e1c19cB68B30aAD73846eFfDf78a5863319"),
			Version: 3,
			Block:   16215519,
		},
		{
			Address: common.HexToAddress("0xff31A1B020c868F6eA3f61Eb953344920EeCA3af"),
			Version: 4,
			Block:   19072527,
		},
		{
			Address: common.HexToAddress("0xE9E8C89c8Fc7E8b8F23425688eb68987231178e5"),
			Version: 5,
			Block:   18385809,
		},
	},
	ExtraVaults: []models.TVaultsFromRegistry{
		{
			//yvAjnaDAI-harvester (ysyvAjnaDAI)
			ChainID:         1,
			Address:         common.HexToAddress(`0x082a5743aAdf3d0Daf750EeF24652b36a68B1e9C`),
			RegistryAddress: common.HexToAddress(`0xE9E8C89c8Fc7E8b8F23425688eb68987231178e5`),
			TokenAddress:    common.HexToAddress(`0xe24BA27551aBE96Ca401D39761cA2319Ea14e3CB`),
			APIVersion:      `3.0.1`,
			BlockNumber:     18981463,
			Type:            models.TTokenType(models.TokenTypeStandardVault),
			Kind:            models.TVaultKind(models.VaultKindMultiple),
		},
	},
	BlacklistedVaults: []common.Address{
		common.HexToAddress("0x662fBF2c1E4b04342EeBA6371ec1C7420042B86F"), // Test deployment - Nothing
		common.HexToAddress("0x9C13e225AE007731caA49Fd17A41379ab1a489F4"), // Test deployment - Nothing
		common.HexToAddress("0xBF7AA989192b020a8d3e1C65a558e123834325cA"), // HBTC yVault - Not Yearn - PPS 0
		common.HexToAddress("0x4Fdd1B06eF986238446be0F3EA163C1b6Fe28cC1"), // GUSD yVault - Not Yearn - PPS 100
		common.HexToAddress("0x8a0889d47f9Aa0Fac1cC718ba34E26b867437880"), // Old st-yCRV
		common.HexToAddress("0x61f46C65E403429266e8b569F23f70dD75d9BeE7"), // Old lp-yCRV
		//
		common.HexToAddress("0x98E86Ed5b0E48734430BfBe92101156C75418cad"), // yearn BAL - Disabled for now
		common.HexToAddress("0xc09cfb625e586B117282399433257a1C0841edf3"), // Staked Yearn BAL Vault - Disabled for now
		common.HexToAddress("0xD725F5742047B4B4A3110D0b38284227fcaB041e"), // LP Yearn BAL Vault - Disabled for now
		common.HexToAddress("0xD61e198e139369a40818FE05F5d5e6e045Cd6eaF"), // Balancer yBAL Stable Pool - Disabled for now
		//
		common.HexToAddress("0xe1d0549702C0beC3d3F918159449542738e6F942"), // Test V3 Vault
		//
		// common.HexToAddress(`0x39CAF13a104FF567f71fd2A4c68C026FDB6E740B`), //DebtRatio 0 - To retire
		// common.HexToAddress(`0x4560b99C904aAD03027B5178CCa81584744AC01f`), //DebtRatio 0 - To retire
		// common.HexToAddress(`0x5c0A86A32c129538D62C106Eb8115a8b02358d57`), //DebtRatio 0 - To retire
		// common.HexToAddress(`0x2D5D4869381C4Fce34789BC1D38aCCe747E295AE`), //DebtRatio 0 - To retire
		// common.HexToAddress(`0xd88dBBA3f9c4391Ee46f5FF548f289054db6E51C`), //DebtRatio 0 - To retire
		// common.HexToAddress(`0xC4dAf3b5e2A9e93861c3FBDd25f1e943B8D87417`), //DebtRatio 0 - To retire
		// common.HexToAddress(`0x801Ab06154Bf539dea4385a39f5fa8534fB53073`), //DebtRatio 0 - To retire
		// common.HexToAddress(`0xF59D66c1d593Fb10e2f8c2a6fD2C958792434B9c`), //DebtRatio 0 - To retire
		// common.HexToAddress(`0x341bb10D8f5947f3066502DC8125d9b8949FD3D6`), //DebtRatio 0 - To retire
		// common.HexToAddress(`0xb4D1Be44BfF40ad6e506edf43156577a3f8672eC`), //DebtRatio 0 - To retire
		// common.HexToAddress(`0x7AB4a7BE740131BdE216521B54ADddD672F44A05`), //DebtRatio 0 - To retire
		// common.HexToAddress(`0x59518884EeBFb03e90a18ADBAAAB770d4666471e`), //DebtRatio 0 - To retire
		// common.HexToAddress(`0xdb25cA703181E7484a155DD612b06f57E12Be5F0`), //DebtRatio 0 - To retire
		// common.HexToAddress(`0x1b905331F7dE2748F4D6a0678e1521E20347643F`), //DebtRatio 0 - To retire
		// common.HexToAddress(`0x9d409a0A012CFbA9B15F6D4B36Ac57A46966Ab9a`), //DebtRatio 0 - To retire
	},
	ExtraTokens: []common.Address{
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
		common.HexToAddress(`0x583019fF0f430721aDa9cfb4fac8F06cA104d0B4`), // st-yETH
		common.HexToAddress(`0x41252E8691e964f7DE35156B68493bAb6797a275`), // dYFI
		common.HexToAddress(`0xd084944d3c05cd115c09d072b9f44ba3e0e45921`), // Fold
		common.HexToAddress(`0xdA47862a83dac0c112BA89c6abC2159b95afd71C`), // Prisma
		common.HexToAddress(`0xe3668873d944e4a949da05fc8bde419eff543882`), // yPrisma
		common.HexToAddress(`0x69833361991ed76f9e8DBBcdf9ea1520fEbFb4a7`), // lp-Prisma
		common.HexToAddress(`0x9a96ec9B57Fb64FbC60B423d1f4da7691Bd35079`), // Ajna
	},
	IgnoredTokens: []common.Address{
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
		//
		common.HexToAddress("0x98E86Ed5b0E48734430BfBe92101156C75418cad"), // yearn BAL - Disabled for now
		common.HexToAddress("0xc09cfb625e586B117282399433257a1C0841edf3"), // Staked Yearn BAL Vault - Disabled for now
		common.HexToAddress("0xD725F5742047B4B4A3110D0b38284227fcaB041e"), // LP Yearn BAL Vault - Disabled for now
		common.HexToAddress("0xD61e198e139369a40818FE05F5d5e6e045Cd6eaF"), // Balancer yBAL Stable Pool - Disabled for now
	},
	Curve: TChainCurve{
		RegistryAddress: common.HexToAddress(`0x90E00ACe148ca3b23Ac1bC8C240C2a7Dd9c2d7f5`),
		FactoryAddress:  common.HexToAddress(`0xF18056Bbd320E96A48e3Fbf8bC061322531aac99`),
		FactoryURIs: []string{
			`https://api.curve.fi/api/getPools/ethereum/factory`,
			`https://api.curve.fi/api/getPools/ethereum/factory-crypto`,
		},
		PoolsURIs: []string{
			`https://api.curve.fi/api/getPools/ethereum/main`,
			`https://api.curve.fi/api/getPools/ethereum/crypto`,
			`https://api.curve.fi/api/getPools/ethereum/factory`,
			`https://api.curve.fi/api/getPools/ethereum/factory-crypto`,
		},
	},
}
