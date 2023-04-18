package store

import (
	"sync"

	"github.com/yearn/ydaemon/common/env"
)

type TDBType string

const (
	DBBadger TDBType = "badger"
	DBMysql  TDBType = "mysql"
)

var _dbType TDBType
var _blockTimeSyncMap = make(map[uint64]*sync.Map)
var _historicalPriceSyncMap = make(map[uint64]*sync.Map)
var _newVaultsFromRegistrySyncMap = make(map[uint64]*sync.Map)
var _vaultsSyncMap = make(map[uint64]*sync.Map)

/**************************************************************************************************
** The init function is a special function triggered directly on execution of the package.
** It is used to initialize the package.
** This init is responsible of loading the store
***************************************************************************************************/
func init() {
	if shouldUseMySQLDB := initializeMySQLDatabase(); shouldUseMySQLDB {
		_dbType = DBMysql
	} else {
		_dbType = DBBadger
	}

	for _, chainID := range env.SUPPORTED_CHAIN_IDS {
		_blockTimeSyncMap[chainID] = &sync.Map{}
		_historicalPriceSyncMap[chainID] = &sync.Map{}
		_newVaultsFromRegistrySyncMap[chainID] = &sync.Map{}
		_vaultsSyncMap[chainID] = &sync.Map{}
	}

	wg := &sync.WaitGroup{}
	for _, chainID := range env.SUPPORTED_CHAIN_IDS {
		wg.Add(4)
		go LoadBlockTime(chainID, wg)
		go LoadHistoricalPrice(chainID, wg)
		go LoadNewVaultsFromRegistry(chainID, wg)
		go LoadVaults(chainID, wg)
	}
	wg.Wait()
}
