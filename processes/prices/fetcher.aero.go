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
)

var AERO_SUGAR_ADDRESS = common.HexToAddress(`0x2073D8035bB2b0F2e85aAF5a8732C6f397F9ff9b`)
var AERO_SUGAR_ORACLE_ADDRESS = common.HexToAddress(`0xB98fB4C9C99dE155cCbF5A14af0dBBAd96033D6f`)
var BASE_WETH_ADDRESS = common.HexToAddress(`0x4200000000000000000000000000000000000006`)
var BASE_USDC_ADDRESS = common.HexToAddress(`0x833589fCD6eDb6E08f4c7C32D4f71b54bdA02913`)

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
		ratesConnector := []common.Address{pair.Token0, pair.Token1, BASE_WETH_ADDRESS, BASE_USDC_ADDRESS}
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

		if !pairPrice.IsZero() {
			newPairPriceMap[pair.Lp] = pairPrice.Int()
		}
		if !token0Price.IsZero() {
			newTokensPriceMap[pair.Token0] = token0Price
		}
		if !token1Price.IsZero() {
			newTokensPriceMap[pair.Token1] = token1Price
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
