package partners

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/common/bigNumber"
)

type TExternalPartnersFromFile struct {
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

type TExternalPartnersWrapper struct {
	Vault     common.MixedcaseAddress `json:"vault,omitempty"`
	Wrapper   common.MixedcaseAddress `json:"wrapper"`
	BalanceOf *bigNumber.Int          `json:"balanceOf,omitempty"`
	Name      string                  `json:"name"`
	Type      string                  `json:"type"`
}

type TPartners struct {
	Name            string                     `json:"name"`
	FullName        string                     `json:"full_name"`
	Description     string                     `json:"description"`
	StartDate       string                     `json:"start_date"`
	Treasury        common.MixedcaseAddress    `json:"treasury"`
	RetiredTreasury []string                   `json:"retired_treasury"`
	Wrappers        []TExternalPartnersWrapper `json:"wrappers"`
}
