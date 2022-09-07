package metaDaemons

import (
	"encoding/json"
	"strconv"
	"sync"

	"github.com/yearn/ydaemon/internal/logs"
	"github.com/yearn/ydaemon/internal/models"
	"github.com/yearn/ydaemon/internal/store"
	"github.com/yearn/ydaemon/internal/utils"
)

// FetchProtocolsFromMeta fetches the protocols information from the data/meta folder for a given chainID
// and store the result to the global variable ProtocolsFromMeta for later use.
func FetchProtocolsFromMeta(chainID uint64) {
	protocols := []models.TProtocolsFromMeta{}
	chainIDStr := strconv.FormatUint(chainID, 10)
	content, _, err := utils.ReadAllFilesInDir(utils.META_BASE_PATH+`/protocols/`+chainIDStr+`/`, `.json`)
	if err != nil {
		logs.Error("Error fetching meta information from the Yearn Meta API", err)
		return
	}
	for _, content := range content {
		protocol := models.TProtocolsFromMeta{}
		if err := json.Unmarshal(content, &protocol); err != nil {
			logs.Error("Error unmarshalling response body from the Yearn Meta API", err)
			continue
		}
		protocols = append(protocols, protocol)
	}
	store.RawMetaProtocols[chainID] = protocols

	if store.ProtocolsFromMeta[chainID] == nil {
		store.ProtocolsFromMeta[chainID] = make(map[string]models.TProtocolsFromMeta)
	}
	for _, protocol := range protocols {
		store.ProtocolsFromMeta[chainID][protocol.Name] = protocol
	}
}

// LoadMetaProtocols is kept in order to have the same behavior everywhere, but as the data
// exists in the same directory as yDaemon, saving the data in the DB is not necessary.
func LoadMetaProtocols(chainID uint64, wg *sync.WaitGroup) {
	_ = chainID
	wg.Done()
}
