package bribes

import (
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/common/bigNumber"
	"github.com/yearn/ydaemon/common/types/common"
)

// TEventAdded contains the rewardAdded event data for the yBribeV3 contract
type TEventAdded struct {
	Amount      *bigNumber.Int `json:"amount"`
	Briber      common.Address `json:"briber"`
	Gauge       common.Address `json:"gauge"`
	RewardToken common.Address `json:"rewardToken"`
	TxHash      ethcommon.Hash `json:"txHash"`
	Timestamp   uint64         `json:"timestamp"`
	BlockNumber uint64         `json:"blockNumber"`
	TxIndex     uint           `json:"-"`
	LogIndex    uint           `json:"-"`
}

/**********************************************************************************************
** Set of functions to store and retrieve the tokens from the cache and/or database and being
** able to access them from the rest of the application.
** The _yBribeRewardAdded variable is not exported and is only used internally by the
** functions below.
**********************************************************************************************/
var _yBribeRewardAdded = make(map[uint64]map[uint64]*TEventAdded)

/**********************************************************************************************
** ListRewardAdded will, for a given chainID, return the list of all the events stored in
** _yBribeRewardAdded.
**********************************************************************************************/
func ListRewardAdded(chainID uint64) []*TEventAdded {
	var events []*TEventAdded
	for _, event := range _yBribeRewardAdded[chainID] {
		events = append(events, event)
	}
	return events
}

/**********************************************************************************************
** SetInRewardAddedMap will allow us to safely populate the _yBribeRewardAdded variable.
**********************************************************************************************/
func SetInRewardAddedMap(chainID uint64, blockNumber uint64, event *TEventAdded) {
	if _, ok := _yBribeRewardAdded[chainID]; !ok {
		_yBribeRewardAdded[chainID] = make(map[uint64]*TEventAdded)
	}
	_yBribeRewardAdded[chainID][blockNumber] = event
}
