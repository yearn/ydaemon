package daemons

import (
	"math"
	"math/big"
	"strconv"
	"sync"

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
	maxBatch := math.MaxInt64
	if chainID == 250 {
		maxBatch = 3
	}
	response := caller.ExecuteByBatch(calls, maxBatch)
	if store.TokenPrices[chainID] == nil {
		store.TokenPrices[chainID] = make(map[common.Address]uint64)
	}
	for key, value := range response {
		store.TokenPrices[chainID][common.HexToAddress(key)] = new(big.Int).SetBytes(value.ReturnData).Uint64()
	}
	store.SaveInDBForChainID(`TokenPrices`, chainID, store.TokenPrices[chainID])
}

// LoadLens will reload the lens data store from the last state of the local Badger Database
func LoadLens(chainID uint64, wg *sync.WaitGroup) {
	defer wg.Done()
	temp := make(map[common.Address]uint64)
	err := store.LoadFromDBForChainID(`TokenPrices`, chainID, &temp)
	if err != nil {
		if err.Error() == "Key not found" {
			logs.Warning("No tokenPrices data found for chainID: " + strconv.FormatUint(chainID, 10))
			return
		}
		logs.Error(err)
		return
	}
	if temp != nil && (len(temp) > 0) {
		store.TokenPrices[chainID] = temp
		logs.Success("Data loaded for the lens data store for chainID: " + strconv.FormatUint(chainID, 10))
	} else {
		logs.Warning("No lens data found for chainID: " + strconv.FormatUint(chainID, 10))
	}
}
