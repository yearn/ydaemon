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
		Type:                      models.TokenTypeNative,
		Name:                      `Ether`,
		Symbol:                    `ETH`,
		DisplayName:               `Ether`,
		DisplaySymbol:             `ETH`,
		Description:               `Ether is the native cryptocurrency for the Ethereum blockchain.`,
		Icon:                      env.BASE_ASSET_URL + strconv.FormatUint(1, 10) + `/` + defaultCoinAddress.Hex() + `/logo-128.png`,
		Decimals:                  18,
		ChainID:                   1,
	}
	COIN_PER_CHAIN[10] = models.TERC20Token{
		Address:                   defaultCoinAddress,
		UnderlyingTokensAddresses: []common.Address{},
		Type:                      models.TokenTypeNative,
		Name:                      `Ether`,
		Symbol:                    `ETH`,
		DisplayName:               `Ether`,
		DisplaySymbol:             `ETH`,
		Description:               `Optimism is a Layer 2 scaling solution for Ethereum.`,
		Icon:                      env.BASE_ASSET_URL + strconv.FormatUint(1, 10) + `/` + defaultCoinAddress.Hex() + `/logo-128.png`,
		Decimals:                  18,
		ChainID:                   10,
	}
	COIN_PER_CHAIN[250] = models.TERC20Token{
		Address:                   defaultCoinAddress,
		UnderlyingTokensAddresses: []common.Address{},
		Type:                      models.TokenTypeNative,
		Name:                      `Fantom`,
		Symbol:                    `FTM`,
		DisplayName:               `Fantom`,
		DisplaySymbol:             `FTM`,
		Description:               `Fantom is a Layer 2 scaling solution for Ethereum.`,
		Icon:                      env.BASE_ASSET_URL + strconv.FormatUint(1, 10) + `/` + defaultCoinAddress.Hex() + `/logo-128.png`,
		Decimals:                  18,
		ChainID:                   250,
	}
	COIN_PER_CHAIN[42161] = models.TERC20Token{
		Address:                   defaultCoinAddress,
		UnderlyingTokensAddresses: []common.Address{},
		Type:                      models.TokenTypeNative,
		Name:                      `Arbitrum`,
		Symbol:                    `ARB`,
		DisplayName:               `Arbitrum`,
		DisplaySymbol:             `ARB`,
		Description:               `Arbitrum is a Layer 2 scaling solution for Ethereum.`,
		Icon:                      env.BASE_ASSET_URL + strconv.FormatUint(1, 10) + `/` + defaultCoinAddress.Hex() + `/logo-128.png`,
		Decimals:                  18,
		ChainID:                   42161,
	}
}
