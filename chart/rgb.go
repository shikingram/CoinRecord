package chart

import "math/rand"

var aryNum = []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "A", "B", "C", "D", "E", "F"}

func RandomColor() string {
	var str = "#"
	for i := 0; i < 6; i++ {
		random := rand.Intn(16)
		str += aryNum[random]
	}
	return str
}
