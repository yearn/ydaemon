package main

import (
	"bytes"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yearn/ydaemon/internal/daemons"
	"github.com/yearn/ydaemon/internal/utils/store"
)

func TestEnvironment(t *testing.T) {
	store.OpenDB()
	defer store.CloseDB()

	//Init the server as non-blocking mode
	go NewRouter().Run()

	//Init the daemons as blocking mode: we want to wait for them to complete before we continue
	daemons.SummonDaemons(1, 0)
	daemons.SummonDaemons(10, 0)
	daemons.SummonDaemons(250, 0)
	daemons.SummonDaemons(42161, 0)

	//Testing a valid request for ChainID == 1
	resp, err := http.Get(`http://localhost:8080/1/vaults/all?skip=1&limit=300&orderBy=id&order=asc`)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	resp, err = http.Get(`http://localhost:8080/1/vaults/all?strategiesDetails=withDetails&strategiesRisk=withRisk`)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	//Testing a valid request for ChainID == 10
	resp, err = http.Get(`http://localhost:8080/10/vaults/all`)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	//Testing a valid request for ChainID == 250
	resp, err = http.Get(`http://localhost:8080/250/vaults/all?classification=all&strategiesCondition=invalid`)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	//Testing a valid request for ChainID == 42161
	resp, err = http.Get(`http://localhost:8080/42161/vaults/all?strategiesCondition=inQueue`)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	//Testing with an invalid chainID
	resp, err = http.Get(`http://localhost:8080/hello/vaults/all`)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)

	//Testing with an not supported chainID
	resp, err = http.Get(`http://localhost:8080/2/vaults/all`)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)

	//Testing with one specific vault
	resp, err = http.Get(`http://localhost:8080/1/vaults/0x6A5468752f8DB94134B6508dAbAC54D3b45efCE6`)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	resp, err = http.Get(`http://localhost:8080/hello/vaults/0x6A5468752f8DB94134B6508dAbAC54D3b45efCE6`)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
	resp, err = http.Get(`http://localhost:8080/1/vaults/0x662fBF2c1E4b04342EeBA6371ec1C7420042B86F`)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)

	//Testing the reports endpoint
	resp, err = http.Get(`http://localhost:8080/250/reports/0xb99d6662127d9A3c68Bc949679f364E1739E2CA1`)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	resp, err = http.Get(`http://localhost:8080/hello/reports/0xb99d6662127d9A3c68Bc949679f364E1739E2CA1`)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)

	//Testing the info/chains route
	resp, err = http.Get(`http://localhost:8080/info/chains`)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	//Testing the info/vaults/blacklisted route
	resp, err = http.Get(`http://localhost:8080/info/vaults/blacklisted`)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	resp, err = http.Get(`http://localhost:8080/info/vaults/blacklisted?chainID=10`)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	//Testing the a basic graphQL request
	jsonStr := []byte(`{
		tokens(first: 5) {
		  id
		  decimals
		  name
		  symbol
		}
	}`)
	req, err := http.NewRequest("POST", `http://localhost:8080/1/graph`, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	assert.Nil(t, err)
	client := &http.Client{}
	resp, err = client.Do(req)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	jsonInvalidStr := []byte(`{\n tokens(first: 5) {\n id \n decimals \n name \n symbol \n } \n }`)
	req, err = http.NewRequest("POST", `http://localhost:8080/1/graph`, bytes.NewBuffer(jsonInvalidStr))
	req.Header.Set("Content-Type", "application/json")
	assert.Nil(t, err)
	client = &http.Client{}
	resp, err = client.Do(req)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)

	req, err = http.NewRequest("POST", `http://localhost:8080/hello/graph`, bytes.NewBuffer(jsonInvalidStr))
	req.Header.Set("Content-Type", "application/json")
	assert.Nil(t, err)
	client = &http.Client{}
	resp, err = client.Do(req)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)

	//Testing vaults TVL
	resp, err = http.Get(`http://localhost:8080/1/vaults/tvl`)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	resp, err = http.Get(`http://localhost:8080/100/vaults/tvl`)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)

	//Testing partners
	resp, err = http.Get(`http://localhost:8080/partners/count`)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	resp, err = http.Get(`http://localhost:8080/partners/all`)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	resp, err = http.Get(`http://localhost:8080/1/partners/all`)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	resp, err = http.Get(`http://localhost:8080/250/partners/all`)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	resp, err = http.Get(`http://localhost:8080/1/partners/inverse`)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	resp, err = http.Get(`http://localhost:8080/1/partners/0x926dF14a23BE491164dCF93f4c468A50ef659D5B`)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	resp, err = http.Get(`http://localhost:8080/1/partners/0x926dF14a23BE491164dCF93f4c468A50ef659D5A`)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
	resp, err = http.Get(`http://localhost:8080/1/partners/major`)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
	resp, err = http.Get(`http://localhost:8080/250/partners/inverse`)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
}
