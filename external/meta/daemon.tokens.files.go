package meta

import (
	"encoding/json"
	"strconv"
	"strings"
	"sync"

	"github.com/yearn/ydaemon/common/env"
	"github.com/yearn/ydaemon/common/helpers"
	"github.com/yearn/ydaemon/common/logs"
	"github.com/yearn/ydaemon/common/types/common"
)

// FetchTokensFromMeta fetches the tokens information from the Yearn Meta API for a given chainID
// and store the result to the global variable TokensFromMeta for later use.
func FetchTokensFromMeta(chainID uint64) {
	tokens := []TTokenFromMeta{}
	chainIDStr := strconv.FormatUint(chainID, 10)
	content, filenames, err := helpers.ReadAllFilesInDir(env.BASE_DATA_PATH+`/meta/tokens/`+chainIDStr+`/`, `.json`)
	if err != nil {
		logs.Warning("Error fetching meta information from the Yearn Meta API for chain", chainID)
		return
	}

	for index, content := range content {
		token := TTokenFromMeta{}
		if err := json.Unmarshal(content, &token); err != nil {
			logs.Warning("Error unmarshalling response body from the Yearn Meta API for chain", chainID)
			continue
		}
		token.Address = common.HexToAddress(strings.TrimSuffix(filenames[index], `.json`))
		tokens = append(tokens, token)
	}
	Store.RawMetaTokens[chainID] = tokens

	// To provide faster access to the data, we index the mapping by the vault address, aka
	// {[vaultAddress]: TTokenFromMeta} if we were working with JS/TS
	if Store.TokensFromMeta[chainID] == nil {
		Store.TokensFromMeta[chainID] = make(map[common.Address]TTokenFromMeta)
	}
	for _, token := range tokens {
		Store.TokensFromMeta[chainID][token.Address] = token
	}
}

// LoadMetaTokens is kept in order to have the same behavior everywhere, but as the data
// exists in the same directory as yDaemon, saving the data in the DB is not necessary.
func LoadMetaTokens(chainID uint64, wg *sync.WaitGroup) {
	_ = chainID
	wg.Done()
}
