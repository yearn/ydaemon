package risks

import (
	"errors"
	"fmt"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/common/env"
	"github.com/yearn/ydaemon/common/helpers"
	"github.com/yearn/ydaemon/internal/models"
)

type TRiskScoreYsec struct {
	RiskLevel int8              `json:"riskLevel"` // The risk level of the vault (1 to 5, -1 if not set)
	RiskScore models.TRiskScore `json:"riskScore"` // The risk score of the vault
}

var (
	allRisksScores = make(map[uint64]map[common.Address]TRiskScoreYsec)
	riskScoresMtx  sync.RWMutex
)

func RetrieveAvailableRiskScores(chainID uint64) map[common.Address]bool {
	result := make(map[common.Address]bool)

	manifestURL := env.RISK_CDN_URL + fmt.Sprintf("vaults/%d.json", chainID)
	manifest, err := helpers.FetchJSONWithReject[map[string]TRiskScoreYsec](manifestURL)
	if err == nil {
		riskScoresMtx.Lock()
		if allRisksScores[chainID] == nil {
			allRisksScores[chainID] = make(map[common.Address]TRiskScoreYsec)
		}
		for addressStr, riskScore := range manifest {
			address := common.HexToAddress(addressStr)
			allRisksScores[chainID][address] = riskScore
			result[address] = true
		}
		riskScoresMtx.Unlock()
	}

	return result
}

/**************************************************************************************************
** GetCachedRiskScore will fetch the risk scores for a specific vault from a GitHub repository.
** The scores are stored in JSON files organized by chain ID and vault address.
** If the risk scores are not found, it will return an empty TRiskScore structure.
**
** Arguments:
** - chainID: the chain ID of the network we are working on
** - vaultAddress: the address of the vault to fetch the risk scores for
**
** Returns:
** - a TRiskScoreYsec structure containing the risk scores for the vault
**************************************************************************************************/
func GetCachedRiskScore(chainID uint64, vaultAddress common.Address) (TRiskScoreYsec, error) {
	riskScoresMtx.RLock()
	defer riskScoresMtx.RUnlock()

	if _, ok := allRisksScores[chainID]; !ok {
		return TRiskScoreYsec{}, errors.New("chainID not found")
	}
	if _, ok := allRisksScores[chainID][vaultAddress]; !ok {
		return TRiskScoreYsec{}, errors.New("vaultAddress not found")
	}
	return allRisksScores[chainID][vaultAddress], nil
}
