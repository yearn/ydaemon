package meta

import (
	"encoding/json"
	"strconv"

	"github.com/yearn/ydaemon/common/env"
	"github.com/yearn/ydaemon/common/helpers"
	"github.com/yearn/ydaemon/common/traces"
)

// TProtocolsFromMeta is the structure of data for the protocols metadata stored in data/meta/protocols
type TProtocolsFromMeta struct {
	Name         string         `json:"name"`
	Description  string         `json:"description"`
	ChainID      uint64         `json:"chainID"`
	Localization *TLocalization `json:"localization,omitempty"`
}

/**********************************************************************************************
** Set of functions to store and retrieve the tokens from the cache and/or database and being
** able to access them from the rest of the application.
** The _vaultMap variable is not exported and is only used internally by the functions below.
**********************************************************************************************/
var _metaProtocolMap = make(map[uint64]map[string]*TProtocolsFromMeta)

/**********************************************************************************************
** setProtocolInMap will put a TProtocolsFromMeta in the _metaProtocolMap variable.
**********************************************************************************************/
func setProtocolInMap(chainID uint64, protocol *TProtocolsFromMeta) {
	if _, ok := _metaProtocolMap[chainID]; !ok {
		_metaProtocolMap[chainID] = make(map[string]*TProtocolsFromMeta)
	}
	_metaProtocolMap[chainID][protocol.Name] = protocol
}

/**********************************************************************************************
** GetMetaProtocol will, for a given chainID, try to retrieve the protocol from the
** _metaProtocolMap variable.
** It will return the protocol if found, and a boolean indicating if the protocol was found or
** not.
**********************************************************************************************/
func GetMetaProtocol(chainID uint64, protocolName string) (*TProtocolsFromMeta, bool) {
	if protocolsForChain, ok := _metaProtocolMap[chainID]; ok {
		if protocol, ok := protocolsForChain[protocolName]; ok {
			return protocol, true
		}
	}
	return nil, false
}

/**********************************************************************************************
** ListMetaProtocol will, for a given chainID, list all the protocols from the _metaProtocolMap
** variable.
**********************************************************************************************/
func ListMetaProtocol(chainID uint64) []*TProtocolsFromMeta {
	var protocols []*TProtocolsFromMeta
	for _, protocol := range _metaProtocolMap[chainID] {
		protocols = append(protocols, protocol)
	}
	return protocols
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
		traces.
			Capture(`warn`, `impossible to read meta files for protocols on chain `+chainIDStr).
			SetEntity(`meta`).
			SetExtra(`error`, err.Error()).
			SetTag(`chainID`, strconv.FormatUint(chainID, 10)).
			Send()
		return
	}
	for _, content := range content {
		protocol := TProtocolsFromMeta{}
		if err := json.Unmarshal(content, &protocol); err != nil {
			traces.
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
