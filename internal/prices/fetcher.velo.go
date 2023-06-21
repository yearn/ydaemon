package prices

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/common/bigNumber"
	"github.com/yearn/ydaemon/common/traces"
	"github.com/yearn/ydaemon/internal/models"
)

var VELO_PAIR_URI = `https://api.velodrome.finance/api/v1/pairs`

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
