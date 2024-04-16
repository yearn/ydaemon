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

/**************************************************************************************************
** IndexJuicedStakingContract is an indexer responsible for indexing the deployed Juiced staking
** contracts. Theses staking contracts are used in the Yearn ecosystem to allow users to stake some
** vaults tokens in it and earn some extra reards
**************************************************************************************************/
func IndexJuicedStakingContract(chainID uint64) (stakingMap map[common.Address]common.Address) {
	/**********************************************************************************************
	** Some staking contracts might be deployed outside of the Yearn ecosystem and are not indexed
	** in the Yearn registry. We need to index them manually.
	**********************************************************************************************/
	extraStakingContracts := env.CHAINS[chainID].ExtraStakingContracts
	if len(extraStakingContracts) > 0 {
		for _, contract := range extraStakingContracts {
			storage.StoreJuicedStaking(chainID, contract.VaultAddress, contract.StakingAddress)
		}
	}

	stakingContracts := env.CHAINS[chainID].StakingRewardRegistry
	if len(stakingContracts) == 0 {
		logs.Error(`No staking contract`)
		return
	}

	/**********************************************************************************************
	** Find the staking registry for the VEYFI gauges
	**********************************************************************************************/
	registry := env.TContractData{}
	for _, contract := range stakingContracts {
		if contract.Tag == `JUICED` {
			registry = contract
			break
		}
	}

	if (registry.Address == common.Address{}) {
		logs.Error(`No registry`)
		return
	}

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
	numberOfGauges, err := stakingRegistry.NumTokens(nil)
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
	for i := int64(0); i < int64(numberOfGauges.Int64()); i++ {
		calls = append(calls, multicalls.GetStakingTokenByIndex(
			strconv.FormatInt(i, 10),
			registry.Address,
			big.NewInt(i),
		))
	}
	callResponse := multicalls.Perform(chainID, calls, nil)
	for i := int64(0); i < int64(numberOfGauges.Int64()); i++ {
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

	/******************************************************************************************
	** Then, via a multicall, we need to call the `asset` method from the individual gauge
	** contract. This will give us the address of the vault assigned to that specific gauge,
	** and we will then be able to link the gauge to the vault.
	******************************************************************************************/
	assetResponse := multicalls.Perform(chainID, assetCalls, nil)
	for i := int64(0); i < int64(numberOfGauges.Int64()); i++ {
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
		storage.StoreJuicedStaking(chainID, vaultAddress, stakingAddress)
	}

	/**********************************************************************************************
	** This part is only accessed once, once the `historical` pools, aka that are already deployed
	** and indexed, are retrieved. New deployed pools will not hit this part, but will be handle
	** in the above functions directly
	**********************************************************************************************/
	stakingMap = storage.ListJuicedStaking(chainID)
	return stakingMap
}
