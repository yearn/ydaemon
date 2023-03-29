package tokens

import (
	"github.com/ethereum/go-ethereum/common"
)

// TERC20Token contains the basic information of an ERC20 token
type TERC20Token struct {
	Address                   common.Address   `json:"address"`
	UnderlyingTokensAddresses []common.Address `json:"underlyingTokensAddresses"`
	Name                      string           `json:"name"`
	Symbol                    string           `json:"symbol"`
	Type                      string           `json:"type"`
	DisplayName               string           `json:"display_name"`
	DisplaySymbol             string           `json:"display_symbol"`
	Description               string           `json:"description"`
	Icon                      string           `json:"icon"`
	Decimals                  uint64           `json:"decimals"`
}

// IsVault will return true if the token is a Yearn Vault token.
func (t *TERC20Token) IsVault() bool {
	return t.Type == "Yearn Vault"
}

// VaultUnderlying will return the underlying token of a vault. A vault can only and
// will only have one underlying token.
func (t *TERC20Token) VaultUnderlying() common.Address {
	if len(t.UnderlyingTokensAddresses) > 0 {
		return t.UnderlyingTokensAddresses[0]
	}
	return common.Address{}
}

/**********************************************************************************************
** Set of functions to store and retrieve the tokens from the cache and/or database and being
** able to access them from the rest of the application.
** The _tokenMap variable is not exported and is only used internally by the functions below.
**********************************************************************************************/
var _tokenMap = make(map[uint64]map[common.Address]*TERC20Token)

/**********************************************************************************************
** ListTokens will, for a given chainID, return the list of all the tokens stored in _tokenMap.
**********************************************************************************************/
func ListTokens(chainID uint64) []*TERC20Token {
	var tokens []*TERC20Token
	for _, token := range _tokenMap[chainID] {
		tokens = append(tokens, token)
	}
	return tokens
}

/**********************************************************************************************
** ListTokensAddresses will, for a given chainID, return the list of addresses of all the
** tokens stored in _tokenMap.
**********************************************************************************************/
func ListTokensAddresses(chainID uint64) []common.Address {
	var addresses []common.Address
	for address := range _tokenMap[chainID] {
		addresses = append(addresses, address)
	}
	return addresses
}

/**********************************************************************************************
** ListVaults will, for a given chainID, return the list of tokens with the "Yearn Vault" type
** stored in _tokenMap.
**********************************************************************************************/
func ListVaults(chainID uint64) []*TERC20Token {
	var tokens []*TERC20Token
	for _, token := range _tokenMap[chainID] {
		if token.IsVault() {
			tokens = append(tokens, token)
		}
	}
	return tokens
}

/**********************************************************************************************
** MapVaults will, for a given chainID, return a map of tokens with the "Yearn Vault" type
** stored in _tokenMap.
** The map will be of the form: map[vaultAddress]vaultToken
**********************************************************************************************/
func MapVaults(chainID uint64) map[common.Address]*TERC20Token {
	tokens := make(map[common.Address]*TERC20Token)
	for _, token := range _tokenMap[chainID] {
		if token.IsVault() {
			tokens[token.Address] = token
		}
	}
	return tokens
}

/**********************************************************************************************
** FindToken will, for a given chainID, try to find the provided tokenAddress stored in _tokenMap.
** It will return the token if found, and a boolean indicating if the token was found or not.
**********************************************************************************************/
func FindToken(chainID uint64, tokenAddress common.Address) (*TERC20Token, bool) {
	token, ok := _tokenMap[chainID][tokenAddress]
	if !ok {
		return nil, false
	}
	return token, true
}

/**********************************************************************************************
** FindUnderlyingForVault will, for a given chainID, try to find the first underlying token
** of the provided vaultAddress stored in _tokenMap.
** As the Yearn's vaults can only have one underlying token, it will return the first one found
** of the list of underlying tokens.
** It will return the token if found, and a boolean indicating if the token was found or not.
**********************************************************************************************/
func FindUnderlyingForVault(chainID uint64, vaultAddress common.Address) (*TERC20Token, bool) {
	token, ok := _tokenMap[chainID][vaultAddress]
	if !ok {
		return nil, false
	}
	return FindToken(chainID, token.VaultUnderlying())
}
