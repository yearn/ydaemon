package tokensList

import (
	"strconv"
	"time"

	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/common/bigNumber"
	"github.com/yearn/ydaemon/common/helpers"
	"github.com/yearn/ydaemon/common/logs"
)

func setSupportedByCowSwap(chainID uint64) {
	tokenMap := MapTokenList(chainID)
	index := 0
	for address, token := range tokenMap {
		if helpers.Contains(token.SupportedZaps, Cowswap) {
			index++
			continue
		}
		logs.Success(`Testing token: ` + token.Symbol + ` (` + strconv.Itoa(index+1) + `/` + strconv.Itoa(len(tokenMap)) + `)`)
		index++

		time.Sleep(400)
		_, errReason, err := tryQuote(address.Hex())
		if err == nil || (err == nil && errReason == `NoLiquidity`) {
			if value, ok := tokenMap[address]; ok {
				if !helpers.Contains(tokenMap[address].SupportedZaps, Cowswap) {
					tokenMap[address].SupportedZaps = append(value.SupportedZaps, Cowswap)
				}
			}
		}
	}
}

func BuildTokenList(chainID uint64) {
	supportedTokenMap := make(map[string]DefaultTokenListToken)
	for _, token := range AggregatedTokenList.Tokens {
		supportedTokenMap[token.Address] = token
	}

	minCountToInclude := 2
	if chainID == 1 {
		minCountToInclude = 3
	}

	if len(WidoTokenList.Tokens) > 0 {
		for _, token := range WidoTokenList.Tokens {
			tokenAddress := ethcommon.HexToAddress(token.Address)
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
			setTokenFromList(chainID, &YTokenFromList{
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
			tokenAddress := ethcommon.HexToAddress(token.Address)
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
			setTokenFromList(chainID, &YTokenFromList{
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

	if len(AggregatedTokenList.Tokens) > 0 {
		for _, token := range AggregatedTokenList.Tokens {
			tokenAddress := ethcommon.HexToAddress(token.Address)
			if token.Count < minCountToInclude {
				continue
			}
			if _, exist := GetTokenFromList(chainID, tokenAddress); exist {
				continue
			}
			if uint64(token.ChainID) != chainID {
				continue
			}
			setTokenFromList(chainID, &YTokenFromList{
				ChainID:       uint64(token.ChainID),
				Address:       tokenAddress.Hex(),
				Name:          token.Name,
				Symbol:        token.Symbol,
				Decimals:      token.Decimals,
				LogoURI:       token.LogoURI,
				Balance:       bigNumber.NewInt(0),
				SupportedZaps: []SupportedZap{},
			})
		}
	}

	if chainID == 1 {
		setSupportedByCowSwap(chainID)
	}
	// saveTokensListToJSON(chainID, MapTokenList(chainID))
}
