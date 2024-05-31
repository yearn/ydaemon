package storage

import (
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/common/addresses"
	"github.com/yearn/ydaemon/common/bigNumber"
)

var _juicedStakingSyncMap = make(map[uint64]*sync.Map)

type TRewardToken struct {
	Address      common.Address
	Name         string
	Symbol       string
	Decimals     uint64
	Duration     uint64
	Rate         uint64
	PeriodFinish uint64
	IsFinished   bool
	APR          *bigNumber.Float
}
type TStakingData struct {
	VaultAddress   common.Address
	StakingAddress common.Address
	RewardTokens   []TRewardToken
}

/**************************************************************************************************
** StoreJuicedStaking will add a new entry in the _juicedStakingSyncMap syncMap
**************************************************************************************************/
func StoreJuicedStaking(chainID uint64, key string, stakingElement TStakingData) {
	safeSyncMap(_juicedStakingSyncMap, chainID).Store(key, stakingElement)
}

/**************************************************************************************************
** ListJuicedStakingData will retrieve the all the vaults added to the registries from the
** given chainID
**************************************************************************************************/
func ListJuicedStakingData(chainID uint64) (asMap map[common.Address]TStakingData) {
	asMap = make(map[common.Address]TStakingData) // make to avoid nil map

	/**********************************************************************************************
	** We can just iterate over the syncMap and add the stakingContracts to the map and slice.
	**********************************************************************************************/
	safeSyncMap(_juicedStakingSyncMap, chainID).Range(func(key, value interface{}) bool {
		stakingElement := value.(TStakingData)
		asMap[stakingElement.VaultAddress] = stakingElement
		return true
	})
	return asMap
}

/**************************************************************************************************
** GetJuicedStakingDataForVault
**************************************************************************************************/
func GetJuicedStakingDataForVault(chainID uint64, vault common.Address) (TStakingData, bool) {
	/**********************************************************************************************
	** We can just iterate over the syncMap until we find the stakingContract for the vault.
	**********************************************************************************************/
	var stakingContract TStakingData
	found := false

	safeSyncMap(_juicedStakingSyncMap, chainID).Range(func(key, value interface{}) bool {
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
** AssignJuicedStakingRewardAPR will update an entry in the _juicedStakingSyncMap syncMap to assign the
** APR to the staking pool.
**************************************************************************************************/
func AssignJuicedStakingRewardAPR(chainID uint64, vault common.Address, rewardToken common.Address, apr *bigNumber.Float) {
	/**********************************************************************************************
	** We can just iterate over the syncMap until we find the stakingContract for the vault.
	**********************************************************************************************/
	safeSyncMap(_juicedStakingSyncMap, chainID).Range(func(key, value interface{}) bool {
		stakingElement := value.(TStakingData)
		if addresses.Equals(stakingElement.VaultAddress, vault) {
			for i, reward := range stakingElement.RewardTokens {
				if addresses.Equals(reward.Address, rewardToken) {
					stakingElement.RewardTokens[i].APR = bigNumber.NewFloat().Clone(apr)
					safeSyncMap(_juicedStakingSyncMap, chainID).Store(key, stakingElement)
					return false
				}
			}
		}
		return true
	})
}
