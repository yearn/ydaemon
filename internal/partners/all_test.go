package partners

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
	helpers.SUPPORTED_CHAIN_IDS = append(helpers.SUPPORTED_CHAIN_IDS, 4545456)

	{
		c := Controller{}
		router.GET(`partners/count`, c.CountAllPartners)
		router.GET(`partners/all`, c.GetAllPartners)
		router.GET(`:chainID/partners/all`, c.GetPartners)
		router.GET(`:chainID/partners/:addressOrName`, c.GetPartner)
	}

	return router
}

func TestPartners(t *testing.T) {
	store.OpenDB()
	defer store.CloseDB()

	//Init the server as non-blocking mode
	go newRouter().Run()
	time.Sleep(3 * time.Second)

	//Init the daemons as blocking mode: we want to wait for them to complete before we continue
	wg := sync.WaitGroup{}
	wg.Add(5)
	LoadPartners(1, &wg)
	LoadPartners(10, &wg)
	LoadPartners(250, &wg)
	LoadPartners(42161, &wg)
	LoadPartners(42, &wg)
	wg.Wait()

	FetchPartnersFromFiles(1)
	FetchPartnersFromFiles(10)
	FetchPartnersFromFiles(250)
	FetchPartnersFromFiles(42161)
	FetchPartnersFromFiles(42)

	//Testing the a meta paths - Legacy
	resp, err := http.Get(`http://localhost:8080/partners/count`)
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
