package storage

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/common/env"
	"github.com/yearn/ydaemon/common/logs"
)

/**************************************************************************************************
** The API can be found here: https://docs.pendle.finance/Developers/Helpers/Backend
** We are using Pendle in order to calculate the APRs of some v3 vaults, using Pendle to get Yield.
** From this API, we are consuming the tokens endpoint, which give us:
** - A list of tokens for a given chain
** - The price for each one of theses tokens
**************************************************************************************************/

// This struct represents the relevant response from the Pendle token API.
type TPendleTokenAPIResp struct {
	ChainID  uint64 `json:"chainId"`
	Address  string `json:"address"`
	Name     string `json:"name"`
	Symbol   string `json:"symbol"`
	Decimals uint64 `json:"decimals"`
	Price    struct {
		USD float64 `json:"usd"`
	} `json:"price"`
}

// This struct represents the relevant response from the Pendle market API.
type TPendleMarketAPIResp struct {
	ChainID       uint64  `json:"chainId"`
	Address       string  `json:"address"`
	Name          string  `json:"name"`
	Symbol        string  `json:"symbol"`
	AggretatedAPY float64 `json:"aggregatedApy"`
	MaxBoostedAPY float64 `json:"maxBoostedApy"`
}

/**************************************************************************************************
** As the API response is quite large, we cache it to avoid unnecessary requests.
** This will be refreshed everytime we globaly refresh the APR system
**************************************************************************************************/
var cachedPendleTokens map[uint64]map[string]TPendleTokenAPIResp
var cachedPendleMarkets map[uint64]map[string]TPendleMarketAPIResp

func init() {
	cachedPendleTokens = make(map[uint64]map[string]TPendleTokenAPIResp)
	cachedPendleMarkets = make(map[uint64]map[string]TPendleMarketAPIResp)
}

/**************************************************************************************************
** RefreshPendleMarkets will refresh the cached Pendle markets for a given chainID.
**************************************************************************************************/
func RefreshPendleMarkets(chainID uint64) {
	updatedData, ok := RetrievePendleMarkets(chainID)
	if !ok {
		return
	}
	cachedPendleMarkets[chainID] = updatedData
}

/**************************************************************************************************
** GetCachedPendleTokens retrieves the cached Pendle tokens for a given chainID. It returns a map
** of Pendle tokens and a boolean indicating if the data was found.
**************************************************************************************************/
func GetCachedPendleTokens(chainID uint64) (map[string]TPendleTokenAPIResp, bool) {
	value, ok := cachedPendleTokens[chainID]
	return value, ok
}

/**************************************************************************************************
** GetCachedPendleMarkets retrieves the cached Pendle markets for a given chainID. It returns a map
** of Pendle markets and a boolean indicating if the data was found.
**************************************************************************************************/
func GetCachedPendleMarkets(chainID uint64) (map[string]TPendleMarketAPIResp, bool) {
	value, ok := cachedPendleMarkets[chainID]
	return value, ok
}

/**************************************************************************
** This function retrieves the Pendle token API response.
**************************************************************************/
func RetrievePendleTokens(chainID uint64) (map[string]TPendleTokenAPIResp, bool) {
	if _, ok := cachedPendleTokens[chainID]; !ok {
		cachedPendleTokens[chainID] = map[string]TPendleTokenAPIResp{}
	}
	chain, ok := env.GetChain(chainID)
	if !ok {
		return map[string]TPendleTokenAPIResp{}, false
	}

	if chain.ExtraURI.PendleCoreURI == `` {
		return map[string]TPendleTokenAPIResp{}, false
	}

	tokens := map[string]TPendleTokenAPIResp{}
	chainIDStr := strconv.FormatUint(chainID, 10)
	resp, err := http.Get(chain.ExtraURI.PendleCoreURI + `/assets/all`)
	if err != nil {
		logs.Error(`Error fetching Pendle tokens for chain ` + chainIDStr + `: ` + err.Error())
		return tokens, false
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		logs.Error(`Error reading Pendle tokens response body for chain ` + chainIDStr + `: ` + err.Error())
		return tokens, false
	}
	var pendleTokensApiResp []TPendleTokenAPIResp
	if err := json.Unmarshal(body, &pendleTokensApiResp); err != nil {
		logs.Error(`Error unmarshalling Pendle tokens response body for chain ` + chainIDStr + `: ` + err.Error())
		return tokens, false
	}

	for _, token := range pendleTokensApiResp {
		tokens[common.HexToAddress(token.Address).Hex()] = token
	}
	cachedPendleTokens[chainID] = tokens
	return tokens, true
}

/**************************************************************************
** This function retrieves the Pendle market API response.
**************************************************************************/
func RetrievePendleMarkets(chainID uint64) (map[string]TPendleMarketAPIResp, bool) {
	if _, ok := cachedPendleMarkets[chainID]; !ok {
		cachedPendleMarkets[chainID] = map[string]TPendleMarketAPIResp{}
	}
	chain, ok := env.GetChain(chainID)
	if !ok {
		return map[string]TPendleMarketAPIResp{}, false
	}

	if chain.ExtraURI.PendleCoreURI == `` {
		return map[string]TPendleMarketAPIResp{}, false
	}

	type TPendleMarketsAPIResp struct {
		Total   int                    `json:"total"`
		Results []TPendleMarketAPIResp `json:"results"`
	}

	markets := map[string]TPendleMarketAPIResp{}
	chainIDStr := strconv.FormatUint(chainID, 10)
	baseURI := chain.ExtraURI.PendleCoreURI + `/markets?select=simple`
	skip := 0
	limit := 100
	totalMarketCount := 100 // This is a dummy value to start the loop
	for skip < totalMarketCount {
		resp, err := http.Get(baseURI + `&skip=` + strconv.Itoa(skip) + `&limit=` + strconv.Itoa(limit))
		if err != nil {
			logs.Error(`Error fetching Pendle markets for chain ` + chainIDStr + `: ` + err.Error())
			return markets, false
		}
		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			logs.Error(`Error reading Pendle markets response body for chain ` + chainIDStr + `: ` + err.Error())
			return markets, false
		}
		var pendleMarketsApiResp TPendleMarketsAPIResp
		if err := json.Unmarshal(body, &pendleMarketsApiResp); err != nil {
			logs.Error(`Error unmarshalling Pendle markets response body for chain ` + chainIDStr + `: ` + err.Error())
			return markets, false
		}

		totalMarketCount = pendleMarketsApiResp.Total
		for _, market := range pendleMarketsApiResp.Results {
			markets[common.HexToAddress(market.Address).Hex()] = market
		}
		skip += limit
	}
	cachedPendleMarkets[chainID] = markets
	return markets, true
}
