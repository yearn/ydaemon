package apr

import (
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/common/env"
	"github.com/yearn/ydaemon/internal/storage"
	"github.com/yearn/ydaemon/processes/initDailyBlock"
)

var COMPUTED_APY = make(map[uint64]*sync.Map)

func init() {
	for chainID := range env.GetChains() {
		if _, ok := COMPUTED_APY[chainID]; !ok {
			COMPUTED_APY[chainID] = &sync.Map{}
		}
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
	allVaults, _ := storage.ListVaults(chainID)
	gauges := storage.FetchCurveGauges(chainID)
	pools := retrieveCurveGetPools(chainID)
	subgraphData := retrieveCurveSubgraphData(chainID)
	fraxPools := retrieveFraxPools()
	storage.RefreshGammaCalls(chainID)

	isOnGnosis := (chainID == 100)
	for _, vault := range allVaults {
		if vault.Metadata.IsRetired && !isOnGnosis {
			continue
		}
		vaultToken, ok := storage.GetERC20(vault.ChainID, vault.Address)
		if !ok {
			continue
		}

		allStrategiesForVault, _ := storage.ListStrategiesForVault(chainID, vault.Address)
		vaultAPY := TVaultAPY{}
		if isV3Vault(vault) {
			if vault.Metadata.ShouldUseV2APR {
				vaultAPY = computeCurrentV2VaultAPY(vault, vaultToken)
			} else {
				vaultAPY = computeCurrentV3VaultAPY(vault, vaultToken)
			}
			vaultAPY.ForwardAPY = computeVaultV3ForwardAPY(
				vault,
				allStrategiesForVault,
			)
		} else {
			vaultAPY = computeCurrentV2VaultAPY(vault, vaultToken)
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

		/**********************************************************************************************
		** If it's a Curve Vault (has a Curve, Convex or Frax strategy), we can estimate the forward
		** APY, aka the expected APY we will get for the upcoming period.
		** We need to compute it and store it in our ForwardAPY structure.
		**********************************************************************************************/
		if isCurveVault(allStrategiesForVault) {
			vaultAPY.ForwardAPY = computeCurveLikeForwardAPY(
				vault,
				allStrategiesForVault,
				gauges,
				pools,
				subgraphData,
				fraxPools,
			)
		}

		/**********************************************************************************************
		** If it's a Velo Vault (has a Velo or Aero strategy), we can estimate the forward APY, aka
		** the expected APY we will get for the upcoming period.
		** We need to compute it and store it in our ForwardAPY structure.
		**********************************************************************************************/
		if veloPool, ok := isVeloVault(chainID, vault); ok {
			vaultAPY.ForwardAPY = computeVeloLikeForwardAPY(
				vault,
				allStrategiesForVault,
				veloPool,
			)
		}
		if aeroPool, ok := isAeroVault(chainID, vault); ok {
			vaultAPY.ForwardAPY = computeVeloLikeForwardAPY(
				vault,
				allStrategiesForVault,
				aeroPool,
			)
		}

		/**********************************************************************************************
		** If it's a Gamma Vault, we can get the feeAPR as an estimate for the upcoming period, and we
		** can retrieve the extraReward APRs.
		**********************************************************************************************/
		if isGammaVault(chainID, vault) {
			if _, extaRewardAPY, ok := calculateGammaExtraRewards(chainID, vault.AssetAddress); ok {
				vaultAPY.Extra.GammaRewardAPY = extaRewardAPY
			}
			vaultAPY.ForwardAPY = computeGammaForwardAPY(
				vault,
				allStrategiesForVault,
			)
			vaultAPY.ForwardAPY.Composite.RewardsAPY = vaultAPY.Extra.GammaRewardAPY
		}

		if isPendleVault(chainID, vault) {
			vaultAPY.ForwardAPY = computePendleForwardAPY(
				vault,
				allStrategiesForVault,
			)
		}

		safeSyncMap(COMPUTED_APY, chainID).Store(vault.Address, vaultAPY)
	}
}

func Run(chainID uint64) {
	initYearnEcosystem(chainID)
	initDailyBlock.Run(chainID)
	ComputeChainAPY(chainID)
}
