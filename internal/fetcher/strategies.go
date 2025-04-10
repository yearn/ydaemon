package fetcher

import (
	"errors"
	"math"
	"strings"
	"sync"
	"time"

	"github.com/yearn/ydaemon/common/addresses"
	"github.com/yearn/ydaemon/common/bigNumber"
	"github.com/yearn/ydaemon/common/ethereum"
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

/**************************************************************************************************
** trimIgnoredStrategies filters the strategies map to include only those that should be processed.
**
** This function optimizes performance by filtering out strategies that don't need to be processed:
** 1. Keeps strategies explicitly marked for refresh regardless of other conditions
** 2. Filters out strategies whose vaults cannot be found in storage
** 3. Filters out strategies from retired vaults (with exceptions for specific Alchemix-related vaults)
**
** This pre-filtering step significantly reduces the number of multicalls and processing needed
** for strategies that don't require updates.
**
** @param strategiesMap map[string]models.TStrategy - The map of strategies to filter
** @return map[string]models.TStrategy - The filtered map containing only relevant strategies
**************************************************************************************************/
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

/**************************************************************************************************
** getStrategyAPR calculates and returns the appropriate APR for a strategy based on its version.
**
** This function determines and retrieves the Annual Percentage Rate (APR) for a strategy, with
** behavior that varies depending on the vault version. For v3 vaults, it first attempts to use
** forward-looking APR, then falls back to current APR if that fails. For earlier versions, it
** always uses current APR.
**
** The function includes performance optimizations to avoid unnecessary calculations:
** - Skips APR calculation for inactive strategies (not active, not in queue, or retired)
** - Skips APR calculation when there's no financial activity (zero debt, gain, and loss)
**
** @param chainID uint64 - The blockchain network ID
** @param versionMajor string - Major version of the vault ("3" for v3, etc.)
** @param strategy models.TStrategy - The strategy to calculate APR for
** @return *bigNumber.Float - The calculated APR value
** @return models.TStrategyAPRType - The type of APR (forward or current)
** @return error - Any error encountered during calculation
**************************************************************************************************/
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

	/******************************************************************************************
	** For v3 vaults: First try to get forward-looking APR, fall back to current APR if that fails
	** For earlier versions: Only try to get current APR
	******************************************************************************************/
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

/**************************************************************************************************
** assignStrategy processes a strategy, updates its properties, and calculates its APR.
**
** This function handles the full lifecycle of strategy data processing:
** 1. Checks cache to avoid redundant processing
** 2. Updates strategy properties based on vault information
** 3. Applies version-specific processing (V2 vs V3)
** 4. Handles special cases for specific strategies
** 5. Sets the strategy status based on its activity and allocation state
** 6. Calculates and assigns APR values
** 7. Updates the cache and persistent storage
**
** The strategy status is categorized as:
** - active: The default state for active strategies with allocated funds
** - not_active: For retired strategies or those explicitly marked as inactive
** - unallocated: For strategies that are active but have no funds allocated (LastTotalDebt = 0)
**
** @param chainID uint64 - The blockchain network ID
** @param strategy models.TStrategy - The strategy to process
** @param response map[string][]interface{} - Multicall responses containing strategy data
** @return string - The strategy key (composite of address and vault address)
** @return models.TStrategy - The updated strategy object
** @return bool - Whether a cached version was used (true) or a fresh update was performed (false)
**************************************************************************************************/
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

	/******************************************************************************************
	** Initialize strategy properties and sync with vault information
	******************************************************************************************/
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

	/******************************************************************************************
	** Apply version-specific processing to extract strategy metrics
	******************************************************************************************/
	versionMajor := strings.Split(vault.Version, `.`)[0]
	if versionMajor == `3` {
		newStrategy = handleV3StrategyCalls(newStrategy, response)
	} else {
		newStrategy = handleV2StrategyCalls(newStrategy, response)
	}

	/******************************************************************************************
	** Handle some specific cases for particular strategies
	******************************************************************************************/
	if chainID == 1 && addresses.Equals(newStrategy.Address, `0xE7863292dd8eE5d215eC6D75ac00911D06E59B2d`) {
		styCRVVault, ok := storage.GetVault(chainID, newStrategy.VaultAddress)
		if ok {
			newStrategy.LastDebtRatio = bigNumber.NewUint64(10000)
			newStrategy.LastTotalDebt = styCRVVault.LastTotalAssets
		}
	}

	/******************************************************************************************
	** Set the strategy status based on its properties
	** - active: The default state for active strategies with allocated funds
	** - not_active: When it's retired or IsActive is false
	** - unallocated: When it's active but LastTotalDebt is 0
	******************************************************************************************/
	if newStrategy.IsRetired || !newStrategy.IsActive {
		newStrategy.Status = models.StrategyStatusNotActive
	} else if newStrategy.LastTotalDebt.IsZero() {
		newStrategy.Status = models.StrategyStatusUnallocated
	} else {
		newStrategy.Status = models.StrategyStatusActive
	}

	/******************************************************************************************
	** Calculate and assign APR values
	******************************************************************************************/
	netAPR, aprType, err := getStrategyAPR(chainID, versionMajor, newStrategy)
	if err == nil {
		netAPRFloat, _ := netAPR.Float64()
		if math.IsInf(netAPRFloat, 0) {
			netAPRFloat = math.MaxFloat64
		}
		newStrategy.NetAPR = netAPRFloat
		newStrategy.APRType = aprType
	} else {
		// logs.Error(`Error while computing APR for ` + newStrategy.Address.Hex() + ` | ` + newStrategy.VaultAddress.Hex() + `: ` + err.Error())
	}

	/******************************************************************************************
	** Update cache and persistent storage
	******************************************************************************************/
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
