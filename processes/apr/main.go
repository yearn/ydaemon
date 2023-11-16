package apr

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/internal/storage"
	"github.com/yearn/ydaemon/processes/initDailyBlock"
)

var COMPUTED_APR = make(map[uint64]map[common.Address]TVaultAPR)

func ComputeChainAPR(chainID uint64) {
	allVaults, _ := storage.ListVaults(chainID)
	gauges := retrieveCurveGauges(chainID)
	pools := retrieveCurveGetPools(chainID)
	subgraphData := retrieveCurveSubgraphData(chainID)
	fraxPools := retrieveFraxPools(chainID)

	if COMPUTED_APR[chainID] == nil {
		COMPUTED_APR[chainID] = make(map[common.Address]TVaultAPR)
	}

	for _, vault := range allVaults {
		if vault.IsRetired {
			continue
		}
		vaultToken, ok := storage.GetERC20(vault.ChainID, vault.Address)
		if !ok {
			continue
		}

		allStrategiesForVault, _ := storage.ListStrategiesForVault(chainID, vault.Address)
		vaultAPR := TVaultAPR{}
		if isV3Vault(vault) {
			vaultAPR = computeCurrentV3VaultAPR(vault, vaultToken)
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

		COMPUTED_APR[chainID][vault.Address] = vaultAPR
	}
}

func Run(chainID uint64) {
	initYearnEcosystem(chainID)
	initDailyBlock.Run(chainID)
	ComputeChainAPR(chainID)
}
