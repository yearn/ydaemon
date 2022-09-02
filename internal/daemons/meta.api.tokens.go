package daemons

import (
	"encoding/json"
	"strconv"
	"strings"
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
	tokens := []models.TTokenFromMeta{}
	chainIDStr := strconv.FormatUint(chainID, 10)
	content, filenames, err := utils.ReadAllFilesInDir(utils.META_BASE_PATH+`/tokens/`+chainIDStr+`/`, `.json`)
	if err != nil {
		logs.Error("Error fetching meta information from the Yearn Meta API", err)
		return
	}

	for index, content := range content {
		token := models.TTokenFromMeta{}
		if err := json.Unmarshal(content, &token); err != nil {
			logs.Error("Error unmarshalling response body from the Yearn Meta API", err)
			continue
		}
		token.Address = strings.TrimSuffix(filenames[index], `.json`)
		token.Address = common.HexToAddress(token.Address).String()
		tokens = append(tokens, token)
	}
	store.RawMetaTokens[chainID] = tokens

	// To provide faster access to the data, we index the mapping by the vault address, aka
	// {[vaultAddress]: TTokenFromMeta} if we were working with JS/TS
	if store.TokensFromMeta[chainID] == nil {
		store.TokensFromMeta[chainID] = make(map[common.Address]models.TTokenFromMeta)
	}
	for _, token := range tokens {
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
			logs.Warning("No metaTokens data found for chainID: " + strconv.FormatUint(chainID, 10))
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
