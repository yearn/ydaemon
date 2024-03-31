package env

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/internal/models"
)

type TChainCurve struct {
	RegistryAddress common.Address
	FactoryAddress  common.Address
	PoolsURIs       []string
}
type TChainExtraURI struct {
	GammaMerklURI      string
	GammaHypervisorURI []string
}
type TChain struct {
	ID                     uint64
	RpcURI                 string
	SubgraphURI            string
	EtherscanURI           string
	MaxBlockRange          uint64
	MaxBatchSize           uint64
	AvgBlocksPerDay        int
	LensContract           TContractData
	MulticallContract      TContractData
	StackingRewardContract TContractData
	YBribeV3Contract       TContractData
	PartnerContract        TContractData
	APROracleContract      TContractData
	Coin                   models.TERC20Token
	Registries             []TContractData
	ExtraVaults            []models.TVaultsFromRegistry
	BlacklistedVaults      []common.Address
	ExtraTokens            []common.Address
	IgnoredTokens          []common.Address
	Curve                  TChainCurve
	ExtraURI               TChainExtraURI
}

var DEFAULT_COIN_ADDRESS = common.HexToAddress(`0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE`)

// CHAINS is the list of supported chains
var CHAINS = map[uint64]TChain{
	1:     ETHEREUM,
	10:    OPTIMISM,
	137:   POLYGON,
	250:   FANTOM,
	8453:  BASE,
	42161: ARBITRUM,
}

var SUPPORTED_CHAIN_IDS = []uint64{}

func init() {
	for k := range CHAINS {
		SUPPORTED_CHAIN_IDS = append(SUPPORTED_CHAIN_IDS, k)
	}
}
