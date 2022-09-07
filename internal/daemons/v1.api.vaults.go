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

// FetchVaultsFromV1 fetches the vaults information from the Yearn V1 API for a given chainID
// and store the result to the global variable VaultsFromAPIV1 for later use.
func FetchVaultsFromV1(chainID uint64) {
	// Get the meta information from the Yearn Meta API
	chainIDStr := strconv.FormatUint(chainID, 10)
	resp, err := http.Get(utils.API_V1_BASE_URL + chainIDStr + `/vaults/all`)
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

	// Unmarshal the response body into the variable VaultsFromAPIV1. Body is a byte array,
	// with this manipulation we are putting it in the correct TAPIV1Vault struct format
	vaults := []models.TAPIV1Vault{}
	if err := json.Unmarshal(body, &vaults); err != nil {
		logs.Error("Error unmarshalling response body from the Yearn Meta API", err)
		return
	}

	// To provide faster access to the data, we index the mapping by the vault address, aka
	// {[vaultAddress]: TAPIV1Vault} if we were working with JS/TS
	if store.VaultsFromAPIV1[chainID] == nil {
		store.VaultsFromAPIV1[chainID] = make(map[common.Address]models.TAPIV1Vault)
	}
	for _, vault := range vaults {
		// common.HexToAddress(vault.Address).String() asserts that the address is a valid
		// chacksummed hex string
		store.VaultsFromAPIV1[chainID][common.HexToAddress(vault.Address)] = vault
	}
	store.SaveInDBForChainID(`VaultsFromAPIV1`, chainID, store.VaultsFromAPIV1[chainID])
}

// LoadAPIV1Vaults will reload the vaults from the v1 API data store from the last state of the local Badger Database
func LoadAPIV1Vaults(chainID uint64, wg *sync.WaitGroup) {
	defer wg.Done()
	temp := make(map[common.Address]models.TAPIV1Vault)
	err := store.LoadFromDBForChainID(`VaultsFromAPIV1`, chainID, &temp)
	if err != nil {
		if err.Error() == "Key not found" {
			logs.Warning("No metaVaults data found for chainID: " + strconv.FormatUint(chainID, 10))
			return
		}
		logs.Error(err)
		return
	}
	if temp != nil && (len(temp) > 0) {
		store.VaultsFromAPIV1[chainID] = temp
		logs.Success("Data loaded for the APIv1 Vaults data store for chainID: " + strconv.FormatUint(chainID, 10))
	} else {
		logs.Warning("No APIv1 Vaults data found for chainID: " + strconv.FormatUint(chainID, 10))
	}
}
