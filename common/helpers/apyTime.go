package helpers

import (
	"time"

	"github.com/yearn/ydaemon/common/bigNumber"
)

func GetPPSToday(ppsPerTime map[uint64]*bigNumber.Int, decimals uint64) *bigNumber.Float {
	now := time.Now()
	noonUTC := time.Date(now.Year(), now.Month(), now.Day(), 12, 0, 0, 0, time.UTC)
	ppsNow := bigNumber.NewFloat(0)

	if data, ok := ppsPerTime[uint64(noonUTC.Unix())]; ok {
		ppsNow = ToNormalizedAmount(data, decimals)
	} else if data, ok := ppsPerTime[uint64(noonUTC.AddDate(0, 0, -1).Unix())]; ok {
		ppsNow = ToNormalizedAmount(data, decimals)
	}

	return ppsNow
}

func GetPPSLastWeek(ppsPerTime map[uint64]*bigNumber.Int, decimals uint64) *bigNumber.Float {
	now := time.Now()
	noonUTC := time.Date(now.Year(), now.Month(), now.Day(), 12, 0, 0, 0, time.UTC)
	lastWeek := noonUTC.AddDate(0, 0, -7)
	if now.Before(noonUTC) {
		lastWeek = noonUTC.AddDate(0, 0, -8)
	}
	ppsWeek := bigNumber.NewFloat(0)

	if data, ok := ppsPerTime[uint64(lastWeek.Unix())]; ok {
		ppsWeek = ToNormalizedAmount(data, decimals)
		return ppsWeek
	}
	for i := 1; i < 7; i++ {
		dayToCheck := noonUTC.AddDate(0, 0, i-7)
		if data, ok := ppsPerTime[uint64(dayToCheck.Unix())]; ok {
			ppsWeek = ToNormalizedAmount(data, decimals)
			break
		}
	}

	return ppsWeek
}

func GetPPSLastMonth(ppsPerTime map[uint64]*bigNumber.Int, decimals uint64) *bigNumber.Float {
	now := time.Now()
	noonUTC := time.Date(now.Year(), now.Month(), now.Day(), 12, 0, 0, 0, time.UTC)
	lastMonth := noonUTC.AddDate(0, -1, 0)
	if now.Before(noonUTC) {
		lastMonth = noonUTC.AddDate(0, -1, -1)
	}
	ppsMonth := bigNumber.NewFloat(0)

	if data, ok := ppsPerTime[uint64(lastMonth.Unix())]; ok {
		ppsMonth = ToNormalizedAmount(data, decimals)
		return ppsMonth
	}
	for i := 1; i < 30; i++ {
		dayToCheck := noonUTC.AddDate(0, 0, i-30)
		if data, ok := ppsPerTime[uint64(dayToCheck.Unix())]; ok {
			ppsMonth = ToNormalizedAmount(data, decimals)
			break
		}
	}

	return ppsMonth
}

/************************************************************************************************
** GetAPR calculates the Annual Percentage Rate (APR) based on the final value (vf), initial
** value (vi), and the number of days (days). The formula used is:
** APR = ((vf - vi) / vi) / days * 365
************************************************************************************************/
func GetAPR(vf *bigNumber.Float, vi *bigNumber.Float, days *bigNumber.Float) *bigNumber.Float {
	apr := bigNumber.NewFloat(0).Sub(vf, vi)
	apr = bigNumber.NewFloat(0).Div(apr, vi)
	apr = bigNumber.NewFloat(0).Div(apr, days)
	apr = bigNumber.NewFloat(0).Mul(apr, bigNumber.NewFloat(365))

	return apr
}

/************************************************************************************************
** GetAPY calculates the Annual Percentage Yield (APY) based on the final value (vf), initial
** value (vi), and the number of days (days). The formula used is:
** APY = ((vf / vi) ^ (365 / days) - 1)
************************************************************************************************/
func GetAPY(vf *bigNumber.Float, vi *bigNumber.Float, days *bigNumber.Float) *bigNumber.Float {
	if vf == nil || vi == nil || days == nil || vi.IsZero() || days.IsZero() {
		return bigNumber.NewFloat(0)
	}

	ratio := bigNumber.NewFloat(0).Quo(vf, vi)
	exponent := bigNumber.NewFloat(0).Quo(bigNumber.NewFloat(365), days)
	exponentUint64, _ := exponent.Uint64()
	apy := bigNumber.NewFloat(0).Pow(ratio, exponentUint64)
	apy = bigNumber.NewFloat(0).Sub(apy, bigNumber.NewFloat(1))

	return apy
}
