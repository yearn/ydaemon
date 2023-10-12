package fetcher

import (
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/common/addresses"
	"github.com/yearn/ydaemon/common/bigNumber"
	"github.com/yearn/ydaemon/common/ethereum"
	"github.com/yearn/ydaemon/common/helpers"
	"github.com/yearn/ydaemon/internal/models"
	"github.com/yearn/ydaemon/internal/multicalls"
	"github.com/yearn/ydaemon/internal/storage"
)

/**************************************************************************************************
** fetchStrategiesBasicInformations will, for a list of strategies, fetch all the relevant basic
** information about the strategy.
**
** Arguments:
** - chainID: the chain ID of the network we are working on
** - strategies: a list of strategies we want to fetch the information for
**
** Returns:
** - a list of TStrategy containing the basic information for the strategies
**************************************************************************************************/
func fetchStrategiesBasicInformations(chainID uint64, strategiesMap map[common.Address]models.TStrategy) map[common.Address]models.TStrategy {
	/**********************************************************************************************
	** The first step is to prepare the multicall, connecting to the multicall instance and
	** preparing the array of calls to send. All calls for all vaults will be send in a single
	** multicall and will later be accessible via a concatened string `stratAddress + methodName`.
	**********************************************************************************************/
	calls := []ethereum.Call{}
	for _, strat := range strategiesMap {
		if vault, ok := storage.GetVault(chainID, strat.VaultAddress); ok && !strat.ShouldRefresh {
			if vault.IsRetired {
				continue
			}
		}

		if time.Since(strat.LastUpdate).Hours() > 24 || strat.ShouldRefresh {
			// If the last vault update was more than 24 hour ago, we will do a full update
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
			calls = append(calls, multicalls.GetDoHealthCheck(strat.Address.Hex(), strat.Address, strat.VaultVersion))
			calls = append(calls, multicalls.GetEmergencyExit(strat.Address.Hex(), strat.Address, strat.VaultVersion))
		} else if time.Since(strat.LastUpdate).Hours() > 1 {
			// If the last vault update was more than 1 hour ago, we will do a partial update
			calls = append(calls, multicalls.GetStategyKeepCRV(strat.Address.Hex(), strat.Address, strat.VaultVersion))
			calls = append(calls, multicalls.GetStategyKeepCRVPercent(strat.Address.Hex(), strat.Address, strat.VaultVersion))
			calls = append(calls, multicalls.GetStategyKeepCVX(strat.Address.Hex(), strat.Address, strat.VaultVersion))
			calls = append(calls, multicalls.GetEmergencyExit(strat.Address.Hex(), strat.Address, strat.VaultVersion))
			calls = append(calls, multicalls.GetCreditAvailable(strat.Address.Hex(), strat.VaultAddress, strat.Address, strat.VaultVersion))
			calls = append(calls, multicalls.GetDebtOutstanding(strat.Address.Hex(), strat.VaultAddress, strat.Address, strat.VaultVersion))
			calls = append(calls, multicalls.GetExpectedReturn(strat.Address.Hex(), strat.VaultAddress, strat.Address, strat.VaultVersion))
			calls = append(calls, multicalls.GetStrategies(strat.Address.Hex(), strat.VaultAddress, strat.Address, strat.VaultVersion))
			calls = append(calls, multicalls.GetStategyEstimatedTotalAsset(strat.Address.Hex(), strat.Address, strat.VaultVersion))
			calls = append(calls, multicalls.GetStategyDelegatedAssets(strat.Address.Hex(), strat.Address, strat.VaultVersion))
		} else {
			// Else, for every loop, do at least this
			calls = append(calls, multicalls.GetCreditAvailable(strat.Address.Hex(), strat.VaultAddress, strat.Address, strat.VaultVersion))
			calls = append(calls, multicalls.GetDebtOutstanding(strat.Address.Hex(), strat.VaultAddress, strat.Address, strat.VaultVersion))
			calls = append(calls, multicalls.GetExpectedReturn(strat.Address.Hex(), strat.VaultAddress, strat.Address, strat.VaultVersion))
			calls = append(calls, multicalls.GetStrategies(strat.Address.Hex(), strat.VaultAddress, strat.Address, strat.VaultVersion))
			calls = append(calls, multicalls.GetStategyEstimatedTotalAsset(strat.Address.Hex(), strat.Address, strat.VaultVersion))
			calls = append(calls, multicalls.GetStategyDelegatedAssets(strat.Address.Hex(), strat.Address, strat.VaultVersion))
		}
	}

	/**********************************************************************************************
	** Then we can proceed the responses. Some date will already be available from the list of
	** tokens, so we can already play with that.
	**********************************************************************************************/
	response := multicalls.Perform(chainID, calls, nil)
	for _, strat := range strategiesMap {
		vault, ok := storage.GetVault(chainID, strat.VaultAddress)
		if ok && !strat.ShouldRefresh {
			if vault.IsRetired {
				continue
			}
		}

		stratAddress := strat.Address
		rawCreditAvailable0 := response[stratAddress.Hex()+`creditAvailable0`]
		rawDebtOutstanding0 := response[stratAddress.Hex()+`debtOutstanding0`]
		rawExpectedReturn := response[stratAddress.Hex()+`expectedReturn0`]
		rawStrategies := response[stratAddress.Hex()+`strategies`]
		rawEstimatedTotalAssets := response[stratAddress.Hex()+`estimatedTotalAssets`]
		rawIsActive := response[stratAddress.Hex()+`isActive`]
		rawKeepCRV := response[stratAddress.Hex()+`keepCRV`]
		rawKeepCRVPercent := response[stratAddress.Hex()+`keepCrvPercent`]
		rawKeepCVX := response[stratAddress.Hex()+`keepCVX`]
		rawDelegatedAssets := response[stratAddress.Hex()+`delegatedAssets`]
		rawName := response[stratAddress.Hex()+`name`]
		rawKeeper := response[stratAddress.Hex()+`keeper`]
		rawStrategist := response[stratAddress.Hex()+`strategist`]
		rawRewards := response[stratAddress.Hex()+`rewards`]
		rawHealthCheck := response[stratAddress.Hex()+`healthCheck`]
		rawDoHealthCheck := response[stratAddress.Hex()+`doHealthCheck`]
		rawEmergencyExit := response[stratAddress.Hex()+`emergencyExit`]

		withdrawalQueuePosition := int64(-1)
		isInQueue := false
		for i, stratInQueue := range vault.LastActiveStrategies {
			if stratInQueue.Hex() == stratAddress.Hex() {
				withdrawalQueuePosition = int64(i)
				isInQueue = true
				break
			}
		}

		/******************************************************************************************
		** Preparing our new TStrategy object
		******************************************************************************************/
		newStrategy := strat
		newStrategy.LastUpdate = time.Now()
		newStrategy.IsInQueue = isInQueue
		newStrategy.ShouldRefresh = false
		newStrategy.WithdrawalQueuePosition = bigNumber.NewInt(withdrawalQueuePosition)
		newStrategy.LastCreditAvailable = helpers.DecodeBigInt(rawCreditAvailable0)
		newStrategy.LastDebtOutstanding = helpers.DecodeBigInt(rawDebtOutstanding0)
		newStrategy.LastExpectedReturn = helpers.DecodeBigInt(rawExpectedReturn)
		newStrategy.LastEstimatedTotalAssets = helpers.DecodeBigInt(rawEstimatedTotalAssets)
		newStrategy.LastDelegatedAssets = helpers.DecodeBigInt(rawDelegatedAssets)
		if strat.VaultVersion == `0.2.2` && len(rawStrategies) == 8 {
			newStrategy.LastReport = bigNumber.SetInt(rawStrategies[4].(*big.Int))
			newStrategy.LastTotalDebt = bigNumber.SetInt(rawStrategies[5].(*big.Int))
			newStrategy.LastTotalGain = bigNumber.SetInt(rawStrategies[6].(*big.Int))
			newStrategy.LastTotalLoss = bigNumber.SetInt(rawStrategies[7].(*big.Int))
			newStrategy.LastPerformanceFee = bigNumber.SetInt(rawStrategies[0].(*big.Int))
			newStrategy.LastDebtLimit = bigNumber.SetInt(rawStrategies[2].(*big.Int))
			newStrategy.LastRateLimit = bigNumber.SetInt(rawStrategies[3].(*big.Int))
			newStrategy.TimeActivated = bigNumber.SetInt(rawStrategies[1].(*big.Int))
		} else if (strat.VaultVersion == `0.3.0` || strat.VaultVersion == `0.3.1`) && len(rawStrategies) == 8 {
			newStrategy.LastReport = bigNumber.SetInt(rawStrategies[4].(*big.Int))
			newStrategy.LastTotalDebt = bigNumber.SetInt(rawStrategies[5].(*big.Int))
			newStrategy.LastTotalGain = bigNumber.SetInt(rawStrategies[6].(*big.Int))
			newStrategy.LastTotalLoss = bigNumber.SetInt(rawStrategies[7].(*big.Int))
			newStrategy.LastPerformanceFee = bigNumber.SetInt(rawStrategies[0].(*big.Int))
			newStrategy.LastDebtRatio = bigNumber.SetInt(rawStrategies[2].(*big.Int))
			newStrategy.LastRateLimit = bigNumber.SetInt(rawStrategies[3].(*big.Int))
			newStrategy.TimeActivated = bigNumber.SetInt(rawStrategies[1].(*big.Int))
		} else if len(rawStrategies) == 9 {
			newStrategy.LastReport = bigNumber.SetInt(rawStrategies[5].(*big.Int))
			newStrategy.LastTotalDebt = bigNumber.SetInt(rawStrategies[6].(*big.Int))
			newStrategy.LastTotalGain = bigNumber.SetInt(rawStrategies[7].(*big.Int))
			newStrategy.LastTotalLoss = bigNumber.SetInt(rawStrategies[8].(*big.Int))
			newStrategy.LastPerformanceFee = bigNumber.SetInt(rawStrategies[0].(*big.Int))
			newStrategy.LastDebtRatio = bigNumber.SetInt(rawStrategies[2].(*big.Int))
			newStrategy.LastMinDebtPerHarvest = bigNumber.SetInt(rawStrategies[3].(*big.Int))
			newStrategy.LastMaxDebtPerHarvest = bigNumber.SetInt(rawStrategies[4].(*big.Int))
			newStrategy.TimeActivated = bigNumber.SetInt(rawStrategies[1].(*big.Int))
		}

		if len(rawKeepCRV) > 0 {
			newStrategy.KeepCRV = helpers.DecodeBigInt(rawKeepCRV)
		}
		if len(rawKeepCRVPercent) > 0 {
			newStrategy.KeepCRVPercent = helpers.DecodeBigInt(rawKeepCRVPercent)
		}
		if len(rawKeepCVX) > 0 {
			newStrategy.KeepCVX = helpers.DecodeBigInt(rawKeepCVX)
		}
		if len(rawName) > 0 {
			newStrategy.Name = helpers.DecodeString(rawName)
		}
		if len(rawKeeper) > 0 {
			newStrategy.KeeperAddress = helpers.DecodeAddress(rawKeeper)
		}
		if len(rawStrategist) > 0 {
			newStrategy.StrategistAddress = helpers.DecodeAddress(rawStrategist)
		}
		if len(rawRewards) > 0 {
			newStrategy.RewardsAddress = helpers.DecodeAddress(rawRewards)
		}
		if len(rawHealthCheck) > 0 {
			newStrategy.HealthCheckAddress = helpers.DecodeAddress(rawHealthCheck)
		}
		if len(rawIsActive) > 0 {
			newStrategy.IsActive = helpers.DecodeBool(rawIsActive)
		}
		if len(rawDoHealthCheck) > 0 {
			newStrategy.DoHealthCheck = helpers.DecodeBool(rawDoHealthCheck)
		}
		if len(rawEmergencyExit) > 0 {
			newStrategy.EmergencyExit = helpers.DecodeBool(rawEmergencyExit)
		}
		if !newStrategy.IsActive {
			newStrategy.IsRetired = true
		}

		if chainID == 1 && addresses.Equals(newStrategy.Address, `0xE7863292dd8eE5d215eC6D75ac00911D06E59B2d`) {
			styCRVVault, ok := storage.GetVault(chainID, newStrategy.VaultAddress)
			if ok {
				newStrategy.LastDebtRatio = bigNumber.NewUint64(10000)
				newStrategy.LastTotalDebt = styCRVVault.LastTotalAssets
			}
		}

		strategiesMap[stratAddress] = newStrategy
		storage.StoreStrategy(chainID, newStrategy)
	}

	return strategiesMap
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
	strategies map[common.Address]models.TStrategy,
) map[common.Address]models.TStrategy {
	fetchStrategiesBasicInformations(chainID, strategies)

	strategyMap, _ := storage.ListStrategies(chainID)
	storage.StoreStrategiesToJson(chainID, strategyMap)
	return strategyMap
}
