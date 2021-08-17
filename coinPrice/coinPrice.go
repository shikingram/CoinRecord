package coinPrice

import (
	"CoinRecord/httpUtil"
	"CoinRecord/models"
	"encoding/json"
	"fmt"
	"time"
)

var res models.Result

func GetAllCoinPrice() error {
	//api比较卡，简单加个重试机制
	var count int
	var params = make(map[string]string)
	var headers = make(map[string]string)

	for {
		if count > 10 {
			break
		}
		resBytes, err := httpUtil.HttpGet("https://api.coincap.io/v2/assets", params, headers)
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

func GetCoinPrice(id string) models.CoinInfo {
	for _, coinInfo := range res.CoinInfos {
		if coinInfo.Name == id || coinInfo.ID == id {
			return coinInfo
		}
	}

	return models.CoinInfo{}
}
