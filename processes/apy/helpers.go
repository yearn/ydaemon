package apy

import (
	"math"

	"github.com/davecgh/go-spew/spew"
	"github.com/ethereum/go-ethereum/common"
	"github.com/fatih/color"
	"github.com/yearn/ydaemon/common/bigNumber"
	"github.com/yearn/ydaemon/common/helpers"
	"github.com/yearn/ydaemon/common/logs"
	"github.com/yearn/ydaemon/internal/events"
	"github.com/yearn/ydaemon/internal/indexer"
	"github.com/yearn/ydaemon/internal/prices"
	"github.com/yearn/ydaemon/internal/registries"
	"github.com/yearn/ydaemon/internal/strategies"
	"github.com/yearn/ydaemon/internal/tokens"
	"github.com/yearn/ydaemon/internal/vaults"
)

func checkDiff(key string, legacy float64, neo *bigNumber.Float) {
	var blue = color.New(color.FgBlue).Add(color.Bold).SprintFunc()
	var green = color.New(color.FgGreen).Add(color.Bold).SprintFunc()
	var yellow = color.New(color.FgYellow).Add(color.Bold).SprintFunc()
	var red = color.New(color.FgRed).Add(color.Bold).SprintFunc()

	if neo == nil {
		neo = bigNumber.NewFloat(0)
	}

	neoFloat, _ := neo.Float64()
	diffPercent := (neoFloat - legacy) / legacy * 100
	diffPercent = math.Round(diffPercent*100) / 100
	sign := "-"
	if diffPercent > 0 {
		sign = "+"
	}
	diffPercent = math.Abs(diffPercent)
	if legacy == 0 {
		diffPercent = 100
		sign = "+"
	} else if neoFloat == 0 {
		diffPercent = 100
		sign = "-"
	}
	if legacy == neoFloat {
		diffPercent = 0
		sign = " "
	}

	//check if diff > 10%
	if diffPercent < 10 {
		spew.Printf("%s: %-32s vs %-32s | %s%s%%\n", green(key), blue(legacy), blue(neoFloat), green(sign), green(diffPercent))
	} else if diffPercent <= 25 {
		spew.Printf("%s: %-32s vs %-32s | %s%s%%\n", green(key), blue(legacy), blue(neoFloat), yellow(sign), yellow(diffPercent))
	} else {
		spew.Printf("%s: %-32s vs %-32s | %s%s%%\n", green(key), blue(legacy), blue(neoFloat), red(sign), red(diffPercent))
	}
}

/**************************************************************************************************
** getTokenPrice is an helper function which will try to retrieve the price of a token from the
** prices package. If the token is not found, it will return 0.
** The price is returned in base 2 (humanized)
**************************************************************************************************/
func getTokenPrice(chainID uint64, tokenAddr common.Address) *bigNumber.Float {
	price, ok := prices.FindPrice(chainID, tokenAddr)
	if !ok {
		logs.Warning("Could not find price for token " + tokenAddr.Hex())
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
	vaultsMap := registries.RegisterAllVaults(chainID, 0, nil)
	tokens.RetrieveAllTokens(chainID, vaultsMap)
	prices.RetrieveAllPrices(chainID)
	vaults.RetrieveAllVaults(chainID, vaultsMap)
	strategiesAddedList := events.HandleStrategyAdded(chainID, vaultsMap, 0, nil)
	strategies.RetrieveAllStrategies(chainID, strategiesAddedList)
	indexer.PostProcessStrategies(chainID)
}
