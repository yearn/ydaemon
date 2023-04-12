package tokens

import (
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/common/env"
	"github.com/yearn/ydaemon/internal/models"
)

var COIN_PER_CHAIN = map[uint64]models.TERC20Token{}

func init() {
	defaultCoinAddress := common.HexToAddress(`0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE`)
	COIN_PER_CHAIN[1] = models.TERC20Token{
		Address:                   defaultCoinAddress,
		UnderlyingTokensAddresses: []common.Address{},
		Name:                      `Ether`,
		Symbol:                    `ETH`,
		Type:                      `nativeCoin`,
		DisplayName:               `Ether`,
		DisplaySymbol:             `ETH`,
		Description:               `Ether is the native cryptocurrency for the Ethereum blockchain.`,
		Icon:                      env.GITHUB_ICON_BASE_URL + strconv.FormatUint(1, 10) + `/` + defaultCoinAddress.Hex() + `/logo-128.png`,
		Decimals:                  18,
	}
	COIN_PER_CHAIN[10] = models.TERC20Token{
		Address:                   defaultCoinAddress,
		UnderlyingTokensAddresses: []common.Address{},
		Name:                      `Ether`,
		Symbol:                    `ETH`,
		Type:                      `nativeCoin`,
		DisplayName:               `Ether`,
		DisplaySymbol:             `ETH`,
		Description:               `Optimism is a Layer 2 scaling solution for Ethereum.`,
		Icon:                      env.GITHUB_ICON_BASE_URL + strconv.FormatUint(1, 10) + `/` + defaultCoinAddress.Hex() + `/logo-128.png`,
		Decimals:                  18,
	}
	COIN_PER_CHAIN[250] = models.TERC20Token{
		Address:                   defaultCoinAddress,
		UnderlyingTokensAddresses: []common.Address{},
		Name:                      `Fantom`,
		Symbol:                    `FTM`,
		Type:                      `nativeCoin`,
		DisplayName:               `Fantom`,
		DisplaySymbol:             `FTM`,
		Description:               `Fantom is a Layer 2 scaling solution for Ethereum.`,
		Icon:                      env.GITHUB_ICON_BASE_URL + strconv.FormatUint(1, 10) + `/` + defaultCoinAddress.Hex() + `/logo-128.png`,
		Decimals:                  18,
	}
	COIN_PER_CHAIN[42161] = models.TERC20Token{
		Address:                   defaultCoinAddress,
		UnderlyingTokensAddresses: []common.Address{},
		Name:                      `Arbitrum`,
		Symbol:                    `ARB`,
		Type:                      `nativeCoin`,
		DisplayName:               `Arbitrum`,
		DisplaySymbol:             `ARB`,
		Description:               `Arbitrum is a Layer 2 scaling solution for Ethereum.`,
		Icon:                      env.GITHUB_ICON_BASE_URL + strconv.FormatUint(1, 10) + `/` + defaultCoinAddress.Hex() + `/logo-128.png`,
		Decimals:                  18,
	}
}
