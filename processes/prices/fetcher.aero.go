package prices

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/common/addresses"
	"github.com/yearn/ydaemon/common/bigNumber"
	"github.com/yearn/ydaemon/common/contracts"
	"github.com/yearn/ydaemon/common/ethereum"
	"github.com/yearn/ydaemon/common/helpers"
	"github.com/yearn/ydaemon/internal/multicalls"
	"github.com/yearn/ydaemon/internal/storage"
)

var AERO_SUGAR_ADDRESS = common.HexToAddress(`0x2073D8035bB2b0F2e85aAF5a8732C6f397F9ff9b`)
var AERO_SUGAR_ORACLE_ADDRESS = common.HexToAddress(`0xB98fB4C9C99dE155cCbF5A14af0dBBAd96033D6f`)
var BASE_WETH_ADDRESS = common.HexToAddress(`0x4200000000000000000000000000000000000006`)
var BASE_USDC_ADDRESS = common.HexToAddress(`0x833589fCD6eDb6E08f4c7C32D4f71b54bdA02913`)

var BASE_RATE_CONNECTORS = []common.Address{
	common.HexToAddress(`0xbf1aea8670d2528e08334083616dd9c5f3b087ae`),
	common.HexToAddress(`0xe3b53af74a4bf62ae5511055290838050bf764df`),
	common.HexToAddress(`0xf544251d25f3d243a36b07e7e7962a678f952691`),
	common.HexToAddress(`0x4a3a6dd60a34bb2aba60d73b4c88315e9ceb6a3d`),
	common.HexToAddress(`0xc5102fe9359fd9a28f877a67e36b0f050d81a3cc`),
	common.HexToAddress(`0x65a2508c429a6078a7bc2f7df81ab575bd9d9275`),
	common.HexToAddress(`0xb79dd08ea68a908a97220c76d19a6aa9cbde4376`),
	common.HexToAddress(`0xde5ed76e7c05ec5e4572cfc88d1acea165109e44`),
	common.HexToAddress(`0x50c5725949a6f0c72e6c4a641f24049a917db0cb`),
	common.HexToAddress(`0x940181a94a35a4569e4529a3cdfb74e38fd98631`),
	common.HexToAddress(`0x9e53e88dcff56d3062510a745952dec4cefdff9e`),
	common.HexToAddress(`0xba5e6fa2f33f3955f0cef50c63dcc84861eab663`),
	common.HexToAddress(`0x8901cb2e82cc95c01e42206f8d1f417fe53e7af0`),
	common.HexToAddress(`0x2ae3f1ec7f1f5012cfeab0185bfc7aa3cf0dec22`),
	common.HexToAddress(`0x9483ab65847a447e36d21af1cab8c87e9712ff93`),
	common.HexToAddress(`0x74ccbe53f77b08632ce0cb91d3a545bf6b8e0979`),
	common.HexToAddress(`0xff8adec2221f9f4d8dfbafa6b9a297d17603493d`),
	common.HexToAddress(`0xf34e0cff046e154cafcae502c7541b9e5fd8c249`),
	common.HexToAddress(`0xa61beb4a3d02decb01039e378237032b351125b4`),
	common.HexToAddress(`0x22a2488fe295047ba13bd8cccdbc8361dbd8cf7c`),
	common.HexToAddress(`0xc142171b138db17a1b7cb999c44526094a4dae05`),
	common.HexToAddress(`0x12063cc18a7096d170e5fc410d8623ad97ee24b3`),
	common.HexToAddress(`0xc19669a405067927865b40ea045a2baabbbe57f5`),
	common.HexToAddress(`0x9cbd543f1b1166b2df36b68eb6bb1dce24e6abdf`),
	common.HexToAddress(`0x9cc2fc2f75768b0307925c7935396ec9d94bba44`),
	common.HexToAddress(`0x8ae125e8653821e851f12a49f7765db9a9ce7384`),
	common.HexToAddress(`0x833589fcd6edb6e08f4c7c32d4f71b54bda02913`),
	common.HexToAddress(`0x96e890c6b2501a69cad5dba402bfb871a2a2874c`),
	common.HexToAddress(`0xeb466342c4d449bc9f53a865d5cb90586f405215`),
	common.HexToAddress(`0xa3d1a8deb97b111454b294e2324efad13a9d8396`),
	common.HexToAddress(`0x833589fcd6edb6e08f4c7c32d4f71b54bda02913`),
	common.HexToAddress(`0x940181a94a35a4569e4529a3cdfb74e38fd98631`),
	common.HexToAddress(`0x50c5725949a6f0c72e6c4a641f24049a917db0cb`),
	common.HexToAddress(`0x4621b7a9c75199271f773ebd9a499dbd165c3191`),
	common.HexToAddress(`0x4200000000000000000000000000000000000006`),
	common.HexToAddress(`0xb79dd08ea68a908a97220c76d19a6aa9cbde4376`),
	common.HexToAddress(`0xf7a0dd3317535ec4f4d29adf9d620b3d8d5d5069`),
	common.HexToAddress(`0xd9aaec86b65d86f6a7b5b1b0c42ffa531710b6ca`),
}

// var BASE_BASE_ADDRESS = common.HexToAddress(`0x4200000000000000000000000000000000000042`)

// fetchPricesFromAeroSugar is used to fetch prices from the sugar APR (Base).
func fetchPricesFromAeroSugar(chainID uint64, blockNumber *uint64, tokens []common.Address) map[common.Address]*bigNumber.Int {
	newPriceMap := make(map[common.Address]*bigNumber.Int)
	if chainID != 8453 {
		return newPriceMap
	}
	newPairPriceMap := make(map[common.Address]*bigNumber.Int)
	newTokensPriceMap := make(map[common.Address]*bigNumber.Int)

	/**********************************************************************************************
	** The first step is to prepare the multicall, connecting to the multicall instance and
	** preparing the array of calls to send. All calls for all tokens will be send in a single
	** multicall and will later be accessible via a concatened string `tokenAddress + methodName`.
	**********************************************************************************************/
	client := ethereum.GetRPC(chainID)
	sugar, _ := contracts.NewAeroSugar(AERO_SUGAR_ADDRESS, client)
	// sugarOracle, _ := contracts.NewVeloSugarOracleCaller(AERO_SUGAR_ORACLE_ADDRESS, client)
	allSugar, _ := sugar.All(nil, big.NewInt(1000), big.NewInt(0), common.Address{})

	for _, pair := range allSugar {
		ratesConnector := []common.Address{pair.Token0, pair.Token1}
		ratesConnector = append(ratesConnector, BASE_RATE_CONNECTORS...)
		calls := []ethereum.Call{}
		calls = append(calls, multicalls.GetManyRatesWithConnectors(AERO_SUGAR_ORACLE_ADDRESS.Hex(), AERO_SUGAR_ORACLE_ADDRESS, 2, ratesConnector))
		calls = append(calls, multicalls.GetDecimals(pair.Token0.Hex(), pair.Token0))
		calls = append(calls, multicalls.GetDecimals(pair.Token1.Hex(), pair.Token1))
		response := multicalls.Perform(chainID, calls, nil)

		prices := helpers.DecodeBigInts(response[AERO_SUGAR_ORACLE_ADDRESS.Hex()+`getManyRatesWithConnectors`])
		token0Decimals := helpers.DecodeUint64(response[pair.Token0.Hex()+`decimals`])
		token1Decimals := helpers.DecodeUint64(response[pair.Token1.Hex()+`decimals`])
		token0Price := prices[0]
		token1Price := prices[1]

		if token0Price.IsZero() && addresses.Equals(pair.Token0, BASE_USDC_ADDRESS) {
			token0Price = bigNumber.NewInt(1e18)
		}
		if token1Price.IsZero() && addresses.Equals(pair.Token1, BASE_USDC_ADDRESS) {
			token1Price = bigNumber.NewInt(1e18)
		}

		token0PriceUSD := helpers.ToNormalizedAmount(token0Price, 18)
		token1PriceUSD := helpers.ToNormalizedAmount(token1Price, 18)
		token0totalSupply := helpers.ToNormalizedAmount(bigNumber.SetInt(pair.Reserve0), uint64(token0Decimals))
		token1totalSupply := helpers.ToNormalizedAmount(bigNumber.SetInt(pair.Reserve1), uint64(token1Decimals))
		pairTotalSupply := helpers.ToNormalizedAmount(bigNumber.SetInt(pair.TotalSupply), 18)

		token0ValueInPair := bigNumber.NewFloat(0).Mul(token0PriceUSD, token0totalSupply)
		token1ValueInPair := bigNumber.NewFloat(0).Mul(token1PriceUSD, token1totalSupply)
		valueInPair := bigNumber.NewFloat(0).Add(token0ValueInPair, token1ValueInPair)

		pairPrice := bigNumber.NewFloat(0).Div(valueInPair, pairTotalSupply)
		pairPrice = bigNumber.NewFloat(0).Mul(pairPrice, bigNumber.NewFloat(1e6))
		token0Price = bigNumber.NewFloat(0).Mul(token0PriceUSD, bigNumber.NewFloat(1e6)).Int()
		token1Price = bigNumber.NewFloat(0).Mul(token1PriceUSD, bigNumber.NewFloat(1e6)).Int()

		if pairToken, _ := storage.GetERC20(chainID, pair.Lp); !pairToken.IsVaultLike() {
			if !pairPrice.IsZero() {
				newPairPriceMap[pair.Lp] = pairPrice.Int()
			}
		}

		if token0, _ := storage.GetERC20(chainID, pair.Token0); !token0.IsVaultLike() {
			if !token0Price.IsZero() {
				newTokensPriceMap[pair.Token0] = token0Price
			}
		}

		if token1, _ := storage.GetERC20(chainID, pair.Token1); !token1.IsVaultLike() {
			if !token1Price.IsZero() {
				newTokensPriceMap[pair.Token1] = token1Price
			}
		}
	}

	for token, price := range newPairPriceMap {
		newPriceMap[token] = price
	}
	for token, price := range newTokensPriceMap {
		if _, ok := newPriceMap[token]; !ok {
			newPriceMap[token] = price
		}
	}

	return newPriceMap
}
