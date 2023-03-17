package partnerTracker

import (
	"github.com/yearn/ydaemon/common/types/common"
)

type TContractData struct {
	Address common.Address
	Block   uint64
}

// PARTNER_TRACKERS_ADDRESSES contains the address of the yBribe contract
var PARTNER_TRACKERS_ADDRESSES = map[uint64]TContractData{
	1: {
		Address: common.HexToAddress(`0x8ee392a4787397126C163Cb9844d7c447da419D8`),
		Block:   uint64(14166636),
	},
	10: {
		Address: common.HexToAddress(`0x7E08735690028cdF3D81e7165493F1C34065AbA2`),
		Block:   uint64(29675215),
	},
	250: {
		Address: common.HexToAddress(`0x086865B2983320b36C42E48086DaDc786c9Ac73B`),
		Block:   uint64(40499061),
	},
	42161: {
		Address: common.HexToAddress(`0x0e5b46E4b2a05fd53F5a4cD974eb98a9a613bcb7`),
		Block:   uint64(30385403),
	},
}

var KnownRefferers = map[uint64]map[common.Address]string{
	1: {
		common.HexToAddress(`0x558247e365be655f9144e1a0140D793984372Ef3`): `Ledger`,
		common.HexToAddress(`0xFEB4acf3df3cDEA7399794D0869ef76A6EfAff52`): `Yearn.finance`,
		common.HexToAddress(`0x3CE37278de6388532C3949ce4e886F365B14fB56`): `Zapper?`,
		common.HexToAddress(`0xFBD4C3D8bE6B15b7cf428Db2838bb44C0054fCd2`): `Portals?`,
	},
	10:    {},
	250:   {},
	42161: {},
}
