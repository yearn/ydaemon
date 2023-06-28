package tokens

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
	"github.com/yearn/ydaemon/common/store"
	"github.com/yearn/ydaemon/internal/meta"
	"github.com/yearn/ydaemon/internal/models"
	"github.com/yearn/ydaemon/internal/multicalls"
)

/**************************************************************************************************
** fetchBasicInformations will, for a list of addresses, fetch all the relevant basic information
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
func fetchBasicInformations(
	chainID uint64,
	tokens []common.Address,
	curveFactoryPoolMap map[string][]common.Address,
) (tokenList []models.TERC20Token) {
	/**********************************************************************************************
	** The first step is to prepare the multicall, connecting to the multicall instance and
	** preparing the array of calls to send. All calls for all tokens will be send in a single
	** multicall and will later be accessible via a concatened string `tokenAddress + methodName`.
	**********************************************************************************************/
	caller := ethereum.MulticallClientForChainID[chainID]
	calls := []ethereum.Call{}
	for _, token := range tokens {
		calls = append(calls, multicalls.GetName(token.Hex(), token))
		calls = append(calls, multicalls.GetSymbol(token.Hex(), token))
		calls = append(calls, multicalls.GetDecimals(token.Hex(), token))
		calls = append(calls, multicalls.GetToken(token.Hex(), token))
		calls = append(calls, multicalls.GetPoolFromLpToken(token.Hex(), env.CURVE_REGISTRY_ADDRESSES[chainID], token))
		calls = append(calls, multicalls.GetCompoundUnderlying(token.Hex(), token))
		calls = append(calls, multicalls.GetAaveV1Underlying(token.Hex(), token))
		calls = append(calls, multicalls.GetAaveV2Underlying(token.Hex(), token))
		calls = append(calls, multicalls.GetCurveMinter(token.Hex(), token))
		calls = append(calls, multicalls.GetCurveCoin(token.Hex()+`_coin0`, token, big.NewInt(0)))
		calls = append(calls, multicalls.GetCurveCoin(token.Hex()+`_coin1`, token, big.NewInt(1)))
	}

	/**********************************************************************************************
	** Then we can proceed the responses. We will create a new relatedTokensList to be able to know
	** which token to fetch then (ex: for aDAI, we also need to fetch the DAI token).
	** Nb: A special case is for Ethereum coin, which is defaulted as address 0xEeeee....EEeE.
	**********************************************************************************************/
	relatedTokensList := []common.Address{}
	response := multicalls.Perform(chainID, calls, nil)
	for _, tokenAddress := range tokens {
		if addresses.Equals(tokenAddress, `0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE`) {
			tokenList = append(tokenList, models.TERC20Token{
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
			Address:  tokenAddress,
			Name:     helpers.DecodeString(rawName),
			Symbol:   helpers.DecodeString(rawSymbol),
			Decimals: helpers.DecodeUint64(rawDecimals),
			Icon:     env.BASE_ASSET_URL + strconv.FormatUint(chainID, 10) + `/` + tokenAddress.Hex() + `/logo-128.png`,
			ChainID:  chainID,
		}

		metaTokenName := newToken.Name
		metaTokenSymbol := newToken.Symbol
		metaTokenDescription := ``
		if tokenFromMeta, ok := meta.GetMetaToken(chainID, tokenAddress); ok {
			metaTokenName = strings.Replace(tokenFromMeta.Name, "\"", "", -1)
			metaTokenSymbol = strings.Replace(tokenFromMeta.Symbol, "\"", "", -1)
			metaTokenDescription = tokenFromMeta.Description
		}
		newToken.DisplayName = metaTokenName
		newToken.DisplaySymbol = metaTokenSymbol
		newToken.Description = metaTokenDescription

		/******************************************************************************************
		** Checking if the token is a Yearn Vault. We can determined that if we got a valid
		** response from the `token` RPC call.
		** If so, we set the token type to `Yearn Vault`, we fetch the Coins from the pool and we
		** add the coins to the newToken UnderlyingTokensAddresses.
		** We can also add the coins to the relatedTokensList, so we can fetch their information
		** later.
		******************************************************************************************/
		isYearnVault := rawYearnVaultToken != nil && helpers.DecodeAddress(rawYearnVaultToken) != common.Address{}
		if isYearnVault {
			newToken.Type = models.TokenTypeVault
			coin := helpers.DecodeAddress(rawYearnVaultToken)
			if (coin != common.Address{}) {
				relatedTokensList = append(relatedTokensList, coin)
				newToken.UnderlyingTokensAddresses = append(newToken.UnderlyingTokensAddresses, coin)
			}
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
			curvePoolCaller, _ := contracts.NewCurvePoolRegistryCaller(env.CURVE_REGISTRY_ADDRESSES[chainID], caller.Client)
			poolCoins, _ := curvePoolCaller.GetCoins(&bind.CallOpts{}, helpers.DecodeAddress(rawPoolFromLpToken))
			for _, coin := range poolCoins {
				if (coin != common.Address{}) {
					relatedTokensList = append(relatedTokensList, coin)
					newToken.UnderlyingTokensAddresses = append(newToken.UnderlyingTokensAddresses, coin)
				}
			}
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

		tokenList = append(tokenList, newToken)
	}

	if len(relatedTokensList) > 0 {
		tokenList = append(tokenList, fetchBasicInformations(chainID, helpers.UniqueArrayAddress(relatedTokensList), curveFactoryPoolMap)...)
	}

	return tokenList
}

/**************************************************************************************************
** findAllTokens is simply a wrapper around our fetchBasicInformations function to make it easier
** to read. It will take the tokenMap, extract the individual addresses, and then call the
** fetchBasicInformations function to get the tokens information, before assigning them to a new
** map.
**
** Arguments:
** - chainID: the chain ID of the network we are working on
** - tokenMap: our map of tokenAddress -> TTokens
**
** Returns:
** - a map of tokenAddress -> TTokens
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

	newtokenMap := fetchBasicInformations(chainID, tokenMapAddresses, curveFactoryPoolMap)
	for _, token := range newtokenMap {
		newMap[token.Address] = token
	}

	return newMap
}

func loadCurvePools(chainID uint64) map[string][]common.Address {
	coinsForPool := make(map[string][]common.Address)
	/**********************************************************************************************
	** The first step is to prepare the multicall, connecting to the multicall instance and
	** preparing the array of calls to send. All calls for all tokens will be send in a single
	** multicall and will later be accessible via a concatened string `tokenAddress + methodName`.
	**********************************************************************************************/
	client := ethereum.GetRPC(chainID)
	curvePoolFactory, _ := contracts.NewCurvePoolFactory(env.CURVE_FACTORIES_ADDRESSES[chainID], client)
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
		calls = append(calls, multicalls.GetCurveFactoryPool(
			strconv.Itoa(i),
			env.CURVE_FACTORIES_ADDRESSES[chainID],
			big.NewInt(int64(i)),
		))
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
			env.CURVE_FACTORIES_ADDRESSES[chainID],
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
** Yearn vaults play with at least 2 tokens: the yVaultToken (aka the vault) and the underlying
** token. This underlying token can be a token or a LP token and therefore can have multiple sub
** tokens.
** The goal of this function is to get a list of all the tokens living in the Yearn Extended
** Ecosystem.
**
** Arguments:
** - chainID: the chain ID of the network we are working on
** - vaults: the array of TVaultsFromRegistry
**
** Returns:
** - a map of tokenAddress -> TTokens
**************************************************************************************************/
func RetrieveAllTokens(
	chainID uint64,
	vaults map[common.Address]models.TVaultsFromRegistry,
) map[common.Address]models.TERC20Token {
	/**********************************************************************************************
	** First, try to retrieve the list of tokens from the database to exclude the one existing
	** from the upcoming calls
	**********************************************************************************************/
	tokenMap, _ := store.ListAllERC20(chainID)

	/**********************************************************************************************
	** From the vault registry we have the first batch of tokens: the vault tokens and the
	** underlying tokens for theses vaults. In order to proceed, we will create a map of ERC20
	** tokens, which will be stored in the DB. This action should only be performed once per token
	** per chainID as the token information are not expected to change and will be retrieve on
	** subsequent reboots.
	**********************************************************************************************/
	updatedTokenMap := make(map[common.Address]models.TERC20Token)
	for _, currentVault := range vaults {
		updatedTokenMap[currentVault.Address] = models.TERC20Token{
			Address: currentVault.Address,
			ChainID: chainID,
		}
	}

	// RESET ALL DB
	for _, currentVault := range vaults {
		if _, ok := tokenMap[currentVault.Address]; !ok {
			updatedTokenMap[currentVault.Address] = models.TERC20Token{
				Address: currentVault.Address,
				ChainID: chainID,
			}
		}

		if _, ok := tokenMap[currentVault.TokenAddress]; !ok {
			updatedTokenMap[currentVault.TokenAddress] = models.TERC20Token{
				Address: currentVault.TokenAddress,
				ChainID: chainID,
			}
		}
	}
	curveFactoryPools := loadCurvePools(chainID)

	/**********************************************************************************************
	** Somehow, some vaults are not in the registries, but we still need the tokens data for them.
	** We will add them manually here.
	**********************************************************************************************/
	extraTokens := map[uint64][]string{
		1: {
			`0x34fe2a45D8df28459d7705F37eD13d7aE4382009`, // yvWBTC
			`0xD533a949740bb3306d119CC777fa900bA034cd52`, // CRV - used by yBribe UI
			`0x090185f2135308BaD17527004364eBcC2D37e5F6`, // Spell - used by yBribe UI
			`0xCdF7028ceAB81fA0C6971208e83fa7872994beE5`, // TNT - used by yBribe UI
			`0xba100000625a3754423978a60c9317c58a424e3D`, // BAL - used by yBAL UI
			`0x5c6Ee304399DBdB9C8Ef030aB642B10820DB8F56`, // BALWETH - used by yBAL UI
		},
		10:    {},
		250:   {},
		42161: {},
	}
	for _, tokenAddress := range extraTokens[chainID] {
		tokenAddress := common.HexToAddress(tokenAddress)
		if _, ok := tokenMap[tokenAddress]; !ok {
			updatedTokenMap[tokenAddress] = models.TERC20Token{
				Address: tokenAddress,
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

		/**********************************************************************************************
		** Once everything is setup, we will store each token in the DB. The storage is set as a map
		** of tokenAddress -> TERC20Token. All tokens will be retrievable from the store.Interate()
		** func.
		**********************************************************************************************/
		for _, token := range updatedTokenMap {
			store.AppendInERC20(chainID, token)
			if _, ok := tokenMap[token.Address]; !ok {
				store.StoreERC20(chainID, token)
			}
		}
		tokenMap, _ = store.ListAllERC20(chainID)
	}

	_tokenMap[chainID] = tokenMap
	return tokenMap
}
