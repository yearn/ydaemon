package risks

import (
	"errors"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/common/helpers"
	"github.com/yearn/ydaemon/common/logs"
	"github.com/yearn/ydaemon/internal/models"
)

type TRiskScoreYsec struct {
	RiskLevel int8              `json:"riskLevel"` // The risk level of the vault (1 to 5, -1 if not set)
	RiskScore models.TRiskScore `json:"riskScore"` // The risk score of the vault
}

var allRisksScores = make(map[uint64]map[common.Address]TRiskScoreYsec)

/**************************************************************************************************
** fetchVaultsRiskScore will fetch the risk scores for a specific vault from a GitHub repository.
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
func fetchVaultsRiskScore(chainID uint64, vaultAddress common.Address) (TRiskScoreYsec, error) {
	baseURL := "https://raw.githubusercontent.com/spalen0/risk-score/refs/heads/master/strategy/"
	uri := baseURL + strconv.FormatUint(chainID, 10) + "/" + vaultAddress.Hex() + ".json"
	riskScores, err := helpers.FetchJSONWithReject[TRiskScoreYsec](uri)
	if err != nil {
		logs.Error(err)
		return TRiskScoreYsec{}, err
	}
	return riskScores, nil
}

/**************************************************************************************************
** RetrieveAllRiskScores will fetch the risk scores for a list of vaults from a GitHub repository.
** The scores are stored in JSON files organized by chain ID and vault address.
** If the risk scores are not found, it will return an empty TRiskScore structure.
**
** Arguments:
** - chainID: the chain ID of the network we are working on
** - vaults: the list of vaults to fetch the risk scores for
**
** Returns:
** - a map of vaultAddress -> TRiskScoreYsec
**************************************************************************************************/
func RetrieveAllRiskScores(chainID uint64, vaults map[common.Address]models.TVault) map[common.Address]TRiskScoreYsec {
	riskScores := make(map[common.Address]TRiskScoreYsec)
	if _, ok := allRisksScores[chainID]; !ok {
		allRisksScores[chainID] = make(map[common.Address]TRiskScoreYsec)
	}

	for _, vault := range vaults {
		result, err := fetchVaultsRiskScore(chainID, vault.Address)
		if err != nil {
			logs.Error(err)
			continue
		}
		riskScores[vault.Address] = result
		allRisksScores[chainID][vault.Address] = riskScores[vault.Address]
	}
	return riskScores
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
	if _, ok := allRisksScores[chainID]; !ok {
		return TRiskScoreYsec{}, errors.New("chainID not found")
	}
	if _, ok := allRisksScores[chainID][vaultAddress]; !ok {
		return TRiskScoreYsec{}, errors.New("vaultAddress not found")
	}
	return allRisksScores[chainID][vaultAddress], nil
}
