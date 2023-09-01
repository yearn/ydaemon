package prices

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/common/addresses"
	"github.com/yearn/ydaemon/common/bigNumber"
	"github.com/yearn/ydaemon/common/contracts"
	"github.com/yearn/ydaemon/common/ethereum"
	"github.com/yearn/ydaemon/common/helpers"
	"github.com/yearn/ydaemon/common/logs"
)

var AERO_SUGAR_ADDRESS = common.HexToAddress(`0x2073d8035bb2b0f2e85aaf5a8732c6f397f9ff9b`)
var AERO_SUGAR_ORACLE_ADDRESS = common.HexToAddress(`0xB98fB4C9C99dE155cCbF5A14af0dBBAd96033D6f`)
var BASE_WETH_ADDRESS = common.HexToAddress(`0x4200000000000000000000000000000000000006`)
var BASE_USDBC_ADDRESS = common.HexToAddress(`0xd9aAEc86B65D86f6A7B5B1b0c42FFA531710b6CA`)

// fetchPricesFromAeroSugar is used to fetch prices from the sugar APR (velo).
func fetchPricesFromAeroSugar(chainID uint64, blockNumber *uint64, tokens []common.Address) map[common.Address]*bigNumber.Int {
	newPriceMap := make(map[common.Address]*bigNumber.Int)
	if chainID != 8453 {
		return newPriceMap
	}

	/**********************************************************************************************
	** The first step is to prepare the multicall, connecting to the multicall instance and
	** preparing the array of calls to send. All calls for all tokens will be send in a single
	** multicall and will later be accessible via a concatened string `tokenAddress + methodName`.
	**********************************************************************************************/
	client := ethereum.GetRPC(chainID)
	sugar, _ := contracts.NewVeloSugarCaller(AERO_SUGAR_ADDRESS, client)
	sugarOracle, _ := contracts.NewVeloSugarOracleCaller(AERO_SUGAR_ORACLE_ADDRESS, client)
	allSugar, _ := sugar.All(nil, big.NewInt(200), big.NewInt(0), common.Address{})

	logs.Pretty(len(allSugar))
	for _, pair := range allSugar {
		ratesConnector := []common.Address{pair.Token0, pair.Token1, BASE_WETH_ADDRESS, BASE_USDBC_ADDRESS}
		prices, err := sugarOracle.GetManyRatesWithConnectors(nil, 2, ratesConnector)
		if err != nil {
			logs.Error("Error fetching prices from sugar oracle", err)
			continue
		}
		token0Price := bigNumber.SetInt(prices[0])
		if token0Price.IsZero() && addresses.Equals(pair.Token0, BASE_USDBC_ADDRESS) {
			token0Price = bigNumber.NewInt(1e18)
		}
		token1Price := bigNumber.SetInt(prices[1])
		if token1Price.IsZero() && addresses.Equals(pair.Token1, BASE_USDBC_ADDRESS) {
			token1Price = bigNumber.NewInt(1e18)
		}
		token0Contract, err := contracts.NewERC20Caller(pair.Token0, client)
		if err != nil {
			logs.Error("Error fetching token0 contract", err)
			continue
		}
		token0Decimals, err := token0Contract.Decimals(nil)
		if err != nil {
			logs.Error("Error fetching token0 decimals", err)
			continue
		}
		token1Contract, err := contracts.NewERC20Caller(pair.Token1, client)
		if err != nil {
			logs.Error("Error fetching token1 contract", err)
			continue
		}
		token1Decimals, err := token1Contract.Decimals(nil)
		if err != nil {
			logs.Error("Error fetching token1 decimals", err)
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
