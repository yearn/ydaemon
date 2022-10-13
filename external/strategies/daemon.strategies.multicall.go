package strategies

import (
	"math"
	"math/big"
	"strconv"
	"strings"
	"sync"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/yearn/ydaemon/common/bigNumber"
	"github.com/yearn/ydaemon/common/contracts"
	"github.com/yearn/ydaemon/common/ethereum"
	"github.com/yearn/ydaemon/common/helpers"
	"github.com/yearn/ydaemon/common/logs"
	"github.com/yearn/ydaemon/common/models"
	"github.com/yearn/ydaemon/common/store"
	"github.com/yearn/ydaemon/common/types/common"
)

// yearnVaultABI takes the ABI of the standard Yearn Vault contract and prepare it for the multicall
var yearnVaultABI, _ = contracts.YearnVaultMetaData.GetAbi()
var yearnStrategyABI, _ = contracts.StrategyBaseMetaData.GetAbi()

func getCreditAvailable(name string, contractAddress common.Address, strategyAddress common.Address, version string) ethereum.Call {
	parsedData, _ := yearnVaultABI.Pack("creditAvailable0", strategyAddress.Address)
	return ethereum.Call{
		Target:   contractAddress.Address,
		Abi:      yearnVaultABI,
		Method:   `creditAvailable0`,
		CallData: parsedData,
		Name:     name,
		Version:  version,
	}
}

func getDebtOutstanding(name string, contractAddress common.Address, strategyAddress common.Address, version string) ethereum.Call {
	parsedData, _ := yearnVaultABI.Pack("debtOutstanding0", strategyAddress.Address)
	return ethereum.Call{
		Target:   contractAddress.Address,
		Abi:      yearnVaultABI,
		Method:   `debtOutstanding0`,
		CallData: parsedData,
		Name:     name,
		Version:  version,
	}
}

func getStrategies(name string, contractAddress common.Address, strategyAddress common.Address, version string) ethereum.Call {
	abiToUse := yearnVaultABI

	switch version {
	case `0.2.2`:
		_abi, _ := abi.JSON(strings.NewReader(helpers.YEARN_VAULT_V022_ABI))
		abiToUse = &_abi
	case `0.3.0`, `0.3.1`:
		_abi, _ := abi.JSON(strings.NewReader(helpers.YEARN_VAULT_V030_ABI))
		abiToUse = &_abi
	}

	parsedData, _ := abiToUse.Pack("strategies", strategyAddress.Address)
	return ethereum.Call{
		Name:     name,
		Target:   contractAddress.Address,
		Abi:      abiToUse,
		Method:   `strategies`,
		CallData: parsedData,
		Version:  version,
	}
}

func getExpectedReturn(name string, contractAddress common.Address, strategyAddress common.Address, version string) ethereum.Call {
	parsedData, _ := yearnVaultABI.Pack("expectedReturn0", strategyAddress.Address)
	return ethereum.Call{
		Target:   contractAddress.Address,
		Abi:      yearnVaultABI,
		Method:   `expectedReturn0`,
		CallData: parsedData,
		Name:     name,
		Version:  version,
	}
}

func getStategyEstimatedTotalAsset(name string, contractAddress common.Address, version string) ethereum.Call {
	parsedData, _ := yearnStrategyABI.Pack("estimatedTotalAssets")
	return ethereum.Call{
		Target:   contractAddress.Address,
		Abi:      yearnStrategyABI,
		Method:   `estimatedTotalAssets`,
		CallData: parsedData,
		Name:     name,
		Version:  version,
	}
}

func getStategyIsActive(name string, contractAddress common.Address, version string) ethereum.Call {
	parsedData, _ := yearnStrategyABI.Pack("isActive")
	return ethereum.Call{
		Target:   contractAddress.Address,
		Abi:      yearnStrategyABI,
		Method:   `isActive`,
		CallData: parsedData,
		Name:     name,
		Version:  version,
	}
}

func getStategyKeepCRV(name string, contractAddress common.Address, version string) ethereum.Call {
	parsedData, _ := yearnStrategyABI.Pack("keepCRV")
	return ethereum.Call{
		Target:   contractAddress.Address,
		Abi:      yearnStrategyABI,
		Method:   `keepCRV`,
		CallData: parsedData,
		Name:     name,
		Version:  version,
	}
}

func getStategyDelegatedAssets(name string, contractAddress common.Address, version string) ethereum.Call {
	parsedData, _ := yearnStrategyABI.Pack("delegatedAssets")
	return ethereum.Call{
		Target:   contractAddress.Address,
		Abi:      yearnStrategyABI,
		Method:   `delegatedAssets`,
		CallData: parsedData,
		Name:     name,
		Version:  version,
	}
}

// FetchStrategiesMulticallData will perform a multicall to get some specific data from on-chain for a specific list of strategies
func FetchStrategiesMulticallData(chainID uint64) {
	// First we connect to the multicall client, stored in memory via the initializer.
	caller := ethereum.MulticallClientForChainID[chainID]

	// As the withdrawal queue is missing in the subgraph, we need, for each strategy, to determine that
	withdrawQueueForStrategies := Store.WithdrawalQueueMultiCallData[chainID]

	// Then, for each token listed for our current chainID, we prepare the calls
	var calls = make([]ethereum.Call, 0)
	for _, strat := range Store.StrategyList[chainID] {
		calls = append(calls, getCreditAvailable(strat.Strategy.String(), strat.Vault, strat.Strategy, strat.VaultVersion))
		calls = append(calls, getDebtOutstanding(strat.Strategy.String(), strat.Vault, strat.Strategy, strat.VaultVersion))
		calls = append(calls, getExpectedReturn(strat.Strategy.String(), strat.Vault, strat.Strategy, strat.VaultVersion))
		calls = append(calls, getStrategies(strat.Strategy.String(), strat.Vault, strat.Strategy, strat.VaultVersion))
		calls = append(calls, getStategyEstimatedTotalAsset(strat.Strategy.String(), strat.Strategy, strat.VaultVersion))
		calls = append(calls, getStategyIsActive(strat.Strategy.String(), strat.Strategy, strat.VaultVersion))
		calls = append(calls, getStategyKeepCRV(strat.Strategy.String(), strat.Strategy, strat.VaultVersion))
		calls = append(calls, getStategyDelegatedAssets(strat.Strategy.String(), strat.Strategy, strat.VaultVersion))
	}

	if len(calls) == 0 {
		logs.Error("No strategies found.")
		return
	}

	// Then, we execute the multicall and store the prices in the TokenPrices map
	maxBatch := math.MaxInt64
	response := caller.ExecuteByBatch(calls, maxBatch, nil)
	if Store.StrategyMultiCallData[chainID] == nil {
		Store.StrategyMultiCallData[chainID] = make(map[common.Address]models.TStrategyMultiCallData)
	}
	for _, strat := range Store.StrategyList[chainID] {
		creditAvailable0 := response[strat.Strategy.String()+`creditAvailable0`]
		debtOutstanding0 := response[strat.Strategy.String()+`debtOutstanding0`]
		expectedReturn := response[strat.Strategy.String()+`expectedReturn0`]
		strategies := response[strat.Strategy.String()+`strategies`]
		estimatedTotalAssets := response[strat.Strategy.String()+`estimatedTotalAssets`]
		isActive := response[strat.Strategy.String()+`isActive`]
		keepCRV := response[strat.Strategy.String()+`keepCRV`]
		delegatedAssets := response[strat.Strategy.String()+`delegatedAssets`]
		withdrawalQueuePosition, ok := withdrawQueueForStrategies[strat.Strategy]
		if !ok {
			withdrawalQueuePosition = -1
		}

		data := models.TStrategyMultiCallData{
			CreditAvailable:         bigNumber.NewInt(0),
			DebtOutstanding:         bigNumber.NewInt(0),
			ExpectedReturn:          bigNumber.NewInt(0),
			EstimatedTotalAssets:    bigNumber.NewInt(0),
			KeepCRV:                 bigNumber.NewInt(0),
			DelegatedAssets:         bigNumber.NewInt(0),
			IsActive:                false,
			WithdrawalQueuePosition: bigNumber.NewInt(withdrawalQueuePosition),
		}
		if len(creditAvailable0) == 1 {
			data.CreditAvailable = bigNumber.SetInt(creditAvailable0[0].(*big.Int))
		}
		if len(debtOutstanding0) == 1 {
			data.DebtOutstanding = bigNumber.SetInt(debtOutstanding0[0].(*big.Int))
		}
		if len(expectedReturn) == 1 {
			data.ExpectedReturn = bigNumber.SetInt(expectedReturn[0].(*big.Int))
		}
		if len(estimatedTotalAssets) == 1 {
			data.EstimatedTotalAssets = bigNumber.SetInt(estimatedTotalAssets[0].(*big.Int))
		}
		if len(keepCRV) == 1 {
			data.KeepCRV = bigNumber.SetInt(keepCRV[0].(*big.Int))
		}
		if len(delegatedAssets) == 1 {
			data.DelegatedAssets = bigNumber.SetInt(delegatedAssets[0].(*big.Int))
		}
		if len(isActive) == 1 {
			data.IsActive = isActive[0].(bool)
		}
		if strat.VaultVersion == `0.2.2` && len(strategies) == 8 {
			data.PerformanceFee = bigNumber.SetInt(strategies[0].(*big.Int))
			data.Activation = bigNumber.SetInt(strategies[1].(*big.Int))
			data.DebtLimit = bigNumber.SetInt(strategies[2].(*big.Int))
			data.RateLimit = bigNumber.SetInt(strategies[3].(*big.Int))
			data.LastReport = bigNumber.SetInt(strategies[4].(*big.Int))
			data.TotalDebt = bigNumber.SetInt(strategies[5].(*big.Int))
			data.TotalGain = bigNumber.SetInt(strategies[6].(*big.Int))
			data.TotalLoss = bigNumber.SetInt(strategies[7].(*big.Int))
		} else if (strat.VaultVersion == `0.3.0` || strat.VaultVersion == `0.3.1`) && len(strategies) == 8 {
			data.PerformanceFee = bigNumber.SetInt(strategies[0].(*big.Int))
			data.Activation = bigNumber.SetInt(strategies[1].(*big.Int))
			data.DebtRatio = bigNumber.SetInt(strategies[2].(*big.Int))
			data.RateLimit = bigNumber.SetInt(strategies[3].(*big.Int))
			data.LastReport = bigNumber.SetInt(strategies[4].(*big.Int))
			data.TotalDebt = bigNumber.SetInt(strategies[5].(*big.Int))
			data.TotalGain = bigNumber.SetInt(strategies[6].(*big.Int))
			data.TotalLoss = bigNumber.SetInt(strategies[7].(*big.Int))
		} else if len(strategies) == 9 {
			data.PerformanceFee = bigNumber.SetInt(strategies[0].(*big.Int))
			data.Activation = bigNumber.SetInt(strategies[1].(*big.Int))
			data.DebtRatio = bigNumber.SetInt(strategies[2].(*big.Int))
			data.MinDebtPerHarvest = bigNumber.SetInt(strategies[3].(*big.Int))
			data.MaxDebtPerHarvest = bigNumber.SetInt(strategies[4].(*big.Int))
			data.LastReport = bigNumber.SetInt(strategies[5].(*big.Int))
			data.TotalDebt = bigNumber.SetInt(strategies[6].(*big.Int))
			data.TotalGain = bigNumber.SetInt(strategies[7].(*big.Int))
			data.TotalLoss = bigNumber.SetInt(strategies[8].(*big.Int))
		}
		Store.StrategyMultiCallData[chainID][common.HexToAddress(strat.Strategy.String())] = data
	}
	store.SaveInDBForChainID(store.KEYS.StrategiesMultiCallData, chainID, Store.StrategyMultiCallData[chainID])
}

// LoadStrategyMulticallData will reload the multicall data store from the last state of the local Badger Database
func LoadStrategyMulticallData(chainID uint64, wg *sync.WaitGroup) {
	defer wg.Done()
	temp := make(map[common.Address]models.TStrategyMultiCallData)
	if err := store.LoadFromDBForChainID(store.KEYS.StrategiesMultiCallData, chainID, &temp); err != nil {
		return
	}
	if temp != nil && (len(temp) > 0) {
		Store.StrategyMultiCallData[chainID] = temp
		logs.Success("Data loaded for the strategy multicall data store for chainID: " + strconv.FormatUint(chainID, 10))
	}
}
