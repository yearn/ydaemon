package fetcher

import (
	"math/big"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/common/addresses"
	"github.com/yearn/ydaemon/common/contracts"
	"github.com/yearn/ydaemon/common/env"
	"github.com/yearn/ydaemon/common/ethereum"
	"github.com/yearn/ydaemon/common/helpers"
	"github.com/yearn/ydaemon/common/logs"
	"github.com/yearn/ydaemon/internal/models"
	"github.com/yearn/ydaemon/internal/multicalls"
	"github.com/yearn/ydaemon/internal/storage"
)

/**************************************************************************************************
** fetchTokensBasicInformations will, for a list of addresses, fetch all the relevant basic information
** for the related token. This includes the name, the symbol, the decimals and the details about
** the "underlying tokens" (ex: dai for aDAI). Supported underlying are Curve, AAVE, Compound.
** Yearn's underlying are already in the initial token list.
**
** Arguments:
** - chainID: the chain ID of the network we are working on
** - tokens: a list of addresses of the tokens we want to fetch the information for
**
** Returns:
** - a list of TERC20Token containing the basic information for the tokens
**************************************************************************************************/
func fetchTokensBasicInformations(
	chainID uint64,
	tokens []common.Address,
	toSkip []common.Address,
	curveFactoryPoolMap map[string][]common.Address,
) (sliceOfTokens []models.TERC20Token) {
	chunkSize := 50
	chunks := [][]common.Address{}
	for i := 0; i < len(tokens); i += chunkSize {
		end := i + chunkSize
		if end > len(tokens) {
			end = len(tokens)
		}
		chunks = append(chunks, tokens[i:end])
	}

	chain, ok := env.GetChain(chainID)
	if !ok {
		return sliceOfTokens
	}

	for _, chunk := range chunks {
		/**********************************************************************************************
		** The first step is to prepare the multicall, connecting to the multicall instance and
		** preparing the array of calls to send. All calls for all tokens will be send in a single
		** multicall and will later be accessible via a concatened string `tokenAddress + methodName`.
		**********************************************************************************************/
		caller := ethereum.MulticallClientForChainID[chainID]
		calls := []ethereum.Call{}
		for _, tokenAddress := range chunk {
			if helpers.Contains(toSkip, tokenAddress) {
				continue
			}
			chain, ok := env.GetChain(chainID)
			if !ok {
				continue
			}
			calls = append(calls, multicalls.GetName(tokenAddress.Hex(), tokenAddress))
			calls = append(calls, multicalls.GetSymbol(tokenAddress.Hex(), tokenAddress))
			calls = append(calls, multicalls.GetDecimals(tokenAddress.Hex(), tokenAddress))
			calls = append(calls, multicalls.GetToken(tokenAddress.Hex(), tokenAddress))
			calls = append(calls, multicalls.GetAsset(tokenAddress.Hex(), tokenAddress))
			calls = append(calls, multicalls.GetPoolFromLpToken(tokenAddress.Hex(), chain.Curve.RegistryAddress, tokenAddress))
			calls = append(calls, multicalls.GetCompoundUnderlying(tokenAddress.Hex(), tokenAddress))
			calls = append(calls, multicalls.GetAaveV1Underlying(tokenAddress.Hex(), tokenAddress))
			calls = append(calls, multicalls.GetAaveV2Underlying(tokenAddress.Hex(), tokenAddress))
			calls = append(calls, multicalls.GetCurveMinter(tokenAddress.Hex(), tokenAddress))
			calls = append(calls, multicalls.GetCurveCoin(tokenAddress.Hex()+`_coin0`, tokenAddress, big.NewInt(0)))
			calls = append(calls, multicalls.GetCurveCoin(tokenAddress.Hex()+`_coin1`, tokenAddress, big.NewInt(1)))
		}

		/**********************************************************************************************
		** Then we can proceed the responses. We will create a new relatedTokensList to be able to know
		** which token to fetch then (ex: for aDAI, we also need to fetch the DAI token).
		** Nb: A special case is for Ethereum coin, which is defaulted as address 0xEeeee....EEeE.
		**********************************************************************************************/
		relatedTokensList := []common.Address{}
		response := multicalls.Perform(chainID, calls, nil)
		for _, tokenAddress := range chunk {
			if helpers.Contains(toSkip, tokenAddress) {
				continue
			}
			if addresses.Equals(tokenAddress, `0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE`) {
				sliceOfTokens = append(sliceOfTokens, models.TERC20Token{
					Address:  tokenAddress,
					Name:     `Ethereum`,
					Symbol:   `ETH`,
					Decimals: 18,
					ChainID:  chainID,
				})
				continue
			}
			rawName := response[tokenAddress.Hex()+`name`]
			rawSymbol := response[tokenAddress.Hex()+`symbol`]
			rawDecimals := response[tokenAddress.Hex()+`decimals`]
			rawYearnVaultToken := response[tokenAddress.Hex()+`token`]
			rawYearnVaultAsset := response[tokenAddress.Hex()+`asset`]
			rawPoolFromLpToken := response[tokenAddress.Hex()+`get_pool_from_lp_token`]
			rawCurveCoinMinterToken := response[tokenAddress.Hex()+`minter`]
			rawCUnderlying := response[tokenAddress.Hex()+`underlying`]
			rawAV1Underlying := response[tokenAddress.Hex()+`underlyingAssetAddress`]
			rawAV2Underlying := response[tokenAddress.Hex()+`UNDERLYING_ASSET_ADDRESS`]
			rawCurveCoin0 := response[tokenAddress.Hex()+`_coin0coins`]
			rawCurveCoin1 := response[tokenAddress.Hex()+`_coin1coins`]

			/******************************************************************************************
			** Preparing our new ERC20Token object
			******************************************************************************************/
			newToken := models.TERC20Token{
				Address:                   tokenAddress,
				Name:                      helpers.DecodeString(rawName),
				DisplayName:               helpers.DecodeString(rawName),
				Symbol:                    helpers.DecodeString(rawSymbol),
				DisplaySymbol:             helpers.DecodeString(rawSymbol),
				Decimals:                  helpers.DecodeUint64(rawDecimals),
				UnderlyingTokensAddresses: []common.Address{},
				Icon:                      env.BASE_ASSET_URL + strconv.FormatUint(chainID, 10) + `/` + tokenAddress.Hex() + `/logo-128.png`,
				ChainID:                   chainID,
			}

			/******************************************************************************************
			** Checking if the token is a Yearn Vault. We can determined that if we got a valid
			** response from the `token` RPC call.
			** If so, we set the token type to `Yearn Vault`, we fetch the Coins from the pool and we
			** add the coins to the newToken UnderlyingTokensAddresses.
			** We can also add the coins to the relatedTokensList, so we can fetch their information
			** later.
			******************************************************************************************/
			isV2YearnVault := rawYearnVaultToken != nil && helpers.DecodeAddress(rawYearnVaultToken) != common.Address{}
			isV3YearnVault := rawYearnVaultAsset != nil && helpers.DecodeAddress(rawYearnVaultAsset) != common.Address{}
			if isV2YearnVault || isV3YearnVault {
				vault, ok := storage.GetVaultFromRegistry(chainID, tokenAddress)
				if !ok {
					newToken.Type = models.TokenTypeStandardVault
				} else {
					newToken.Type = vault.Type
				}
				if isV2YearnVault {
					coin := helpers.DecodeAddress(rawYearnVaultToken)
					if (coin != common.Address{}) {
						relatedTokensList = append(relatedTokensList, coin)
						newToken.UnderlyingTokensAddresses = append(newToken.UnderlyingTokensAddresses, coin)
					}
				} else {
					coin := helpers.DecodeAddress(rawYearnVaultAsset)
					if (coin != common.Address{}) {
						relatedTokensList = append(relatedTokensList, coin)
						newToken.UnderlyingTokensAddresses = append(newToken.UnderlyingTokensAddresses, coin)
					}
				}
			}
			if newToken.Category == `` && (isV2YearnVault || isV3YearnVault) {
				newToken.Category = `yVault`
			}

			/******************************************************************************************
			** Checking if the token is a Curve LP token. We can determined that if we got a valid
			** response from the `get_pool_from_lp_token` RPC call.
			** If so, we set the token type to `Curve LP`, we fetch the Coins from the pool and we
			** add the coins to the newToken UnderlyingTokensAddresses.
			** We can also add the coins to the relatedTokensList, so we can fetch their information
			** later.
			******************************************************************************************/
			isCurveLpToken := rawPoolFromLpToken != nil && helpers.DecodeAddress(rawPoolFromLpToken) != common.Address{}
			if isCurveLpToken {
				newToken.Type = models.TokenTypeCurveLP
				curvePoolCaller, _ := contracts.NewCurvePoolRegistryCaller(chain.Curve.RegistryAddress, caller.Client)
				poolCoins, _ := curvePoolCaller.GetCoins(&bind.CallOpts{}, helpers.DecodeAddress(rawPoolFromLpToken))
				for _, coin := range poolCoins {
					if (coin != common.Address{}) {
						relatedTokensList = append(relatedTokensList, coin)
						newToken.UnderlyingTokensAddresses = append(newToken.UnderlyingTokensAddresses, coin)
					}
				}
			}
			if newToken.Category == `` && isCurveLpToken {
				newToken.Category = `Curve`
			}

			/******************************************************************************************
			** Checking if the token is a Compound token. We can determined that if we got a valid
			** response from the `underlying` RPC call.
			** If so, we set the token type to `Compound` and we add the coin we got as response for
			** this RPC call in the newToken UnderlyingTokensAddresses.
			** We can also add the coins to the relatedTokensList, so we can fetch it's information
			** later.
			******************************************************************************************/
			isCToken := rawCUnderlying != nil && helpers.DecodeAddress(rawCUnderlying) != common.Address{}
			if isCToken {
				newToken.Type = models.TokenTypeCompound
				coin := helpers.DecodeAddress(rawCUnderlying)
				if (coin != common.Address{}) {
					relatedTokensList = append(relatedTokensList, coin)
					newToken.UnderlyingTokensAddresses = append(newToken.UnderlyingTokensAddresses, coin)
				}
			}

			/******************************************************************************************
			** Checking if the token is a AAVE V1 token. We can determined that if we got a valid
			** response from the `underlyingAssetAddress` RPC call.
			** If so, we set the token type to `AAVE V1` and we add the coin we got as response for
			** this RPC call in the newToken UnderlyingTokensAddresses.
			** We can also add the coins to the relatedTokensList, so we can fetch it's information
			** later.
			******************************************************************************************/
			isAV1Token := rawAV1Underlying != nil && helpers.DecodeAddress(rawAV1Underlying) != common.Address{}
			if isAV1Token {
				newToken.Type = models.TokenTypeAaveV1
				coin := helpers.DecodeAddress(rawAV1Underlying)
				if (coin != common.Address{}) {
					relatedTokensList = append(relatedTokensList, coin)
					newToken.UnderlyingTokensAddresses = append(newToken.UnderlyingTokensAddresses, coin)
				}
			}

			/******************************************************************************************
			** Checking if the token is a AAVE V2 token. We can determined that if we got a valid
			** response from the `UNDERLYING_ASSET_ADDRESS` RPC call.
			** If so, we set the token type to `AAVE V2` and we add the coin we got as response for
			** this RPC call in the newToken UnderlyingTokensAddresses.
			** We can also add the coins to the relatedTokensList, so we can fetch it's information
			** later.
			******************************************************************************************/
			isAV2Token := rawAV2Underlying != nil && helpers.DecodeAddress(rawAV2Underlying) != common.Address{}
			if isAV2Token {
				newToken.Type = models.TokenTypeAaveV2
				coin := helpers.DecodeAddress(rawAV2Underlying)
				if (coin != common.Address{}) {
					relatedTokensList = append(relatedTokensList, coin)
					newToken.UnderlyingTokensAddresses = append(newToken.UnderlyingTokensAddresses, coin)
				}
			}

			/******************************************************************************************
			** Checking if the token is a Curve LP token. We can determined that if we got a valid
			** response from the `minter` RPC call.
			** This is used for some of the Curve LP tokens.
			******************************************************************************************/
			isCurveLPCoin := rawCurveCoinMinterToken != nil && helpers.DecodeAddress(rawCurveCoinMinterToken) != common.Address{}
			if isCurveLPCoin {
				newToken.Type = models.TokenTypeCurveLP
				minter := helpers.DecodeAddress(rawCurveCoinMinterToken)
				coins, ok := curveFactoryPoolMap[minter.Hex()]
				if ok {
					for _, coin := range coins {
						relatedTokensList = append(relatedTokensList, coin)
						newToken.UnderlyingTokensAddresses = append(newToken.UnderlyingTokensAddresses, coin)
					}
				}
			}
			if newToken.Category == `` && isCurveLPCoin {
				newToken.Category = `Curve`
			}

			/******************************************************************************************
			** Checking if the token is a Curve LP token. We can determined that if we got a valid
			** response from the `coin` RPC call.
			** This is used for some of the Curve LP tokens.
			******************************************************************************************/
			isCurveLPCoinFromCoins := (rawCurveCoin0 != nil && helpers.DecodeAddress(rawCurveCoin0) != common.Address{}) && (rawCurveCoin1 != nil && helpers.DecodeAddress(rawCurveCoin1) != common.Address{})
			if isCurveLPCoinFromCoins {
				newToken.Type = models.TokenTypeCurveLP
				coin0 := helpers.DecodeAddress(rawCurveCoin0)
				coin1 := helpers.DecodeAddress(rawCurveCoin1)
				relatedTokensList = append(relatedTokensList, coin0)
				relatedTokensList = append(relatedTokensList, coin1)
				newToken.UnderlyingTokensAddresses = append(newToken.UnderlyingTokensAddresses, coin0)
				newToken.UnderlyingTokensAddresses = append(newToken.UnderlyingTokensAddresses, coin1)
			}

			sliceOfTokens = append(sliceOfTokens, newToken)
			toSkip = append(toSkip, tokenAddress)
		}

		if len(relatedTokensList) > 0 {
			sliceOfTokens = append(sliceOfTokens, fetchTokensBasicInformations(
				chainID,
				helpers.UniqueArrayAddress(relatedTokensList),
				helpers.UniqueArrayAddress(toSkip),
				curveFactoryPoolMap,
			)...)
		}
	}

	return sliceOfTokens
}

/**************************************************************************************************
** findAllTokens constructs a complete token information map from a list of token addresses.
**
** This function serves as a wrapper around fetchTokensBasicInformations that simplifies token
** information retrieval and organization. It performs these key operations:
**
** 1. Extracts individual token addresses from the input map
** 2. Calls fetchTokensBasicInformations to retrieve detailed token data
** 3. Constructs a new map with the fetched token information
**
** The function maintains clean separation of concerns by isolating the token address extraction,
** information fetching, and map construction steps. This makes the code more maintainable and
** the token retrieval process more understandable.
**
** @param chainID uint64 - The chain ID of the network we are working on
** @param tokenMap map[common.Address]models.TERC20Token - Map of token addresses to existing token data
** @param curveFactoryPoolMap map[string][]common.Address - Map of Curve pool addresses to their underlying tokens
** @return map[common.Address]models.TERC20Token - Updated map of token addresses to complete token information
**************************************************************************************************/
func findAllTokens(
	chainID uint64,
	tokenMap map[common.Address]models.TERC20Token,
	curveFactoryPoolMap map[string][]common.Address,
) map[common.Address]models.TERC20Token {
	newMap := make(map[common.Address]models.TERC20Token)
	tokenMapAddresses := []common.Address{}
	for tokenAddress := range tokenMap {
		tokenMapAddresses = append(tokenMapAddresses, tokenAddress)
	}

	newtokenMap := fetchTokensBasicInformations(chainID, tokenMapAddresses, []common.Address{}, curveFactoryPoolMap)
	for _, token := range newtokenMap {
		newMap[token.Address] = token
	}
	return newMap
}

/**************************************************************************************************
** loadCurvePools retrieves Curve liquidity pool data including underlying tokens.
**
** This function fetches data from the Curve Factory to identify all deployed pools and their
** constituent tokens. The process involves:
**
** 1. Connecting to the appropriate Curve Factory contract for the specified chain
** 2. Determining the total number of pools via the poolCount method
** 3. Executing a multicall to efficiently retrieve all pool addresses in a single RPC request
** 4. Executing a second multicall to fetch the underlying tokens for each pool
** 5. Organizing the results into a map of pool addresses to token arrays
**
** The function includes special handling for chains with RPC limitations (particularly Fantom)
** by optimizing the multicall batch sizes. This ensures reliable operation across all supported
** networks regardless of RPC endpoint constraints.
**
** @param chainID uint64 - The blockchain network ID
** @return map[string][]common.Address - A map where keys are pool addresses (as hex strings) and
**                                       values are arrays of token addresses in that pool
**************************************************************************************************/
func loadCurvePools(chainID uint64) map[string][]common.Address {
	chain, ok := env.GetChain(chainID)
	if !ok {
		return map[string][]common.Address{}
	}

	coinsForPool := make(map[string][]common.Address)

	/**********************************************************************************************
	** The first step is to prepare the multicall, connecting to the multicall instance and
	** preparing the array of calls to send. All calls for all tokens will be send in a single
	** multicall and will later be accessible via a concatened string `tokenAddress + methodName`.
	**********************************************************************************************/
	client := ethereum.GetRPC(chainID)
	curvePoolFactory, _ := contracts.NewCurvePoolFactory(chain.Curve.FactoryAddress, client)
	poolCount, err := curvePoolFactory.PoolCount(&bind.CallOpts{})
	if err != nil {
		return coinsForPool
	}

	/**********************************************************************************************
	** The first step is to prepare the multicall, connecting to the multicall instance and
	** preparing the array of calls to send.
	**********************************************************************************************/
	calls := []ethereum.Call{}
	for i := 0; i < int(poolCount.Int64()); i++ {
		calls = append(calls, multicalls.GetCurveFactoryPool(strconv.Itoa(i), chain.Curve.FactoryAddress, big.NewInt(int64(i))))
	}

	/**********************************************************************************************
	** Regular fix for Fantom's RPC, which limit the number of calls in a multicall to a very low
	** number. We split the multicall in multiple calls of 3 calls each.
	** Otherwise, we just send the multicall as is.
	**********************************************************************************************/
	response := multicalls.Perform(chainID, calls, nil)
	calls = []ethereum.Call{}
	for i := 0; i < int(poolCount.Int64()); i++ {
		poolAddress := helpers.DecodeAddress(response[strconv.Itoa(i)+`pool_list`])
		calls = append(calls, multicalls.GetCurveFactoryCoin(
			poolAddress.Hex(),
			chain.Curve.FactoryAddress,
			poolAddress,
		))
	}

	response = multicalls.Perform(chainID, calls, nil)
	for poolAddressRaw, coinAddressesRaw := range response {
		if len(coinAddressesRaw) == 0 {
			continue
		}
		poolAddress := common.HexToAddress(strings.TrimRight(poolAddressRaw, `get_coins`))
		coinAddresses := coinAddressesRaw[0].([2]common.Address)
		for _, coinAddress := range coinAddresses {
			if (coinAddress == common.Address{}) {
				continue
			}
			coinsForPool[poolAddress.Hex()] = append(coinsForPool[poolAddress.Hex()], coinAddress)
		}
	}
	return coinsForPool
}

/**************************************************************************************************
** loadGammaPools retrieves Gamma liquidity pool data including underlying tokens.
**
** This function fetches Gamma concentrated liquidity pool information from storage, which was
** previously loaded from the Gamma API. For each pool, it extracts:
**
** 1. The pool address (which is the LP token address)
** 2. The two underlying token addresses (token0 and token1)
**
** It then creates simple TERC20Token objects for each address (the pool token and both underlying
** tokens) and organizes them into a map. These token objects contain minimal information (just
** the address and chain ID) as they will be enriched with additional data later in the token
** processing pipeline.
**
** @param chainID uint64 - The blockchain network ID
** @return map[common.Address]models.TERC20Token - A map of token addresses to basic token objects,
**                                               including both pool tokens and their underlying tokens
**************************************************************************************************/
func loadGammaPools(chainID uint64) map[common.Address]models.TERC20Token {
	coinsForPools := make(map[common.Address]models.TERC20Token)
	pools, ok := storage.RetrieveGammaAllData(chainID)
	if !ok {
		return coinsForPools
	}

	/**********************************************************************************************
	** We will loop over the pools and add the tokens to the coinsForPools map.
	**********************************************************************************************/
	for poolAddress, poolData := range pools {
		pool := common.HexToAddress(poolAddress)
		token0 := common.HexToAddress(poolData.Token0)
		token1 := common.HexToAddress(poolData.Token1)
		coinsForPools[pool] = models.TERC20Token{
			Address: pool,
			ChainID: chainID,
		}
		coinsForPools[token0] = models.TERC20Token{
			Address: token0,
			ChainID: chainID,
		}
		coinsForPools[token1] = models.TERC20Token{
			Address: token1,
			ChainID: chainID,
		}
	}
	return coinsForPools
}

/**************************************************************************************************
** loadPendleTokens retrieves token addresses from the Pendle ecosystem.
**
** This function fetches Pendle token information from storage, which was previously loaded
** from the Pendle API or subgraph. It performs these key operations:
**
** 1. Attempts to retrieve the Pendle token data for the specified chain
** 2. If no data is found, returns an empty slice
** 3. If data is found, extracts token addresses and converts them to common.Address format
**
** Pendle tokens include yield-bearing tokens and PT (Principal Tokens) that are part of
** Pendle's yield trading infrastructure. These tokens are important for Yearn strategies
** that may interact with the Pendle ecosystem.
**
** @param chainID uint64 - The blockchain network ID
** @return []common.Address - An array of Pendle-related token addresses
**************************************************************************************************/
func loadPendleTokens(chainID uint64) []common.Address {
	tokens := []common.Address{}
	pendleTokens, ok := storage.RetrievePendleTokens(chainID)
	if !ok {
		return []common.Address{}
	}
	for _, address := range pendleTokens {
		tokens = append(tokens, common.HexToAddress(address.Address))
	}
	return tokens
}

/**************************************************************************************************
** loadVeloTokens retrieves token addresses from the Velodrome ecosystem.
**
** This function fetches token information from the Velodrome Sugar contract, which provides
** a simplified interface to access Velodrome pool data. It performs these steps:
**
** 1. Checks if the chain ID is 10 (Optimism), as Velodrome only exists on Optimism
** 2. Connects to the Velodrome Sugar contract at a specific address
** 3. Iteratively calls the "all" method to retrieve batches of pool data
** 4. Extracts token0, token1, and LP token addresses from each pool
** 5. Implements an adaptive batch size approach to handle potential RPC limitations
**
** The function uses a resilient fetching strategy that reduces the batch size if errors occur,
** ensuring maximum data retrieval even under suboptimal network conditions.
**
** @param chainID uint64 - The blockchain network ID (only processes data if chainID is 10)
** @return []common.Address - An array containing all unique token addresses from the Velodrome ecosystem
**************************************************************************************************/
func loadVeloTokens(chainID uint64) []common.Address {
	veloTokens := []common.Address{}
	if chainID != 10 {
		return veloTokens
	}

	var VELO_SUGAR_ADDRESS = common.HexToAddress(`0x3e532BC1998584fe18e357B5187897ad0110ED3A`)

	/**********************************************************************************************
	** The first step is to prepare the multicall, connecting to the multicall instance and
	** preparing the array of calls to send. All calls for all tokens will be send in a single
	** multicall and will later be accessible via a concatened string `tokenAddress + methodName`.
	**********************************************************************************************/
	client := ethereum.GetRPC(chainID)
	sugar, err := contracts.NewVeloSugarCaller(VELO_SUGAR_ADDRESS, client)
	if err != nil {
		return veloTokens
	}

	/**********************************************************************************************
	** Fetch the tokens
	**********************************************************************************************/
	callSize := 25
	for i := 0; i < 30; i++ {
		start := int64(i * callSize)
		allSugar, err := sugar.All(nil, big.NewInt(int64(callSize)), big.NewInt(start))
		if len(allSugar) == 0 || err != nil {
			if err != nil {
				logs.Error(`error fetching velo sugar`, err)
				if callSize > 1 {
					callSize /= 2
					i += 1
					continue
				}
			}
			break
		}
		callSize = 25

		for _, pair := range allSugar {
			if !helpers.Contains(veloTokens, pair.Token0) {
				veloTokens = append(veloTokens, pair.Token0)
			}
			if !helpers.Contains(veloTokens, pair.Token1) {
				veloTokens = append(veloTokens, pair.Token1)
			}
			if !helpers.Contains(veloTokens, pair.Lp) {
				veloTokens = append(veloTokens, pair.Lp)
			}
		}
	}
	return veloTokens
}

/**************************************************************************************************
** RetrieveAllTokens collects comprehensive token information across the Yearn ecosystem.
**
** This function orchestrates the collection of token data from multiple sources to create
** a complete view of the token ecosystem that Yearn interacts with. The process involves:
**
** 1. Token Source Collection:
**    - Extracts vault tokens and their underlying assets from vault registry
**    - Includes chain-specific extra tokens defined in configuration
**    - Fetches Velodrome ecosystem tokens (Optimism only)
**    - Retrieves Curve Factory pool tokens and their underlying assets
**    - Includes Gamma concentrated liquidity pools and their constituent tokens
**    - Adds Pendle yield tokens to the ecosystem
**
** 2. Token Information Retrieval:
**    - Processes new tokens not already in storage
**    - Refreshes token information when explicitly requested
**    - Fetches basic information (name, symbol, decimals) via multicalls
**    - Determines token types (Yearn vault, Curve LP, Compound, AAVE, etc.)
**    - Identifies underlying tokens recursively
**
** 3. Storage Integration:
**    - Stores token information persistently for future reference
**    - Exports token data to JSON for external consumption
**
** The function implements optimizations to minimize RPC calls by:
** - Utilizing existing token information from storage when available
** - Processing tokens in batches
** - Using multicalls to fetch multiple properties in single RPC requests
**
** @param chainID uint64 - The blockchain network ID
** @param vaults map[common.Address]models.TVault - Map of vault addresses to vault models
** @return map[common.Address]models.TERC20Token - Complete map of token addresses to token information
**************************************************************************************************/
func RetrieveAllTokens(
	chainID uint64,
	vaults map[common.Address]models.TVault,
) map[common.Address]models.TERC20Token {
	chain, ok := env.GetChain(chainID)
	if !ok {
		return map[common.Address]models.TERC20Token{}
	}

	/**********************************************************************************************
	** First, try to retrieve the list of tokens from the database to exclude the one existing
	** from the upcoming calls
	**********************************************************************************************/
	tokenMap, _ := storage.ListERC20(chainID)
	metadata := storage.GetTokensJsonMetadata(chainID)
	shouldRefresh := metadata.ShouldRefresh

	/**********************************************************************************************
	** From the vault registry we have the first batch of tokens: the vault tokens and the
	** underlying tokens for theses vaults. In order to proceed, we will create a map of ERC20
	** tokens, which will be stored in the DB. This action should only be performed once per token
	** per chainID as the token information are not expected to change and will be retrieve on
	** subsequent reboots.
	**********************************************************************************************/
	updatedTokenMap := make(map[common.Address]models.TERC20Token)
	for _, currentVault := range vaults {
		if _, ok := tokenMap[currentVault.Address]; !ok || shouldRefresh {
			updatedTokenMap[currentVault.Address] = models.TERC20Token{
				Address: currentVault.Address,
				Type:    currentVault.Type,
				ChainID: chainID,
			}
		}

		if _, ok := tokenMap[currentVault.AssetAddress]; !ok || shouldRefresh {
			updatedTokenMap[currentVault.AssetAddress] = models.TERC20Token{
				Address: currentVault.AssetAddress,
				Type:    currentVault.Type,
				ChainID: chainID,
			}
		}
	}

	/**********************************************************************************************
	** We may have a list of extra tokens to add to the list. This is used for some tokens that
	** are not in the registry but that we still want to track.
	**********************************************************************************************/
	for _, extraToken := range chain.ExtraTokens {
		if _, ok := tokenMap[extraToken]; !ok || shouldRefresh {
			updatedTokenMap[extraToken] = models.TERC20Token{
				Address: extraToken,
				ChainID: chainID,
			}
		}
	}

	/**********************************************************************************************
	** Fetch the tokens from the Velodrome API.
	**********************************************************************************************/
	veloTokens := loadVeloTokens(chainID)
	for _, veloToken := range veloTokens {
		if _, ok := tokenMap[veloToken]; !ok || shouldRefresh {
			updatedTokenMap[veloToken] = models.TERC20Token{
				Address: veloToken,
				ChainID: chainID,
			}
		}
	}

	/**********************************************************************************************
	** Fetch the tokens from the Curve Factory.
	**********************************************************************************************/
	curveFactoryPools := loadCurvePools(chainID)
	for _, poolCoins := range curveFactoryPools {
		for _, coinAddress := range poolCoins {
			if _, ok := tokenMap[coinAddress]; !ok || shouldRefresh {
				updatedTokenMap[coinAddress] = models.TERC20Token{
					Address: coinAddress,
					ChainID: chainID,
				}
			}
		}
	}

	/**********************************************************************************************
	** Fetch the tokens from the Gamma API.
	**********************************************************************************************/
	gammaPools := loadGammaPools(chainID)
	for _, token := range gammaPools {
		if _, ok := tokenMap[token.Address]; !ok || shouldRefresh {
			updatedTokenMap[token.Address] = token
		}
	}

	/**********************************************************************************************
	** Fetch the tokens from the Pendle V1 API.
	**********************************************************************************************/
	pendleTokens := loadPendleTokens(chainID)
	for _, token := range pendleTokens {
		if _, ok := tokenMap[token]; !ok || shouldRefresh {
			updatedTokenMap[token] = models.TERC20Token{
				Address: token,
				ChainID: chainID,
			}
		}
	}

	/**********************************************************************************************
	** Once we got out basic list, we will fetch theses basics informations. This includes:
	** - the name
	** - the symbol
	** - the decimals
	** Based on the type of token (aave, compound, curve, etc), we will also fetch the underlyings
	** tokens.
	**********************************************************************************************/
	if len(updatedTokenMap) > 0 {
		updatedTokenMap = findAllTokens(chainID, updatedTokenMap, curveFactoryPools)
		for _, token := range updatedTokenMap {
			storage.StoreERC20(chainID, token)
		}
	}
	tokenMap, _ = storage.ListERC20(chainID)
	storage.StoreTokensToJson(chainID, tokenMap)
	return tokenMap
}
