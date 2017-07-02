package main

import (
	"fmt"

	"macd.me/ch6/intset"
)

func main() {
	var x, y intset.IntSet
	x.Add(1)
	x.Add(2)
	x.Remove(1)

	y.Add(0)
	y.Add(3)
	fmt.Println(x.String())
	fmt.Println(y.String())
}
