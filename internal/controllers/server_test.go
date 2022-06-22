package controllers

import (
	"net/http"
	"testing"

	"github.com/majorfi/ydaemon/internal/daemons"
	"github.com/stretchr/testify/assert"
)

func TestEnvironment(t *testing.T) {
	//Init the server as non-blocking mode
	go NewRouter().Run()

	//Init the daemons as blocking mode: we want to wait for them to complete before we continue
	daemons.SummonDaemons(1, 0)
	daemons.SummonDaemons(250, 0)
	daemons.SummonDaemons(42161, 0)

	//Testing a valid request for ChainID == 1
	resp, err := http.Get(`http://localhost:8080/1/vaults/all?skip=1&limit=300&orderBy=id&order=asc`)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	//Testing a valid request for ChainID == 250
	resp, err = http.Get(`http://localhost:8080/250/vaults/all`)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	//Testing a valid request for ChainID == 42161
	resp, err = http.Get(`http://localhost:8080/42161/vaults/all`)
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

	//Testing the info/chains route
	resp, err = http.Get(`http://localhost:8080/info/chains`)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	//Testing the info/vaults/blacklisted route
	resp, err = http.Get(`http://localhost:8080/info/vaults/blacklisted`)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}
