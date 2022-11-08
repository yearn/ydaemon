package strategies

import (
	"github.com/yearn/ydaemon/common/types/common"
)

type TStore struct {
	// StrategiesFromRisk holds the data for the strategies from the Yearn Risk Framework for each chain.
	StrategiesFromRisk map[uint64]map[common.Address]TStrategyFromRisk

	// StrategyGroupFromRisk holds the risk groups from the risk framework
	StrategyGroupFromRisk map[uint64][]*TStrategyGroupFromRisk
}

// Store holds the data for the partners for each chain.
var Store = TStore{}

func init() {
	Store.StrategiesFromRisk = make(map[uint64]map[common.Address]TStrategyFromRisk)
	Store.StrategyGroupFromRisk = make(map[uint64][]*TStrategyGroupFromRisk)
}
