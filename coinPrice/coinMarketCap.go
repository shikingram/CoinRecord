package coinPrice

type CoinMarketCapLastListResult struct {
	Data   []SingleCoinInfo `json:"data"`
	Status State            `json:"status"`
}

type State struct {
	Timestamp    string `json:"timestamp"`
	ErrorCode    int    `json:"error_code"`
	ErrorMessage string `json:"error_message"`
	Elapsed      int    `json:"elapsed"`
	CreditCount  int    `json:"credit_count"`
}

type SingleCoinInfo struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Symbol string `json:"symbol"`
	Slug   string `json:"slug"`
	//市值排名
	CmcRank           int     `json:"cmc_rank"`
	NumMarketPairs    int     `json:"num_market_pairs"`
	CirculatingSupply float64 `json:"circulating_supply"`
	//目前存在的硬币总量
	TotalSupply float64 `json:"total_supply"`

	MaxSupply   float64    `json:"max_supply"`
	LastUpdated string     `json:"last_updated"`
	DateAdded   string     `json:"date_added"`
	Tags        []string   `json:"tags"`
	Platform    quoteInfo  `json:"-"`
	Quote       quoteInfos `json:"quote"`
}

type quoteInfos struct {
	BTC quoteInfo
	USD quoteInfo
}

type quoteInfo struct {
	//价格
	Price float64 `json:"price"`
	// 24小时交易量
	Volume24H float64 `json:"volume_24h"`
	//1小时变化
	PercentChange1H float64 `json:"percent_change_1h"`
	//24小时变化
	PercentChange24H float64 `json:"percent_change_24h"`
	//7天变化
	PercentChange7D float64 `json:"percent_change_7d"`
	//市值
	MarketCap             float64 `json:"market_cap"`
	MarketCapDominance    float64 `json:"market_cap_dominance"`
	FullyDilutedMarketCap float64 `json:"fully_diluted_market_cap"`
	LastUpdated           string  `json:"last_updated"`
}

//===========
type GlobalMetricsResult struct {
	Data        GlobalMetricsData `json:"data"`
	Status      State            `json:"status"`
	LastUpdated string           `json:"last_updated"`
}

type GlobalMetricsData struct {
	ActiveCryptocurrencies          int                    `json:"active_cryptocurrencies"`
	TotalCryptocurrencies           int                    `json:"total_cryptocurrencies"`
	ActiveMarketPairs               int                    `json:"active_market_pairs"`
	ActiveExchanges                 int                    `json:"active_exchanges"`
	TotalExchanges                  int                    `json:"total_exchanges"`
	EthDominance                    float64                `json:"eth_dominance"`
	BtcDominance                    float64                `json:"btc_dominance"`
	EthDominanceYesterday           float64                `json:"eth_dominance_yesterday"`
	BtcDominanceYesterday           float64                `json:"btc_dominance_yesterday"`
	EthDominance24HPercentageChange float64                `json:"eth_dominance_24h_percentage_change"`
	BtcDominance24HPercentageChange float64                `json:"btc_dominance_24h_percentage_change"`
	DefiVolume24H                   float64                `json:"defi_volume_24h"`
	DefiVolume24HReported           float64                `json:"defi_volume_24h_reported"`
	DefiMarketCap                   float64                `json:"defi_market_cap"`
	Defi24HPercentageChange         float64                `json:"defi_24h_percentage_change"`
	StablecoinVolume24H             float64                `json:"stablecoin_volume_24h"`
	StablecoinVolume24HReported     float64                `json:"stablecoin_volume_24h_reported"`
	StablecoinMarketCap             float64                `json:"stablecoin_market_cap"`
	Stablecoin24HPercentageChange   float64                `json:"stablecoin_24h_percentage_change"`
	DerivativesVolume24H            float64                `json:"derivatives_volume_24h"`
	DerivativesVolume24HReported    float64                `json:"derivatives_volume_24h_reported"`
	Derivatives24HPercentageChange  float64                `json:"derivatives_24h_percentage_change"`
	Quote                           GlobalMetricsQuoteInfo `json:"quote"`
}

type GlobalMetricsQuoteInfo struct {
	USD GlobalMetricsQuoteInfoUsd
}

type GlobalMetricsQuoteInfoUsd struct {
	TotalMarketCap                          float64 `json:"total_market_cap"`
	TotalVolume24H                          float64 `json:"total_volume_24h"`
	TotalVolume24HReported                  float64 `json:"total_volume_24h_reported"`
	AltcoinVolume24H                        float64 `json:"altcoin_volume_24h"`
	AltcoinVolume24HReported                float64 `json:"altcoin_volume_24h_reported"`
	AltcoinMarketCap                        float64 `json:"altcoin_market_cap"`
	DefiVolume24H                           float64 `json:"defi_volume_24h"`
	DefiVolume24HReported                   float64 `json:"defi_volume_24h_reported"`
	Defi24HPercentageChange                 float64 `json:"defi_24h_percentage_change"`
	DefiMarketCap                           float64 `json:"defi_market_cap"`
	StablecoinVolume24H                     float64 `json:"stablecoin_volume_24h"`
	StablecoinVolume24HReported             float64 `json:"stablecoin_volume_24h_reported"`
	Stablecoin24HPercentageChange           float64 `json:"stablecoin_24h_percentage_change"`
	StablecoinMarketCap                     float64 `json:"stablecoin_market_cap"`
	DerivativesVolume24H                    float64 `json:"derivatives_volume_24h"`
	DerivativesVolume24HReported            float64 `json:"derivatives_volume_24h_reported"`
	Derivatives24HPercentageChange          float64 `json:"derivatives_24h_percentage_change"`
	LastUpdated                             string  `json:"last_updated"`
	TotalMarketCapYesterday                 float64 `json:"total_market_cap_yesterday"`
	TotalVolume24HYesterday                 float64 `json:"total_volume_24h_yesterday"`
	TotalMarketCapYesterdayPercentageChange float64 `json:"total_market_cap_yesterday_percentage_change"`
	TotalVolume24HYesterdayPercentageChange float64 `json:"total_volume_24h_yesterday_percentage_change"`
}
