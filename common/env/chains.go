package env

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/internal/models"
)

/**************************************************************************************************
** TExtraStakingContracts represents additional staking contracts that may not be detected through
** the standard registry scanning process. This structure allows for manual inclusion of staking
** opportunities in the system.
**
** @field VaultAddress The address of the vault associated with the staking contract
** @field StakingAddress The address of the staking contract itself
** @field Tag A label to categorize the staking contract (e.g., "JUICED")
**************************************************************************************************/
type TExtraStakingContracts struct {
	VaultAddress   common.Address
	StakingAddress common.Address
	Tag            string
}

/**************************************************************************************************
** TChainCurve contains Curve protocol specific addresses and endpoints for a particular chain.
** Curve is a major DeFi protocol that Yearn integrates with, requiring specific configuration.
**
** @field RegistryAddress The address of the Curve registry contract
** @field FactoryAddress The address of the Curve factory contract
** @field PoolsURIs URLs for fetching Curve pool information
** @field GaugesURI URL for fetching Curve gauge information
**************************************************************************************************/
type TChainCurve struct {
	RegistryAddress common.Address
	FactoryAddress  common.Address
	PoolsURIs       []string
	GaugesURI       string
}

/**************************************************************************************************
** TChainExtraURI contains additional URIs for external protocols that Yearn integrates with.
** These endpoints provide data for specialized integrations like Gamma, Pendle, etc.
**
** @field GammaMerklURI URL for Gamma Merkl protocol data
** @field GammaHypervisorURI URLs for Gamma Hypervisor protocol data
** @field PendleCoreURI URL for Pendle Core protocol data
**************************************************************************************************/
type TChainExtraURI struct {
	GammaMerklURI      string
	GammaHypervisorURI []string
	PendleCoreURI      string
}

/**************************************************************************************************
** TChain is the primary configuration structure for a blockchain network supported by yDaemon.
** It contains all the necessary information to interact with a specific chain, including:
** - Network identification and connection details
** - Contract addresses for various Yearn and partner systems
** - Protocol-specific configuration
** - Lists of special tokens, vaults, and staking opportunities
**
** This comprehensive structure centralizes all chain-specific settings to ensure consistent
** behavior across the application.
**************************************************************************************************/
type TChain struct {
	ID                    uint64
	RpcURI                string
	SubgraphURI           string
	EtherscanURI          string
	MaxBlockRange         uint64
	MaxBatchSize          uint64
	AvgBlocksPerDay       int
	CanUseWebsocket       bool
	LensContract          TContractData
	MulticallContract     TContractData
	YBribeV3Contract      TContractData
	PartnerContract       TContractData
	Coin                  models.TERC20Token
	StakingRewardRegistry []TContractData
	Registries            []TContractData
	YearnXRegistries      []TContractData
	ExtraStakingContracts []TExtraStakingContracts
	ExtraVaults           []models.TVaultsFromRegistry
	BlacklistedVaults     []common.Address
	ExtraTokens           []common.Address
	IgnoredTokens         []common.Address
	Curve                 TChainCurve
	ExtraURI              TChainExtraURI
}

/**************************************************************************************************
** DEFAULT_COIN_ADDRESS represents the standard address used for the native coin of a blockchain.
** This is a special address (0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE) that represents ETH
** on Ethereum, and equivalent native currencies on other chains.
**************************************************************************************************/
var DEFAULT_COIN_ADDRESS = common.HexToAddress(`0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE`)

/**************************************************************************************************
** CHAINS is a global map that stores configuration for all supported blockchain networks.
** The map is indexed by chain ID (uint64) for easy lookup of specific chain configurations.
**************************************************************************************************/
var CHAINS = map[uint64]TChain{}

/**************************************************************************************************
** SUPPORTED_CHAIN_IDS contains a list of all chain IDs supported by yDaemon.
** This list is populated during initialization and can be used to iterate through all
** supported chains without having to extract keys from the CHAINS map.
**************************************************************************************************/
var SUPPORTED_CHAIN_IDS = []uint64{}

/**************************************************************************************************
** GetChains returns the complete map of all supported chain configurations.
** This is useful when you need to access or iterate through all chain configurations.
**
** @return map[uint64]TChain A map of chain IDs to their respective configurations
**************************************************************************************************/
func GetChains() map[uint64]TChain {
	return CHAINS
}

/**************************************************************************************************
** GetChain retrieves the configuration for a specific chain by its ID.
** This is the preferred method to access chain configuration as it handles the case
** where a chain may not be supported.
**
** @param chainID The unique identifier for the blockchain network
** @return TChain The chain configuration
** @return bool True if the chain is supported, false otherwise
**************************************************************************************************/
func GetChain(chainID uint64) (TChain, bool) {
	chain, ok := CHAINS[chainID]
	return chain, ok
}
