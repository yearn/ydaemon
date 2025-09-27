package risks

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/common/helpers"
	"github.com/yearn/ydaemon/common/logs"
	"github.com/yearn/ydaemon/internal/models"
)

type TRiskScoreYsec struct {
	RiskLevel int8              `json:"riskLevel"` // The risk level of the vault (1 to 5, -1 if not set)
	RiskScore models.TRiskScore `json:"riskScore"` // The risk score of the vault
}

var (
	allRisksScores         = make(map[uint64]map[common.Address]TRiskScoreYsec)
	availableRiskScores    = make(map[uint64]map[common.Address]bool)
	riskScoresMtx          sync.RWMutex
	availableRiskScoresMtx sync.RWMutex
)

type TGithubTreeResponse struct {
	Tree []struct {
		Path string `json:"path"`
	} `json:"tree"`
}

/**************************************************************************************************
** fetchVaultsRiskScore will fetch the risk scores for a specific vault from a GitHub repository.
** The scores are stored in JSON files organized by chain ID and vault address.
** This function tries multiple locations and address formats in order:
** 1. strategy/ directory with checksummed address (legacy location)
** 2. strategy/ directory with lowercase address (legacy location)
** 3. vaults/ directory with checksummed address (new location)
** 4. vaults/ directory with lowercase address (new location)
** If the risk scores are not found in any location/format, it will return an error.
**
** Arguments:
** - chainID: the chain ID of the network we are working on
** - vaultAddress: the address of the vault to fetch the risk scores for
**
** Returns:
** - a TRiskScoreYsec structure containing the risk scores for the vault
**************************************************************************************************/
func fetchVaultsRiskScore(chainID uint64, vaultAddress common.Address) (TRiskScoreYsec, error) {
	// baseURL := "https://raw.githubusercontent.com/yearn/risk-score/refs/heads/master/"
	baseURL := "https://risk.yearn.fi/cdn/"
	chainIDStr := strconv.FormatUint(chainID, 10)
	vaultHex := vaultAddress.Hex()
	vaultHexLower := strings.ToLower(vaultHex)

	vaultsURILowercase := baseURL + "vaults/" + chainIDStr + "/" + vaultHexLower + ".json"
	riskScores, err := helpers.FetchJSONWithReject[TRiskScoreYsec](vaultsURILowercase)
	if err == nil {
		logs.Success("Fetch risk score (lowercase)/")
		return riskScores, nil
	}

	vaultsURIChecksummed := baseURL + "vaults/" + chainIDStr + "/" + vaultHex + ".json"
	riskScores, err = helpers.FetchJSONWithReject[TRiskScoreYsec](vaultsURIChecksummed)
	if err == nil {
		logs.Success("Fetch risk score (checksummed)/")
		return riskScores, nil
	}

	logs.Warning("No risk scores", vaultAddress.Hex())
	return TRiskScoreYsec{}, err
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

	riskScoresMtx.Lock()
	if _, ok := allRisksScores[chainID]; !ok {
		allRisksScores[chainID] = make(map[common.Address]TRiskScoreYsec)
	}
	riskScoresMtx.Unlock()

	availableRiskScoresMtx.RLock()
	isAvailable := availableRiskScores[chainID]
	availableRiskScoresMtx.RUnlock()

	for _, vault := range vaults {
		if !isAvailable[vault.Address] {
			continue
		}
		result, err := fetchVaultsRiskScore(chainID, vault.Address)
		if err != nil {
			logs.Error(err)
			continue
		}
		riskScores[vault.Address] = result

		riskScoresMtx.Lock()
		allRisksScores[chainID][vault.Address] = result
		riskScoresMtx.Unlock()
	}
	return riskScores
}

func RetrieveAvailableRiskScores(chainID uint64) map[common.Address]bool {
	availableRiskScoresMtx.Lock()
	defer availableRiskScoresMtx.Unlock()

	if availableRiskScores[chainID] == nil {
		availableRiskScores[chainID] = make(map[common.Address]bool)
	}

	// Fetch the GitHub tree
	treeResponse := helpers.FetchJSON[TGithubTreeResponse]("https://api.github.com/repos/yearn/risk-score/git/trees/master?recursive=1")

	// Parse the tree to find risk scores for this chain in both directories
	strategyPrefix := fmt.Sprintf("strategy/%d/", chainID)
	vaultsPrefix := fmt.Sprintf("vaults/%d/", chainID)

	for _, item := range treeResponse.Tree {
		var addressStr string

		// Check strategy directory (legacy location)
		if strings.HasPrefix(item.Path, strategyPrefix) && strings.HasSuffix(item.Path, ".json") {
			addressStr = strings.TrimSuffix(strings.TrimPrefix(item.Path, strategyPrefix), ".json")
		}
		// Check vaults directory (new location)
		if strings.HasPrefix(item.Path, vaultsPrefix) && strings.HasSuffix(item.Path, ".json") {
			addressStr = strings.TrimSuffix(strings.TrimPrefix(item.Path, vaultsPrefix), ".json")
		}

		if addressStr != "" {
			address := common.HexToAddress(addressStr)
			availableRiskScores[chainID][address] = true
		}
	}

	// Convert to the required return type
	result := make(map[common.Address]bool)
	for address := range availableRiskScores[chainID] {
		result[address] = true
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
