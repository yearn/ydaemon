package daemons

import (
	"context"
	"strconv"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/machinebox/graphql"
	"github.com/majorfi/ydaemon/internal/ethereum"
	"github.com/majorfi/ydaemon/internal/logs"
	"github.com/majorfi/ydaemon/internal/models"
	"github.com/majorfi/ydaemon/internal/store"
	"github.com/majorfi/ydaemon/internal/utils"
)

// FetchTokenList is an utility function that will query the subgraph in order to
// extract the list of tokens (yvTokens, aka share tokens, and underlying tokens)
// used by the Yearn system in order to be able to play with them (e.g. get the
// price though the lens contract).
func FetchTokenList(chainID uint64) {
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
		return
	}

	for _, vault := range response.Vaults {
		tokenList = append(tokenList, common.HexToAddress(vault.ShareToken.Id))
		tokenList = append(tokenList, common.HexToAddress(vault.Token.Id))
	}
	store.TokenList[chainID] = utils.UniqueArrayAddress(tokenList)
	store.SaveInDBForChainID(`TokenList`, chainID, store.TokenList[chainID])
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
