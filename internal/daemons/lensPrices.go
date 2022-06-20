package daemons

import (
	"math/big"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/majorfi/ydaemon/internal/contracts"
	"github.com/majorfi/ydaemon/internal/ethereum"
	"github.com/majorfi/ydaemon/internal/logs"
	"github.com/majorfi/ydaemon/internal/store"
)

// lensABI takes the ABI of the lens contract and prepare it for the multicall
var lensABI, _ = contracts.OracleMetaData.GetAbi()

// getPriceUsdcRecommendedCall will pack the transaction with it's argument and return the
// ethereum.Call object that can be used to execute the transaction.
func getPriceUsdcRecommendedCall(name string, contractAddress common.Address, tokenAddress common.Address) ethereum.Call {
	parsedData, err := lensABI.Pack("getPriceUsdcRecommended", tokenAddress)
	if err != nil {
		return ethereum.Call{
			Target:   contractAddress,
			CallData: nil,
			Name:     name,
		}
	}
	return ethereum.Call{
		Target:   contractAddress,
		CallData: parsedData,
		Name:     name,
	}
}

// FetchLens will fetch the prices of the yvTokens and the tokens that are listed in the
// corresponding subgraph, execute a multicall on the given chainID to retreive them,
// and will store the prices in the TokenPrices map.
func FetchLens(chainID uint64) {
	// First we connect to the multicall client, stored in memory via the initializer.
	caller := multicallClientForChainID[chainID]

	// Then, for the given chainID, we need to select the correct lens contract address,
	// aka the endpoint we will use to perform the read transaction.
	lensAddress := ethereum.GetLensAddress(chainID)

	// Then, for each token listed for our current chainID, we prepare the calls
	var calls = make([]ethereum.Call, 0)
	for _, token := range store.TokenList[chainID] {
		calls = append(calls, getPriceUsdcRecommendedCall(token.String(), lensAddress, token))
	}

	if len(calls) == 0 {
		logs.Error("No tokens found.")
		return
	}

	// Then, we execute the multicall and store the prices in the TokenPrices map
	if chainID == 250 {
		response := caller.ExecuteByBatch(calls, 5)
		if store.TokenPrices[chainID] == nil {
			store.TokenPrices[chainID] = make(map[common.Address]uint64)
		}
		for key, value := range response {
			store.TokenPrices[chainID][common.HexToAddress(key)] = new(big.Int).SetBytes(value.ReturnData).Uint64()
		}
	} else {
		response := caller.Execute(calls)
		if store.TokenPrices[chainID] == nil {
			store.TokenPrices[chainID] = make(map[common.Address]uint64)
		}
		for key, value := range response {
			store.TokenPrices[chainID][common.HexToAddress(key)] = new(big.Int).SetBytes(value.ReturnData).Uint64()
		}
	}
}

// RunLens is a goroutine that periodically execute a multicall on the given chainID
// to retreive the prices of the yvTokens and the tokens that are listed in the
// corresponding subgraph.
func RunLens(chainID uint64, wg *sync.WaitGroup) {
	// First, we fetch the initial tokenList for this chain
	store.TokenList[chainID] = uniqueArrayAddress(fetchTokenList(chainID))

	// Then, we fetch the prices of the tokens, in loop.
	isDone := false
	for {
		FetchLens(chainID)
		if !isDone {
			isDone = true
			logs.Info("Lens prices oracle is ready")
			wg.Done()
		}
		time.Sleep(30 * time.Second)
	}
}
