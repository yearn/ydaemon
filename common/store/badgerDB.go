package store

import (
	"encoding/json"
	"errors"
	"log"
	"reflect"
	"strconv"
	"sync"

	"github.com/dgraph-io/badger/v3"
	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/common/addresses"
	"github.com/yearn/ydaemon/common/env"
	"github.com/yearn/ydaemon/common/logs"
	"github.com/yearn/ydaemon/common/traces"
)

// DB is the badger database
var _assertMutex sync.Mutex
var _isLoaded bool
var DB = make(map[uint64]map[string]*badger.DB)
var DBMutex = make(map[uint64]map[string]*sync.Mutex)

func _assertMap() {
	if !_isLoaded {
		_assertMutex.Lock()
		defer _assertMutex.Unlock()
		if DB == nil {
			DB = make(map[uint64]map[string]*badger.DB)
		}
		if DBMutex == nil {
			DBMutex = make(map[uint64]map[string]*sync.Mutex)
		}

		for _, chainID := range env.SUPPORTED_CHAIN_IDS {
			if DB[chainID] == nil {
				DB[chainID] = make(map[string]*badger.DB)
			}
			if DBMutex[chainID] == nil {
				DBMutex[chainID] = make(map[string]*sync.Mutex)
				DBMutex[chainID][TABLES.BLOCK_TIME] = &sync.Mutex{}
				DBMutex[chainID][TABLES.PRICES] = &sync.Mutex{}
				DBMutex[chainID][TABLES.STRATEGIES] = &sync.Mutex{}
				DBMutex[chainID][TABLES.TOKENS] = &sync.Mutex{}
				DBMutex[chainID][TABLES.VAULTS] = &sync.Mutex{}
				DBMutex[chainID][TABLES.VAULTS_LEGACY] = &sync.Mutex{}
			}
		}
		_isLoaded = true
	}
}

// OpenDB opens the badger database
func OpenDB(chainID uint64, dbKey string) *badger.DB {
	_assertMap()
	if DB[chainID][dbKey] == nil {
		DBMutex[chainID][dbKey].Lock()
		defer DBMutex[chainID][dbKey].Unlock()
		chainStr := strconv.FormatUint(chainID, 10)
		opts := badger.DefaultOptions(`./data/store/` + chainStr + `/` + dbKey)
		opts = opts.WithMetricsEnabled(false)
		opts = opts.WithLogger(nil)
		opts = opts.WithNumGoroutines(16)
		opts = opts.WithValueLogFileSize((1 << 20) * 10)
		db, err := badger.Open(opts)
		if err != nil {
			log.Fatal(err)
		}
		DB[chainID][dbKey] = db
	}
	return DB[chainID][dbKey]
}

// CloseDB closes the badger database
func CloseDB(chainID uint64, dbKey string) {
	DB[chainID][dbKey].Close()
}

// SaveInDB saves a specific data in the badger database
func SaveInDB(chainID uint64, dbKey string, dataKey string, value interface{}) {
	currentStore := OpenDB(chainID, dbKey) //Safely open the DB if it's not already open

	err := currentStore.Update(func(txn *badger.Txn) error {
		vaultsBytes, err := json.Marshal(value)
		if err != nil {
			return err
		}
		return txn.Set([]byte(dataKey), vaultsBytes)
	})
	if err != nil {
		traces.
			Capture(`error`, `impossible to save in DB`).
			SetEntity(`database`).
			SetExtra(`error`, err.Error()).
			SetExtra(`value`, value).
			SetTag(`chainID`, strconv.FormatUint(chainID, 10)).
			SetTag(`dbKey`, dbKey).
			SetTag(`dataKey`, dataKey).
			Send()
	}
}

// LoadFromDB saves a specific data in the badger database
func LoadFromDB(chainID uint64, dbKey string, dataKey string, dest interface{}) error {
	currentStore := OpenDB(chainID, dbKey) //Safely open the DB if it's not already open

	return currentStore.View(func(txn *badger.Txn) error {
		marshalizedData := []byte{}
		dataFromKey, err := txn.Get([]byte(dataKey))
		if err != nil {
			if err.Error() == "Key not found" {
				logs.Warning(dataKey + " not found for chainID: " + strconv.FormatUint(chainID, 10))
				return err
			}
			logs.Error(err)
			return err
		}

		if err = dataFromKey.Value(func(v []byte) error {
			marshalizedData = v
			return nil
		}); err != nil {
			logs.Error(err)
			return err
		}

		return json.Unmarshal([]byte(marshalizedData), &dest)
	})
}

/**************************************************************************************************
** Iterate tries to load the data from the DB for a specific table Key to the provided dest
** pointer.
**************************************************************************************************/
func Iterate(chainID uint64, dbKey string, dest interface{}) error {
	currentStore := OpenDB(chainID, dbKey) //Safely open the DB if it's not already open

	return currentStore.View(func(txn *badger.Txn) error {
		it := txn.NewIterator(badger.DefaultIteratorOptions)
		defer it.Close()

		t := reflect.TypeOf(dest)
		if t.Kind() != reflect.Ptr {
			return errors.New("dest must be a pointer to a slice")
		}

		elem := t.Elem()
		if elem.Kind() == reflect.Map {
			elemKey := elem.Key()
			elemValue := elem.Elem()
			destMap := reflect.MakeMap(elem)
			for it.Rewind(); it.Valid(); it.Next() {
				item := it.Item()
				k := item.Key()
				v, err := item.ValueCopy(nil)
				if err != nil {
					traces.
						Capture(`error`, `impossible to get value for key: `+string(k)).
						SetEntity(`database`).
						SetExtra(`error`, err.Error()).
						SetTag(`chainID`, strconv.FormatUint(chainID, 10)).
						SetTag(`dbKey`, dbKey).
						Send()
					continue
				}

				/**********************************************************************************
				** Detect the key we are working with and assign it to the correct type of Type.
				** Supported types are string and common.Address
				**********************************************************************************/
				keyType := reflect.New(elemKey)
				key := keyType.Interface()
				if elemKey.Kind() == reflect.String {
					key = string(k)
				} else if elemKey.Kind() == reflect.Struct && elemKey.Name() == `MixedcaseAddress` {
					key = addresses.ToMixedcase(string(k))
				} else if elemKey.Kind() == reflect.Struct && elemKey.Name() == `Address` {
					key = common.HexToAddress(string(k))
				} else if elemKey.Kind() == reflect.Array && elemKey.Name() == `Address` {
					var addressString string
					for _, v := range k {
						addressString += string(v)
					}
					key = common.HexToAddress(addressString)
				} else if elemKey.Kind() == reflect.Uint64 {
					var asString string
					for _, v := range k {
						asString += string(v)
					}
					key, _ = strconv.ParseUint(asString, 10, 64)
				} else {
					_ = key
					logs.Pretty(k)
					logs.Error(`unsupported key type: `+elemKey.Name()+` for kind: `, elemKey.Kind())
					continue
				}

				/**********************************************************************************
				** Assign the value to the destination expected type via a standard unmarshal
				** function. Only JSON kind of values are supported.
				**********************************************************************************/
				valueType := reflect.New(elemValue)
				value := valueType.Interface()
				if err := json.Unmarshal(v, &value); err != nil {
					logs.Pretty(value, valueType, elemValue)
					logs.Error(`impossible to unmarshal value for key: ` + string(k))
					continue
				}

				destMap.SetMapIndex(
					reflect.ValueOf(key),
					reflect.ValueOf(reflect.ValueOf(value).Elem().Interface()),
				)
			}
			/**************************************************************************************
			** Assign our reflected value in our dest pointer
			**************************************************************************************/
			reflect.ValueOf(dest).Elem().Set(destMap)
		}
		return nil
	})
}
