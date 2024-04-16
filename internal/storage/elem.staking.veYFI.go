package storage

import (
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/common/addresses"
)

var _veYFIStakingSyncMap = make(map[uint64]*sync.Map)

/**************************************************************************************************
** StoreVeYFIStaking will add a new entry in the _veYFIStakingSyncMap syncMap
**************************************************************************************************/
func StoreVeYFIStaking(chainID uint64, vault common.Address, gauge common.Address) {
	safeSyncMap(_veYFIStakingSyncMap, chainID).Store(vault, gauge)
}

/**************************************************************************************************
** ListVeYFIStaking will retrieve the all the vaults added to the registries from the
** given chainID
**************************************************************************************************/
func ListVeYFIStaking(chainID uint64) (asMap map[common.Address]common.Address) {
	asMap = make(map[common.Address]common.Address) // make to avoid nil map

	/**********************************************************************************************
	** We can just iterate over the syncMap and add the stakingPools to the map and slice.
	**********************************************************************************************/
	safeSyncMap(_veYFIStakingSyncMap, chainID).Range(func(key, value interface{}) bool {
		vaultAddress := key.(common.Address)
		gaugeAddress := value.(common.Address)
		asMap[vaultAddress] = gaugeAddress
		return true
	})
	return asMap
}

/**************************************************************************************************
** GetVeYFIStakingForVault
**************************************************************************************************/
func GetVeYFIStakingForVault(chainID uint64, vault common.Address) (common.Address, bool) {
	/**********************************************************************************************
	** We can just iterate over the syncMap until we find the stakingPool for the vault.
	**********************************************************************************************/
	var stakingPool common.Address
	safeSyncMap(_veYFIStakingSyncMap, chainID).Range(func(key, value interface{}) bool {
		vaultAddress := key.(common.Address)
		gaugeAddress := value.(common.Address)
		if addresses.Equals(vaultAddress, vault) {
			stakingPool = gaugeAddress
			return false
		}
		return true
	})
	return stakingPool, stakingPool != (common.Address{})
}
