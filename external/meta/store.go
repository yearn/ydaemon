package meta

import "github.com/yearn/ydaemon/common/types/common"

type TStore struct {
	// VaultsFromMeta holds the data for the vaults from the Yearn Meta API for each chain.
	VaultsFromMeta map[uint64]map[common.Address]TVaultFromMeta
	// TokensFromMeta holds the data for the tokens from the Yearn Meta API for each chain.
	TokensFromMeta map[uint64]map[common.Address]TTokenFromMeta
	// StrategiesFromMeta holds the data for the strategies from the Yearn Meta API for each chain.
	StrategiesFromMeta map[uint64]map[common.Address]TStrategyFromMeta
	// ProtocolsFromMeta holds the data for the protocols from the Yearn Meta API for each chain.
	ProtocolsFromMeta map[uint64]map[string]TProtocolsFromMeta
	// RawMetaStrategies contains the raw response from the meta endpoint for the strategies
	RawMetaStrategies map[uint64][]TStrategyFromMeta
	// RawMetaTokens contains the raw response from the meta endpoint for the tokens
	RawMetaTokens map[uint64][]TTokenFromMeta
	// RawMetaVaults contains the raw response from the meta endpoint for the vaults
	RawMetaVaults map[uint64][]TVaultFromMeta
	// RawMetaProtocols contains the raw response from the meta endpoint for the protocols
	RawMetaProtocols map[uint64][]TProtocolsFromMeta
}

// Store holds the meta data for each chain.
var Store = TStore{}

func init() {
	Store.VaultsFromMeta = make(map[uint64]map[common.Address]TVaultFromMeta)
	Store.TokensFromMeta = make(map[uint64]map[common.Address]TTokenFromMeta)
	Store.StrategiesFromMeta = make(map[uint64]map[common.Address]TStrategyFromMeta)
	Store.ProtocolsFromMeta = make(map[uint64]map[string]TProtocolsFromMeta)
	Store.RawMetaStrategies = make(map[uint64][]TStrategyFromMeta)
	Store.RawMetaTokens = make(map[uint64][]TTokenFromMeta)
	Store.RawMetaVaults = make(map[uint64][]TVaultFromMeta)
	Store.RawMetaProtocols = make(map[uint64][]TProtocolsFromMeta)
}
