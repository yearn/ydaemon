package apy

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/common/bigNumber"
)

var ZERO = bigNumber.NewFloat(0)
var ONE = bigNumber.NewFloat(1)
var WEI = bigNumber.NewFloat(1e18)

var YEARN_VOTER_ADDRESS = map[uint64]common.Address{
	1:     common.HexToAddress(`0xF147b8125d2ef93FB6965Db97D6746952a133934`),
	10:    common.HexToAddress(`0xea3a15df68fcdbe44fdb0db675b2b3a14a148b26`),
	250:   common.HexToAddress(`0x72a34AbafAB09b15E7191822A679f28E067C4a16`),
	42161: common.HexToAddress(`0x6346282DB8323A54E840c6C772B4399C9c655C0d`),
}

var CONVEX_VOTER_ADDRESS = map[uint64]common.Address{
	1:     common.HexToAddress(`0x989AEb4d175e16225E39E87d0D97A3360524AD80`),
	10:    {},
	250:   {},
	42161: {},
}

var CVX_BOOSTER_ADDRESS = map[uint64]common.Address{
	1:     common.HexToAddress(`0xF403C135812408BFbE8713b5A23a04b3D48AAE31`),
	10:    {},
	250:   {},
	42161: {},
}

var CRV_TOKEN_ADDRESS = map[uint64]common.Address{
	1:     common.HexToAddress(`0xD533a949740bb3306d119CC777fa900bA034cd52`),
	10:    common.HexToAddress(`0x0994206dfE8De6Ec6920FF4D779B0d950605Fb53`),
	250:   common.HexToAddress(`0x1E4F97b9f9F913c46F1632781732927B9019C68b`),
	42161: common.HexToAddress(`0x11cDb42B0EB46D95f990BeDD4695A6e3fA034978`),
}

var CVX_TOKEN_ADDRESS = map[uint64]common.Address{
	1:     common.HexToAddress(`0x4e3FBD56CD56c3e72c1403e103b45Db9da5B9D2B`),
	10:    {},
	250:   {},
	42161: {},
}

// CURVE_GAUGES_URI contains the URI of the Curve gauges to use
var CURVE_GAUGES_URI = map[uint64]string{
	1:     `https://api.curve.fi/api/getAllGauges?blockchainId=ethereum`,
	10:    `https://api.curve.fi/api/getAllGauges?blockchainId=optimism`,
	250:   `https://api.curve.fi/api/getAllGauges?blockchainId=fantom`,
	42161: `https://api.curve.fi/api/getAllGauges?blockchainId=arbitrum`,
}

// CURVE_SUBGRAPHDATA_URI contains the URI of the Curve gauges to use
var CURVE_SUBGRAPHDATA_URI = map[uint64]string{
	1:     `https://api.curve.fi/api/getSubgraphData/ethereum`,
	10:    `https://api.curve.fi/api/getSubgraphData/optimism`,
	250:   `https://api.curve.fi/api/getSubgraphData/fantom`,
	42161: `https://api.curve.fi/api/getSubgraphData/arbitrum`,
}
