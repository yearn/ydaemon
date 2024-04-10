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
** The function `IndexStakingPools` ...
**************************************************************************************************/
func IndexVeYFIGauges(chainID uint64) (veYFIGaugesMap map[common.Address]common.Address) {
	stakingContracts := env.CHAINS[chainID].StakingRewardContract
	if len(stakingContracts) == 0 {
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

	if (registry.Address == common.Address{}) {
		return
	}

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
		storage.StoreVeYFIGauge(chainID, assetAddress, gaugeAddress)
	}

	/**********************************************************************************************
	** This part is only accessed once, once the `historical` pools, aka that are already deployed
	** and indexed, are retrieved. New deployed pools will not hit this part, but will be handle
	** in the above functions directly
	**********************************************************************************************/
	veYFIGaugesMap = storage.ListVeYFIGauges(chainID)
	return veYFIGaugesMap
}
