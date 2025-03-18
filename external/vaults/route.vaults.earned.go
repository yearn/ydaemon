package vaults

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/machinebox/graphql"
	"github.com/yearn/ydaemon/common/addresses"
	"github.com/yearn/ydaemon/common/bigNumber"
	"github.com/yearn/ydaemon/common/env"
	"github.com/yearn/ydaemon/common/helpers"
	"github.com/yearn/ydaemon/common/logs"
	"github.com/yearn/ydaemon/internal/models"
	"github.com/yearn/ydaemon/internal/storage"
)

/**************************************************************************************************
** graphQLRequestForUser builds a GraphQL query to fetch vault position data for a specific user.
**
** This helper function constructs a GraphQL query that retrieves FIFO (First In, First Out) data
** for a user's positions in vaults. It can be used to query all vaults for a user or filtered to
** specific vaults provided in the vaultAddresses parameter.
**
** @param userAddress string - The Ethereum address of the user whose data is being queried
** @param vaultAddresses []string - Optional list of vault addresses to filter the query
** @return *graphql.Request - A prepared GraphQL request object ready to be executed
**************************************************************************************************/
func graphQLRequestForUser(userAddress string, vaultAddresses []string) *graphql.Request {
	if len(vaultAddresses) == 0 {
		return graphql.NewRequest(`{
			accountVaultPositions(where:{account: "` + strings.ToLower(userAddress) + `"}) {
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

/**************************************************************************************************
** TEarned represents the earnings data for a user's position in a vault.
**
** This structure contains both realized and unrealized gains data in both token and USD values.
** Realized gains represent profits that have been actualized through withdrawals, while
** unrealized gains represent the current value increase that has not yet been withdrawn.
**
** @field RealizedGains string - The amount of realized gains in token units
** @field RealizedGainsUSD float64 - The USD value of realized gains
** @field UnrealizedGains string - The amount of unrealized gains in token units
** @field UnrealizedGainsUSD float64 - The USD value of unrealized gains
**************************************************************************************************/
type TEarned struct {
	RealizedGains      string  `json:"realizedGains"`
	RealizedGainsUSD   float64 `json:"realizedGainsUSD"`
	UnrealizedGains    string  `json:"unrealizedGains"`
	UnrealizedGainsUSD float64 `json:"unrealizedGainsUSD"`
}

/**************************************************************************************************
** GetEarnedPerVaultPerUser calculates and returns earnings data for a specific user in specific
** vaults on a given chain.
**
** This endpoint performs FIFO (First In, First Out) calculations to determine how much a user
** has earned in specified vaults. It retrieves transaction history from a subgraph, processes
** deposit and withdrawal events, and calculates both realized and unrealized gains.
**
** The calculation considers token price changes between deposits and withdrawals to accurately
** determine profits. Results are returned both in token amount and USD value.
**
** Endpoint: GET /chains/:chainID/vaults/earned/:address/:vaults
**
** @param c *gin.Context - The Gin context containing the HTTP request with parameters:
**   - chainID: The blockchain network ID
**   - address: The user's wallet address
**   - vaults: Comma-separated list of vault addresses to calculate earnings for
**
** @return JSON response with:
**   - totalRealizedGainsUSD: Sum of all realized gains in USD
**   - totalUnrealizedGainsUSD: Sum of all unrealized gains in USD
**   - earned: Map of vault addresses to their respective TEarned objects
**************************************************************************************************/
func (y Controller) GetEarnedPerVaultPerUser(c *gin.Context) {
	// Validate chain ID using the utility function
	chainID, ok := validateChainID(c, "chainID")
	if !ok {
		return
	}

	// Validate user address using the utility function
	userAddress, ok := validateAddress(c, "address", chainID)
	if !ok {
		return
	}

	// Validate vault addresses parameter
	vaultsParam := c.Param("vaults")
	vaultsAddressesStr, ok := validateAddressesParam(c, vaultsParam, "GetEarnedPerVaultPerUser")
	if !ok {
		return
	}

	// Get chain configuration
	chain, ok := env.GetChain(chainID)
	if !ok {
		handleError(c, fmt.Errorf("chain configuration not found for chainID %d", chainID),
			http.StatusInternalServerError, "Internal configuration error", "GetEarnedPerVaultPerUser")
		return
	}

	// Validate subgraph endpoint availability
	graphQLEndpoint := chain.SubgraphURI
	if graphQLEndpoint == "" {
		handleError(c, fmt.Errorf("no graph endpoint configured for chainID %d", chainID),
			http.StatusInternalServerError, "Subgraph not available", "GetEarnedPerVaultPerUser")
		return
	}

	// Create GraphQL request
	client := graphql.NewClient(graphQLEndpoint)
	request := graphQLRequestForUser(userAddress.Hex(), vaultsAddressesStr)

	// Execute GraphQL request with timeout context
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var response models.TFIFOForUserForVault
	if err := client.Run(ctx, request, &response); err != nil {
		wrappedErr := fmt.Errorf("failed to execute GraphQL request for user %s: %w", userAddress.Hex(), err)
		handleError(c, wrappedErr, http.StatusInternalServerError,
			"Failed to fetch data from subgraph", "GetEarnedPerVaultPerUser")
		return
	}

	// Verify response contains data
	if response.AccountVaultPositions == nil {
		// Not an error, just no data for this user
		c.JSON(http.StatusOK, gin.H{
			"totalRealizedGainsUSD":   0,
			"totalUnrealizedGainsUSD": 0,
			"earned":                  gin.H{},
		})
		return
	}

	earnedMap := make(map[string]*TEarned)
	listOfVaults := response.AccountVaultPositions
	totalRealizedGainsUSD := 0.0
	totalUnrealizedGainsUSD := 0.0
	for _, currentVault := range listOfVaults {
		vaultAddress := currentVault.Vault.Id
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

		token, _ := storage.GetUnderlyingERC20(chainID, addresses.ToAddress(vaultAddress))
		tokenPrice, _ := storage.GetPrice(chainID, addresses.ToAddress(vaultAddress))
		realizedGainsUSD := helpers.GetHumanizedValue(realizedGains, int(token.Decimals), tokenPrice.Price)
		unrealizedGainsUSD := helpers.GetHumanizedValue(unrealizedGains, int(token.Decimals), tokenPrice.Price)
		if realizedGains.Lt(bigNumber.NewInt(0)) {
			realizedGains = bigNumber.NewInt(0)
		}
		if unrealizedGains.Lt(bigNumber.NewInt(0)) {
			unrealizedGains = bigNumber.NewInt(0)
		}
		if realizedGainsUSD < 0 {
			realizedGainsUSD = 0
		}
		if unrealizedGainsUSD < 0 {
			unrealizedGainsUSD = 0
		}

		earnedMap[vaultAddress] = &TEarned{
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

/**************************************************************************************************
** GetEarnedPerUser calculates and returns earnings data for a specific user across all vaults
** on a given chain.
**
** This endpoint performs FIFO (First In, First Out) calculations to determine how much a user
** has earned across all vaults they've interacted with on the specified blockchain. It retrieves
** transaction history from a subgraph, processes all deposit and withdrawal events, and calculates
** both realized and unrealized gains for each vault.
**
** The calculation considers token price changes between deposits and withdrawals to accurately
** determine profits. Results are aggregated by vault and returned both in token amount and USD value.
**
** Endpoint: GET /chains/:chainID/vaults/earned/:address
**
** @param c *gin.Context - The Gin context containing the HTTP request with parameters:
**   - chainID: The blockchain network ID
**   - address: The user's wallet address
**
** @return JSON response with:
**   - totalRealizedGainsUSD: Sum of all realized gains in USD across all vaults
**   - totalUnrealizedGainsUSD: Sum of all unrealized gains in USD across all vaults
**   - earned: Map of vault addresses to their respective TEarned objects
**************************************************************************************************/
func (y Controller) GetEarnedPerUser(c *gin.Context) {
	// Validate chain ID using the utility function
	chainID, ok := validateChainID(c, "chainID")
	if !ok {
		return
	}

	// Validate user address using the utility function
	userAddress, ok := validateAddress(c, "address", chainID)
	if !ok {
		return
	}

	chain, ok := env.GetChain(chainID)
	if !ok {
		return
	}
	graphQLEndpoint := chain.SubgraphURI
	if graphQLEndpoint == "" {
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
		vaultAddress := currentVault.Vault.Id
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
		token, ok := storage.GetUnderlyingERC20(chainID, addresses.ToAddress(vaultAddress))
		tokenPrice, _ := storage.GetPrice(chainID, addresses.ToAddress(vaultAddress))
		if !ok {
			tokenDecimal = int(token.Decimals)
		}
		realizedGainsUSD := helpers.GetHumanizedValue(realizedGains, tokenDecimal, tokenPrice.Price)
		unrealizedGainsUSD := helpers.GetHumanizedValue(unrealizedGains, tokenDecimal, tokenPrice.Price)
		if realizedGains.Lt(bigNumber.NewInt(0)) {
			realizedGains = bigNumber.NewInt(0)
		}
		if unrealizedGains.Lt(bigNumber.NewInt(0)) {
			unrealizedGains = bigNumber.NewInt(0)
		}
		if realizedGainsUSD < 0 {
			realizedGainsUSD = 0
		}
		if unrealizedGainsUSD < 0 {
			unrealizedGainsUSD = 0
		}

		earnedMap[vaultAddress] = &TEarned{
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

/**************************************************************************************************
** GetEarnedPerUserForAllChains calculates and returns earnings data for a specific user across
** all vaults on multiple chains.
**
** This endpoint aggregates earnings data across multiple blockchain networks, performing FIFO
** calculations for each chain and combining the results. It allows querying specific chains via
** the chainIDs query parameter or defaults to all supported chains if not specified.
**
** The function processes deposit and withdrawal events from each chain's subgraph, calculating
** realized and unrealized gains for every vault the user has interacted with. Results are
** organized hierarchically by chain ID and vault address.
**
** Endpoint: GET /vaults/earned/:address
**
** Query Parameters:
**   - chainIDs: Optional comma-separated list of chain IDs to include (defaults to all supported chains)
**
** @param c *gin.Context - The Gin context containing the HTTP request with parameters:
**   - address: The user's wallet address
**
** @return JSON response with:
**   - totalRealizedGainsUSD: Sum of all realized gains in USD across all chains and vaults
**   - totalUnrealizedGainsUSD: Sum of all unrealized gains in USD across all chains and vaults
**   - earned: Nested map of chain IDs to vault addresses to their respective TEarned objects
**************************************************************************************************/
func (y Controller) GetEarnedPerUserForAllChains(c *gin.Context) {
	// Validate address using the utility function - we use chain ID 1 as default for format validation
	userAddress, ok := validateAddress(c, "address", 1)
	if !ok {
		return
	}

	/** 🔵 - Yearn *************************************************************************************
	** chainsStr: A string that represents the chain IDs for which the vaults are to be returned. It is
	** obtained from the 'chainIDs' query parameter in the request. The string is split by commas to
	** obtain an array of chain IDs.
	**
	** chains: An array of uint64 values that represents the chain IDs for which the vaults are to be
	** returned. If the 'chains' query parameter is not provided or is empty, this array defaults to
	** all supported chain IDs.
	**
	** The 'chains' array is populated by iterating over the 'chainsStr' array and converting each
	** chain ID to a uint64 value. If the conversion is not successful, the chain ID is ignored.
	**
	** The 'chains' array is used to filter the vaults that are returned in the response.
	**************************************************************************************************/
	chainsStr := strings.Split(getQueryParam(c, `chainIDs`), `,`)
	chains := []uint64{}
	if len(chainsStr) == 0 || (len(chainsStr) == 1 && chainsStr[0] == ``) {
		chains = env.SUPPORTED_CHAIN_IDS
	} else {
		for _, chainStr := range chainsStr {
			chain, ok := helpers.AssertChainID(chainStr)
			if !ok {
				continue
			}
			chains = append(chains, chain)
		}
	}

	earnedMap := make(map[uint64]map[string]*TEarned)
	totalRealizedGainsUSD := 0.0
	totalUnrealizedGainsUSD := 0.0
	for _, chainID := range chains {
		earnedMap[chainID] = make(map[string]*TEarned)
		chain, ok := env.GetChain(chainID)
		if !ok {
			return
		}
		graphQLEndpoint := chain.SubgraphURI
		if graphQLEndpoint == "" {
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

		listOfVaults := response.AccountVaultPositions
		for _, currentVault := range listOfVaults {
			vaultAddress := currentVault.Vault.Id
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
			token, ok := storage.GetUnderlyingERC20(chainID, addresses.ToAddress(vaultAddress))
			tokenPrice, _ := storage.GetPrice(chainID, addresses.ToAddress(vaultAddress))
			if !ok {
				tokenDecimal = int(token.Decimals)
			}
			realizedGainsUSD := helpers.GetHumanizedValue(realizedGains, tokenDecimal, tokenPrice.Price)
			unrealizedGainsUSD := helpers.GetHumanizedValue(unrealizedGains, tokenDecimal, tokenPrice.Price)
			if realizedGains.Lt(bigNumber.NewInt(0)) {
				realizedGains = bigNumber.NewInt(0)
			}
			if unrealizedGains.Lt(bigNumber.NewInt(0)) {
				unrealizedGains = bigNumber.NewInt(0)
			}
			if realizedGainsUSD < 0 {
				realizedGainsUSD = 0
			}
			if unrealizedGainsUSD < 0 {
				unrealizedGainsUSD = 0
			}

			earnedMap[chainID][vaultAddress] = &TEarned{
				RealizedGains:      realizedGains.String(),
				RealizedGainsUSD:   realizedGainsUSD,
				UnrealizedGains:    unrealizedGains.String(),
				UnrealizedGainsUSD: unrealizedGainsUSD,
			}
			totalRealizedGainsUSD += realizedGainsUSD
			totalUnrealizedGainsUSD += unrealizedGainsUSD
		}
	}

	c.JSON(http.StatusOK, gin.H{
		`totalRealizedGainsUSD`:   totalRealizedGainsUSD,
		`totalUnrealizedGainsUSD`: totalUnrealizedGainsUSD,
		`earned`:                  earnedMap,
	})
}
