package models

import "encoding/json"

type TCurveFactoriesPoolData struct {
	Name        string  `json:"name"`
	Symbol      string  `json:"symbol"`
	Address     string  `json:"address"`
	LPAddress   string  `json:"lpTokenAddress"`
	TotalSupply string  `json:"totalSupply"`
	USDTotal    float64 `json:"usdTotal"`
	Coins       []struct {
		Address  string          `json:"address"`
		Decimals json.RawMessage `json:"decimals"`
		Symbol   string          `json:"symbol"`
		USDPrice float64         `json:"usdPrice"`
	} `json:"coins"`
}

type TCurveFactories struct {
	Data struct {
		PoolData []TCurveFactoriesPoolData `json:"poolData"`
	} `json:"data"`
}

type TLlamaPriceData struct {
	Decimals int     `json:"decimals"`
	Price    float64 `json:"price"`
	Symbol   string  `json:"symbol"`
}

type TLlamaPrice struct {
	Coins map[string]TLlamaPriceData `json:"coins"`
}

type TGeckoPrice map[string]struct {
	USDPrice float64 `json:"usd"`
}
