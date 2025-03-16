package helpers

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/yearn/ydaemon/common/logs"
)

/**************************************************************************************************
** The fetch package provides HTTP utility functions for fetching and processing JSON data from
** external APIs. These functions simplify common tasks like making HTTP requests, handling
** responses, and parsing JSON data, while also providing special handling for certain API
** endpoints that require custom headers.
**
** This package supports both error-swallowing and error-returning variants of the fetch
** functions to accommodate different error handling requirements across the application.
**************************************************************************************************/

/**************************************************************************************************
** FetchJSON makes an HTTP GET request to the specified URI and unmarshals the JSON response
** into the provided generic type. This function handles all aspects of the HTTP request,
** including setting appropriate headers for certain APIs that require them.
**
** This function has "swallowing" error behavior - it logs errors but does not return them,
** instead returning a zero value of the requested type if any error occurs.
**
** @param uri The URL to fetch JSON data from
** @return data The unmarshaled JSON data as the specified generic type T
**************************************************************************************************/
func FetchJSON[T any](uri string) (data T) {
	var resp *http.Response
	var err error

	if strings.Contains(uri, `api.portals.fi`) ||
		strings.Contains(uri, `api.1inch.io`) ||
		strings.Contains(uri, `api.joinwido.com`) {
		req, _ := http.NewRequest("GET", uri, nil)
		req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36")
		resp, err = http.DefaultClient.Do(req)
	} else {
		resp, err = http.Get(uri)
	}
	if err != nil {
		logs.Error(err)
		return data
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logs.Error(err)
		return data
	}

	if (resp.StatusCode < 200) || (resp.StatusCode > 299) {
		logs.Error(`failed to fetch: ` + uri)
		return data
	}

	if err := json.Unmarshal(body, &data); err != nil {
		logs.Error(err)
		return data
	}
	return data
}

/**************************************************************************************************
** FetchJSONWithReject makes an HTTP GET request to the specified URI and unmarshals the JSON
** response into the provided generic type, but unlike FetchJSON, it returns any errors that
** occur during the process.
**
** This function is useful when the caller needs to know if and why a request failed, allowing
** for more sophisticated error handling and potential retry logic.
**
** @param uri The URL to fetch JSON data from
** @return data The unmarshaled JSON data as the specified generic type T
** @return err An error if any step of the fetch process failed, or nil on success
**************************************************************************************************/
func FetchJSONWithReject[T any](uri string) (data T, err error) {
	var resp *http.Response

	if strings.Contains(uri, `api.portals.fi`) ||
		strings.Contains(uri, `api.1inch.io`) ||
		strings.Contains(uri, `api.joinwido.com`) {
		req, _ := http.NewRequest("GET", uri, nil)
		req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36")
		resp, err = http.DefaultClient.Do(req)
	} else {
		resp, err = http.Get(uri)
	}
	if err != nil {
		logs.Error(err)
		return data, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logs.Error(err)
		return data, err
	}

	if (resp.StatusCode < 200) || (resp.StatusCode > 299) {
		logs.Error(`failed to fetch: ` + uri)
		return data, errors.New(`failed to fetch: ` + uri)
	}

	if err := json.Unmarshal(body, &data); err != nil {
		logs.Error(err)
		return data, err
	}
	return data, nil
}
