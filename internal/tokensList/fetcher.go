package tokensList

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/yearn/ydaemon/common/logs"
	"github.com/yearn/ydaemon/common/sort"
	"github.com/yearn/ydaemon/common/types/common"
)

var WidoTokenList DefaultTokenListData
var PortalsTokenList DefaultTokenListData
var AggregatedTokenList DefaultTokenListData

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
		value.Address = common.HexToAddress(value.Address).Hex()
		sortedTokenList.Tokens = append(sortedTokenList.Tokens, value)
	}
	return sortedTokenList
}
