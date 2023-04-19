package tokenList

import (
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/common/bigNumber"
	"github.com/yearn/ydaemon/common/env"
	"github.com/yearn/ydaemon/common/helpers"
	"github.com/yearn/ydaemon/internal/models"
)

func BuildTokenList(chainID uint64) {
	if len(WidoTokenList.Tokens) > 0 {
		for _, token := range WidoTokenList.Tokens {
			tokenAddress := common.HexToAddress(token.Address)
			if uint64(token.ChainID) != chainID {
				continue
			}
			if value, exist := GetTokenFromList(chainID, tokenAddress); exist {
				if helpers.Contains(value.SupportedZaps, models.Wido) {
					continue
				}
				value.SupportedZaps = append(value.SupportedZaps, models.Wido)
				setTokenFromList(chainID, value)
				continue
			}
			setTokenFromList(chainID, models.TYearnTokenListToken{
				TTokenListToken: models.TTokenListToken{
					ChainID:  int(token.ChainID),
					Address:  tokenAddress.Hex(),
					Name:     token.Name,
					Symbol:   token.Symbol,
					Decimals: token.Decimals,
					LogoURI:  token.LogoURI,
				},
				Balance:       bigNumber.NewInt(0),
				SupportedZaps: []models.SupportedZap{models.Wido},
			})
		}
	}

	if len(PortalsTokenList.Tokens) > 0 {
		for _, token := range PortalsTokenList.Tokens {
			tokenAddress := common.HexToAddress(token.Address)
			if uint64(token.ChainID) != chainID {
				continue
			}
			if value, exist := GetTokenFromList(chainID, tokenAddress); exist {
				if helpers.Contains(value.SupportedZaps, models.Portals) {
					continue
				}
				value.SupportedZaps = append(value.SupportedZaps, models.Portals)
				setTokenFromList(chainID, value)
				continue
			}
			setTokenFromList(chainID, models.TYearnTokenListToken{
				TTokenListToken: models.TTokenListToken{
					ChainID:  int(token.ChainID),
					Address:  tokenAddress.Hex(),
					Name:     token.Name,
					Symbol:   token.Symbol,
					Decimals: token.Decimals,
					LogoURI:  token.LogoURI,
				},
				Balance:       bigNumber.NewInt(0),
				SupportedZaps: []models.SupportedZap{models.Portals},
			})
		}
	}

	if chainID == 1 {
		setSupportedByCowSwap(chainID)
	}

	chainIDStr := strconv.FormatUint(chainID, 10)
	saveTokensListToJSON(MapTokenList(chainID), env.BASE_DATA_PATH+`/tokensList/`+chainIDStr+`.json`)
}
