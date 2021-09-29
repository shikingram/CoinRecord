package main

import (
	"fmt"
	"math/big"
	"strconv"
)

func ProcessFloat(f float64) float64 {
	num2, _ := strconv.ParseFloat(fmt.Sprintf("%.8f", f), 64)
	return num2
}

//big.NewRat(1, 1).SetFloat64(coinPrice.GRes.Data.Quote.USD.TotalVolume24H).FloatString(0)
func ConvertFloatToString(f float64) string {
	return big.NewRat(1, 1).SetFloat64(f).FloatString(0)
}
