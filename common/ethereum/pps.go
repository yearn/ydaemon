package ethereum

import (
	"math/big"

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
	estBlockLastWeek := GetBlockNumberXDaysAgo(chainID, 7)
	opts := &bind.CallOpts{
		BlockNumber: big.NewInt(int64(estBlockLastWeek)),
	}
	pps, _ := vaultContract.PricePerShare(opts)
	ppsLastWeek := helpers.ToNormalizedAmount(bigNumber.SetInt(pps), decimals)
	return ppsLastWeek
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
	estBlockLastWeek := GetBlockNumberXDaysAgo(chainID, 30)
	opts := &bind.CallOpts{
		BlockNumber: big.NewInt(int64(estBlockLastWeek)),
	}
	pps, _ := vaultContract.PricePerShare(opts)
	ppsToday := helpers.ToNormalizedAmount(bigNumber.SetInt(pps), decimals)
	return ppsToday
}
