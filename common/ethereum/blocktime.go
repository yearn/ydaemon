package ethereum

import (
	"context"
	"math/big"
	"os"
	"strconv"
	"time"

	"github.com/yearn/ydaemon/common/env"
	"github.com/yearn/ydaemon/common/helpers"
	"github.com/yearn/ydaemon/common/logs"
	"github.com/yearn/ydaemon/common/store"
)

var blockTimeMap = make(map[uint64]map[uint64]uint64)

func InitBlockTimestamp(chainID uint64) {
	type TScanResult struct {
		Status  string `json:"status"`
		Message string `json:"message"`
		Result  string `json:"result"`
	}
	var err error
	now := time.Now()
	noonUTC := time.Date(now.Year(), now.Month(), now.Day(), 12, 0, 0, 0, time.UTC)
	lastWeekTimestamp := noonUTC.AddDate(0, 0, -7).Unix()
	lastMonthTimestamp := noonUTC.AddDate(0, -1, 0).Unix()

	APIKey := os.Getenv("SCAN_API_KEY_FOR_" + strconv.FormatUint(chainID, 10))
	lastWeekBlock := helpers.FetchJSON[TScanResult](env.CHAINS[chainID].EtherscanURI + `?module=block&action=getblocknobytime&timestamp=` + strconv.FormatInt(lastWeekTimestamp, 10) + `&closest=before&apikey=` + APIKey)
	lastMonthBlock := helpers.FetchJSON[TScanResult](env.CHAINS[chainID].EtherscanURI + `?module=block&action=getblocknobytime&timestamp=` + strconv.FormatInt(lastMonthTimestamp, 10) + `&closest=before&apikey=` + APIKey)

	if blockTimeMap[chainID] == nil {
		blockTimeMap[chainID] = make(map[uint64]uint64)
	}
	if lastWeekBlock.Status == "1" {
		blockTimeMap[chainID][7], err = strconv.ParseUint(lastWeekBlock.Result, 10, 64)
		if err != nil {
			logs.Error(err)
			blockTimeMap[chainID][7] = uint64(env.CHAINS[chainID].AvgBlocksPerDay * 7)
		}
	} else {
		blockTimeMap[chainID][7] = uint64(env.CHAINS[chainID].AvgBlocksPerDay * 7)
	}

	if lastMonthBlock.Status == "1" {
		blockTimeMap[chainID][30], err = strconv.ParseUint(lastMonthBlock.Result, 10, 64)
		if err != nil {
			logs.Error(err)
			blockTimeMap[chainID][30] = uint64(env.CHAINS[chainID].AvgBlocksPerDay * 30)
		}
	} else {
		blockTimeMap[chainID][30] = uint64(env.CHAINS[chainID].AvgBlocksPerDay * 30)
	}
}

func GetBlockNumberXDaysAgo(chainID uint64, days uint64) uint64 {
	if (days != 7 && days != 30) || blockTimeMap[chainID] == nil {
		return 0
	}
	return blockTimeMap[chainID][days]
}

/**************************************************************************************************
** GetBlockTime will try, for a specific blockNumber on a specific chain, to find its execution
** timestamp. This timestamp should be available in the blockTimeSyncMap. If it is not, it will
** fetch it from the blockchain and store it in the blockTimeSyncMap.
**************************************************************************************************/
func GetBlockTime(chainID uint64, blockNumber uint64) (blockTime uint64) {
	blockTimeData, ok := store.GetBlockTime(chainID, blockNumber)
	if !ok {
		client := RPC[chainID]
		block, err := client.HeaderByNumber(context.Background(), big.NewInt(int64(blockNumber)))
		if err != nil {
			logs.
				Capture(`warn`, `impossible to retrieve block `+strconv.FormatUint(blockNumber, 10)+` on chain `+strconv.FormatUint(chainID, 10)).
				SetEntity(`block`).
				SetExtra(`error`, err.Error()).
				SetTag(`chainID`, strconv.FormatUint(chainID, 10)).
				SetTag(`rpcURI`, GetRPCURI(chainID)).
				SetTag(`blockNumber`, strconv.FormatUint(blockNumber, 10)).
				Send()
			return 0
		}
		store.StoreBlockTime(chainID, blockNumber, block.Time)
		return block.Time
	}
	return blockTimeData
}
