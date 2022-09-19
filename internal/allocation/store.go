package allocation

import "github.com/ethereum/go-ethereum/common"

type TStore struct {
	// StrategyGroupAllocation holds the TVL for strategies in a risk group
	StrategyGroupAllocation map[uint64]map[common.Address]*TStrategyAllocation
}

// Store holds the data for the partners for each chain.
var Store = TStore{}

func init() {
	Store.StrategyGroupAllocation = make(map[uint64]map[common.Address]*TStrategyAllocation)
}
