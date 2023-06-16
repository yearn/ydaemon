package helpers

import (
	"time"

	"github.com/yearn/ydaemon/common/bigNumber"
)

func GetToday(ppsPerTime map[uint64]*bigNumber.Int, decimals uint64) *bigNumber.Float {
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

func GetLastWeek(ppsPerTime map[uint64]*bigNumber.Int, decimals uint64) *bigNumber.Float {
	now := time.Now()
	noonUTC := time.Date(now.Year(), now.Month(), now.Day(), 12, 0, 0, 0, time.UTC)
	lastWeek := noonUTC.AddDate(0, 0, -7)
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

func GetLastMonth(ppsPerTime map[uint64]*bigNumber.Int, decimals uint64) *bigNumber.Float {
	now := time.Now()
	noonUTC := time.Date(now.Year(), now.Month(), now.Day(), 12, 0, 0, 0, time.UTC)
	lastMonth := noonUTC.AddDate(0, -1, 0)
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

func GetLastYear(ppsPerTime map[uint64]*bigNumber.Int, decimals uint64) *bigNumber.Float {
	now := time.Now()
	noonUTC := time.Date(now.Year(), now.Month(), now.Day(), 12, 0, 0, 0, time.UTC)
	lastYear := noonUTC.AddDate(-1, 0, 0)
	ppsYear := bigNumber.NewFloat(0)

	if data, ok := ppsPerTime[uint64(lastYear.Unix())]; ok {
		ppsYear = ToNormalizedAmount(data, decimals)
		return ppsYear
	}
	for i := 1; i < 365; i++ {
		dayToCheck := noonUTC.AddDate(0, 0, i-365)
		if data, ok := ppsPerTime[uint64(dayToCheck.Unix())]; ok {
			ppsYear = ToNormalizedAmount(data, decimals)
			break
		}
	}

	return ppsYear
}
