package vaults

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"sync"

	"github.com/yearn/ydaemon/common/env"
	"github.com/yearn/ydaemon/common/logs"
	"github.com/yearn/ydaemon/common/store"
	"github.com/yearn/ydaemon/common/types/common"
)

// FetchVaultsFromV1 fetches the vaults information from the Yearn V1 API for a given chainID
// and store the result to the global variable AggregatedVault for later use.
func FetchVaultsFromV1(chainID uint64) {
	// Get the meta information from the Yearn Meta API
	chainIDStr := strconv.FormatUint(chainID, 10)
	resp, err := http.Get(env.API_V1_BASE_URL + chainIDStr + `/vaults/all`)
	if err != nil || resp.StatusCode != 200 {
		logs.Warning("Error fetching vaults from V1 for chain", chainID)
		return
	}

	// Defer the closing of the response body to avoid memory leaks
	defer resp.Body.Close()

	// Read the response body and store it in the body variable
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logs.Warning("Error unmarshalling response body from the vaults from V1 for chain", chainID)
		return
	}

	// Unmarshal the response body into the variable AggregatedVault. Body is a byte array,
	// with this manipulation we are putting it in the correct TLegacyAPI struct format
	vaults := []TLegacyAPI{}
	if err := json.Unmarshal(body, &vaults); err != nil {
		logs.Warning("Error unmarshalling response body from the Yearn Meta API")
		return
	}

	// To provide faster access to the data, we index the mapping by the vault address
	if Store.AggregatedVault[chainID] == nil {
		Store.AggregatedVault[chainID] = make(map[common.Address]*TAggregatedVault)
	}
	for _, vault := range vaults {
		Store.AggregatedVault[chainID][vault.Address] = &TAggregatedVault{
			Address:   vault.Address,
			LegacyAPY: vault.APY,
		}
		go store.SaveInDB(
			chainID,
			store.TABLES.VAULTS,
			vault.Address.String(),
			Store.AggregatedVault[chainID][vault.Address],
		)
	}
}

// LoadAggregatedVaults will reload the vaults from the v1 API data store from the last state of the local Badger Database
func LoadAggregatedVaults(chainID uint64, wg *sync.WaitGroup) {
	defer wg.Done()
	temp := make(map[common.Address]*TAggregatedVault)
	store.Interate(chainID, store.TABLES.VAULTS, &temp)

	if temp != nil && (len(temp) > 0) {
		Store.AggregatedVault[chainID] = temp
		logs.Success("Data loaded for the AggregatedVault store for chainID: " + strconv.FormatUint(chainID, 10))
	}
}
