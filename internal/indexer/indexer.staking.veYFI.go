package indexer

import (
	"math/big"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/common/contracts"
	"github.com/yearn/ydaemon/common/env"
	"github.com/yearn/ydaemon/common/ethereum"
	"github.com/yearn/ydaemon/common/helpers"
	"github.com/yearn/ydaemon/common/logs"
	"github.com/yearn/ydaemon/internal/multicalls"
	"github.com/yearn/ydaemon/internal/storage"
)

func retrieveVEYFIStakingData(chainID uint64, vaultAddresses []common.Address, stakingAddresses []common.Address) []storage.TStakingData {
	/**********************************************************************************************
	** First, for all staking contracts, we need to get the reward tokens. This is done via a
	** multicall, as we need to get the reward tokens for the first 10 indexes for each staking
	** contract.
	**********************************************************************************************/
	rewardTokens := map[common.Address][]common.Address{}
	for _, stakingContract := range stakingAddresses {
		if _, ok := rewardTokens[stakingContract]; !ok {
			rewardTokens[stakingContract] = []common.Address{}
		}
		rewardTokens[stakingContract] = append(rewardTokens[stakingContract], common.HexToAddress(`0x41252E8691e964f7DE35156B68493bAb6797a275`))
	}

	/******************************************************************************************
	** Once we have that, the flow is pretty easy: one token, multiple info, but no need to
	** look at multiple rewards tokens.
	******************************************************************************************/
	rewardTokensCalls := []ethereum.Call{}
	for _, stakingContract := range stakingAddresses {
		for _, rewardToken := range rewardTokens[stakingContract] {
			rewardTokensCalls = append(rewardTokensCalls, multicalls.GetName(stakingContract.Hex()+rewardToken.Hex(), rewardToken))
			rewardTokensCalls = append(rewardTokensCalls, multicalls.GetSymbol(stakingContract.Hex()+rewardToken.Hex(), rewardToken))
			rewardTokensCalls = append(rewardTokensCalls, multicalls.GetDecimals(stakingContract.Hex()+rewardToken.Hex(), rewardToken))
			rewardTokensCalls = append(rewardTokensCalls, multicalls.GetRewardsDuration(stakingContract.Hex()+rewardToken.Hex(), stakingContract))
			rewardTokensCalls = append(rewardTokensCalls, multicalls.GetRewardRate(stakingContract.Hex()+rewardToken.Hex(), stakingContract))
			rewardTokensCalls = append(rewardTokensCalls, multicalls.GetPeriodFinish(stakingContract.Hex()+rewardToken.Hex(), stakingContract))
		}
	}
	rewardTokensResponse := multicalls.Perform(chainID, rewardTokensCalls, nil)
	allStakingData := []storage.TStakingData{}
	for i, stakingContract := range stakingAddresses {
		staking := storage.TStakingData{
			VaultAddress:   vaultAddresses[i],
			StakingAddress: stakingContract,
			RewardTokens:   []storage.TRewardToken{},
		}
		for _, rewardToken := range rewardTokens[stakingContract] {
			duration := helpers.DecodeBigInt(rewardTokensResponse[stakingContract.Hex()+rewardToken.Hex()+`rewardsDuration`])
			rate := helpers.DecodeBigInt(rewardTokensResponse[stakingContract.Hex()+rewardToken.Hex()+`rewardRate`])
			periodFinish := helpers.DecodeBigInt(rewardTokensResponse[stakingContract.Hex()+rewardToken.Hex()+`periodFinish`])
			rewardTokenData := storage.TRewardToken{
				Address:      rewardToken,
				Name:         helpers.DecodeString(rewardTokensResponse[stakingContract.Hex()+rewardToken.Hex()+`name`]),
				Symbol:       helpers.DecodeString(rewardTokensResponse[stakingContract.Hex()+rewardToken.Hex()+`symbol`]),
				Decimals:     helpers.DecodeUint64(rewardTokensResponse[stakingContract.Hex()+rewardToken.Hex()+`decimals`]),
				Duration:     duration.Uint64(),
				Rate:         rate.Uint64(),
				PeriodFinish: periodFinish.Uint64(),
				IsFinished:   false,
			}
			staking.RewardTokens = append(staking.RewardTokens, rewardTokenData)
		}
		allStakingData = append(allStakingData, staking)
	}
	return allStakingData
}

/**************************************************************************************************
** IndexVeYFIStakingContract is an indexer responsible for indexing the deployed veYFI staking
** contracts. Theses staking contracts, also called gauges, are used in the Yearn ecosystem to
** allow users to stake some vaults tokens in it and earn some dYFI rewards.
**************************************************************************************************/
func IndexVeYFIStakingContract(chainID uint64) (veYFIGaugesMap map[common.Address]storage.TStakingData) {
	allVaults := []common.Address{}
	allStaking := []common.Address{}

	/**********************************************************************************************
	** Some staking contracts might be deployed outside of the Yearn ecosystem and are not indexed
	** in the Yearn registry. We need to index them manually.
	**********************************************************************************************/
	extraStakingContracts := env.CHAINS[chainID].ExtraStakingContracts
	if len(extraStakingContracts) > 0 {
		for _, contract := range extraStakingContracts {
			if contract.Tag == `VEYFI` {
				allVaults = append(allVaults, contract.VaultAddress)
				allStaking = append(allStaking, contract.StakingAddress)
			}
		}
	}

	stakingContracts := env.CHAINS[chainID].StakingRewardRegistry
	if len(stakingContracts) == 0 && len(allVaults) == 0 {
		return
	}

	/**********************************************************************************************
	** Find the staking registry for the VEYFI gauges
	**********************************************************************************************/
	registry := env.TContractData{}
	for _, contract := range stakingContracts {
		if contract.Tag == `VEYFI` {
			registry = contract
			break
		}
	}

	if (registry.Address != common.Address{}) {
		/******************************************************************************************
		** In order to get the gauges, 3 different calls are required, 2 of them are multicalls.
		** They are performed one after the other, as the result of the first call is required for
		** the second call.
		******************************************************************************************/
		client := ethereum.GetRPC(chainID)
		gaugeRegistry, err := contracts.NewYGaugeRegistry(registry.Address, client)
		if err != nil {
			logs.Error(`Failed to connect to the gauge registry contract`, err)
			return
		}

		/******************************************************************************************
		** We first need to get the number of gauges via the gaugeRegistry `VaultCount` method.
		** This will give us the number of vaults having a gauges.
		******************************************************************************************/
		numberOfGauges, err := gaugeRegistry.VaultCount(nil)
		if err != nil {
			logs.Error(`Failed to retrieve the number of gauges`, err)
			return
		}

		/******************************************************************************************
		** Then, via a multicall, we need to call the `gauges(idx)` method from the gaugeRegistry
		** contract. This will give us the address of the gauge for a given index. Once we have
		** that, we can know two things: the index of a gauge and the address of that gauge.
		******************************************************************************************/
		calls := []ethereum.Call{}
		assetCalls := []ethereum.Call{}
		for i := int64(0); i < int64(numberOfGauges.Int64()); i++ {
			calls = append(calls, multicalls.GetVeYFIGaugeByIndex(
				strconv.FormatInt(i, 10),
				registry.Address,
				big.NewInt(i),
			))
		}
		callResponse := multicalls.Perform(chainID, calls, nil)
		for i := int64(0); i < int64(numberOfGauges.Int64()); i++ {
			rawGaugeAddress := callResponse[strconv.FormatInt(i, 10)+`gauges`]
			if len(rawGaugeAddress) == 0 {
				continue
			}
			gaugeAddress := helpers.DecodeAddress(rawGaugeAddress)
			assetCalls = append(assetCalls, multicalls.GetAsset(gaugeAddress.Hex(), gaugeAddress))
		}

		/******************************************************************************************
		** Then, via a multicall, we need to call the `asset` method from the individual gauge
		** contract. This will give us the address of the vault assigned to that specific gauge,
		** and we will then be able to link the gauge to the vault.
		******************************************************************************************/
		assetResponse := multicalls.Perform(chainID, assetCalls, nil)
		for i := int64(0); i < int64(numberOfGauges.Int64()); i++ {
			rawGaugeAddress := callResponse[strconv.FormatInt(i, 10)+`gauges`]
			if len(rawGaugeAddress) == 0 {
				continue
			}
			gaugeAddress := helpers.DecodeAddress(rawGaugeAddress)

			rawAssetAddress := assetResponse[gaugeAddress.Hex()+`asset`]
			if len(rawAssetAddress) == 0 {
				continue
			}
			assetAddress := helpers.DecodeAddress(rawAssetAddress)
			allVaults = append(allVaults, assetAddress)
			allStaking = append(allStaking, gaugeAddress)
		}
	}

	result := retrieveVEYFIStakingData(chainID, allVaults, allStaking)
	/******************************************************************************************
	** Then, via a multicall, we need to call the `asset` method from the individual gauge
	** contract. This will give us the address of the vault assigned to that specific gauge,
	** and we will then be able to link the gauge to the vault.
	******************************************************************************************/
	for _, stakingElement := range result {
		key := stakingElement.VaultAddress.Hex() + stakingElement.StakingAddress.Hex()
		storage.StoreVeYFIStaking(chainID, key, stakingElement)
	}

	/** 🔵 - Yearn *********************************************************************************
	** This part is only accessed once, once the `historical` pools, aka that are already deployed
	** and indexed, are retrieved. New deployed pools will not hit this part, but will be handle
	** in the above functions directly
	**********************************************************************************************/
	stakingMap := storage.ListVeYFIStaking(chainID)
	return stakingMap
}
