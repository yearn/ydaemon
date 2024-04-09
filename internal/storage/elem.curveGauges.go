package storage

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/common/env"
	"github.com/yearn/ydaemon/internal/models"
)

var _curveGaugesSyncMap = make(map[uint64][]models.CurveGauge)

/**************************************************************************************************
** StoreCurveGauges
**************************************************************************************************/
func StoreCurveGauges(chainID uint64, gauges []models.CurveGauge) {
	if _curveGaugesSyncMap[chainID] == nil {
		_curveGaugesSyncMap[chainID] = []models.CurveGauge{}
	}
	_curveGaugesSyncMap[chainID] = gauges
}

/**************************************************************************************************
** ListGauges
**************************************************************************************************/
func ListGauges(chainID uint64) (asSlice []models.CurveGauge) {
	return _curveGaugesSyncMap[chainID]
}

/**************************************************************************************************
** GetGauge
**************************************************************************************************/
func GetGauge(chainID uint64, underlyingAddress common.Address) (models.CurveGauge, bool) {
	for _, gauge := range _curveGaugesSyncMap[chainID] {
		if common.HexToAddress(gauge.SwapToken).Hex() == underlyingAddress.Hex() {
			return gauge, true
		}
	}
	return models.CurveGauge{}, false
}

func FetchGauges() {
	for _, chain := range env.CHAINS {
		FetchCurveGauges(chain.ID)
	}
}
