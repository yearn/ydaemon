package storage

import (
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/common/addresses"
)

var _veYFIGaugesSyncMap = make(map[uint64]*sync.Map)

/**************************************************************************************************
** StoreVeYFIGauge will add a new vault in the _veYFIGaugesSyncMap
**************************************************************************************************/
func StoreVeYFIGauge(chainID uint64, vault common.Address, gauge common.Address) {
	safeSyncMap(_veYFIGaugesSyncMap, chainID).Store(vault, gauge)
}

/**************************************************************************************************
** ListVeYFIGauges will retrieve the all the vaults added to the registries from the
** given chainID
**************************************************************************************************/
func ListVeYFIGauges(chainID uint64) (asMap map[common.Address]common.Address) {
	asMap = make(map[common.Address]common.Address) // make to avoid nil map

	/**********************************************************************************************
	** We can just iterate over the syncMap and add the stakingPools to the map and slice.
	**********************************************************************************************/
	safeSyncMap(_veYFIGaugesSyncMap, chainID).Range(func(key, value interface{}) bool {
		vaultAddress := key.(common.Address)
		gaugeAddress := value.(common.Address)
		asMap[vaultAddress] = gaugeAddress
		return true
	})
	return asMap
}

/**************************************************************************************************
** GetGaugeForVault
**************************************************************************************************/
func GetGaugeForVault(chainID uint64, vault common.Address) (common.Address, bool) {
	/**********************************************************************************************
	** We can just iterate over the syncMap until we find the stakingPool for the vault.
	**********************************************************************************************/
	var stakingPool common.Address
	safeSyncMap(_veYFIGaugesSyncMap, chainID).Range(func(key, value interface{}) bool {
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
