package chart

import (
	"fmt"
	"testing"
)

func TestRandomColor(t *testing.T) {
	a := getRandomColorArray(6)
	fmt.Println(len(a))
	fmt.Printf("%+v", a)
}
