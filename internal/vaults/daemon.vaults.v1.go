package vaults

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"sync"

	"github.com/yearn/ydaemon/internal/types/common"
	"github.com/yearn/ydaemon/internal/utils/helpers"
	"github.com/yearn/ydaemon/internal/utils/logs"
	"github.com/yearn/ydaemon/internal/utils/models"
	"github.com/yearn/ydaemon/internal/utils/store"
)

// FetchVaultsFromV1 fetches the vaults information from the Yearn V1 API for a given chainID
// and store the result to the global variable VaultsFromAPIV1 for later use.
func FetchVaultsFromV1(chainID uint64) {
	// Get the meta information from the Yearn Meta API
	chainIDStr := strconv.FormatUint(chainID, 10)
	resp, err := http.Get(helpers.API_V1_BASE_URL + chainIDStr + `/vaults/all`)
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

	// Unmarshal the response body into the variable VaultsFromAPIV1. Body is a byte array,
	// with this manipulation we are putting it in the correct TAPIV1Vault struct format
	vaults := []models.TAPIV1Vault{}
	if err := json.Unmarshal(body, &vaults); err != nil {
		logs.Warning("Error unmarshalling response body from the Yearn Meta API")
		return
	}

	// To provide faster access to the data, we index the mapping by the vault address
	if Store.VaultsFromAPIV1[chainID] == nil {
		Store.VaultsFromAPIV1[chainID] = make(map[common.Address]models.TAPIV1Vault)
	}
	for _, vault := range vaults {
		Store.VaultsFromAPIV1[chainID][common.HexToAddress(vault.Address)] = vault
	}
	store.SaveInDBForChainID(store.KEYS.VaultsFromAPIV1, chainID, Store.VaultsFromAPIV1[chainID])
}

// LoadAPIV1Vaults will reload the vaults from the v1 API data store from the last state of the local Badger Database
func LoadAPIV1Vaults(chainID uint64, wg *sync.WaitGroup) {
	defer wg.Done()
	temp := make(map[common.Address]models.TAPIV1Vault)
	if err := store.LoadFromDBForChainID(store.KEYS.VaultsFromAPIV1, chainID, &temp); err != nil {
		return
	}
	if temp != nil && (len(temp) > 0) {
		Store.VaultsFromAPIV1[chainID] = temp
		logs.Success("Data loaded for the APIv1 Vaults data store for chainID: " + strconv.FormatUint(chainID, 10))
	}
}
