package metaDaemons

import (
	"encoding/json"
	"strconv"
	"strings"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/internal/meta"
	"github.com/yearn/ydaemon/internal/utils/helpers"
	"github.com/yearn/ydaemon/internal/utils/logs"
)

// FetchTokensFromMeta fetches the tokens information from the Yearn Meta API for a given chainID
// and store the result to the global variable TokensFromMeta for later use.
func FetchTokensFromMeta(chainID uint64) {
	tokens := []meta.TTokenFromMeta{}
	chainIDStr := strconv.FormatUint(chainID, 10)
	content, filenames, err := helpers.ReadAllFilesInDir(helpers.BASE_DATA_PATH+`/meta/tokens/`+chainIDStr+`/`, `.json`)
	if err != nil {
		logs.Warning("Error fetching meta information from the Yearn Meta API")
		return
	}

	for index, content := range content {
		token := meta.TTokenFromMeta{}
		if err := json.Unmarshal(content, &token); err != nil {
			logs.Warning("Error unmarshalling response body from the Yearn Meta API")
			continue
		}
		token.Address = strings.TrimSuffix(filenames[index], `.json`)
		token.Address = common.HexToAddress(token.Address).String()
		tokens = append(tokens, token)
	}
	meta.Store.RawMetaTokens[chainID] = tokens

	// To provide faster access to the data, we index the mapping by the vault address, aka
	// {[vaultAddress]: TTokenFromMeta} if we were working with JS/TS
	if meta.Store.TokensFromMeta[chainID] == nil {
		meta.Store.TokensFromMeta[chainID] = make(map[common.Address]meta.TTokenFromMeta)
	}
	for _, token := range tokens {
		meta.Store.TokensFromMeta[chainID][common.HexToAddress(token.Address)] = token
	}
}

// LoadMetaTokens is kept in order to have the same behavior everywhere, but as the data
// exists in the same directory as yDaemon, saving the data in the DB is not necessary.
func LoadMetaTokens(chainID uint64, wg *sync.WaitGroup) {
	_ = chainID
	wg.Done()
}
