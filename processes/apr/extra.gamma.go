package apr

import (
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/common/addresses"
	"github.com/yearn/ydaemon/common/bigNumber"
)

/**************************************************************************
** This struct represents the relevant response from the Gamma Merkl API.
**************************************************************************/
type TGammaMerklAPIResp struct {
	Pools map[string]struct {
		ALM map[string]struct {
			ALMAPR     float64 `json:"almAPR"`
			ALMAddress string  `json:"almAddress"`
		} `json:"alm"`
	} `json:"pools"`
}

/**************************************************************************
** As the API response is quite large, we cache it to avoid unnecessary
** requests.
** This will be refreshed everytime we globaly refresh the APR system
**************************************************************************/
var cachedGammaMerkl map[uint64]map[string]TGammaMerklAPIResp

func init() {
	cachedGammaMerkl = make(map[uint64]map[string]TGammaMerklAPIResp)
}
func refreshGammaMerkl(chainID uint64) {
	updatedValues, ok := retrieveGammaMerklData(chainID)
	if !ok {
		return
	}
	cachedGammaMerkl[chainID] = updatedValues
}

/**************************************************************************
** To "calculate" the Gamma Extra Rewards, we will just retrieve the
** relevant information from the Gamma Merkl API cached response.
** We need to make sure we have it cached (or we will retrieve it) and that
** we have relevant information for the given network.
**
** Once this is done, we are looking for the ALLMAPR for the given token.
**************************************************************************/
func calculateGammaExtraRewards(chainID uint64, tokenAddress common.Address) (*bigNumber.Float, bool) {
	chainIDStr := strconv.FormatUint(chainID, 10)
	if _, ok := cachedGammaMerkl[chainID]; !ok {
		refreshGammaMerkl(chainID)
	}
	if _, ok := cachedGammaMerkl[chainID]; !ok {
		return bigNumber.NewFloat(0), false
	}

	gammaMerkl := cachedGammaMerkl[chainID]
	if _, ok := gammaMerkl[chainIDStr]; !ok {
		return bigNumber.NewFloat(0), false
	}

	for _, pool := range gammaMerkl[chainIDStr].Pools {
		for lpTokenAddress, alm := range pool.ALM {
			if addresses.Equals(tokenAddress, lpTokenAddress) {
				return bigNumber.NewFloat(0).Div(bigNumber.NewFloat(alm.ALMAPR), bigNumber.NewFloat(100)), true
			}
		}

	}

	return bigNumber.NewFloat(0), false
}
