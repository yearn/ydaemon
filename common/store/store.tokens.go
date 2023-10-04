package store

import (
	"strconv"
	"strings"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/common/addresses"
	"github.com/yearn/ydaemon/common/env"
	"github.com/yearn/ydaemon/common/logs"
	"github.com/yearn/ydaemon/internal/models"
	"gorm.io/gorm"
)

var _erc20SyncMap = make(map[uint64]*sync.Map)

/**************************************************************************************************
** LoadERC20 will retrieve the all the ERC20 tokens added to the configured DB and store them in
** the _erc20SyncMap for fast access during that
**************************************************************************************************/
func LoadERC20(chainID uint64, wg *sync.WaitGroup) {
	if wg != nil {
		defer wg.Done()
	}
	syncMap := _erc20SyncMap[chainID]
	if syncMap == nil {
		syncMap = &sync.Map{}
		_erc20SyncMap[chainID] = syncMap
	}

	switch _dbType {
	case DBBadger:
		logs.Warning(`BadgerDB is deprecated for LoadERC20`)
	case DBSql:
		var temp []DBERC20
		DATABASE.Table(`db_erc20`).
			Where("chain_id = ?", chainID).
			FindInBatches(&temp, 10_000, func(tx *gorm.DB, batch int) error {
				for _, tokenFromDB := range temp {
					token := models.TERC20Token{
						Address:       addresses.ToAddress(tokenFromDB.Address),
						Type:          tokenFromDB.Type,
						Name:          tokenFromDB.Name,
						Symbol:        tokenFromDB.Symbol,
						DisplayName:   tokenFromDB.DisplayName,
						DisplaySymbol: tokenFromDB.DisplaySymbol,
						Description:   tokenFromDB.Description,
						Icon:          env.BASE_ASSET_URL + strconv.FormatUint(chainID, 10) + `/` + tokenFromDB.Address + `/logo-128.png`,
						Decimals:      tokenFromDB.Decimals,
						ChainID:       tokenFromDB.ChainID,
					}
					allUnderlyingAsString := strings.Split(tokenFromDB.UnderlyingTokensAddresses, ",")
					for _, addressStr := range allUnderlyingAsString {
						token.UnderlyingTokensAddresses = append(token.UnderlyingTokensAddresses, common.HexToAddress(addressStr))
					}
					syncMap.Store(tokenFromDB.Address, token)
				}
				return nil
			})
	}
}

/**************************************************************************************************
** AppendToERC20 will add a new erc20 token in the _erc20SyncMap
**************************************************************************************************/
func AppendToERC20(chainID uint64, token models.TERC20Token) {
	syncMap := _erc20SyncMap[chainID]
	key := token.Address.Hex()
	syncMap.Store(key, token)
}

/**************************************************************************************************
** StoreERC20 will store a new erc20 token in the _erc20SyncMap for fast access during that same
** execution, and will store it in the configured DB for future executions.
**************************************************************************************************/
func StoreERC20(chainID uint64, token models.TERC20Token) {
	AppendToERC20(chainID, token)

	switch _dbType {
	case DBBadger:
		// LEGACY: Deprecated
		logs.Warning(`BadgerDB is deprecated for StoreERC20`)
	case DBSql:
		go func() {
			allUnderlyingAsString := []string{}
			for _, address := range token.UnderlyingTokensAddresses {
				allUnderlyingAsString = append(allUnderlyingAsString, address.Hex())
			}
			newItem := &DBERC20{
				UUID:                      getUUID(token.Address.Hex()),
				Address:                   token.Address.Hex(),
				Name:                      token.Name,
				Symbol:                    token.Symbol,
				Type:                      token.Type,
				DisplayName:               token.DisplayName,
				DisplaySymbol:             token.DisplaySymbol,
				Description:               token.Description,
				Decimals:                  token.Decimals,
				ChainID:                   token.ChainID,
				UnderlyingTokensAddresses: strings.Join(allUnderlyingAsString, ","),
			}
			wait(`StoreERC20`)
			DATABASE.
				Table(`db_erc20`).
				Where(`address = ? AND chain_id = ?`, newItem.Address, newItem.ChainID).
				Assign(newItem).
				FirstOrCreate(newItem)
		}()
	}
}

/**************************************************************************************************
** ListERC20 will return a list of all the ERC20 stored in the caching system for a given
** chainID. Both a map and a slice are returned.
**************************************************************************************************/
func ListERC20(chainID uint64) (asMap map[common.Address]models.TERC20Token, asSlice []models.TERC20Token) {
	asMap = make(map[common.Address]models.TERC20Token) // make to avoid nil map

	/**********************************************************************************************
	** We first retrieve the syncMap. This syncMap should be initialized first via the `LoadERC20`
	** function which take the data from the database/badger and store it in it.
	**********************************************************************************************/
	syncMap := _erc20SyncMap[chainID]
	if syncMap == nil {
		syncMap = &sync.Map{}
		_erc20SyncMap[chainID] = syncMap
	}

	/**********************************************************************************************
	** We can just iterate over the syncMap and add the vaults to the map and slice.
	** As the stored vault data are only a subset of static, we need to use the actual structure
	** and not the DBVault one.
	**********************************************************************************************/
	syncMap.Range(func(key, value interface{}) bool {
		token := value.(models.TERC20Token)
		asMap[token.Address] = token
		asSlice = append(asSlice, token)
		return true
	})

	return asMap, asSlice
}

/**************************************************************************************************
** ListVaultsLike will return a list of all the ERC20 stored in the caching system for a given
** chainID. Both a map and a slice are returned. This will filter the list to only return the
** vault-like tokens.
**************************************************************************************************/
func ListVaultsLike(chainID uint64) (asMap map[common.Address]models.TERC20Token, asSlice []models.TERC20Token) {
	asMap = make(map[common.Address]models.TERC20Token) // make to avoid nil map

	/**********************************************************************************************
	** We first retrieve the syncMap. This syncMap should be initialized first via the `LoadERC20`
	** function which take the data from the database/badger and store it in it.
	**********************************************************************************************/
	syncMap := _erc20SyncMap[chainID]
	if syncMap == nil {
		syncMap = &sync.Map{}
		_erc20SyncMap[chainID] = syncMap
	}

	/**********************************************************************************************
	** We can just iterate over the syncMap and add the vaults to the map and slice.
	** As the stored vault data are only a subset of static, we need to use the actual structure
	** and not the DBVault one.
	**********************************************************************************************/
	syncMap.Range(func(key, value interface{}) bool {
		token := value.(models.TERC20Token)
		if IsVaultLike(token) {
			asMap[token.Address] = token
			asSlice = append(asSlice, token)
		}
		return true
	})

	return asMap, asSlice
}

/**************************************************************************************************
** GetERC20 will try, for a specific chain, to find the provided token address as ERC20 struct.
**************************************************************************************************/
func GetERC20(chainID uint64, tokenAddress common.Address) (token models.TERC20Token, ok bool) {
	tokenFromSyncMap, ok := _erc20SyncMap[chainID].Load(tokenAddress.Hex())
	if !ok {
		return models.TERC20Token{}, false
	}
	return tokenFromSyncMap.(models.TERC20Token), true
}

/**********************************************************************************************
** GetUnderlyingERC20 will, for a given chainID, try to find the first underlying token
** of the provided vaultAddress.
** As the Yearn's vaults can only have one underlying token, it will return the first one found
** of the list of underlying tokens.
** It will return the token if found, and a boolean indicating if the token was found or not.
**********************************************************************************************/
func GetUnderlyingERC20(chainID uint64, vaultAddress common.Address) (models.TERC20Token, bool) {
	token, ok := GetERC20(chainID, vaultAddress)
	if !ok {
		return models.TERC20Token{}, false
	}
	return GetERC20(chainID, GetVaultUnderlying(token))
}

/**********************************************************************************************
** ListERC20Addresses will, for a given chainID, return the list of all addresses of the ERC20
** tokens stored in _erc20SyncMap.
**********************************************************************************************/
func ListERC20Addresses(chainID uint64) []common.Address {
	var addresses []common.Address
	_, asSlice := ListERC20(chainID)
	for _, token := range asSlice {
		addresses = append(addresses, token.Address)
	}
	return addresses
}

/**********************************************************************************************
** IsVaultLike will check if the provided token is of the "Yearn Vault" type.
** It takes a TERC20Token as input and returns a boolean value.
** If the token type matches with "Yearn Vault", it returns true. Otherwise, it returns false.
**********************************************************************************************/
func IsVaultLike(t models.TERC20Token) bool {
	return t.Type == models.TokenTypeStandardVault || IsExperimentalVault(t) || IsAutomatedVault(t)
}

/** 🔵 - Yearn *************************************************************************************
** This function checks if the provided token is of the "Experimental Vault" type.
** It takes a TERC20Token as input and returns a boolean value.
**************************************************************************************************/
func IsExperimentalVault(t models.TERC20Token) bool {
	return t.Type == models.TokenTypeExperimentalVault || t.Type == models.TokenTypeLegacyExperimentalVault
}

/** 🔵 - Yearn *************************************************************************************
** This function checks if the provided token is of the "Automated Vault" type.
** It takes a TERC20Token as input and returns a boolean value.
**************************************************************************************************/
func IsAutomatedVault(t models.TERC20Token) bool {
	return t.Type == models.TokenTypeAutomatedVault || t.Type == models.TokenTypeLegacyAutomatedVault
}

/** 🔵 - Yearn *************************************************************************************
** GetVaultUnderlyingAddress will, for a given TERC20Token, return the address of the underlying
** token of the vault. If the vault has no underlying token, it will return an empty address.
** This function is useful when we need to retrieve the underlying token of a vault for further
** operations or computations.
**************************************************************************************************/
func GetVaultUnderlying(t models.TERC20Token) common.Address {
	if len(t.UnderlyingTokensAddresses) > 0 {
		return t.UnderlyingTokensAddresses[0]
	}
	return common.Address{}
}
