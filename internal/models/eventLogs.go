package models

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/common/bigNumber"
)

// TEventRewardAdded contains the rewardAdded event data for the yBribeV3 contract
type TEventRewardAdded struct {
	Amount      *bigNumber.Int `json:"amount"`
	Briber      common.Address `json:"briber"`
	Gauge       common.Address `json:"gauge"`
	RewardToken common.Address `json:"rewardToken"`
	TxHash      common.Hash    `json:"txHash"`
	Timestamp   uint64         `json:"timestamp"`
	BlockNumber uint64         `json:"blockNumber"`
	TxIndex     uint           `json:"-"`
	LogIndex    uint           `json:"-"`
}

// TEventReferredBalanceIncreased contains the ReferredBalanceIncreased event data
// for the partner tracker contract
type TEventReferredBalanceIncreased struct {
	Amount         *bigNumber.Int `json:"amount"`
	TotalDeposited *bigNumber.Int `json:"totalDeposited"`
	PartnerID      common.Address `json:"partnerID"`
	Vault          common.Address `json:"vault"`
	Depositer      common.Address `json:"depositer"`
	TxHash         common.Hash    `json:"txHash"`
	BlockNumber    uint64         `json:"blockNumber"`
	TxIndex        uint           `json:"-"`
	LogIndex       uint           `json:"-"`
}
