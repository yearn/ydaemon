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

	wg := &sync.WaitGroup{}
	for _, chainID := range env.SUPPORTED_CHAIN_IDS {
		wg.Add(2)
		go ListBlockTime(chainID, wg)
		go ListHistoricalPrice(chainID, wg)
	}
	wg.Wait()
}
