package tokensList

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/yearn/ydaemon/common/logs"
)

var WidoTokenList DefaultTokenListData
var PortalsTokenList DefaultTokenListData

// fetchDefaultTokenList is used to fetch a tokenList from theses https://tokenlists.org/
func fetchDefaultTokenList(url string) DefaultTokenListData {
	resp, err := http.Get(url)
	if err != nil {
		logs.Error(err)
		return DefaultTokenListData{}
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logs.Error(err)
		return DefaultTokenListData{}
	}

	if (resp.StatusCode < 200) || (resp.StatusCode > 299) {
		logs.Error(`request failed for url: ` + url)
		return DefaultTokenListData{}
	}

	var list DefaultTokenListData
	if err := json.Unmarshal(body, &list); err != nil {
		logs.Error(err, url)
		return DefaultTokenListData{}
	}
	return list
}
