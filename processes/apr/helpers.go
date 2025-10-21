package apr

import (
	"math"
	"strconv"

	"github.com/yearn/ydaemon/common/bigNumber"
	"github.com/yearn/ydaemon/common/logs"
)

/**************************************************************************************************
** Convert APR to APY using the formula: APY = (1 + APR/n)^n - 1
** where n is the number of compounding periods per year.
**************************************************************************************************/
func convertAPRToAPY(apr *bigNumber.Float, compoundingPeriods *bigNumber.Float) *bigNumber.Float {
	if apr == nil || compoundingPeriods == nil || compoundingPeriods.IsZero() {
		return bigNumber.NewFloat(0)
	}

	// Calculate (1 + APR/n)
	onePlusAPRDivN := bigNumber.NewFloat(0).Div(apr, compoundingPeriods)
	onePlusAPRDivN = bigNumber.NewFloat(0).Add(onePlusAPRDivN, bigNumber.NewFloat(1))

	// Calculate (1 + APR/n)^n
	compoundingPeriodsUint64, _ := compoundingPeriods.Uint64()
	apy := bigNumber.NewFloat(0).Pow(onePlusAPRDivN, compoundingPeriodsUint64)

	// Subtract 1 to get the final APY
	apy = bigNumber.NewFloat(0).Sub(apy, bigNumber.NewFloat(1))

	return apy
}

func convertFloatAPRToAPY(apr float64, periodsPerYear float64) float64 {

	// Convert APR to decimal form
	aprDecimal := apr / 100.0

	// APY = (1 + r/n)^n - 1
	// where r is the APR in decimal form and n is the number of compounding periods
	apy := math.Pow(1+(aprDecimal/periodsPerYear), periodsPerYear) - 1

	// Convert back to percentage
	return apy * 100
}

/**************************************************************************************************
** parseKongFloatAPY safely parses Kong float APY values (weeklyNet, monthlyNet, inceptionNet)
** to bigNumber.Float. Returns zero on nil or error.
**************************************************************************************************/
func parseKongFloatAPY(value *float64) *bigNumber.Float {
	if value == nil {
		return bigNumber.NewFloat(0)
	}
	return bigNumber.NewFloat(*value)
}

/**************************************************************************************************
** parseKongStringPPS safely parses Kong string PPS values (pricePerShare, weeklyPricePerShare,
** monthlyPricePerShare) to bigNumber.Float, normalizing by decimals. Returns zero on empty string or error.
**************************************************************************************************/
func parseKongStringPPS(value string, decimals uint64) *bigNumber.Float {
	if value == "" {
		return bigNumber.NewFloat(0)
	}

	// Parse as integer first (Kong returns BigInt as string)
	intVal, err := strconv.ParseUint(value, 10, 64)
	if err != nil {
		logs.Warning("Failed to parse Kong PPS value '%s': %v", value, err)
		return bigNumber.NewFloat(0)
	}

	// Normalize by decimals: divide by 10^decimals
	divisor := 1.0
	for i := uint64(0); i < decimals; i++ {
		divisor *= 10.0
	}
	floatVal := float64(intVal) / divisor

	return bigNumber.NewFloat(floatVal)
}
