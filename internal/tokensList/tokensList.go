package tokensList

import (
	"strconv"
	"time"

	"github.com/yearn/ydaemon/common/bigNumber"
	"github.com/yearn/ydaemon/common/env"
	"github.com/yearn/ydaemon/common/helpers"
	"github.com/yearn/ydaemon/common/logs"
	"github.com/yearn/ydaemon/common/types/common"
)

func setSupportedByCowSwap(chainID uint64) {
	tokenMap := MapTokenList(chainID)
	index := 0
	for address, token := range tokenMap {
		if helpers.Contains(token.SupportedZaps, Cowswap) {
			index++
			continue
		}
		logs.Info(`Testing token: ` + token.Symbol + ` (` + strconv.Itoa(index+1) + `/` + strconv.Itoa(len(tokenMap)) + `)`)
		index++

		time.Sleep(400)
		_, errReason, err := tryQuote(address.Hex())
		if err == nil || (err == nil && errReason == `NoLiquidity`) {
			if value, ok := tokenMap[address]; ok {
				if !helpers.Contains(tokenMap[address].SupportedZaps, Cowswap) {
					value.SupportedZaps = append(value.SupportedZaps, Cowswap)
					setTokenFromList(chainID, value)
				}
			}
		}
	}
}

func BuildTokenList(chainID uint64) {
	lastUpdate := getLastUpdate(chainID)
	if !lastUpdate.IsZero() && lastUpdate.Before(time.Now().Add(-time.Hour*24)) {
		return
	}
	logs.Info(`Reloading tokenLists...`)

	supportedTokenMap := make(map[string]DefaultTokenListToken)

	if len(WidoTokenList.Tokens) > 0 {
		for _, token := range WidoTokenList.Tokens {
			tokenAddress := common.HexToAddress(token.Address)
			if _, exists := supportedTokenMap[tokenAddress.Hex()]; !exists {
				continue
			}
			if uint64(token.ChainID) != chainID {
				continue
			}
			if value, exist := GetTokenFromList(chainID, tokenAddress); exist {
				if helpers.Contains(value.SupportedZaps, Wido) {
					continue
				}
				value.SupportedZaps = append(value.SupportedZaps, Wido)
				setTokenFromList(chainID, value)
				continue
			}
			setTokenFromList(chainID, YTokenFromList{
				ChainID:       uint64(token.ChainID),
				Address:       tokenAddress.Hex(),
				Name:          token.Name,
				Symbol:        token.Symbol,
				Decimals:      token.Decimals,
				LogoURI:       token.LogoURI,
				Balance:       bigNumber.NewInt(0),
				SupportedZaps: []SupportedZap{Wido},
			})
		}
	}

	if len(PortalsTokenList.Tokens) > 0 {
		for _, token := range PortalsTokenList.Tokens {
			tokenAddress := common.HexToAddress(token.Address)
			if _, exists := supportedTokenMap[tokenAddress.Hex()]; !exists {
				continue
			}
			if uint64(token.ChainID) != chainID {
				continue
			}
			if value, exist := GetTokenFromList(chainID, tokenAddress); exist {
				if helpers.Contains(value.SupportedZaps, Portals) {
					continue
				}
				value.SupportedZaps = append(value.SupportedZaps, Portals)
				setTokenFromList(chainID, value)
				continue
			}
			setTokenFromList(chainID, YTokenFromList{
				ChainID:       uint64(token.ChainID),
				Address:       tokenAddress.Hex(),
				Name:          token.Name,
				Symbol:        token.Symbol,
				Decimals:      token.Decimals,
				LogoURI:       token.LogoURI,
				Balance:       bigNumber.NewInt(0),
				SupportedZaps: []SupportedZap{Portals},
			})
		}
	}

	if chainID == 1 {
		setSupportedByCowSwap(chainID)
	}

	if lastUpdate.Before(time.Now().Add(-time.Hour * 24 * 7)) {
		logs.Warning(`Please update the tokenList json files`)
	}
	chainIDStr := strconv.FormatUint(chainID, 10)
	saveTokensListToJSON(MapTokenList(chainID), env.BASE_DATA_PATH+`/tokensList/`+chainIDStr+`.json`)
}
