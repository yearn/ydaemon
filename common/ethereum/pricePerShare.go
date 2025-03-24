package ethereum

import (
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/common/bigNumber"
	"github.com/yearn/ydaemon/common/contracts"
	"github.com/yearn/ydaemon/common/helpers"
	"github.com/yearn/ydaemon/common/logs"
)

/**************************************************************************************************
** Price Per Share (PPS) is a core mechanism in Yearn's vault system that represents the value of
** each share in a vault. This file contains functions for calculating PPS values and APY from
** existing datasets, complementing the blockchain fetching functions in pps.go.
**
** This file primarily focuses on:
** 1. Working with pre-fetched time-series PPS data
** 2. Calculating APY values based on PPS changes over time
** 3. Providing helper functions for typical time periods (weekly, monthly)
**************************************************************************************************/

/**************************************************************************************************
** FetchPPSToday retrieves the current price per share (PPS) for a Yearn vault at the latest block.
**
** This function connects to the specified vault contract on the given chain and calls the
** pricePerShare method to retrieve the current value. The result is normalized according to
** the vault's decimals to return a proper floating point value.
**
** @param chainID The ID of the blockchain where the vault exists
** @param vaultAddress The Ethereum address of the vault contract
** @param decimals The number of decimals used by the vault token
** @return *bigNumber.Float The normalized price per share as a floating point number
**************************************************************************************************/
func FetchPPSToday(
	chainID uint64,
	vaultAddress common.Address,
	vaultActivation uint64,
	decimals uint64,
) *bigNumber.Float {
	vaultContract, err := contracts.NewYearnVaultCaller(vaultAddress, GetRPC(chainID))
	if err != nil {
		logs.Error(fmt.Sprintf("Chain %d - Could not get vault contract for %s: %v", chainID, vaultAddress.Hex(), err))
		return bigNumber.NewFloat(0)
	}

	pps, err := vaultContract.PricePerShare(nil)
	if err != nil {
		logs.Error(fmt.Sprintf("Chain %d - Failed to fetch current PPS for vault %s: %v", chainID, vaultAddress.Hex(), err))
		return bigNumber.NewFloat(0)
	}

	ppsToday := helpers.ToNormalizedAmount(bigNumber.SetInt(pps), decimals)
	return ppsToday
}

/**************************************************************************************************
** FetchPPSLastWeek retrieves the price per share (PPS) for a Yearn vault from approximately
** 7 days ago, using historical block number data.
**
** This function uses GetBlockNumberByPeriod to determine the block number from 7 days ago,
** then queries the vault contract at that specific block to get the historical PPS value.
** The result is normalized according to the vault's decimals.
**
** @param chainID The ID of the blockchain where the vault exists
** @param vaultAddress The Ethereum address of the vault contract
** @param decimals The number of decimals used by the vault token
** @return *bigNumber.Float The normalized historical price per share as a floating point number
**************************************************************************************************/
func FetchPPSLastWeek(
	chainID uint64,
	vaultAddress common.Address,
	vaultDeploymentBlock uint64,
	decimals uint64,
) *bigNumber.Float {
	vaultContract, err := contracts.NewYearnVaultCaller(vaultAddress, GetRPC(chainID))
	if err != nil {
		logs.Error(fmt.Sprintf("Chain %d - Could not get vault contract for %s: %v", chainID, vaultAddress.Hex(), err))
		return bigNumber.NewFloat(0)
	}

	estBlockLastWeek := GetBlockNumberByPeriod(chainID, 7)
	if estBlockLastWeek == 0 {
		logs.Warning(fmt.Sprintf("Chain %d - Failed to get block number from 7 days ago for vault %s", chainID, vaultAddress.Hex()))
		return bigNumber.NewFloat(0)
	}

	if vaultDeploymentBlock > 0 && estBlockLastWeek < vaultDeploymentBlock {
		// logs.Warning(`Chain ` + strconv.Itoa(int(chainID)) + ` - Vault ` + vaultAddress.Hex() + ` was not deployed last week, using inception PPS`)
		return bigNumber.NewFloat(1)
	}

	opts := &bind.CallOpts{
		BlockNumber: big.NewInt(int64(estBlockLastWeek)),
	}

	pps, err := vaultContract.PricePerShare(opts)
	if err != nil {
		logs.Error(fmt.Sprintf("Chain %d - Failed to fetch PPS from 7 days ago (block %d) for vault %s: %v",
			chainID, estBlockLastWeek, vaultAddress.Hex(), err))
		return bigNumber.NewFloat(0)
	}

	ppsLastWeek := helpers.ToNormalizedAmount(bigNumber.SetInt(pps), decimals)
	return ppsLastWeek
}

/**************************************************************************************************
** FetchPPSLastMonth retrieves the price per share (PPS) for a Yearn vault from approximately
** 30 days ago, using historical block number data.
**
** This function uses GetBlockNumberByPeriod to determine the block number from 30 days ago,
** then queries the vault contract at that specific block to get the historical PPS value.
** The result is normalized according to the vault's decimals.
**
** @param chainID The ID of the blockchain where the vault exists
** @param vaultAddress The Ethereum address of the vault contract
** @param decimals The number of decimals used by the vault token
** @return *bigNumber.Float The normalized historical price per share as a floating point number
**************************************************************************************************/
func FetchPPSLastMonth(
	chainID uint64,
	vaultAddress common.Address,
	vaultDeploymentBlock uint64,
	decimals uint64,
) *bigNumber.Float {
	vaultContract, err := contracts.NewYearnVaultCaller(vaultAddress, GetRPC(chainID))
	if err != nil {
		logs.Error(fmt.Sprintf("Chain %d - Could not get vault contract for %s: %v", chainID, vaultAddress.Hex(), err))
		return bigNumber.NewFloat(0)
	}

	estBlockLastMonth := GetBlockNumberByPeriod(chainID, 30)
	if estBlockLastMonth == 0 {
		logs.Warning(fmt.Sprintf("Chain %d - Failed to get block number from 30 days ago for vault %s", chainID, vaultAddress.Hex()))
		return bigNumber.NewFloat(0)
	}

	if vaultDeploymentBlock > 0 && estBlockLastMonth < vaultDeploymentBlock {
		// logs.Warning(`Chain ` + strconv.Itoa(int(chainID)) + ` - Vault ` + vaultAddress.Hex() + ` was not deployed last month, using inception PPS`)
		return bigNumber.NewFloat(1)
	}

	opts := &bind.CallOpts{
		BlockNumber: big.NewInt(int64(estBlockLastMonth)),
	}

	pps, err := vaultContract.PricePerShare(opts)
	if err != nil {
		logs.Error(fmt.Sprintf("Chain %d - Failed to fetch PPS from 30 days ago (block %d) for vault %s: %v",
			chainID, estBlockLastMonth, vaultAddress.Hex(), err))
		return bigNumber.NewFloat(0)
	}

	ppsLastMonth := helpers.ToNormalizedAmount(bigNumber.SetInt(pps), decimals)
	return ppsLastMonth
}

/**************************************************************************************************
** GetPPSToday retrieves the price per share (PPS) for today from a map of historical PPS values.
**
** This function first tries to get the PPS for today at noon UTC. If that is not available,
** it falls back to yesterday's noon value. This provides a consistent approach for getting
** the "current" PPS from historical data.
**
** @param ppsPerTime A map of timestamps to PPS values
** @param decimals The number of decimals used by the vault token
** @return *bigNumber.Float The normalized price per share as a floating point number
**************************************************************************************************/
func GetPPSToday(ppsPerTime map[uint64]*bigNumber.Int, decimals uint64) *bigNumber.Float {
	now := time.Now()
	noonUTC := time.Date(now.Year(), now.Month(), now.Day(), 12, 0, 0, 0, time.UTC)
	ppsNow := bigNumber.NewFloat(0)

	if data, ok := ppsPerTime[uint64(noonUTC.Unix())]; ok {
		ppsNow = helpers.ToNormalizedAmount(data, decimals)
	} else if data, ok := ppsPerTime[uint64(noonUTC.AddDate(0, 0, -1).Unix())]; ok {
		ppsNow = helpers.ToNormalizedAmount(data, decimals)
	}

	return ppsNow
}

/**************************************************************************************************
** GetPPSLastWeek retrieves the price per share (PPS) from approximately 7 days ago from
** a map of historical PPS values.
**
** This function first attempts to get the exact value from 7 days ago at noon UTC. If that
** specific timestamp is not available, it searches nearby days (up to 7 days before today)
** to find the closest available data point.
**
** @param ppsPerTime A map of timestamps to PPS values
** @param decimals The number of decimals used by the vault token
** @return *bigNumber.Float The normalized price per share as a floating point number
**************************************************************************************************/
func GetPPSLastWeek(ppsPerTime map[uint64]*bigNumber.Int, decimals uint64) *bigNumber.Float {
	now := time.Now()
	noonUTC := time.Date(now.Year(), now.Month(), now.Day(), 12, 0, 0, 0, time.UTC)
	lastWeek := noonUTC.AddDate(0, 0, -7)
	if now.Before(noonUTC) {
		lastWeek = noonUTC.AddDate(0, 0, -8)
	}
	ppsWeek := bigNumber.NewFloat(0)

	if data, ok := ppsPerTime[uint64(lastWeek.Unix())]; ok {
		ppsWeek = helpers.ToNormalizedAmount(data, decimals)
		return ppsWeek
	}
	for i := 1; i < 7; i++ {
		dayToCheck := noonUTC.AddDate(0, 0, i-7)
		if data, ok := ppsPerTime[uint64(dayToCheck.Unix())]; ok {
			ppsWeek = helpers.ToNormalizedAmount(data, decimals)
			break
		}
	}

	return ppsWeek
}

/**************************************************************************************************
** GetPPSLastMonth retrieves the price per share (PPS) from approximately 30 days ago from
** a map of historical PPS values.
**
** This function first attempts to get the exact value from 30 days ago at noon UTC. If that
** specific timestamp is not available, it searches nearby days (up to 30 days before today)
** to find the closest available data point.
**
** @param ppsPerTime A map of timestamps to PPS values
** @param decimals The number of decimals used by the vault token
** @return *bigNumber.Float The normalized price per share as a floating point number
**************************************************************************************************/
func GetPPSLastMonth(ppsPerTime map[uint64]*bigNumber.Int, decimals uint64) *bigNumber.Float {
	now := time.Now()
	noonUTC := time.Date(now.Year(), now.Month(), now.Day(), 12, 0, 0, 0, time.UTC)
	lastMonth := noonUTC.AddDate(0, -1, 0)
	if now.Before(noonUTC) {
		lastMonth = noonUTC.AddDate(0, -1, -1)
	}
	ppsMonth := bigNumber.NewFloat(0)

	if data, ok := ppsPerTime[uint64(lastMonth.Unix())]; ok {
		ppsMonth = helpers.ToNormalizedAmount(data, decimals)
		return ppsMonth
	}
	for i := 1; i < 30; i++ {
		dayToCheck := noonUTC.AddDate(0, 0, i-30)
		if data, ok := ppsPerTime[uint64(dayToCheck.Unix())]; ok {
			ppsMonth = helpers.ToNormalizedAmount(data, decimals)
			break
		}
	}

	return ppsMonth
}

/**************************************************************************************************
** CalculateAPY calculates the Annual Percentage Yield (APY) between two PPS values over a given
** number of days. This is a generalized function that can be used for weekly, monthly, or any
** custom time period APY calculations.
**
** The formula used is: APY = ((currentPPS - historicalPPS) / historicalPPS) / days * 365
**
** @param currentPPS The current price per share value
** @param historicalPPS The historical price per share value
** @param days The number of days between the two PPS values
** @return *bigNumber.Float The calculated APY as a decimal (e.g., 0.10 for 10%)
**************************************************************************************************/
func CalculateAPY(currentPPS, historicalPPS *bigNumber.Float, days int) *bigNumber.Float {
	// If either PPS is zero, or they're the same, APY is zero
	if historicalPPS.IsZero() || currentPPS.String() == historicalPPS.String() {
		return bigNumber.NewFloat(0)
	}

	// Calculate the raw change in PPS
	ppsChange := bigNumber.NewFloat(0).Sub(currentPPS, historicalPPS)

	// Calculate the percentage change (change / historicalPPS)
	percentageChange := bigNumber.NewFloat(0).Div(ppsChange, historicalPPS)

	// Annualize by dividing by days and multiplying by 365
	daysFloat := bigNumber.NewFloat(float64(days))
	annualizedChange := bigNumber.NewFloat(0).Div(percentageChange, daysFloat)
	apy := bigNumber.NewFloat(0).Mul(annualizedChange, bigNumber.NewFloat(365))

	return apy
}

/**************************************************************************************************
** CalculateWeeklyAPY calculates the annualized APY based on the price per share change over
** the past 7 days.
**
** @param currentPPS The current price per share value
** @param weekAgoPPS The price per share value from 7 days ago
** @return *bigNumber.Float The calculated weekly APY as a decimal
**************************************************************************************************/
func CalculateWeeklyAPY(currentPPS, weekAgoPPS *bigNumber.Float) *bigNumber.Float {
	return CalculateAPY(currentPPS, weekAgoPPS, 7)
}

/**************************************************************************************************
** CalculateMonthlyAPY calculates the annualized APY based on the price per share change over
** the past 30 days.
**
** @param currentPPS The current price per share value
** @param monthAgoPPS The price per share value from 30 days ago
** @return *bigNumber.Float The calculated monthly APY as a decimal
**************************************************************************************************/
func CalculateMonthlyAPY(currentPPS, monthAgoPPS *bigNumber.Float) *bigNumber.Float {
	return CalculateAPY(currentPPS, monthAgoPPS, 30)
}

/**************************************************************************************************
** CalculateYearlyAPY calculates the annualized APY based on the price per share change over
** the past 365 days.
**
** @param currentPPS The current price per share value
** @param yearlyPPS The price per share value from 365 days ago
** @return *bigNumber.Float The calculated yearly APY as a decimal
**************************************************************************************************/
func CalculateYearlyAPY(currentPPS, yearlyPPS *bigNumber.Float) *bigNumber.Float {
	return CalculateAPY(currentPPS, yearlyPPS, 365)
}
