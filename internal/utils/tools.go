package utils

import (
	"crypto/sha256"
	"encoding/json"
	"io/ioutil"
	"math/big"
	"net/http"
	"time"

	"github.com/google/uuid"
)

//FormatUnits convert a bigInt to a decimal float64. Not 100% accurate, but good enough for now
func FormatUnits(wei *big.Int, decimals float64) *big.Float {
	f := new(big.Float)
	f.SetPrec(236) //  IEEE 754 octuple-precision binary floating-point format: binary256
	f.SetMode(big.ToNearestEven)
	fWei := new(big.Float)
	fWei.SetPrec(236) //  IEEE 754 octuple-precision binary floating-point format: binary256
	fWei.SetMode(big.ToNearestEven)
	return f.Quo(fWei.SetInt(wei), big.NewFloat(decimals))
}

//GetPublicIDString is an helper function used to generate an UUID based on a string
func GetPublicIDString(str string) string {
	hash := sha256.Sum256([]byte(str))
	trimmedHash := hash[:16]
	finalUUID, _ := uuid.FromBytes(trimmedHash)
	return finalUUID.String()
}

//PriceEthereum is the price for the ethereum token
var PriceEthereum = float32(0.0)

func fetchPrices() {
	resp, err := http.Get("https://api.coingecko.com/api/v3/simple/price?ids=ethereum&vs_currencies=usd")
	if err != nil {
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	t := make(map[string]map[string]float32)
	err = json.Unmarshal(body, &t)
	if err != nil {
		return
	}
	PriceEthereum = t[`ethereum`][`usd`]
}

//RefreshPrices refresh the ethereum price from coingecko
func RefreshPrices() {
	for {
		fetchPrices()
		time.Sleep(15 * time.Second)
	}
}
