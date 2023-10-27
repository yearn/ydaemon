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
	"github.com/yearn/ydaemon/common/helpers"
	"github.com/yearn/ydaemon/common/logs"
	"github.com/yearn/ydaemon/internal/models"
)

// LLAMA_CHAIN_NAMES contains the chain identifiers for the DeFiLlama API
var LLAMA_CHAIN_NAMES = map[uint64]string{
	1:     `ethereum`,
	10:    `optimism`,
	137:   `polygon`,
	250:   `fantom`,
	8453:  `base`,
	42161: `arbitrum`,
}

/**************************************************************************************************
** fetchPriceFromLlama tries to fetch the price for a given token from
** the DeFiLlama pricing API, returns nil if there is no data returned
**************************************************************************************************/
func fetchPricesFromLlama(chainID uint64, tokens []models.TERC20Token) map[common.Address]models.TPrices {
	priceMap := make(map[common.Address]models.TPrices)
	chunkSize := 100
	timeToSleep := rand.Intn(600-100) + 100
	for i := 0; i < len(tokens); i += chunkSize {
		time.Sleep(time.Duration(timeToSleep) * time.Millisecond)

		end := i + chunkSize
		if end > len(tokens) {
			end = len(tokens)
		}

		tokensFromChunk := tokens[i:end]
		var tokenString []string
		for _, token := range tokensFromChunk {
			tokenString = append(tokenString, LLAMA_CHAIN_NAMES[chainID]+`:`+token.Address.Hex())
		}
		resp, err := http.Get(env.LLAMA_PRICE_URL + strings.Join(tokenString, ","))
		if err != nil || resp.StatusCode != 200 {
			logs.Error(err, resp.StatusCode, resp.Status)
			logs.Warning("Error fetching prices from DeFiLlama for chain", chainID)
			return priceMap
		}
		// Defer the closing of the response body to avoid memory leaks
		defer resp.Body.Close()

		// Read the response body
		body, err := io.ReadAll(resp.Body)
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
				humanizedPrice := helpers.ToNormalizedAmount(price.Int(), 6)
				priceMap[common.HexToAddress(tokenAddressStr)] = models.TPrices{
					Address:        common.HexToAddress(tokenAddressStr),
					Price:          bigNumber.NewFloat().Mul(price, decimalsUSDC).Int(),
					HumanizedPrice: humanizedPrice,
					Source:         `defillama`,
				}
			}
		}
	}
	return priceMap
}
