package prices

import (
	"encoding/json"
	"io"
	"math"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/common/bigNumber"
	"github.com/yearn/ydaemon/common/env"
	"github.com/yearn/ydaemon/common/logs"
	"github.com/yearn/ydaemon/internal/models"
)

// GECKO_CHAIN_NAMES contains the chain identifiers for the CoinGecko API
var GECKO_CHAIN_NAMES = map[uint64]string{
	1:      `ethereum`,
	10:     `optimistic-ethereum`,
	100:    `xdai`,
	137:    `polygon-pos`,
	146:    `sonic`,
	250:    `fantom`,
	8453:   `base`,
	42161:  `arbitrum-one`,
	747474: `katana`,
}

var keyIndexToUse uint64 = 0

func useCGAPIKey() string {
	if len(env.CG_DEMO_KEYS) == 0 {
		return ``
	}
	if keyIndexToUse >= uint64(len(env.CG_DEMO_KEYS)) {
		keyIndexToUse = 0
	}
	keyToUse := env.CG_DEMO_KEYS[keyIndexToUse]
	return keyToUse
}

/**************************************************************************************************
** fetchPriceFromGecko tries to fetch the price for a given token from
** the CoinGecko API, returns nil if there is no data returned
**************************************************************************************************/
func fetchPricesFromGecko(chainID uint64, tokens []models.TERC20Token) map[common.Address]models.TPrices {
	priceMap := make(map[common.Address]models.TPrices)
	chunkSize := 100
	timeToSleep := rand.Intn(2000-200) + 200
	cgKey := useCGAPIKey()

	for i := 0; i < len(tokens); i += chunkSize {
		time.Sleep(time.Duration(timeToSleep) * time.Millisecond)
		end := i + chunkSize
		if end > len(tokens) {
			end = len(tokens)
		}

		tokensFromChunk := tokens[i:end]
		var tokenString []string
		for _, token := range tokensFromChunk {
			tokenString = append(tokenString, strings.ToLower(token.Address.Hex()))
		}
		req, err := http.NewRequest("GET", env.GECKO_PRICE_URL+GECKO_CHAIN_NAMES[chainID], nil)
		if err != nil {
			logs.Warning("Error fetching prices from CoinGecko for chain", chainID)
			return priceMap
		}
		req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36")
		q := req.URL.Query()
		q.Add("contract_addresses", strings.Join(tokenString, ","))
		q.Add("vs_currencies", "usd")
		if cgKey != `` {
			q.Add("x_cg_demo_api_key", cgKey)
		}
		req.URL.RawQuery = q.Encode()
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			logs.Error(err)
			logs.Warning("Error fetching prices from CoinGecko for chain", chainID)
			return priceMap
		}
		defer resp.Body.Close()

		/******************************************************************************************
		** Decode the body of the response from the API of CoinGecko
		******************************************************************************************/
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			logs.Warning("Error reading response body from the API of CoinGecko for chain", chainID)
			return priceMap
		}

		/******************************************************************************************
		** If we are getting an error from the coingecko call, it's probably linked to the rate
		** limit. We will just switch to the next key and retry the call.
		******************************************************************************************/
		apiKeyStatus := TGeckoAPIKeyStatus{}
		if err := json.Unmarshal(body, &apiKeyStatus); err == nil {
			if apiKeyStatus.Status.ErrorCode == 429 || apiKeyStatus.Status.ErrorCode == 1020 || apiKeyStatus.Status.ErrorCode == 10002 {
				keyIndexToUse++
				i -= chunkSize
				continue
			}
		}

		/******************************************************************************************
		** If we are not hitting the rate limit, we will just parse the response and store the
		** prices in the priceMap
		******************************************************************************************/
		priceData := TGeckoPrice{}
		if err := json.Unmarshal(body, &priceData); err != nil {
			logs.Warning("Error unmarshalling response body from the API of CoinGecko for chain", chainID)
			return priceMap
		}

		/******************************************************************************************
		** For consistency, we will convert the price into USDC decimals, aka 10^6 as "raw" prices.
		** It will not be 100% accurate, as we will miss some decimals, but it's good enough for
		** our use case.
		******************************************************************************************/
		decimalsUSDC := bigNumber.NewFloat(math.Pow10(6))
		for _, tokenAddressStr := range tokenString {
			data, ok := priceData[tokenAddressStr]
			if ok { // Convert price into USDC decimals
				price := bigNumber.NewFloat(data.USDPrice)
				humanizedPrice := price
				priceMap[common.HexToAddress(tokenAddressStr)] = models.TPrices{
					Address:        common.HexToAddress(tokenAddressStr),
					Price:          bigNumber.NewFloat().Mul(price, decimalsUSDC).Int(),
					HumanizedPrice: humanizedPrice,
					Source:         `coinGecko`,
				}
			}
		}
	}
	return priceMap
}
