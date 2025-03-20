package apr

import (
	"math"

	"github.com/yearn/ydaemon/common/bigNumber"
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
