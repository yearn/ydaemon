package strategies

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/internal/models"
)

/**********************************************************************************************
** Set of functions to store and retrieve the strategies from the cache and/or database and
** being able to access them from the rest of the application.
** The _strategyMap variable is not exported and is only used internally by the functions below.
**********************************************************************************************/
var _strategyMap = make(map[uint64]map[common.Address]*models.TStrategy)
var _strategyWithdrawalQueueMap = make(map[uint64]map[common.Address][]common.Address)

/**********************************************************************************************
** ListStrategies will, for a given chainID, return the list of all the strategies stored in
** the _strategyMap.
**********************************************************************************************/
func ListStrategies(chainID uint64) []*models.TStrategy {
	var strategies []*models.TStrategy
	for _, strategy := range _strategyMap[chainID] {
		strategies = append(strategies, strategy)
	}
	return strategies
}

/**********************************************************************************************
** ListStrategiesForVault will, for a given chainID and a given vault address, return the list
** of all the strategies stored in the _strategyMap.
**********************************************************************************************/
func ListStrategiesForVault(chainID uint64, vaultAddress common.Address) []*models.TStrategy {
	var strategies []*models.TStrategy
	for _, strategy := range _strategyMap[chainID] {
		if strategy.VaultAddress.Hex() == vaultAddress.Hex() {
			strategies = append(strategies, strategy)
		}
	}
	return strategies
}

/**********************************************************************************************
** ListStrategiesAddresses will, for a given chainID, return the list of addresses of all the
** strategies stored in _strategyMap.
**********************************************************************************************/
func ListStrategiesAddresses(chainID uint64) []common.Address {
	var addresses []common.Address
	for address := range _strategyMap[chainID] {
		addresses = append(addresses, address)
	}
	return addresses
}

/**********************************************************************************************
** FindStrategy will, for a given chainID, try to find the provided strategyAddress stored in
** the _strategyMap. It will return the strategy if found, and a boolean indicating if the
** strategy was found or not.
**********************************************************************************************/
func FindStrategy(chainID uint64, strategyAddress common.Address) (*models.TStrategy, bool) {
	strategy, ok := _strategyMap[chainID][strategyAddress]
	if !ok {
		return nil, false
	}
	return strategy, true
}

/**********************************************************************************************
** ToVaultMap will transform a list of TStrategyAdded into a map of
** TStrategyAdded per vault address:
** - [vaultAddress] - [strategyAddress] - TStrategyAdded
**********************************************************************************************/
func ToVaultMap(strategiesAddedList []*models.TStrategy) map[common.Address]map[common.Address]*models.TStrategy {
	strategiesAddedPerVault := make(map[common.Address]map[common.Address]*models.TStrategy)
	for _, strategy := range strategiesAddedList {
		if _, ok := strategiesAddedPerVault[strategy.VaultAddress]; !ok {
			strategiesAddedPerVault[strategy.VaultAddress] = make(map[common.Address]*models.TStrategy)
		}
		strategiesAddedPerVault[strategy.VaultAddress][strategy.Address] = strategy
	}
	return strategiesAddedPerVault
}

/**********************************************************************************************
** SetInStrategiesWithdrawalQueue will mirror a vault withdrawal queue in the strategies
** package to avoid import circles.
**********************************************************************************************/
func SetInStrategiesWithdrawalQueue(chainID uint64, vaultAddress common.Address, queue []common.Address) {
	if _, ok := _strategyWithdrawalQueueMap[chainID]; !ok {
		_strategyWithdrawalQueueMap[chainID] = make(map[common.Address][]common.Address)
	}
	_strategyWithdrawalQueueMap[chainID][vaultAddress] = queue
}

/**********************************************************************************************
** FindWithdrawalQueueForVault will retrieve the withdrawal queue for a given vault address.
**********************************************************************************************/
func FindWithdrawalQueueForVault(chainID uint64, vaultAddress common.Address) ([]common.Address, bool) {
	if _, ok := _strategyWithdrawalQueueMap[chainID]; !ok {
		return nil, false
	}
	return _strategyWithdrawalQueueMap[chainID][vaultAddress], true
}
