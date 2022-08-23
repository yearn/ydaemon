package daemons

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/majorfi/ydaemon/internal/logs"
	"github.com/majorfi/ydaemon/internal/models"
	"github.com/majorfi/ydaemon/internal/store"
	"github.com/majorfi/ydaemon/internal/utils"
)

// FetchTokensFromMeta fetches the tokens information from the Yearn Meta API for a given chainID
// and store the result to the global variable TokensFromMeta for later use.
func FetchTokensFromMeta(chainID uint64) {
	// Get the meta information from the Yearn Meta API
	chainIDStr := strconv.FormatUint(chainID, 10)
	resp, err := http.Get(utils.META_BASE_URL + chainIDStr + `/tokens/all`)
	if err != nil {
		logs.Error("Error fetching meta information from the Yearn Meta API", err)
		return
	}

	// Defer the closing of the response body to avoid memory leaks
	defer resp.Body.Close()

	// Read the response body and store it in the body variable
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logs.Error("Error reading response body from the Yearn Meta API", err)
		return
	}

	// Unmarshal the response body into the variable TokensFromMeta. Body is a byte array,
	// with this manipulation we are putting it in the correct TTokenFromMeta struct format
	tokens := []models.TTokenFromMeta{}
	if err := json.Unmarshal(body, &tokens); err != nil {
		logs.Error("Error unmarshalling response body from the Yearn Meta API", err)
		return
	}
	store.RawMetaTokens[chainID] = tokens

	// To provide faster access to the data, we index the mapping by the vault address, aka
	// {[vaultAddress]: TTokenFromMeta} if we were working with JS/TS
	if store.TokensFromMeta[chainID] == nil {
		store.TokensFromMeta[chainID] = make(map[common.Address]models.TTokenFromMeta)
	}
	for _, token := range tokens {
		// The address is checksummed
		store.TokensFromMeta[chainID][common.HexToAddress(token.Address)] = token
	}
	store.SaveInDBForChainID(`TokensFromMeta`, chainID, store.TokensFromMeta[chainID])
}

// LoadMetaTokens will reload the meta tokens data store from the last state of the local Badger Database
func LoadMetaTokens(chainID uint64, wg *sync.WaitGroup) {
	defer wg.Done()
	temp := make(map[common.Address]models.TTokenFromMeta)
	err := store.LoadFromDBForChainID(`TokensFromMeta`, chainID, &temp)
	if err != nil {
		if err.Error() == "Key not found" {
			logs.Warning("No metaVaults data found for chainID: " + strconv.FormatUint(chainID, 10))
			return
		}
		logs.Error(err)
		return
	}
	if temp != nil && (len(temp) > 0) {
		store.TokensFromMeta[chainID] = temp
		logs.Success("Data loaded for the metaTokens data store for chainID: " + strconv.FormatUint(chainID, 10))
	} else {
		logs.Warning("No metaTokens data found for chainID: " + strconv.FormatUint(chainID, 10))
	}
}
