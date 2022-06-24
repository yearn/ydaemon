package daemons

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"sync"
	"time"

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

	// To provide faster access to the data, we index the mapping by the vault address, aka
	// {[vaultAddress]: TTokenFromMeta} if we were working with JS/TS
	if store.TokensFromMeta[chainID] == nil {
		store.TokensFromMeta[chainID] = make(map[string]models.TTokenFromMeta)
	}
	for _, token := range tokens {
		// common.HexToAddress(token.Address).String() asserts that the address is a valid
		// chacksummed hex string
		store.TokensFromMeta[chainID][common.HexToAddress(token.Address).String()] = token
	}
}

// RunMetaTokens is a goroutine that periodically fetches the tokens information from the
// Yearn Meta API.
// The data is updated every _at least_ 24 hours.
func RunMetaTokens(chainID uint64, wg *sync.WaitGroup) {
	isDone := false
	for {
		FetchTokensFromMeta(chainID)
		if !isDone {
			isDone = true
			wg.Done()
		}
		time.Sleep(24 * time.Hour)
	}
}
