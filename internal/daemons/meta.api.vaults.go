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

// FetchVaultsFromMeta fetches the meta information from the Yearn Meta API for a given chainID
// and store the result to the global variable VaultsFromMeta for later use.
func FetchVaultsFromMeta(chainID uint64) {
	// Get the meta information from the Yearn Meta API
	chainIDStr := strconv.FormatUint(chainID, 10)
	resp, err := http.Get(utils.META_BASE_URL + chainIDStr + `/vaults/all`)
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

	// Unmarshal the response body into the variable VaultsFromMeta. Body is a byte array,
	// with this manipulation we are putting it in the correct TVaultFromMeta struct format
	vaults := []models.TVaultFromMeta{}
	if err := json.Unmarshal(body, &vaults); err != nil {
		logs.Error("Error unmarshalling response body from the Yearn Meta API", err)
		return
	}

	// To provide faster access to the data, we index the mapping by the vault address, aka
	// {[vaultAddress]: TVaultFromMeta} if we were working with JS/TS
	if store.VaultsFromMeta[chainID] == nil {
		store.VaultsFromMeta[chainID] = make(map[string]models.TVaultFromMeta)
	}
	for _, vault := range vaults {
		// The address is checksummed
		store.VaultsFromMeta[chainID][common.HexToAddress(vault.Address).String()] = vault
	}
	store.SaveInDBForChainID(`VaultsFromMeta`, chainID, store.VaultsFromMeta[chainID])
}

// LoadMetaVaults will reload the meta vaults data store from the last state of the local Badger Database
func LoadMetaVaults(chainID uint64, wg *sync.WaitGroup) {
	defer wg.Done()
	temp := make(map[string]models.TVaultFromMeta)
	err := store.LoadFromDBForChainID(`VaultsFromMeta`, chainID, &temp)
	if err != nil {
		if err.Error() == "Key not found" {
			logs.Warning("No metaVaults data found for chainID: " + strconv.FormatUint(chainID, 10))
			return
		}
		logs.Error(err)
		return
	}
	if temp != nil && (len(temp) > 0) {
		store.VaultsFromMeta[chainID] = temp
		logs.Success("Data loaded for the metaVaults data store for chainID: " + strconv.FormatUint(chainID, 10))
	} else {
		logs.Warning("No metaVaults data found for chainID: " + strconv.FormatUint(chainID, 10))
	}
}
