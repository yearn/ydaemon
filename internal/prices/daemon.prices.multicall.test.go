package prices

import (
	"context"
	"math"
	"math/big"
	"sync"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/machinebox/graphql"
	"github.com/yearn/ydaemon/internal/tokens"
	"github.com/yearn/ydaemon/internal/utils/ethereum"
	"github.com/yearn/ydaemon/internal/utils/helpers"
	"github.com/yearn/ydaemon/internal/utils/logs"
	"github.com/yearn/ydaemon/internal/utils/models"
)

func fetchTokenList(chainID uint64) []common.Address {
	tokenList := []common.Address{}
	client := graphql.NewClient(ethereum.GetGraphURI(chainID))
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

func testFetchLens(chainID uint64) {
	// First we connect to the multicall client, stored in memory via the initializer.
	caller := ethereum.NewMulticall(ethereum.GetRPCURI(chainID), ethereum.GetMulticallAddress(chainID))

	// Then, for the given chainID, we need to select the correct lens contract address,
	// aka the endpoint we will use to perform the read transaction.
	lensAddress := ethereum.GetLensAddress(chainID)

	// Then, for each token listed for our current chainID, we prepare the calls
	var calls = make([]ethereum.Call, 0)
	for _, token := range tokens.Store.TokenList[chainID] {
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
	if Store.TokenPrices[chainID] == nil {
		Store.TokenPrices[chainID] = make(map[common.Address]*big.Int)
	}
	for key, value := range response {
		Store.TokenPrices[chainID][common.HexToAddress(key)] = value[0].(*big.Int)
	}
}

func TestMulticall(t *testing.T) {
	var wg sync.WaitGroup

	helpers.SetEnv(`../../cmd/.env`)
	wg.Add(5)

	//Testing for chainID == 1
	go func(wg *sync.WaitGroup) {
		tokens.Store.TokenList[1] = helpers.UniqueArrayAddress(fetchTokenList(1))
		testFetchLens(1)
		wg.Done()
	}(&wg)

	//Testing for chainID == 4
	go func(wg *sync.WaitGroup) {
		tokens.Store.TokenList[4] = helpers.UniqueArrayAddress(fetchTokenList(4))
		testFetchLens(4)
		wg.Done()
	}(&wg)

	//Testing for chainID == 10
	go func(wg *sync.WaitGroup) {
		tokens.Store.TokenList[10] = helpers.UniqueArrayAddress(fetchTokenList(10))
		testFetchLens(10)
		wg.Done()
	}(&wg)

	//Testing for chainID == 250
	go func(wg *sync.WaitGroup) {
		tokens.Store.TokenList[250] = helpers.UniqueArrayAddress(fetchTokenList(250))
		testFetchLens(250)
		wg.Done()
	}(&wg)

	//Testing for chainID == 42161
	go func(wg *sync.WaitGroup) {
		tokens.Store.TokenList[42161] = helpers.UniqueArrayAddress(fetchTokenList(42161))
		testFetchLens(42161)
		wg.Done()
	}(&wg)

	wg.Wait()
}
