package meta

import (
	"encoding/json"
	"strconv"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/common/env"
	"github.com/yearn/ydaemon/common/helpers"
	"github.com/yearn/ydaemon/common/traces"
	"github.com/yearn/ydaemon/internal/models"
)

/**********************************************************************************************
** Set of functions to store and retrieve the tokens from the cache and/or database and being
** able to access them from the rest of the application.
** The _metaStrategyMap variable is not exported and is only used internally by the functions
** below.
**********************************************************************************************/
var _metaStrategyMap = make(map[uint64]*sync.Map)

func initOrGetMetaStrategyMap(chainID uint64) *sync.Map {
	syncMap := _metaStrategyMap[chainID]
	if syncMap == nil {
		syncMap = &sync.Map{}
		_metaStrategyMap[chainID] = syncMap
	}
	return syncMap
}

func init() {
	for _, chainID := range env.SUPPORTED_CHAIN_IDS {
		if _, ok := _metaStrategyMap[chainID]; !ok {
			_metaStrategyMap[chainID] = &sync.Map{}
		}
	}
}

/**********************************************************************************************
** setStrategyInMap will put a TStrategyFromMeta in the _metaStrategyMap variable.
**********************************************************************************************/
func setStrategyInMap(chainID uint64, strategy *models.TStrategyFromMeta) {
	syncMap := initOrGetMetaStrategyMap(chainID)
	syncMap.Store(strategy.Address, strategy)
}

/**********************************************************************************************
** GetMetaStrategy will, for a given chainID, try to retrieve the strategy from the
** _metaStrategyMap variable.
** It will return the strategy if found, and a boolean indicating if the strategy was found or
** not.
**********************************************************************************************/
func GetMetaStrategy(chainID uint64, strategyAddress common.Address) (*models.TStrategyFromMeta, bool) {
	syncMap := initOrGetMetaStrategyMap(chainID)
	data, ok := syncMap.Load(strategyAddress)
	if !ok {
		return nil, false
	}
	return data.(*models.TStrategyFromMeta), true
}

/**********************************************************************************************
** ListMetaStrategies will, for a given chainID, list all the strategies from the
** _metaStrategyMap variable.
**********************************************************************************************/
func ListMetaStrategies(chainID uint64) []*models.TStrategyFromMeta {
	syncMap := initOrGetMetaProtocolMap(chainID)
	var retValue []*models.TStrategyFromMeta
	syncMap.Range(func(_, data interface{}) bool {
		retValue = append(retValue, data.(*models.TStrategyFromMeta))
		return true
	})
	return retValue
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
		strategy := models.TStrategyFromMeta{}
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
