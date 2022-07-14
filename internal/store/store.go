package store

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/majorfi/ydaemon/internal/models"
)

// TokenList contains the list of tokens we will need to fetch prices for.
// List is accessible through TokenList[chainID]
var TokenList = make(map[uint64][]common.Address)

// StrategyList contains the list of strategies we will need to multicall more info.
// List is accessible through StrategyList[chainID]
var StrategyList = make(map[uint64]map[common.Address]models.TStrategyList)

// TokenPrices contains, for each network, the prices for each token we fetched
// the prices for. Price is accessible through TokenPrices[chainID][tokenAddress]
var TokenPrices = make(map[uint64]map[common.Address]*big.Int)

// VaultsFromMeta holds the data for the vaults from the Yearn Meta API for each chain.
// The data is updated every 15 seconds.
var VaultsFromMeta = make(map[uint64]map[string]models.TVaultFromMeta)

// TokensFromMeta holds the data for the tokens from the Yearn Meta API for each chain.
// The data is updated every 15 seconds.
var TokensFromMeta = make(map[uint64]map[string]models.TTokenFromMeta)

// StrategiesFromMeta holds the data for the strategies from the Yearn Meta API for each chain.
// The data is updated every 15 seconds.
var StrategiesFromMeta = make(map[uint64]map[string]models.TStrategyFromMeta)

// VaultsFromAPIV1 holds the data about the Vaults from the Yearn API for each chain.
// The data is updated every 15 seconds.
var VaultsFromAPIV1 = make(map[uint64]map[string]models.TAPIV1Vault)

// StrategyMultiCallData holds the details about the strategies based on a multicall
// The data is updated every 10 minutes
var StrategyMultiCallData = make(map[uint64]map[common.Address]models.TStrategyMultiCallData)
