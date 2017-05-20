package main

import (
	"fmt"
	"os"
	"strings"
)

// 处理正整数部分
func comma1(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}

	return comma1(s[:n-3]) + "," + s[n-3:]
}

// 处理小数部分
func comma2(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}

	return s[:3] + "," + comma2(s[3:])
}

func comma(s string) string {
	var strs []string

	strs = strings.Split(s, ".")

	l := len(strs)

	if l > 1 {
		if strings.Index(strs[0], "-") >= 0 {
			return "-" + comma1(strs[0][1:]) + "." + comma2(strs[1])
		} else {
			return comma1(strs[0]) + "." + comma2(strs[1])
		}
	} else {
		if strings.Index(strs[0], "-") >= 0 {
			return "-" + comma1(strs[0][1:])
		} else {
			return comma1(strs[0])
		}
	}
}

func main() {
	for _, arg := range os.Args[1:] {
		fmt.Printf("%s\n", comma(arg))
	}
}
