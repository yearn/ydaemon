package tokensList

import (
	"math"
	"net/http"

	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"github.com/yearn/ydaemon/common/ethereum"
	"github.com/yearn/ydaemon/common/helpers"
)

// GetTokenList will, for a given chainID, return the Yearn's Token List
func GetTokenList(c *gin.Context) {
	/**********************************************************************************************
	** The AssertChainID function takes a string as input and returns a uint64 value and a boolean
	** indicating whether the input is a valid chain ID. If the input is not a valid chain ID, the
	** function returns false.
	**********************************************************************************************/
	chainID, ok := helpers.AssertChainID(c.Param("chainID"))
	if !ok {
		c.String(http.StatusBadRequest, "invalid chainID")
		return
	}

	/**********************************************************************************************
	** The AssertAddress function takes a string as input and returns a common.Address value (which
	** is a type defined in the Ethereum Go library) and a boolean indicating whether the input is
	** a valid Ethereum address. If the input is not a valid Ethereum address, the function returns
	** false.
	**********************************************************************************************/
	userAddress, ok := helpers.AssertAddress(c.Param("address"), chainID)
	if !ok {
		c.String(http.StatusBadRequest, "invalid address")
		return
	}

	/**********************************************************************************************
	** This function makes an HTTP request to the specified URL and return a struct of type
	** DefaultTokenList representing the JSON data returned by the request, which is a default
	** token list from Uniswap Standards.
	** Then, for each token in the tokenList, the code checks if the ChainID field of the token
	** matches the chainID variable. If the ChainID field does match the chainID variable, the
	** token is added to the tokensFromListMap map using the token's address as the key and the
	** token struct as the value.
	**********************************************************************************************/
	tokensFromListMap := make(map[string]YTokenList)
	for _, token := range yTokenMap {
		if uint64(token.ChainID) == chainID {
			tokensFromListMap[token.Address] = token
		}
	}

	/**********************************************************************************************
	** Once we have our list of tokens, we can proceed to fetch the balances of the tokens for the
	** user. We create a slice of calls, which will be used to make a multicall to the blockchain.
	** For each token in the tokensFromListMap map, we add a call to the slice of calls, for the
	** balanceOf function of the ERC20 token.
	**********************************************************************************************/
	caller := ethereum.MulticallClientForChainID[chainID]
	calls := []ethereum.Call{}
	for _, token := range tokensFromListMap {
		calls = append(calls, getBalance(token.Address, ethcommon.HexToAddress(token.Address), userAddress))
	}

	/**********************************************************************************************
	** Regular fix for Fantom's RPC, which limit the number of calls in a multicall to a very low
	** number. We split the multicall in multiple calls of 3 calls each.
	** Otherwise, we just send the multicall as is.
	**********************************************************************************************/
	maxBatch := math.MaxInt64

	/**********************************************************************************************
	** The following code will execute the multicall and then map the results to the tokens in the
	** returned tokenBalanceMap map, which is a map of token addresses to YTokenList structs.
	**********************************************************************************************/
	tokenBalanceMap := make(map[string]YTokenList)
	response := caller.ExecuteByBatch(calls, maxBatch, nil)
	for _, token := range tokensFromListMap {
		rawBalance := response[token.Address+`balanceOf`]
		if len(rawBalance) == 0 {
			continue
		}
		tokenBalance := helpers.DecodeBigInt(rawBalance)
		if tokenBalance.IsZero() {
			continue
		}
		token.Balance = tokenBalance
		tokenBalanceMap[token.Address] = token
	}

	c.JSON(http.StatusOK, tokenBalanceMap)
}
