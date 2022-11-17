package prices

import (
	"encoding/json"

	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/common/bigNumber"
	"github.com/yearn/ydaemon/common/types/common"
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

/**********************************************************************************************
** Set of functions to store and retrieve the prices from the cache and/or database and being
** able to access them from the rest of the application.
** The _priceMap variable is not exported and is only used internally by the functions below.
**********************************************************************************************/
var _priceMap = make(map[uint64]map[ethcommon.Address]*bigNumber.Int)

/**********************************************************************************************
** MapPrices will, for a given chainID, return a map of prices.
** The map will be of the form: map[vaultAddress]bigNumber.Int
**********************************************************************************************/
func MapPrices(chainID uint64) map[common.Address]*bigNumber.Int {
	prices := make(map[common.Address]*bigNumber.Int)
	for key, price := range _priceMap[chainID] {
		prices[common.FromAddress(key)] = price
	}
	return prices
}

/**********************************************************************************************
** FindPrice will, for a given chainID, try to find the price of the provided tokenAddress
** stored in _priceMap.
** It will return the price if found, and a boolean indicating if the token was found or not.
**********************************************************************************************/
func FindPrice(chainID uint64, tokenAddress common.Address) (*bigNumber.Int, bool) {
	price, ok := _priceMap[chainID][tokenAddress.ToAddress()]
	if !ok {
		return nil, false
	}
	return price, true
}
