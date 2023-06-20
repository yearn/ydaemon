package vaults

import (
	"context"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/machinebox/graphql"
	"github.com/yearn/ydaemon/common/addresses"
	"github.com/yearn/ydaemon/common/bigNumber"
	"github.com/yearn/ydaemon/common/env"
	"github.com/yearn/ydaemon/common/helpers"
	"github.com/yearn/ydaemon/common/logs"
	"github.com/yearn/ydaemon/common/models"
	"github.com/yearn/ydaemon/internal/prices"
	"github.com/yearn/ydaemon/internal/tokens"
)

func graphQLRequestForUser(userAddress string, vaultAddresses []string) *graphql.Request {
	if len(vaultAddresses) == 0 {
		return graphql.NewRequest(`{
			accountVaultPositions(account: "` + strings.ToLower(userAddress) + `") {
				` + helpers.GetFIFOForUser() + `
			}
		}`)
	}
	return graphql.NewRequest(`{
		accountVaultPositions(where:{account: "` + strings.ToLower(userAddress) + `", vault_in: ["` + strings.Join(vaultAddresses, `", "`) + `"]}) {
			` + helpers.GetFIFOForUser() + `
		}
	}`)
}

type TEarned struct {
	RealizedGains      string  `json:"realizedGains"`
	RealizedGainsUSD   float64 `json:"realizedGainsUSD"`
	UnrealizedGains    string  `json:"unrealizedGains"`
	UnrealizedGainsUSD float64 `json:"unrealizedGainsUSD"`
}

// GetEarnedPerVaultPerUser will, for a given chainID, return the amount earned by an user in a vault from a FIFO perspective
func (y Controller) GetEarnedPerVaultPerUser(c *gin.Context) {
	chainID, ok := helpers.AssertChainID(c.Param("chainID"))
	if !ok {
		c.String(http.StatusBadRequest, "invalid chainID")
		return
	}
	userAddress, ok := helpers.AssertAddress(c.Param("address"), chainID)
	if !ok {
		c.String(http.StatusBadRequest, "invalid address")
		return
	}
	vaultsAddressesStr := strings.Split(strings.ToLower(c.Param(`vaults`)), `,`)

	graphQLEndpoint, ok := env.THEGRAPH_ENDPOINTS[chainID]
	if !ok {
		logs.Error(`No graph endpoint for chainID`, chainID)
		c.String(http.StatusInternalServerError, `impossible to fetch subgraph`)
		return
	}

	client := graphql.NewClient(graphQLEndpoint)
	request := graphQLRequestForUser(userAddress.Hex(), vaultsAddressesStr)
	var response models.TFIFOForUserForVault
	if err := client.Run(context.Background(), request, &response); err != nil {
		logs.Error(err)
		c.String(http.StatusInternalServerError, `invalid graphQL response`)
		return
	}

	earnedMap := make(map[string]*TEarned)
	listOfVaults := response.AccountVaultPositions
	totalRealizedGainsUSD := 0.0
	totalUnrealizedGainsUSD := 0.0
	for _, currentVault := range listOfVaults {
		vaultAddress := addresses.ToMixedcase(currentVault.Vault.Id)
		pricePerShare := bigNumber.NewInt(0).SetString(currentVault.Vault.LatestUpdate.PricePerShare)
		decimals := bigNumber.NewInt(currentVault.Vault.ShareToken.Decimals)
		decimals = bigNumber.NewInt(0).Exp(bigNumber.NewInt(10), decimals, nil)
		decimalsFloat := bigNumber.NewFloat(0).SetInt(decimals)
		allDepositsPricePerShare := []*bigNumber.Int{}
		allWithdrawalsPricePerShare := []*bigNumber.Int{}
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
		for _, update := range currentVault.Updates {
			deposit := bigNumber.NewInt().SetString(update.Deposits)
			withdrawal := bigNumber.NewInt().SetString(update.Withdrawals)
			sharesBurnt := bigNumber.NewInt().SetString(update.SharesBurnt)
			sharesMinted := bigNumber.NewInt().SetString(update.SharesMinted)
			sharesReceived := bigNumber.NewInt().SetString(update.SharesReceived)
			sharesSent := bigNumber.NewInt().SetString(update.SharesSent)
			tokensSent := bigNumber.NewInt().SetString(update.TokensSent)
			tokensReceived := bigNumber.NewInt().SetString(update.TokensReceived)

			if deposit != nil && !deposit.IsZero() {
				pricePerShareAtThisTime := bigNumber.NewInt(0).Mul(deposit, decimals).Div(sharesMinted)
				allDepositsPricePerShare = append(allDepositsPricePerShare, pricePerShareAtThisTime)
				allShareMinted = append(allShareMinted, sharesMinted)
			}
			if withdrawal != nil && !withdrawal.IsZero() {
				pricePerShareAtThisTime := bigNumber.NewInt(0).Mul(withdrawal, decimals).Div(sharesBurnt)
				allWithdrawalsPricePerShare = append(allWithdrawalsPricePerShare, pricePerShareAtThisTime)
				allShareBurned = append(allShareBurned, sharesBurnt)
			}

			if sharesReceived != nil && !sharesReceived.IsZero() {
				pricePerShareAtThisTime := bigNumber.NewInt(0).Mul(tokensReceived, decimals).Div(sharesReceived)
				allDepositsPricePerShare = append(allDepositsPricePerShare, pricePerShareAtThisTime)
				allShareMinted = append(allShareMinted, sharesReceived)
			}
			if sharesSent != nil && !sharesSent.IsZero() {
				pricePerShareAtThisTime := bigNumber.NewInt(0).Mul(tokensSent, decimals).Div(sharesSent)
				allWithdrawalsPricePerShare = append(allWithdrawalsPricePerShare, pricePerShareAtThisTime)
				allShareBurned = append(allShareBurned, sharesSent)
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
		fifoRealizedGains := []*bigNumber.Int{}
		latestMintIndex := 0
		for index := 0; index < len(allShareBurned); {
			currentShareBurned := allShareBurned[index]
			if latestMintIndex >= len(allShareMinted) {
				break
			}
			latestSharesMinted := allShareMinted[latestMintIndex]

			latestSharesMintedPPS := allDepositsPricePerShare[latestMintIndex]
			currentShareBurnedPPS := allWithdrawalsPricePerShare[index]
			currentShareBurnedValue := bigNumber.NewInt(0).Mul(currentShareBurned, currentShareBurnedPPS).Div(decimals)

			gainOrLoss := bigNumber.NewInt(0)
			if currentShareBurned.Gt(latestSharesMinted) {
				mintedShareValueForRemainingAmountOfShares := bigNumber.NewInt(0).
					Mul(latestSharesMinted, latestSharesMintedPPS).
					Div(decimals)

				gainOrLoss = bigNumber.NewInt(0).Sub(currentShareBurnedValue, mintedShareValueForRemainingAmountOfShares)
				allShareMinted[latestMintIndex] = bigNumber.NewInt(0)

				latestMintIndex++
			} else if currentShareBurned.Lte(latestSharesMinted) {
				mintedShareValueForThisAmountOfBurnedShares := bigNumber.NewInt(0).
					Mul(currentShareBurned, latestSharesMintedPPS).
					Div(decimals)

				gainOrLoss = bigNumber.NewInt(0).Sub(currentShareBurnedValue, mintedShareValueForThisAmountOfBurnedShares)
				allShareMinted[latestMintIndex] = bigNumber.NewInt(0).Sub(allShareMinted[latestMintIndex], currentShareBurned)

				index++
				if currentShareBurned.Eq(latestSharesMinted) {
					latestMintIndex++
				}
			}

			fifoRealizedGains = append(fifoRealizedGains, gainOrLoss)
		}
		realizedGains := bigNumber.NewInt(0)
		for _, gain := range fifoRealizedGains {
			realizedGains.Add(realizedGains, gain)
		}

		/**************************************************************************************************
		** Second step is to do almost the same thing for the unrealized gains. We will proceed just like
		** we did for the realized gains, but we will use the current price per share and all the
		** remaining shares.
		**************************************************************************************************/
		fifoUnRealizedGains := []*bigNumber.Int{}
		currentPPS := bigNumber.NewFloat(0).Div(bigNumber.NewFloat().SetInt(pricePerShare), decimalsFloat)
		_ = currentPPS
		for index := range allShareMinted {
			latestSharesMinted := allShareMinted[index]
			latestSharesMintedPPS := allDepositsPricePerShare[index]

			mintedValue := bigNumber.NewInt(0).
				Mul(latestSharesMinted, latestSharesMintedPPS).
				Div(decimals)

			scaledValue := bigNumber.NewInt(0).
				Mul(latestSharesMinted, pricePerShare).
				Div(decimals)

			fifoUnRealizedGains = append(fifoUnRealizedGains, bigNumber.NewInt(0).Sub(scaledValue, mintedValue))
		}
		unrealizedGains := bigNumber.NewInt(0)
		for _, gain := range fifoUnRealizedGains {
			unrealizedGains.Add(unrealizedGains, gain)
		}

		token, _ := tokens.FindUnderlyingForVault(chainID, addresses.ToAddress(vaultAddress))
		tokenPrice, _ := prices.FindPrice(chainID, addresses.ToAddress(vaultAddress))
		realizedGainsUSD := helpers.GetHumanizedValue(realizedGains, int(token.Decimals), tokenPrice)
		unrealizedGainsUSD := helpers.GetHumanizedValue(unrealizedGains, int(token.Decimals), tokenPrice)

		earnedMap[vaultAddress.Address().Hex()] = &TEarned{
			RealizedGains:      realizedGains.String(),
			RealizedGainsUSD:   realizedGainsUSD,
			UnrealizedGains:    unrealizedGains.String(),
			UnrealizedGainsUSD: unrealizedGainsUSD,
		}
		totalRealizedGainsUSD += realizedGainsUSD
		totalUnrealizedGainsUSD += unrealizedGainsUSD
	}

	c.JSON(http.StatusOK, gin.H{
		`totalRealizedGainsUSD`:   totalRealizedGainsUSD,
		`totalUnrealizedGainsUSD`: totalUnrealizedGainsUSD,
		`earned`:                  earnedMap,
	})
}

// GetEarnedPerUser will, for a given chainID, return the amount earned by an user in all vaults
func (y Controller) GetEarnedPerUser(c *gin.Context) {
	chainID, ok := helpers.AssertChainID(c.Param("chainID"))
	if !ok {
		c.String(http.StatusBadRequest, "invalid chainID")
		return
	}
	userAddress, ok := helpers.AssertAddress(c.Param("address"), chainID)
	if !ok {
		c.String(http.StatusBadRequest, "invalid address")
		return
	}

	graphQLEndpoint, ok := env.THEGRAPH_ENDPOINTS[chainID]
	if !ok {
		logs.Error(`No graph endpoint for chainID`, chainID)
		c.String(http.StatusInternalServerError, `impossible to fetch subgraph`)
		return
	}

	client := graphql.NewClient(graphQLEndpoint)
	request := graphQLRequestForUser(userAddress.Hex(), []string{})
	var response models.TFIFOForUserForVault
	if err := client.Run(context.Background(), request, &response); err != nil {
		logs.Error(err)
		c.String(http.StatusInternalServerError, `invalid graphQL response`)
		return
	}

	earnedMap := make(map[string]*TEarned)
	totalRealizedGainsUSD := 0.0
	totalUnrealizedGainsUSD := 0.0
	listOfVaults := response.AccountVaultPositions
	for _, currentVault := range listOfVaults {
		vaultAddress := addresses.ToMixedcase(currentVault.Vault.Id)
		pricePerShare := bigNumber.NewInt(0).SetString(currentVault.Vault.LatestUpdate.PricePerShare)
		decimals := bigNumber.NewInt(currentVault.Vault.ShareToken.Decimals)
		decimals = bigNumber.NewInt(0).Exp(bigNumber.NewInt(10), decimals, nil)
		decimalsFloat := bigNumber.NewFloat(0).SetInt(decimals)
		allDepositsPricePerShare := []*bigNumber.Int{}
		allWithdrawalsPricePerShare := []*bigNumber.Int{}
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
		for _, update := range currentVault.Updates {
			deposit := bigNumber.NewInt().SetString(update.Deposits)
			withdrawal := bigNumber.NewInt().SetString(update.Withdrawals)
			sharesBurnt := bigNumber.NewInt().SetString(update.SharesBurnt)
			sharesMinted := bigNumber.NewInt().SetString(update.SharesMinted)
			sharesReceived := bigNumber.NewInt().SetString(update.SharesReceived)
			sharesSent := bigNumber.NewInt().SetString(update.SharesSent)
			tokensSent := bigNumber.NewInt().SetString(update.TokensSent)
			tokensReceived := bigNumber.NewInt().SetString(update.TokensReceived)

			if deposit != nil && !deposit.IsZero() {
				pricePerShareAtThisTime := bigNumber.NewInt(0).Mul(deposit, decimals).Div(sharesMinted)
				allDepositsPricePerShare = append(allDepositsPricePerShare, pricePerShareAtThisTime)
				allShareMinted = append(allShareMinted, sharesMinted)
			}
			if withdrawal != nil && !withdrawal.IsZero() {
				pricePerShareAtThisTime := bigNumber.NewInt(0).Mul(withdrawal, decimals).Div(sharesBurnt)
				allWithdrawalsPricePerShare = append(allWithdrawalsPricePerShare, pricePerShareAtThisTime)
				allShareBurned = append(allShareBurned, sharesBurnt)
			}

			if sharesReceived != nil && !sharesReceived.IsZero() {
				pricePerShareAtThisTime := bigNumber.NewInt(0).Mul(tokensReceived, decimals).Div(sharesReceived)
				allDepositsPricePerShare = append(allDepositsPricePerShare, pricePerShareAtThisTime)
				allShareMinted = append(allShareMinted, sharesReceived)
			}
			if sharesSent != nil && !sharesSent.IsZero() {
				pricePerShareAtThisTime := bigNumber.NewInt(0).Mul(tokensSent, decimals).Div(sharesSent)
				allWithdrawalsPricePerShare = append(allWithdrawalsPricePerShare, pricePerShareAtThisTime)
				allShareBurned = append(allShareBurned, sharesSent)
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
		fifoRealizedGains := []*bigNumber.Int{}
		latestMintIndex := 0
		for index := 0; index < len(allShareBurned); {
			currentShareBurned := allShareBurned[index]
			if latestMintIndex >= len(allShareMinted) {
				break
			}
			latestSharesMinted := allShareMinted[latestMintIndex]

			latestSharesMintedPPS := allDepositsPricePerShare[latestMintIndex]
			currentShareBurnedPPS := allWithdrawalsPricePerShare[index]
			currentShareBurnedValue := bigNumber.NewInt(0).Mul(currentShareBurned, currentShareBurnedPPS).Div(decimals)

			gainOrLoss := bigNumber.NewInt(0)
			if currentShareBurned.Gt(latestSharesMinted) {
				mintedShareValueForRemainingAmountOfShares := bigNumber.NewInt(0).
					Mul(latestSharesMinted, latestSharesMintedPPS).
					Div(decimals)

				gainOrLoss = bigNumber.NewInt(0).Sub(currentShareBurnedValue, mintedShareValueForRemainingAmountOfShares)
				allShareMinted[latestMintIndex] = bigNumber.NewInt(0)

				latestMintIndex++
			} else if currentShareBurned.Lte(latestSharesMinted) {
				mintedShareValueForThisAmountOfBurnedShares := bigNumber.NewInt(0).
					Mul(currentShareBurned, latestSharesMintedPPS).
					Div(decimals)

				gainOrLoss = bigNumber.NewInt(0).Sub(currentShareBurnedValue, mintedShareValueForThisAmountOfBurnedShares)
				allShareMinted[latestMintIndex] = bigNumber.NewInt(0).Sub(allShareMinted[latestMintIndex], currentShareBurned)

				index++
				if currentShareBurned.Eq(latestSharesMinted) {
					latestMintIndex++
				}
			}

			fifoRealizedGains = append(fifoRealizedGains, gainOrLoss)
		}
		realizedGains := bigNumber.NewInt(0)
		for _, gain := range fifoRealizedGains {
			realizedGains.Add(realizedGains, gain)
		}

		/**************************************************************************************************
		** Second step is to do almost the same thing for the unrealized gains. We will proceed just like
		** we did for the realized gains, but we will use the current price per share and all the
		** remaining shares.
		**************************************************************************************************/
		fifoUnRealizedGains := []*bigNumber.Int{}
		currentPPS := bigNumber.NewFloat(0).Div(bigNumber.NewFloat().SetInt(pricePerShare), decimalsFloat)
		_ = currentPPS
		for index := range allShareMinted {
			latestSharesMinted := allShareMinted[index]
			latestSharesMintedPPS := allDepositsPricePerShare[index]

			mintedValue := bigNumber.NewInt(0).
				Mul(latestSharesMinted, latestSharesMintedPPS).
				Div(decimals)

			scaledValue := bigNumber.NewInt(0).
				Mul(latestSharesMinted, pricePerShare).
				Div(decimals)

			fifoUnRealizedGains = append(fifoUnRealizedGains, bigNumber.NewInt(0).Sub(scaledValue, mintedValue))
		}
		unrealizedGains := bigNumber.NewInt(0)
		for _, gain := range fifoUnRealizedGains {
			unrealizedGains.Add(unrealizedGains, gain)
		}

		tokenDecimal := 18
		token, ok := tokens.FindUnderlyingForVault(chainID, addresses.ToAddress(vaultAddress))
		tokenPrice, _ := prices.FindPrice(chainID, addresses.ToAddress(vaultAddress))
		if !ok {
			tokenDecimal = int(token.Decimals)
		}
		if tokenPrice == nil {
			tokenPrice = bigNumber.NewInt(0)
		}
		realizedGainsUSD := helpers.GetHumanizedValue(realizedGains, tokenDecimal, tokenPrice)
		unrealizedGainsUSD := helpers.GetHumanizedValue(unrealizedGains, tokenDecimal, tokenPrice)

		earnedMap[vaultAddress.Address().Hex()] = &TEarned{
			RealizedGains:      realizedGains.String(),
			RealizedGainsUSD:   realizedGainsUSD,
			UnrealizedGains:    unrealizedGains.String(),
			UnrealizedGainsUSD: unrealizedGainsUSD,
		}
		totalRealizedGainsUSD += realizedGainsUSD
		totalUnrealizedGainsUSD += unrealizedGainsUSD
	}

	c.JSON(http.StatusOK, gin.H{
		`totalRealizedGainsUSD`:   totalRealizedGainsUSD,
		`totalUnrealizedGainsUSD`: totalUnrealizedGainsUSD,
		`earned`:                  earnedMap,
	})
}
