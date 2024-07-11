package indexer

import (
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/common/contracts"
	"github.com/yearn/ydaemon/common/ethereum"
	"github.com/yearn/ydaemon/common/logs"
	"github.com/yearn/ydaemon/internal/models"
)

var _alreadyIndexedPoolTogetherVaults = make(map[uint64]*sync.Map)

/** 🔵 - Yearn *************************************************************************************
** The function `IndexNewStrategies` is responsible for indexing new strategies for a given chain ID.
**************************************************************************************************/
func IndexYearnXPoolTogetherVaults(chainID uint64) (historicalStrategiesMap map[string]models.TStrategy) {
	if _, ok := _alreadyIndexedPoolTogetherVaults[chainID]; !ok {
		_alreadyIndexedPoolTogetherVaults[chainID] = &sync.Map{}
	}

	if chainID != 10 {
		return historicalStrategiesMap
	}

	/** 🔵 - Yearn *********************************************************************************
	** Loop over all the known vaults for the specified chain ID.
	**********************************************************************************************/
	client := ethereum.GetRPC(chainID)
	registry, err := contracts.NewPoolTogetherRegistry(
		common.HexToAddress(`0x0c379e9b71ba7079084ada0d1c1aeb85d24dfd39`),
		client,
	)
	if err != nil {
		return historicalStrategiesMap
	}
	totalVaults, err := registry.TotalVaults(nil)
	if err != nil {
		return historicalStrategiesMap
	}
	logs.Pretty(`Total vaults: ` + totalVaults.String())

	return historicalStrategiesMap
}
