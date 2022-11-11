package meta

import (
	"encoding/json"
	"strconv"
	"strings"

	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/common/env"
	"github.com/yearn/ydaemon/common/helpers"
	"github.com/yearn/ydaemon/common/logs"
	"github.com/yearn/ydaemon/common/types/common"
)

// TVaultFromMeta is the structure of data we receive when calling meta.yearn.finance/api/1/vaults/all
type TVaultFromMeta struct {
	Address              common.Address `json:"address"`
	MigrationTargetVault common.Address `json:"migrationTargetVault"`
	MigrationContract    common.Address `json:"migrationContract"`
	DisplayName          string         `json:"displayName"`
	Comment              string         `json:"comment"`
	APYTypeOverride      string         `json:"apyTypeOverride"`
	APYOverride          float64        `json:"apyOverride"`
	Order                float32        `json:"order"`
	ChainID              uint64         `json:"chainID"`
	HideAlways           bool           `json:"hideAlways"`
	DepositsDisabled     bool           `json:"depositsDisabled"`
	WithdrawalsDisabled  bool           `json:"withdrawalsDisabled"`
	MigrationAvailable   bool           `json:"migrationAvailable"`
	AllowZapIn           bool           `json:"allowZapIn"`
	AllowZapOut          bool           `json:"allowZapOut"`
	Retired              bool           `json:"retired"`
}

/**********************************************************************************************
** Set of functions to store and retrieve the tokens from the cache and/or database and being
** able to access them from the rest of the application.
** The _vaultMap variable is not exported and is only used internally by the functions below.
**********************************************************************************************/
var _metaVaultMap = make(map[uint64]map[ethcommon.Address]*TVaultFromMeta)

/**********************************************************************************************
** setVaultInMap will put a TVaultFromMeta in the _vaultMap variable.
**********************************************************************************************/
func setVaultInMap(chainID uint64, vault *TVaultFromMeta) {
	if _, ok := _metaVaultMap[chainID]; !ok {
		_metaVaultMap[chainID] = make(map[ethcommon.Address]*TVaultFromMeta)
	}
	_metaVaultMap[chainID][vault.Address.Address] = vault
}

/**********************************************************************************************
** GetMetaVault will, for a given chainID, try to retrieve the vault from the _metaVaultMap
** variable.
** It will return the vault if found, and a boolean indicating if the vault was found or not.
**********************************************************************************************/
func GetMetaVault(chainID uint64, vaultAddress common.Address) (*TVaultFromMeta, bool) {
	if vaultsForChain, ok := _metaVaultMap[chainID]; ok {
		if vault, ok := vaultsForChain[vaultAddress.ToAddress()]; ok {
			return vault, true
		}
	}
	return nil, false
}

/**********************************************************************************************
** ListMetaVaults will, for a given chainID, list all the vaults from the _metaVaultMap
** variable.
**********************************************************************************************/
func ListMetaVaults(chainID uint64) []*TVaultFromMeta {
	var vaults []*TVaultFromMeta
	for _, vault := range _metaVaultMap[chainID] {
		vaults = append(vaults, vault)
	}
	return vaults
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
		vault.ChainID = chainID
		setVaultInMap(chainID, &vault)
	}
}
