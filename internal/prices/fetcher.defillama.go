package prices

import (
	"encoding/json"
	"io/ioutil"
	"math"
	"net/http"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/common/bigNumber"
	"github.com/yearn/ydaemon/common/logs"
)

/**************************************************************************************************
** fetchPriceFromLlama tries to fetch the price for a given token from
** the DeFiLlama pricing API, returns nil if there is no data returned
**************************************************************************************************/
func fetchPricesFromLlama(chainID uint64, tokens []common.Address) map[common.Address]*bigNumber.Int {
	priceMap := make(map[common.Address]*bigNumber.Int)

	var tokenString []string
	for _, token := range tokens {
		tokenString = append(tokenString, LLAMA_CHAIN_NAMES[chainID]+`:`+token.String())
	}
	resp, err := http.Get(LLAMA_PRICE_URL + strings.Join(tokenString, ","))
	if err != nil || resp.StatusCode != 200 {
		logs.Warning("Error fetching prices from DeFiLlama for chain", chainID)
		return priceMap
	}
	// Defer the closing of the response body to avoid memory leaks
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logs.Warning("Error unmarshalling response body from the pricing API of DeFiLlama for chain", chainID)
		return priceMap
	}
	priceData := TLlamaPrice{}
	if err := json.Unmarshal(body, &priceData); err != nil {
		logs.Warning("Error unmarshalling response body from the pricing API of DeFiLlama for chain", chainID)
		return priceMap
	}

	// Parse response
	decimalsUSDC := bigNumber.NewFloat(math.Pow10(6))
	for _, tokenStr := range tokenString {
		tokenAddressStr := strings.Split(tokenStr, ":")[1]
		data, ok := priceData.Coins[tokenStr]
		if ok { // Convert price into USDC decimals
			price := bigNumber.NewFloat(data.Price)
			priceMap[common.HexToAddress(tokenAddressStr)] = bigNumber.NewFloat().Mul(price, decimalsUSDC).Int()
		}
	}
	return priceMap
}
