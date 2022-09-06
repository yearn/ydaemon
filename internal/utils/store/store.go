package store

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/internal/utils/models"
)

// TokenList contains the list of tokens we will need to fetch prices for.
// List is accessible through TokenList[chainID]
var TokenList = make(map[uint64][]common.Address)

// Tokens contains the list of address for a specific chain and some minimal information about the token
// List is accessible through Tokens[chainID]
var Tokens = make(map[uint64]map[common.Address]models.TERC20Token)

// StrategyList contains the list of strategies we will need to multicall more info.
// List is accessible through StrategyList[chainID]
var StrategyList = make(map[uint64]map[common.Address]models.TStrategyList)

// TokenPrices contains, for each network, the prices for each token we fetched
// the prices for. Price is accessible through TokenPrices[chainID][tokenAddress]
var TokenPrices = make(map[uint64]map[common.Address]*big.Int)

// StrategiesFromRisk holds the data for the strategies from the Yearn Risk Framework for each chain.
// The data is updated every 1 hour.
var StrategiesFromRisk = make(map[uint64]map[common.Address]models.TStrategyFromRisk)

// VaultsFromAPIV1 holds the data about the Vaults from the Yearn API for each chain.
// The data is updated every 10 minutes.
var VaultsFromAPIV1 = make(map[uint64]map[common.Address]models.TAPIV1Vault)

// StrategyMultiCallData holds the details about the strategies based on a multicall
// The data is updated every 10 minutes
var StrategyMultiCallData = make(map[uint64]map[common.Address]models.TStrategyMultiCallData)
