package tokens

import (
	"context"
	"strconv"
	"sync"

	"github.com/machinebox/graphql"
	"github.com/yearn/ydaemon/internal/utils/env"
	"github.com/yearn/ydaemon/internal/utils/helpers"
	"github.com/yearn/ydaemon/internal/utils/logs"
	"github.com/yearn/ydaemon/internal/utils/models"
	"github.com/yearn/ydaemon/internal/utils/store"
	"github.com/yearn/ydaemon/internal/utils/types/common"
)

// FetchTokenList is an utility function that will query the subgraph in order to
// extract the list of tokens (yvTokens, aka share tokens, and underlying tokens)
// used by the Yearn system in order to be able to play with them (e.g. get the
// price though the lens contract).
func FetchTokenList(chainID uint64) {
	tokenList := []common.Address{}
	tokenData := make(map[common.Address]*TERC20Token)
	graphQLEndpoint, ok := env.THEGRAPH_ENDPOINTS[chainID]
	if !ok {
		logs.Error("No graph endpoint for chainID", chainID)
		return
	}

	client := graphql.NewClient(graphQLEndpoint)
	request := graphql.NewRequest(`
        {
			vaults(first: 1000) {
				id
				shareToken {
					id
					decimals
					name
					symbol
				}
				token {
					id
					decimals
					name
					symbol
				}
			}
        }
    `)
	var response models.TGraphQueryResponseForVaults
	if err := client.Run(context.Background(), request, &response); err != nil {
		logs.Error(`Error fetching token list from the graph: `, err)
		return
	}

	for _, vault := range response.Vaults {
		tokenList = append(tokenList, vault.ShareToken.Id)
		tokenList = append(tokenList, vault.Token.Id)

		token := &TERC20Token{
			Address:  vault.Token.Id,
			Decimals: vault.Token.Decimals,
			Name:     vault.Token.Name,
			Symbol:   vault.Token.Symbol,
			IsVault:  false,
		}
		if helpers.AddressIsValid(vault.ShareToken.Id, chainID) {
			tokenData[vault.ShareToken.Id] = &TERC20Token{
				Address:                vault.ShareToken.Id,
				UnderlyingTokenAddress: vault.Token.Id,
				Decimals:               vault.ShareToken.Decimals,
				Name:                   vault.ShareToken.Name,
				Symbol:                 vault.ShareToken.Symbol,
				IsVault:                true,
			}
		}
		if helpers.AddressIsValid(vault.Token.Id, chainID) {
			tokenData[vault.Token.Id] = token
		}

		tokenData[vault.Token.Id] = token
		if Store.VaultToToken[chainID] == nil {
			Store.VaultToToken[chainID] = make(map[common.Address]common.Address)
		}

		Store.VaultToToken[chainID][vault.Id] = vault.Token.Id
	}

	Store.Tokens[chainID] = tokenData
	Store.TokenList[chainID] = helpers.UniqueArrayAddress(tokenList)
	store.SaveInDBForChainID(store.KEYS.TokenData, chainID, Store.Tokens[chainID])
	store.SaveInDBForChainID(store.KEYS.TokenList, chainID, Store.TokenList[chainID])
}

// LoadTokenList will reload the tokenList data store from the last state of the local Badger Database
func LoadTokenList(chainID uint64, wg *sync.WaitGroup) {
	defer wg.Done()
	{
		temp := []common.Address{}
		if err := store.LoadFromDBForChainID(store.KEYS.TokenList, chainID, &temp); err != nil {
			return
		}
		if temp != nil && (len(temp) > 0) {
			Store.TokenList[chainID] = temp
			logs.Success("Data loaded for the tokenList store for chainID: " + strconv.FormatUint(chainID, 10))
		} else {
			logs.Warning("No tokenList data found for chainID: " + strconv.FormatUint(chainID, 10))
		}
	}

	{
		temp := make(map[common.Address]*TERC20Token)
		if err := store.LoadFromDBForChainID(store.KEYS.TokenData, chainID, &temp); err != nil {
			return
		}
		if temp != nil && (len(temp) > 0) {
			Store.Tokens[chainID] = temp
			logs.Success("Data loaded for the TokenData store for chainID: " + strconv.FormatUint(chainID, 10))
		} else {
			logs.Warning("No TokenData data found for chainID: " + strconv.FormatUint(chainID, 10))
		}
	}
}
