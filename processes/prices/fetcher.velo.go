package prices

import (
	"encoding/json"
	"io/ioutil"
	"math/big"
	"net/http"

	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/common/addresses"
	"github.com/yearn/ydaemon/common/bigNumber"
	"github.com/yearn/ydaemon/common/contracts"
	"github.com/yearn/ydaemon/common/ethereum"
	"github.com/yearn/ydaemon/common/helpers"
	"github.com/yearn/ydaemon/common/logs"
	"github.com/yearn/ydaemon/internal/multicalls"
)

var VELO_PAIR_URI = `https://api.velodrome.finance/api/v1/pairs`
var VELO_SUGAR_ADDRESS = common.HexToAddress(`0x4D996E294B00cE8287C16A2b9A4e637ecA5c939f`)
var VELO_SUGAR_ORACLE_ADDRESS = common.HexToAddress(`0x395942c2049604a314d39f370dfb8d87aac89e16`)
var OPT_WETH_ADDRESS = common.HexToAddress(`0x4200000000000000000000000000000000000006`)
var OPT_OP_ADDRESS = common.HexToAddress(`0x4200000000000000000000000000000000000042`)
var OPT_USDC_ADDRESS = common.HexToAddress(`0x7F5c764cBc14f9669B88837ca1490cCa17c31607`)

func fetchVelo(url string) []TVeloPairData {
	resp, err := http.Get(url)
	if err != nil {
		logs.Error(`impossible to get velo URL`, err.Error())
		return []TVeloPairData{}
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logs.Error(`impossible to read velo Get body`, err.Error())
		return []TVeloPairData{}
	}
	var factories TVeloPairs
	if err := json.Unmarshal(body, &factories); err != nil {
		logs.Error(`impossible to unmarshal velo Get body`, err.Error())
		return []TVeloPairData{}
	}
	return factories.Data
}

// getPricesFromVeloPairsAPI is used after setting the prices from the multicall, aka the lens contract.
// Some missing prices may exists and we can try to call the VELO API to get them
func getPricesFromVeloPairsAPI(chainID uint64) map[common.Address]*bigNumber.Int {
	newPriceMap := make(map[common.Address]*bigNumber.Int)
	if chainID != 10 {
		return newPriceMap
	}
	veloPairs := fetchVelo(VELO_PAIR_URI)

	// For each pool, we calculate the price per token and assign it to the token
	// if the Store price is 0
	for _, pair := range veloPairs {
		coins := []TVeloToken{}
		coins = append(coins, pair.Token0)
		coins = append(coins, pair.Token1)
		for _, bribes := range pair.Gauge.Bribes {
			coins = append(coins, bribes.Token)
		}
		for _, coin := range coins {
			coinAddress := common.HexToAddress(coin.Address)
			coinPrice := bigNumber.NewFloat(coin.Price)
			coinPrice = coinPrice.Mul(coinPrice, bigNumber.NewFloat(1e6))
			coinPriceBigInt := coinPrice.Int()
			newPriceMap[coinAddress] = coinPriceBigInt
		}

		if pair.TotalSupply > 0 {
			if (pair.Token0.Price == 0) || (pair.Reserve1 == 0) {
				continue
			}

			priceOfWrapper := (pair.Token1.Price * pair.Reserve1) / (pair.Token0.Price * pair.Reserve0)
			wrapperPrice := bigNumber.NewFloat(priceOfWrapper)
			wrapperPrice = wrapperPrice.Mul(wrapperPrice, bigNumber.NewFloat(1e6))
			wrapperPriceBigInt := wrapperPrice.Int()
			newPriceMap[common.HexToAddress(pair.Address)] = wrapperPriceBigInt
		}

	}
	return newPriceMap
}

// fetchPricesFromSugar is used to fetch prices from the sugar APR (velo).
func fetchPricesFromSugar(chainID uint64, blockNumber *uint64, tokens []common.Address) map[common.Address]*bigNumber.Int {
	newPriceMap := make(map[common.Address]*bigNumber.Int)
	if chainID != 10 {
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
	sugar, _ := contracts.NewVeloSugarCaller(VELO_SUGAR_ADDRESS, client)
	// sugarOracle, _ := contracts.NewVeloSugarOracleCaller(VELO_SUGAR_ORACLE_ADDRESS, client)
	allSugar, _ := sugar.All(nil, big.NewInt(10000), big.NewInt(0), common.Address{})

	for _, pair := range allSugar {
		ratesConnector := []common.Address{pair.Token0, pair.Token1, OPT_WETH_ADDRESS, OPT_OP_ADDRESS, OPT_USDC_ADDRESS}
		calls := []ethereum.Call{}
		calls = append(calls, multicalls.GetManyRatesWithConnectors(VELO_SUGAR_ORACLE_ADDRESS.Hex(), VELO_SUGAR_ORACLE_ADDRESS, 2, ratesConnector))
		calls = append(calls, multicalls.GetDecimals(pair.Token0.Hex(), pair.Token0))
		calls = append(calls, multicalls.GetDecimals(pair.Token1.Hex(), pair.Token1))
		response := multicalls.Perform(chainID, calls, nil)

		prices := helpers.DecodeBigInts(response[VELO_SUGAR_ORACLE_ADDRESS.Hex()+`getManyRatesWithConnectors`])
		token0Decimals := helpers.DecodeUint64(response[pair.Token0.Hex()+`decimals`])
		token1Decimals := helpers.DecodeUint64(response[pair.Token1.Hex()+`decimals`])
		token0Price := prices[0]
		token1Price := prices[1]

		if token0Price.IsZero() && addresses.Equals(pair.Token0, OPT_USDC_ADDRESS) {
			token0Price = bigNumber.NewInt(1e18)
		}
		if token1Price.IsZero() && addresses.Equals(pair.Token1, OPT_USDC_ADDRESS) {
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
			newPairPriceMap[pair.PairAddress] = pairPrice.Int()
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
