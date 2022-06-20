package daemons

import (
	"github.com/majorfi/ydaemon/internal/ethereum"
)

var multicallClientForChainID = make(map[uint64]ethereum.TEthMultiCaller)

// init is fired directly on app start and prepare the multicall clients
func init() {
	multicallClientForChainID[1] = ethereum.NewMulticall(
		ethereum.GetRPCURI(1),
		ethereum.GetMulticallAddress(1),
	)
	multicallClientForChainID[250] = ethereum.NewMulticall(
		ethereum.GetRPCURI(250),
		ethereum.GetMulticallAddress(250),
	)
	multicallClientForChainID[42161] = ethereum.NewMulticall(
		ethereum.GetRPCURI(42161),
		ethereum.GetMulticallAddress(42161),
	)
}
