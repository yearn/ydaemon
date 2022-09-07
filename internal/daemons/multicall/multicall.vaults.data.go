package multicallDaemons

import (
	"math"
	"math/big"
	"strconv"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/internal/ethereum"
	"github.com/yearn/ydaemon/internal/utils/logs"
	"github.com/yearn/ydaemon/internal/utils/store"
)

func getPricePerShare(contractAddress common.Address) ethereum.Call {
	parsedData, _ := yearnVaultABI.Pack("pricePerShare")
	return ethereum.Call{
		Target:   contractAddress,
		Abi:      yearnVaultABI,
		Method:   `pricePerShare`,
		CallData: parsedData,
		Name:     contractAddress.String(),
	}
}

// FetchVaultMulticallData will perform a multicall to get some specific data from on-chain for a specific list of vaults
func FetchVaultMulticallData(chainID uint64) {
	// First we connect to the multicall client, stored in memory via the initializer.
	caller := multicallClientForChainID[chainID]

	// Then, for each token listed for our current chainID, we prepare the calls
	var calls = make([]ethereum.Call, 0)
	for _, vault := range store.Tokens[chainID] {
		if !vault.IsVault {
			continue
		}
		calls = append(calls, getPricePerShare(vault.Address))
	}

	if len(calls) == 0 {
		logs.Error("No vault found.")
		return
	}

	// Then, we execute the multicall and store the prices in the TokenPrices map
	maxBatch := math.MaxInt64
	response := caller.ExecuteByBatch(calls, maxBatch)
	if store.VaultPricePerShare[chainID] == nil {
		store.VaultPricePerShare[chainID] = make(map[common.Address]*big.Int)
	}
	for _, vault := range store.Tokens[chainID] {
		if !vault.IsVault {
			continue
		}
		pricePerShareRaw := response[vault.Address.String()+`pricePerShare`]
		pricePerShare := new(big.Int)
		if len(pricePerShareRaw) == 1 {
			pricePerShare = pricePerShareRaw[0].(*big.Int)
		}
		store.VaultPricePerShare[chainID][vault.Address] = pricePerShare
	}
	store.SaveInDBForChainID(`VaultMultiCallData`, chainID, store.VaultPricePerShare[chainID])
}

// LoadVaultMulticallData will reload the multicall data store from the last state of the local Badger Database
func LoadVaultMulticallData(chainID uint64, wg *sync.WaitGroup) {
	defer wg.Done()
	temp := make(map[common.Address]*big.Int)
	err := store.LoadFromDBForChainID(`VaultMultiCallData`, chainID, &temp)
	if err != nil {
		if err.Error() == "Key not found" {
			logs.Warning("No vault multicall data found for chainID: " + strconv.FormatUint(chainID, 10))
			return
		}
		logs.Error(err)
		return
	}
	if temp != nil && (len(temp) > 0) {
		store.VaultPricePerShare[chainID] = temp
		logs.Success("Data loaded for the vault multicall data store for chainID: " + strconv.FormatUint(chainID, 10))
	} else {
		logs.Warning("No vault multicall data found for chainID: " + strconv.FormatUint(chainID, 10))
	}
}
