package helpers

import (
	"path"
	"path/filepath"
	"runtime"

	"github.com/ethereum/go-ethereum/common"
)

func getCurrentPath() string {
	_, filename, _, _ := runtime.Caller(1)

	return path.Dir(filename)
}

// GITHUB_ICON_BASE_URL is the base URL to access the tokens icons from the yearn-assets repository on GitHub
var GITHUB_ICON_BASE_URL = `https://raw.githubusercontent.com/yearn/yearn-assets/master/icons/multichain-tokens/`

// BASE_DATA_PATH is the base path to access the data informations
var BASE_DATA_PATH, _ = filepath.Abs(getCurrentPath() + `../../../../data/`)

// API_V1_BASE_URL is the base URL to access query the legacy Yearn's api
var API_V1_BASE_URL = `https://api.yearn.finance/v1/chains/`

// RISK_BASE_URL is the base URL to access the risk framework
var RISK_BASE_URL = `https://d3971bp2359cnv.cloudfront.net/api/strategies/`

// SUPPORTED_CHAIN_IDS is the list of supported chain IDs
var SUPPORTED_CHAIN_IDS = []uint64{1, 10, 250, 42161}

// BLACKLISTED_VAULTS contains the vault we should not work with
var BLACKLISTED_VAULTS = map[uint64][]common.Address{
	1: {
		common.HexToAddress("0x662fBF2c1E4b04342EeBA6371ec1C7420042B86F"),
		common.HexToAddress("0x9C13e225AE007731caA49Fd17A41379ab1a489F4"),
	},
	10: {
		common.HexToAddress("0x0000000000000000000000000000000000000000"),
	},
	250: {
		common.HexToAddress("0x03B82e4070cA32FF63A03F2EcfC16c0165689a9d"),
	},
	42161: {
		common.HexToAddress("0x0000000000000000000000000000000000000000"),
	},
}

// CURVE_FACTORY_URI contains the the URI of the Curve Factory to use
var CURVE_FACTORY_URI = map[uint64][]string{
	1: {
		`https://api.curve.fi/api/getPools/ethereum/factory`,
		`https://api.curve.fi/api/getPools/ethereum/factory-crypto`,
	},
	10: {
		`https://api.curve.fi/api/getPools/optimism/factory`,
	},
	250: {
		`https://api.curve.fi/api/getPools/fantom/factory`,
	},
	42161: {},
}

// STRATEGIES_CONDITIONS contains the possible conditions to determine which strategies should
// be used in the return value.
// If the strategy is `absolute`, an active strategy will be isActive, inQueue and with a debt > 0
// If the strategy is `inQueue`, we will check if the vault has the strategy in it's withdrawal queue
// If the strategy is `debtLimit`, we will check if the vault has allocated a debtLimit to the strategy
var STRATEGIES_CONDITIONS = []string{`inQueue`, `debtLimit`, `absolute`}
