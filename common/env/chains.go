package env

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/internal/models"
)

type TChainCurve struct {
	RegistryAddress common.Address
	FactoryAddress  common.Address
	FactoryURIs     []string
	PoolsURIs       []string
}
type TChain struct {
	ID                     uint64
	RpcURI                 string
	SubgraphURI            string
	MaxBlockRange          uint64
	MaxBatchSize           int
	LensContract           TContractData
	MulticallContract      TContractData
	StackingRewardContract TContractData
	YBribeV3Contract       TContractData
	PartnerContract        TContractData
	Coin                   models.TERC20Token
	Registries             []TContractData
	ExtraVaults            []models.TVaultsFromRegistry
	BlacklistedVaults      []common.Address
	ExtraTokens            []common.Address
	IgnoredTokens          []common.Address
	Curve                  TChainCurve
}

var DEFAULT_COIN_ADDRESS = common.HexToAddress(`0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE`)

// SUPPORTED_CHAIN_IDS is the list of supported chain IDs
var SUPPORTED_CHAIN_IDS = []uint64{1, 10, 137, 250, 8453, 42161}

// CHAINS is the list of supported chains
var CHAINS = map[uint64]TChain{
	1:     ETHEREUM,
	10:    OPTIMISM,
	137:   POLYGON,
	250:   FANTOM,
	8453:  BASE,
	42161: ARBITRUM,
}
