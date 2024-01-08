package store

import (
	"sync"

	"github.com/yearn/ydaemon/common/env"
	"github.com/yearn/ydaemon/common/logs"
	"github.com/yearn/ydaemon/internal/storage"
)

type TDBType string

const (
	DBSql TDBType = "sql"
)

var _dbType TDBType

/**************************************************************************************************
** The init function is a special function triggered directly on execution of the package.
** It is used to initialize the package.
** This init is responsible of loading the store
***************************************************************************************************/
func init() {
	if shouldUseMySQLDB := initializeMySQLDatabase(); shouldUseMySQLDB {
		_dbType = DBSql
	}

	for chainID := range env.CHAINS {
		_blockTimeSyncMap[chainID] = &sync.Map{}
		_timeBlockSyncMap[chainID] = &sync.Map{}
		_vaultsPricePerShareSyncMap[chainID] = &sync.Map{}
	}

	wg := &sync.WaitGroup{}
	for chainID := range env.CHAINS {
		storage.LoadRegistries(chainID, nil)
		storage.LoadVaults(chainID, nil)
		storage.LoadStrategies(chainID, nil)
		storage.LoadERC20(chainID, nil)

		wg.Add(2)
		go LoadBlockTime(chainID, wg)
		go LoadPricePerShare(chainID, wg)
	}
	wg.Wait()
	logs.Success(`Initialized the store`)
}
