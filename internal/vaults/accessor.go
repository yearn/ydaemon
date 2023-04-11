package vaults

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/internal/models"
)

/**********************************************************************************************
** Set of functions to store and retrieve the tokens from the cache and/or database and being
** able to access them from the rest of the application.
** The _vaultMap variable is not exported and is only used internally by the functions below.
**********************************************************************************************/
var _vaultMap = make(map[uint64]map[common.Address]*models.TVault)

/**********************************************************************************************
** GetVault will, for a given chainID, try to retrieve the vault from the _vaultMap variable.
** It will return the vault if found, and a boolean indicating if the vault was found or not.
**********************************************************************************************/
func GetVault(chainID uint64, vaultAddress common.Address) (*models.TVault, bool) {
	if vaultsForChain, ok := _vaultMap[chainID]; ok {
		if vault, ok := vaultsForChain[vaultAddress]; ok {
			return vault, true
		}
	}
	return nil, false
}

/**********************************************************************************************
** ListVaults will, for a given chainID, return the list of all the vaults stored in _vaultMap.
**********************************************************************************************/
func ListVaults(chainID uint64) []*models.TVault {
	var vaults []*models.TVault
	for _, vault := range _vaultMap[chainID] {
		vaults = append(vaults, vault)
	}
	return vaults
}

/**********************************************************************************************
** MapVaults will, for a given chainID, return the map of all the vaults stored in _vaultMap.
**********************************************************************************************/
func MapVaults(chainID uint64) map[common.Address]*models.TVault {
	var vaults map[common.Address]*models.TVault
	if vaultsForChain, ok := _vaultMap[chainID]; ok {
		vaults = vaultsForChain
	}

	return vaults
}

/**********************************************************************************************
** ListVaultsAddresses will, for a given chainID, return the list of addresses of all the
** vaults stored in _vaultMap.
**********************************************************************************************/
func ListVaultsAddresses(chainID uint64) []common.Address {
	var addresses []common.Address
	for address := range _vaultMap[chainID] {
		addresses = append(addresses, address)
	}
	return addresses
}

/**********************************************************************************************
** FindVault will, for a given chainID, try to find the provided vaultAddress stored in
** _vaultMap. It will return the vault if found, and a boolean indicating if the vault was
** found or not.
**********************************************************************************************/
func FindVault(chainID uint64, vaultAddress common.Address) (*models.TVault, bool) {
	token, ok := _vaultMap[chainID][vaultAddress]
	if !ok {
		return nil, false
	}
	return token, true
}
