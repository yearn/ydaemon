package tokenList

import (
	"encoding/json"
	"io/ioutil"
	"strconv"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/common/env"
	"github.com/yearn/ydaemon/internal/models"
)

/**********************************************************************************************
** Set of functions to store and retrieve the tokensList from the cache and/or database and
** being able to access them from the rest of the application.
** The _tokenListMap variable is not exported and is only used internally by the functions
** below.
**********************************************************************************************/
var _tokenListMap = make(map[uint64]map[common.Address]models.TYearnTokenListToken)
var _tokenListUpdateMap = make(map[uint64]time.Time)

/**********************************************************************************************
** MapTokenList will, for a given chainID, return the tokenList stored in _tokenListMap.
**********************************************************************************************/
func MapTokenList(chainID uint64) map[common.Address]models.TYearnTokenListToken {
	if _, ok := _tokenListMap[chainID]; !ok {
		_tokenListMap[chainID] = make(map[common.Address]models.TYearnTokenListToken)
	}
	return _tokenListMap[chainID]
}

/**********************************************************************************************
** GetTokenFromList will, for a given chainID, return the desired token stored in _tokenListMap.
**********************************************************************************************/
func GetTokenFromList(chainID uint64, tokenAddress common.Address) (models.TYearnTokenListToken, bool) {
	if value, exist := MapTokenList(chainID)[tokenAddress]; exist {
		return value, true
	}
	return models.TYearnTokenListToken{}, false
}

/**********************************************************************************************
** getLastUpdate will, for a given chainID, return the last time the list was updated and
** stored in _tokenListUpdateMap.
**********************************************************************************************/
func getLastUpdate(chainID uint64) time.Time {
	if _, ok := _tokenListUpdateMap[chainID]; !ok {
		_tokenListUpdateMap[chainID] = time.Time{}
	}
	return _tokenListUpdateMap[chainID]
}

/**********************************************************************************************
** setTokenFromList will, for a given chainID, update or set a token stored in _tokenListMap.
**********************************************************************************************/
func setTokenFromList(chainID uint64, newTokenValue models.TYearnTokenListToken) {
	if _, ok := _tokenListMap[chainID]; !ok {
		_tokenListMap[chainID] = make(map[common.Address]models.TYearnTokenListToken)
	}
	_tokenListMap[chainID][common.HexToAddress(newTokenValue.Address)] = newTokenValue
}

/**********************************************************************************************
** init will, on startup, fetch the token list and update it every 24 hours. This will fetch the
** tokens for our supported zap solvers (Wido and Portals), but also the aggregated list of
** tokens from the top 10 token lists, with a forced threshold of 3 appearances in the top 10
** lists.
**********************************************************************************************/
func loadTokensListFromJSON(chainID uint64) (map[common.Address]models.TYearnTokenListToken, time.Time) {
	chainIDStr := strconv.FormatUint(chainID, 10)
	file, err := ioutil.ReadFile(env.BASE_DATA_PATH + `/tokensList/` + chainIDStr + `.json`)
	var lastUpdate time.Time
	if err != nil {
		return make(map[common.Address]models.TYearnTokenListToken), lastUpdate
	}

	var tokenList models.YTokenList
	if err := json.Unmarshal(file, &tokenList); err != nil {
		return make(map[common.Address]models.TYearnTokenListToken), lastUpdate
	}

	tokenListMap := make(map[common.Address]models.TYearnTokenListToken)
	for _, token := range tokenList.Tokens {
		tokenListMap[common.HexToAddress(token.Address)] = token
	}
	return tokenListMap, tokenList.Timestamp
}

func saveTokensListToJSON(YTokenMap map[common.Address]models.TYearnTokenListToken, path string) {
	tokenList := models.YTokenList{
		Name:      "Yearn Token List",
		LogoURI:   "https://raw.githubusercontent.com/yearn/yearn-assets/master/icons/multichain-tokens/1/0x0bc529c00C6401aEF6D220BE8C6Ea1667F6Ad93e/logo.svg",
		Timestamp: time.Now(),
	}
	tokenList.Version.Major = 1
	tokenList.Version.Minor = 0
	tokenList.Version.Patch = 0
	tokenList.Tokens = make([]models.TYearnTokenListToken, 0)
	for _, token := range YTokenMap {
		tokenList.Tokens = append(tokenList.Tokens, token)
	}

	jsonData, err := json.MarshalIndent(tokenList, "", "  ")
	if err != nil {
		return
	}
	err = ioutil.WriteFile(path, jsonData, 0644)
	if err != nil {
		return
	}
}

func init() {
	for chainID := range env.CHAINS {
		yTokenMap, lastUpdate := loadTokensListFromJSON(chainID)
		_tokenListMap[chainID] = yTokenMap
		_tokenListUpdateMap[chainID] = lastUpdate
	}
	wg := sync.WaitGroup{}
	isDone := false

	wg.Add(1)
	go func() {
		for {
			WidoTokenList = fetchDefaultTokenList(`https://raw.githubusercontent.com/Migratooor/tokenLists/main/lists/wido.json`)
			PortalsTokenList = fetchDefaultTokenList(`https://raw.githubusercontent.com/Migratooor/tokenLists/main/lists/portals.json`)
			if !isDone {
				wg.Done()
				isDone = true
			}
			time.Sleep(24 * time.Hour)
		}
	}()
	wg.Wait()
}
