package daemons

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/machinebox/graphql"
	"github.com/majorfi/ydaemon/internal/ethereum"
	"github.com/majorfi/ydaemon/internal/logs"
	"github.com/majorfi/ydaemon/internal/models"
)

// fetchTokenList is an utility function that will query the subgraph in order to
// extract the list of tokens (yvTokens, aka share tokens, and underlying tokens)
// used by the Yearn system in order to be able to play with them (e.g. get the
// price though the lens contract).
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
