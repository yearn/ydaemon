package partners

import (
	"github.com/yearn/ydaemon/internal/types/bigNumber"
	"github.com/yearn/ydaemon/internal/types/common"
)

type TPartnersFromFile struct {
	Name       string `json:"name"`
	StartBlock uint64 `json:"start_block"`
	Treasury   string `json:"treasury"`
	Wrappers   []struct {
		Vault   string `json:"vault,omitempty"`
		Wrapper string `json:"wrapper"`
		Name    string `json:"name"`
		Type    string `json:"type"`
	} `json:"wrappers"`
}

type TPartnersWrapper struct {
	Vault     common.Address `json:"vault,omitempty"`
	Wrapper   common.Address `json:"wrapper"`
	BalanceOf *bigNumber.Int `json:"balanceOf,omitempty"`
	Name      string         `json:"name"`
	Type      string         `json:"type"`
}

type TPartners struct {
	Name       string             `json:"name"`
	StartBlock uint64             `json:"start_block"`
	Treasury   common.Address     `json:"treasury"`
	Wrappers   []TPartnersWrapper `json:"wrappers"`
}
