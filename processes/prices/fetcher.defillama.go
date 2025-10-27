package prices

import (
	"encoding/json"
	"fmt"
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
	"github.com/yearn/ydaemon/internal/storage"
)

// LLAMA_CHAIN_NAMES contains the chain identifiers for the DeFiLlama API
var LLAMA_CHAIN_NAMES = map[uint64]string{
	1:      `ethereum`,
	10:     `optimism`,
	100:    `xdai`,
	137:    `polygon`,
	250:    `fantom`,
	8453:   `base`,
	42161:  `arbitrum`,
	747474: `katana`,
}
var AJNA_TOKENS = map[uint64]common.Address{
	1:     common.HexToAddress(`0x9a96ec9B57Fb64FbC60B423d1f4da7691Bd35079`),
	10:    common.HexToAddress(`0x6c518f9D1a163379235816c543E62922a79863Fa`),
	100:   common.HexToAddress(`0x67Ee2155601e168F7777F169Cd74f3E22BB5E0cE`),
	137:   common.HexToAddress(`0xA63b19647787Da652D0826424460D1BBf43Bf9c6'`),
	8453:  common.HexToAddress(`0xf0f326af3b1Ed943ab95C29470730CC8Cf66ae47'`),
	42161: common.HexToAddress(`0xA98c94d67D9dF259Bee2E7b519dF75aB00E3E2A8'`),
}

var KATANA_TOKEN_NAMES_TO_MAINNET_NAMES = map[string]string{
	"Vault Bridge WBTC": "Wrapped BTC",
	"Vault Bridge USDC": "USD Coin",
	"Vault Bridge USDT": "Tether USD",
	"Vault Bridge USDS": "USDS Stablecoin",
	"Vault Bridge ETH":  "Wrapped Ether",
	"AUSD":              "AUSD",
}

// getMainnetAddressForTokenName finds the mainnet address for a token by its name
func getMainnetAddressForTokenName(tokenName string) (common.Address, bool) {
	mainnetTokens := storage.LoadTokensFromJson(uint64(1))
	for _, token := range mainnetTokens.Tokens {
		if token.Name == tokenName {
			logs.Info("hit", tokenName, "on mainnet is", token.Name, "at", token.Address.Hex())
			return token.Address, true
		}
	}
	return common.Address{}, false
}

/**************************************************************************************************
** fetchPricesFromLlama tries to fetch the price for a given token from
** the DeFiLlama pricing API, returns nil if there is no data returned
**************************************************************************************************/
func fetchPricesFromLlama(chainID uint64, tokens []models.TERC20Token) map[common.Address]models.TPrices {
	priceMap := make(map[common.Address]models.TPrices)
	chunkSize := 100
	timeToSleep := rand.Intn(600-100) + 100
	for i := 0; i < len(tokens); i += chunkSize {
		chunkStart := i
		time.Sleep(time.Duration(timeToSleep) * time.Millisecond)

		end := i + chunkSize
		if end > len(tokens) {
			end = len(tokens)
		}

		tokensFromChunk := tokens[i:end]
		var tokenString []string
		katanaToMainnetMapping := make(map[common.Address]common.Address) // maps mainnet address to katana address

		for _, token := range tokensFromChunk {
			lowerHex := strings.ToLower(token.Address.Hex())
			// Handle Katana chain special token mapping
			if chainID == 747474 {
				if mainnetTokenName, exists := KATANA_TOKEN_NAMES_TO_MAINNET_NAMES[token.Name]; exists {
					logs.Info("Katana token", token.Name, "mapped to mainnet token", mainnetTokenName)
					if mainnetAddress, found := getMainnetAddressForTokenName(mainnetTokenName); found {
						katanaToMainnetMapping[mainnetAddress] = token.Address
						tokenString = append(tokenString, LLAMA_CHAIN_NAMES[1]+`:`+strings.ToLower(mainnetAddress.Hex()))
						continue // Skip the normal token addition for this token
					}
				}
			}

			// Normal token handling
			tokenString = append(tokenString, LLAMA_CHAIN_NAMES[chainID]+`:`+lowerHex)

			// Handle AJNA tokens
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

		llamaRequest := env.LLAMA_PRICE_URL + strings.Join(tokenString, ",")
		t0 := time.Now()
		logs.Info("🦙 [LLAMA CHUNK] start", "chain", chainID, "range", fmt.Sprintf("%d-%d", chunkStart, end))
		resp, err := http.Get(llamaRequest)
		if err != nil {
			logs.Warning("Error fetching prices from DeFiLlama for chain", chainID)
			logs.Error(err)
			return priceMap
		}
		if resp.StatusCode != 200 {
			logs.Warning("🦙 [LLAMA] non-200", "chain", chainID, "status", resp.StatusCode)
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
		zeroCount := 0
		for _, tokenStr := range tokenString {
			tokenAddressStr := strings.Split(tokenStr, ":")[1]
			finalTokenAddress := tokenAddressStr

			// Handle Katana token mapping in response
			if chainID == 747474 {
				mainnetAddr := common.HexToAddress(tokenAddressStr)
				if katanaAddr, exists := katanaToMainnetMapping[mainnetAddr]; exists {
					finalTokenAddress = katanaAddr.Hex()
				}
			}

			// Handle AJNA token mapping in response
			if addresses.Equals(tokenAddressStr, AJNA_TOKENS[1]) && chainID == 10 {
				finalTokenAddress = AJNA_TOKENS[10].Hex()
			}
			if addresses.Equals(tokenAddressStr, AJNA_TOKENS[1]) && chainID == 100 {
				finalTokenAddress = AJNA_TOKENS[100].Hex()
			}
			if addresses.Equals(tokenAddressStr, AJNA_TOKENS[1]) && chainID == 137 {
				finalTokenAddress = AJNA_TOKENS[137].Hex()
			}
			if addresses.Equals(tokenAddressStr, AJNA_TOKENS[1]) && chainID == 8453 {
				finalTokenAddress = AJNA_TOKENS[8453].Hex()
			}
			if addresses.Equals(tokenAddressStr, AJNA_TOKENS[1]) && chainID == 42161 {
				finalTokenAddress = AJNA_TOKENS[42161].Hex()
			}

			key := strings.ToLower(tokenStr)
			data, ok := priceData.Coins[key]
			if ok {
				price := bigNumber.NewFloat(data.Price)
				if price.IsZero() {
					if zeroCount < 5 {
						logs.Warning("🦙 [LLAMA ZERO]", "chain", chainID, "token", finalTokenAddress, "symbol", data.Symbol)
					}
					zeroCount++
					continue
				}

				humanizedPrice := price
				priceMap[common.HexToAddress(finalTokenAddress)] = models.TPrices{
					Address:        common.HexToAddress(finalTokenAddress),
					Price:          bigNumber.NewFloat().Mul(price, decimalsUSDC).Int(),
					HumanizedPrice: humanizedPrice,
					Source:         `defillama`,
				}
			}
		}
		if zeroCount > 5 {
			logs.Warning("🦙 [LLAMA ZERO]", "chain", chainID, "extraZero", zeroCount-5)
		}

		took := time.Since(t0)
		if took > 2*time.Second {
			logs.Warning("🦙 [LLAMA CHUNK] slow", "chain", chainID, "took", took)
		}
		missingCount := 0
		for _, token := range tokensFromChunk {
			price, exists := priceMap[token.Address]
			if !exists || price.Price == nil || price.Price.IsZero() {
				if missingCount < 5 {
					logs.Warning("🦙 [LLAMA MISS]", "chain", chainID, "token", token.Address.Hex(), "symbol", token.Symbol, "name", token.Name)
				}
				missingCount++
			}
		}
		if missingCount > 5 {
			logs.Warning("🦙 [LLAMA MISS]", "chain", chainID, "extraMissing", missingCount-5)
		}
		logs.Success("🦙 [LLAMA CHUNK] done", "chain", chainID, "took", took, "missing", missingCount)
	}

	return priceMap
}
