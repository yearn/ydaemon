package store

import (
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/common/addresses"
	"github.com/yearn/ydaemon/internal/models"
)

var _stackingPoolsSyncMap = make(map[uint64]*sync.Map)

type TStakingKey string

const (
	PerVault TStakingKey = "PerVault"
	PerPool  TStakingKey = "PerPool"
)

/**************************************************************************************************
** AppendToSakingPoolMap will add a new stakingPool in the _stackingPoolsSyncMap
**************************************************************************************************/
func AppendToSakingPoolMap(chainID uint64, pool models.TStakingPoolAdded) {
	syncMap, ok := _stackingPoolsSyncMap[chainID]
	if !ok {
		syncMap = &sync.Map{}
		_stackingPoolsSyncMap[chainID] = syncMap
	}
	syncMap.Store(pool.StackingPoolAddress.Hex(), pool)
}

/**************************************************************************************************
** ListAllStakingPools will returns the registered staking pool for a given chain
** key should be perVault or perPool
**************************************************************************************************/
func ListAllStakingPools(chainID uint64, key TStakingKey) (asMap map[common.Address]models.TStakingPoolAdded, asSlice []models.TStakingPoolAdded) {
	asMap = make(map[common.Address]models.TStakingPoolAdded) // make to avoid nil map

	/**********************************************************************************************
	** We first retrieve the syncMap.
	**********************************************************************************************/
	syncMap := _stackingPoolsSyncMap[chainID]
	if syncMap == nil {
		syncMap = &sync.Map{}
		_stackingPoolsSyncMap[chainID] = syncMap
	}

	/**********************************************************************************************
	** We can just iterate over the syncMap and add the stakingPools to the map and slice.
	**********************************************************************************************/
	syncMap.Range(func(key, value interface{}) bool {
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
** GetStakingPoolForVault will returns the registered staking pool for a given chain and vault
**************************************************************************************************/
func GetStakingPoolForVault(chainID uint64, vault common.Address) (models.TStakingPoolAdded, bool) {
	/**********************************************************************************************
	** We first retrieve the syncMap.
	**********************************************************************************************/
	syncMap := _stackingPoolsSyncMap[chainID]
	if syncMap == nil {
		syncMap = &sync.Map{}
		_stackingPoolsSyncMap[chainID] = syncMap
	}

	/**********************************************************************************************
	** We can just iterate over the syncMap until we find the stakingPool for the vault.
	**********************************************************************************************/
	var stakingPool models.TStakingPoolAdded
	syncMap.Range(func(key, value interface{}) bool {
		stakingContract := value.(models.TStakingPoolAdded)
		if addresses.Equals(stakingContract.VaultAddress, vault) {
			stakingPool = stakingContract
			return false
		}
		return true
	})
	return stakingPool, stakingPool != (models.TStakingPoolAdded{})
}
