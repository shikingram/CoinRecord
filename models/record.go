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

type CoinInfoVo struct {
	//币种
	Name string `json:"name"`
	//持有数量
	Amount float64 `json:"amount"`
	//现价
	NowPrice float64 `json:"nowPrice"`
	//现在的总金额
	Sum float64 `json:"sum"`
	//成本单价
	AvgPrice float64 `json:"avgCost"`
	//成本
	Cost float64 `json:"cost"`
	//利润
	Profit float64 `json:"profit"`
	//收益率
	Yield float64 `json:"yield"`
}
