package storage

import (
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/common/addresses"
)

var _juicedStakingSyncMap = make(map[uint64]*sync.Map)

/**************************************************************************************************
** StoreJuicedStaking will add a new entry in the _juicedStakingSyncMap syncMap
**************************************************************************************************/
func StoreJuicedStaking(chainID uint64, vault common.Address, stakingContract common.Address) {
	safeSyncMap(_juicedStakingSyncMap, chainID).Store(vault, stakingContract)
}

/**************************************************************************************************
** ListJuicedStaking will retrieve the all the vaults added to the registries from the
** given chainID
**************************************************************************************************/
func ListJuicedStaking(chainID uint64) (asMap map[common.Address]common.Address) {
	asMap = make(map[common.Address]common.Address) // make to avoid nil map

	/**********************************************************************************************
	** We can just iterate over the syncMap and add the stakingContracts to the map and slice.
	**********************************************************************************************/
	safeSyncMap(_juicedStakingSyncMap, chainID).Range(func(key, value interface{}) bool {
		vaultAddress := key.(common.Address)
		stakingContractAddress := value.(common.Address)
		asMap[vaultAddress] = stakingContractAddress
		return true
	})
	return asMap
}

/**************************************************************************************************
** GetJuicedStakingForVault
**************************************************************************************************/
func GetJuicedStakingForVault(chainID uint64, vault common.Address) (common.Address, bool) {
	/**********************************************************************************************
	** We can just iterate over the syncMap until we find the stakingContract for the vault.
	**********************************************************************************************/
	var stakingContract common.Address
	safeSyncMap(_juicedStakingSyncMap, chainID).Range(func(key, value interface{}) bool {
		vaultAddress := key.(common.Address)
		stakingContractAddress := value.(common.Address)
		if addresses.Equals(vaultAddress, vault) {
			stakingContract = stakingContractAddress
			return false
		}
		return true
	})
	return stakingContract, stakingContract != (common.Address{})
}
