package meta

import (
	"net/http"
	"sync"
	"testing"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/yearn/ydaemon/internal/utils/helpers"
	"github.com/yearn/ydaemon/internal/utils/store"
)

// NewRouter create the routes and setup the server
func NewRouter() *gin.Engine {
	gin.EnableJsonDecoderDisallowUnknownFields()
	router := gin.New()
	router.Use(gin.Recovery())
	corsConf := cors.Config{
		AllowAllOrigins: true,
		AllowMethods:    []string{"GET", "HEAD"},
		AllowHeaders:    []string{`Origin`, `Content-Length`, `Content-Type`, `Authorization`},
	}
	router.Use(cors.New(corsConf))
	helpers.SUPPORTED_CHAIN_IDS = append(helpers.SUPPORTED_CHAIN_IDS, 4545456)

	{
		meta := Controller{}

		// Proxy meta strategies
		router.GET(`api/:chainID/strategies/all`, meta.GetMetaStrategiesLegacy)
		router.GET(`:chainID/meta/strategies`, meta.GetMetaStrategies)
		router.GET(`api/:chainID/strategies/:address`, meta.GetMetaStrategy)
		router.GET(`:chainID/meta/strategies/:address`, meta.GetMetaStrategy)
		router.GET(`:chainID/meta/strategy/:address`, meta.GetMetaStrategy)

		// Proxy meta tokens
		router.GET(`api/:chainID/tokens/all`, meta.GetMetaTokensLegacy)
		router.GET(`:chainID/meta/tokens`, meta.GetMetaTokens)
		router.GET(`api/:chainID/tokens/:address`, meta.GetMetaToken)
		router.GET(`:chainID/meta/tokens/:address`, meta.GetMetaToken)
		router.GET(`:chainID/meta/token/:address`, meta.GetMetaToken)

		// Proxy meta vaults
		router.GET(`api/:chainID/vaults/all`, meta.GetMetaVaultsLegacy)
		router.GET(`:chainID/meta/vaults`, meta.GetMetaVaults)
		router.GET(`api/:chainID/vaults/:address`, meta.GetMetaVault)
		router.GET(`:chainID/meta/vaults/:address`, meta.GetMetaVault)
		router.GET(`:chainID/meta/vault/:address`, meta.GetMetaVault)

		// Proxy meta protocols
		router.GET(`api/:chainID/protocols/all`, meta.GetMetaProtocolsLegacy)
		router.GET(`:chainID/meta/protocols`, meta.GetMetaProtocols)
		router.GET(`api/:chainID/protocols/:name`, meta.GetMetaProtocol)
		router.GET(`:chainID/meta/protocols/:name`, meta.GetMetaProtocol)
		router.GET(`:chainID/meta/protocol/:name`, meta.GetMetaProtocol)
	}

	return router
}

func TestTokens(t *testing.T) {
	store.OpenDB()
	defer store.CloseDB()

	//Init the server as non-blocking mode
	go NewRouter().Run()
	time.Sleep(3 * time.Second)

	//Init the daemons as blocking mode: we want to wait for them to complete before we continue
	wg := sync.WaitGroup{}
	wg.Add(5)
	LoadMetaTokens(1, &wg)
	LoadMetaTokens(10, &wg)
	LoadMetaTokens(250, &wg)
	LoadMetaTokens(42161, &wg)
	LoadMetaTokens(42, &wg)
	wg.Wait()

	FetchTokensFromMeta(1)
	FetchTokensFromMeta(10)
	FetchTokensFromMeta(250)
	FetchTokensFromMeta(42161)
	FetchTokensFromMeta(42)

	//Testing the a meta paths - Legacy
	resp, err := http.Get(`http://localhost:8080/api/1/tokens/all`)
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
	resp, err = http.Get(`http://localhost:8080/1/meta/tokens/hello`)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
}

func TestStrategies(t *testing.T) {
	store.OpenDB()
	defer store.CloseDB()

	//Init the server as non-blocking mode
	go NewRouter().Run()
	time.Sleep(3 * time.Second)

	//Init the daemons as blocking mode: we want to wait for them to complete before we continue
	wg := sync.WaitGroup{}
	wg.Add(5)
	LoadMetaStrategies(1, &wg)
	LoadMetaStrategies(10, &wg)
	LoadMetaStrategies(250, &wg)
	LoadMetaStrategies(42161, &wg)
	LoadMetaStrategies(42, &wg)
	wg.Wait()

	FetchStrategiesFromMeta(1)
	FetchStrategiesFromMeta(10)
	FetchStrategiesFromMeta(250)
	FetchStrategiesFromMeta(42161)
	FetchStrategiesFromMeta(42)

	//Testing the a meta paths - Legacy
	resp, err := http.Get(`http://localhost:8080/api/1/strategies/all`)
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
	resp, err = http.Get(`http://localhost:8080/1/meta/strategies/hello`)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
}

func TestVaults(t *testing.T) {
	store.OpenDB()
	defer store.CloseDB()

	//Init the server as non-blocking mode
	go NewRouter().Run()
	time.Sleep(3 * time.Second)

	//Init the daemons as blocking mode: we want to wait for them to complete before we continue
	wg := sync.WaitGroup{}
	wg.Add(5)
	LoadMetaVaults(1, &wg)
	LoadMetaVaults(10, &wg)
	LoadMetaVaults(250, &wg)
	LoadMetaVaults(42161, &wg)
	LoadMetaVaults(42, &wg)
	wg.Wait()

	FetchVaultsFromMeta(1)
	FetchVaultsFromMeta(10)
	FetchVaultsFromMeta(250)
	FetchVaultsFromMeta(42161)
	FetchVaultsFromMeta(42)

	resp, err := http.Get(`http://localhost:8080/api/1/vaults/all`)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	resp, err = http.Get(`http://localhost:8080/api/hello/vaults/all`)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
	resp, err = http.Get(`http://localhost:8080/api/4545456/vaults/all`)
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
	resp, err = http.Get(`http://localhost:8080/1/meta/vaults/hello`)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
}

func TestEnvironment(t *testing.T) {
	store.OpenDB()
	defer store.CloseDB()

	//Init the server as non-blocking mode
	go NewRouter().Run()
	time.Sleep(3 * time.Second)

	//Init the daemons as blocking mode: we want to wait for them to complete before we continue
	wg := sync.WaitGroup{}
	wg.Add(5)
	LoadMetaProtocols(1, &wg)
	LoadMetaProtocols(10, &wg)
	LoadMetaProtocols(250, &wg)
	LoadMetaProtocols(42161, &wg)
	LoadMetaProtocols(42, &wg)
	wg.Wait()

	FetchProtocolsFromMeta(1)
	FetchProtocolsFromMeta(10)
	FetchProtocolsFromMeta(250)
	FetchProtocolsFromMeta(42161)
	FetchProtocolsFromMeta(42)

	resp, err := http.Get(`http://localhost:8080/api/1/protocols/all`)
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
}
