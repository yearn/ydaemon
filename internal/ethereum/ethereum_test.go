package ethereum

import (
	"context"
	"math"
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
	maxBatch := math.MaxInt64
	if chainID == 250 {
		maxBatch = 3
	}
	response := caller.ExecuteByBatch(calls, maxBatch)
	if store.TokenPrices[chainID] == nil {
		store.TokenPrices[chainID] = make(map[common.Address]*big.Int)
	}
	for key, value := range response {
		store.TokenPrices[chainID][common.HexToAddress(key)] = value[0].(*big.Int)
	}
}

func TestMulticall(t *testing.T) {
	var wg sync.WaitGroup

	utils.SetEnv(`../../cmd/.env`)

	//Testing for chainID == 1
	go func(wg *sync.WaitGroup) {
		store.TokenList[1] = utils.UniqueArrayAddress(fetchTokenList(1))
		testFetchLens(1)
		wg.Done()
	}(&wg)

	//Testing for chainID == 4
	go func(wg *sync.WaitGroup) {
		store.TokenList[4] = utils.UniqueArrayAddress(fetchTokenList(4))
		testFetchLens(4)
		wg.Done()
	}(&wg)

	//Testing for chainID == 250
	go func(wg *sync.WaitGroup) {
		store.TokenList[250] = utils.UniqueArrayAddress(fetchTokenList(250))
		testFetchLens(250)
		wg.Done()
	}(&wg)

	//Testing for chainID == 42161
	go func(wg *sync.WaitGroup) {
		store.TokenList[42161] = utils.UniqueArrayAddress(fetchTokenList(42161))
		testFetchLens(42161)
		wg.Done()
	}(&wg)

	wg.Add(4)
	wg.Wait()
}
