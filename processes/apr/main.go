package apr

import (
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/common/env"
	"github.com/yearn/ydaemon/internal/storage"
	"github.com/yearn/ydaemon/processes/initDailyBlock"
)

var COMPUTED_APR = make(map[uint64]*sync.Map)

func init() {
	for chainID := range env.CHAINS {
		if _, ok := COMPUTED_APR[chainID]; !ok {
			COMPUTED_APR[chainID] = &sync.Map{}
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

func GetComputedAPR(chainID uint64, vaultAddress common.Address) (any, bool) {
	return safeSyncMap(COMPUTED_APR, chainID).Load(vaultAddress)
}

/**************************************************************************
** Function to calculate the APR for all the vaults in a chain.
**************************************************************************/
func ComputeChainAPR(chainID uint64) {
	allVaults, _ := storage.ListVaults(chainID)
	gauges := storage.FetchCurveGauges(chainID)
	pools := retrieveCurveGetPools(chainID)
	subgraphData := retrieveCurveSubgraphData(chainID)
	fraxPools := retrieveFraxPools()
	storage.RefreshGammaCalls(chainID)

	for _, vault := range allVaults {
		if vault.Metadata.IsRetired {
			continue
		}
		vaultToken, ok := storage.GetERC20(vault.ChainID, vault.Address)
		if !ok {
			continue
		}

		allStrategiesForVault, _ := storage.ListStrategiesForVault(chainID, vault.Address)
		vaultAPR := TVaultAPR{}
		if isV3Vault(vault) {
			if vault.Metadata.ShouldUseV2APR {
				vaultAPR = computeCurrentV2VaultAPR(vault, vaultToken)
			} else {
				vaultAPR = computeCurrentV3VaultAPR(vault, vaultToken)
			}
			vaultAPR.ForwardAPR = computeVaultV3ForwardAPR(
				vault,
				allStrategiesForVault,
			)
		} else {
			vaultAPR = computeCurrentV2VaultAPR(vault, vaultToken)
		}

		/**********************************************************************************************
		** Some vaults may have a staking rewards system. If so, we need to calculate the APR for
		** this staking rewards system and add it to the netAPY.
		**********************************************************************************************/
		stakingRewardAPR := computeStakingRewardsAPR(chainID, vault)
		vaultAPR.Extra.StakingRewardsAPR = stakingRewardAPR

		/**********************************************************************************************
		** If it's a Curve Vault (has a Curve, Convex or Frax strategy), we can estimate the forward
		** APR, aka the expected APR we will get for the upcoming period.
		** We need to compute it and store it in our ForwardAPR structure.
		**********************************************************************************************/
		if isCurveVault(allStrategiesForVault) {
			vaultAPR.ForwardAPR = computeCurveLikeForwardAPR(
				vault,
				allStrategiesForVault,
				gauges,
				pools,
				subgraphData,
				fraxPools,
			)
		}

		/**********************************************************************************************
		** If it's a Velo Vault (has a Velo or Aero strategy), we can estimate the forward APR, aka
		** the expected APR we will get for the upcoming period.
		** We need to compute it and store it in our ForwardAPR structure.
		**********************************************************************************************/
		if veloPool, ok := isVeloVault(chainID, vault); ok {
			vaultAPR.ForwardAPR = computeVeloLikeForwardAPR(
				vault,
				allStrategiesForVault,
				veloPool,
			)
		}
		if aeroPool, ok := isAeroVault(chainID, vault); ok {
			vaultAPR.ForwardAPR = computeVeloLikeForwardAPR(
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
			if extaRewardAPR, ok := calculateGammaExtraRewards(chainID, vault.AssetAddress); ok {
				vaultAPR.Extra.GammaRewardAPR = extaRewardAPR
			}
			vaultAPR.ForwardAPR = computeGammaForwardAPR(
				vault,
				allStrategiesForVault,
			)
			vaultAPR.ForwardAPR.Composite.RewardsAPR = vaultAPR.Extra.GammaRewardAPR
		}

		safeSyncMap(COMPUTED_APR, chainID).Store(vault.Address, vaultAPR)
	}
}

func Run(chainID uint64) {
	initYearnEcosystem(chainID)
	initDailyBlock.Run(chainID)
	ComputeChainAPR(chainID)
}
