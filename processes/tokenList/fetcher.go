package tokenList

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/yearn/ydaemon/common/logs"
	"github.com/yearn/ydaemon/internal/models"
)

var WidoTokenList models.TTokenList
var PortalsTokenList models.TTokenList

// fetchDefaultTokenList is used to fetch a tokenList from theses https://tokenlists.org/
func fetchDefaultTokenList(url string) models.TTokenList {
	resp, err := http.Get(url)
	if err != nil {
		logs.Error(err)
		return models.TTokenList{}
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		logs.Error(err)
		return models.TTokenList{}
	}

	if (resp.StatusCode < 200) || (resp.StatusCode > 299) {
		logs.Error(`request failed for url: ` + url)
		return models.TTokenList{}
	}

	var list models.TTokenList
	if err := json.Unmarshal(body, &list); err != nil {
		logs.Error(err, url)
		return models.TTokenList{}
	}
	return list
}
