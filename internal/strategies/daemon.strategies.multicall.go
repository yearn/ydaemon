package strategies

import (
	"math"
	"math/big"
	"strconv"
	"strings"
	"sync"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/internal/utils/contracts"
	"github.com/yearn/ydaemon/internal/utils/ethereum"
	"github.com/yearn/ydaemon/internal/utils/helpers"
	"github.com/yearn/ydaemon/internal/utils/logs"
	"github.com/yearn/ydaemon/internal/utils/models"
	"github.com/yearn/ydaemon/internal/utils/store"
)

// yearnVaultABI takes the ABI of the standard Yearn Vault contract and prepare it for the multicall
var yearnVaultABI, _ = contracts.YearnVaultMetaData.GetAbi()
var yearnStrategyABI, _ = contracts.StrategyBaseMetaData.GetAbi()

func getCreditAvailable(name string, contractAddress common.Address, strategyAddress common.Address, version string) ethereum.Call {
	parsedData, _ := yearnVaultABI.Pack("creditAvailable0", strategyAddress)
	return ethereum.Call{
		Target:   contractAddress,
		Abi:      yearnVaultABI,
		Method:   `creditAvailable0`,
		CallData: parsedData,
		Name:     name,
		Version:  version,
	}
}

func getDebtOutstanding(name string, contractAddress common.Address, strategyAddress common.Address, version string) ethereum.Call {
	parsedData, _ := yearnVaultABI.Pack("debtOutstanding0", strategyAddress)
	return ethereum.Call{
		Target:   contractAddress,
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

	parsedData, _ := abiToUse.Pack("strategies", strategyAddress)
	return ethereum.Call{
		Name:     name,
		Target:   contractAddress,
		Abi:      abiToUse,
		Method:   `strategies`,
		CallData: parsedData,
		Version:  version,
	}
}

func getExpectedReturn(name string, contractAddress common.Address, strategyAddress common.Address, version string) ethereum.Call {
	parsedData, _ := yearnVaultABI.Pack("expectedReturn0", strategyAddress)
	return ethereum.Call{
		Target:   contractAddress,
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
		Target:   contractAddress,
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
		Target:   contractAddress,
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
		Target:   contractAddress,
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
		Target:   contractAddress,
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
	response := caller.ExecuteByBatch(calls, maxBatch)
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

		data := models.TStrategyMultiCallData{
			CreditAvailable:      big.NewInt(0),
			DebtOutstanding:      big.NewInt(0),
			ExpectedReturn:       big.NewInt(0),
			EstimatedTotalAssets: big.NewInt(0),
			KeepCRV:              big.NewInt(0),
			DelegatedAssets:      big.NewInt(0),
			IsActive:             false,
		}
		if len(creditAvailable0) == 1 {
			data.CreditAvailable = creditAvailable0[0].(*big.Int)
		}
		if len(debtOutstanding0) == 1 {
			data.DebtOutstanding = debtOutstanding0[0].(*big.Int)
		}
		if len(expectedReturn) == 1 {
			data.ExpectedReturn = expectedReturn[0].(*big.Int)
		}
		if len(estimatedTotalAssets) == 1 {
			data.EstimatedTotalAssets = estimatedTotalAssets[0].(*big.Int)
		}
		if len(keepCRV) == 1 {
			data.KeepCRV = keepCRV[0].(*big.Int)
		}
		if len(delegatedAssets) == 1 {
			data.DelegatedAssets = delegatedAssets[0].(*big.Int)
		}
		if len(isActive) == 1 {
			data.IsActive = isActive[0].(bool)
		}
		if strat.VaultVersion == `0.2.2` && len(strategies) == 8 {
			data.PerformanceFee = strategies[0].(*big.Int)
			data.Activation = strategies[1].(*big.Int)
			data.DebtLimit = strategies[2].(*big.Int)
			data.RateLimit = strategies[3].(*big.Int)
			data.LastReport = strategies[4].(*big.Int)
			data.TotalDebt = strategies[5].(*big.Int)
			data.TotalGain = strategies[6].(*big.Int)
			data.TotalLoss = strategies[7].(*big.Int)
		} else if (strat.VaultVersion == `0.3.0` || strat.VaultVersion == `0.3.1`) && len(strategies) == 8 {
			data.PerformanceFee = strategies[0].(*big.Int)
			data.Activation = strategies[1].(*big.Int)
			data.DebtRatio = strategies[2].(*big.Int)
			data.RateLimit = strategies[3].(*big.Int)
			data.LastReport = strategies[4].(*big.Int)
			data.TotalDebt = strategies[5].(*big.Int)
			data.TotalGain = strategies[6].(*big.Int)
			data.TotalLoss = strategies[7].(*big.Int)
		} else if len(strategies) == 9 {
			data.PerformanceFee = strategies[0].(*big.Int)
			data.Activation = strategies[1].(*big.Int)
			data.DebtRatio = strategies[2].(*big.Int)
			data.MinDebtPerHarvest = strategies[3].(*big.Int)
			data.MaxDebtPerHarvest = strategies[4].(*big.Int)
			data.LastReport = strategies[5].(*big.Int)
			data.TotalDebt = strategies[6].(*big.Int)
			data.TotalGain = strategies[7].(*big.Int)
			data.TotalLoss = strategies[8].(*big.Int)
		}
		Store.StrategyMultiCallData[chainID][common.HexToAddress(strat.Strategy.String())] = data
	}
	store.SaveInDBForChainID(`StrategiesMultiCallData`, chainID, Store.StrategyMultiCallData[chainID])
}

// LoadStrategyMulticallData will reload the multicall data store from the last state of the local Badger Database
func LoadStrategyMulticallData(chainID uint64, wg *sync.WaitGroup) {
	defer wg.Done()
	temp := make(map[common.Address]models.TStrategyMultiCallData)
	if err := store.LoadFromDBForChainID(`StrategiesMultiCallData`, chainID, &temp); err != nil {
		return
	}
	if temp != nil && (len(temp) > 0) {
		Store.StrategyMultiCallData[chainID] = temp
		logs.Success("Data loaded for the strategy multicall data store for chainID: " + strconv.FormatUint(chainID, 10))
	}
}
