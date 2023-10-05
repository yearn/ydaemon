package prices

import (
	"encoding/json"
	"io/ioutil"
	"math"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/common/bigNumber"
	"github.com/yearn/ydaemon/common/env"
	"github.com/yearn/ydaemon/common/logs"
)

// GECKO_CHAIN_NAMES contains the chain identifiers for the CoinGecko API
var GECKO_CHAIN_NAMES = map[uint64]string{
	1:     `ethereum`,
	10:    `optimistic-ethereum`,
	137:   `polygon`,
	250:   `fantom`,
	8453:  `base`,
	42161: `arbitrum-one`,
}

/**************************************************************************************************
** fetchPriceFromGecko tries to fetch the price for a given token from
** the CoinGecko API, returns nil if there is no data returned
**************************************************************************************************/
func fetchPricesFromGecko(chainID uint64, tokens []common.Address) map[common.Address]*bigNumber.Int {
	priceMap := make(map[common.Address]*bigNumber.Int)
	chunkSize := 100
	timeToSleep := rand.Intn(2000-200) + 200
	for i := 0; i < len(tokens); i += chunkSize {
		time.Sleep(time.Duration(timeToSleep) * time.Millisecond)
		end := i + chunkSize
		if end > len(tokens) {
			end = len(tokens)
		}

		tokensFromChunk := tokens[i:end]
		var tokenString []string
		for _, token := range tokensFromChunk {
			tokenString = append(tokenString, strings.ToLower(token.Hex()))
		}
		req, err := http.NewRequest("GET", env.GECKO_PRICE_URL+GECKO_CHAIN_NAMES[chainID], nil)
		if err != nil {
			logs.Error(err)
			logs.Warning("Error fetching prices from CoinGecko for chain", chainID)
			return priceMap
		}
		q := req.URL.Query()
		q.Add("contract_addresses", strings.Join(tokenString, ","))
		q.Add("vs_currencies", "usd")
		req.URL.RawQuery = q.Encode()
		resp, err := http.DefaultClient.Do(req)
		if err != nil || resp.StatusCode != 200 {
			logs.Error(err, resp.StatusCode, resp.Status)
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
		priceData := TGeckoPrice{}
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
	}
	return priceMap
}
