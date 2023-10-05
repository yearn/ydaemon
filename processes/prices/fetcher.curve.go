package prices

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/common/bigNumber"
	"github.com/yearn/ydaemon/common/env"
	"github.com/yearn/ydaemon/common/helpers"
	"github.com/yearn/ydaemon/common/traces"
)

func fetchCurve(url string) []TCurveFactoriesPoolData {
	resp, err := http.Get(url)
	if err != nil {
		traces.
			Capture(`error`, `impossible to get curve URL`).
			SetEntity(`prices`).
			SetExtra(`error`, err.Error()).
			SetTag(`url`, url).
			Send()
		return []TCurveFactoriesPoolData{}
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		traces.
			Capture(`error`, `impossible to read curve Get body`).
			SetEntity(`prices`).
			SetExtra(`error`, err.Error()).
			SetTag(`url`, url).
			Send()
		return []TCurveFactoriesPoolData{}
	}
	var factories TCurveFactories
	if err := json.Unmarshal(body, &factories); err != nil {
		traces.
			Capture(`error`, `impossible to unmarshal curve Get body`).
			SetEntity(`prices`).
			SetExtra(`error`, err.Error()).
			SetTag(`url`, url).
			Send()
		return []TCurveFactoriesPoolData{}
	}
	return factories.Data.PoolData
}

// getPricesFromCurveFactoriesAPI is used after setting the prices from the multicall, aka the lens contract.
// Some missing prices may be for Curve LP tokens or Curve tokens.
// The Curve API provides the total supply and the total USD value of the LP tokens. Based on that
// we can calculate the price per token, and assign it to the token if the Store price is 0
func getPricesFromCurveFactoriesAPI(chainID uint64) map[common.Address]*bigNumber.Int {
	newPriceMap := make(map[common.Address]*bigNumber.Int)
	curveFactoryPoolData := []TCurveFactoriesPoolData{}

	// Running a sync group to execute all fetch at the same time
	wg := sync.WaitGroup{}
	for _, url := range env.CHAINS[chainID].Curve.FactoryURIs {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			curveFactoryPoolData = append(curveFactoryPoolData, fetchCurve(url)...)
		}(url)
	}
	wg.Wait()

	// For each pool, we calculate the price per token and assign it to the token
	// if the Store price is 0
	for _, fact := range curveFactoryPoolData {
		//First add the underlying tokens
		for _, coin := range fact.Coins {
			coinAddress := common.HexToAddress(coin.Address)
			coinPrice := bigNumber.NewFloat(coin.USDPrice)
			coinPrice = coinPrice.Mul(coinPrice, bigNumber.NewFloat(1e6))
			coinPriceBigInt := coinPrice.Int()
			newPriceMap[coinAddress] = coinPriceBigInt
		}

		pricePerToken := 0.0
		formatedTotalSupply, _ := helpers.FormatAmount(fact.TotalSupply, 18)
		if formatedTotalSupply > 0 && fact.USDTotal > 0 {
			pricePerToken = fact.USDTotal / formatedTotalSupply
			pricePerTokenBigFloat := bigNumber.NewFloat(pricePerToken)
			pricePerTokenBigFloat = pricePerTokenBigFloat.Mul(pricePerTokenBigFloat, bigNumber.NewFloat(1e6))
			pricePerTokenBigInt := pricePerTokenBigFloat.Int()
			addressToUse := fact.LPAddress
			if addressToUse == `` {
				addressToUse = fact.Address
			}
			newPriceMap[common.HexToAddress(addressToUse)] = pricePerTokenBigInt
		}
	}
	return newPriceMap
}
