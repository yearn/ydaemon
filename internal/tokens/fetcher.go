package tokens

import (
	"math"
	"strconv"
	"strings"
	"sync"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/common/contracts"
	"github.com/yearn/ydaemon/common/env"
	"github.com/yearn/ydaemon/common/ethereum"
	"github.com/yearn/ydaemon/common/helpers"
	"github.com/yearn/ydaemon/common/store"
	"github.com/yearn/ydaemon/common/traces"
	"github.com/yearn/ydaemon/common/types/common"
	"github.com/yearn/ydaemon/internal/meta"
	"github.com/yearn/ydaemon/internal/utils"
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
func fetchBasicInformations(chainID uint64, tokens []ethcommon.Address) (tokenList []*TERC20Token) {
	/**********************************************************************************************
	** The first step is to prepare the multicall, connecting to the multicall instance and
	** preparing the array of calls to send. All calls for all tokens will be send in a single
	** multicall and will later be accessible via a concatened string `tokenAddress + methodName`.
	**********************************************************************************************/
	caller := ethereum.MulticallClientForChainID[chainID]
	calls := []ethereum.Call{}
	for _, token := range tokens {
		calls = append(calls, getName(token.String(), token))
		calls = append(calls, getSymbol(token.String(), token))
		calls = append(calls, getDecimals(token.String(), token))
		calls = append(calls, getToken(token.String(), token))
		calls = append(calls, getPoolFromLpToken(token.String(), env.CURVE_REGISTRY_ADDRESSES[chainID].ToAddress(), token))
		calls = append(calls, getCompoundUnderlying(token.String(), token))
		calls = append(calls, getAaveV1Underlying(token.String(), token))
		calls = append(calls, getAaveV2Underlying(token.String(), token))
	}

	/**********************************************************************************************
	** Regular fix for Fantom's RPC, which limit the number of calls in a multicall to a very low
	** number. We split the multicall in multiple calls of 3 calls each.
	** Otherwise, we just send the multicall as is.
	**********************************************************************************************/
	maxBatch := math.MaxInt64
	if chainID == 250 {
		maxBatch = 3
	}

	/**********************************************************************************************
	** Then we can proceed the responses. We will create a new relatedTokensList to be able to know
	** which token to fetch then (ex: for aDAI, we also need to fetch the DAI token).
	** Nb: A special case is for Ethereum coin, which is defaulted as address 0xEeeee....EEeE.
	**********************************************************************************************/
	relatedTokensList := []ethcommon.Address{}
	response := caller.ExecuteByBatch(calls, maxBatch, nil)
	for _, token := range tokens {
		if token == ethcommon.HexToAddress(`0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE`) {
			tokenList = append(tokenList, &TERC20Token{
				Address:  token,
				Name:     `Ethereum`,
				Symbol:   `ETH`,
				Decimals: 18,
			})
			continue
		}
		rawName := response[token.String()+`name`]
		rawSymbol := response[token.String()+`symbol`]
		rawDecimals := response[token.String()+`decimals`]
		rawYearnVaultToken := response[token.String()+`token`]
		rawPoolFromLpToken := response[token.String()+`get_pool_from_lp_token`]
		rawCUnderlying := response[token.String()+`underlying`]
		rawAV1Underlying := response[token.String()+`underlyingAssetAddress`]
		rawAV2Underlying := response[token.String()+`UNDERLYING_ASSET_ADDRESS`]

		/******************************************************************************************
		** Preparing our new ERC20Token object
		******************************************************************************************/
		newToken := &TERC20Token{
			Address:  token,
			Name:     rawName[0].(string),
			Symbol:   rawSymbol[0].(string),
			Decimals: uint64(rawDecimals[0].(uint8)),
			Icon:     env.GITHUB_ICON_BASE_URL + strconv.FormatUint(chainID, 10) + `/` + token.Hex() + `/logo-128.png`,
		}

		metaTokenName := newToken.Name
		metaTokenSymbol := newToken.Symbol
		metaTokenDescription := ``
		if tokenFromMeta, ok := meta.GetMetaToken(chainID, common.FromAddress(token)); ok {
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
		** If so, we set the token type to `Yeearn Vault`, we fetch the Coins from the pool and we
		** add the coins to the newToken UnderlyingTokensAddresses.
		** We can also add the coins to the relatedTokensList, so we can fetch their information
		** later.
		******************************************************************************************/
		isYearnVault := rawYearnVaultToken != nil && rawYearnVaultToken[0].(ethcommon.Address) != ethcommon.Address{}
		if isYearnVault {
			newToken.Type = `Yearn Vault`
			coin := rawYearnVaultToken[0].(ethcommon.Address)
			if (coin != ethcommon.Address{}) {
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
		isCurveLpToken := rawPoolFromLpToken != nil && rawPoolFromLpToken[0].(ethcommon.Address) != ethcommon.Address{}
		if isCurveLpToken {
			newToken.Type = `Curve LP`
			curvePoolCaller, _ := contracts.NewCurvePoolRegistryCaller(env.CURVE_REGISTRY_ADDRESSES[chainID].ToAddress(), caller.Client)
			poolCoins, _ := curvePoolCaller.GetCoins(&bind.CallOpts{}, rawPoolFromLpToken[0].(ethcommon.Address))
			for _, coin := range poolCoins {
				if (coin != ethcommon.Address{}) {
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
		isCToken := rawCUnderlying != nil && rawCUnderlying[0].(ethcommon.Address) != ethcommon.Address{}
		if isCToken {
			newToken.Type = `Compound`
			coin := rawCUnderlying[0].(ethcommon.Address)
			if (coin != ethcommon.Address{}) {
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
		isAV1Token := rawAV1Underlying != nil && rawAV1Underlying[0].(ethcommon.Address) != ethcommon.Address{}
		if isAV1Token {
			newToken.Type = `AAVE V1`
			coin := rawAV1Underlying[0].(ethcommon.Address)
			if (coin != ethcommon.Address{}) {
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
		isAV2Token := rawAV2Underlying != nil && rawAV2Underlying[0].(ethcommon.Address) != ethcommon.Address{}
		if isAV2Token {
			newToken.Type = `AAVE V2`
			coin := rawAV2Underlying[0].(ethcommon.Address)
			if (coin != ethcommon.Address{}) {
				relatedTokensList = append(relatedTokensList, coin)
				newToken.UnderlyingTokensAddresses = append(newToken.UnderlyingTokensAddresses, coin)
			}
		}

		tokenList = append(tokenList, newToken)
	}

	if len(relatedTokensList) > 0 {
		tokenList = append(tokenList, fetchBasicInformations(chainID, helpers.UniqueArrayAddress(relatedTokensList))...)
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
	tokenMap map[ethcommon.Address]*TERC20Token,
) map[ethcommon.Address]*TERC20Token {
	newMap := make(map[ethcommon.Address]*TERC20Token)
	tokenMapAddresses := []ethcommon.Address{}
	for tokenAddress := range tokenMap {
		tokenMapAddresses = append(tokenMapAddresses, tokenAddress)
	}

	newtokenMap := fetchBasicInformations(chainID, tokenMapAddresses)
	for _, token := range newtokenMap {
		newMap[token.Address] = token
	}

	return newMap
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
	vaults map[ethcommon.Address]utils.TVaultsFromRegistry,
) map[ethcommon.Address]*TERC20Token {
	trace := traces.Init(`app.indexer.tokens.multicall_data`).
		SetTag(`chainID`, strconv.FormatUint(chainID, 10)).
		SetTag(`rpcURI`, ethereum.GetRPCURI(chainID)).
		SetTag(`entity`, `tokens`).
		SetTag(`subsystem`, `daemon`)
	defer trace.Finish()

	/**********************************************************************************************
	** First, try to retrieve the list of tokens from the database to exclude the one existing
	** from the upcoming calls
	**********************************************************************************************/
	tokenMap := make(map[ethcommon.Address]*TERC20Token)
	store.Iterate(chainID, store.TABLES.TOKENS, &tokenMap)

	/**********************************************************************************************
	** From the vault registry we have the first batch of tokens: the vault tokens and the
	** underlying tokens for theses vaults. In order to proceed, we will create a map of ERC20
	** tokens, which will be stored in the DB. This action should only be performed once per token
	** per chainID as the token information are not expected to change and will be retrieve on
	** subsequent reboots.
	**********************************************************************************************/
	updatedTokenMap := make(map[ethcommon.Address]*TERC20Token)
	for _, currentVault := range vaults {
		if _, ok := tokenMap[currentVault.VaultsAddress]; !ok {
			updatedTokenMap[currentVault.VaultsAddress] = &TERC20Token{
				Address: currentVault.VaultsAddress,
			}
		}

		if _, ok := tokenMap[currentVault.TokenAddress]; !ok {
			updatedTokenMap[currentVault.TokenAddress] = &TERC20Token{
				Address: currentVault.TokenAddress,
			}
		}
	}

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
		},
		10:    {},
		250:   {},
		42161: {},
	}
	for _, tokenAddress := range extraTokens[chainID] {
		tokenAddress := ethcommon.HexToAddress(tokenAddress)
		if _, ok := tokenMap[tokenAddress]; !ok {
			updatedTokenMap[tokenAddress] = &TERC20Token{
				Address: tokenAddress,
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
		updatedTokenMap = findAllTokens(chainID, updatedTokenMap)

		/**********************************************************************************************
		** Once everything is setup, we will store each token in the DB. The storage is set as a map
		** of tokenAddress -> TERC20Token. All tokens will be retrievable from the store.Interate()
		** func.
		**********************************************************************************************/
		wg := sync.WaitGroup{}
		wg.Add(len(updatedTokenMap))
		for _, token := range updatedTokenMap {
			go func(_token *TERC20Token) {
				defer wg.Done()
				store.SaveInDB(
					chainID,
					store.TABLES.TOKENS,
					_token.Address.String(),
					_token,
				)
			}(token)
		}
		wg.Wait()
		store.Iterate(chainID, store.TABLES.TOKENS, &tokenMap)
	}

	_tokenMap[chainID] = tokenMap
	return tokenMap
}
