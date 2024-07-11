package partners

import (
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
	Vault     string         `json:"vault,omitempty"`
	Wrapper   string         `json:"wrapper"`
	BalanceOf *bigNumber.Int `json:"balanceOf,omitempty"`
	Name      string         `json:"name"`
	Type      string         `json:"type"`
}

type TPartners struct {
	Name            string                     `json:"name"`
	FullName        string                     `json:"full_name"`
	Description     string                     `json:"description"`
	StartDate       string                     `json:"start_date"`
	Treasury        string                     `json:"treasury"`
	RetiredTreasury []string                   `json:"retired_treasury"`
	Wrappers        []TExternalPartnersWrapper `json:"wrappers"`
}

var partnersByAddress map[uint64]map[string]*TPartners
var partnersByName map[uint64]map[string]*TPartners

func init() {
	partnersByAddress = make(map[uint64]map[string]*TPartners)
	partnersByName = make(map[uint64]map[string]*TPartners)
}
