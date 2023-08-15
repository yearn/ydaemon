package store

import (
	"strconv"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/common/bigNumber"
	"github.com/yearn/ydaemon/common/logs"
	"github.com/yearn/ydaemon/internal/models"
	"gorm.io/gorm"
)

var _strategiesSyncMap = make(map[uint64]*sync.Map)

/**************************************************************************************************
** LoadStrategies will retrieve the all the strategies from the configured DB and store them in the
** _strategiesSyncMap for fast access during that same execution.
**************************************************************************************************/
func LoadStrategies(chainID uint64, wg *sync.WaitGroup) {
	if wg != nil {
		defer wg.Done()
	}
	syncMap := _strategiesSyncMap[chainID]
	if syncMap == nil {
		syncMap = &sync.Map{}
		_strategiesSyncMap[chainID] = syncMap
	}

	switch _dbType {
	case DBBadger:
		logs.Warning(`BadgerDB is deprecated for LoadStrategies`)
	case DBSql:
		var temp []DBStrategy
		DATABASE.Table(`db_strategies`).
			Where("chain_id = ?", chainID).
			FindInBatches(&temp, 10_000, func(tx *gorm.DB, batch int) error {
				for _, stratFromDB := range temp {
					strat := models.TStrategyAdded{
						VaultAddress:      common.HexToAddress(stratFromDB.VaultAddress),
						StrategyAddress:   common.HexToAddress(stratFromDB.StrategyAddress),
						TxHash:            common.HexToHash(stratFromDB.TxHash),
						DebtRatio:         bigNumber.NewInt(0).SetString(stratFromDB.DebtRatio),
						MaxDebtPerHarvest: bigNumber.NewInt(0).SetString(stratFromDB.MaxDebtPerHarvest),
						MinDebtPerHarvest: bigNumber.NewInt(0).SetString(stratFromDB.MinDebtPerHarvest),
						PerformanceFee:    bigNumber.NewInt(0).SetString(stratFromDB.PerformanceFee),
						DebtLimit:         bigNumber.NewInt(0).SetString(stratFromDB.DebtLimit),
						RateLimit:         bigNumber.NewInt(0).SetString(stratFromDB.RateLimit),
						VaultVersion:      stratFromDB.VaultVersion,
						ChainID:           stratFromDB.ChainID,
						BlockNumber:       stratFromDB.BlockNumber,
						TxIndex:           stratFromDB.TxIndex,
						LogIndex:          stratFromDB.LogIndex,
					}
					syncMap.Store(strat.StrategyAddress, strat)
				}
				return nil
			})
	}
}

/**************************************************************************************************
** CheckAndAppendToStrategyMap will add a new vault in the _strategiesSyncMap if it doesn't exist
**************************************************************************************************/
func CheckAndAppendToStrategyMap(chainID uint64, strat models.TStrategyAdded) bool {
	syncMap := _strategiesSyncMap[chainID]
	if syncMap == nil {
		syncMap = &sync.Map{}
		_strategiesSyncMap[chainID] = syncMap
	}

	key := strat.StrategyAddress.Hex() + "_" + strat.VaultAddress.Hex() + "_" + strat.TxHash.Hex() + "_" + strconv.FormatUint(strat.ChainID, 10)
	if data, exists := syncMap.Load(key); exists {
		if Compare(Hash(data), Hash(strat)) {
			return false
		}
	}

	syncMap.Store(strat.StrategyAddress, strat)
	return true
}

/**************************************************************************************************
** AppendToStrategyMap will add a new vault in the _strategiesSyncMap
**************************************************************************************************/
func AppendToStrategyMap(chainID uint64, strat models.TStrategyAdded) {
	syncMap := _strategiesSyncMap[chainID]
	if syncMap == nil {
		syncMap = &sync.Map{}
		_strategiesSyncMap[chainID] = syncMap
	}

	syncMap.Store(strat.StrategyAddress, strat)
}

/**************************************************************************************************
** StoreStrategies will store the new strategies in the _strategiesSyncMap for fast access during
** that same execution, and will store it in the configured DB for future executions.
**************************************************************************************************/
func StoreStrategies(chainID uint64, strat models.TStrategyAdded) {
	canAppend := CheckAndAppendToStrategyMap(chainID, strat)
	if !canAppend {
		return
	}

	key := strat.StrategyAddress.Hex() + "_" + strat.VaultAddress.Hex() + "_" + strat.TxHash.Hex() + "_" + strconv.FormatUint(strat.ChainID, 10)
	switch _dbType {
	case DBBadger:
		logs.Warning(`BadgerDB is deprecated for StoreStrategies`)
	case DBSql:
		go func() {
			newItem := &DBStrategy{}
			newItem.UUID = getUUID(key)
			newItem.VaultAddress = strat.VaultAddress.Hex()
			newItem.StrategyAddress = strat.StrategyAddress.Hex()
			newItem.TxHash = strat.TxHash.Hex()
			newItem.DebtRatio = strat.DebtRatio.String()
			newItem.MaxDebtPerHarvest = strat.MaxDebtPerHarvest.String()
			newItem.MinDebtPerHarvest = strat.MinDebtPerHarvest.String()
			newItem.PerformanceFee = strat.PerformanceFee.String()
			newItem.DebtLimit = strat.DebtLimit.String()
			newItem.RateLimit = strat.RateLimit.String()
			newItem.VaultVersion = strat.VaultVersion
			newItem.ChainID = strat.ChainID
			newItem.BlockNumber = strat.BlockNumber
			newItem.TxIndex = strat.TxIndex
			newItem.LogIndex = strat.LogIndex
			wait(`StoreStrategies`)
			DATABASE.
				Table(`db_strategies`).
				FirstOrCreate(newItem)
		}()
	}
}

/**************************************************************************************************
** ListAllStrategies will return a list of all the strategies stored in the caching system for a
** given chainID. Both a map and a slice are returned.
** Note: It's for the TStrategyAdded structure.
**************************************************************************************************/
func ListAllStrategies(chainID uint64) (
	asMap map[common.Address]models.TStrategyAdded,
	asSlice []models.TStrategyAdded,
) {
	asMap = make(map[common.Address]models.TStrategyAdded) // make to avoid nil map

	/**********************************************************************************************
	** We first retrieve the syncMap. This syncMap should be initialized first via the
	** `LoadStrategies` function which take the data from the database/badger and store it in it.
	**********************************************************************************************/
	syncMap := _strategiesSyncMap[chainID]
	if syncMap == nil {
		syncMap = &sync.Map{}
		_strategiesSyncMap[chainID] = syncMap
	}

	/**********************************************************************************************
	** We can just iterate over the syncMap and add the strategies to the map and slice.
	** As the stored vault data are only a subset of static, we need to use the actual structure
	** and not the DBVault one.
	**********************************************************************************************/
	syncMap.Range(func(key, value interface{}) bool {
		vault := value.(models.TStrategyAdded)
		asMap[vault.StrategyAddress] = vault
		asSlice = append(asSlice, vault)
		return true
	})

	return asMap, asSlice
}

/**************************************************************************************************
** GetStrategy will return a strategy stored in the caching system for a given chainID and address.
**************************************************************************************************/
func GetStrategy(chainID uint64, strategyAddress common.Address) (models.TStrategyAdded, bool) {
	/**********************************************************************************************
	** We first retrieve the syncMap. This syncMap should be initialized first via the
	** `LoadStrategies` function which take the data from the database/badger and store it in it.
	**********************************************************************************************/
	syncMap := _strategiesSyncMap[chainID]
	if syncMap == nil {
		syncMap = &sync.Map{}
		_strategiesSyncMap[chainID] = syncMap
	}

	/**********************************************************************************************
	** We can just iterate over the syncMap and add the strategies to the map and slice.
	** As the stored vault data are only a subset of static, we need to use the actual structure
	** and not the DBVault one.
	**********************************************************************************************/
	strategy, ok := syncMap.Load(strategyAddress)
	return strategy.(models.TStrategyAdded), ok
}
