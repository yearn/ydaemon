package tokens

import "github.com/yearn/ydaemon/internal/types/common"

// TERC20Token contains the basic information of an ERC20 token
type TERC20Token struct {
	Address                common.Address `json:"address"`
	UnderlyingTokenAddress common.Address `json:"underlyingTokenAddress,omitempty"`
	Name                   string         `json:"name"`
	Symbol                 string         `json:"symbol"`
	Price                  float64        `json:"price"`
	Decimals               uint64         `json:"decimals"`
	IsVault                bool           `json:"isVault"`
}
