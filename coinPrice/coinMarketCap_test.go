package coinPrice

import (
	"CoinRecord/httpUtil"
	"encoding/json"
	"fmt"
	"testing"
)

func TestGetCoinPrice(t *testing.T) {
	url := "https://pro-api.coinmarketcap.com/v1/global-metrics/quotes/latest"
	m := make(map[string]string)
	m2 := make(map[string]string)
	//m["start"] = "1"
	//m["limit"] = "100"
	b, err := httpUtil.HttpGet(url, m, m2)
	if err != nil {
		fmt.Printf("%+v", err)
	}
	//fmt.Println(string(b))
	res := &GlobalMetricsResult{}
	err = json.Unmarshal(b, res)
	if err != nil {
		fmt.Printf("%+v",err)
	}
	fmt.Printf("%+v",res)
}
