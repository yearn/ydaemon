package daemons

import (
	"encoding/json"
	"strconv"
	"strings"
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
	vaults := []models.TVaultFromMeta{}
	chainIDStr := strconv.FormatUint(chainID, 10)
	content, filenames, err := utils.ReadAllFilesInDir(utils.META_BASE_PATH+`/vaults/`+chainIDStr+`/`, `.json`)
	if err != nil {
		logs.Error("Error fetching meta information from the Yearn Meta API", err)
		return
	}
	for index, content := range content {
		vault := models.TVaultFromMeta{}
		if err := json.Unmarshal(content, &vault); err != nil {
			logs.Error("Error unmarshalling response body from the Yearn Meta API", err)
			continue
		}
		vault.Address = strings.TrimSuffix(filenames[index], `.json`)
		vaults = append(vaults, vault)
	}

	// To provide faster access to the data, we index the mapping by the vault address, aka
	// {[vaultAddress]: TVaultFromMeta} if we were working with JS/TS
	if store.VaultsFromMeta[chainID] == nil {
		store.VaultsFromMeta[chainID] = make(map[common.Address]models.TVaultFromMeta)
	}
	for _, vault := range vaults {
		vault.MigrationContract = common.HexToAddress(vault.MigrationContract).Hex()
		vault.MigrationTargetVault = common.HexToAddress(vault.MigrationTargetVault).Hex()
		store.VaultsFromMeta[chainID][common.HexToAddress(vault.Address)] = vault
		store.RawMetaVaults[chainID] = append(store.RawMetaVaults[chainID], vault)
	}
	store.SaveInDBForChainID(`VaultsFromMeta`, chainID, store.VaultsFromMeta[chainID])
}

// LoadMetaVaults will reload the meta vaults data store from the last state of the local Badger Database
func LoadMetaVaults(chainID uint64, wg *sync.WaitGroup) {
	defer wg.Done()
	temp := make(map[common.Address]models.TVaultFromMeta)
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
