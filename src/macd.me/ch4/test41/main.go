package main

import (
	"crypto/sha256"
	"fmt"
	"os"
	"strconv"

	"macd.me/ch2/popcount"
)

func cmpDiff(a []byte, b []byte) int {
	c1 := sha256.Sum256(a)
	c2 := sha256.Sum256(b)
	fmt.Printf("%x , %x\n", c1, c2)
	t1, err := strconv.ParseUint(fmt.Sprintf("%x", c1), 16, 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
	t2, err := strconv.ParseUint(fmt.Sprintf("%x", c2), 16, 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
	fmt.Printf("%x ^ %x = %x\n", t1, t2, t1^t2)
	return popcount.PopCount(t1 ^ t2)
}

func main() {
	if len(os.Args) == 3 {
		fmt.Printf("字符串：%s和字符串：%s的哈希值中有%d个bit不同\n", os.Args[1], os.Args[2], cmpDiff([]byte(os.Args[1]), []byte(os.Args[2])))
	}
}
