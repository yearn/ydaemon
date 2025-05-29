package env

import (
	"math"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/internal/models"
)

var SONIC = TChain{
	ID:              146,
	RpcURI:          `https://sonic.drpc.org`,
	SubgraphURI:     ``,
	EtherscanURI:    `https://api.etherscan.io/v2/api`,
	MaxBlockRange:   100_000_000,
	MaxBatchSize:    math.MaxInt64,
	AvgBlocksPerDay: 45_000,
	CanUseWebsocket: true,
	LensContract:    TContractData{},
	MulticallContract: TContractData{
		Address: common.HexToAddress(`0xca11bde05977b3631167028862be2a173976ca11`),
		Block:   60,
	},
	PartnerContract: TContractData{},
	Coin: models.TERC20Token{
		Address:                   DEFAULT_COIN_ADDRESS,
		UnderlyingTokensAddresses: []common.Address{},
		Type:                      models.TokenTypeNative,
		Name:                      `Sonic`,
		Symbol:                    `S`,
		DisplayName:               `Sonic`,
		DisplaySymbol:             `S`,
		Description:               `Sonic is a fast Layer 1.`,
		Icon:                      BASE_ASSET_URL + strconv.FormatUint(250, 10) + `/` + DEFAULT_COIN_ADDRESS.Hex() + `/logo-128.png`,
		Decimals:                  18,
		ChainID:                   146,
	},
	Registries: []TContractData{
		{
			Address: common.HexToAddress("0xd40ecF29e001c76Dcc4cC0D9cd50520CE845B038"),
			Version: 4,
			Block:   4459165,
			Label:   `YEARN`,
		},
		{
			Address: common.HexToAddress("0x770D0d1Fb036483Ed4AbB6d53c1C88fb277D812F"),
			Version: 5,
			Block:   307787,
			Tag:     `STEALTH`,
			Label:   `PUBLIC_ERC4626`,
		},
	},
	ExtraVaults:       []models.TVaultsFromRegistry{},
	BlacklistedVaults: []common.Address{},
	ExtraTokens:       []common.Address{},
	IgnoredTokens:     []common.Address{},
	Curve:             TChainCurve{},
	ExtraURI:          TChainExtraURI{},
}
