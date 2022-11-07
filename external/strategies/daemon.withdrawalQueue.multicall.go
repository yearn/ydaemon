package strategies

import (
	"math"
	"math/big"
	"strconv"

	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/common/ethereum"
	"github.com/yearn/ydaemon/common/helpers"
	"github.com/yearn/ydaemon/common/types/common"
)

func getVaultWithdrawalQueue(name string, index int64, contractAddress common.Address) ethereum.Call {
	parsedData, _ := yearnVaultABI.Pack("withdrawalQueue", big.NewInt(index))
	return ethereum.Call{
		Target:   contractAddress.Address,
		Abi:      yearnVaultABI,
		Method:   `withdrawalQueue`,
		CallData: parsedData,
		Name:     name + strconv.FormatInt(int64(index), 10),
	}
}

func getAllUniqueVaults(chainID uint64) []common.Address {
	var vaults []common.Address
	for _, strat := range Store.AggregatedStrategies[chainID] {
		vaults = append(vaults, strat.Vault)
	}
	return helpers.UniqueArrayAddress(vaults)
}

// FetchWithdrawalQueueMulticallData is used to get, for each available vaults, the withdrawalQueue order for it's strategies.
// We first grab all vaults from the strategies list (as it's the final object we want to update), remove all
// duplicates and then for each vault we fetch the withdrawalQueue for the 20 possible strategies per vault.
// Once done, we remove all 0x0 addresses from the list and return the result.
// We can then access, for each strategy, the withdrawal queue order.
func FetchWithdrawalQueueMulticallData(chainID uint64) {
	calls := make([]ethereum.Call, 0)
	caller := ethereum.MulticallClientForChainID[chainID]
	vaultList := getAllUniqueVaults(chainID)
	maxStrategiesPerVault := 20
	for _, vault := range vaultList {
		for i := 0; i < maxStrategiesPerVault; i++ {
			calls = append(calls, getVaultWithdrawalQueue(vault.String(), int64(i), vault))
		}
	}

	maxBatch := math.MaxInt64
	response := caller.ExecuteByBatch(calls, maxBatch, nil)
	withdrawQueueForStrategies := make(map[common.Address]int64)
	for _, vault := range vaultList {
		for i := 0; i < maxStrategiesPerVault; i++ {
			result := response[vault.String()+strconv.FormatInt(int64(i), 10)+`withdrawalQueue`]
			if len(result) == 1 {
				strategyAddress := common.FromAddress(result[0].(ethcommon.Address))
				if helpers.AddressIsValid(strategyAddress, chainID) {
					withdrawQueueForStrategies[strategyAddress] = int64(i)
				}
			}
		}
	}

	Store.WithdrawalQueueMultiCallData[chainID] = withdrawQueueForStrategies
	//No store in DB as it will be save via the AggregatedStrategies
}
