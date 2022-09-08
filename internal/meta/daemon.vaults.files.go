package meta

import (
	"encoding/json"
	"strconv"
	"strings"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/internal/utils/helpers"
	"github.com/yearn/ydaemon/internal/utils/logs"
)

// FetchVaultsFromMeta fetches the meta information from the Yearn Meta API for a given chainID
// and store the result to the global variable VaultsFromMeta for later use.
func FetchVaultsFromMeta(chainID uint64) {
	vaults := []TVaultFromMeta{}
	chainIDStr := strconv.FormatUint(chainID, 10)
	content, filenames, err := helpers.ReadAllFilesInDir(helpers.BASE_DATA_PATH+`/meta/vaults/`+chainIDStr+`/`, `.json`)
	if err != nil {
		logs.Warning("Error fetching meta information from the Yearn Meta API")
		return
	}
	for index, content := range content {
		vault := TVaultFromMeta{}
		if err := json.Unmarshal(content, &vault); err != nil {
			logs.Warning("Error unmarshalling response body from the Yearn Meta API")
			continue
		}
		vault.Address = strings.TrimSuffix(filenames[index], `.json`)
		vault.Address = common.HexToAddress(vault.Address).String()
		vaults = append(vaults, vault)
	}

	// To provide faster access to the data, we index the mapping by the vault address, aka
	// {[vaultAddress]: TVaultFromMeta} if we were working with JS/TS
	if Store.VaultsFromMeta[chainID] == nil {
		Store.VaultsFromMeta[chainID] = make(map[common.Address]TVaultFromMeta)
	}
	for _, vault := range vaults {
		vault.MigrationContract = common.HexToAddress(vault.MigrationContract).Hex()
		vault.MigrationTargetVault = common.HexToAddress(vault.MigrationTargetVault).Hex()
		Store.VaultsFromMeta[chainID][common.HexToAddress(vault.Address)] = vault
		Store.RawMetaVaults[chainID] = append(Store.RawMetaVaults[chainID], vault)
	}
}

// LoadMetaVaults is kept in order to have the same behavior everywhere, but as the data
// exists in the same directory as yDaemon, saving the data in the DB is not necessary.
func LoadMetaVaults(chainID uint64, wg *sync.WaitGroup) {
	_ = chainID
	wg.Done()
}
