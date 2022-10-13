package vaults

import (
	"math"
	"math/big"
	"strconv"
	"sync"

	"github.com/yearn/ydaemon/common/bigNumber"
	"github.com/yearn/ydaemon/common/contracts"
	"github.com/yearn/ydaemon/common/ethereum"
	"github.com/yearn/ydaemon/common/logs"
	"github.com/yearn/ydaemon/common/store"
	"github.com/yearn/ydaemon/common/types/common"
	"github.com/yearn/ydaemon/external/prices"
	"github.com/yearn/ydaemon/external/tokens"
)

// yearnVaultABI takes the ABI of the standard Yearn Vault contract and prepare it for the multicall
var yearnVaultABI, _ = contracts.YearnVaultMetaData.GetAbi()

func getPricePerShare(contractAddress common.Address) ethereum.Call {
	parsedData, _ := yearnVaultABI.Pack("pricePerShare")
	return ethereum.Call{
		Target:   contractAddress.Address,
		Abi:      yearnVaultABI,
		Method:   `pricePerShare`,
		CallData: parsedData,
		Name:     contractAddress.String(),
	}
}

// FetchVaultMulticallData will perform a multicall to get some specific data from on-chain for a specific list of vaults
func FetchVaultMulticallData(chainID uint64) {
	// First we connect to the multicall client, stored in memory via the initializer.
	caller := ethereum.MulticallClientForChainID[chainID]

	// Then, for each token listed for our current chainID, we prepare the calls
	var calls = make([]ethereum.Call, 0)
	for _, vault := range tokens.Store.Tokens[chainID] {
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
	response := caller.ExecuteByBatch(calls, maxBatch, nil)
	if prices.Store.VaultPricePerShare[chainID] == nil {
		prices.Store.VaultPricePerShare[chainID] = make(map[common.Address]*bigNumber.Int)
	}
	for _, vault := range tokens.Store.Tokens[chainID] {
		if !vault.IsVault {
			continue
		}
		pricePerShareRaw := response[vault.Address.String()+`pricePerShare`]
		pricePerShare := bigNumber.NewInt()
		if len(pricePerShareRaw) == 1 {
			pricePerShare = bigNumber.SetInt(pricePerShareRaw[0].(*big.Int))
		}
		prices.Store.VaultPricePerShare[chainID][vault.Address] = pricePerShare
	}
	store.SaveInDBForChainID(store.KEYS.VaultMultiCallData, chainID, prices.Store.VaultPricePerShare[chainID])
}

// LoadVaultMulticallData will reload the multicall data store from the last state of the local Badger Database
func LoadVaultMulticallData(chainID uint64, wg *sync.WaitGroup) {
	defer wg.Done()
	temp := make(map[common.Address]*bigNumber.Int)
	if err := store.LoadFromDBForChainID(store.KEYS.VaultMultiCallData, chainID, &temp); err != nil {
		return
	}
	if temp != nil && (len(temp) > 0) {
		prices.Store.VaultPricePerShare[chainID] = temp
		logs.Success("Data loaded for the vault multicall data store for chainID: " + strconv.FormatUint(chainID, 10))
	}
}
