package prices

import (
	"encoding/json"
	"strconv"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/common/bigNumber"
	"github.com/yearn/ydaemon/common/env"
	"github.com/yearn/ydaemon/common/logs"
	"github.com/yearn/ydaemon/common/store"
)

type TCurveFactoriesPoolData struct {
	Name        string  `json:"name"`
	Symbol      string  `json:"symbol"`
	Address     string  `json:"address"`
	LPAddress   string  `json:"lpTokenAddress"`
	TotalSupply string  `json:"totalSupply"`
	USDTotal    float64 `json:"usdTotal"`
	Coins       []struct {
		Address  string          `json:"address"`
		Decimals json.RawMessage `json:"decimals"`
		Symbol   string          `json:"symbol"`
		USDPrice float64         `json:"usdPrice"`
	} `json:"coins"`
}

type TCurveFactories struct {
	Data struct {
		PoolData []TCurveFactoriesPoolData `json:"poolData"`
	} `json:"data"`
}

type TLlamaPriceData struct {
	Decimals int     `json:"decimals"`
	Price    float64 `json:"price"`
	Symbol   string  `json:"symbol"`
}

type TLlamaPrice struct {
	Coins map[string]TLlamaPriceData `json:"coins"`
}

type TGeckoPrice map[string]struct {
	USDPrice float64 `json:"usd"`
}

/**********************************************************************************************
** Set of functions to store and retrieve the prices from the cache and/or database and being
** able to access them from the rest of the application.
** The _priceMap variable is not exported and is only used internally by the functions below.
**********************************************************************************************/
var _priceMap = make(map[uint64]map[common.Address]*bigNumber.Int)

/**********************************************************************************************
** The _historicalPriceMap variable is not exported and is only used internally by the
** functions below. It is used to store the historical prices of the tokens:
** map[chainID]map[blockNumber]map[tokenAddress]price
**********************************************************************************************/
var _historicalPriceMap = sync.Map{}

/**************************************************************************************************
** LoadHistoricalPrices will retrieve the historical prices values from the local DB and store
** them in the _historicalPriceMap.
**************************************************************************************************/
func LoadHistoricalPrices(chainID uint64, wg *sync.WaitGroup) {
	defer wg.Done()
	temp := make(map[string]string)
	store.Iterate(chainID, store.TABLES.HISTORICAL_PRICES, &temp)
	if temp != nil && (len(temp) > 0) {
		for key, value := range temp {
			_historicalPriceMap.Store(key, bigNumber.NewInt(0).SetString(value))
		}
	}
}

/**********************************************************************************************
** MapPrices will, for a given chainID, return a map of prices.
** The map will be of the form: map[vaultAddress]bigNumber.Int
**********************************************************************************************/
func MapPrices(chainID uint64) map[common.Address]*bigNumber.Int {
	prices := make(map[common.Address]*bigNumber.Int)
	for key, price := range _priceMap[chainID] {
		prices[key] = price
	}
	return prices
}

/**********************************************************************************************
** FindPrice will, for a given chainID, try to find the price of the provided tokenAddress
** stored in _priceMap.
** It will return the price if found, and a boolean indicating if the token was found or not.
**********************************************************************************************/
func FindPrice(chainID uint64, tokenAddress common.Address) (*bigNumber.Int, bool) {
	price, ok := _priceMap[chainID][tokenAddress]
	if !ok {
		return nil, false
	}
	return price, true
}

/**********************************************************************************************
** FindPriceOnBlock will, for a given chainID, try to find the price of the provided
** tokenAddress for the provided blockNumber stored in _historicalPriceMap.
** It will return the price if found, and a boolean indicating if the token was found or not.
**********************************************************************************************/
func FindPriceOnBlock(chainID uint64, blockNumber uint64, tokenAddress common.Address) (*bigNumber.Int, bool) {
	key := strconv.FormatUint(chainID, 10) + "_" + strconv.FormatUint(blockNumber, 10) + "_" + tokenAddress.Hex()
	price, ok := _historicalPriceMap.Load(key)
	if !ok {
		logs.Info(`Fetching historical price for token`, tokenAddress.Hex(), `on block`, blockNumber)
		newPrice := FetchPricesOnBlock(chainID, blockNumber, []common.Address{tokenAddress})
		if newPrice == nil {
			return nil, false
		}
		price = newPrice[tokenAddress]
		_historicalPriceMap.Store(key, price)
	}
	return price.(*bigNumber.Int), true
}

/**********************************************************************************************
** StorePriceOnBlock will, for a given chainID, store the price of the provided tokenAddress
** for the provided blockNumber in _historicalPriceMap.
**********************************************************************************************/
func StorePriceOnBlock(chainID uint64, blockNumber uint64, tokenAddress common.Address, value *bigNumber.Int) {
	key := strconv.FormatUint(chainID, 10) + "_" + strconv.FormatUint(blockNumber, 10) + "_" + tokenAddress.Hex()
	_historicalPriceMap.Store(key, value)
	go store.SaveInDB(chainID, store.TABLES.HISTORICAL_PRICES, key, value.String())
}

func init() {
	wg := &sync.WaitGroup{}
	for _, chainID := range env.SUPPORTED_CHAIN_IDS {
		wg.Add(1)
		go LoadHistoricalPrices(chainID, wg)
	}
	wg.Wait()
}
