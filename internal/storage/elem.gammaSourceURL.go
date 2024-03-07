package storage

import (
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/go-co-op/gocron"
	"github.com/yearn/ydaemon/common/logs"
)

type TGammaSourceURL struct {
	Vault  string
	Source string
}

var _gammaSourceURLSyncMap = make(map[uint64][]TGammaSourceURL)

/**************************************************************************************************
** StoreGammaSourceURL
**************************************************************************************************/
func StoreGammaSourceURL(chainID uint64, gauges []TGammaSourceURL) {
	if _gammaSourceURLSyncMap[chainID] == nil {
		_gammaSourceURLSyncMap[chainID] = []TGammaSourceURL{}
	}
	_gammaSourceURLSyncMap[chainID] = gauges
}

/**************************************************************************************************
** ListGammaSourceURL
**************************************************************************************************/
func ListGammaSourceURL(chainID uint64) (asSlice []TGammaSourceURL) {
	return _gammaSourceURLSyncMap[chainID]
}

/**************************************************************************************************
** GetGammaSourceURL
**************************************************************************************************/
func GetGammaSourceURL(chainID uint64, vaultAddress common.Address) (string, bool) {
	for _, src := range _gammaSourceURLSyncMap[chainID] {
		if common.HexToAddress(src.Vault).Hex() == vaultAddress.Hex() {
			return src.Source, true
		}
	}
	return ``, false
}

func FetchGammaSourceURL() {
	resp, err := http.Get(`https://raw.githubusercontent.com/mil0xeth/yearn-v3-factory-helper/main/gamma-polygon.json`)
	if err != nil {
		logs.Error(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		logs.Error(err)
	}
	var gammaSourceURL []TGammaSourceURL
	if err := json.Unmarshal(body, &gammaSourceURL); err != nil {
		logs.Error(err)
	}
	StoreGammaSourceURL(137, gammaSourceURL)
}

func init() {
	cron := gocron.NewScheduler(time.UTC)
	cron.Every(30).Minute().Do(FetchGammaSourceURL)
	cron.StartAsync()
}
