package vaultPricePerShare

import (
	"encoding/json"
	"io/ioutil"
	"math"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/yearn/ydaemon/common/env"
	"github.com/yearn/ydaemon/common/ethereum"
	"github.com/yearn/ydaemon/common/logs"
	"github.com/yearn/ydaemon/common/store"
)

var _dailyBlockNumber = make(map[uint64]*sync.Map)

type TLlamaBlock struct {
	Height    uint64 `json:"height"`
	Timestamp uint64 `json:"timestamp"`
}

func chainIDToName(chainID uint64) string {
	switch chainID {
	case 1:
		return "ethereum"
	case 10:
		return "optimism"
	case 100:
		return "gnosis"
	case 137:
		return "polygon"
	case 250:
		return "fantom"
	case 42161:
		return "arbitrum"
	case 43114:
		return "avalanche"
	default:
		return "Unknown"
	}
}

func Run(chainID uint64) {
	assertDailyBlockNumber(chainID)
}

func assertDailyBlockNumber(chainID uint64) {
	lastItem := store.DBVaultPricePerShare{}
	syncMap := _dailyBlockNumber[chainID]
	if syncMap == nil {
		_dailyBlockNumber[chainID] = &sync.Map{}
		syncMap = _dailyBlockNumber[chainID]
	}
	/**********************************************************************************************
	** The first step is to find the earliest relevant deployment for yearn on this chain. For us
	** this means the earliest block where we have a registry contract deployed.
	**********************************************************************************************/
	earliestBlock := uint64(lastItem.Block)
	if (earliestBlock == 0) || (earliestBlock == math.MaxUint64) {
		earliestBlock = uint64(math.MaxUint64)
		for _, registry := range env.YEARN_REGISTRIES[chainID] {
			if registry.Block < earliestBlock {
				earliestBlock = registry.Block
			}
		}
	}

	/**********************************************************************************************
	** Then we will play with noon UTC time, starting from the earliest block. We will loop every
	** 24 hours until we reach now. For each day, we will get the block number at noon UTC and
	** store it in the dailyBlockNumber map.
	**********************************************************************************************/
	timestamp := ethereum.GetBlockTime(chainID, earliestBlock)
	timestampAsTime := time.Unix(int64(timestamp), 0)
	noonUTC := time.Date(timestampAsTime.Year(), timestampAsTime.Month(), timestampAsTime.Day(), 12, 0, 0, 0, time.UTC)
	now := time.Now()

	/**********************************************************************************************
	** And here we loop, first checking if we already have the block number for this day, but if
	** not, we will fetch it from DeFiLlama. If we get it, we will store it in the dailyBlockNumber
	** map.
	**********************************************************************************************/
	for noonUTC.Before(now) {
		nextDayNoonUTCTimestamp := noonUTC.Unix()
		data, ok := store.GetTimeBlock(chainID, uint64(nextDayNoonUTCTimestamp))
		if ok {
			syncMap.Store(noonUTC, data)
			noonUTC = noonUTC.AddDate(0, 0, 1)
			continue
		} else {
			defillamaURI := `https://coins.llama.fi/block/` + chainIDToName(chainID) + `/` + strconv.FormatInt(nextDayNoonUTCTimestamp, 10)
			resp, err := http.Get(defillamaURI)
			if err != nil || resp.StatusCode != 200 {
				logs.Warning("Error fetching timestamp from DeFiLlama for chain", chainID)
				logs.Error(err)
				continue
			}
			defer resp.Body.Close()
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				logs.Warning("Error unmarshalling response body from DeFiLlama for chain", chainID)
				continue
			}
			data := TLlamaBlock{}
			if err := json.Unmarshal(body, &data); err != nil {
				logs.Warning("Error unmarshalling response body from DeFiLlama for chain", chainID)
				continue
			}
			store.StoreBlockTime(chainID, data.Height, uint64(nextDayNoonUTCTimestamp))
			syncMap.Store(noonUTC, data.Height)
			noonUTC = noonUTC.AddDate(0, 0, 1)
			time.Sleep(500 * time.Millisecond)
		}
	}
}
