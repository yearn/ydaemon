package daemons

import (
	"strconv"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/majorfi/ydaemon/internal/logs"
	"github.com/majorfi/ydaemon/internal/store"
	"github.com/majorfi/ydaemon/internal/utils"
)

// RunTokenList is a goroutine that periodically execute a graph request for a
// given chain to retreive the list of tokens used in Yearn's ecosystem.
func RunTokenList(chainID uint64, wg *sync.WaitGroup) {
	isDone := false
	for {
		store.TokenList[chainID] = utils.UniqueArrayAddress(fetchTokenList(chainID))
		store.SaveInDBForChainID(`TokenList`, chainID, store.TokenList[chainID])
		if !isDone {
			isDone = true
			wg.Done()
		}
		time.Sleep(1 * time.Hour)
	}
}

// LoadTokenList will reload the tokenList data store from the last state of the local Badger Database
func LoadTokenList(chainID uint64, wg *sync.WaitGroup) {
	defer wg.Done()
	temp := []common.Address{}
	err := store.LoadFromDBForChainID(`TokenList`, chainID, &temp)
	if err != nil {
		if err.Error() == "Key not found" {
			logs.Warning("No TokenList data found for chainID: " + strconv.FormatUint(chainID, 10))
			return
		}
		logs.Error(err)
		return
	}
	if temp != nil && (len(temp) > 0) {
		store.TokenList[chainID] = temp
		logs.Success("Data loaded for the tokenList store for chainID: " + strconv.FormatUint(chainID, 10))
	} else {
		logs.Warning("No tokenList data found for chainID: " + strconv.FormatUint(chainID, 10))
	}
}
