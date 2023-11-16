package meta

import (
	"encoding/json"
	"strconv"
	"sync"

	"github.com/yearn/ydaemon/common/env"
	"github.com/yearn/ydaemon/common/helpers"
	"github.com/yearn/ydaemon/common/logs"
	"github.com/yearn/ydaemon/internal/models"
)

/**********************************************************************************************
** Set of functions to store and retrieve the tokens from the cache and/or database and being
** able to access them from the rest of the application.
** The _metaProtocolMap variable is not exported and is only used internally by the functions
** below.
**********************************************************************************************/
var _metaProtocolMap = make(map[uint64]*sync.Map)

func initOrGetMetaProtocolMap(chainID uint64) *sync.Map {
	syncMap := _metaProtocolMap[chainID]
	if syncMap == nil {
		syncMap = &sync.Map{}
		_metaProtocolMap[chainID] = syncMap
	}
	return syncMap
}

func init() {
	for chainID := range env.CHAINS {
		if _, ok := _metaProtocolMap[chainID]; !ok {
			_metaProtocolMap[chainID] = &sync.Map{}
		}
	}
}

/**********************************************************************************************
** setProtocolInMap will put a TProtocolsFromMeta in the _metaProtocolMap variable.
**********************************************************************************************/
func setProtocolInMap(chainID uint64, protocol *models.TProtocolsFromMeta) {
	syncMap := initOrGetMetaProtocolMap(chainID)
	syncMap.Store(protocol.Name, protocol)
}

/**********************************************************************************************
** GetMetaProtocol will, for a given chainID, try to retrieve the protocol from the
** _metaProtocolMap variable.
** It will return the protocol if found, and a boolean indicating if the protocol was found or
** not.
**********************************************************************************************/
func GetMetaProtocol(chainID uint64, protocolName string) (*models.TProtocolsFromMeta, bool) {
	syncMap := initOrGetMetaProtocolMap(chainID)
	data, ok := syncMap.Load(protocolName)
	if !ok {
		return nil, false
	}
	return data.(*models.TProtocolsFromMeta), true
}

/**********************************************************************************************
** ListMetaProtocol will, for a given chainID, list all the protocols from the _metaProtocolMap
** variable.
**********************************************************************************************/
func ListMetaProtocol(chainID uint64) []*models.TProtocolsFromMeta {
	syncMap := initOrGetMetaProtocolMap(chainID)
	var retValue []*models.TProtocolsFromMeta
	syncMap.Range(func(_, data interface{}) bool {
		retValue = append(retValue, data.(*models.TProtocolsFromMeta))
		return true
	})
	return retValue
}

/**************************************************************************************************
** RetrieveAllProtocolsFromFiles will read all files in the /meta/protocols directory for a given
** chainID and store them in the _metaProtocolMap global variable.
** The goal of this function is to get a list of all protocols information about the vaults for a
** given chainID.
**
** Arguments:
** - chainID: the chain ID of the network we are working on
**************************************************************************************************/
func RetrieveAllProtocolsFromFiles(chainID uint64) {
	chainIDStr := strconv.FormatUint(chainID, 10)
	content, _, err := helpers.ReadAllFilesInDir(env.BASE_DATA_PATH+`/meta/protocols/`+chainIDStr+`/`, `.json`)
	if err != nil {
		logs.
			Capture(`warn`, `impossible to read meta files for protocols on chain `+chainIDStr).
			SetEntity(`meta`).
			SetExtra(`error`, err.Error()).
			SetTag(`chainID`, strconv.FormatUint(chainID, 10)).
			Send()
		return
	}
	for _, content := range content {
		protocol := models.TProtocolsFromMeta{}
		if err := json.Unmarshal(content, &protocol); err != nil {
			logs.
				Capture(`warn`, `impossible to unmarshall meta files for protocols response body `+chainIDStr).
				SetEntity(`meta`).
				SetExtra(`error`, err.Error()).
				SetTag(`chainID`, strconv.FormatUint(chainID, 10)).
				SetExtra(`content`, string(content)).
				Send()
			continue
		}
		protocol.ChainID = chainID
		setProtocolInMap(chainID, &protocol)
	}
}
