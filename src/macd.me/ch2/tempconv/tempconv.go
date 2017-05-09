// Package tempconv test package doc
package tempconv

import (
	"fmt"
)

// Celsius 摄氏度
type Celsius float64

// Fahrenheit 华氏度
type Fahrenheit float64

const (
	// AbsoluteZeroC test
	AbsoluteZeroC Celsius = -273.15
	// FreezingC test
	FreezingC Celsius = 0
	// BoilingC test
	BoilingC Celsius = 100
)

func (c Celsius) String() string {
	return fmt.Sprintf("%gC", c)
}

func (f Fahrenheit) String() string {
	return fmt.Sprintf("%gF", f)
}
