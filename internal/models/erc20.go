package models

import (
	"github.com/ethereum/go-ethereum/common"
)

type TTokenType string

const (
	TokenTypeStandardVault           TTokenType = "Yearn Vault"
	TokenTypeLegagyStandardVault     TTokenType = "Standard"
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
// Category can be "Curve", "Stablecoin", "DeFi", "Balancer", "Currency", "Staking", "Smart Contract Platform", "Special",
type TERC20Token struct {
	Address                   common.Address   `json:"address"`
	UnderlyingTokensAddresses []common.Address `json:"underlyingTokensAddresses"`
	Type                      TTokenType       `json:"type"`
	Name                      string           `json:"name"`
	Symbol                    string           `json:"symbol"`
	DisplayName               string           `json:"displayName"`
	DisplaySymbol             string           `json:"displaySymbol"`
	Description               string           `json:"description"`
	Category                  string           `json:"category"`
	Icon                      string           `json:"icon"`
	Decimals                  uint64           `json:"decimals"`
	ChainID                   uint64           `json:"chainID"`
}

/** 🔵 - Yearn *************************************************************************************
** A few getters for the TERC20Token structure.
** All of theses getter expends the TERC20Token structure with additional information.
**************************************************************************************************/
//IsVaultLike will check if the provided token is of the "Yearn Vault" type.
func (t TERC20Token) IsVaultLike() bool {
	return t.IsStandardVault() || t.IsExperimentalVault() || t.IsAutomatedVault()
}

// IsStandardVault checks if the provided token is of the "Standard Vault" type.
func (t TERC20Token) IsStandardVault() bool {
	return t.Type == TokenTypeStandardVault || t.Type == TokenTypeLegagyStandardVault
}

// IsExperimentalVault checks if the provided token is of the "Experimental Vault" type.
func (t TERC20Token) IsExperimentalVault() bool {
	return t.Type == TokenTypeExperimentalVault || t.Type == TokenTypeLegacyExperimentalVault
}

// IsAutomatedVaultchecks if the provided token is of the "Automated Vault" type.
func (t TERC20Token) IsAutomatedVault() bool {
	return t.Type == TokenTypeAutomatedVault || t.Type == TokenTypeLegacyAutomatedVault
}

// GetVaultUnderlyingAddress will, return the address of the underlying token of the specific ERC20
func (t TERC20Token) GetVaultUnderlying() common.Address {
	if len(t.UnderlyingTokensAddresses) > 0 {
		return t.UnderlyingTokensAddresses[0]
	}
	return common.Address{}
}
