package strategies

import (
	"math"
	"math/big"
	"strconv"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/common/bigNumber"
	"github.com/yearn/ydaemon/common/ethereum"
	"github.com/yearn/ydaemon/common/helpers"
	"github.com/yearn/ydaemon/common/store"
	"github.com/yearn/ydaemon/common/traces"
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
		calls = append(calls, getCreditAvailable(strat.StrategyAddress.String(), strat.VaultAddress, strat.StrategyAddress, strat.VaultVersion))
		calls = append(calls, getDebtOutstanding(strat.StrategyAddress.String(), strat.VaultAddress, strat.StrategyAddress, strat.VaultVersion))
		calls = append(calls, getExpectedReturn(strat.StrategyAddress.String(), strat.VaultAddress, strat.StrategyAddress, strat.VaultVersion))
		calls = append(calls, getStrategies(strat.StrategyAddress.String(), strat.VaultAddress, strat.StrategyAddress, strat.VaultVersion))
		calls = append(calls, getStategyEstimatedTotalAsset(strat.StrategyAddress.String(), strat.StrategyAddress, strat.VaultVersion))
		calls = append(calls, getStategyIsActive(strat.StrategyAddress.String(), strat.StrategyAddress, strat.VaultVersion))
		calls = append(calls, getStategyKeepCRV(strat.StrategyAddress.String(), strat.StrategyAddress, strat.VaultVersion))
		calls = append(calls, getStategyDelegatedAssets(strat.StrategyAddress.String(), strat.StrategyAddress, strat.VaultVersion))
		calls = append(calls, getName(strat.StrategyAddress.String(), strat.StrategyAddress, strat.VaultVersion))
		calls = append(calls, getKeeper(strat.StrategyAddress.String(), strat.StrategyAddress, strat.VaultVersion))
		calls = append(calls, getStrategist(strat.StrategyAddress.String(), strat.StrategyAddress, strat.VaultVersion))
		calls = append(calls, getRewards(strat.StrategyAddress.String(), strat.StrategyAddress, strat.VaultVersion))
		calls = append(calls, getHealthCheck(strat.StrategyAddress.String(), strat.StrategyAddress, strat.VaultVersion))
		calls = append(calls, getAPIVersion(strat.StrategyAddress.String(), strat.StrategyAddress, strat.VaultVersion))
		calls = append(calls, getDoHealthCheck(strat.StrategyAddress.String(), strat.StrategyAddress, strat.VaultVersion))
		calls = append(calls, getEmergencyExit(strat.StrategyAddress.String(), strat.StrategyAddress, strat.VaultVersion))
	}

	/**********************************************************************************************
	** Regular fix for Fantom's RPC, which limit the number of calls in a multicall to a very low
	** number. We split the multicall in multiple calls of 3 calls each.
	** Otherwise, we just send the multicall as is.
	**********************************************************************************************/
	maxBatch := math.MaxInt64

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
		keeper := response[stratAddress.String()+`keeper`]
		strategist := response[stratAddress.String()+`strategist`]
		rewards := response[stratAddress.String()+`rewards`]
		healthCheck := response[stratAddress.String()+`healthCheck`]
		apiVersion := response[stratAddress.String()+`apiVersion`]
		doHealthCheck := response[stratAddress.String()+`doHealthCheck`]
		emergencyExit := response[stratAddress.String()+`emergencyExit`]

		withdrawalQueuePosition := int64(-1)
		isInQueue := false
		if withdrawalQueue, ok := FindWithdrawalQueueForVault(chainID, vaultAddress); ok {
			for i, stratInQueue := range withdrawalQueue {
				if stratInQueue.Hex() == stratAddress.Hex() {
					withdrawalQueuePosition = int64(i)
					isInQueue = true
					break
				}
			}
		}

		newStrategy := &TStrategy{
			Address:                 strat.StrategyAddress,
			VaultAddress:            strat.VaultAddress,
			ChainID:                 chainID,
			VaultVersion:            strat.VaultVersion,
			CreditAvailable:         bigNumber.NewInt(0),
			DebtOutstanding:         bigNumber.NewInt(0),
			ExpectedReturn:          bigNumber.NewInt(0),
			EstimatedTotalAssets:    bigNumber.NewInt(0),
			KeepCRV:                 bigNumber.NewInt(0),
			DelegatedAssets:         bigNumber.NewInt(0),
			IsActive:                false,
			IsInQueue:               isInQueue,
			WithdrawalQueuePosition: bigNumber.NewInt(withdrawalQueuePosition),
			Initialization: TStrategyInitialization{
				TxHash:      strat.TxHash,
				BlockNumber: strat.BlockNumber,
				TxIndex:     strat.TxIndex,
				LogIndex:    strat.LogIndex,
			},
		}
		newStrategy.CreditAvailable = helpers.DecodeBigInt(creditAvailable0)
		newStrategy.DebtOutstanding = helpers.DecodeBigInt(debtOutstanding0)
		newStrategy.ExpectedReturn = helpers.DecodeBigInt(expectedReturn)
		newStrategy.EstimatedTotalAssets = helpers.DecodeBigInt(estimatedTotalAssets)
		newStrategy.KeepCRV = helpers.DecodeBigInt(keepCRV)
		newStrategy.DelegatedAssets = helpers.DecodeBigInt(delegatedAssets)
		newStrategy.IsActive = helpers.DecodeBool(isActive)
		newStrategy.Name = helpers.DecodeString(name)
		newStrategy.KeeperAddress = helpers.DecodeAddress(keeper)
		newStrategy.StrategistAddress = helpers.DecodeAddress(strategist)
		newStrategy.RewardsAddress = helpers.DecodeAddress(rewards)
		newStrategy.HealthCheckAddress = helpers.DecodeAddress(healthCheck)
		newStrategy.APIVersion = helpers.DecodeString(apiVersion)
		newStrategy.DoHealthCheck = helpers.DecodeBool(doHealthCheck)
		newStrategy.EmergencyExit = helpers.DecodeBool(emergencyExit)

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

		if strategyFromMeta, ok := meta.GetMetaStrategy(chainID, stratAddress); ok {
			newStrategy.DisplayName = strategyFromMeta.Name
			newStrategy.Description = strategyFromMeta.Description
			newStrategy.Protocols = strategyFromMeta.Protocols
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
) map[common.Address]*TStrategy {
	newMap := make(map[common.Address]*TStrategy)
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
) map[common.Address]*TStrategy {
	trace := traces.Init(`app.indexer.strategies.multicall_data`).
		SetTag(`chainID`, strconv.FormatUint(chainID, 10)).
		SetTag(`rpcURI`, ethereum.GetRPCURI(chainID)).
		SetTag(`entity`, `strategies`).
		SetTag(`subsystem`, `daemon`)
	defer trace.Finish()

	/**********************************************************************************************
	** First, try to retrieve the list of strategies from the database to exclude the one existing
	** from the upcoming calls
	**********************************************************************************************/
	strategyMap := make(map[common.Address]*TStrategy)
	store.ListFromBadgerDB(chainID, store.TABLES.STRATEGIES, &strategyMap)

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
			go func(_strategy *TStrategy) {
				defer wg.Done()
				store.SaveInBadgerDB(
					chainID,
					store.TABLES.STRATEGIES,
					_strategy.Address.String(),
					_strategy,
				)
			}(token)
		}
		wg.Wait()
		store.ListFromBadgerDB(chainID, store.TABLES.STRATEGIES, &strategyMap)
	}

	_strategyMap[chainID] = strategyMap
	return strategyMap
}
