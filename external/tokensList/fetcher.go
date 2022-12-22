package tokensList

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/yearn/ydaemon/common/bigNumber"
	"github.com/yearn/ydaemon/common/helpers"
	"github.com/yearn/ydaemon/common/logs"
	"github.com/yearn/ydaemon/common/types/common"
)

var WidoTokenList DefaultTokenListData
var CoingeckoTokenList DefaultTokenListData
var yTokenMap map[string]YTokenList

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
	var list DefaultTokenListData
	if err := json.Unmarshal(body, &list); err != nil {
		logs.Error(err)
		return DefaultTokenListData{}
	}
	return list
}

// Fetch the tokenLists to keep them up to date
func init() {
	yTokenMap = make(map[string]YTokenList)
	go func() {
		for {
			WidoTokenList = fetchDefaultTokenList(`https://api.joinwido.com/tokens`)
			if len(WidoTokenList.Tokens) > 0 {
				for _, token := range WidoTokenList.Tokens {
					tokenAddress := common.HexToAddress(token.Address).Hex()
					if value, exist := yTokenMap[tokenAddress]; exist {
						if helpers.Contains(value.SupportedZaps, Wido) {
							continue
						}
						value.SupportedZaps = append(value.SupportedZaps, Wido)
						yTokenMap[tokenAddress] = value
						continue
					}
					yTokenMap[tokenAddress] = YTokenList{
						ChainID:       uint64(token.ChainID),
						Address:       tokenAddress,
						Name:          token.Name,
						Symbol:        token.Symbol,
						Decimals:      token.Decimals,
						LogoURI:       token.LogoURI,
						Balance:       bigNumber.NewInt(0),
						SupportedZaps: []SupportedZap{Wido},
					}
				}
			}

			CoingeckoTokenList = fetchDefaultTokenList(`https://tokens.coingecko.com/uniswap/all.json`)
			if len(CoingeckoTokenList.Tokens) > 0 {
				for _, token := range CoingeckoTokenList.Tokens {
					tokenAddress := common.HexToAddress(token.Address).Hex()
					if value, exist := yTokenMap[tokenAddress]; exist {
						if helpers.Contains(value.SupportedZaps, Cowswap) {
							continue
						}
						value.SupportedZaps = append(value.SupportedZaps, Cowswap)
						yTokenMap[tokenAddress] = value
						continue
					}
					yTokenMap[tokenAddress] = YTokenList{
						ChainID:       uint64(token.ChainID),
						Address:       tokenAddress,
						Name:          token.Name,
						Symbol:        token.Symbol,
						Decimals:      token.Decimals,
						LogoURI:       token.LogoURI,
						Balance:       bigNumber.NewInt(0),
						SupportedZaps: []SupportedZap{Cowswap},
					}
				}
			}

			time.Sleep(24 * time.Hour)
		}
	}()
}
