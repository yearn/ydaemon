package vaults

import (
	"net/http"
	"sync"
	"testing"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/yearn/ydaemon/common/env"
	"github.com/yearn/ydaemon/common/store"
	"github.com/yearn/ydaemon/external/meta"
	"github.com/yearn/ydaemon/external/strategies"
	"github.com/yearn/ydaemon/external/tokens"
)

func newRouter() *gin.Engine {
	gin.EnableJsonDecoderDisallowUnknownFields()
	router := gin.New()
	router.Use(gin.Recovery())
	corsConf := cors.Config{
		AllowAllOrigins: true,
		AllowMethods:    []string{"GET", "HEAD"},
		AllowHeaders:    []string{`Origin`, `Content-Length`, `Content-Type`, `Authorization`},
	}
	router.Use(cors.New(corsConf))
	env.SUPPORTED_CHAIN_IDS = append(env.SUPPORTED_CHAIN_IDS, 4545456)

	{
		c := Controller{}
		// Retrieve the vaults for all chains
		router.GET(`vaults/tvl`, c.GetAllVaultsTVL)

		// Retrieve the vaults for a specific chainID
		router.GET(`:chainID/vaults/tvl`, c.GetVaultsTVL)
		router.GET(`:chainID/vaults/all`, c.GetAllVaults)
		router.GET(`:chainID/vaults/:address`, c.GetVault)
		router.GET(`:chainID/vault/:address`, c.GetVault)

		router.GET(`info/vaults/blacklisted`, c.GetBlacklistedVaults)
	}

	return router
}
func loadMetaDaemons() {
	meta.FetchTokensFromMeta(1)
	meta.FetchTokensFromMeta(10)
	meta.FetchTokensFromMeta(250)
	meta.FetchTokensFromMeta(42161)
	meta.FetchTokensFromMeta(42)

	meta.FetchStrategiesFromMeta(1)
	meta.FetchStrategiesFromMeta(10)
	meta.FetchStrategiesFromMeta(250)
	meta.FetchStrategiesFromMeta(42161)
	meta.FetchStrategiesFromMeta(42)

	meta.FetchVaultsFromMeta(1)
	meta.FetchVaultsFromMeta(10)
	meta.FetchVaultsFromMeta(250)
	meta.FetchVaultsFromMeta(42161)
	meta.FetchVaultsFromMeta(42)

	meta.FetchProtocolsFromMeta(1)
	meta.FetchProtocolsFromMeta(10)
	meta.FetchProtocolsFromMeta(250)
	meta.FetchProtocolsFromMeta(42161)
	meta.FetchProtocolsFromMeta(42)
}
func loadTokensDaemons() {
	tokens.FetchTokenList(1)
	tokens.FetchTokenList(10)
	tokens.FetchTokenList(250)
	tokens.FetchTokenList(42161)
	tokens.FetchTokenList(42)
}
func loadStrategiesDaemons() {
	strategies.FetchWithdrawalQueueMulticallData(1)
	strategies.FetchWithdrawalQueueMulticallData(10)
	strategies.FetchWithdrawalQueueMulticallData(250)
	strategies.FetchWithdrawalQueueMulticallData(42161)
	strategies.FetchWithdrawalQueueMulticallData(42)

	strategies.FetchStrategiesMulticallData(1)
	strategies.FetchStrategiesMulticallData(10)
	strategies.FetchStrategiesMulticallData(250)
	strategies.FetchStrategiesMulticallData(42161)
	strategies.FetchStrategiesMulticallData(42)
}
func loadVaultDaemons() {
	FetchVaultMulticallData(1)
	FetchVaultMulticallData(10)
	FetchVaultMulticallData(250)
	FetchVaultMulticallData(42161)
	FetchVaultMulticallData(42)
	FetchVaultMulticallData(4545456)
	FetchVaultsFromV1(1)
	FetchVaultsFromV1(10)
	FetchVaultsFromV1(250)
	FetchVaultsFromV1(42161)
	FetchVaultsFromV1(42)
	FetchVaultsFromV1(4545456)

	wg := sync.WaitGroup{}
	wg.Add(12)
	go LoadVaultMulticallData(1, &wg)
	go LoadVaultMulticallData(10, &wg)
	go LoadVaultMulticallData(250, &wg)
	go LoadVaultMulticallData(42161, &wg)
	go LoadVaultMulticallData(42, &wg)
	go LoadVaultMulticallData(4545456, &wg)
	go LoadAPIV1Vaults(1, &wg)
	go LoadAPIV1Vaults(10, &wg)
	go LoadAPIV1Vaults(250, &wg)
	go LoadAPIV1Vaults(42161, &wg)
	go LoadAPIV1Vaults(42, &wg)
	go LoadAPIV1Vaults(4545456, &wg)
	wg.Wait()

	FetchVaultMulticallData(1)
	FetchVaultMulticallData(10)
	FetchVaultMulticallData(250)
	FetchVaultMulticallData(42161)
	FetchVaultMulticallData(42)
	FetchVaultMulticallData(4545456)
	FetchVaultsFromV1(1)
	FetchVaultsFromV1(10)
	FetchVaultsFromV1(250)
	FetchVaultsFromV1(42161)
	FetchVaultsFromV1(42)
	FetchVaultsFromV1(4545456)
}

func TestEnvironment(t *testing.T) {
	store.OpenDB()
	defer store.CloseDB()
	loadMetaDaemons()
	loadTokensDaemons()
	loadStrategiesDaemons()

	//Init the server as non-blocking mode
	go newRouter().Run(":8082")
	time.Sleep(3 * time.Second)

	//Init the daemons as blocking mode: we want to wait for them to complete before we continue
	loadVaultDaemons()

	//Testing a valid request for ChainID == 1
	resp, err := http.Get(`http://localhost:8082/1/vaults/all?skip=1&limit=300&orderBy=id&order=asc`)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	resp, err = http.Get(`http://localhost:8082/1/vaults/all?strategiesDetails=withDetails&strategiesRisk=withRisk`)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	//Testing a valid request for ChainID == 10
	resp, err = http.Get(`http://localhost:8082/10/vaults/all`)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	//Testing a valid request for ChainID == 250
	resp, err = http.Get(`http://localhost:8082/250/vaults/all?classification=all&strategiesCondition=invalid`)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	//Testing a valid request for ChainID == 42161
	resp, err = http.Get(`http://localhost:8082/42161/vaults/all?strategiesCondition=inQueue`)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	//Testing with an invalid chainID
	resp, err = http.Get(`http://localhost:8082/hello/vaults/all`)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)

	//Testing with an not supported chainID
	resp, err = http.Get(`http://localhost:8082/2/vaults/all`)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)

	//Testing with one specific vault
	resp, err = http.Get(`http://localhost:8082/1/vaults/0x6A5468752f8DB94134B6508dAbAC54D3b45efCE6`)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	resp, err = http.Get(`http://localhost:8082/hello/vaults/0x6A5468752f8DB94134B6508dAbAC54D3b45efCE6`)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
	resp, err = http.Get(`http://localhost:8082/1/vaults/0x662fBF2c1E4b04342EeBA6371ec1C7420042B86F`)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)

	//Testing the info/vaults/blacklisted route
	resp, err = http.Get(`http://localhost:8082/info/vaults/blacklisted`)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	resp, err = http.Get(`http://localhost:8082/info/vaults/blacklisted?chainID=10`)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	//Testing vaults TVL
	resp, err = http.Get(`http://localhost:8082/vaults/tvl`)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	resp, err = http.Get(`http://localhost:8082/1/vaults/tvl`)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	resp, err = http.Get(`http://localhost:8082/100/vaults/tvl`)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
}
