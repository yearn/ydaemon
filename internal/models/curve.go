package models

type CurveGaugeData struct {
	InflationRate string `json:"inflation_rate"`
	WorkingSupply string `json:"working_supply"`
}

type CurveGaugeController struct {
	GaugeRelativeWeight string `json:"gauge_relative_weight"`
	GetGaugeWeight      string `json:"get_gauge_weight"`
	InflationRate       string `json:"inflation_rate"`
}

type CurveGaugeSwapData struct {
	VirtualPrice any `json:"virtual_price"`
}
type CurveGauge struct {
	Swap             string               `json:"swap"`
	SwapToken        string               `json:"swap_token"`
	Name             string               `json:"name"`
	ShortName        string               `json:"shortName"`
	Gauge            string               `json:"gauge"`
	GaugeData        CurveGaugeData       `json:"gauge_data"`
	GaugeController  CurveGaugeController `json:"gauge_controller"`
	SwapData         CurveGaugeSwapData   `json:"swap_data,omitempty"`
	Factory          bool                 `json:"factory"`
	SideChain        bool                 `json:"side_chain"`
	IsKilled         bool                 `json:"is_killed"`
	HasNoCrv         bool                 `json:"hasNoCrv"`
	Type             string               `json:"type"`
	LpTokenPrice     float64              `json:"lpTokenPrice"`
	IsMetaPool       bool                 `json:"is_meta_pool"`
	MetaPoolAddress  string               `json:"meta_pool_address"`
	MetaPoolName     string               `json:"meta_pool_name"`
	MetaPoolCategory string               `json:"meta_pool_category"`
}

type TCurveGauges struct {
	Data map[string]CurveGauge `json:"data"`
}

type TCurvePools struct {
	Data struct {
		PoolData []CurvePool `json:"poolData"`
	} `json:"data"`
}

type CurvePool struct {
	ID                        string              `json:"id"`
	Address                   string              `json:"address"`
	CoinsAddresses            []string            `json:"coinsAddresses"`
	Decimals                  []string            `json:"decimals"`
	VirtualPrice              any                 `json:"virtual_price"`
	AmplificationCoefficient  string              `json:"amplificationCoefficient"`
	UnderlyingDecimals        []string            `json:"underlyingDecimals"`
	AssetType                 string              `json:"assetType"`
	TotalSupply               string              `json:"totalSupply"`
	Name                      string              `json:"name"`
	LPTokenAddress            string              `json:"lpTokenAddress"`
	Symbol                    string              `json:"symbol"`
	PriceOracle               any                 `json:"priceOracle"`
	Implementation            string              `json:"implementation"`
	AssetTypeName             string              `json:"assetTypeName"`
	Coins                     []CurveGetCoin      `json:"coins"`
	USDTotal                  float64             `json:"usdTotal"`
	IsMetaPool                bool                `json:"isMetaPool"`
	UnderlyingCoins           []CurveGetCoin      `json:"underlyingCoins"`
	USDTotalExcludingBasePool float64             `json:"usdTotalExcludingBasePool"`
	GaugeAddress              string              `json:"gaugeAddress"`
	GaugeRewards              []CurveGetGaugeData `json:"gaugeRewards"`
	GaugeCrvApy               []float64           `json:"gaugeCrvApy"`
}

type CurveGetCoin struct {
	Address           string  `json:"address"`
	USDPrice          float64 `json:"usdPrice"`
	Decimals          any     `json:"decimals"`
	IsBasePoolLpToken bool    `json:"isBasePoolLpToken"`
	Symbol            string  `json:"symbol"`
	PoolBalance       string  `json:"poolBalance"`
}

type CurveGetGaugeData struct {
	GaugeAddress string  `json:"gaugeAddress"`
	TokenPrice   float64 `json:"tokenPrice"`
	Name         string  `json:"name"`
	Symbol       string  `json:"symbol"`
	Decimals     any     `json:"decimals"`
	APY          float64 `json:"apy"`
	TokenAddress string  `json:"tokenAddress"`
}

type TCurveSubgraphData struct {
	Data struct {
		PoolList []CurveSubgraphData `json:"poolList"`
	} `json:"data"`
}

type CurveSubgraphData struct {
	Type            string  `json:"type"`
	Address         string  `json:"address"`
	RawVolume       float64 `json:"rawVolume"`
	VolumeUSD       float64 `json:"volumeUSD"`
	LatestDailyApy  float64 `json:"latestDailyApy"`
	LatestWeeklyApy float64 `json:"latestWeeklyApy"`
}
