package tokenList

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/yearn/ydaemon/common/helpers"
	"github.com/yearn/ydaemon/common/logs"
	"github.com/yearn/ydaemon/internal/models"
)

type QuotePriceResponse struct {
	Price float64 `json:"price"`
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
	body, err := io.ReadAll(resp.Body)
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
