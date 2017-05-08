package main

import (
	"fmt"
	"time"
)

func main() {
	var s, sep string
	fmt.Println(time.Now())
	for i := 1; i < 100000; i++ {
		s += sep + "os.Args[i]"
		sep = " "
	}
	fmt.Println(time.Now())
}
