package tokensList

import (
	"encoding/json"
	"io/ioutil"
	"strconv"
	"sync"
	"time"

	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/common/bigNumber"
	"github.com/yearn/ydaemon/common/env"
	"github.com/yearn/ydaemon/common/logs"
)

// DefaultTokenListToken is the token struct used in the default token list
type DefaultTokenListToken struct {
	ChainID  int    `json:"chainId"`
	Decimals int    `json:"decimals"`
	Count    int    `json:"count"`
	Address  string `json:"address"`
	Name     string `json:"name"`
	Symbol   string `json:"symbol"`
	LogoURI  string `json:"logoURI"`
}

// DefaultTokenListData is the token list struct used in the default token list
type DefaultTokenListData struct {
	Name      string `json:"name"`
	Timestamp string `json:"timestamp"`
	Version   struct {
		Major int `json:"major"`
		Minor int `json:"minor"`
		Patch int `json:"patch"`
	} `json:"version"`
	Tags     map[string]interface{}  `json:"tags"`
	LogoURI  string                  `json:"logoURI"`
	Keywords []string                `json:"keywords"`
	Tokens   []DefaultTokenListToken `json:"tokens"`
}

type PortalsTokenListData struct {
	TotalItems int `json:"totalItems"`
	Tokens     []struct {
		Name      string  `json:"name"`
		Decimals  int     `json:"decimals"`
		Symbol    string  `json:"symbol"`
		Price     float64 `json:"price"`
		Address   string  `json:"address"`
		Platform  string  `json:"platform"`
		Network   string  `json:"network"`
		Liquidity float64 `json:"liquidity"`
		Image     string  `json:"image"`
	} `json:"tokens"`
}

/* 🔵 - Yearn Finance **************************************************
** Yearn's own token list, aka what is kept and returned
**********************************************************************/

// SupportedZap is the type of zap supported by the token (Wido, Portals, Cowswap)
type SupportedZap string

const (
	Wido    SupportedZap = "Wido"
	Portals SupportedZap = "Portals"
	Cowswap SupportedZap = "Cowswap"
)

type YTokenFromList struct {
	ChainID       uint64         `json:"chainID"`
	Address       string         `json:"address"`
	Name          string         `json:"name"`
	Symbol        string         `json:"symbol"`
	LogoURI       string         `json:"logoURI"`
	Decimals      int            `json:"decimals"`
	Balance       *bigNumber.Int `json:"balance"`
	SupportedZaps []SupportedZap `json:"supportedZaps"`
}

/**********************************************************************************************
** Set of functions to store and retrieve the tokensList from the cache and/or database and
** being able to access them from the rest of the application.
** The _tokenListMap variable is not exported and is only used internally by the functions
** below.
**********************************************************************************************/
var _tokenListMap = make(map[uint64]map[ethcommon.Address]*YTokenFromList)

/**********************************************************************************************
** MapTokenList will, for a given chainID, return the tokenList stored in _tokenListMap.
**********************************************************************************************/
func MapTokenList(chainID uint64) map[ethcommon.Address]*YTokenFromList {
	if _, ok := _tokenListMap[chainID]; !ok {
		_tokenListMap[chainID] = make(map[ethcommon.Address]*YTokenFromList)
	}
	return _tokenListMap[chainID]
}

/**********************************************************************************************
** GetTokenFromList will, for a given chainID, return the desired token stored in _tokenListMap.
**********************************************************************************************/
func GetTokenFromList(chainID uint64, tokenAddress ethcommon.Address) (*YTokenFromList, bool) {
	if value, exist := MapTokenList(chainID)[tokenAddress]; exist {
		return value, true
	}
	return nil, false
}

/**********************************************************************************************
** setTokenFromList will, for a given chainID, update or set a token stored in _tokenListMap.
**********************************************************************************************/
func setTokenFromList(chainID uint64, newTokenValue *YTokenFromList) {
	if _, ok := _tokenListMap[chainID]; !ok {
		_tokenListMap[chainID] = make(map[ethcommon.Address]*YTokenFromList)
	}
	_tokenListMap[chainID][ethcommon.HexToAddress(newTokenValue.Address)] = newTokenValue
}

/**********************************************************************************************
** init will, on startup, fetch the token list and update it every 24 hours. This will fetch the
** tokens for our supported zap solvers (Wido and Portals), but also the aggregated list of
** tokens from the top 10 token lists, with a forced threshold of 3 appearances in the top 10
** lists.
**********************************************************************************************/
func loadTokensListFromJSON(chainID uint64) map[ethcommon.Address]*YTokenFromList {
	chainIDStr := strconv.FormatUint(chainID, 10)
	file, err := ioutil.ReadFile(env.BASE_DATA_PATH + `/tokensList/` + chainIDStr + `.json`)
	if err != nil {
		return make(map[ethcommon.Address]*YTokenFromList)
	}

	var tokenList map[ethcommon.Address]*YTokenFromList
	if err := json.Unmarshal(file, &tokenList); err != nil {
		return make(map[ethcommon.Address]*YTokenFromList)
	}

	return tokenList
}

func saveTokensListToJSON(chainID uint64, YTokenMap map[ethcommon.Address]*YTokenFromList) {
	jsonData, err := json.MarshalIndent(YTokenMap, "", "  ")
	if err != nil {
		return
	}

	chainIDStr := strconv.FormatUint(chainID, 10)
	err = ioutil.WriteFile(env.BASE_DATA_PATH+`/tokensList/`+chainIDStr+`.json`, jsonData, 0644)
	if err != nil {
		return
	}
}

func init() {
	for _, chainID := range env.SUPPORTED_CHAIN_IDS {
		yTokenMap := loadTokensListFromJSON(chainID)
		_tokenListMap[chainID] = yTokenMap
		logs.Pretty(len(yTokenMap), "tokens loaded for chainID", chainID)
	}
	wg := sync.WaitGroup{}
	isDone := false

	wg.Add(1)
	go func() {
		for {
			AggregatedTokenList = fetchTokenLists()
			WidoTokenList = fetchDefaultTokenList(`https://api.joinwido.com/tokens`)
			PortalsTokenList = fetchDefaultTokenList(`https://api.portals.fi/v1/tokens/format/uniswap`)
			if !isDone {
				wg.Done()
				isDone = true
			}
			time.Sleep(24 * time.Hour)
		}
	}()
	wg.Wait()
}
