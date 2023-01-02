package tokensList

import "github.com/yearn/ydaemon/common/bigNumber"

// DefaultTokenListToken is the token struct used in the default token list
type DefaultTokenListToken struct {
	ChainID  int    `json:"chainId"`
	Decimals int    `json:"decimals"`
	Count    int    `json:"count"`
	Address  string `json:"address"`
	Name     string `json:"name"`
	Symbol   string `json:"symbol"`
	LogoURI  string `json:"logoURI"`
}

// DefaultTokenListData is the token list struct used in the default token list
type DefaultTokenListData struct {
	Name      string `json:"name"`
	Timestamp string `json:"timestamp"`
	Version   struct {
		Major int `json:"major"`
		Minor int `json:"minor"`
		Patch int `json:"patch"`
	} `json:"version"`
	Tags     map[string]interface{}  `json:"tags"`
	LogoURI  string                  `json:"logoURI"`
	Keywords []string                `json:"keywords"`
	Tokens   []DefaultTokenListToken `json:"tokens"`
}

type PortalsTokenListData struct {
	TotalItems int `json:"totalItems"`
	Tokens     []struct {
		Name      string  `json:"name"`
		Decimals  int     `json:"decimals"`
		Symbol    string  `json:"symbol"`
		Price     float64 `json:"price"`
		Address   string  `json:"address"`
		Platform  string  `json:"platform"`
		Network   string  `json:"network"`
		Liquidity float64 `json:"liquidity"`
		Image     string  `json:"image"`
	} `json:"tokens"`
}

/* 🔵 - Yearn Finance **************************************************
** Yearn's own token list, aka what is kept and returned
**********************************************************************/

// SupportedZap is the type of zap supported by the token (Wido, Portals, Cowswap)
type SupportedZap string

const (
	Wido    SupportedZap = "Wido"
	Portals SupportedZap = "Portals"
	Cowswap SupportedZap = "Cowswap"
)

type YTokenList struct {
	ChainID       uint64         `json:"chainID"`
	Address       string         `json:"address"`
	Name          string         `json:"name"`
	Symbol        string         `json:"symbol"`
	LogoURI       string         `json:"logoURI"`
	Decimals      int            `json:"decimals"`
	Balance       *bigNumber.Int `json:"balance"`
	SupportedZaps []SupportedZap `json:"supportedZaps"`
}
