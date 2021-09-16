package chart

import (
	"fmt"
	"testing"
)

func TestRandomColor(t *testing.T) {

	for i := 0; i < 5; i++ {
		a := getRandomColorArray(6)
		fmt.Println(len(a))
		fmt.Printf("%+v", a)
	}
}
