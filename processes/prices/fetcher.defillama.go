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
	"github.com/yearn/ydaemon/common/addresses"
	"github.com/yearn/ydaemon/common/bigNumber"
	"github.com/yearn/ydaemon/common/env"
	"github.com/yearn/ydaemon/common/logs"
	"github.com/yearn/ydaemon/internal/models"
)

// LLAMA_CHAIN_NAMES contains the chain identifiers for the DeFiLlama API
var LLAMA_CHAIN_NAMES = map[uint64]string{
	1:     `ethereum`,
	10:    `optimism`,
	100:   `xdai`,
	137:   `polygon`,
	250:   `fantom`,
	8453:  `base`,
	42161: `arbitrum`,
}
var AJNA_TOKENS = map[uint64]common.Address{
	1:     common.HexToAddress(`0x9a96ec9B57Fb64FbC60B423d1f4da7691Bd35079`),
	10:    common.HexToAddress(`0x6c518f9D1a163379235816c543E62922a79863Fa`),
	100:   common.HexToAddress(`0x67Ee2155601e168F7777F169Cd74f3E22BB5E0cE`),
	137:   common.HexToAddress(`0xA63b19647787Da652D0826424460D1BBf43Bf9c6'`),
	8453:  common.HexToAddress(`0xf0f326af3b1Ed943ab95C29470730CC8Cf66ae47'`),
	42161: common.HexToAddress(`0xA98c94d67D9dF259Bee2E7b519dF75aB00E3E2A8'`),
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
			if addresses.Equals(token.Address, AJNA_TOKENS[10]) && chainID == 10 {
				tokenString = append(tokenString, LLAMA_CHAIN_NAMES[1]+`:`+AJNA_TOKENS[1].Hex())
			}
			if addresses.Equals(token.Address, AJNA_TOKENS[100]) && chainID == 100 {
				tokenString = append(tokenString, LLAMA_CHAIN_NAMES[1]+`:`+AJNA_TOKENS[1].Hex())
			}
			if addresses.Equals(token.Address, AJNA_TOKENS[137]) && chainID == 137 {
				tokenString = append(tokenString, LLAMA_CHAIN_NAMES[1]+`:`+AJNA_TOKENS[1].Hex())
			}
			if addresses.Equals(token.Address, AJNA_TOKENS[8453]) && chainID == 8453 {
				tokenString = append(tokenString, LLAMA_CHAIN_NAMES[1]+`:`+AJNA_TOKENS[1].Hex())
			}
			if addresses.Equals(token.Address, AJNA_TOKENS[42161]) && chainID == 42161 {
				tokenString = append(tokenString, LLAMA_CHAIN_NAMES[1]+`:`+AJNA_TOKENS[1].Hex())
			}
		}
		resp, err := http.Get(env.LLAMA_PRICE_URL + strings.Join(tokenString, ","))
		if err != nil {
			logs.Warning("Error fetching prices from DeFiLlama for chain", chainID)
			logs.Error(err)
			return priceMap
		}
		if resp.StatusCode != 200 {
			logs.Warning("Error fetching prices from DeFiLlama for chain", chainID)
			logs.Error(resp.StatusCode, resp.Status)
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

			if addresses.Equals(tokenAddressStr, AJNA_TOKENS[1]) && chainID == 10 {
				tokenAddressStr = AJNA_TOKENS[1].Hex()
			}
			if addresses.Equals(tokenAddressStr, AJNA_TOKENS[1]) && chainID == 100 {
				tokenAddressStr = AJNA_TOKENS[100].Hex()
			}
			if addresses.Equals(tokenAddressStr, AJNA_TOKENS[1]) && chainID == 137 {
				tokenAddressStr = AJNA_TOKENS[137].Hex()
			}
			if addresses.Equals(tokenAddressStr, AJNA_TOKENS[1]) && chainID == 8453 {
				tokenAddressStr = AJNA_TOKENS[8453].Hex()
			}
			if addresses.Equals(tokenAddressStr, AJNA_TOKENS[1]) && chainID == 42161 {
				tokenAddressStr = AJNA_TOKENS[42161].Hex()
			}

			data, ok := priceData.Coins[tokenStr]
			if ok { // Convert price into USDC decimals
				price := bigNumber.NewFloat(data.Price)
				humanizedPrice := price
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
