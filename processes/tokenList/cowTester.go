package tokenList

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/yearn/ydaemon/common/helpers"
	"github.com/yearn/ydaemon/common/logs"
	"github.com/yearn/ydaemon/internal/models"
)

type QuoteRequest struct {
	SellToken           string `json:"sellToken"`
	BuyToken            string `json:"buyToken"`
	Receiver            string `json:"receiver"`
	AppData             string `json:"appData"`
	PartiallyFillable   bool   `json:"partiallyFillable"`
	SellTokenBalance    string `json:"sellTokenBalance"`
	BuyTokenBalance     string `json:"buyTokenBalance"`
	From                string `json:"from"`
	PriceQuality        string `json:"priceQuality"`
	SigningScheme       string `json:"signingScheme"`
	OnchainOrder        bool   `json:"onchainOrder"`
	Kind                string `json:"kind"`
	SellAmountBeforeFee string `json:"sellAmountBeforeFee"`
}
type QuotePriceResponse struct {
	Price float64 `json:"price"`
}
type QuoteResponse struct {
	Quote struct {
		SellToken         string `json:"sellToken"`
		BuyToken          string `json:"buyToken"`
		Receiver          string `json:"receiver"`
		SellAmount        string `json:"sellAmount"`
		BuyAmount         string `json:"buyAmount"`
		ValidTo           int    `json:"validTo"`
		AppData           string `json:"appData"`
		FeeAmount         string `json:"feeAmount"`
		Kind              string `json:"kind"`
		PartiallyFillable bool   `json:"partiallyFillable"`
		SellTokenBalance  string `json:"sellTokenBalance"`
		BuyTokenBalance   string `json:"buyTokenBalance"`
		SigningScheme     string `json:"signingScheme"`
	} `json:"quote"`
	From       string `json:"from"`
	Expiration string `json:"expiration"`
	ID         int    `json:"id"`
}
type QuoteResponseError struct {
	ErrorType   string `json:"errorType"`
	Description string `json:"description"`
}

func testCowSwapSupport(fromToken string) (string, string, error) {
	resp, err := http.Get(`https://api.cow.fi/mainnet/api/v1/token/` + fromToken + `/native_price`)
	if err != nil {
		logs.Error(err)
		return ``, ``, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logs.Error(err)
		return ``, ``, err
	}

	if resp.StatusCode == 200 {
		var quote QuotePriceResponse
		if err := json.Unmarshal(body, &quote); err != nil {
			logs.Error(err)
			return ``, ``, err
		}
		return strconv.FormatFloat(quote.Price, 'f', -1, 64), ``, nil
	} else {
		var quote QuoteResponseError
		if err := json.Unmarshal(body, &quote); err != nil {
			logs.Error(err)
			return ``, ``, err
		}
		return quote.Description, quote.ErrorType, errors.New(quote.Description)
	}
}

func setSupportedByCowSwap(chainID uint64) {
	tokenMap := MapTokenList(chainID)
	index := 0
	for address, token := range tokenMap {
		if helpers.Contains(token.SupportedZaps, models.Cowswap) {
			index++
			continue
		}
		logs.Info(`Testing token: ` + token.Symbol + ` (` + strconv.Itoa(index+1) + `/` + strconv.Itoa(len(tokenMap)) + `)`)
		index++

		time.Sleep(400)
		_, errReason, err := testCowSwapSupport(address.Hex())
		if err == nil || (err == nil && errReason == `NoLiquidity`) {
			if value, ok := tokenMap[address]; ok {
				if !helpers.Contains(tokenMap[address].SupportedZaps, models.Cowswap) {
					value.SupportedZaps = append(value.SupportedZaps, models.Cowswap)
					setTokenFromList(chainID, value)
				}
			}
		}
	}
}
