package store

import (
	"sync"

	"github.com/yearn/ydaemon/common/env"
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
		_newVaultsFromRegistrySyncMap[chainID] = &sync.Map{}
		_vaultsSyncMap[chainID] = &sync.Map{}
		_erc20SyncMap[chainID] = &sync.Map{}
		_strategiesSyncMap[chainID] = &sync.Map{}
		_vaultsPricePerShareSyncMap[chainID] = &sync.Map{}
		_vaultsActivations[chainID] = &sync.Map{}
		_stackingPoolsSyncMap[chainID] = &sync.Map{}
	}

	wg := &sync.WaitGroup{}
	for chainID := range env.CHAINS {
		wg.Add(5)
		go LoadBlockTime(chainID, nil)
		// go LoadHistoricalPrice(chainID, nil)
		go LoadNewVaultsFromRegistry(chainID, wg)
		go LoadStrategies(chainID, wg)
		LoadERC20(chainID, nil) //This is a blocking function, required for the next function to work
		go LoadVaults(chainID, wg)
		go LoadPricePerShare(chainID, wg)
		go LoadVaultsActivation(chainID, wg)
	}
	wg.Wait()
}
