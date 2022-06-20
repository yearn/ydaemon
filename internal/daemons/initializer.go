package daemons

import (
	"github.com/majorfi/ydaemon/internal/ethereum"
	"github.com/majorfi/ydaemon/internal/logs"
)

var multicallClientForChainID = make(map[uint64]ethereum.TEthMultiCaller)

// init is fired directly on app start and prepare the multicall clients
func init() {
	logs.Info("Initializing lens Oracle")
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
	logs.Success("Lens Oracle initialized!")
}
