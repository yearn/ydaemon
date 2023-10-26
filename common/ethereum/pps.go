package ethereum

import (
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/common/bigNumber"
	"github.com/yearn/ydaemon/common/contracts"
	"github.com/yearn/ydaemon/common/helpers"
	"github.com/yearn/ydaemon/common/logs"
)

func FetchPPSToday(
	chainID uint64,
	vaultAddress common.Address,
	decimals uint64,
) *bigNumber.Float {
	vaultContract, err := contracts.NewYearnVaultCaller(vaultAddress, GetRPC(chainID))
	if err != nil {
		logs.Error("Could not get vault contract for " + vaultAddress.Hex())
		return bigNumber.NewFloat(0)
	}
	pps, _ := vaultContract.PricePerShare(nil)
	ppsToday := helpers.ToNormalizedAmount(bigNumber.SetInt(pps), decimals)
	return ppsToday
}

func FetchPPSLastWeek(
	chainID uint64,
	vaultAddress common.Address,
	decimals uint64,
) *bigNumber.Float {
	vaultContract, err := contracts.NewYearnVaultCaller(vaultAddress, GetRPC(chainID))
	if err != nil {
		logs.Error("Could not get vault contract for " + vaultAddress.Hex())
		return bigNumber.NewFloat(0)
	}
	blocksPerDay := 7150
	estBlockLastWeek := blocksPerDay * 7
	opts := &bind.CallOpts{
		BlockNumber: big.NewInt(int64(estBlockLastWeek)),
	}
	pps, _ := vaultContract.PricePerShare(opts)
	ppsToday := helpers.ToNormalizedAmount(bigNumber.SetInt(pps), decimals)
	return ppsToday
}

func FetchPPSLastMonth(
	chainID uint64,
	vaultAddress common.Address,
	decimals uint64,
) *bigNumber.Float {
	vaultContract, err := contracts.NewYearnVaultCaller(vaultAddress, GetRPC(chainID))
	if err != nil {
		logs.Error("Could not get vault contract for " + vaultAddress.Hex())
		return bigNumber.NewFloat(0)
	}
	blocksPerDay := 7150
	estBlockLastWeek := blocksPerDay * 30
	opts := &bind.CallOpts{
		BlockNumber: big.NewInt(int64(estBlockLastWeek)),
	}
	pps, _ := vaultContract.PricePerShare(opts)
	ppsToday := helpers.ToNormalizedAmount(bigNumber.SetInt(pps), decimals)
	return ppsToday
}

func GetLastYear(ppsPerTime map[uint64]*bigNumber.Int, decimals uint64) *bigNumber.Float {
	now := time.Now()
	noonUTC := time.Date(now.Year(), now.Month(), now.Day(), 12, 0, 0, 0, time.UTC)
	lastYear := noonUTC.AddDate(-1, 0, 0)
	if now.Before(noonUTC) {
		lastYear = noonUTC.AddDate(-1, 0, -1)
	}
	ppsYear := bigNumber.NewFloat(0)

	if data, ok := ppsPerTime[uint64(lastYear.Unix())]; ok {
		ppsYear = helpers.ToNormalizedAmount(data, decimals)
		return ppsYear
	}
	for i := 1; i < 365; i++ {
		dayToCheck := noonUTC.AddDate(0, 0, i-365)
		if data, ok := ppsPerTime[uint64(dayToCheck.Unix())]; ok {
			ppsYear = helpers.ToNormalizedAmount(data, decimals)
			break
		}
	}

	return ppsYear
}
