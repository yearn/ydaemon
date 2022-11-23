package vaults

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/internal/prices"
	"github.com/yearn/ydaemon/internal/strategies"
	"github.com/yearn/ydaemon/internal/tokens"
	"github.com/yearn/ydaemon/internal/utils"
)

func ProcessNewVault(chainID uint64, vaultsMap map[common.Address]utils.TVaultsFromRegistry) {
	tokens.RetrieveAllTokens(chainID, vaultsMap)
	prices.RetrieveAllPrices(chainID)
	RetrieveAllVaults(chainID, vaultsMap)
	strategiesAddedList := strategies.RetrieveAllStrategiesAdded(chainID, vaultsMap)
	strategies.RetrieveAllStrategies(chainID, strategiesAddedList)
	go strategies.IndexStrategyAdded(chainID, vaultsMap)
}
