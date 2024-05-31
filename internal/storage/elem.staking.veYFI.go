package storage

import (
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/common/addresses"
	"github.com/yearn/ydaemon/common/bigNumber"
)

var _veYFIStakingSyncMap = make(map[uint64]*sync.Map)

/**************************************************************************************************
** StoreVeYFIStaking will add a new entry in the _veYFIStakingSyncMap syncMap
**************************************************************************************************/
func StoreVeYFIStaking(chainID uint64, key string, stakingElement TStakingData) {
	safeSyncMap(_veYFIStakingSyncMap, chainID).Store(key, stakingElement)
}

/**************************************************************************************************
** ListVeYFIStaking will retrieve the all the vaults added to the registries from the
** given chainID
**************************************************************************************************/
func ListVeYFIStaking(chainID uint64) (asMap map[common.Address]TStakingData) {
	asMap = make(map[common.Address]TStakingData) // make to avoid nil map

	/**********************************************************************************************
	** We can just iterate over the syncMap and add the stakingPools to the map and slice.
	**********************************************************************************************/
	safeSyncMap(_veYFIStakingSyncMap, chainID).Range(func(key, value interface{}) bool {
		stakingElement := value.(TStakingData)
		asMap[stakingElement.VaultAddress] = stakingElement
		return true
	})
	return asMap
}

/**************************************************************************************************
** GetVeYFIStakingForVault
**************************************************************************************************/
func GetVeYFIStakingForVault(chainID uint64, vault common.Address) (TStakingData, bool) {
	/**********************************************************************************************
	** We can just iterate over the syncMap until we find the stakingPool for the vault.
	**********************************************************************************************/
	var stakingContract TStakingData
	found := false

	safeSyncMap(_veYFIStakingSyncMap, chainID).Range(func(key, value interface{}) bool {
		stakingElement := value.(TStakingData)
		if addresses.Equals(stakingElement.VaultAddress, vault) {
			stakingContract = stakingElement
			found = true
			return false
		}
		return true
	})
	return stakingContract, found
}

/**************************************************************************************************
** AssignVEYFIStakingRewardAPR will update an entry in the _veYFIStakingSyncMap syncMap to assign
** the APR to the staking pool.
**************************************************************************************************/
func AssignVEYFIStakingRewardAPR(chainID uint64, vault common.Address, rewardToken common.Address, apr *bigNumber.Float) {
	/**********************************************************************************************
	** We can just iterate over the syncMap until we find the stakingContract for the vault.
	**********************************************************************************************/
	safeSyncMap(_veYFIStakingSyncMap, chainID).Range(func(key, value interface{}) bool {
		stakingElement := value.(TStakingData)
		if addresses.Equals(stakingElement.VaultAddress, vault) {
			for i, reward := range stakingElement.RewardTokens {
				if addresses.Equals(reward.Address, rewardToken) {
					stakingElement.RewardTokens[i].APR = bigNumber.NewFloat().Clone(apr)
					safeSyncMap(_veYFIStakingSyncMap, chainID).Store(key, stakingElement)
					return false
				}
			}
		}
		return true
	})
}
