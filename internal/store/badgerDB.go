package store

import (
	"encoding/json"
	"log"
	"strconv"

	"github.com/dgraph-io/badger/v3"
	"github.com/majorfi/ydaemon/internal/logs"
)

// DB is the badger database
var DB *badger.DB

// OpenDB opens the badger database
func OpenDB() {
	db, err := badger.Open(badger.DefaultOptions("./ystore"))
	if err != nil {
		log.Fatal(err)
	}
	DB = db
}

// CloseDB closes the badger database
func CloseDB() {
	DB.Close()
}

// SaveInDB saves a specific data in the badger database
func SaveInDB(key string, value interface{}) {
	err := DB.Update(func(txn *badger.Txn) error {
		vaultsBytes, err := json.Marshal(value)
		if err != nil {
			return err
		}
		return txn.Set([]byte(key), vaultsBytes)
	})
	if err != nil {
		logs.Error(err)
	}
}

// SaveInDBForChainID saves a specific data in the badger database for a specific chainID
func SaveInDBForChainID(key string, chainID uint64, value interface{}) {
	SaveInDB(key+strconv.FormatUint(chainID, 10), value)
}

// LoadFromDB saves a specific data in the badger database
func LoadFromDB(key string, dest interface{}) error {
	return DB.View(func(txn *badger.Txn) error {
		marshalizedData := []byte{}
		dataFromKey, err := txn.Get([]byte(key))
		if err != nil {
			return err
		}

		err = dataFromKey.Value(func(v []byte) error {
			marshalizedData = v
			return nil
		})
		if err != nil {
			return err
		}

		return json.Unmarshal([]byte(marshalizedData), &dest)
	})
}

// LoadFromDBForChainID saves a specific data in the badger database for a specific chainID
func LoadFromDBForChainID(key string, chainID uint64, dest interface{}) error {
	err := LoadFromDB(key+strconv.FormatUint(chainID, 10), &dest)
	if err != nil {
		return err
	}
	return nil
}
