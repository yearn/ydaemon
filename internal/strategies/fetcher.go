package strategies

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/common/bigNumber"
	"github.com/yearn/ydaemon/common/ethereum"
	"github.com/yearn/ydaemon/common/helpers"
	"github.com/yearn/ydaemon/internal/meta"
	"github.com/yearn/ydaemon/internal/models"
	"github.com/yearn/ydaemon/internal/multicalls"
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
	strategyAddedList []models.TStrategyAdded,
) (strategyList []*models.TStrategy) {
	/**********************************************************************************************
	** The first step is to prepare the multicall, connecting to the multicall instance and
	** preparing the array of calls to send. All calls for all vaults will be send in a single
	** multicall and will later be accessible via a concatened string `stratAddress + methodName`.
	**********************************************************************************************/
	calls := []ethereum.Call{}
	for _, strat := range strategyAddedList {
		calls = append(calls, multicalls.GetCreditAvailable(strat.Address.Hex(), strat.VaultAddress, strat.Address, strat.VaultVersion))
		calls = append(calls, multicalls.GetDebtOutstanding(strat.Address.Hex(), strat.VaultAddress, strat.Address, strat.VaultVersion))
		calls = append(calls, multicalls.GetExpectedReturn(strat.Address.Hex(), strat.VaultAddress, strat.Address, strat.VaultVersion))
		calls = append(calls, multicalls.GetStrategies(strat.Address.Hex(), strat.VaultAddress, strat.Address, strat.VaultVersion))
		calls = append(calls, multicalls.GetStategyEstimatedTotalAsset(strat.Address.Hex(), strat.Address, strat.VaultVersion))
		calls = append(calls, multicalls.GetStategyIsActive(strat.Address.Hex(), strat.Address, strat.VaultVersion))
		calls = append(calls, multicalls.GetStategyKeepCRV(strat.Address.Hex(), strat.Address, strat.VaultVersion))
		calls = append(calls, multicalls.GetStategyKeepCRVPercent(strat.Address.Hex(), strat.Address, strat.VaultVersion))
		calls = append(calls, multicalls.GetStategyKeepCVX(strat.Address.Hex(), strat.Address, strat.VaultVersion))
		calls = append(calls, multicalls.GetStategyDelegatedAssets(strat.Address.Hex(), strat.Address, strat.VaultVersion))
		calls = append(calls, multicalls.GetStrategyName(strat.Address.Hex(), strat.Address, strat.VaultVersion))
		calls = append(calls, multicalls.GetKeeper(strat.Address.Hex(), strat.Address, strat.VaultVersion))
		calls = append(calls, multicalls.GetStrategist(strat.Address.Hex(), strat.Address, strat.VaultVersion))
		calls = append(calls, multicalls.GetStrategyRewards(strat.Address.Hex(), strat.Address, strat.VaultVersion))
		calls = append(calls, multicalls.GetHealthCheck(strat.Address.Hex(), strat.Address, strat.VaultVersion))
		calls = append(calls, multicalls.GetStrategyAPIVersion(strat.Address.Hex(), strat.Address, strat.VaultVersion))
		calls = append(calls, multicalls.GetDoHealthCheck(strat.Address.Hex(), strat.Address, strat.VaultVersion))
		calls = append(calls, multicalls.GetEmergencyExit(strat.Address.Hex(), strat.Address, strat.VaultVersion))
	}

	/**********************************************************************************************
	** Then we can proceed the responses. Some date will already be available from the list of
	** tokens, so we can already play with that.
	**********************************************************************************************/
	response := multicalls.Perform(chainID, calls, nil)
	for _, strat := range strategyAddedList {
		vaultAddress := strat.VaultAddress
		stratAddress := strat.Address
		creditAvailable0 := response[stratAddress.Hex()+`creditAvailable0`]
		debtOutstanding0 := response[stratAddress.Hex()+`debtOutstanding0`]
		expectedReturn := response[stratAddress.Hex()+`expectedReturn0`]
		strategies := response[stratAddress.Hex()+`strategies`]
		estimatedTotalAssets := response[stratAddress.Hex()+`estimatedTotalAssets`]
		isActive := response[stratAddress.Hex()+`isActive`]
		keepCRV := response[stratAddress.Hex()+`keepCRV`]
		keepCRVPercent := response[stratAddress.Hex()+`keepCrvPercent`]
		keepCVX := response[stratAddress.Hex()+`keepCVX`]
		delegatedAssets := response[stratAddress.Hex()+`delegatedAssets`]
		name := response[stratAddress.Hex()+`name`]
		keeper := response[stratAddress.Hex()+`keeper`]
		strategist := response[stratAddress.Hex()+`strategist`]
		rewards := response[stratAddress.Hex()+`rewards`]
		healthCheck := response[stratAddress.Hex()+`healthCheck`]
		apiVersion := response[stratAddress.Hex()+`apiVersion`]
		doHealthCheck := response[stratAddress.Hex()+`doHealthCheck`]
		emergencyExit := response[stratAddress.Hex()+`emergencyExit`]

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

		newStrategy := &models.TStrategy{
			Address:                 strat.Address,
			VaultAddress:            strat.VaultAddress,
			ChainID:                 chainID,
			VaultVersion:            strat.VaultVersion,
			CreditAvailable:         bigNumber.NewInt(0),
			DebtOutstanding:         bigNumber.NewInt(0),
			ExpectedReturn:          bigNumber.NewInt(0),
			EstimatedTotalAssets:    bigNumber.NewInt(0),
			KeepCRV:                 bigNumber.NewInt(0),
			KeepCRVPercent:          bigNumber.NewInt(0),
			KeepCVX:                 bigNumber.NewInt(0),
			DelegatedAssets:         bigNumber.NewInt(0),
			IsActive:                false,
			IsInQueue:               isInQueue,
			WithdrawalQueuePosition: bigNumber.NewInt(withdrawalQueuePosition),
		}
		newStrategy.CreditAvailable = helpers.DecodeBigInt(creditAvailable0)
		newStrategy.DebtOutstanding = helpers.DecodeBigInt(debtOutstanding0)
		newStrategy.ExpectedReturn = helpers.DecodeBigInt(expectedReturn)
		newStrategy.EstimatedTotalAssets = helpers.DecodeBigInt(estimatedTotalAssets)
		newStrategy.KeepCRV = helpers.DecodeBigInt(keepCRV)
		newStrategy.KeepCRVPercent = helpers.DecodeBigInt(keepCRVPercent)
		newStrategy.KeepCVX = helpers.DecodeBigInt(keepCVX)
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
	strategyAddedList []models.TStrategyAdded,
) map[common.Address]*models.TStrategy {
	newMap := make(map[common.Address]*models.TStrategy)
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
	strategyAddedList []models.TStrategyAdded,
) map[common.Address]*models.TStrategy {
	updatedStrategiesMap := findAllStrategies(chainID, strategyAddedList)

	StoreStrategies(chainID, updatedStrategiesMap)
	return updatedStrategiesMap
}
