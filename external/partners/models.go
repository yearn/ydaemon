package partners

import (
	"github.com/yearn/ydaemon/common/bigNumber"
	"github.com/yearn/ydaemon/common/types/common"
)

type TPartnersFromFile struct {
	Name            string   `json:"name"`
	FullName        string   `json:"full_name"`
	Description     string   `json:"description"`
	StartDate       string   `json:"start_date"`
	RetiredTreasury []string `json:"retired_treasury"`
	Treasury        string   `json:"treasury"`
	Wrappers        []struct {
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
	Name            string             `json:"name"`
	FullName        string             `json:"full_name"`
	Description     string             `json:"description"`
	StartDate       string             `json:"start_date"`
	Treasury        common.Address     `json:"treasury"`
	RetiredTreasury []string           `json:"retired_treasury"`
	Wrappers        []TPartnersWrapper `json:"wrappers"`
}
