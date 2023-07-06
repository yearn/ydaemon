package strategies

import (
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/common/env"
	"github.com/yearn/ydaemon/internal/models"
)

/**********************************************************************************************
** Set of functions to store and retrieve the strategies from the cache and/or database and
** being able to access them from the rest of the application.
** The _strategyMap variable is not exported and is only used internally by the functions below.
**********************************************************************************************/
var _strategyMap = make(map[uint64]*sync.Map)

func initOrGetStrategyMap(chainID uint64) *sync.Map {
	syncMap := _strategyMap[chainID]
	if syncMap == nil {
		syncMap = &sync.Map{}
		_strategyMap[chainID] = syncMap
	}
	return syncMap
}

var _strategyWithdrawalQueueMap = make(map[uint64]*sync.Map)

func initOrGetStrategyWithdrawalQueueMap(chainID uint64) *sync.Map {
	syncMap := _strategyWithdrawalQueueMap[chainID]
	if syncMap == nil {
		syncMap = &sync.Map{}
		_strategyWithdrawalQueueMap[chainID] = syncMap
	}
	return syncMap
}

func init() {
	for _, chainID := range env.SUPPORTED_CHAIN_IDS {
		if _, ok := _strategyMap[chainID]; !ok {
			_strategyMap[chainID] = &sync.Map{}
			_strategyWithdrawalQueueMap[chainID] = &sync.Map{}
		}
	}
}

/**********************************************************************************************
** ListStrategies will, for a given chainID, return the list of all the strategies stored in
** the _strategyMap.
**********************************************************************************************/
func ListStrategies(chainID uint64) []*models.TStrategy {
	var strategies []*models.TStrategy
	syncMap := initOrGetStrategyMap(chainID)
	syncMap.Range(func(_, value interface{}) bool {
		strategies = append(strategies, value.(*models.TStrategy))
		return true
	})
	return strategies
}

/**********************************************************************************************
** ListStrategiesForVault will, for a given chainID and a given vault address, return the list
** of all the strategies stored in the _strategyMap.
**********************************************************************************************/
func ListStrategiesForVault(chainID uint64, vaultAddress common.Address) []*models.TStrategy {
	var strategies []*models.TStrategy
	syncMap := initOrGetStrategyMap(chainID)
	syncMap.Range(func(_, value interface{}) bool {
		strategy := value.(*models.TStrategy)
		if strategy.VaultAddress.Hex() == vaultAddress.Hex() {
			strategies = append(strategies, value.(*models.TStrategy))
		}
		return true
	})
	return strategies
}

/**********************************************************************************************
** ListStrategiesAddresses will, for a given chainID, return the list of addresses of all the
** strategies stored in _strategyMap.
**********************************************************************************************/
func ListStrategiesAddresses(chainID uint64) []common.Address {
	var addresses []common.Address
	syncMap := initOrGetStrategyMap(chainID)
	syncMap.Range(func(key, _ interface{}) bool {
		addresses = append(addresses, common.HexToAddress(key.(string)))
		return true
	})
	return addresses
}

/**********************************************************************************************
** FindStrategy will, for a given chainID, try to find the provided strategyAddress stored in
** the _strategyMap. It will return the strategy if found, and a boolean indicating if the
** strategy was found or not.
**********************************************************************************************/
func FindStrategy(chainID uint64, strategyAddress common.Address) (*models.TStrategy, bool) {
	syncMap := initOrGetStrategyMap(chainID)
	strategy, ok := syncMap.Load(strategyAddress.Hex())
	if !ok {
		return nil, false
	}
	return strategy.(*models.TStrategy), true
}

func StoreStrategies(chainID uint64, strategies map[common.Address]*models.TStrategy) {
	syncMap := initOrGetStrategyMap(chainID)
	for addr, strategy := range strategies {
		syncMap.Store(addr.Hex(), strategy)
	}
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
	syncMap := initOrGetStrategyWithdrawalQueueMap(chainID)
	syncMap.Store(vaultAddress.Hex(), queue)
}

/**********************************************************************************************
** FindWithdrawalQueueForVault will retrieve the withdrawal queue for a given vault address.
**********************************************************************************************/
func FindWithdrawalQueueForVault(chainID uint64, vaultAddress common.Address) ([]common.Address, bool) {
	syncMap := initOrGetStrategyWithdrawalQueueMap(chainID)
	queue, ok := syncMap.Load(vaultAddress.Hex())
	if !ok {
		return nil, false
	}
	return queue.([]common.Address), true
}
