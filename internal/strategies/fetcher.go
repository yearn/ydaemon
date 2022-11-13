package strategies

import (
	"context"
	"math"
	"math/big"
	"sync"
	"time"

	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/getsentry/sentry-go"
	"github.com/yearn/ydaemon/common/bigNumber"
	"github.com/yearn/ydaemon/common/ethereum"
	"github.com/yearn/ydaemon/common/logs"
	"github.com/yearn/ydaemon/common/store"
	"github.com/yearn/ydaemon/common/types/common"
	"github.com/yearn/ydaemon/internal/meta"
)

/**************************************************************************************************
** fetchBasicInformations will, for a list of TStrategyAdded, fetch all the relevant basic
** information about the strategy.
**
** Arguments:
** - chainID: the chain ID of the network we are working on
** - strategyAddedList: a list of TStrategyAdded we want to fetch the information for
**
** Returns:
** - a list of TStrategy containing the basic information for the strategies
**************************************************************************************************/
func fetchBasicInformations(
	chainID uint64,
	strategyAddedList []TStrategyAdded,
) (strategyList []*TStrategy) {
	/**********************************************************************************************
	** The first step is to prepare the multicall, connecting to the multicall instance and
	** preparing the array of calls to send. All calls for all vaults will be send in a single
	** multicall and will later be accessible via a concatened string `stratAddress + methodName`.
	**********************************************************************************************/
	caller := ethereum.MulticallClientForChainID[chainID]
	calls := []ethereum.Call{}
	for _, strat := range strategyAddedList {
		vaultAddress := common.FromAddress(strat.VaultAddress)
		stratAddress := common.FromAddress(strat.StrategyAddress)
		calls = append(calls, getCreditAvailable(strat.StrategyAddress.String(), vaultAddress, stratAddress, strat.VaultVersion))
		calls = append(calls, getDebtOutstanding(strat.StrategyAddress.String(), vaultAddress, stratAddress, strat.VaultVersion))
		calls = append(calls, getExpectedReturn(strat.StrategyAddress.String(), vaultAddress, stratAddress, strat.VaultVersion))
		calls = append(calls, getStrategies(strat.StrategyAddress.String(), vaultAddress, stratAddress, strat.VaultVersion))
		calls = append(calls, getStategyEstimatedTotalAsset(strat.StrategyAddress.String(), stratAddress, strat.VaultVersion))
		calls = append(calls, getStategyIsActive(strat.StrategyAddress.String(), stratAddress, strat.VaultVersion))
		calls = append(calls, getStategyKeepCRV(strat.StrategyAddress.String(), stratAddress, strat.VaultVersion))
		calls = append(calls, getStategyDelegatedAssets(strat.StrategyAddress.String(), stratAddress, strat.VaultVersion))
		calls = append(calls, getName(strat.StrategyAddress.String(), stratAddress, strat.VaultVersion))
	}

	/**********************************************************************************************
	** Regular fix for Fantom's RPC, which limit the number of calls in a multicall to a very low
	** number. We split the multicall in multiple calls of 3 calls each.
	** Otherwise, we just send the multicall as is.
	**********************************************************************************************/
	maxBatch := math.MaxInt64
	if chainID == 250 {
		maxBatch = 3
	}

	/**********************************************************************************************
	** Then we can proceed the responses. Some date will already be available from the list of
	** tokens, so we can already play with that.
	**********************************************************************************************/
	response := caller.ExecuteByBatch(calls, maxBatch, nil)
	for _, strat := range strategyAddedList {
		vaultAddress := strat.VaultAddress
		stratAddress := strat.StrategyAddress
		creditAvailable0 := response[stratAddress.String()+`creditAvailable0`]
		debtOutstanding0 := response[stratAddress.String()+`debtOutstanding0`]
		expectedReturn := response[stratAddress.String()+`expectedReturn0`]
		strategies := response[stratAddress.String()+`strategies`]
		estimatedTotalAssets := response[stratAddress.String()+`estimatedTotalAssets`]
		isActive := response[stratAddress.String()+`isActive`]
		keepCRV := response[stratAddress.String()+`keepCRV`]
		delegatedAssets := response[stratAddress.String()+`delegatedAssets`]
		name := response[stratAddress.String()+`name`]

		withdrawalQueuePosition := int64(-1)
		if withdrawalQueue, ok := FindWithdrawalQueueForVault(chainID, vaultAddress); ok {
			for i, stratInQueue := range withdrawalQueue {
				if stratInQueue.Hex() == stratAddress.Hex() {
					withdrawalQueuePosition = int64(i)
					break
				}
			}
		}

		newStrategy := &TStrategy{
			Address:                 strat.StrategyAddress,
			VaultAddress:            strat.VaultAddress,
			VaultVersion:            strat.VaultVersion,
			CreditAvailable:         bigNumber.NewInt(0),
			DebtOutstanding:         bigNumber.NewInt(0),
			ExpectedReturn:          bigNumber.NewInt(0),
			EstimatedTotalAssets:    bigNumber.NewInt(0),
			KeepCRV:                 bigNumber.NewInt(0),
			DelegatedAssets:         bigNumber.NewInt(0),
			IsActive:                false,
			WithdrawalQueuePosition: bigNumber.NewInt(withdrawalQueuePosition),
			Initialization: TStrategyInitialization{
				TxHash:      strat.TxHash,
				BlockNumber: strat.BlockNumber,
				TxIndex:     strat.TxIndex,
				LogIndex:    strat.LogIndex,
			},
		}
		if len(creditAvailable0) == 1 {
			newStrategy.CreditAvailable = bigNumber.SetInt(creditAvailable0[0].(*big.Int))
		}
		if len(debtOutstanding0) == 1 {
			newStrategy.DebtOutstanding = bigNumber.SetInt(debtOutstanding0[0].(*big.Int))
		}
		if len(expectedReturn) == 1 {
			newStrategy.ExpectedReturn = bigNumber.SetInt(expectedReturn[0].(*big.Int))
		}
		if len(estimatedTotalAssets) == 1 {
			newStrategy.EstimatedTotalAssets = bigNumber.SetInt(estimatedTotalAssets[0].(*big.Int))
		}
		if len(keepCRV) == 1 {
			newStrategy.KeepCRV = bigNumber.SetInt(keepCRV[0].(*big.Int))
		}
		if len(delegatedAssets) == 1 {
			newStrategy.DelegatedAssets = bigNumber.SetInt(delegatedAssets[0].(*big.Int))
		}
		if len(isActive) == 1 {
			newStrategy.IsActive = isActive[0].(bool)
		}
		if len(name) == 1 {
			newStrategy.Name = name[0].(string)
		}
		if strat.VaultVersion == `0.2.2` && len(strategies) == 8 {
			newStrategy.PerformanceFee = bigNumber.SetInt(strategies[0].(*big.Int))
			newStrategy.Activation = bigNumber.SetInt(strategies[1].(*big.Int))
			newStrategy.DebtLimit = bigNumber.SetInt(strategies[2].(*big.Int))
			newStrategy.RateLimit = bigNumber.SetInt(strategies[3].(*big.Int))
			newStrategy.LastReport = bigNumber.SetInt(strategies[4].(*big.Int))
			newStrategy.TotalDebt = bigNumber.SetInt(strategies[5].(*big.Int))
			newStrategy.TotalGain = bigNumber.SetInt(strategies[6].(*big.Int))
			newStrategy.TotalLoss = bigNumber.SetInt(strategies[7].(*big.Int))
		} else if (strat.VaultVersion == `0.3.0` || strat.VaultVersion == `0.3.1`) && len(strategies) == 8 {
			newStrategy.PerformanceFee = bigNumber.SetInt(strategies[0].(*big.Int))
			newStrategy.Activation = bigNumber.SetInt(strategies[1].(*big.Int))
			newStrategy.DebtRatio = bigNumber.SetInt(strategies[2].(*big.Int))
			newStrategy.RateLimit = bigNumber.SetInt(strategies[3].(*big.Int))
			newStrategy.LastReport = bigNumber.SetInt(strategies[4].(*big.Int))
			newStrategy.TotalDebt = bigNumber.SetInt(strategies[5].(*big.Int))
			newStrategy.TotalGain = bigNumber.SetInt(strategies[6].(*big.Int))
			newStrategy.TotalLoss = bigNumber.SetInt(strategies[7].(*big.Int))
		} else if len(strategies) == 9 {
			newStrategy.PerformanceFee = bigNumber.SetInt(strategies[0].(*big.Int))
			newStrategy.Activation = bigNumber.SetInt(strategies[1].(*big.Int))
			newStrategy.DebtRatio = bigNumber.SetInt(strategies[2].(*big.Int))
			newStrategy.MinDebtPerHarvest = bigNumber.SetInt(strategies[3].(*big.Int))
			newStrategy.MaxDebtPerHarvest = bigNumber.SetInt(strategies[4].(*big.Int))
			newStrategy.LastReport = bigNumber.SetInt(strategies[5].(*big.Int))
			newStrategy.TotalDebt = bigNumber.SetInt(strategies[6].(*big.Int))
			newStrategy.TotalGain = bigNumber.SetInt(strategies[7].(*big.Int))
			newStrategy.TotalLoss = bigNumber.SetInt(strategies[8].(*big.Int))
		}

		if strategyFromMeta, ok := meta.GetMetaStrategy(chainID, common.FromAddress(stratAddress)); ok {
			newStrategy.Name = strategyFromMeta.Name
			newStrategy.Description = strategyFromMeta.Description
		}

		strategyList = append(strategyList, newStrategy)
	}

	return strategyList
}

/**************************************************************************************************
** findAllStrategies is simply a wrapper around our fetchBasicInformations function to make it
** easier to read. It will take the strategyAddedList and call the fetchBasicInformations function
** to get the strategies information, before assigning them to a new map.
**
** Arguments:
** - chainID: the chain ID of the network we are working on
** - strategyAddedList: a list of TStrategyAdded we want to fetch the information for
**
** Returns:
** - a map of stratAddress -> TStrategy
**************************************************************************************************/
func findAllStrategies(
	chainID uint64,
	strategyAddedList []TStrategyAdded,
) map[ethcommon.Address]*TStrategy {
	newMap := make(map[ethcommon.Address]*TStrategy)
	newStrategyList := fetchBasicInformations(chainID, strategyAddedList)
	for _, strat := range newStrategyList {
		newMap[strat.Address] = strat
	}

	return newMap
}

/**************************************************************************************************
** The base of Yearn are the vaults. They are the smart contracts that are used to manage the
** deposits and the withdrawals of the users.
** The goal of this function is to get a list of all the Vaults living in the Yearn's Ecosystem.
**
** Arguments:
** - chainID: the chain ID of the network we are working on
**
** Returns:
** - a map of vaultAddress -> TVault
**************************************************************************************************/
func RetrieveAllStrategies(
	chainID uint64,
	strategyAddedList []TStrategyAdded,
) map[ethcommon.Address]*TStrategy {
	span := sentry.StartSpan(context.Background(), "app.fetch",
		sentry.TransactionName("Fetch Strategies"))
	span.SetTag("subsystem", "daemon")
	defer span.Finish()

	timeBefore := time.Now()

	/**********************************************************************************************
	** First, try to retrieve the list of strategies from the database to exclude the one existing
	** from the upcoming calls
	**********************************************************************************************/
	strategyMap := make(map[ethcommon.Address]*TStrategy)
	store.Iterate(chainID, store.TABLES.STRATEGIES, &strategyMap)

	/**********************************************************************************************
	** Once we got out basic list, we will fetch theses basics informations.
	**********************************************************************************************/
	if len(strategyAddedList) > 0 {
		updatedStrategiesMap := findAllStrategies(chainID, strategyAddedList)

		/**********************************************************************************************
		** Once everything is setup, we will store each token in the DB. The storage is set as a map
		** of strategyAddress -> TStrategy. All the strategies will be retrievable from the
		** store.Interate() func.
		**********************************************************************************************/
		wg := sync.WaitGroup{}
		wg.Add(len(updatedStrategiesMap))
		for _, token := range updatedStrategiesMap {
			go func(_token *TStrategy) {
				defer wg.Done()
				store.SaveInDB(
					chainID,
					store.TABLES.STRATEGIES,
					_token.Address.String(),
					_token,
				)
			}(token)
		}
		wg.Wait()
		store.Iterate(chainID, store.TABLES.STRATEGIES, &strategyMap)
	}

	logs.Success(`It tooks`, time.Since(timeBefore), `to retrieve`, len(strategyMap), `strategies`)
	_strategyMap[chainID] = strategyMap
	return strategyMap
}
