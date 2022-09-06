package daemons

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/internal/logs"
	"github.com/yearn/ydaemon/internal/models"
	"github.com/yearn/ydaemon/internal/store"
	"github.com/yearn/ydaemon/internal/utils"
)

// FetchStrategiesFromRisk fetches the strategies information from the Risk Framework for a given chainID
// and store the result to the global variable StrategiesFromRisk for later use.
func FetchStrategiesFromRisk(chainID uint64) {
	// Get the information from the Risk Framework
	chainIDStr := strconv.FormatUint(chainID, 10)
	resp, err := http.Get(utils.RISK_BASE_URL + chainIDStr)
	if err != nil {
		logs.Error("Error fetching information from the Risk Framework", err)
		return
	}

	// Defer the closing of the response body to avoid memory leaks
	defer resp.Body.Close()

	// Read the response body and store it in the body variable
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logs.Error("Error reading response body from the Risk Framework", err)
		return
	}

	// Unmarshal the response body into the variable StrategiesFromRisk. Body is a byte array,
	// with this manipulation we are putting it in the correct TStrategyFromRisk struct format
	strategies := make(map[common.Address]models.TStrategyFromRisk)
	if err := json.Unmarshal(body, &strategies); err != nil {
		logs.Error("Error unmarshalling response body from the Risk Framework", err)
		return
	}

	// To provide faster access to the data, we index the mapping by the vault address, aka
	// {[vaultAddress]: TStrategyFromRisk} if we were working with JS/TS
	if store.StrategiesFromRisk[chainID] == nil {
		store.StrategiesFromRisk[chainID] = make(map[common.Address]models.TStrategyFromRisk)
	}
	for address, strategy := range strategies {
		store.StrategiesFromRisk[chainID][address] = strategy
	}
	store.SaveInDBForChainID(`StrategiesFromRisk`, chainID, store.StrategiesFromRisk[chainID])
}

// LoadRiskStrategies will reload the risk strategies data store from the last state of the local Badger Database
func LoadRiskStrategies(chainID uint64, wg *sync.WaitGroup) {
	defer wg.Done()
	temp := make(map[common.Address]models.TStrategyFromRisk)
	err := store.LoadFromDBForChainID(`StrategiesFromRisk`, chainID, &temp)
	if err != nil {
		if err.Error() == "Key not found" {
			logs.Warning("No risk data found for chainID: " + strconv.FormatUint(chainID, 10))
			return
		}
		logs.Error(err)
		return
	}
	if temp != nil && (len(temp) > 0) {
		store.StrategiesFromRisk[chainID] = temp
		logs.Success("Data loaded for the risk data store for chainID: " + strconv.FormatUint(chainID, 10))
	} else {
		logs.Warning("No risk data found for chainID: " + strconv.FormatUint(chainID, 10))
	}
}
