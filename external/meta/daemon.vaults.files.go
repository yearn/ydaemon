package meta

import (
	"encoding/json"
	"strconv"
	"strings"
	"sync"

	"github.com/yearn/ydaemon/common/env"
	"github.com/yearn/ydaemon/common/helpers"
	"github.com/yearn/ydaemon/common/logs"
	"github.com/yearn/ydaemon/common/types/common"
)

// FetchVaultsFromMeta fetches the meta information from the Yearn Meta API for a given chainID
// and store the result to the global variable VaultsFromMeta for later use.
func FetchVaultsFromMeta(chainID uint64) {
	vaults := []TVaultFromMeta{}
	chainIDStr := strconv.FormatUint(chainID, 10)
	content, filenames, err := helpers.ReadAllFilesInDir(env.BASE_DATA_PATH+`/meta/vaults/`+chainIDStr+`/`, `.json`)
	if err != nil {
		logs.Warning("Error fetching meta information from the Yearn Meta API for chain", chainID)
		return
	}
	for index, content := range content {
		vault := TVaultFromMeta{}
		if err := json.Unmarshal(content, &vault); err != nil {
			logs.Warning("Error unmarshalling response body from the Yearn Meta API for chain", chainID)
			continue
		}
		vault.Address = common.HexToAddress(strings.TrimSuffix(filenames[index], `.json`))
		vaults = append(vaults, vault)
	}

	// To provide faster access to the data, we index the mapping by the vault address
	if Store.VaultsFromMeta[chainID] == nil {
		Store.VaultsFromMeta[chainID] = make(map[common.Address]TVaultFromMeta)
	}
	for _, vault := range vaults {
		Store.VaultsFromMeta[chainID][vault.Address] = vault
		Store.RawMetaVaults[chainID] = append(Store.RawMetaVaults[chainID], vault)
	}
}

// LoadMetaVaults is kept in order to have the same behavior everywhere, but as the data
// exists in the same directory as yDaemon, saving the data in the DB is not necessary.
func LoadMetaVaults(chainID uint64, wg *sync.WaitGroup) {
	_ = chainID
	wg.Done()
}
