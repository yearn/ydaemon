package apy

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/common/bigNumber"
	"github.com/yearn/ydaemon/common/helpers"
	"github.com/yearn/ydaemon/common/logs"
	"github.com/yearn/ydaemon/internal/events"
	"github.com/yearn/ydaemon/internal/indexer"
	priceAccessor "github.com/yearn/ydaemon/internal/prices"
	"github.com/yearn/ydaemon/internal/strategies"
	"github.com/yearn/ydaemon/internal/tokens"
	"github.com/yearn/ydaemon/internal/vaults"
	"github.com/yearn/ydaemon/processes/prices"
)

/**************************************************************************************************
** getTokenPrice is an helper function which will try to retrieve the price of a token from the
** prices package. If the token is not found, it will return 0.
** The price is returned in base 2 (humanized)
**************************************************************************************************/
func getTokenPrice(chainID uint64, tokenAddr common.Address) *bigNumber.Float {
	price, ok := priceAccessor.FindPrice(chainID, tokenAddr)
	if !ok {
		logs.Warning("Could not find price for token "+tokenAddr.Hex()+" on chain", chainID)
		price = bigNumber.NewInt(0)
	}
	return helpers.ToNormalizedAmount(price, 6)
}

/**************************************************************************************************
** initYearnEcosystem is used to initialize the yearn ecosystem data, aka fetching all the vaults,
** strategies, tokens, prices, etc.
** Based on that, we have everything ready to compute the fees for each partner.
**************************************************************************************************/
func initYearnEcosystem(chainID uint64) {
	registries := indexer.IndexNewVaults(chainID)
	vaultMap := vaults.RetrieveAllVaults(chainID, registries)
	strategiesSlice := indexer.IndexNewStrategies(chainID, vaultMap)
	tokens.RetrieveAllTokens(chainID, registries)

	prices.Run(chainID)
	logs.Info(`loading staking pools...`)
	events.HandleStakingPoolAdded(chainID, 0, nil)
	logs.Info(`loading strategies...`)
	strategies.RetrieveAllStrategies(chainID, strategiesSlice)
	indexer.PostProcess(chainID)
}
