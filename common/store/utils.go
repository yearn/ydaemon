package store

import (
	"bytes"
	"context"
	"encoding/gob"
	"encoding/json"
	"strconv"

	"github.com/dgraph-io/badger/v3"
	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/common/logs"
	"golang.org/x/time/rate"
)

var storeRateLimiter = rate.NewLimiter(4, 8)

func wait(name string) {
	// logs.Warning(`Limiter for ` + name + `: ` + strconv.FormatFloat(storeRateLimiter.Tokens(), 'f', 2, 64))
	storeRateLimiter.Wait(context.Background())
}

func StoreRateLimiter() *rate.Limiter {
	return storeRateLimiter
}

func Compare(a, b []byte) bool {
	a = append(a, b...)
	c := 0
	for _, x := range a {
		c ^= int(x)
	}
	return c == 0
}

func Hash(s any) []byte {
	var b bytes.Buffer
	gob.NewEncoder(&b).Encode(s)
	return b.Bytes()
}

/**************************************************************************************************
** LoadSyncStrategiesAdded will try to retrieve all the sync for vaults/strategies for a given
** chain on the local DB.
**************************************************************************************************/
func LoadSyncStrategiesAdded(chainID uint64) []DBStrategyAddedSync {
	syncMap := make(map[string]DBStrategyAddedSync)
	ListFromBadgerDB(chainID, TABLES.STRATEGIES_FROM_VAULT_SYNC, &syncMap)
	arr := []DBStrategyAddedSync{}
	for _, v := range syncMap {
		arr = append(arr, v)
	}
	return arr
}

/**************************************************************************************************
** LoadSyncRegistry will try to retrieve all the sync for registry/vault for a given chain on local
** DB.
**************************************************************************************************/
func LoadSyncRegistry(chainID uint64) []DBRegistrySync {
	syncMap := make(map[string]DBRegistrySync)
	ListFromBadgerDB(chainID, TABLES.REGISTRY_SYNC, &syncMap)
	arr := []DBRegistrySync{}
	for _, v := range syncMap {
		arr = append(arr, v)
	}
	return arr
}

/**************************************************************************************************
** StoreSyncRegistry will store the sync status indicating we went up to the block number to check
** for new vaults.
**************************************************************************************************/
func StoreSyncRegistry(chainID uint64, registryAddess common.Address, end *uint64) {
	OpenBadgerDB(chainID, TABLES.REGISTRY_SYNC).Update(func(txn *badger.Txn) error {
		if getter, err := txn.Get([]byte(registryAddess.Hex())); err == nil {
			if previousValue, err := getter.ValueCopy(nil); err == nil {
				previousSync := &DBRegistrySync{}
				if err := json.Unmarshal(previousValue, previousSync); err != nil {
					logs.Error(`StoreSyncRegistry: json.Unmarshal(previousValue, previousSync)`, err)
					return err
				}
				if previousSync.Block > *end {
					logs.Info(`StoreSyncRegistry: previousSync.Block > *end`, previousSync.Block, *end)
					return nil
				}
			}
		}

		data := &DBRegistrySync{
			ChainID:  chainID,
			Block:    *end,
			Registry: registryAddess.Hex(),
			UUID:     GetUUID(registryAddess.Hex() + strconv.FormatUint(chainID, 10)),
		}
		dataByte, err := json.Marshal(data)
		if err != nil {
			return err
		}
		return txn.Set([]byte(data.Registry), dataByte)
	})
}

/**************************************************************************************************
** StoreSyncStrategiesAdded will store the sync status indicating we went up to the block number
** to check for new strategies added.
**************************************************************************************************/
func StoreSyncStrategiesAdded(chainID uint64, vaultAddress common.Address, end uint64) {
	OpenBadgerDB(
		chainID,
		TABLES.STRATEGIES_FROM_VAULT_SYNC,
	).Update(func(txn *badger.Txn) error {
		if getter, err := txn.Get([]byte(vaultAddress.Hex())); err == nil {
			if previousValue, err := getter.ValueCopy(nil); err == nil {
				previousSync := &DBRegistrySync{}
				if err := json.Unmarshal(previousValue, previousSync); err != nil {
					logs.Error(`StoreSyncStrategiesAdded: json.Unmarshal(previousValue, previousSync)`, err)
					return err
				}
				if previousSync.Block > end {
					logs.Info(`StoreSyncStrategiesAdded: previousSync.Block > end`, previousSync.Block, end)
					return nil
				}
			}
		}

		data := &DBStrategyAddedSync{
			ChainID: chainID,
			Block:   end,
			Vault:   vaultAddress.Hex(),
			UUID:    GetUUID(vaultAddress.Hex() + strconv.FormatUint(chainID, 10)),
		}
		dataByte, err := json.Marshal(data)
		if err != nil {
			logs.Error(`StoreSyncStrategiesAdded: json.Marshal(data)`, err)
			return err
		}
		return txn.Set([]byte(data.Vault), dataByte)
	})
}
