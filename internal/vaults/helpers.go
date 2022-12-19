package vaults

import (
	"math"

	"github.com/yearn/ydaemon/common/bigNumber"
	"github.com/yearn/ydaemon/common/helpers"
	"github.com/yearn/ydaemon/common/types/common"
	"github.com/yearn/ydaemon/internal/prices"
)

// Get the price of the underlying asset. This is tricky because of the decimals. The prices are fetched
// using the lens oracle daemon, stored in the TokenPrices map, and based on the USDC token, aka with 6
// decimals. We first need to parse the Int Price to a float64, then divide by 10^6 to get the price
// in an human readable USDC format.
func getHumanizedTokenPrice(chainID uint64, tokenAddress common.Address) (*bigNumber.Float, float64) {
	fPrice := new(bigNumber.Float)
	price, ok := prices.FindPrice(chainID, tokenAddress)
	if ok {
		fPrice.SetInt(price)
		humanizedPrice := new(bigNumber.Float).Quo(fPrice, bigNumber.NewFloat(math.Pow10(int(6))))
		fHumanizedPrice, _ := humanizedPrice.Float64()
		return humanizedPrice, fHumanizedPrice
	}
	return bigNumber.NewFloat(), 0.0
}

// Get the total assets locked in this vault. This is tricky because of the decimals. The total asset value
// is a string which will be formated as a float64 and scaled with the underlying token decimals. With that
// we should have a human readable Total Asset value, and we should be able to get the Total Value Locked
// in the vault thanks to the above humanizedPrice value.
func getHumanizedValue(balanceToken *bigNumber.Int, decimals int, humanizedPrice *bigNumber.Float) float64 {
	_, humanizedValue := helpers.FormatAmount(balanceToken.String(), decimals)
	fHumanizedValue, _ := bigNumber.NewFloat().Mul(humanizedValue, humanizedPrice).Float64()
	return fHumanizedValue
}
