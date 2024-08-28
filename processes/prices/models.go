package prices

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
		USDPrice any             `json:"usdPrice"`
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
type TGeckoAPIKeyStatus struct {
	Status struct {
		ErrorCode    int    `json:"error_code"`
		ErrorMessage string `json:"error_message"`
	} `json:"status"`
}

type TVeloToken struct {
	Price              float64 `json:"price"`
	NativeChainAddress string  `json:"nativeChainAddress"`
	NativeChainID      int     `json:"nativeChainId"`
	Address            string  `json:"address"`
	Name               string  `json:"name"`
	Symbol             string  `json:"symbol"`
	Decimals           int     `json:"decimals"`
	LogoURI            string  `json:"logoURI"`
}
type TVeloPairData struct {
	TVL           float64    `json:"tvl"`
	APR           float64    `json:"apr"`
	Address       string     `json:"address"`
	Symbol        string     `json:"symbol"`
	Decimals      int        `json:"decimals"`
	Stable        bool       `json:"stable"`
	Reserve0      float64    `json:"reserve0"`
	Reserve1      float64    `json:"reserve1"`
	Token0Address string     `json:"token0_address"`
	Token1Address string     `json:"token1_address"`
	GaugeAddress  string     `json:"gauge_address"`
	IsStable      bool       `json:"isStable"`
	TotalSupply   float64    `json:"totalSupply"`
	Token0        TVeloToken `json:"token0"`
	Token1        TVeloToken `json:"token1"`
	Gauge         struct {
		Bribes []struct {
			Token TVeloToken `json:"token"`
		}
	} `json:"gauge"`
}

type TVeloPairs struct {
	Data []TVeloPairData `json:"data"`
}
