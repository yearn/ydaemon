package prices

import (
	"encoding/json"
	"io/ioutil"
	"math"
	"net/http"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/common/bigNumber"
	"github.com/yearn/ydaemon/common/env"
	"github.com/yearn/ydaemon/common/logs"
	"github.com/yearn/ydaemon/internal/models"
)

/**************************************************************************************************
** fetchPriceFromGecko tries to fetch the price for a given token from
** the CoinGecko API, returns nil if there is no data returned
**************************************************************************************************/
func fetchPricesFromGecko(chainID uint64, tokens []common.Address) map[common.Address]*bigNumber.Int {
	priceMap := make(map[common.Address]*bigNumber.Int)

	var tokenString []string
	for _, token := range tokens {
		tokenString = append(tokenString, strings.ToLower(token.Hex()))
	}
	req, err := http.NewRequest("GET", env.GECKO_PRICE_URL+env.GECKO_CHAIN_NAMES[chainID], nil)
	if err != nil {
		logs.Warning("Error fetching prices from CoinGecko for chain", chainID)
		return priceMap
	}
	q := req.URL.Query()
	q.Add("contract_addresses", strings.Join(tokenString, ","))
	q.Add("vs_currencies", "usd")
	req.URL.RawQuery = q.Encode()
	resp, err := http.DefaultClient.Do(req)
	if err != nil || resp.StatusCode != 200 {
		logs.Warning("Error fetching prices from CoinGecko for chain", chainID)
		return priceMap
	}

	// Defer the closing of the response body to avoid memory leaks
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logs.Warning("Error unmarshalling response body from the API of CoinGecko for chain", chainID)
		return priceMap
	}
	priceData := models.TGeckoPrice{}
	if err := json.Unmarshal(body, &priceData); err != nil {
		logs.Warning("Error unmarshalling response body from the API of CoinGecko for chain", chainID)
		return priceMap
	}

	// Parse response
	decimalsUSDC := bigNumber.NewFloat(math.Pow10(6))
	for _, tokenStr := range tokenString {
		data, ok := priceData[tokenStr]
		if ok { // Convert price into USDC decimals
			price := bigNumber.NewFloat(data.USDPrice)
			priceMap[common.HexToAddress(tokenStr)] = bigNumber.NewFloat().Mul(price, decimalsUSDC).Int()
		}
	}
	return priceMap
}
