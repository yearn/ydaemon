package main

import (
	"bytes"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestEnvironment(t *testing.T) {
	//Init the server as non-blocking mode
	go NewRouter().Run()
	time.Sleep(3 * time.Second)

	//Init the daemons as blocking mode: we want to wait for them to complete before we continue
	SummonDaemons(1)
	SummonDaemons(10)
	SummonDaemons(250)
	SummonDaemons(42161)

	//Testing the reports endpoint
	resp, err := http.Get(`http://localhost:8080/250/reports/0xb99d6662127d9A3c68Bc949679f364E1739E2CA1`)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	resp, err = http.Get(`http://localhost:8080/hello/reports/0xb99d6662127d9A3c68Bc949679f364E1739E2CA1`)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)

	//Testing the info/chains route
	resp, err = http.Get(`http://localhost:8080/info/chains`)
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
	assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)

	req, err = http.NewRequest("POST", `http://localhost:8080/hello/graph`, bytes.NewBuffer(jsonInvalidStr))
	req.Header.Set("Content-Type", "application/json")
	assert.Nil(t, err)
	client = &http.Client{}
	resp, err = client.Do(req)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
}
