package tokensList

import (
	"net/http"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"github.com/yearn/ydaemon/common/addresses"
	"github.com/yearn/ydaemon/common/bigNumber"
	"github.com/yearn/ydaemon/common/env"
	"github.com/yearn/ydaemon/common/ethereum"
	"github.com/yearn/ydaemon/common/helpers"
	"github.com/yearn/ydaemon/common/logs"
	"github.com/yearn/ydaemon/internal/multicalls"
	"github.com/yearn/ydaemon/internal/storage"
)

// SupportedZap is the type of zap supported by the token (Wido, Portals, Cowswap)
type SupportedZap string

const (
	Wido    SupportedZap = "Wido"    // Wido indicated that [Wido](https://www.joinwido.com) supports this token
	Portals SupportedZap = "Portals" // Portals indicated that [Portals](https://portals.fi) supports this token
	Cowswap SupportedZap = "Cowswap" // Cowswap indicated that [Cowswap](https://cowswap.exchange) supports this token
)

// TTokenListToken is the structure used for the tokens inside a token list.
// This is a standard and alterations should be avoided.
type TTokenListToken struct {
	ChainID       int            `json:"chainID"`           // ChainID indicates on which chain the token is deployed
	Decimals      int            `json:"decimals"`          // Decimals is the number of decimals the token uses
	Address       string         `json:"address"`           // Address is the address of the token contract
	Name          string         `json:"name"`              // Name is the name of the token
	Symbol        string         `json:"symbol"`            // Symbol is the symbol of the token
	LogoURI       string         `json:"logoURI"`           // LogoURI is the URI of the token logo
	Balance       *bigNumber.Int `json:"balance,omitempty"` // Balance is the balance of the token for the user
	Price         *bigNumber.Int `json:"price,omitempty"`   // Price is the price of the token
	SupportedZaps []SupportedZap `json:"supportedZaps"`     // SupportedZaps is the list of zaps supported by the token
}

// TTokenList is the token list struct used in the default token list.
// This is a standard and alterations should be avoided.
// More details: https://tokenlistooor.com/tokenlistooor
type TTokenList struct {
	Name      string                 `json:"name"`      // Name is the name of the token list
	Timestamp string                 `json:"timestamp"` // Timestamp indicates when the token list was last updated
	LogoURI   string                 `json:"logoURI"`   // LogoURI is the URI of the token list logo
	Keywords  []string               `json:"keywords"`  // Keywords is a list of keywords used to find the token list
	Tags      map[string]interface{} `json:"tags"`      // Tags is a map of tags used to categorize the token list
	Tokens    []TTokenListToken      `json:"tokens"`    // Tokens is a list of tokens in that token list
}

var existingTokenLists = map[string]map[string]TTokenListToken{}

func init() {
	yearnList := helpers.FetchJSON[TTokenList](`https://raw.githubusercontent.com/SmolDapp/tokenLists/main/lists/yearn.json`)
	cowswapList := helpers.FetchJSON[TTokenList](`https://raw.githubusercontent.com/SmolDapp/tokenLists/main/lists/cowswap.json`)
	widoList := helpers.FetchJSON[TTokenList](`https://raw.githubusercontent.com/SmolDapp/tokenLists/main/lists/wido.json`)
	portals := helpers.FetchJSON[TTokenList](`https://raw.githubusercontent.com/SmolDapp/tokenLists/main/lists/portals.json`)

	yearnMap := make(map[string]TTokenListToken)
	cowswapMap := make(map[string]TTokenListToken)
	widoMap := make(map[string]TTokenListToken)
	portalsMap := make(map[string]TTokenListToken)

	for _, token := range yearnList.Tokens {
		yearnMap[strconv.FormatInt(int64(token.ChainID), 10)+common.HexToAddress(token.Address).Hex()] = token
	}
	existingTokenLists[`yearn`] = yearnMap

	for _, token := range cowswapList.Tokens {
		cowswapMap[strconv.FormatInt(int64(token.ChainID), 10)+common.HexToAddress(token.Address).Hex()] = token
	}
	existingTokenLists[`cowswap`] = cowswapMap

	for _, token := range widoList.Tokens {
		widoMap[strconv.FormatInt(int64(token.ChainID), 10)+common.HexToAddress(token.Address).Hex()] = token
	}
	existingTokenLists[`wido`] = widoMap

	for _, token := range portals.Tokens {
		portalsMap[strconv.FormatInt(int64(token.ChainID), 10)+common.HexToAddress(token.Address).Hex()] = token
	}
	existingTokenLists[`portals`] = portalsMap

	logs.Pretty(len(portalsMap), len(widoMap), len(cowswapMap), len(yearnMap))
}

func getSupportedZaps(chainID uint64, tokenAddress common.Address) []SupportedZap {
	supportedZaps := []SupportedZap{}
	if _, ok := existingTokenLists[`wido`][strconv.FormatInt(int64(chainID), 10)+tokenAddress.Hex()]; ok {
		supportedZaps = append(supportedZaps, Wido)
	}
	if _, ok := existingTokenLists[`cowswap`][strconv.FormatInt(int64(chainID), 10)+tokenAddress.Hex()]; ok {
		supportedZaps = append(supportedZaps, Cowswap)
	}
	if _, ok := existingTokenLists[`portals`][strconv.FormatInt(int64(chainID), 10)+tokenAddress.Hex()]; ok {
		supportedZaps = append(supportedZaps, Portals)
	}
	return supportedZaps
}

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
	** Once we have our list of tokens, we can proceed to fetch the balances of the tokens for the
	** user. We create a slice of calls, which will be used to make a multicall to the blockchain.
	** For each token in the tokensFromListMap map, we add a call to the slice of calls, for the
	** balanceOf function of the ERC20 token.
	**********************************************************************************************/
	calls := []ethereum.Call{}
	for _, token := range existingTokenLists[`yearn`] {
		calls = append(calls, getBalance(
			token.Address,
			addresses.ToMixedcase(token.Address),
			addresses.ToMixedcase(userAddress),
		))
	}

	/**********************************************************************************************
	** The following code will execute the multicall and then map the results to the tokens in the
	** returned tokenBalanceMap map, which is a map of token addresses to TTokenListToken structs.
	**********************************************************************************************/
	tokenBalanceMap := make(map[string]TTokenListToken)

	/**********************************************************************************************
	** We first need to load the current chainCoin, which is the native coin of the chain. We then
	** check if the price of the chainCoin is available in the prices map. If the price is
	** available, we add the price to the chainToken struct. Then, we add the chainToken struct to
	** the tokenBalanceMap map using the chainCoin address as the key.
	**********************************************************************************************/
	chainCoin := env.CHAINS[chainID].Coin
	chainCoinPrice, ok := storage.GetPrice(chainID, chainCoin.Address)
	chainToken := TTokenListToken{
		ChainID:       int(chainID),
		Address:       chainCoin.Address.Hex(),
		Name:          chainCoin.DisplayName,
		Symbol:        chainCoin.DisplaySymbol,
		LogoURI:       chainCoin.Icon,
		Decimals:      int(chainCoin.Decimals),
		Balance:       bigNumber.NewInt(0),
		Price:         bigNumber.NewInt(0),
		SupportedZaps: []SupportedZap{},
	}
	if ok {
		chainToken.Price = chainCoinPrice.Price
	}
	tokenBalanceMap[chainCoin.Address.Hex()] = chainToken

	/**********************************************************************************************
	** And we can finally execute the multicall.
	**********************************************************************************************/
	response := multicalls.Perform(chainID, calls, nil)
	for _, token := range existingTokenLists[`yearn`] {
		rawBalance := response[token.Address+`balanceOf`]
		if len(rawBalance) == 0 {
			continue
		}
		tokenBalance := helpers.DecodeBigInt(rawBalance)
		if tokenBalance.IsZero() {
			continue
		}
		token.Balance = tokenBalance
		if tokenPrice, ok := storage.GetPrice(
			chainID,
			addresses.ToAddress(token.Address),
		); ok {
			token.Price = tokenPrice.Price
			token.SupportedZaps = getSupportedZaps(chainID, addresses.ToAddress(token.Address))
			tokenBalanceMap[token.Address] = token
		}
	}

	c.JSON(http.StatusOK, tokenBalanceMap)
}
