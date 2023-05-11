package models

import (
	"time"

	"github.com/yearn/ydaemon/common/bigNumber"
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
	ChainID  int    `json:"chainID"`  // ChainID indicates on which chain the token is deployed
	Decimals int    `json:"decimals"` // Decimals is the number of decimals the token uses
	Address  string `json:"address"`  // Address is the address of the token contract
	Name     string `json:"name"`     // Name is the name of the token
	Symbol   string `json:"symbol"`   // Symbol is the symbol of the token
	LogoURI  string `json:"logoURI"`  // LogoURI is the URI of the token logo
}

// TTokenListVersion is the structure used to indicate the version of a token list.
// It's divided in 3 parts: Major, Minor and Patch.
// Major: Indicates that a token was removed from the list
// Minor: Indicates that a token was added to the list
// Patch: Indicates that a token was updated in the list
type TTokenListVersion struct {
	Major int `json:"major"`
	Minor int `json:"minor"`
	Patch int `json:"patch"`
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
	Version   TTokenListVersion      `json:"version"`   // Version specifies the version of the token list
	Tokens    []TTokenListToken      `json:"tokens"`    // Tokens is a list of tokens in that token list
}

// Yearn's own token list, aka what is kept and returned
type TYearnTokenListToken struct {
	TTokenListToken
	// ChainID       uint64         `json:"chainID"`
	// Address       string         `json:"address"`
	// Name          string         `json:"name"`
	// Symbol        string         `json:"symbol"`
	// LogoURI       string         `json:"logoURI"`
	// Decimals      int            `json:"decimals"`
	Balance       *bigNumber.Int `json:"balance"`
	Price         *bigNumber.Int `json:"price,omitempty"`
	SupportedZaps []SupportedZap `json:"supportedZaps"`
}
type YTokenList struct {
	Name      string                 `json:"name"`
	LogoURI   string                 `json:"logoURI"`
	Timestamp time.Time              `json:"timestamp"`
	Tokens    []TYearnTokenListToken `json:"tokens"`
	Version   struct {
		Major int `json:"major"`
		Minor int `json:"minor"`
		Patch int `json:"patch"`
	} `json:"version"`
}
