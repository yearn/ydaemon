package storage

import (
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/common/addresses"
	"github.com/yearn/ydaemon/common/bigNumber"
)

var _v3StakingSyncMap = make(map[uint64]*sync.Map)

/**************************************************************************************************
** StoreV3Staking will add a new entry in the _v3StakingSyncMap syncMap
**************************************************************************************************/
func StoreV3Staking(chainID uint64, key string, stakingElement TStakingData) {
	safeSyncMap(_v3StakingSyncMap, chainID).Store(key, stakingElement)
}

/**************************************************************************************************
** ListV3StakingData will retrieve the all the vaults added to the registries from the
** given chainID
**************************************************************************************************/
func ListV3StakingData(chainID uint64) (asMap map[common.Address]TStakingData) {
	asMap = make(map[common.Address]TStakingData) // make to avoid nil map

	/**********************************************************************************************
	** We can just iterate over the syncMap and add the stakingContracts to the map and slice.
	**********************************************************************************************/
	safeSyncMap(_v3StakingSyncMap, chainID).Range(func(key, value interface{}) bool {
		stakingElement := value.(TStakingData)
		asMap[stakingElement.VaultAddress] = stakingElement
		return true
	})
	return asMap
}

/**************************************************************************************************
** GetV3StakingDataForVault
**************************************************************************************************/
func GetV3StakingDataForVault(chainID uint64, vault common.Address) (TStakingData, bool) {
	/**********************************************************************************************
	** We can just iterate over the syncMap until we find the stakingContract for the vault.
	**********************************************************************************************/
	var stakingContract TStakingData
	found := false

	safeSyncMap(_v3StakingSyncMap, chainID).Range(func(key, value interface{}) bool {
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
** AssignV3StakingRewardAPY will update an entry in the _v3StakingSyncMap syncMap to assign the
** APR to the staking pool.
**************************************************************************************************/
func AssignV3StakingRewardAPY(chainID uint64, vault common.Address, rewardToken common.Address, apy *bigNumber.Float) {
	/**********************************************************************************************
	** We can just iterate over the syncMap until we find the stakingContract for the vault.
	**********************************************************************************************/
	safeSyncMap(_v3StakingSyncMap, chainID).Range(func(key, value interface{}) bool {
		stakingElement := value.(TStakingData)
		if addresses.Equals(stakingElement.VaultAddress, vault) {
			for i, reward := range stakingElement.RewardTokens {
				if addresses.Equals(reward.Address, rewardToken) {
					stakingElement.RewardTokens[i].APR = bigNumber.NewFloat().Clone(apy)
					safeSyncMap(_v3StakingSyncMap, chainID).Store(key, stakingElement)
					return false
				}
			}
		}
		return true
	})
}
