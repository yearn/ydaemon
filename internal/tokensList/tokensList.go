package tokensList

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"

	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/common/bigNumber"
	"github.com/yearn/ydaemon/common/helpers"
	"github.com/yearn/ydaemon/common/logs"
	"github.com/yearn/ydaemon/common/sort"
	"github.com/yearn/ydaemon/common/types/common"
)

var WidoTokenList DefaultTokenListData
var PortalsTokenList DefaultTokenListData
var AggregatedTokenList DefaultTokenListData
var YTokenMap map[ethcommon.Address]*YTokenList

// fetchDefaultTokenList is used to fetch a tokenList from theses https://tokenlists.org/
func fetchDefaultTokenList(url string) DefaultTokenListData {
	resp, err := http.Get(url)
	if err != nil {
		logs.Error(err)
		return DefaultTokenListData{}
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logs.Error(err)
		return DefaultTokenListData{}
	}

	if (resp.StatusCode < 200) || (resp.StatusCode > 299) {
		logs.Error(`request failed for url: ` + url)
		return DefaultTokenListData{}
	}

	var list DefaultTokenListData
	if err := json.Unmarshal(body, &list); err != nil {
		logs.Error(err, url)
		return DefaultTokenListData{}
	}
	return list
}

func fetchPortalsTokenList(url string, chainID int) DefaultTokenListData {
	resp, err := http.Get(url)
	if err != nil {
		logs.Error(err)
		return DefaultTokenListData{}
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logs.Error(err)
		return DefaultTokenListData{}
	}

	if (resp.StatusCode < 200) || (resp.StatusCode > 299) {
		logs.Error(`request failed for url: ` + url)
		return DefaultTokenListData{}
	}

	var portalsList PortalsTokenListData
	if err := json.Unmarshal(body, &portalsList); err != nil {
		logs.Error(err, url)
		return DefaultTokenListData{}
	}

	//convert to DefaultTokenListData
	var list DefaultTokenListData
	list.Name = `Portals`
	list.Timestamp = time.Now().Format(time.RFC3339)
	list.Version.Major = 1
	list.Version.Minor = 0
	list.Version.Patch = 0
	for _, token := range portalsList.Tokens {
		list.Tokens = append(list.Tokens, DefaultTokenListToken{
			ChainID:  chainID,
			Decimals: token.Decimals,
			Count:    0,
			Address:  common.HexToAddress(token.Address).Hex(),
			Name:     token.Name,
			Symbol:   token.Symbol,
			LogoURI:  token.Image,
		})
	}

	return list
}

func fetchList(url string, tokenList *DefaultTokenListData, wg *sync.WaitGroup) {
	defer wg.Done()
	list := fetchDefaultTokenList(url)

	for _, token := range list.Tokens {
		found := false
		token.Address = common.HexToAddress(token.Address).Hex()
		for i, t := range tokenList.Tokens {
			if strings.EqualFold(t.Address, token.Address) {
				found = true
				tokenList.Tokens[i].Count++
				break
			}
		}
		if !found {
			token.Count = 1
			tokenList.Tokens = append(tokenList.Tokens, token)
		}
	}
}

func fetchTokenLists() DefaultTokenListData {
	var tokenList DefaultTokenListData

	wg := sync.WaitGroup{}
	wg.Add(10)
	go fetchList(`https://tokens.coingecko.com/uniswap/all.json`, &tokenList, &wg)
	go fetchList(`https://tokens.coingecko.com/uniswap/all.json`, &tokenList, &wg)
	go fetchList(`https://defi.cmc.eth.link`, &tokenList, &wg)
	go fetchList(`https://stablecoin.cmc.eth.link`, &tokenList, &wg)
	go fetchList(`https://erc20.cmc.eth.link`, &tokenList, &wg)
	go fetchList(`https://tokenlist.aave.eth.link`, &tokenList, &wg)
	go fetchList(`https://uniswap.mycryptoapi.com/`, &tokenList, &wg)
	go fetchList(`https://gateway.ipfs.io/ipns/tokens.uniswap.org`, &tokenList, &wg)
	go fetchList(`https://tokenlist.zerion.eth.link`, &tokenList, &wg)
	go fetchList(`https://tokens.1inch.eth.link`, &tokenList, &wg)
	wg.Wait()

	// Merge the two lists
	tokenList.Name = "Yearn Token List"
	tokenList.Keywords = []string{"yearn", "yfi", "yvault", "ytoken", "ycurve", "yprotocol"}
	tokenList.LogoURI = "https://raw.githubusercontent.com/yearn/yearn-assets/3ec995a8b19cd95e56a1a42b18d394d667e0e2cd/icons/multichain-tokens/1/0x0bc529c00C6401aEF6D220BE8C6Ea1667F6Ad93e/logo.svg"
	tokenList.Timestamp = time.Now().Format(time.RFC3339)
	tokenList.Version.Major = 1
	tokenList.Version.Minor = 0
	tokenList.Version.Patch = 0

	var sortedData = make([]interface{}, len(tokenList.Tokens))
	for i, d := range tokenList.Tokens {
		sortedData[i] = d
	}

	sort.SortBy(sortedData, `count`, `desc`)
	var sortedTokenList DefaultTokenListData
	for _, d := range sortedData {
		value := d.(DefaultTokenListToken)
		if value.Count < 3 {
			continue
		}
		value.Address = common.HexToAddress(value.Address).Hex()
		sortedTokenList.Tokens = append(sortedTokenList.Tokens, value)
	}
	return sortedTokenList
}

// init will, on startup, fetch the token list and update it every 24 hours. This will fetch the tokens for our supported zap
// solvers (Wido and Portals), but also the aggregated list of tokens from the top 10 token lists, with a forced threshold of
// 3 appearances in the top 10 lists.
func init() {
	YTokenMap = make(map[ethcommon.Address]*YTokenList)
	wg := sync.WaitGroup{}
	isDone := false

	wg.Add(1)
	go func() {
		for {
			AggregatedTokenList = fetchTokenLists()
			WidoTokenList = fetchDefaultTokenList(`https://api.joinwido.com/tokens`)
			PortalsTokenList = fetchPortalsTokenList(`https://api.portals.fi/v1/tokens/ethereum`, 1)
			if !isDone {
				wg.Done()
				isDone = true
			}
			time.Sleep(24 * time.Hour)
		}
	}()
	wg.Wait()
}

func setSupportedByCowSwap(chainID uint64, tokenMap map[ethcommon.Address]*YTokenList) {
	index := 0
	for address, token := range tokenMap {
		logs.Success(`Testing token: ` + token.Symbol + ` (` + strconv.Itoa(index+1) + `/` + strconv.Itoa(len(tokenMap)) + `)`)
		index++

		time.Sleep(400)
		_, errReason, err := tryQuote(address.Hex())
		if err == nil || (err == nil && errReason == `NoLiquidity`) {
			if value, ok := tokenMap[address]; ok {
				if !helpers.Contains(tokenMap[address].SupportedZaps, Cowswap) {
					tokenMap[address].SupportedZaps = append(value.SupportedZaps, Cowswap)
				}
			}
		}
	}
}

func BuildTokenList(chainID uint64) {
	supportedTokenMap := make(map[string]DefaultTokenListToken)
	for _, token := range AggregatedTokenList.Tokens {
		supportedTokenMap[token.Address] = token
	}

	if len(WidoTokenList.Tokens) > 0 {
		for _, token := range WidoTokenList.Tokens {
			tokenAddress := ethcommon.HexToAddress(token.Address)
			if _, exists := supportedTokenMap[tokenAddress.Hex()]; !exists {
				continue
			}
			if uint64(token.ChainID) != chainID {
				continue
			}
			if value, exist := YTokenMap[tokenAddress]; exist {
				if helpers.Contains(value.SupportedZaps, Wido) {
					continue
				}
				value.SupportedZaps = append(value.SupportedZaps, Wido)
				YTokenMap[tokenAddress] = value
				continue
			}
			YTokenMap[tokenAddress] = &YTokenList{
				ChainID:       uint64(token.ChainID),
				Address:       tokenAddress.Hex(),
				Name:          token.Name,
				Symbol:        token.Symbol,
				Decimals:      token.Decimals,
				LogoURI:       token.LogoURI,
				Balance:       bigNumber.NewInt(0),
				SupportedZaps: []SupportedZap{Wido},
			}
		}
	}

	if len(PortalsTokenList.Tokens) > 0 {
		for _, token := range PortalsTokenList.Tokens {
			tokenAddress := ethcommon.HexToAddress(token.Address)
			if _, exists := supportedTokenMap[tokenAddress.Hex()]; !exists {
				continue
			}
			if uint64(token.ChainID) != chainID {
				continue
			}
			if value, exist := YTokenMap[tokenAddress]; exist {
				if helpers.Contains(value.SupportedZaps, Portals) {
					continue
				}
				value.SupportedZaps = append(value.SupportedZaps, Portals)
				YTokenMap[tokenAddress] = value
				continue
			}
			YTokenMap[tokenAddress] = &YTokenList{
				ChainID:       uint64(token.ChainID),
				Address:       tokenAddress.Hex(),
				Name:          token.Name,
				Symbol:        token.Symbol,
				Decimals:      token.Decimals,
				LogoURI:       token.LogoURI,
				Balance:       bigNumber.NewInt(0),
				SupportedZaps: []SupportedZap{Portals},
			}
		}
	}

	if len(AggregatedTokenList.Tokens) > 0 {
		for _, token := range AggregatedTokenList.Tokens {
			tokenAddress := ethcommon.HexToAddress(token.Address)
			if _, exist := YTokenMap[tokenAddress]; exist {
				continue
			}
			if uint64(token.ChainID) != chainID {
				continue
			}
			YTokenMap[tokenAddress] = &YTokenList{
				ChainID:       uint64(token.ChainID),
				Address:       tokenAddress.Hex(),
				Name:          token.Name,
				Symbol:        token.Symbol,
				Decimals:      token.Decimals,
				LogoURI:       token.LogoURI,
				Balance:       bigNumber.NewInt(0),
				SupportedZaps: []SupportedZap{},
			}
		}
	}

	setSupportedByCowSwap(chainID, YTokenMap)
	saveToJSON(YTokenMap)
}

func saveToJSON(YTokenMap map[ethcommon.Address]*YTokenList) {
	jsonData, err := json.MarshalIndent(YTokenMap, "", "  ")
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile("YTokenMap.json", jsonData, 0644)
	if err != nil {
		panic(err)
	}
}
