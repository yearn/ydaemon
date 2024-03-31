package apr

import (
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/common/addresses"
	"github.com/yearn/ydaemon/common/bigNumber"
	"github.com/yearn/ydaemon/internal/storage"
)

/**************************************************************************************************
** To "calculate" the Gamma Extra Rewards, we will just retrieve the relevant information from the
** Gamma Merkl API cached response. We need to make sure we have it cached (or we will retrieve it)
** and that we have relevant information for the given network.
**
** Once this is done, we are looking for the ALLMAPR for the given token.
**************************************************************************************************/
func calculateGammaExtraRewards(chainID uint64, tokenAddress common.Address) (*bigNumber.Float, bool) {
	chainIDStr := strconv.FormatUint(chainID, 10)
	if _, ok := storage.GetCachedGammaMerkl(chainID); !ok {
		storage.RefreshGammaCalls(chainID)
	}
	if _, ok := storage.GetCachedGammaMerkl(chainID); !ok {
		return bigNumber.NewFloat(0), false
	}

	gammaMerkl, _ := storage.GetCachedGammaMerkl(chainID)
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
