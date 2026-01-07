package env

import (
	"math"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/internal/models"
)

var KATANA = TChain{
	ID:              747474,
	RpcURI:          `https://rpc.katana.network`,
	SubgraphURI:     ``,
	EtherscanURI:    `https://api.etherscan.io/v2/api`,
	MaxBlockRange:   100_000_000,
	MaxBatchSize:    math.MaxInt64,
	AvgBlocksPerDay: 86_400,
	CanUseWebsocket: true, //CHECK!
	LensContract:    TContractData{},
	// // this one is multicall1
	// MulticallContract: TContractData{
	// 	Address: common.HexToAddress(`0x1F4c1E0afBeb5b5B86d7722549274434b29884F6`),
	// 	Block:   1295447,
	// },
	//this one is multicall2
	MulticallContract: TContractData{
		Address: common.HexToAddress(`0xe9128E672bc08E12deb1C2048E9f91e6D6E08e74`),
		Block:   1295938,
	},
	// // this one is multicall3
	// Multicall3Contract: TContractData{
	// 	Address: common.HexToAddress(`0xcA11bde05977b3631167028862bE2a173976CA11`),
	// 	Block:    1898013,
	// },
	PartnerContract: TContractData{},
	Coin: models.TERC20Token{
		Address:                   DEFAULT_COIN_ADDRESS, // unclear if this should be KAT or ETH (which is the gas token)
		UnderlyingTokensAddresses: []common.Address{},
		Type:                      models.TokenTypeNative,
		Name:                      `Ether`, // unclear if this should be KAT or ETH (which is the gas token)
		Symbol:                    `ETH`,   // unclear if this should be KAT or ETH (which is the gas token)
		DisplayName:               `Ether`, // unclear if this should be KAT or ETH (which is the gas token)
		DisplaySymbol:             `ETH`,   // unclear if this should be KAT or ETH (which is the gas token)
		Description:               `Katana is a DeFi chain for deep liquidity and high yield.`,
		Icon:                      BASE_ASSET_URL + strconv.FormatUint(747474, 10) + `/` + strings.ToLower(DEFAULT_COIN_ADDRESS.Hex()) + `/logo-128.png`,
		Decimals:                  18,
		ChainID:                   747474,
	},
	Registries: []TContractData{
		{
			Address: common.HexToAddress("0xd40ecF29e001c76Dcc4cC0D9cd50520CE845B038"),
			Version: 4,
			Block:   2852031,
			Label:   `YEARN`,
		},
		{
			Address: common.HexToAddress("0x770D0d1Fb036483Ed4AbB6d53c1C88fb277D812F"),
			Version: 5,
			Block:   2236950,
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
