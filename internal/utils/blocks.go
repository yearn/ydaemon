package utils

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/common/bigNumber"
)

type TEventBlock struct {
	EventType   string
	TxHash      common.Hash
	BlockNumber uint64
	TxIndex     uint
	LogIndex    uint
	Value       *bigNumber.Int
}
