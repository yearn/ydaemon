package models

import (
	"github.com/ethereum/go-ethereum/common"
)

type TTokenType string

const (
	TokenTypeStandardVault           TTokenType = "Yearn Vault"
	TokenTypeExperimentalVault       TTokenType = "Experimental Yearn Vault"
	TokenTypeLegacyExperimentalVault TTokenType = "Experimental"
	TokenTypeAutomatedVault          TTokenType = "Automated Yearn Vault"
	TokenTypeLegacyAutomatedVault    TTokenType = "Automated"
	TokenTypeNative                  TTokenType = "Native"
	TokenTypeCurveLP                 TTokenType = "Curve LP"
	TokenTypeCompound                TTokenType = "Compound"
	TokenTypeAaveV1                  TTokenType = "AAVE V1"
	TokenTypeAaveV2                  TTokenType = "AAVE V2"
)

// TERC20Token contains the basic information of an ERC20 token
type TERC20Token struct {
	Address                   common.Address   `json:"address"`
	UnderlyingTokensAddresses []common.Address `json:"underlyingTokensAddresses"`
	Type                      TTokenType       `json:"type"`
	Name                      string           `json:"name"`
	Symbol                    string           `json:"symbol"`
	DisplayName               string           `json:"display_name"`
	DisplaySymbol             string           `json:"display_symbol"`
	Description               string           `json:"description"`
	Icon                      string           `json:"icon"`
	Decimals                  uint64           `json:"decimals"`
	ChainID                   uint64           `json:"chainID"`
}

// TTokenFromMeta is the structure of data we receive when calling meta.yearn.fi/api/1/tokens/all
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
