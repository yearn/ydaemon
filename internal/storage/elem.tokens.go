package storage

import (
	"encoding/json"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/common/env"
	"github.com/yearn/ydaemon/common/helpers"
	"github.com/yearn/ydaemon/common/logs"
	"github.com/yearn/ydaemon/internal/models"
)

var _erc20SyncMap = make(map[uint64]*sync.Map)
var _erc20JSONMetadataSyncMap = sync.Map{}

/** 🔵 - Yearn *************************************************************************************
** The function `loadTokensFromJson` is responsible for loading tokens from a JSON file.
**************************************************************************************************/
func LoadTokensFromJson(chainID uint64) TJsonERC20Storage {
	var tokens TJsonERC20Storage
	chainIDStr := strconv.FormatUint(chainID, 10)

	// Load the JSON file
	fileName := env.BASE_DATA_PATH + "/meta/tokens/" + chainIDStr + ".json"
	file, err := os.Open(env.BASE_DATA_PATH + "/meta/tokens/" + chainIDStr + ".json")
	if err != nil {
		return TJsonERC20Storage{}
	}
	defer file.Close()

	// Decode the JSON file into the map
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&tokens)
	if err != nil {
		logs.Error("Failed to decode tokens JSON file " + fileName + ": " + err.Error())
		return TJsonERC20Storage{}
	}

	return tokens
}

/** 🔵 - Yearn *************************************************************************************
** The function `storeTokensToJson` is responsible for storing tokens to a JSON file.
**************************************************************************************************/
func StoreTokensToJson(chainID uint64, tokens map[common.Address]models.TERC20Token) {
	chainIDStr := strconv.FormatUint(chainID, 10)
	previousTokens := LoadTokensFromJson(chainID)
	version := detectVersionUpdate(chainID, previousTokens.Version, previousTokens.Tokens, tokens)

	data := TJsonERC20Storage{
		TJsonMetadata: TJsonMetadata{
			LastUpdate: time.Now(),
			Version:    version,
		},
		Tokens: tokens,
	}
	_erc20JSONMetadataSyncMap.Store(chainID, TJsonMetadata{
		data.LastUpdate,
		data.Version,
		data.ShouldRefresh,
	})

	file, _ := json.MarshalIndent(data, "", "\t")
	if _, err := os.Stat(env.BASE_DATA_PATH + "/meta/tokens"); os.IsNotExist(err) {
		os.MkdirAll(env.BASE_DATA_PATH+"/meta/tokens", 0755)
	}
	err := os.WriteFile(env.BASE_DATA_PATH+"/meta/tokens/"+chainIDStr+".json", file, 0644)
	if err != nil {
		logs.Error("Failed to write vaults JSON file: " + err.Error())
	}
}

/**************************************************************************************************
** Retrieve the last time the ERC20s were updated for a specific chainID
**************************************************************************************************/
func GetTokensJsonMetadata(chainID uint64) TJsonMetadata {
	if jsonMetadata, ok := _erc20JSONMetadataSyncMap.Load(chainID); ok {
		return jsonMetadata.(TJsonMetadata)
	}
	return TJsonMetadata{}
}

/**************************************************************************************************
** LoadERC20 will retrieve the all the ERC20 tokens added to the configured DB and store them in
** the _erc20SyncMap for fast access during that
**************************************************************************************************/
func LoadERC20(chainID uint64, wg *sync.WaitGroup) {
	if wg != nil {
		defer wg.Done()
	}
	file := LoadTokensFromJson(chainID)
	_erc20JSONMetadataSyncMap.Store(chainID, TJsonMetadata{
		file.LastUpdate,
		file.Version,
		file.ShouldRefresh,
	})
	for _, token := range file.Tokens {
		safeSyncMap(_erc20SyncMap, chainID).Store(token.Address, token)
	}
}

/**************************************************************************************************
** StoreERC20 will add a new erc20 token in the _erc20SyncMap
**************************************************************************************************/
func StoreERC20(chainID uint64, token models.TERC20Token) {
	if helpers.Contains(env.CHAINS[chainID].IgnoredTokens, token.Address) {
		return
	}
	safeSyncMap(_erc20SyncMap, chainID).Store(token.Address, token)
}

/**************************************************************************************************
** ListERC20 will return a list of all the ERC20 stored in the caching system for a given
** chainID. Both a map and a slice are returned.
**************************************************************************************************/
func ListERC20(chainID uint64) (asMap map[common.Address]models.TERC20Token, asSlice []models.TERC20Token) {
	asMap = make(map[common.Address]models.TERC20Token) // make to avoid nil map

	/**********************************************************************************************
	** We can just iterate over the syncMap and add the vaults to the map and slice.
	** As the stored vault data are only a subset of static, we need to use the actual structure

	**********************************************************************************************/
	safeSyncMap(_erc20SyncMap, chainID).Range(func(key, value interface{}) bool {
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
	** We can just iterate over the syncMap and add the vaults to the map and slice.
	** As the stored vault data are only a subset of static, we need to use the actual structure

	**********************************************************************************************/
	safeSyncMap(_erc20SyncMap, chainID).Range(func(key, value interface{}) bool {
		token := value.(models.TERC20Token)
		if token.IsVaultLike() {
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
	tokenFromSyncMap, ok := safeSyncMap(_erc20SyncMap, chainID).Load(tokenAddress)
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
	return GetERC20(chainID, token.GetVaultUnderlying())
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

func ERC20MapToSlice(tokensMap map[common.Address]models.TERC20Token) []models.TERC20Token {
	var tokensSlice []models.TERC20Token
	for _, token := range tokensMap {
		tokensSlice = append(tokensSlice, token)
	}
	return tokensSlice
}
