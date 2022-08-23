package daemons

import (
	"encoding/json"
	"strconv"
	"sync"

	"github.com/majorfi/ydaemon/internal/logs"
	"github.com/majorfi/ydaemon/internal/models"
	"github.com/majorfi/ydaemon/internal/store"
	"github.com/majorfi/ydaemon/internal/utils"
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
	store.SaveInDBForChainID(`ProtocolsFromMeta`, chainID, store.ProtocolsFromMeta[chainID])
}

// LoadMetaProtocols will reload the meta protocols data store from the last state of the local Badger Database
func LoadMetaProtocols(chainID uint64, wg *sync.WaitGroup) {
	defer wg.Done()
	temp := make(map[string]models.TProtocolsFromMeta)
	err := store.LoadFromDBForChainID(`ProtocolsFromMeta`, chainID, &temp)
	if err != nil {
		if err.Error() == "Key not found" {
			logs.Warning("No metaProtocols data found for chainID: " + strconv.FormatUint(chainID, 10))
			return
		}
		logs.Error(err)
		return
	}
	if temp != nil && (len(temp) > 0) {
		store.ProtocolsFromMeta[chainID] = temp
		logs.Success("Data loaded for the metaProtocols data store for chainID: " + strconv.FormatUint(chainID, 10))
	} else {
		logs.Warning("No metaProtocols data found for chainID: " + strconv.FormatUint(chainID, 10))
	}
}
