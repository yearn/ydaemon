package tokensList

import "github.com/yearn/ydaemon/common/bigNumber"

// DefaultTokenListToken is the token struct used in the default token list
type DefaultTokenListToken struct {
	ChainID    int    `json:"chainId"`
	Address    string `json:"address"`
	Name       string `json:"name"`
	Symbol     string `json:"symbol"`
	Decimals   int    `json:"decimals"`
	LogoURI    string `json:"logoURI"`
	Extensions struct {
		BridgeInfo map[string]struct {
			TokenAddress string `json:"tokenAddress"`
		} `json:"bridgeInfo"`
	} `json:"extensions"`
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

// YTokenList is the structure used for our tokenList override with the SupportedZaps field
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
