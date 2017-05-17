package main

import (
	"fmt"
	"os"
)

func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}

	return comma(s[:n-3]) + "," + s[n-3:]
}

func main() {
	for _, arg := range os.Args[1:] {
		fmt.Printf("%s ", comma(arg))
	}
	fmt.Println()
}
