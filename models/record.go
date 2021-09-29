package models

type Holds struct {
	Name    string   `json:"name"`
	Records []Record `json:"records"`
}

type Record struct {
	Operate string  `json:"operate"`
	Amount  float64 `json:"amount"`
	Sum     float64 `json:"sum"`
}

const (
	Green        = "#00EC00"
	ShallowGreen = "#F0FFF0"
	Red          = "#FF0000"
	ShallowRed   = "#FFECEC"
)

type CoinInfoVo struct {
	//币种
	Name string
	//持有数量
	Amount float64
	//现价
	NowPrice float64
	//现在的总金额
	Sum float64
	//成本单价
	AvgPrice float64
	//成本
	Cost float64
	//利润
	Profit float64
	//收益率
	Yield string
	//颜色
	YieldColor         string
	DeepYieldColor     string
	Change24hColor     string
	DeepChange24hColor string
	DeepChange7DColor  string
	//排名
	CmcRank int
	//市值
	MarketCap string
	//24小时交易量
	Volume24H string
	//24小时变化
	PercentChange24H string
	//7天变化
	PercentChange7D string
	Profit24h       float64
}

type TemplateValue struct {
	// 总市值
	TotalMarketCap string
	//总交易量
	TotalVolume24h   string
	TotalVolumeColor string
	TotalMarketColor string
	//较昨日总交易量变化百分比
	TotalVolume24hYesterdayPercentageChange string
	//较昨日总市值变化百分比
	TotalMarketCapYesterdayPercentageChange string
	// defi总市值
	DefiMarketCap float64
	//defi交易量
	DefiVolume24h float64
	//defi总市值变化
	Defi24hPercentageChange string
	//开始时间
	StartTime string
	//定投天数
	Days int
	//当前时间
	NowTime string
	//今天日期
	NowDay string
	//24h收益
	YesterdayProfit      float64
	YesterdayProfitColor string
	//总利润
	TotalProfit float64
	//总成本
	TotalCost float64
	//总收益率
	TotalYield string
	//总价值
	TotalSum float64
	//字体颜色
	TotalYieldClolor string
	//收益数据
	IncomeData []CoinInfoVo
	//pieString
	PieContent string
	BarContent string
}
