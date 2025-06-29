package storage

import (
	"github.com/yearn/ydaemon/common/env"
	"github.com/yearn/ydaemon/common/logs"
)

/**************************************************************************************************
** The init function is a special function triggered directly on execution of the package.
** It is used to initialize the package.
** This init is responsible of loading the store
***************************************************************************************************/
func InitializeStorage() {
	for chainID := range env.GetChains() {
		LoadRegistries(chainID, nil)
		LoadVaults(chainID, nil)
		LoadStrategies(chainID, nil)
		LoadERC20(chainID, nil)
		LoadAPY(chainID, nil)
	}
	logs.Success(`Initialized the store`)
}
