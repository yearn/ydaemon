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

// FetchStrategiesFromMeta fetches the strategies information from the Yearn Meta API for a given chainID
// and store the result to the global variable StrategiesFromMeta for later use.
func FetchStrategiesFromMeta(chainID uint64) {
	// Get the meta information from the Yearn Meta API
	chainIDStr := strconv.FormatUint(chainID, 10)
	resp, err := http.Get(utils.META_BASE_URL + chainIDStr + `/strategies/all`)
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

	// Unmarshal the response body into the variable StrategiesFromMeta. Body is a byte array,
	// with this manipulation we are putting it in the correct TStrategyFromMeta struct format
	strategies := []models.TStrategyFromMeta{}
	if err := json.Unmarshal(body, &strategies); err != nil {
		logs.Error("Error unmarshalling response body from the Yearn Meta API", err)
		return
	}

	// To provide faster access to the data, we index the mapping by the vault address, aka
	// {[vaultAddress]: TStrategyFromMeta} if we were working with JS/TS
	if store.StrategiesFromMeta[chainID] == nil {
		store.StrategiesFromMeta[chainID] = make(map[string]models.TStrategyFromMeta)
	}
	for _, strategy := range strategies {
		for _, strategyAddress := range strategy.Addresses {
			// The address is checksummed
			store.StrategiesFromMeta[chainID][common.HexToAddress(strategyAddress).String()] = strategy
		}
	}
	store.SaveInDBForChainID(`StrategiesFromMeta`, chainID, store.StrategiesFromMeta[chainID])
}

// RunMetaStrategies is a goroutine that periodically fetches the strategies information from the
// Yearn Meta API.
// The data is updated every _at least_ 24 hours.
func RunMetaStrategies(chainID uint64, wg *sync.WaitGroup) {
	isDone := false
	for {
		FetchStrategiesFromMeta(chainID)
		if !isDone {
			isDone = true
			wg.Done()
		}
		time.Sleep(24 * time.Hour)
	}
}

// LoadMetaStrategies will reload the meta strategies data store from the last state of the local Badger Database
func LoadMetaStrategies(chainID uint64, wg *sync.WaitGroup) {
	defer wg.Done()
	temp := make(map[string]models.TStrategyFromMeta)
	err := store.LoadFromDBForChainID(`StrategiesFromMeta`, chainID, &temp)
	if err != nil {
		if err.Error() == "Key not found" {
			logs.Warning("No metaVaults data found for chainID: " + strconv.FormatUint(chainID, 10))
			return
		}
		logs.Error(err)
		return
	}
	if temp != nil && (len(temp) > 0) {
		store.StrategiesFromMeta[chainID] = temp
		logs.Success("Data loaded for the metaStrategies data store for chainID: " + strconv.FormatUint(chainID, 10))
	} else {
		logs.Warning("No metaStrategies data found for chainID: " + strconv.FormatUint(chainID, 10))
	}
}
