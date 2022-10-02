package meta

import (
	"encoding/json"
	"strconv"
	"sync"

	"github.com/yearn/ydaemon/internal/utils/env"
	"github.com/yearn/ydaemon/internal/utils/helpers"
	"github.com/yearn/ydaemon/internal/utils/logs"
)

// FetchProtocolsFromMeta fetches the protocols information from the data/meta folder for a given chainID
// and store the result to the global variable ProtocolsFromMeta for later use.
func FetchProtocolsFromMeta(chainID uint64) {
	protocols := []TProtocolsFromMeta{}
	chainIDStr := strconv.FormatUint(chainID, 10)
	content, _, err := helpers.ReadAllFilesInDir(env.BASE_DATA_PATH+`/meta/protocols/`+chainIDStr+`/`, `.json`)
	if err != nil {
		logs.Warning("Error fetching meta information from the Yearn Meta API for chain", chainID)
		return
	}
	for _, content := range content {
		protocol := TProtocolsFromMeta{}
		if err := json.Unmarshal(content, &protocol); err != nil {
			logs.Warning("Error unmarshalling response body from the Yearn Meta API")
			continue
		}
		protocols = append(protocols, protocol)
	}
	Store.RawMetaProtocols[chainID] = protocols

	if Store.ProtocolsFromMeta[chainID] == nil {
		Store.ProtocolsFromMeta[chainID] = make(map[string]TProtocolsFromMeta)
	}
	for _, protocol := range protocols {
		Store.ProtocolsFromMeta[chainID][protocol.Name] = protocol
	}
}

// LoadMetaProtocols is kept in order to have the same behavior everywhere, but as the data
// exists in the same directory as yDaemon, saving the data in the DB is not necessary.
func LoadMetaProtocols(chainID uint64, wg *sync.WaitGroup) {
	_ = chainID
	wg.Done()
}
