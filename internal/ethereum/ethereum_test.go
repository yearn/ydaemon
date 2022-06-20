package ethereum

import (
	"context"
	"math/big"
	"sync"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/machinebox/graphql"
	"github.com/majorfi/ydaemon/internal/contracts"
	"github.com/majorfi/ydaemon/internal/logs"
	"github.com/majorfi/ydaemon/internal/models"
	"github.com/majorfi/ydaemon/internal/store"
	"github.com/majorfi/ydaemon/internal/utils"
)

var lensABI, _ = contracts.OracleMetaData.GetAbi()

func uniqueArrayAddress(arr []common.Address) []common.Address {
	occurred := map[common.Address]bool{}
	result := []common.Address{}
	for i := range arr {
		if !occurred[arr[i]] {
			occurred[arr[i]] = true
			result = append(result, arr[i])
		}
	}
	return result
}

func fetchTokenList(chainID uint64) []common.Address {
	tokenList := []common.Address{}
	client := graphql.NewClient(GetGraphURI(chainID))
	request := graphql.NewRequest(`
        {
			vaults(first: 1000) {
				shareToken {id}
				token {id}
			}
        }
    `)
	var response models.TGraphQueryResponseForVaults
	if err := client.Run(context.Background(), request, &response); err != nil {
		logs.Error(`Error fetching token list from the graph: `, err)
		return tokenList
	}

	for _, vault := range response.Vaults {
		tokenList = append(tokenList, common.HexToAddress(vault.ShareToken.Id))
		tokenList = append(tokenList, common.HexToAddress(vault.Token.Id))
	}
	return tokenList
}

func getPriceUsdcRecommendedCall(name string, contractAddress common.Address, tokenAddress common.Address) Call {
	parsedData, err := lensABI.Pack("getPriceUsdcRecommended", tokenAddress)
	if err != nil {
		return Call{
			Target:   contractAddress,
			CallData: nil,
			Name:     name,
		}
	}
	return Call{
		Target:   contractAddress,
		CallData: parsedData,
		Name:     name,
	}
}

func testFetchLens(chainID uint64) {
	// First we connect to the multicall client, stored in memory via the initializer.
	caller := NewMulticall(GetRPCURI(chainID), GetMulticallAddress(chainID))

	// Then, for the given chainID, we need to select the correct lens contract address,
	// aka the endpoint we will use to perform the read transaction.
	lensAddress := GetLensAddress(chainID)

	// Then, for each token listed for our current chainID, we prepare the calls
	var calls = make([]Call, 0)
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

func TestMulticall(t *testing.T) {
	var wg sync.WaitGroup

	utils.SetEnv(`../../cmd/.env`)

	//Testing for chainID == 1
	go func(wg *sync.WaitGroup) {
		store.TokenList[1] = uniqueArrayAddress(fetchTokenList(1))
		testFetchLens(1)
		wg.Done()
	}(&wg)

	//Testing for chainID == 4
	go func(wg *sync.WaitGroup) {
		store.TokenList[4] = uniqueArrayAddress(fetchTokenList(4))
		testFetchLens(4)
		wg.Done()
	}(&wg)

	//Testing for chainID == 250
	go func(wg *sync.WaitGroup) {
		store.TokenList[250] = uniqueArrayAddress(fetchTokenList(250))
		testFetchLens(250)
		wg.Done()
	}(&wg)

	//Testing for chainID == 42161
	go func(wg *sync.WaitGroup) {
		store.TokenList[42161] = uniqueArrayAddress(fetchTokenList(42161))
		testFetchLens(42161)
		wg.Done()
	}(&wg)

	wg.Add(4)
	wg.Wait()
}
