package prices

import (
	"math/big"
	"strconv"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/common/bigNumber"
	"github.com/yearn/ydaemon/common/contracts"
	"github.com/yearn/ydaemon/common/env"
	"github.com/yearn/ydaemon/common/ethereum"
	"github.com/yearn/ydaemon/common/helpers"
	"github.com/yearn/ydaemon/common/store"
	"github.com/yearn/ydaemon/internal/multicalls"
)

func FetchPricesOnBlock(chainID uint64, blockNumber uint64, tokenList []common.Address) map[common.Address]*bigNumber.Int {
	newPriceMap := make(map[common.Address]*bigNumber.Int)

	/**********************************************************************************************
	** The first step is to prepare the multicall, connecting to the multicall instance and
	** preparing the array of calls to send. All calls for all tokens will be send in a single
	** multicall and will later be accessible via a concatened string `tokenAddress + methodName`.
	**********************************************************************************************/
	calls := []ethereum.Call{}
	for _, tokenAddress := range tokenList {
		if tokenPrice, ok := store.GetBlockPrice(chainID, blockNumber, tokenAddress); ok {
			newPriceMap[tokenAddress] = tokenPrice
			continue
		}
		calls = append(calls, multicalls.GetPriceUsdcRecommendedCall(tokenAddress.Hex(), env.LENS_ADDRESSES[chainID], tokenAddress))
	}

	/**********************************************************************************************
	** Then we can proceed the responses. We loop over the responses and check if the price is
	** available. If it is, we add it to the map. If it's not, we try to fetch it from an external
	** API.
	**********************************************************************************************/
	response := multicalls.Perform(chainID, calls, big.NewInt(int64(blockNumber)))
	itemsToUpsert := []store.DBHistoricalPrice{}
	for _, tokenAddress := range tokenList {
		rawTokenPrice := response[tokenAddress.Hex()+`getPriceUsdcRecommended`]
		if len(rawTokenPrice) == 0 {
			continue
		}
		tokenPrice := bigNumber.SetInt(rawTokenPrice[0].(*big.Int))
		if tokenPrice.IsZero() {
			continue
		}
		newPriceMap[tokenAddress] = helpers.DecodeBigInt(rawTokenPrice)

		/******************************************************************************************
		** To save some process, we will batch the saving to the database in one call.
		** The following code is creating the upsert array.
		******************************************************************************************/
		DBbaseSchema := store.DBBaseSchema{
			UUID:    store.GetUUID(strconv.FormatUint(chainID, 10) + strconv.FormatUint(blockNumber, 10) + tokenAddress.Hex()),
			Block:   blockNumber,
			ChainID: chainID,
		}
		humanizedPrice, _ := helpers.ToNormalizedAmount(newPriceMap[tokenAddress], 6).Float64()
		itemsToUpsert = append(itemsToUpsert, store.DBHistoricalPrice{
			DBBaseSchema:   DBbaseSchema,
			Token:          tokenAddress.Hex(),
			Price:          newPriceMap[tokenAddress].String(),
			HumanizedPrice: humanizedPrice,
		})
		store.AppendToHistoricalPrice(chainID, blockNumber, tokenAddress, newPriceMap[tokenAddress])
	}

	store.StoreManyHistoricalPrice(itemsToUpsert)
	return newPriceMap
}

func GetPricesOnBlock(chainID uint64, blockNumber uint64, tokenAddress common.Address) (*bigNumber.Int, bool) {
	tokenPrice, ok := store.GetBlockPrice(chainID, blockNumber, tokenAddress)
	if ok {
		return tokenPrice, true
	}

	client := ethereum.GetRPC(chainID)
	contractAddress := env.LENS_ADDRESSES[chainID]
	oracleContract, _ := contracts.NewOracleCaller(contractAddress, client)
	tokenPriceFromOracle, err := oracleContract.GetPriceUsdcRecommended(&bind.CallOpts{BlockNumber: big.NewInt(int64(blockNumber))}, tokenAddress)
	if err != nil {
		return bigNumber.NewInt(0), false
	}
	store.StoreHistoricalPrice(chainID, blockNumber, tokenAddress, bigNumber.NewInt(0).Set(tokenPriceFromOracle))
	return bigNumber.NewInt(0).Set(tokenPriceFromOracle), true
}
