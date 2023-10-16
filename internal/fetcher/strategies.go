package fetcher

import (
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/common/addresses"
	"github.com/yearn/ydaemon/common/bigNumber"
	"github.com/yearn/ydaemon/common/ethereum"
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
		vault, ok := storage.GetVault(chainID, strat.VaultAddress)
		if ok && !strat.ShouldRefresh && vault.IsRetired {
			continue
		}

		if vault.Version == `3.0.0` {
			calls = append(calls, getV3StrategyCalls(strat)...)
		} else {
			calls = append(calls, getV2StrategyCalls(strat)...)
		}
	}

	/**********************************************************************************************
	** Then we can proceed the responses. Some date will already be available from the list of
	** tokens, so we can already play with that.
	**********************************************************************************************/
	response := multicalls.Perform(chainID, calls, nil)
	for _, strat := range strategiesMap {
		vault, ok := storage.GetVault(chainID, strat.VaultAddress)
		if ok && !strat.ShouldRefresh && vault.IsRetired {
			continue
		}

		newStrategy := strat
		newStrategy.ChainID = chainID
		newStrategy.LastUpdate = time.Now()
		newStrategy.ShouldRefresh = false
		for _, stratInQueue := range vault.LastActiveStrategies {
			if stratInQueue.Hex() == strat.Address.Hex() {
				newStrategy.IsInQueue = true
				break
			}
		}

		if vault.Version == `3.0.0` {
			newStrategy = handleV3StrategyCalls(strat, response)
		} else {
			newStrategy = handleV2StrategyCalls(strat, response)
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

		strategiesMap[strat.Address] = newStrategy
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
