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
	pricePerShare := bigNumber.NewInt(0).SetString(response.Vault.LatestUpdate.PricePerShare)
	// Based on the deposits, withdrawals, shares burned and shares minted, we can compute the amount of token he has earned
	// earned := bigNumber.NewInt(0)
	sumOfDeposits := bigNumber.NewInt(0)
	sumOfWithdrawals := bigNumber.NewInt(0)
	sumOfSharesBurned := bigNumber.NewInt(0)
	sumOfSharesMinted := bigNumber.NewInt(0)
	allDepositsPricePerShare := []*bigNumber.Float{}
	allWithdrawalsPricePerShare := []*bigNumber.Float{}
	allDeposits := []*bigNumber.Int{}
	allWithdrawals := []*bigNumber.Int{}
	allShareMinted := []*bigNumber.Int{}
	allShareBurned := []*bigNumber.Int{}

	for _, update := range allUpdates {
		deposit := bigNumber.NewInt().SetString(update.Deposits)
		withdrawal := bigNumber.NewInt().SetString(update.Withdrawals)
		sharesBurnt := bigNumber.NewInt().SetString(update.SharesBurnt)
		sharesMinted := bigNumber.NewInt().SetString(update.SharesMinted)

		if deposit != nil && !deposit.IsZero() {
			pricePerShareAtThisTime := bigNumber.NewFloat(0).Quo(bigNumber.NewFloat().SetInt(deposit), bigNumber.NewFloat().SetInt(sharesMinted))
			logs.Info(`user deposited`, deposit, `shares minted`, sharesMinted, `price per share`, pricePerShareAtThisTime)
			allDepositsPricePerShare = append(allDepositsPricePerShare, pricePerShareAtThisTime)
			allDeposits = append(allDeposits, deposit)
			allShareMinted = append(allShareMinted, sharesMinted)
		}
		if withdrawal != nil && !withdrawal.IsZero() {
			pricePerShareAtThisTime := bigNumber.NewFloat(0).Quo(bigNumber.NewFloat().SetInt(withdrawal), bigNumber.NewFloat().SetInt(sharesBurnt))
			logs.Info(`user withdrew`, withdrawal, `shares burnt`, sharesBurnt, `price per share`, pricePerShareAtThisTime)
			allWithdrawalsPricePerShare = append(allWithdrawalsPricePerShare, pricePerShareAtThisTime)
			allWithdrawals = append(allWithdrawals, withdrawal)
			allShareBurned = append(allShareBurned, sharesBurnt)
		}

		sumOfDeposits.Add(sumOfDeposits, deposit)
		sumOfWithdrawals.Add(sumOfWithdrawals, withdrawal)
		sumOfSharesBurned.Add(sumOfSharesBurned, sharesBurnt)
		sumOfSharesMinted.Add(sumOfSharesMinted, sharesMinted)
	}

	earnedForFifoStep := []*bigNumber.Float{}
	latestMintIndex := 0
	for index := 0; index < len(allShareBurned); {
		currentShareBurned := allShareBurned[index]
		latestSharesMinted := allShareMinted[latestMintIndex]
		if currentShareBurned.Eq(latestSharesMinted) {
			latestSharesMintedPPS := allDepositsPricePerShare[latestMintIndex]
			currentShareBurnedPPS := allWithdrawalsPricePerShare[index]

			latestSharesMintedValue := bigNumber.NewFloat(0).Mul(bigNumber.NewFloat().SetInt(latestSharesMinted), latestSharesMintedPPS)
			currentShareBurnedValue := bigNumber.NewFloat(0).Mul(bigNumber.NewFloat().SetInt(currentShareBurned), currentShareBurnedPPS)
			latestSharesMintedValue = bigNumber.NewFloat(0).Quo(latestSharesMintedValue, bigNumber.NewFloat(1e18))
			currentShareBurnedValue = bigNumber.NewFloat(0).Quo(currentShareBurnedValue, bigNumber.NewFloat(1e18))

			earnedForFifoStep = append(earnedForFifoStep, bigNumber.NewFloat(0).Sub(currentShareBurnedValue, latestSharesMintedValue))
			logs.Success(`currentShareBurnedPPS`, currentShareBurnedPPS)
			logs.Success(`latestSharesMintedPPS`, latestSharesMintedPPS)
			logs.Success(`latestSharesMintedValue`, latestSharesMintedValue)
			logs.Success(`currentShareBurnedValue`, currentShareBurnedValue)
			logs.Success(`earnedForFifoStep`, bigNumber.NewFloat(0).Sub(currentShareBurnedValue, latestSharesMintedValue))
			index++
		}

		for currentShareBurned.Gt(latestSharesMinted) {
			latestSharesMinted := allShareMinted[latestMintIndex]

			latestSharesMintedPPS := allDepositsPricePerShare[latestMintIndex]
			currentShareBurnedPPS := allWithdrawalsPricePerShare[index]

			latestSharesMintedValue := bigNumber.NewFloat(0).Mul(bigNumber.NewFloat().SetInt(latestSharesMinted), latestSharesMintedPPS)
			currentShareBurnedValue := bigNumber.NewFloat(0).Mul(bigNumber.NewFloat().SetInt(currentShareBurned), currentShareBurnedPPS)
			latestSharesMintedValue = bigNumber.NewFloat(0).Quo(latestSharesMintedValue, bigNumber.NewFloat(1e18))
			currentShareBurnedValue = bigNumber.NewFloat(0).Quo(currentShareBurnedValue, bigNumber.NewFloat(1e18))

			earnedForFifoStep = append(earnedForFifoStep, bigNumber.NewFloat(0).Sub(currentShareBurnedValue, latestSharesMintedValue))
			logs.Success(`currentShareBurnedPPS`, currentShareBurnedPPS)
			logs.Success(`latestSharesMintedPPS`, latestSharesMintedPPS)
			logs.Success(`latestSharesMintedValue`, latestSharesMintedValue)
			logs.Success(`currentShareBurnedValue`, currentShareBurnedValue)
			logs.Success(`earnedForFifoStep`, bigNumber.NewFloat(0).Sub(currentShareBurnedValue, latestSharesMintedValue))
			latestMintIndex++
		}

		if currentShareBurned.Lt(latestSharesMinted) {
			latestSharesMintedPPS := allDepositsPricePerShare[latestMintIndex]
			currentShareBurnedPPS := allWithdrawalsPricePerShare[index]

			latestSharesMintedValue := bigNumber.NewFloat(0).Mul(bigNumber.NewFloat().SetInt(latestSharesMinted), latestSharesMintedPPS)
			currentShareBurnedValue := bigNumber.NewFloat(0).Mul(bigNumber.NewFloat().SetInt(currentShareBurned), currentShareBurnedPPS)
			latestSharesMintedValue = bigNumber.NewFloat(0).Quo(latestSharesMintedValue, bigNumber.NewFloat(1e18))
			currentShareBurnedValue = bigNumber.NewFloat(0).Quo(currentShareBurnedValue, bigNumber.NewFloat(1e18))

			earnedForFifoStep = append(earnedForFifoStep, bigNumber.NewFloat(0).Sub(currentShareBurnedValue, latestSharesMintedValue))
			logs.Success(`currentShareBurnedPPS`, currentShareBurnedPPS)
			logs.Success(`latestSharesMintedPPS`, latestSharesMintedPPS)
			logs.Success(`latestSharesMintedValue`, latestSharesMintedValue)
			logs.Success(`currentShareBurnedValue`, currentShareBurnedValue)
			logs.Success(`earnedForFifoStep`, bigNumber.NewFloat(0).Sub(currentShareBurnedValue, latestSharesMintedValue))
			allShareMinted[latestMintIndex] = bigNumber.NewInt(0).Sub(allShareMinted[latestMintIndex], currentShareBurned)
			index++
		}

	}

	realizedGains := bigNumber.NewInt(0).Sub(sumOfWithdrawals, sumOfDeposits)
	remainingShares := bigNumber.NewInt(0).Sub(sumOfSharesMinted, sumOfSharesBurned)

	_ = pricePerShare
	// The unrealized gains are the remaining shares multiplied by the price per share. The pricePerShare used for the unrealized gains is the latest pricePerShare
	// unrealizedGains := bigNumber.NewInt(0).Mul(remainingShares, pricePerShare)

	logs.Info(`sum of deposits`, sumOfDeposits, `sum of withdrawals`, sumOfWithdrawals, `sum of shares minted`, sumOfSharesMinted, `sum of shares burnt`, sumOfSharesBurned)
	logs.Info(`deposits minus withdrawals`, realizedGains, `remaining shares`, remainingShares)

	c.JSON(http.StatusOK, earnedForFifoStep)
}
