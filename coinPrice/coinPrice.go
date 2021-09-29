package coinPrice

import (
	"CoinRecord/httpUtil"
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

// coincap.io
//var url = "https://api.coincap.io/v2/assets"
var url = "https://pro-api.coinmarketcap.com/v1/cryptocurrency/listings/latest"
var coinMarketCapUrl = "https://pro-api.coinmarketcap.com/v1/global-metrics/quotes/latest"

var res CoinMarketCapLastListResult
var GRes GlobalMetricsResult

func GetAllCoinPrice() error {
	//api比较卡，简单加个重试机制
	var count int
	var params = make(map[string]string)

	params["limit"] = "1000"
	params["start"] = "1"
	var headers = make(map[string]string)
	for {
		if count > 10 {
			break
		}
		resBytes, err := httpUtil.HttpGet(url, params, headers)

		if err != nil {
			count++
			fmt.Printf("%+v \n", err)
			time.Sleep(time.Millisecond * 10)
			continue
		}

		err = json.Unmarshal(resBytes, &res)
		if err != nil {
			count++
			fmt.Printf("%+v \n", err)
			time.Sleep(time.Millisecond * 10)
			continue
		}
		break
	}
	return nil
}

func GetCoinPrice(id string) SingleCoinInfo {
	for _, coinInfo := range res.Data {
		if coinInfo.Name == id || coinInfo.Slug == id || coinInfo.Symbol == strings.ToUpper(id) {
			return coinInfo
		}
	}

	return SingleCoinInfo{}
}

func GetCoinMarketCap() error {
	var params = make(map[string]string)
	var headers = make(map[string]string)
	resBytes, err := httpUtil.HttpGet(coinMarketCapUrl, params, headers)
	if err != nil {
		return err
	}
	err = json.Unmarshal(resBytes, &GRes)
	if err != nil {
		return err
	}
	return nil
}
