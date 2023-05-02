package tokensList

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/yearn/ydaemon/common/addresses"
	"github.com/yearn/ydaemon/common/bigNumber"
	"github.com/yearn/ydaemon/common/env"
	"github.com/yearn/ydaemon/common/ethereum"
	"github.com/yearn/ydaemon/common/helpers"
	"github.com/yearn/ydaemon/internal/models"
	"github.com/yearn/ydaemon/internal/multicalls"
	"github.com/yearn/ydaemon/internal/prices"
	"github.com/yearn/ydaemon/internal/tokens"
	"github.com/yearn/ydaemon/processes/tokenList"
)

// GetYearnTokenList will, for a given chainID, return the Yearn's Token List
func GetYearnTokenList(c *gin.Context) {
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
	** The AssertAddress function takes a string as input and returns a mixedcaseAddress value (which
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
	tokensFromListMap := tokenList.MapTokenList(chainID)

	/**********************************************************************************************
	** Once we have our list of tokens, we can proceed to fetch the balances of the tokens for the
	** user. We create a slice of calls, which will be used to make a multicall to the blockchain.
	** For each token in the tokensFromListMap map, we add a call to the slice of calls, for the
	** balanceOf function of the ERC20 token.
	**********************************************************************************************/
	calls := []ethereum.Call{}
	for _, token := range tokensFromListMap {
		calls = append(calls, getBalance(
			token.Address,
			addresses.ToMixedcase(token.Address),
			addresses.ToMixedcase(userAddress),
		))
	}

	/**********************************************************************************************
	** The following code will execute the multicall and then map the results to the tokens in the
	** returned tokenBalanceMap map, which is a map of token addresses to TYearnTokenListToken
	** structs.
	**********************************************************************************************/
	tokenBalanceMap := make(map[string]models.TYearnTokenListToken)

	/**********************************************************************************************
	** We first need to load the current chainCoin, which is the native coin of the chain. We then
	** check if the price of the chainCoin is available in the prices map. If the price is
	** available, we add the price to the chainToken struct. Then, we add the chainToken struct to
	** the tokenBalanceMap map using the chainCoin address as the key.
	**********************************************************************************************/
	chainCoin := tokens.COIN_PER_CHAIN[chainID]
	chainCoinPrice, ok := prices.FindPrice(chainID, chainCoin.Address)
	chainToken := models.TYearnTokenListToken{
		TTokenListToken: models.TTokenListToken{
			ChainID:  int(chainID),
			Address:  chainCoin.Address.Hex(),
			Name:     chainCoin.DisplayName,
			Symbol:   chainCoin.DisplaySymbol,
			LogoURI:  chainCoin.Icon,
			Decimals: int(chainCoin.Decimals),
		},
		Balance:       bigNumber.NewInt(0),
		Price:         bigNumber.NewInt(0),
		SupportedZaps: []models.SupportedZap{models.Wido},
	}
	if ok {
		chainToken.Price = chainCoinPrice
	}
	tokenBalanceMap[chainCoin.Address.Hex()] = chainToken

	/**********************************************************************************************
	** And we can finally execute the multicall.
	**********************************************************************************************/
	response := multicalls.Perform(chainID, calls, nil)
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
		if tokenPrice, ok := prices.FindPrice(
			chainID,
			addresses.ToAddress(token.Address),
		); ok {
			token.Price = tokenPrice
			tokenBalanceMap[token.Address] = token
		}
	}

	c.JSON(http.StatusOK, tokenBalanceMap)
}

// GetTokenList will return the Yearn's Token List
func GetTokenList(c *gin.Context) {
	var list models.TTokenList

	list.Name = "Yearn Token List"
	list.Keywords = []string{"yearn", "yfi", "yvault", "ytoken", "ycurve", "yprotocol", "vaults"}
	list.LogoURI = "https://raw.githubusercontent.com/yearn/yearn-assets/3ec995a8b19cd95e56a1a42b18d394d667e0e2cd/icons/multichain-tokens/1/0x0bc529c00C6401aEF6D220BE8C6Ea1667F6Ad93e/logo.svg"
	list.Timestamp = time.Now().Format(time.RFC3339)
	list.Version.Major = 1
	list.Version.Minor = 0
	list.Version.Patch = 0

	/**********************************************************************************************
	** Retrieve the MapTokenList for each chainID and return it as a map of chainID to token list.
	**********************************************************************************************/
	for _, chainID := range env.SUPPORTED_CHAIN_IDS {
		allTokens := tokens.ListTokens(chainID)
		for _, token := range allTokens {
			list.Tokens = append(list.Tokens, models.TTokenListToken{
				ChainID:  int(chainID),
				Address:  token.Address.Hex(),
				Name:     token.Name,
				Symbol:   token.Symbol,
				Decimals: int(token.Decimals),
				LogoURI:  env.BASE_ASSET_URL + strconv.FormatUint(chainID, 10) + `/` + token.Address.Hex() + `/logo-128.png`,
			})
		}
	}

	c.JSON(http.StatusOK, list)
}
