package models

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

// TTokenFromMeta is the structure of data we receive when calling meta.yearn.finance/api/1/tokens/all
type TTokenFromMeta struct {
	Address      common.Address `json:"address"`
	Name         string         `json:"name"`
	Symbol       string         `json:"symbol"`
	Description  string         `json:"description"`
	Website      string         `json:"website"`
	Categories   []string       `json:"categories"`
	ChainID      uint64         `json:"chainID"`
	Localization *TLocalization `json:"localization,omitempty"`
}
