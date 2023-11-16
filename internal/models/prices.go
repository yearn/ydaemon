package models

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/common/bigNumber"
)

type TPrices struct {
	Address        common.Address   `json:"address"`
	Price          *bigNumber.Int   `json:"price"`
	HumanizedPrice *bigNumber.Float `json:"humanizedPrice"`
	Source         string           `json:"source"`
}
