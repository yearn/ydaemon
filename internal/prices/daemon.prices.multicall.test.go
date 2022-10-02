package prices

import (
	"context"
	"math"
	"math/big"
	"sync"
	"testing"

	"github.com/machinebox/graphql"
	"github.com/yearn/ydaemon/internal/tokens"
	"github.com/yearn/ydaemon/internal/utils/env"
	"github.com/yearn/ydaemon/internal/utils/ethereum"
	"github.com/yearn/ydaemon/internal/utils/helpers"
	"github.com/yearn/ydaemon/internal/utils/logs"
	"github.com/yearn/ydaemon/internal/utils/models"
	"github.com/yearn/ydaemon/internal/utils/types/bigNumber"
	"github.com/yearn/ydaemon/internal/utils/types/common"
)

func fetchTokenList(chainID uint64) []common.Address {
	tokenList := []common.Address{}
	graphQLEndpoint, ok := env.THEGRAPH_ENDPOINTS[chainID]
	if !ok {
		logs.Error("No graph endpoint for chainID", chainID)
		return tokenList
	}

	client := graphql.NewClient(graphQLEndpoint)
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
		tokenList = append(tokenList, vault.ShareToken.Id)
		tokenList = append(tokenList, vault.Token.Id)
	}
	return tokenList
}

func testFetchLens(chainID uint64) {
	multiCall, ok := env.MULTICALL_ADDRESSES[chainID]
	if !ok {
		logs.Error(`Multicall address not found for chainID: `, chainID)
		return
	}

	caller := ethereum.NewMulticall(ethereum.GetRPCURI(chainID), multiCall)
	lensAddress, ok := env.LENS_ADDRESSES[chainID]
	if !ok {
		logs.Error(`Lens address not found for chainID: `, chainID)
		return
	}

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
	response := caller.ExecuteByBatch(calls, maxBatch, nil)
	if Store.TokenPrices[chainID] == nil {
		Store.TokenPrices[chainID] = make(map[common.Address]*bigNumber.Int)
	}
	for key, value := range response {
		Store.TokenPrices[chainID][common.HexToAddress(key)] = bigNumber.SetInt(value[0].(*big.Int))
	}
}

func TestMulticall(t *testing.T) {
	var wg sync.WaitGroup

	env.SetEnv(`../../cmd/.env`)
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
