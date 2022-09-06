package partnersDaemons

import (
	"encoding/json"
	"strconv"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/internal/logs"
	"github.com/yearn/ydaemon/internal/models"
	"github.com/yearn/ydaemon/internal/store"
	"github.com/yearn/ydaemon/internal/utils"
)

// FetchPartnersFromFiles fetches the partners information from the Yearn Meta API for a given chainID
// and store the result to the global variable Partners for later use.
func FetchPartnersFromFiles(chainID uint64) {
	partners := []models.TPartners{}
	chainIDStr := strconv.FormatUint(chainID, 10)
	content, _, err := utils.ReadAllFilesInDir(utils.PARTNERS_BASE_PATH+`/networks/`+chainIDStr+`/`, `.json`)
	if err != nil {
		logs.Error("Error fetching meta information from the Yearn Meta API", err)
		return
	}
	for _, content := range content {
		partner := models.TPartners{}
		partnerFromFile := models.TPartnersFromFile{}
		if err := json.Unmarshal(content, &partnerFromFile); err != nil {
			logs.Error("Error unmarshalling response body from the Yearn Meta API", err)
			continue
		}

		partner.Name = partnerFromFile.Name
		partner.StartBlock = partnerFromFile.StartBlock
		partner.Treasury = common.HexToAddress(partnerFromFile.Treasury)
		partner.Wrappers = make([]models.TPartnersWrapper, len(partnerFromFile.Wrappers))
		for i, v := range partnerFromFile.Wrappers {
			if v.Vault != `` {
				partner.Wrappers[i].Vault = common.HexToAddress(v.Vault)
			}
			partner.Wrappers[i].Wrapper = common.HexToAddress(v.Wrapper)
			partner.Wrappers[i].Name = v.Name
			partner.Wrappers[i].Type = v.Type
		}
		partners = append(partners, partner)
	}

	// To provide faster access to the data, we index the mapping by the vault address, aka
	// {[vaultAddress]: TPartners} if we were working with JS/TS
	if store.PartnersByAddress[chainID] == nil {
		store.PartnersByAddress[chainID] = make(map[common.Address]models.TPartners)
		store.PartnersByName[chainID] = make(map[string]models.TPartners)
	}
	for _, partner := range partners {
		store.PartnersByAddress[chainID][partner.Treasury] = partner
		store.PartnersByName[chainID][partner.Name] = partner
	}
}

// LoadPartners is kept in order to have the same behavior everywhere, but as the data
// exists in the same directory as yDaemon, saving the data in the DB is not necessary.
func LoadPartners(chainID uint64, wg *sync.WaitGroup) {
	_ = chainID
	wg.Done()
}
