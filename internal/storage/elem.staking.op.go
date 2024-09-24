package storage

import (
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/common/addresses"
	"github.com/yearn/ydaemon/common/bigNumber"
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
func StoreOPStaking(chainID uint64, key string, stakingElement TStakingData) {
	safeSyncMap(_opStakingSyncMap, chainID).Store(key, stakingElement)
}

/**************************************************************************************************
** ListOPStaking will retrieve the all the vaults added to the registries from the
** given chainID
**************************************************************************************************/
func ListOPStaking(chainID uint64, key TStakingKey) (asMap map[common.Address]TStakingData, asSlice []TStakingData) {
	asMap = make(map[common.Address]TStakingData) // make to avoid nil map

	/**********************************************************************************************
	** We can just iterate over the syncMap and add the stakingPools to the map and slice.
	**********************************************************************************************/
	safeSyncMap(_opStakingSyncMap, chainID).Range(func(key, value interface{}) bool {
		stakingElement := value.(TStakingData)
		if key == PerVault {
			asMap[stakingElement.VaultAddress] = stakingElement
		} else {
			asMap[stakingElement.StakingAddress] = stakingElement
		}
		asSlice = append(asSlice, stakingElement)
		return true
	})
	return asMap, asSlice
}

/**************************************************************************************************
** GetOPStakingForVault
**************************************************************************************************/
func GetOPStakingForVault(chainID uint64, vault common.Address) (TStakingData, bool) {
	/**********************************************************************************************
	** We can just iterate over the syncMap until we find the stakingPool for the vault.
	**********************************************************************************************/
	var stakingPool TStakingData
	found := false

	safeSyncMap(_opStakingSyncMap, chainID).Range(func(key, value interface{}) bool {
		stakingElement := value.(TStakingData)
		if addresses.Equals(stakingElement.VaultAddress, vault) {
			stakingPool = stakingElement
			found = true
			return false
		}
		return true
	})
	return stakingPool, found
}

/**************************************************************************************************
** AssignOPStakingRewardAPY will update an entry in the _opStakingSyncMap syncMap to assign the
** APR to the staking pool.
**************************************************************************************************/
func AssignOPStakingRewardAPY(chainID uint64, vault common.Address, rewardToken common.Address, apy *bigNumber.Float) {
	/**********************************************************************************************
	** We can just iterate over the syncMap until we find the stakingContract for the vault.
	**********************************************************************************************/
	safeSyncMap(_opStakingSyncMap, chainID).Range(func(key, value interface{}) bool {
		stakingElement := value.(TStakingData)
		if addresses.Equals(stakingElement.VaultAddress, vault) {
			for i, reward := range stakingElement.RewardTokens {
				if addresses.Equals(reward.Address, rewardToken) {
					stakingElement.RewardTokens[i].APR = bigNumber.NewFloat().Clone(apy)
					safeSyncMap(_opStakingSyncMap, chainID).Store(key, stakingElement)
					return false
				}
			}
		}
		return true
	})
}
