package strategies

import (
	"strconv"

	"github.com/yearn/ydaemon/common/bigNumber"
	"github.com/yearn/ydaemon/common/helpers"
	"github.com/yearn/ydaemon/common/logs"
	"github.com/yearn/ydaemon/internal/models"
	"github.com/yearn/ydaemon/internal/prices"
	"github.com/yearn/ydaemon/internal/tokens"
)

/**********************************************************************************************
** ComputeTVL, calculates the Total Value Locked (TVL) for a given strategy. It retrieves the
** token data associated with the vault of the strategy. If the token data cannot be found,
** it logs a warning and returns 0 for the TVL, amount, and price.
**
** It then attempts to find the price for the token. If the price cannot be found, it logs a
** warning. The function formats the token price and the estimated total assets of the strategy
** to a standard format.
**
** The TVL is calculated by multiplying the amount of tokens in the vault by the price of the
** token. The function returns the calculated TVL, the amount of tokens, and the price of the
** token.
**********************************************************************************************/
func computeTVL(t *models.TStrategy) (*bigNumber.Float, *bigNumber.Float, *bigNumber.Float) {
	tokenData, ok := tokens.FindUnderlyingForVault(t.ChainID, t.VaultAddress)
	if !ok {
		logs.Warning(`impossible to find tokenData for vault ` + t.VaultAddress.Hex() + ` on chain ` + strconv.FormatUint(t.ChainID, 10))
		return bigNumber.NewFloat(0), bigNumber.NewFloat(0), bigNumber.NewFloat(0)
	}

	_tokenPrice, ok := prices.FindPrice(t.ChainID, tokenData.Address)
	if !ok {
		logs.Warning(`impossible to find price for token ` + tokenData.Address.Hex() + ` on chain ` + strconv.FormatUint(t.ChainID, 10))
	}
	_, price := helpers.FormatAmount(_tokenPrice.String(), 6)
	_, amount := helpers.FormatAmount(t.EstimatedTotalAssets.String(), int(tokenData.Decimals))
	tvl := bigNumber.NewFloat(0).Mul(amount, price)
	return tvl, amount, price
}
