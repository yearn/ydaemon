package apr

import (
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/common/addresses"
	"github.com/yearn/ydaemon/common/env"
	"github.com/yearn/ydaemon/common/logs"
	"github.com/yearn/ydaemon/internal/storage"
)

func isForwardAPYAllZeros(f TForwardAPY) bool {
	if f.NetAPY != nil && !f.NetAPY.IsZero() {
		return false
	}
	c := f.Composite
	if c.Boost != nil && !c.Boost.IsZero() {
		return false
	}
	if c.PoolAPY != nil && !c.PoolAPY.IsZero() {
		return false
	}
	if c.BoostedAPR != nil && !c.BoostedAPR.IsZero() {
		return false
	}
	if c.BaseAPR != nil && !c.BaseAPR.IsZero() {
		return false
	}
	if c.CvxAPR != nil && !c.CvxAPR.IsZero() {
		return false
	}
	if c.RewardsAPY != nil && !c.RewardsAPY.IsZero() {
		return false
	}
	if c.V3OracleCurrentAPR != nil && !c.V3OracleCurrentAPR.IsZero() {
		return false
	}
	if c.V3OracleStratRatioAPR != nil && !c.V3OracleStratRatioAPR.IsZero() {
		return false
	}
	if c.KeepCRV != nil && !c.KeepCRV.IsZero() {
		return false
	}
	if c.KeepVelo != nil && !c.KeepVelo.IsZero() {
		return false
	}
	return true
}

var COMPUTED_APY = make(map[uint64]*sync.Map)

func init() {
	for chainID := range env.GetChains() {
		if _, ok := COMPUTED_APY[chainID]; !ok {
			COMPUTED_APY[chainID] = &sync.Map{}
		}
	}
}

/**************************************************************************
** LoadPersistedAPY loads the persisted APY data from disk into the COMPUTED_APY map
**************************************************************************/
func LoadPersistedAPY(chainID uint64) {
	apyMap, _ := storage.ListAPY(chainID)
	for vaultAddress, apy := range apyMap {
		safeSyncMap(COMPUTED_APY, chainID).Store(vaultAddress, apy)
	}
}

/**************************************************************************
** Little helper to ensure that the sync map is initialized before use.
**************************************************************************/
func safeSyncMap(source map[uint64]*sync.Map, chainID uint64) *sync.Map {
	syncMap := source[chainID]
	if syncMap == nil {
		syncMap = &sync.Map{}
		source[chainID] = syncMap
	}
	return syncMap
}

func GetComputedAPY(chainID uint64, vaultAddress common.Address) (any, bool) {
	return safeSyncMap(COMPUTED_APY, chainID).Load(vaultAddress)
}

/**************************************************************************
** Function to calculate the APY for all the vaults in a chain.
**************************************************************************/
func ComputeChainAPY(chainID uint64) {
	start := time.Now()
	logs.Warning("📈 [APY START]", "chain", chainID)
	allVaults, _ := storage.ListVaults(chainID)
	storage.RefreshGammaCalls(chainID)

	isOnGnosis := (chainID == 100)
	computedAPYData := make(map[common.Address]TVaultAPY)

	for _, vault := range allVaults {
		// Adding an exception for the vault that is retired but we still want to compute. Alchemix related
		isException := addresses.Equals(vault.Address, "0xaD17A225074191d5c8a37B50FdA1AE278a2EE6A2") ||
			addresses.Equals(vault.Address, "0x5B977577Eb8a480f63e11FC615D6753adB8652Ae") ||
			addresses.Equals(vault.Address, "0x65343F414FFD6c97b0f6add33d16F6845Ac22BAc") ||
			addresses.Equals(vault.Address, "0xFaee21D0f0Af88EE72BB6d68E54a90E6EC2616de")
		shouldSkip := false
		if vault.Metadata.IsRetired {
			shouldSkip = true
			if isOnGnosis {
				shouldSkip = false
			}
			if isException {
				shouldSkip = false
			}
		}
		if shouldSkip {
			continue
		}
		
		allStrategiesForVault, _ := storage.ListStrategiesForVault(chainID, vault.Address)
		vaultAPY := TVaultAPY{}
		if isV3Vault(vault) {
			if vault.Metadata.ShouldUseV2APR {
				vaultAPY = computeCurrentV2VaultAPY(vault)
			} else {
				vaultAPY = computeCurrentV3VaultAPY(vault)
			}
			vaultAPY.ForwardAPY = computeVaultV3ForwardAPY(
				vault,
				allStrategiesForVault,
			)
		} else {
			vaultAPY = computeCurrentV2VaultAPY(vault)
		}

		/**********************************************************************************************
		** Some vaults may have a staking rewards system. If so, we need to calculate the APY for
		** this staking rewards system and add it to the netAPY.
		**********************************************************************************************/
		_, stakingRewardAPY, hasExtraAPR := computeOPBoostStakingRewardsAPY(chainID, vault)
		if hasExtraAPR {
			vaultAPY.Extra.StakingRewardsAPY = stakingRewardAPY
		}

		_, veYFIGaugeStakingAPY, hasExtraAPR := computeVeYFIGaugeStakingRewardsAPY(chainID, vault)
		if hasExtraAPR {
			vaultAPY.Extra.StakingRewardsAPY = veYFIGaugeStakingAPY
		}

		_, juicedStakingAPY, hasExtraAPR := computeJuicedStakingRewardsAPY(chainID, vault)
		if hasExtraAPR {
			vaultAPY.Extra.StakingRewardsAPY = juicedStakingAPY
		}

		_, v3StakingAPY, hasExtraAPR := computeV3StakingRewardsAPY(chainID, vault)
		if hasExtraAPR {
			vaultAPY.Extra.StakingRewardsAPY = v3StakingAPY
		}

		kongForwardAPY, _ := convertKongEstimatedAprToForwardAPY(chainID, vault.Address)
		if isForwardAPYAllZeros(kongForwardAPY) {
			kongForwardAPY.Type = ""
		}
		vaultAPY.ForwardAPY = kongForwardAPY
		

		safeSyncMap(COMPUTED_APY, chainID).Store(vault.Address, vaultAPY)
		computedAPYData[vault.Address] = vaultAPY
	}

	// Save the computed APY data to disk
	storage.StoreAPYToJson(chainID, computedAPYData)
	logs.Success("📈 [APY DONE]", "chain", chainID, "took", time.Since(start))
	logs.Success(chainID, `-`, `ComputeChainAPY ✅`) // Legacy format for deploy workflow detection
}
