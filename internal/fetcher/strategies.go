package fetcher

import (
	"errors"
	"strings"
	"sync"
	"time"

	"github.com/yearn/ydaemon/common/addresses"
	"github.com/yearn/ydaemon/common/bigNumber"
	"github.com/yearn/ydaemon/common/ethereum"
	"github.com/yearn/ydaemon/common/logs"
	"github.com/yearn/ydaemon/internal/models"
	"github.com/yearn/ydaemon/internal/multicalls"
	"github.com/yearn/ydaemon/internal/storage"
	"github.com/yearn/ydaemon/processes/apr"
)

/**************************************************************************************************
** Cache store for strategy fetch times and results to prevent duplicate fetches within a 5-minute window
**************************************************************************************************/
type strategyCache struct {
	lastFetch time.Time
	strategy  models.TStrategy
}

var (
	strategyCacheStore sync.Map
	strategyCacheMutex sync.RWMutex
	cacheDuration      = 5 * time.Minute
)

func trimIgnoredStrategies(strategiesMap map[string]models.TStrategy) map[string]models.TStrategy {
	relevantStrategies := make(map[string]models.TStrategy)
	for key, strat := range strategiesMap {
		if strat.ShouldRefresh {
			relevantStrategies[key] = strat
			continue
		}

		vault, ok := storage.GetVault(strat.ChainID, strat.VaultAddress)
		if !ok {
			continue
		}

		// Adding an exception for the vault that is retired but we still want to compute. Alchemix related
		isException := addresses.Equals(vault.Address, "0xaD17A225074191d5c8a37B50FdA1AE278a2EE6A2") ||
			addresses.Equals(vault.Address, "0x5B977577Eb8a480f63e11FC615D6753adB8652Ae") ||
			addresses.Equals(vault.Address, "0x65343F414FFD6c97b0f6add33d16F6845Ac22BAc") ||
			addresses.Equals(vault.Address, "0xFaee21D0f0Af88EE72BB6d68E54a90E6EC2616de")
		if vault.Metadata.IsRetired && !isException {
			continue
		}
		relevantStrategies[key] = strat
	}
	return relevantStrategies
}

func getStrategyAPR(chainID uint64, versionMajor string, strategy models.TStrategy) (*bigNumber.Float, models.TStrategyAPRType, error) {
	netAPR := bigNumber.NewFloat(0)
	aprType := models.APRTypeCurrent

	/******************************************************************************************
	** Skip APR calculation for inactive strategies to optimize performance
	** If the strategy is not active, not in queue, or retired, we return 0 APR
	** Also if all totalLoss, totalGain, and totalDebt are 0, no need to calculate APR
	******************************************************************************************/
	if !strategy.IsActive || !strategy.IsInQueue || strategy.IsRetired {
		return netAPR, aprType, nil
	}

	if strategy.LastTotalLoss.IsZero() && strategy.LastTotalGain.IsZero() && strategy.LastTotalDebt.IsZero() {
		return netAPR, aprType, nil
	}

	if versionMajor == `3` {
		if forwardAPR, err := apr.ComputeForwardStrategyAPR(strategy); err == nil {
			netAPR = forwardAPR
			aprType = models.APRTypeForward
		} else if currentAPR, err := apr.GetCurrentStrategyAPR(chainID, strategy.Address.Hex()); err == nil {
			netAPR = currentAPR
			aprType = models.APRTypeCurrent
		} else {
			return netAPR, aprType, errors.New(`Error while computing forward APR for ` + strategy.Address.Hex() + ` | ` + strategy.VaultAddress.Hex() + `: ` + err.Error())
		}
	} else {
		if currentAPR, err := apr.GetCurrentStrategyAPR(chainID, strategy.Address.Hex()); err == nil {
			netAPR = currentAPR
			aprType = models.APRTypeCurrent
		} else {
			return netAPR, aprType, errors.New(`Error while computing current APR for ` + strategy.Address.Hex() + ` | ` + strategy.VaultAddress.Hex() + `: ` + err.Error())
		}
	}
	return netAPR, aprType, nil
}

func assignStrategy(chainID uint64, strategy models.TStrategy, response map[string][]interface{}) (string, models.TStrategy, bool) {
	strategyCacheMutex.RLock()
	defer strategyCacheMutex.RUnlock()

	strategyKey := strategy.Address.Hex() + `_` + strategy.VaultAddress.Hex()
	now := time.Now()

	// Use Load instead of direct map access
	if cacheInterface, exists := strategyCacheStore.Load(strategyKey); exists {
		cache := cacheInterface.(strategyCache)
		if !strategy.ShouldRefresh && now.Sub(cache.lastFetch) < cacheDuration {
			return strategyKey, cache.strategy, true
		}

	}

	vault, _ := storage.GetVault(chainID, strategy.VaultAddress) // Checked before with trimIgnoredStrategies
	newStrategy := strategy
	newStrategy.ChainID = chainID
	newStrategy.ShouldRefresh = false
	for _, stratInQueue := range vault.LastActiveStrategies {
		if stratInQueue.Hex() == strategy.Address.Hex() {
			newStrategy.IsInQueue = true
			break
		}
	}

	versionMajor := strings.Split(vault.Version, `.`)[0]
	if versionMajor == `3` {
		newStrategy = handleV3StrategyCalls(newStrategy, response)
	} else {
		newStrategy = handleV2StrategyCalls(newStrategy, response)
	}

	/******************************************************************************************
	** Handle some specific cases
	******************************************************************************************/
	if chainID == 1 && addresses.Equals(newStrategy.Address, `0xE7863292dd8eE5d215eC6D75ac00911D06E59B2d`) {
		styCRVVault, ok := storage.GetVault(chainID, newStrategy.VaultAddress)
		if ok {
			newStrategy.LastDebtRatio = bigNumber.NewUint64(10000)
			newStrategy.LastTotalDebt = styCRVVault.LastTotalAssets
		}
	}

	netAPR, aprType, err := getStrategyAPR(chainID, versionMajor, newStrategy)
	if err == nil {
		newStrategy.NetAPR = netAPR
		newStrategy.APRType = aprType
	} else {
		logs.Pretty(newStrategy)
		logs.Error(`Error while computing APR for ` + newStrategy.Address.Hex() + ` | ` + newStrategy.VaultAddress.Hex() + `: ` + err.Error())
	}

	// Use Store instead of direct map assignment
	strategyCacheStore.Store(strategyKey, strategyCache{now, newStrategy})
	storage.StoreStrategy(chainID, newStrategy)
	return strategyKey, newStrategy, false
}

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
func fetchStrategiesBasicInformations(
	chainID uint64,
	strategiesMap map[string]models.TStrategy,
) map[string]models.TStrategy {
	relevantStrategies := trimIgnoredStrategies(strategiesMap)
	updatedStrategiesMap := make(map[string]models.TStrategy)

	/**********************************************************************************************
	** The first step is to prepare the multicall, connecting to the multicall instance and
	** preparing the array of calls to send. All calls for all vaults will be send in a single
	** multicall and will later be accessible via a concatened string `stratAddress + methodName`.
	**********************************************************************************************/
	calls := []ethereum.Call{}
	now := time.Now()

	for _, strat := range relevantStrategies {
		strategyKey := strat.Address.Hex() + `_` + strat.VaultAddress.Hex()

		// Check cache with read lock
		strategyCacheMutex.RLock()
		if cacheInterface, exists := strategyCacheStore.Load(strategyKey); exists {
			cache := cacheInterface.(strategyCache)
			if !strat.ShouldRefresh && now.Sub(cache.lastFetch) < cacheDuration {
				updatedStrategiesMap[strategyKey] = cache.strategy
				strategyCacheMutex.RUnlock()
				continue
			}
		}

		vault, _ := storage.GetVault(strat.ChainID, strat.VaultAddress) // Checked before with trimIgnoredStrategies
		versionMajor := strings.Split(vault.Version, `.`)[0]
		if versionMajor == `3` {
			calls = append(calls, getV3StrategyCalls(strat)...)
		} else {
			calls = append(calls, getV2StrategyCalls(strat)...)
		}
		strategyCacheMutex.RUnlock()
	}

	/**********************************************************************************************
	** Then we can proceed the responses. Some date will already be available from the list of
	** tokens, so we can already play with that.
	**********************************************************************************************/
	response := multicalls.Perform(chainID, calls, nil)
	for _, strat := range relevantStrategies {
		strategyKey, updatedStrategy, _ := assignStrategy(chainID, strat, response)
		updatedStrategiesMap[strategyKey] = updatedStrategy
	}

	// allstrat, _ := storage.ListStrategies(chainID)
	// //Check apy is 0
	// apy0Count := 0
	// non0Count := 0
	// nilCount := 0
	// if !alreadyRun {
	// 	for _, strat := range allstrat {
	// 		vault, _ := storage.GetVault(chainID, strat.VaultAddress) // Checked before with trimIgnoredStrategies
	// 		versionMajor := strings.Split(vault.Version, `.`)[0]
	// 		if versionMajor == `3` {
	// 			if strat.ForwardAPY == nil {
	// 				nilCount++
	// 				logs.Info(`ForwardAPY is nil for ` + strat.Address.Hex() + ` | ` + strat.VaultAddress.Hex())
	// 			} else if strat.ForwardAPY.IsZero() {
	// 				apy0Count++
	// 				logs.Info(`ForwardAPY is 0 for ` + strat.Address.Hex() + ` | ` + strat.VaultAddress.Hex())
	// 			} else {
	// 				non0Count++
	// 			}
	// 		}
	// 	}
	// 	alreadyRun = true
	// }
	// logs.Info(`Strategy fetch summary - Total: ` + strconv.Itoa(totalCount) + `, Fetched: ` + strconv.Itoa(fetchedCount) + `, Cached: ` + strconv.Itoa(cachedCount) + `, APY0: ` + strconv.Itoa(apy0Count) + `, Non0: ` + strconv.Itoa(non0Count) + `, Nil: ` + strconv.Itoa(nilCount) + `, ForwardRunCount: ` + strconv.Itoa(forwardRunCount))
	return updatedStrategiesMap
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
	strategies map[string]models.TStrategy,
) map[string]models.TStrategy {
	fetchStrategiesBasicInformations(chainID, strategies)

	strategyMap, _ := storage.ListStrategies(chainID)
	storage.StoreStrategiesToJson(chainID, strategyMap)
	return strategyMap
}
