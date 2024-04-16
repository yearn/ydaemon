package storage

import (
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/common/addresses"
	"github.com/yearn/ydaemon/internal/models"
)

var _opStakingSyncMap = make(map[uint64]*sync.Map)

type TStakingKey string

const (
	PerVault TStakingKey = "PerVault"
	PerPool  TStakingKey = "PerPool"
)

/**************************************************************************************************
** StoreOPStaking will add a new vault in the _vaultsSyncMap
**************************************************************************************************/
func StoreOPStaking(chainID uint64, pool models.TStakingPoolAdded) {
	safeSyncMap(_opStakingSyncMap, chainID).Store(pool.StackingPoolAddress, pool)
}

/**************************************************************************************************
** ListOPStaking will retrieve the all the vaults added to the registries from the
** given chainID
**************************************************************************************************/
func ListOPStaking(chainID uint64, key TStakingKey) (asMap map[common.Address]models.TStakingPoolAdded, asSlice []models.TStakingPoolAdded) {
	asMap = make(map[common.Address]models.TStakingPoolAdded) // make to avoid nil map

	/**********************************************************************************************
	** We can just iterate over the syncMap and add the stakingPools to the map and slice.
	**********************************************************************************************/
	safeSyncMap(_opStakingSyncMap, chainID).Range(func(key, value interface{}) bool {
		stakingContract := value.(models.TStakingPoolAdded)
		if key == PerVault {
			asMap[stakingContract.VaultAddress] = stakingContract
		} else {
			asMap[stakingContract.StackingPoolAddress] = stakingContract
		}
		asSlice = append(asSlice, stakingContract)
		return true
	})
	return asMap, asSlice
}

/**************************************************************************************************
** GetOPStakingForVault
**************************************************************************************************/
func GetOPStakingForVault(chainID uint64, vault common.Address) (models.TStakingPoolAdded, bool) {
	/**********************************************************************************************
	** We can just iterate over the syncMap until we find the stakingPool for the vault.
	**********************************************************************************************/
	var stakingPool models.TStakingPoolAdded
	safeSyncMap(_opStakingSyncMap, chainID).Range(func(key, value interface{}) bool {
		stakingContract := value.(models.TStakingPoolAdded)
		if addresses.Equals(stakingContract.VaultAddress, vault) {
			stakingPool = stakingContract
			return false
		}
		return true
	})
	return stakingPool, stakingPool != (models.TStakingPoolAdded{})
}
