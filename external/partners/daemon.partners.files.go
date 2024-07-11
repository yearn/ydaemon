package partners

import (
	"encoding/json"
	"strconv"
	"sync"

	"github.com/yearn/ydaemon/common/env"
	"github.com/yearn/ydaemon/common/helpers"
	"github.com/yearn/ydaemon/common/logs"
)

// FetchPartnersFromFiles fetches the partners information from the Yearn Meta API for a given chainID
// and store the result to the global variable Partners for later use.
func FetchPartnersFromFiles(chainID uint64) {
	allPartners := []*TPartners{}
	chainIDStr := strconv.FormatUint(chainID, 10)
	content, _, err := helpers.ReadAllFilesInDir(env.BASE_DATA_PATH+`/partners/networks/`+chainIDStr+`/`, `.json`)
	if err != nil {
		logs.Warning("Error fetching meta information from the Yearn Meta API for chain", chainID)
		return
	}
	for _, content := range content {
		partner := &TPartners{}
		partnerFromFile := TExternalPartnersFromFile{}
		if err := json.Unmarshal(content, &partnerFromFile); err != nil {
			logs.Warning("Error unmarshalling response body from the Yearn Meta API for chain", chainID)
			continue
		}

		partner.Name = partnerFromFile.Name
		partner.FullName = partnerFromFile.FullName
		partner.Description = partnerFromFile.Description
		partner.StartDate = partnerFromFile.StartDate
		partner.Treasury = partnerFromFile.Treasury
		partner.RetiredTreasury = partnerFromFile.RetiredTreasury
		partner.Wrappers = make([]TExternalPartnersWrapper, len(partnerFromFile.Wrappers))
		for i, v := range partnerFromFile.Wrappers {
			if v.Vault != `` {
				partner.Wrappers[i].Vault = v.Vault
			}
			partner.Wrappers[i].Wrapper = v.Wrapper
			partner.Wrappers[i].Name = v.Name
			partner.Wrappers[i].Type = v.Type
		}
		allPartners = append(allPartners, partner)
	}

	// To provide faster access to the data, we index the mapping by the vault address, aka
	// {[vaultAddress]: TPartners} if we were working with JS/TS
	if partnersByAddress[chainID] == nil {
		partnersByAddress[chainID] = make(map[string]*TPartners)
		partnersByName[chainID] = make(map[string]*TPartners)
	}
	for _, partner := range allPartners {
		partnersByAddress[chainID][partner.Treasury] = partner
		partnersByName[chainID][partner.Name] = partner
	}
}

// LoadPartners is kept in order to have the same behavior everywhere, but as the data
// exists in the same directory as yDaemon, saving the data in the DB is not necessary.
func LoadPartners(chainID uint64, wg *sync.WaitGroup) {
	_ = chainID
	wg.Done()
}
