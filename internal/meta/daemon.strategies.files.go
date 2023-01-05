package meta

import (
	"encoding/json"
	"strconv"

	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/common/env"
	"github.com/yearn/ydaemon/common/helpers"
	"github.com/yearn/ydaemon/common/traces"
	"github.com/yearn/ydaemon/common/types/common"
)

// TStrategyFromMeta is the structure of data we receive when calling meta.yearn.finance/api/1/strategies/all
type TStrategyFromMeta struct {
	Address          common.Address   `json:"address"`
	Name             string           `json:"name"`
	Description      string           `json:"description"`
	RelatedAddresses []common.Address `json:"addresses"`
	Protocols        []string         `json:"protocols"`
	ChainID          uint64           `json:"chainID"`
	Localization     *TLocalization   `json:"localization,omitempty"`
}

/**********************************************************************************************
** Set of functions to store and retrieve the tokens from the cache and/or database and being
** able to access them from the rest of the application.
** The _vaultMap variable is not exported and is only used internally by the functions below.
**********************************************************************************************/
var _metaStrategyMap = make(map[uint64]map[ethcommon.Address]*TStrategyFromMeta)

func init() {
	for _, chainID := range env.SUPPORTED_CHAIN_IDS {
		if _, ok := _metaStrategyMap[chainID]; !ok {
			_metaStrategyMap[chainID] = make(map[ethcommon.Address]*TStrategyFromMeta)
		}
	}
}

/**********************************************************************************************
** setStrategyInMap will put a TStrategyFromMeta in the _metaStrategyMap variable.
**********************************************************************************************/
func setStrategyInMap(chainID uint64, strategy *TStrategyFromMeta) {
	if _, ok := _metaStrategyMap[chainID]; !ok {
		_metaStrategyMap[chainID] = make(map[ethcommon.Address]*TStrategyFromMeta)
	}
	_metaStrategyMap[chainID][strategy.Address.Address] = strategy
}

/**********************************************************************************************
** GetMetaStrategy will, for a given chainID, try to retrieve the strategy from the
** _metaStrategyMap variable.
** It will return the strategy if found, and a boolean indicating if the strategy was found or
** not.
**********************************************************************************************/
func GetMetaStrategy(chainID uint64, strategyAddress common.Address) (*TStrategyFromMeta, bool) {
	if strategysForChain, ok := _metaStrategyMap[chainID]; ok {
		if strategy, ok := strategysForChain[strategyAddress.ToAddress()]; ok {
			return strategy, true
		}
	}
	return nil, false
}

/**********************************************************************************************
** ListMetaStrategies will, for a given chainID, list all the strategies from the
** _metaStrategyMap variable.
**********************************************************************************************/
func ListMetaStrategies(chainID uint64) []*TStrategyFromMeta {
	var strategies []*TStrategyFromMeta
	for _, strategy := range _metaStrategyMap[chainID] {
		strategies = append(strategies, strategy)
	}
	return strategies
}

/**************************************************************************************************
** RetrieveAllStrategiesFromFiles will read all tokens in the /meta/strategies directory for a
** given chainID and store them in the _metaStrategyMap global variable.
** The goal of this function is to get a list of all meta information about the strategies for a
** given chainID.
**
** Arguments:
** - chainID: the chain ID of the network we are working on
**************************************************************************************************/
func RetrieveAllStrategiesFromFiles(chainID uint64) {
	chainIDStr := strconv.FormatUint(chainID, 10)
	content, _, err := helpers.ReadAllFilesInDir(env.BASE_DATA_PATH+`/meta/strategies/`+chainIDStr+`/`, `.json`)
	if err != nil {
		traces.
			Capture(`warn`, `impossible to read meta files for strategies on chain `+chainIDStr).
			SetEntity(`meta`).
			SetExtra(`error`, err.Error()).
			SetTag(`chainID`, strconv.FormatUint(chainID, 10)).
			Send()
		return
	}
	for _, content := range content {
		strategy := TStrategyFromMeta{}
		if err := json.Unmarshal(content, &strategy); err != nil {
			traces.
				Capture(`warn`, `impossible to unmarshall meta files for strategies response body `+chainIDStr).
				SetEntity(`meta`).
				SetExtra(`error`, err.Error()).
				SetTag(`chainID`, strconv.FormatUint(chainID, 10)).
				SetExtra(`content`, string(content)).
				Send()
			continue
		}
		relatedAddresses := append([]common.Address{}, strategy.RelatedAddresses...)
		strategy.RelatedAddresses = relatedAddresses
		strategy.ChainID = chainID
		for _, relatedAddress := range relatedAddresses {
			strategy.Address = relatedAddress
			setStrategyInMap(chainID, &strategy)
		}
	}
}
