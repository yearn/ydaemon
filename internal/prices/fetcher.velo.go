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
	"github.com/yearn/ydaemon/common/traces"
	"github.com/yearn/ydaemon/internal/models"
)

var VELO_PAIR_URI = `https://api.velodrome.finance/api/v1/pairs`
var VELO_SUGAR_ADDRESS = common.HexToAddress(`0x3b21531bd00289f10c7d8b64b9389095f521a4d3`)
var VELO_SUGAR_ORACLE_ADDRESS = common.HexToAddress(`0x07f544813e9fb63d57a92f28fbd3ff0f7136f5ce`)
var OPT_WETH_ADDRESS = common.HexToAddress(`0x4200000000000000000000000000000000000006`)
var OPT_OP_ADDRESS = common.HexToAddress(`0x4200000000000000000000000000000000000042`)
var OPT_USDC_ADDRESS = common.HexToAddress(`0x7F5c764cBc14f9669B88837ca1490cCa17c31607`)

func fetchVelo(url string) []models.TVeloPairData {
	resp, err := http.Get(url)
	if err != nil {
		traces.
			Capture(`error`, `impossible to get velo URL`).
			SetEntity(`prices`).
			SetExtra(`error`, err.Error()).
			SetTag(`url`, url).
			Send()
		return []models.TVeloPairData{}
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		traces.
			Capture(`error`, `impossible to read velo Get body`).
			SetEntity(`prices`).
			SetExtra(`error`, err.Error()).
			SetTag(`url`, url).
			Send()
		return []models.TVeloPairData{}
	}
	var factories models.TVeloPairs
	if err := json.Unmarshal(body, &factories); err != nil {
		traces.
			Capture(`error`, `impossible to unmarshal velo Get body`).
			SetEntity(`prices`).
			SetExtra(`error`, err.Error()).
			SetTag(`url`, url).
			Send()
		return []models.TVeloPairData{}
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
		coins := []models.TVeloToken{}
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

	/**********************************************************************************************
	** The first step is to prepare the multicall, connecting to the multicall instance and
	** preparing the array of calls to send. All calls for all tokens will be send in a single
	** multicall and will later be accessible via a concatened string `tokenAddress + methodName`.
	**********************************************************************************************/
	client := ethereum.GetRPC(chainID)
	sugar, _ := contracts.NewVeloSugarCaller(VELO_SUGAR_ADDRESS, client)
	sugarOracle, _ := contracts.NewVeloSugarOracleCaller(VELO_SUGAR_ORACLE_ADDRESS, client)
	allSugar, _ := sugar.All(nil, big.NewInt(200), big.NewInt(0), common.Address{})

	for _, pair := range allSugar {
		ratesConnector := []common.Address{pair.Token0, pair.Token1, OPT_WETH_ADDRESS, OPT_OP_ADDRESS, OPT_USDC_ADDRESS}
		prices, err := sugarOracle.GetManyRatesWithConnectors(nil, 2, ratesConnector)
		if err != nil {
			continue
		}
		token0Price := bigNumber.SetInt(prices[0])
		if token0Price.IsZero() && addresses.Equals(pair.Token0, OPT_USDC_ADDRESS) {
			token0Price = bigNumber.NewInt(1e18)
		}
		token1Price := bigNumber.SetInt(prices[1])
		if token1Price.IsZero() && addresses.Equals(pair.Token1, OPT_USDC_ADDRESS) {
			token1Price = bigNumber.NewInt(1e18)
		}
		token0Contract, err := contracts.NewERC20Caller(pair.Token0, client)
		if err != nil {
			continue
		}
		token0Decimals, err := token0Contract.Decimals(nil)
		if err != nil {
			continue
		}
		token1Contract, err := contracts.NewERC20Caller(pair.Token1, client)
		if err != nil {
			continue
		}
		token1Decimals, err := token1Contract.Decimals(nil)
		if err != nil {
			continue
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
		newPriceMap[pair.PairAddress] = pairPrice.Int()
		newPriceMap[pair.Token0] = bigNumber.NewFloat(0).Mul(token0PriceUSD, bigNumber.NewFloat(1e6)).Int()
		newPriceMap[pair.Token1] = bigNumber.NewFloat(0).Mul(token1PriceUSD, bigNumber.NewFloat(1e6)).Int()
	}

	return newPriceMap
}
