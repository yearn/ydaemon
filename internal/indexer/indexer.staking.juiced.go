package indexer

import (
	"math/big"
	"strconv"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/common/bigNumber"
	"github.com/yearn/ydaemon/common/contracts"
	"github.com/yearn/ydaemon/common/env"
	"github.com/yearn/ydaemon/common/ethereum"
	"github.com/yearn/ydaemon/common/helpers"
	"github.com/yearn/ydaemon/common/logs"
	"github.com/yearn/ydaemon/internal/multicalls"
	"github.com/yearn/ydaemon/internal/storage"
)

func retrieveJuicedStakingData(chainID uint64, vaultAddresses []common.Address, stakingAddresses []common.Address) []storage.TStakingData {
	/**********************************************************************************************
	** First, for all staking contracts, we need to get the reward tokens. This is done via a
	** multicall, as we need to get the reward tokens for the first 10 indexes for each staking
	** contract.
	**********************************************************************************************/
	calls := []ethereum.Call{}
	rewardTokens := map[common.Address][]common.Address{}
	for _, stakingContract := range stakingAddresses {
		for i := 0; i < 10; i++ {
			calls = append(calls, multicalls.GetRewardTokens(stakingContract.Hex()+strconv.Itoa(i), stakingContract, int64(i)))
		}
	}

	response := multicalls.Perform(chainID, calls, nil)
	for _, stakingContract := range stakingAddresses {
		if _, ok := rewardTokens[stakingContract]; !ok {
			rewardTokens[stakingContract] = []common.Address{}
		}

		for i := 0; i < 10; i++ {
			rewardToken := helpers.DecodeAddress(response[stakingContract.Hex()+strconv.Itoa(i)+`rewardTokens`])
			if (rewardToken != common.Address{}) {
				rewardTokens[stakingContract] = append(rewardTokens[stakingContract], rewardToken)
			}
		}
	}

	/**********************************************************************************************
	** Once we have the indexes of the reward tokens, we can do another multicall to get, for each
	** of them, the reward data, the decimals of the reward token
	**********************************************************************************************/
	rewardTokensCalls := []ethereum.Call{}
	for _, stakingContract := range stakingAddresses {
		for _, rewardToken := range rewardTokens[stakingContract] {
			rewardTokensCalls = append(rewardTokensCalls, multicalls.GetRewardData(stakingContract.Hex()+rewardToken.Hex(), stakingContract, rewardToken))
			rewardTokensCalls = append(rewardTokensCalls, multicalls.GetDecimals(stakingContract.Hex()+rewardToken.Hex(), rewardToken))
			rewardTokensCalls = append(rewardTokensCalls, multicalls.GetName(stakingContract.Hex()+rewardToken.Hex(), rewardToken))
			rewardTokensCalls = append(rewardTokensCalls, multicalls.GetSymbol(stakingContract.Hex()+rewardToken.Hex(), rewardToken))
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
			rewardDataKey := rewardTokensResponse[stakingContract.Hex()+rewardToken.Hex()+`rewardData`]
			rewardRate := bigNumber.NewInt(0)
			rewardDuration := bigNumber.NewInt(0)
			rewardPeriodFinish := bigNumber.NewInt(0)
			if len(rewardDataKey) > 0 {
				rewardDuration = bigNumber.SetInt(rewardDataKey[1].(*big.Int))
				rewardPeriodFinish = bigNumber.SetInt(rewardDataKey[2].(*big.Int))
				rewardRate = bigNumber.SetInt(rewardDataKey[3].(*big.Int))
			}

			now := time.Now().Unix()
			isFinished := rewardPeriodFinish.Uint64() > 0 && rewardPeriodFinish.Uint64() < uint64(now)
			rewardTokenData := storage.TRewardToken{
				Address:      rewardToken,
				Name:         helpers.DecodeString(rewardTokensResponse[stakingContract.Hex()+rewardToken.Hex()+`name`]),
				Symbol:       helpers.DecodeString(rewardTokensResponse[stakingContract.Hex()+rewardToken.Hex()+`symbol`]),
				Decimals:     helpers.DecodeUint64(rewardTokensResponse[stakingContract.Hex()+rewardToken.Hex()+`decimals`]),
				Duration:     rewardDuration.Uint64(),
				Rate:         rewardRate.Uint64(),
				PeriodFinish: rewardPeriodFinish.Uint64(),
				IsFinished:   isFinished,
			}
			staking.RewardTokens = append(staking.RewardTokens, rewardTokenData)
		}
		allStakingData = append(allStakingData, staking)
	}

	return allStakingData
}

/**************************************************************************************************
** IndexJuicedStakingContract is an indexer responsible for indexing the deployed Juiced staking
** contracts. Theses staking contracts are used in the Yearn ecosystem to allow users to stake some
** vaults tokens in it and earn some extra reards
**************************************************************************************************/
func IndexJuicedStakingContract(chainID uint64) (stakingMap map[common.Address]storage.TStakingData) {
	allVaults := []common.Address{}
	allStaking := []common.Address{}
	/**********************************************************************************************
	** Some staking contracts might be deployed outside of the Yearn ecosystem and are not indexed
	** in the Yearn registry. We need to index them manually.
	**********************************************************************************************/
	extraStakingContracts := env.CHAINS[chainID].ExtraStakingContracts
	if len(extraStakingContracts) > 0 {
		for _, contract := range extraStakingContracts {
			if contract.Tag == `JUICED` {
				allVaults = append(allVaults, contract.VaultAddress)
				allStaking = append(allStaking, contract.StakingAddress)
			}
		}
	}

	stakingContracts := env.CHAINS[chainID].StakingRewardRegistry
	if len(stakingContracts) == 0 && len(allVaults) == 0 {
		logs.Error(`No staking contract`)
		return
	}

	/**********************************************************************************************
	** Find the staking registry for the Juiced staking contracts
	**********************************************************************************************/
	registry := env.TContractData{}
	for _, contract := range stakingContracts {
		if contract.Tag == `JUICED` {
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
		stakingRegistry, err := contracts.NewJuicedStakingRewardsRegistry(registry.Address, client)
		if err != nil {
			logs.Error(`Failed to connect to the juiced staking registry contract`, err)
			return
		}

		/******************************************************************************************
		** We first need to get the number of vaults that have a staking contract attached to them.
		******************************************************************************************/
		numberOfTokens, err := stakingRegistry.NumTokens(nil)
		if err != nil {
			logs.Error(`Failed to retrieve the number of staking contract`, err)
			return
		}

		/******************************************************************************************
		** Then, via a multicall, we need to call the `tokens(idx)` method from the stakingRegistry
		** contract. This will give us the address of the vault for a given index. Once we have
		** that, we can know two things: the index of a vault and the address of that vault.
		******************************************************************************************/
		calls := []ethereum.Call{}
		assetCalls := []ethereum.Call{}
		for i := int64(0); i < int64(numberOfTokens.Int64()); i++ {
			calls = append(calls, multicalls.GetStakingTokenByIndex(
				strconv.FormatInt(i, 10),
				registry.Address,
				big.NewInt(i),
			))
		}
		callResponse := multicalls.Perform(chainID, calls, nil)
		for i := int64(0); i < int64(numberOfTokens.Int64()); i++ {
			rawVaultAddress := callResponse[strconv.FormatInt(i, 10)+`tokens`]
			if len(rawVaultAddress) == 0 {
				continue
			}
			vaultAddress := helpers.DecodeAddress(rawVaultAddress)
			assetCalls = append(assetCalls, multicalls.GetStakingPoolForVault(
				vaultAddress.Hex(),
				registry.Address,
				vaultAddress,
			))
		}

		assetResponse := multicalls.Perform(chainID, assetCalls, nil)

		for i := int64(0); i < int64(numberOfTokens.Int64()); i++ {
			rawVaultAddress := callResponse[strconv.FormatInt(i, 10)+`tokens`]
			if len(rawVaultAddress) == 0 {
				continue
			}
			vaultAddress := helpers.DecodeAddress(rawVaultAddress)

			rawStakingAddress := assetResponse[vaultAddress.Hex()+`stakingPool`]
			if len(rawStakingAddress) == 0 {
				continue
			}
			stakingAddress := helpers.DecodeAddress(rawStakingAddress)
			allVaults = append(allVaults, vaultAddress)
			allStaking = append(allStaking, stakingAddress)
		}
	}

	result := retrieveJuicedStakingData(chainID, allVaults, allStaking)

	/******************************************************************************************
	** Then, via a multicall, we need to call the `asset` method from the individual gauge
	** contract. This will give us the address of the vault assigned to that specific gauge,
	** and we will then be able to link the gauge to the vault.
	******************************************************************************************/
	for _, stakingElement := range result {
		key := stakingElement.VaultAddress.Hex() + stakingElement.StakingAddress.Hex()
		storage.StoreJuicedStaking(chainID, key, stakingElement)
	}

	/**********************************************************************************************
	** This part is only accessed once, once the `historical` pools, aka that are already deployed
	** and indexed, are retrieved. New deployed pools will not hit this part, but will be handle
	** in the above functions directly
	**********************************************************************************************/
	stakingMap = storage.ListJuicedStakingData(chainID)
	return stakingMap
}
