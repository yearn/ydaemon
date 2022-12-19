package vaults

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/yearn/ydaemon/common/helpers"
)

type TVisionAPYResp struct {
	Results struct {
		A struct {
			Frames []struct {
				Data struct {
					Values [][]float64
				}
			}
		}
	}
}

func fetchFromVision(chainID uint64, vaultAddress string) float64 {
	client := &http.Client{}
	vault := vaultAddress
	expr := `"(((delta(yearn_vault{network=\"ETH\", param=\"pricePerShare\", address=\"` + vault + `\"}[7d]) + 1) ^ (1/7)) ^ 365 - 1) < 15000"`
	defaultStuff := `[{"datasource":{"uid":"PBFE396EC0B189D67","type":"prometheus"},"queryType":"timeSeriesQuery","maxDataPoints":1000,"expr":` + expr + `}]`

	now := strconv.FormatInt(time.Now().UnixMilli(), 10)
	minutesAgo := strconv.FormatInt(time.Now().Add(15*time.Minute).UnixMilli(), 10)
	data := strings.NewReader(`{"queries":` + defaultStuff + `,"from":"` + minutesAgo + `","to":"` + now + `"}`)
	req, err := http.NewRequest("POST", "https://yearn.vision/api/ds/query", data)
	if err != nil {
		return 0
	}
	req.Header.Set("authority", "yearn.vision")
	req.Header.Set("accept", "application/json, text/plain, */*")
	req.Header.Set("content-type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return 0
	}
	defer resp.Body.Close()
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0
	}

	var respData TVisionAPYResp
	err = json.Unmarshal(bodyText, &respData)
	if err != nil {
		return 0
	}
	frames := respData.Results.A.Frames
	if len(frames) == 0 {
		return 0
	}
	values := frames[0].Data.Values
	if len(values) != 2 {
		return 0
	}
	lastValue := values[1][len(values[1])-1]
	return lastValue
}

//GetVaultsVisionAPY will return the current APY for a given vault
func (y Controller) GetVaultsVisionAPY(c *gin.Context) {
	chainID, ok := helpers.AssertChainID(c.Param("chainID"))
	if !ok {
		c.String(http.StatusBadRequest, "invalid chainID")
		return
	}
	address, ok := helpers.AssertAddress(c.Param("address"), chainID)
	if !ok {
		c.String(http.StatusBadRequest, "invalid address")
		return
	}

	c.JSON(http.StatusOK, fetchFromVision(chainID, address.String()))
}
