package store

import (
	"sync"

	"github.com/yearn/ydaemon/common/env"
	"github.com/yearn/ydaemon/internal/storage"
)

type TDBType string

const (
	DBBadger TDBType = "badger"
	DBSql    TDBType = "sql"
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
	} else {
		_dbType = DBBadger
	}

	for chainID := range env.CHAINS {
		_blockTimeSyncMap[chainID] = &sync.Map{}
		_timeBlockSyncMap[chainID] = &sync.Map{}
		_historicalPriceSyncMap[chainID] = &sync.Map{}
		_vaultsPricePerShareSyncMap[chainID] = &sync.Map{}
		_stackingPoolsSyncMap[chainID] = &sync.Map{}
	}

	wg := &sync.WaitGroup{}
	for chainID := range env.CHAINS {
		storage.LoadERC20(chainID, nil) //This is a blocking function, required for the next function to work

		go LoadBlockTime(chainID, nil) //This does not require a wg
		// go LoadHistoricalPrice(chainID, nil)

		wg.Add(4)
		go storage.LoadRegistries(chainID, wg)
		go storage.LoadVaults(chainID, wg)
		go storage.LoadStrategies(chainID, wg)
		go LoadPricePerShare(chainID, wg)
	}
	wg.Wait()
}
