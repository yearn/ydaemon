package partners

import "github.com/yearn/ydaemon/internal/types/common"

type TStore struct {
	// PartnersByAddress holds the data ordered by address.
	PartnersByAddress map[uint64]map[common.Address]*TPartners
	// PartnersByAddress holds the data ordered by name
	PartnersByName map[uint64]map[string]*TPartners
}

// Store holds the data for the partners for each chain.
var Store = TStore{}

func init() {
	Store.PartnersByAddress = make(map[uint64]map[common.Address]*TPartners)
	Store.PartnersByName = make(map[uint64]map[string]*TPartners)
}
