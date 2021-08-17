package main

import (
	"CoinRecord/coinPrice"
	"CoinRecord/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

func main() {

	err := coinPrice.GetAllCoinPrice()
	if err != nil {
		log.Fatal(err)
	}

	var getCurrentDirectory = func() string {
		dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
		if err != nil {
			log.Fatal(err)
		}
		return strings.Replace(dir, "\\", "/", -1)
	}

	currentDir := getCurrentDirectory()
	fmt.Println("current DIR:", currentDir)

	holdsDir := filepath.Join(currentDir, "holds")

	rd, err := ioutil.ReadDir(holdsDir)
	if err != nil {
		log.Fatal(err)
		return
	}

	var allHolds []models.Holds
	for _, fi := range rd {
		if fi.IsDir() {
			continue
		}
		fmt.Println(fi.Name())
		filePath := filepath.Join(holdsDir, fi.Name())
		f, err := os.Open(filePath)
		if err != nil {
			log.Fatal(err)
			return
		}
		content, err := ioutil.ReadAll(f)
		if err != nil {
			log.Fatal(err)
			return
		}
		var holds models.Holds
		err = json.Unmarshal(content, &holds)
		if err != nil {
			log.Fatal(err)
			return
		}
		fmt.Printf("%+v \n", holds)
		allHolds = append(allHolds, holds)
	}

	var coinInfoVos []models.CoinInfoVo
	for i, v := range allHolds {
		fmt.Printf("id: %d,value:%+v \n", i, v)
		//遍历所有记录计算投资总额，成本价，币的数量
		var coinAmount float64
		var name string
		var avgPrice float64
		var coinCost float64
		var coinSum float64
		var profit float64
		var yield float64
		var priceUsd float64
		for _, coin := range v.Records {
			name = coin.Name
			// 单个币平均成本价格 = 成本价/币的数量
			avgPrice = coin.Sum / coin.Amount
			switch coin.Operate {
			case "+":
				coinAmount += coin.Amount
				coinCost += coin.Sum
			case "-":
				coinAmount -= coin.Amount
				coinCost -= coin.Sum
			}
			coinInfo := coinPrice.GetCoinPrice(name)
			priceUsd,err = strconv.ParseFloat(coinInfo.PriceUsd,64)
			if err != nil {
				log.Fatal(err)
			}
			coinSum = priceUsd * coinAmount
			// 利润 = 持仓现价 - 投资总额
			profit = coinSum - coinCost
			// 收益率 = 利润/投资总额
			yield = profit / coinCost
		}
		coinInfoVo := models.CoinInfoVo{
			Name:     name,
			Amount:   coinAmount,
			Sum:      coinSum,
			NowPrice: priceUsd,
			AvgPrice: avgPrice,
			Cost:     coinCost,
			Profit:   profit,
			Yield:    yield,
		}
		coinInfoVos = append(coinInfoVos, coinInfoVo)
	}

	mdDir := filepath.Join(currentDir, "records")
	_ = os.MkdirAll(mdDir,0666)
	mdFileName := time.Now().Format("20060102150405") + ".md"
	mdFilePath := filepath.Join(mdDir,mdFileName)
	fmt.Println(mdFilePath)
	mdFile,err := os.OpenFile(mdFilePath,os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer mdFile.Close()
	_,err = mdFile.WriteString(models.Head+"\r")
	if err != nil {
		log.Fatal(err)
	}
	for _, v := range coinInfoVos {
		var value = fmt.Sprintf(models.ValueFrame,v.Name,v.Amount,v.NowPrice,v.Sum,v.AvgPrice,v.Cost,v.Profit,fmt.Sprintf("%.2f",v.Yield*100)+"%")
		var str = models.Partition +"\r"+ value
		fmt.Println(str)
		mdFile.WriteString(str)
	}

	fmt.Println("按任意键继续...")
	var input string
	_, _ = fmt.Scanln(&input)
}