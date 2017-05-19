package main

import (
	"bytes"
	"fmt"
	"os"
)

func comma(s string) string {
	n := len(s)
	t := n % 3
	a := 0
	var buf bytes.Buffer

	if t == 0 {
		t = 3
	}

	for i := t; i <= n; i += 3 {
		buf.WriteString(s[a:i])
		buf.WriteByte(',')
		a = i
	}

	return buf.String()[:buf.Len()-1]
}

func main() {
	for _, arg := range os.Args[1:] {
		fmt.Printf("%s ", comma(arg))
	}
	fmt.Println()
}
