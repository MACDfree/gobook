package main

import (
	"fmt"
)

type Celsius float64    // 摄氏度
type Fahrenheit float64 // 华氏度

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

//
func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

//
func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

func (c Celsius) String() string {
	return fmt.Sprintf("%gC", c)
}

func main() {
	fmt.Println(CToF(123))
	fmt.Println(FToC(32))
}
