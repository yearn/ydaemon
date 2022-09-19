package strategies

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/internal/utils/models"
)

type TStore struct {
	// StrategyList contains the list of strategies we will need to multicall more info.
	StrategyList map[uint64]map[common.Address]models.TStrategyList

	// StrategiesFromRisk holds the data for the strategies from the Yearn Risk Framework for each chain.
	StrategiesFromRisk map[uint64]map[common.Address]models.TStrategyFromRisk

	// StrategyMultiCallData holds the details about the strategies based on a multicall
	StrategyMultiCallData map[uint64]map[common.Address]models.TStrategyMultiCallData

	// WithdrawalQueueMultiCallData holds the details about the withdrawalQueue order based on a multicall
	WithdrawalQueueMultiCallData map[uint64]map[common.Address]int64

	// StrategyGroupFromRisk holds the risk groups from the risk framework
	StrategyGroupFromRisk map[uint64][]*TStrategyGroupFromRisk
}

// Store holds the data for the partners for each chain.
var Store = TStore{}

func init() {
	Store.StrategyList = make(map[uint64]map[common.Address]models.TStrategyList)
	Store.StrategiesFromRisk = make(map[uint64]map[common.Address]models.TStrategyFromRisk)
	Store.StrategyMultiCallData = make(map[uint64]map[common.Address]models.TStrategyMultiCallData)
	Store.WithdrawalQueueMultiCallData = make(map[uint64]map[common.Address]int64)
	Store.StrategyGroupFromRisk = make(map[uint64][]*TStrategyGroupFromRisk)
}
