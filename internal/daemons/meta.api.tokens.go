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
}

// LoadMetaTokens is kept in order to have the same behavior everywhere, but as the data
// exists in the same directory as yDaemon, saving the data in the DB is not necessary.
func LoadMetaTokens(chainID uint64, wg *sync.WaitGroup) {
	_ = chainID
	wg.Done()
}
