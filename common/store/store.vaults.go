package store

import (
	"context"
	"strconv"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/common/bigNumber"
	"github.com/yearn/ydaemon/common/env"
	"github.com/yearn/ydaemon/common/logs"
	"github.com/yearn/ydaemon/internal/models"
	"gorm.io/gorm"
)

var _vaultsSyncMap = make(map[uint64]*sync.Map)

/**************************************************************************************************
** LoadVaults will retrieve the all the vaults from the configured DB and store them in the
** _vaultsSyncMap for fast access during that same execution.
**************************************************************************************************/
func LoadVaults(chainID uint64, wg *sync.WaitGroup) {
	if wg != nil {
		defer wg.Done()
	}
	syncMap := _vaultsSyncMap[chainID]
	if syncMap == nil {
		syncMap = &sync.Map{}
		_vaultsSyncMap[chainID] = syncMap
	}

	switch _dbType {
	case DBBadger:
		// LEGACY: Deprecated
		logs.Warning(`BadgerDB is deprecated for LoadVaults`)
	case DBSql:
		var temp []DBVault
		DATABASE.Table(`db_vaults`).
			Where("chain_id = ?", chainID).
			FindInBatches(&temp, 10_000, func(tx *gorm.DB, batch int) error {
				for _, vaultFromDB := range temp {
					vault := models.TVault{
						Address:        common.HexToAddress(vaultFromDB.Address),
						Management:     common.HexToAddress(vaultFromDB.Management),
						Governance:     common.HexToAddress(vaultFromDB.Governance),
						Guardian:       common.HexToAddress(vaultFromDB.Guardian),
						Rewards:        common.HexToAddress(vaultFromDB.Rewards),
						Type:           vaultFromDB.Type,
						Symbol:         vaultFromDB.Symbol,
						DisplaySymbol:  vaultFromDB.DisplaySymbol,
						FormatedSymbol: vaultFromDB.FormatedSymbol,
						Name:           vaultFromDB.Name,
						DisplayName:    vaultFromDB.DisplayName,
						FormatedName:   vaultFromDB.FormatedName,
						Icon:           env.BASE_ASSET_URL + strconv.FormatUint(chainID, 10) + `/` + vaultFromDB.Address + `/logo-128.png`,
						Version:        vaultFromDB.Version,
						ChainID:        vaultFromDB.ChainID,
						Inception:      vaultFromDB.Inception,
						Activation:     vaultFromDB.Activation,
						Decimals:       vaultFromDB.Decimals,
						PerformanceFee: vaultFromDB.PerformanceFee,
						ManagementFee:  vaultFromDB.ManagementFee,
						Endorsed:       vaultFromDB.Endorsed,
					}

					/******************************************************************************
					** As the vaults stored in the DB do not store the underlying token information
					** we need to retrieve it from the _erc20SyncMap and add it to our structure.
					******************************************************************************/
					vault.Token = models.TERC20Token{
						Address: common.HexToAddress(vaultFromDB.Token),
						ChainID: vaultFromDB.ChainID,
					}
					if token, ok := GetERC20(chainID, common.HexToAddress(vaultFromDB.Token)); ok {
						vault.Token.Decimals = token.Decimals
						vault.Token.Symbol = token.Symbol
						vault.Token.Name = token.Name
						vault.Token.Icon = env.BASE_ASSET_URL + strconv.FormatUint(chainID, 10) + `/` + vault.Token.Address.Hex() + `/logo-128.png`
						vault.Token.DisplaySymbol = token.DisplaySymbol
						vault.Token.DisplayName = token.DisplayName
						vault.Token.Description = token.Description
						vault.Token.Type = token.Type
						vault.Token.UnderlyingTokensAddresses = token.UnderlyingTokensAddresses
					}

					/******************************************************************************
					** As the vaults stored in the DB do not store mutable data we will need to
					** retrieve them on the fly from the blockchain.
					** Here are the missing fields:
					******************************************************************************/
					vault.WithdrawalQueue = []common.Address{}
					vault.PricePerShare = bigNumber.NewInt(0)
					vault.DepositLimit = bigNumber.NewInt(0)
					vault.AvailableDepositLimit = bigNumber.NewInt(0)
					vault.TotalAssets = bigNumber.NewInt(0)
					syncMap.Store(vault.Address, vault)
				}
				return nil
			})
	}
}

/**************************************************************************************************
** AppendToVaultMap will add a new vault in the _vaultsSyncMap
**************************************************************************************************/
func AppendToVaultMap(chainID uint64, vault models.TVault) {
	syncMap := _vaultsSyncMap[chainID]
	if syncMap == nil {
		syncMap = &sync.Map{}
		_vaultsSyncMap[chainID] = syncMap
	}
	syncMap.Store(vault.Address, vault)
}

/**************************************************************************************************
** StoreVault will store a new vault in the _vaultsSyncMap for fast access during that same
** execution, and will store it in the configured DB for future executions.
**************************************************************************************************/
func StoreVault(chainID uint64, vault models.TVault) {
	AppendToVaultMap(chainID, vault)
	key := vault.Address.Hex() + "_" + vault.Token.Address.Hex() + "_" + strconv.FormatUint(vault.Activation, 10) + "_" + strconv.FormatUint(vault.ChainID, 10)

	switch _dbType {
	case DBBadger:
		logs.Warning(`BadgerDB is deprecated for StoreVault`)
	case DBSql:
		go func() {
			newItem := &DBVault{}
			newItem.UUID = getUUID(key)
			newItem.Address = vault.Address.Hex()
			newItem.Management = vault.Management.Hex()
			newItem.Governance = vault.Governance.Hex()
			newItem.Guardian = vault.Guardian.Hex()
			newItem.Rewards = vault.Rewards.Hex()
			newItem.Token = vault.Token.Address.Hex()
			newItem.Type = vault.Type
			newItem.Symbol = vault.Symbol
			newItem.DisplaySymbol = vault.DisplaySymbol
			newItem.FormatedSymbol = vault.FormatedSymbol
			newItem.Name = vault.Name
			newItem.DisplayName = vault.DisplayName
			newItem.FormatedName = vault.FormatedName
			newItem.Version = vault.Version
			newItem.ChainID = vault.ChainID
			newItem.Inception = vault.Inception
			newItem.Activation = vault.Activation
			newItem.Decimals = vault.Decimals
			newItem.PerformanceFee = vault.PerformanceFee
			newItem.ManagementFee = vault.ManagementFee
			newItem.Endorsed = vault.Endorsed
			storeRateLimiter.Wait(context.Background())
			DATABASE.
				Table(`db_vaults`).
				FirstOrCreate(newItem)
		}()
	}
}

/**************************************************************************************************
** ListAllVault will return a list of all the vaults stored in the caching system for a given
** chainID. Both a map and a slice are returned.
**************************************************************************************************/
func ListAllVaults(chainID uint64) (asMap map[common.Address]models.TVault, asSlice []models.TVault) {
	asMap = make(map[common.Address]models.TVault) // make to avoid nil map

	/**********************************************************************************************
	** We first retrieve the syncMap. This syncMap should be initialized first via the `LoadVaults`
	** function which take the data from the database/badger and store it in it.
	**********************************************************************************************/
	syncMap := _vaultsSyncMap[chainID]
	if syncMap == nil {
		syncMap = &sync.Map{}
		_vaultsSyncMap[chainID] = syncMap
	}

	/**********************************************************************************************
	** We can just iterate over the syncMap and add the vaults to the map and slice.
	** As the stored vault data are only a subset of static, we need to use the actual structure
	** and not the DBVault one.
	**********************************************************************************************/
	syncMap.Range(func(key, value interface{}) bool {
		vault := value.(models.TVault)
		asMap[vault.Address] = vault
		asSlice = append(asSlice, vault)
		return true
	})

	return asMap, asSlice
}
