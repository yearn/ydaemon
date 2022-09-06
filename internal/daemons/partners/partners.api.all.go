package partnersDaemons

import (
	"encoding/json"
	"strconv"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/internal/partners"
	"github.com/yearn/ydaemon/internal/utils/helpers"
	"github.com/yearn/ydaemon/internal/utils/logs"
)

// FetchPartnersFromFiles fetches the partners information from the Yearn Meta API for a given chainID
// and store the result to the global variable Partners for later use.
func FetchPartnersFromFiles(chainID uint64) {
	allPartners := []partners.TPartners{}
	chainIDStr := strconv.FormatUint(chainID, 10)
	content, _, err := helpers.ReadAllFilesInDir(helpers.BASE_DATA_PATH+`/partners/networks/`+chainIDStr+`/`, `.json`)
	if err != nil {
		logs.Warning("Error fetching meta information from the Yearn Meta API")
		return
	}
	for _, content := range content {
		partner := partners.TPartners{}
		partnerFromFile := partners.TPartnersFromFile{}
		if err := json.Unmarshal(content, &partnerFromFile); err != nil {
			logs.Warning("Error unmarshalling response body from the Yearn Meta API")
			continue
		}

		partner.Name = partnerFromFile.Name
		partner.StartBlock = partnerFromFile.StartBlock
		partner.Treasury = common.HexToAddress(partnerFromFile.Treasury)
		partner.Wrappers = make([]partners.TPartnersWrapper, len(partnerFromFile.Wrappers))
		for i, v := range partnerFromFile.Wrappers {
			if v.Vault != `` {
				partner.Wrappers[i].Vault = common.HexToAddress(v.Vault)
			}
			partner.Wrappers[i].Wrapper = common.HexToAddress(v.Wrapper)
			partner.Wrappers[i].Name = v.Name
			partner.Wrappers[i].Type = v.Type
		}
		allPartners = append(allPartners, partner)
	}

	// To provide faster access to the data, we index the mapping by the vault address, aka
	// {[vaultAddress]: TPartners} if we were working with JS/TS
	if partners.Store.PartnersByAddress[chainID] == nil {
		partners.Store.PartnersByAddress[chainID] = make(map[common.Address]partners.TPartners)
		partners.Store.PartnersByName[chainID] = make(map[string]partners.TPartners)
	}
	for _, partner := range allPartners {
		partners.Store.PartnersByAddress[chainID][partner.Treasury] = partner
		partners.Store.PartnersByName[chainID][partner.Name] = partner
	}
}

// LoadPartners is kept in order to have the same behavior everywhere, but as the data
// exists in the same directory as yDaemon, saving the data in the DB is not necessary.
func LoadPartners(chainID uint64, wg *sync.WaitGroup) {
	_ = chainID
	wg.Done()
}
