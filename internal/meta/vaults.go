package meta

import (
	"encoding/json"
	"strconv"
	"strings"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/common/env"
	"github.com/yearn/ydaemon/common/helpers"
	"github.com/yearn/ydaemon/common/traces"
	"github.com/yearn/ydaemon/internal/models"
)

/**********************************************************************************************
** Set of functions to store and retrieve the tokens from the cache and/or database and being
** able to access them from the rest of the application.
** The _metaVaultMap variable is not exported and is only used internally by the functions
** below.
**********************************************************************************************/
var _metaVaultMap = make(map[uint64]*sync.Map)

func initOrGetMetaVaultsMap(chainID uint64) *sync.Map {
	syncMap := _metaVaultMap[chainID]
	if syncMap == nil {
		syncMap = &sync.Map{}
		_metaVaultMap[chainID] = syncMap
	}
	return syncMap
}

func init() {
	for _, chainID := range env.SUPPORTED_CHAIN_IDS {
		if _, ok := _metaVaultMap[chainID]; !ok {
			_metaVaultMap[chainID] = &sync.Map{}
		}
	}
}

/**********************************************************************************************
** setVaultInMap will put a TInternalVaultFromMeta in the _metaVaultMap variable.
**********************************************************************************************/
func setVaultInMap(chainID uint64, vault *models.TInternalVaultFromMeta) {
	syncMap := initOrGetMetaVaultsMap(chainID)
	syncMap.Store(vault.Address, vault)
}

/**********************************************************************************************
** GetMetaVault will, for a given chainID, try to retrieve the vault from the _metaVaultMap
** variable.
** It will return the vault if found, and a boolean indicating if the vault was found or not.
**********************************************************************************************/
func GetMetaVault(chainID uint64, vaultAddress common.Address) (*models.TInternalVaultFromMeta, bool) {
	syncMap := initOrGetMetaVaultsMap(chainID)
	data, ok := syncMap.Load(vaultAddress)
	if !ok {
		return nil, false
	}
	return data.(*models.TInternalVaultFromMeta), true
}

/**********************************************************************************************
** ListMetaVaults will, for a given chainID, list all the vaults from the _metaVaultMap
** variable.
**********************************************************************************************/
func ListMetaVaults(chainID uint64) []*models.TInternalVaultFromMeta {
	syncMap := initOrGetMetaVaultsMap(chainID)
	var retValue []*models.TInternalVaultFromMeta
	syncMap.Range(func(_, data interface{}) bool {
		retValue = append(retValue, data.(*models.TInternalVaultFromMeta))
		return true
	})
	return retValue
}

/**************************************************************************************************
** RetrieveAllVaultsFromFiles will read all files in the /meta/vaults directory for a given chainID
** and store them in the _metaVaultMap global variable.
** The goal of this function is to get a list of all meta information about the vaults for a given
** chainID.
**
** Arguments:
** - chainID: the chain ID of the network we are working on
**************************************************************************************************/
func RetrieveAllVaultsFromFiles(chainID uint64) {
	chainIDStr := strconv.FormatUint(chainID, 10)
	content, filenames, err := helpers.ReadAllFilesInDir(env.BASE_DATA_PATH+`/meta/vaults/`+chainIDStr+`/`, `.json`)
	if err != nil {
		traces.
			Capture(`warn`, `impossible to read meta files for vaults on chain `+chainIDStr).
			SetEntity(`meta`).
			SetExtra(`error`, err.Error()).
			SetTag(`chainID`, strconv.FormatUint(chainID, 10)).
			Send()
		return
	}
	for index, content := range content {
		vault := models.TInternalVaultFromMeta{}
		if err := json.Unmarshal(content, &vault); err != nil {
			traces.
				Capture(`warn`, `impossible to unmarshall meta files for vaults response body `+chainIDStr).
				SetEntity(`meta`).
				SetExtra(`error`, err.Error()).
				SetTag(`chainID`, strconv.FormatUint(chainID, 10)).
				SetExtra(`content`, string(content)).
				Send()
			continue
		}
		vault.Address = common.HexToAddress(strings.TrimSuffix(filenames[index], `.json`))
		vault.ChainID = chainID
		setVaultInMap(chainID, &vault)
	}
}
