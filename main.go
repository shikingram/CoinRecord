package main

import (
	"CoinRecord/coinPrice"
	"CoinRecord/models"
	"embed"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

const StartTime = "2021-08-19"

//go:embed template/*
var fs embed.FS

func main() {

	err := coinPrice.GetAllCoinPrice()
	err = coinPrice.GetCoinMarketCap()
	if err != nil {
		log.Fatal(err)
		return
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
			switch coin.Operate {
			case "+":
				coinAmount += coin.Amount
				coinCost += coin.Sum
			case "-":
				coinAmount -= coin.Amount
				coinCost -= coin.Sum
			}
		}
		name = v.Name
		// 单个币平均成本价格 = 总成本/币的数量
		avgPrice = coinCost / coinAmount
		coinInfo := coinPrice.GetCoinPrice(name)
		priceUsd = ProcessFloat(coinInfo.Quote.USD.Price)
		coinSum = priceUsd * coinAmount
		// 利润 = 持仓现价 - 投资总额
		profit = coinSum - coinCost
		// 收益率 = 利润/投资总额
		yield = profit / coinCost
		coinInfoVo := models.CoinInfoVo{
			Name:     name,
			Amount:   ProcessFloat(coinAmount),
			Sum:      ProcessFloat(coinSum),
			NowPrice: ProcessFloat(priceUsd),
			AvgPrice: ProcessFloat(avgPrice),
			Cost:     ProcessFloat(coinCost),
			Profit:   ProcessFloat(profit),
			Yield:    fmt.Sprintf("%.2f", yield*100) + "%",
			CmcRank: coinInfo.CmcRank,
		}
		coinInfoVo.MarketCap = ConvertFloatToString(coinInfo.Quote.USD.MarketCap)
		coinInfoVo.Volume24H = ConvertFloatToString(coinInfo.Quote.USD.Volume24H)
		per24h := coinInfo.Quote.USD.PercentChange24H
		per7d := coinInfo.Quote.USD.PercentChange7D
		coinInfoVo.PercentChange24H= fmt.Sprintf("%.2f",per24h)+"%"
		coinInfoVo.PercentChange7D= fmt.Sprintf("%.2f",per7d)+"%"
		//原单价 = 现单价 / （1+昨日涨幅）
		//（现单价 - 原单价） * 持有数量 = 昨日收益
		yesPrice := coinInfo.Quote.USD.Price / (1+(per24h * 0.01))
		coinInfoVo.Profit24h = ProcessFloat((coinInfo.Quote.USD.Price - yesPrice) * coinAmount)

		coinInfoVo.Change24hColor = models.ShallowRed
		coinInfoVo.DeepChange24hColor = models.Red
		coinInfoVo.YieldColor = models.ShallowRed
		coinInfoVo.DeepYieldColor = models.Red
		coinInfoVo.DeepChange7DColor = models.Red
		if per24h > 0{
			coinInfoVo.Change24hColor = models.ShallowGreen
			coinInfoVo.DeepChange24hColor = models.Green
		}

		if per7d > 0 {
			coinInfoVo.DeepChange7DColor = models.Green
		}

		if yield > 0 {
			coinInfoVo.YieldColor = models.ShallowGreen
			coinInfoVo.DeepYieldColor = models.Green
		}
		coinInfoVos = append(coinInfoVos, coinInfoVo)
	}

	mdDir := filepath.Join(currentDir, "records")
	_ = os.MkdirAll(mdDir, 0666)
	mdFileName := time.Now().Format("20060102150405") + ".md"
	mdFilePath := filepath.Join(mdDir, mdFileName)
	fmt.Println(mdFilePath)
	mdFile, err := os.OpenFile(mdFilePath, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer mdFile.Close()

	var totalProfit float64
	var totalCost float64
	var totalYield float64
	var totalSum float64
	var yesProfit float64
	for _, v := range coinInfoVos {
		totalProfit += v.Profit
		totalCost += v.Cost
		totalSum += v.Sum
		yesProfit += v.Profit24h
	}
	totalYield = totalProfit / totalCost
	totalYieldString := fmt.Sprintf("%.2f", totalYield*100) + "%"
	var templateValue = models.TemplateValue {
		StartTime:   StartTime,
		NowTime:     TimeFormat(time.Now()),
		NowDay:      time.Now().Format("2006-01-02"),
		TotalProfit: ProcessFloat(totalProfit),
		TotalCost:   ProcessFloat(totalCost),
		TotalYield:  totalYieldString,
		TotalSum:    ProcessFloat(totalSum),
	}
	templateValue.IncomeData = coinInfoVos
	templateValue.YesterdayProfit = yesProfit
	templateValue.Days = timeSub(time.Now(), StringToTime(StartTime))
	templateValue.YesterdayProfitColor = models.Red
	templateValue.TotalYieldClolor = models.Red
	if totalYield > 0 {
		templateValue.TotalYieldClolor = models.Green
	}
	if templateValue.YesterdayProfit> 0 {
		templateValue.YesterdayProfitColor = models.Green
	}
	//保存行情信息

	templateValue.TotalMarketCap = ConvertFloatToString(coinPrice.GRes.Data.Quote.USD.TotalMarketCap)
	tmcpc := coinPrice.GRes.Data.Quote.USD.TotalMarketCapYesterdayPercentageChange
	if tmcpc > 0 {
		templateValue.TotalMarketColor = models.Green
	} else {
		templateValue.TotalMarketColor = models.Red
	}
	templateValue.TotalMarketCapYesterdayPercentageChange = fmt.Sprintf("%.2f",tmcpc)+"%"

	templateValue.TotalVolume24h = ConvertFloatToString(coinPrice.GRes.Data.Quote.USD.TotalVolume24H)
	tvcpc := coinPrice.GRes.Data.Quote.USD.TotalVolume24HYesterdayPercentageChange
	if tvcpc > 0 {
		templateValue.TotalVolumeColor = models.Green
	} else {
		templateValue.TotalVolumeColor = models.Red
	}
	templateValue.TotalVolume24hYesterdayPercentageChange = fmt.Sprintf("%.2f",tmcpc)+"%"

	// 把模板编进二进制文件中
	tmpl, err := template.ParseFS(fs, "template/*.tmpl")
	if err != nil {
		log.Fatal(err)
		return
	}
	err = tmpl.Execute(mdFile, templateValue)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println("按任意键继续...")
	var input string
	_, _ = fmt.Scanln(&input)
}
