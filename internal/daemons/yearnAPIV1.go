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
		store.VaultsFromAPIV1[chainID] = make(map[string]models.TAPIV1Vault)
	}
	for _, vault := range vaults {
		// common.HexToAddress(vault.Address).String() asserts that the address is a valid
		// chacksummed hex string
		store.VaultsFromAPIV1[chainID][common.HexToAddress(vault.Address).String()] = vault
	}
}

// RunAPIV1Vaults is a goroutine that periodically fetches the meta information from the
// Yearn Meta API. It runs forever, every 15 minutes, for the desired chain.
func RunAPIV1Vaults(chainID uint64, wg *sync.WaitGroup) {
	isDone := false
	for {
		FetchVaultsFromV1(chainID)
		if !isDone {
			isDone = true
			wg.Done()
		}
		time.Sleep(10 * time.Minute)
	}
}
