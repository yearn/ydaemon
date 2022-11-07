package vaults

import (
	"context"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/machinebox/graphql"
	"github.com/yearn/ydaemon/internal/utils/env"
	"github.com/yearn/ydaemon/internal/utils/helpers"
	"github.com/yearn/ydaemon/internal/utils/logs"
	"github.com/yearn/ydaemon/internal/utils/models"
	"github.com/yearn/ydaemon/internal/utils/types/bigNumber"
)

func graphQLRequestForUser(userAddress, vaultAddress string) *graphql.Request {
	return graphql.NewRequest(`{
		accountVaultPositions(where:{account: "` + strings.ToLower(userAddress) + `", vault: "` + strings.ToLower(vaultAddress) + `"}) {
			` + helpers.GetFIFOForUser() + `
		}
	}`)
}

//GetEarnedPerVaultPerUser will, for a given chainID, return the amount earned by an user in a vault from a FIFO perspective
func (y Controller) GetEarnedPerVaultPerUser(c *gin.Context) {
	chainID, ok := helpers.AssertChainID(c.Param("chainID"))
	if !ok {
		c.String(http.StatusBadRequest, "invalid chainID")
		return
	}

	userAddress := `0x537f83ea54e89c967cef5fc313c9160ed2a4ae5c` //`0x165abbd2ffe204900b52036b210a0d76b2ac229e`
	addressesStr := `0x27b5739e22ad9033bcbf192059122d163b60349d`

	graphQLEndpoint, ok := env.THEGRAPH_ENDPOINTS[chainID]
	if !ok {
		logs.Error(`No graph endpoint for chainID`, chainID)
		c.String(http.StatusInternalServerError, `impossible to fetch subgraph`)
		return
	}

	client := graphql.NewClient(graphQLEndpoint)
	request := graphQLRequestForUser(userAddress, addressesStr)
	var response models.TFIFOForUserForVault
	if err := client.Run(context.Background(), request, &response); err != nil {
		logs.Error(err)
		c.String(http.StatusInternalServerError, `invalid graphQL response`)
		return
	}

	allUpdates := response.AccountVaultPositions[0].Updates
	pricePerShare := bigNumber.NewInt(0).SetString(response.AccountVaultPositions[0].Vault.LatestUpdate.PricePerShare)
	decimals := bigNumber.NewInt(response.AccountVaultPositions[0].Vault.ShareToken.Decimals)
	decimals = bigNumber.NewInt(0).Exp(bigNumber.NewInt(10), decimals, nil)
	decimalsFloat := bigNumber.NewFloat(0).SetInt(decimals)
	allDepositsPricePerShare := []*bigNumber.Float{}
	allWithdrawalsPricePerShare := []*bigNumber.Float{}
	allShareMinted := []*bigNumber.Int{}
	allShareBurned := []*bigNumber.Int{}

	/**************************************************************************************************
	** Processing all the updates to get:
	** - a list of all deposits price per share
	** - a list of all withdrawals price per share
	** - a list of all share minted
	** - a list of all share burned
	** We will have as many elements in depositsPricePerShare as in sharesMinted and as many elements
	** in withdrawalsPricePerShare as in sharesBurned, which will allow us to find, for a given index
	** in sharesMinted, the corresponding deposit price per share.
	**************************************************************************************************/
	for _, update := range allUpdates {
		deposit := bigNumber.NewInt().SetString(update.Deposits)
		withdrawal := bigNumber.NewInt().SetString(update.Withdrawals)
		sharesBurnt := bigNumber.NewInt().SetString(update.SharesBurnt)
		sharesMinted := bigNumber.NewInt().SetString(update.SharesMinted)
		sharesReceived := bigNumber.NewInt().SetString(update.SharesReceived)
		sharesSent := bigNumber.NewInt().SetString(update.SharesSent)

		if deposit != nil && !deposit.IsZero() {
			pricePerShareAtThisTime := bigNumber.NewFloat(0).Div(bigNumber.NewFloat().SetInt(deposit), bigNumber.NewFloat().SetInt(sharesMinted))
			allDepositsPricePerShare = append(allDepositsPricePerShare, pricePerShareAtThisTime)
			allShareMinted = append(allShareMinted, sharesMinted)
		}
		if sharesReceived != nil && !sharesReceived.IsZero() {
			pricePerShareAtThisTime := bigNumber.NewFloat(0).Div(bigNumber.NewFloat().SetInt(sharesReceived), bigNumber.NewFloat().SetInt(sharesMinted))
			allDepositsPricePerShare = append(allDepositsPricePerShare, pricePerShareAtThisTime)
			allShareMinted = append(allShareMinted, sharesMinted)
		}
		if withdrawal != nil && !withdrawal.IsZero() {
			pricePerShareAtThisTime := bigNumber.NewFloat(0).Div(bigNumber.NewFloat().SetInt(withdrawal), bigNumber.NewFloat().SetInt(sharesBurnt))
			allWithdrawalsPricePerShare = append(allWithdrawalsPricePerShare, pricePerShareAtThisTime)
			allShareBurned = append(allShareBurned, sharesBurnt)
		}
		if sharesSent != nil && !sharesSent.IsZero() {
			pricePerShareAtThisTime := bigNumber.NewFloat(0).Div(bigNumber.NewFloat().SetInt(sharesSent), bigNumber.NewFloat().SetInt(sharesBurnt))
			allWithdrawalsPricePerShare = append(allWithdrawalsPricePerShare, pricePerShareAtThisTime)
			allShareBurned = append(allShareBurned, sharesBurnt)
		}
	}

	/**************************************************************************************************
	** First step is to compute the realized gains or losses. Plan is to use the FIFO method, aka
	** First In First Out. In order to do so, we will iterate over all withdrawals (allSharesBurned)
	** and compare the amount of shares burned with the oldest shares minted. This oldest shares minted
	** will act like if it was a queue, meaning that we will withdraw X shares burned from it, with it's
	** own pricePerShare, until it is empty, then we will move to the next oldest shares minted and
	** so on.
	**
	** Example:
	**
	** Shares minted: [100, 100] [pricePerShare: 1, 1.2]
	** Shares burned: [50, 80] [pricePerShare: 1.5, 1.5]
	**
	** Step 1:
	** - 50 shares burned from the first shares minted, which has a pricePerShare of 1
	** - New Shares minted queue: [50, 100] [pricePerShare: 1, 1.2]
	** - FIFO earned: 50 * 1.5 - 50 * 1 = 50 * 0.5 = 25
	** - Total FIFO earned: 25
	**
	** Step 2:
	** - 50 shares burned from the first shares minted, which has a pricePerShare of 1 (max possible)
	** - 30 shares burned from the second shares minted, which has a pricePerShare of 1.2
	** - New Shares minted queue: [0, 70] [pricePerShare: 1, 1.2]
	** - FIFO earned: 50 * 1.5 - 50 * 1 = 50 * 0.5 = 25
	** - FIFO earned: 30 * 1.5 - 30 * 1.2 = 30 * 0.3 = 9
	** - Total FIFO earned: 25 + 9 = 34
	**************************************************************************************************/
	fifoRealizedGains := []*bigNumber.Float{}
	latestMintIndex := 0
	for index := 0; index < len(allShareBurned); {
		currentShareBurned := allShareBurned[index]
		latestSharesMinted := allShareMinted[latestMintIndex]

		// Take the PPS for the share minted & the share burned
		latestSharesMintedPPS := allDepositsPricePerShare[latestMintIndex]
		currentShareBurnedPPS := allWithdrawalsPricePerShare[index]

		if currentShareBurned.Eq(latestSharesMinted) {
			// Calculate the value of the deposit when it was minted
			currentShareBurnedValue := bigNumber.NewFloat(0).Mul(bigNumber.NewFloat().SetInt(currentShareBurned), currentShareBurnedPPS)
			currentShareBurnedValue = bigNumber.NewFloat(0).Div(currentShareBurnedValue, decimalsFloat)

			// Calculate the value of the deposit with the PPS when the share was burned
			mintedShareValueForThisAmountOfBurnedShares := bigNumber.NewFloat(0).Mul(bigNumber.NewFloat().SetInt(currentShareBurned), latestSharesMintedPPS)
			mintedShareValueForThisAmountOfBurnedShares = bigNumber.NewFloat(0).Div(mintedShareValueForThisAmountOfBurnedShares, decimalsFloat)

			// Do a valueWhenBurned - valueWhenMinted
			gainOrLoss := bigNumber.NewFloat(0).Sub(currentShareBurnedValue, mintedShareValueForThisAmountOfBurnedShares)

			// This is the realized gain or loss for this withdrawal
			fifoRealizedGains = append(fifoRealizedGains, gainOrLoss)
			allShareMinted[latestMintIndex] = bigNumber.NewInt(0)
			index++
			latestMintIndex++
		}

		if currentShareBurned.Gt(latestSharesMinted) {
			// Calculate the value of the deposit when it was minted
			currentShareBurnedValue := bigNumber.NewFloat(0).Mul(bigNumber.NewFloat().SetInt(currentShareBurned), currentShareBurnedPPS)
			currentShareBurnedValue = bigNumber.NewFloat(0).Div(currentShareBurnedValue, decimalsFloat)

			// Calculate the value of the deposit with the PPS when the share was burned
			mintedShareValueForRemainingAmountOfShares := bigNumber.NewFloat(0).Mul(bigNumber.NewFloat().SetInt(latestSharesMinted), latestSharesMintedPPS)
			mintedShareValueForRemainingAmountOfShares = bigNumber.NewFloat(0).Div(mintedShareValueForRemainingAmountOfShares, decimalsFloat)

			// Do a valueWhenBurned - valueWhenMinted
			gainOrLoss := bigNumber.NewFloat(0).Sub(currentShareBurnedValue, mintedShareValueForRemainingAmountOfShares)

			// This is the realized gain or loss for this withdrawal
			fifoRealizedGains = append(fifoRealizedGains, gainOrLoss)

			allShareMinted[latestMintIndex] = bigNumber.NewInt(0)
			latestMintIndex++
		}

		if currentShareBurned.Lt(latestSharesMinted) {
			// Calculate the value of the deposit when it was minted
			currentShareBurnedValue := bigNumber.NewFloat(0).Mul(bigNumber.NewFloat().SetInt(currentShareBurned), currentShareBurnedPPS)
			currentShareBurnedValue = bigNumber.NewFloat(0).Div(currentShareBurnedValue, decimalsFloat)

			// Calculate the value of the deposit with the PPS when the share was burned
			mintedShareValueForThisAmountOfBurnedShares := bigNumber.NewFloat(0).Mul(bigNumber.NewFloat().SetInt(currentShareBurned), latestSharesMintedPPS)
			mintedShareValueForThisAmountOfBurnedShares = bigNumber.NewFloat(0).Div(mintedShareValueForThisAmountOfBurnedShares, decimalsFloat)

			// Do a valueWhenBurned - valueWhenMinted
			gainOrLoss := bigNumber.NewFloat(0).Sub(currentShareBurnedValue, mintedShareValueForThisAmountOfBurnedShares)

			// This is the realized gain or loss for this withdrawal
			fifoRealizedGains = append(fifoRealizedGains, gainOrLoss)

			allShareMinted[latestMintIndex] = bigNumber.NewInt(0).Sub(allShareMinted[latestMintIndex], currentShareBurned)
			index++
		}
	}
	realizedGains := bigNumber.NewFloat(0)
	for _, gain := range fifoRealizedGains {
		realizedGains.Add(realizedGains, gain)
	}

	/**************************************************************************************************
	** Second step is to do almost the same thing for the unrealized gains. We will proceed just like
	** we did for the realized gains, but we will use the current price per share and all the
	** remaining shares.
	**************************************************************************************************/
	fifoUnRealizedGains := []*bigNumber.Float{}
	currentPPS := bigNumber.NewFloat(0).Div(bigNumber.NewFloat().SetInt(pricePerShare), decimalsFloat)
	for index := range allShareMinted {
		latestSharesMinted := allShareMinted[index]
		latestSharesMintedPPS := allDepositsPricePerShare[index]

		mintedValue := bigNumber.NewFloat(0).Mul(bigNumber.NewFloat().SetInt(latestSharesMinted), latestSharesMintedPPS)
		mintedValue = bigNumber.NewFloat(0).Div(mintedValue, decimalsFloat)

		scaledValue := bigNumber.NewFloat(0).Mul(bigNumber.NewFloat().SetInt(latestSharesMinted), currentPPS)
		scaledValue = bigNumber.NewFloat(0).Div(scaledValue, decimalsFloat)

		fifoUnRealizedGains = append(fifoUnRealizedGains, bigNumber.NewFloat(0).Sub(scaledValue, mintedValue))
	}
	unrealizedGains := bigNumber.NewFloat(0)
	for _, gain := range fifoUnRealizedGains {
		unrealizedGains.Add(unrealizedGains, gain)
	}

	c.JSON(http.StatusOK, gin.H{
		"realizedGains":   realizedGains,
		"unrealizedGains": unrealizedGains,
	})
}
