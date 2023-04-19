package meta

import (
	"encoding/json"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/common/env"
	"github.com/yearn/ydaemon/common/helpers"
	"github.com/yearn/ydaemon/common/traces"
	"github.com/yearn/ydaemon/internal/models"
)

/**********************************************************************************************
** Set of functions to store and retrieve the tokens from the cache and/or database and being
** able to access them from the rest of the application.
** The _vaultMap variable is not exported and is only used internally by the functions below.
**********************************************************************************************/
var _metaTokentMap = make(map[uint64]map[common.Address]*models.TTokenFromMeta)

/**********************************************************************************************
** setTokenInMap will put a TTokenFromMeta in the _metaTokentMap variable.
**********************************************************************************************/
func setTokenInMap(chainID uint64, token *models.TTokenFromMeta) {
	if _, ok := _metaTokentMap[chainID]; !ok {
		_metaTokentMap[chainID] = make(map[common.Address]*models.TTokenFromMeta)
	}
	_metaTokentMap[chainID][token.Address] = token
}

/**********************************************************************************************
** GetMetaToken will, for a given chainID, try to retrieve the token from the _metaTokentMap
** variable.
** It will return the token if found, and a boolean indicating if the token was found or not.
**********************************************************************************************/
func GetMetaToken(chainID uint64, tokenAddress common.Address) (*models.TTokenFromMeta, bool) {
	if tokensForChain, ok := _metaTokentMap[chainID]; ok {
		if token, ok := tokensForChain[tokenAddress]; ok {
			return token, true
		}
	}
	return nil, false
}

/**********************************************************************************************
** ListMetaTokens will, for a given chainID, list all the tokens from the _metaTokentMap
** variable.
**********************************************************************************************/
func ListMetaTokens(chainID uint64) []*models.TTokenFromMeta {
	var tokens []*models.TTokenFromMeta
	for _, token := range _metaTokentMap[chainID] {
		tokens = append(tokens, token)
	}
	return tokens
}

/**************************************************************************************************
** RetrieveAllTokensFromFiles will read all tokens in the /meta/tokens directory for a given chainID
** and store them in the _metaTokentMap global variable.
** The goal of this function is to get a list of all meta information about the tokens for a given
** chainID.
**
** Arguments:
** - chainID: the chain ID of the network we are working on
**************************************************************************************************/
func RetrieveAllTokensFromFiles(chainID uint64) {
	chainIDStr := strconv.FormatUint(chainID, 10)
	content, filenames, err := helpers.ReadAllFilesInDir(env.BASE_DATA_PATH+`/meta/tokens/`+chainIDStr+`/`, `.json`)
	if err != nil {
		traces.
			Capture(`warn`, `impossible to read meta files for tokens on chain `+chainIDStr).
			SetEntity(`meta`).
			SetExtra(`error`, err.Error()).
			SetTag(`chainID`, strconv.FormatUint(chainID, 10)).
			Send()
		return
	}

	for index, content := range content {
		token := models.TTokenFromMeta{}
		if err := json.Unmarshal(content, &token); err != nil {
			traces.
				Capture(`warn`, `impossible to unmarshall meta files for tokens response body `+chainIDStr).
				SetEntity(`meta`).
				SetExtra(`error`, err.Error()).
				SetTag(`chainID`, strconv.FormatUint(chainID, 10)).
				SetExtra(`content`, string(content)).
				Send()
			continue
		}
		token.Address = common.HexToAddress(strings.TrimSuffix(filenames[index], `.json`))
		token.ChainID = chainID
		setTokenInMap(chainID, &token)
	}
}
