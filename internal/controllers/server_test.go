package controllers

import (
	"bytes"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yearn/ydaemon/internal/daemons"
	"github.com/yearn/ydaemon/internal/store"
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

	//Testing the a meta paths - Legacy
	resp, err = http.Get(`http://localhost:8080/api/1/strategies/all`)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	resp, err = http.Get(`http://localhost:8080/api/250/strategies/all?loc=all`)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	resp, err = http.Get(`http://localhost:8080/api/hello/strategies/all`)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
	resp, err = http.Get(`http://localhost:8080/api/4545456/strategies/all`)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)

	resp, err = http.Get(`http://localhost:8080/api/1/tokens/all`)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	resp, err = http.Get(`http://localhost:8080/api/250/tokens/all?loc=all`)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	resp, err = http.Get(`http://localhost:8080/api/hello/tokens/all`)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
	resp, err = http.Get(`http://localhost:8080/api/4545456/tokens/all`)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)

	resp, err = http.Get(`http://localhost:8080/api/1/vaults/all`)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	resp, err = http.Get(`http://localhost:8080/api/hello/vaults/all`)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
	resp, err = http.Get(`http://localhost:8080/api/4545456/vaults/all`)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)

	resp, err = http.Get(`http://localhost:8080/api/1/protocols/all`)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	resp, err = http.Get(`http://localhost:8080/api/250/protocols/all?loc=all`)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	resp, err = http.Get(`http://localhost:8080/api/hello/protocols/all`)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
	resp, err = http.Get(`http://localhost:8080/api/4545456/protocols/all`)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)

	//Testing the a meta paths - Regular
	resp, err = http.Get(`http://localhost:8080/1/meta/strategies`)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	resp, err = http.Get(`http://localhost:8080/250/meta/strategies?loc=all`)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	resp, err = http.Get(`http://localhost:8080/hello/meta/strategies`)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
	resp, err = http.Get(`http://localhost:8080/4545456/meta/strategies`)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)

	resp, err = http.Get(`http://localhost:8080/1/meta/tokens`)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	resp, err = http.Get(`http://localhost:8080/250/meta/tokens?loc=all`)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	resp, err = http.Get(`http://localhost:8080/hello/meta/tokens`)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
	resp, err = http.Get(`http://localhost:8080/4545456/meta/tokens`)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)

	resp, err = http.Get(`http://localhost:8080/1/meta/vaults`)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	resp, err = http.Get(`http://localhost:8080/hello/meta/vaults`)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
	resp, err = http.Get(`http://localhost:8080/4545456/meta/vaults`)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)

	resp, err = http.Get(`http://localhost:8080/1/meta/protocols`)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	resp, err = http.Get(`http://localhost:8080/250/meta/protocols?loc=all`)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	resp, err = http.Get(`http://localhost:8080/hello/meta/protocols`)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
	resp, err = http.Get(`http://localhost:8080/4545456/meta/protocols`)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)

	//Testing the a meta paths - Individual
	resp, err = http.Get(`http://localhost:8080/250/meta/strategies/0xb99d6662127d9A3c68Bc949679f364E1739E2CA1`)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	resp, err = http.Get(`http://localhost:8080/250/meta/strategies/0xb99d6662127d9A3c68Bc949679f364E1739E2CA1?loc=all`)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	resp, err = http.Get(`http://localhost:8080/hello/meta/strategies/0xb99d6662127d9A3c68Bc949679f364E1739E2CA1`)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
	resp, err = http.Get(`http://localhost:8080/4545456/meta/strategies/0xb99d6662127d9A3c68Bc949679f364E1739E2CA1`)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)

	resp, err = http.Get(`http://localhost:8080/1/meta/tokens/0x0000000000085d4780B73119b644AE5ecd22b376`)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	resp, err = http.Get(`http://localhost:8080/1/meta/tokens/0x0000000000085d4780B73119b644AE5ecd22b376?loc=all`)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	resp, err = http.Get(`http://localhost:8080/hello/meta/tokens/0x0000000000085d4780B73119b644AE5ecd22b376`)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
	resp, err = http.Get(`http://localhost:8080/4545456/meta/tokens/0x0000000000085d4780B73119b644AE5ecd22b376`)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)

	resp, err = http.Get(`http://localhost:8080/1/meta/vaults/0x04d73c87b20d372cB3240C72eEFB9d79bA5e4959`)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	resp, err = http.Get(`http://localhost:8080/1/meta/vaults/0x04d73c87b20d372cB3240C72eEFB9d79bA5e4959`)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	resp, err = http.Get(`http://localhost:8080/hello/meta/vaults/0x04d73c87b20d372cB3240C72eEFB9d79bA5e4959`)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
	resp, err = http.Get(`http://localhost:8080/4545456/meta/vaults/0x04d73c87b20d372cB3240C72eEFB9d79bA5e4959`)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)

	resp, err = http.Get(`http://localhost:8080/1/meta/protocols/1inch`)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	resp, err = http.Get(`http://localhost:8080/1/meta/protocols/1inch?loc=all`)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	resp, err = http.Get(`http://localhost:8080/hello/meta/protocols/1inch`)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
	resp, err = http.Get(`http://localhost:8080/4545456/meta/protocols/1inch`)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)

	// Testing all locs
	resp, err = http.Get(`http://localhost:8080/1/meta/protocols/1inch?loc=fr`)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	resp, err = http.Get(`http://localhost:8080/1/meta/protocols/1inch?loc=es`)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	resp, err = http.Get(`http://localhost:8080/1/meta/protocols/1inch?loc=de`)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	resp, err = http.Get(`http://localhost:8080/1/meta/protocols/1inch?loc=pt`)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	resp, err = http.Get(`http://localhost:8080/1/meta/protocols/1inch?loc=el`)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	resp, err = http.Get(`http://localhost:8080/1/meta/protocols/1inch?loc=tr`)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	resp, err = http.Get(`http://localhost:8080/1/meta/protocols/1inch?loc=vi`)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	resp, err = http.Get(`http://localhost:8080/1/meta/protocols/1inch?loc=zh`)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	resp, err = http.Get(`http://localhost:8080/1/meta/protocols/1inch?loc=hi`)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	resp, err = http.Get(`http://localhost:8080/1/meta/protocols/1inch?loc=ja`)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	resp, err = http.Get(`http://localhost:8080/1/meta/protocols/1inch?loc=id`)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	resp, err = http.Get(`http://localhost:8080/1/meta/protocols/1inch?loc=ru`)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	resp, err = http.Get(`http://localhost:8080/1/meta/protocols/1inch?loc=none`)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

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
