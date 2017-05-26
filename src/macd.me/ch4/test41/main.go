package main

import (
	"crypto/sha256"
	"fmt"
	"os"

	"macd.me/ch2/popcount"
)

func cmpDiff(a []byte, b []byte) int {
	// 算出哈希值
	c1 := sha256.Sum256(a)
	c2 := sha256.Sum256(b)

	// 计算两个字节数组的异或
	var c3 [len(c1)]byte
	for i := 0; i < len(c1); i++ {
		c3[i] = c1[i] ^ c2[i]
	}
	return popcount.PopCountByts(c3)
}

func main() {
	if len(os.Args) == 3 {
		fmt.Printf("字符串：%s和字符串：%s的哈希值中有%d个bit不同\n", os.Args[1], os.Args[2], cmpDiff([]byte(os.Args[1]), []byte(os.Args[2])))
	}
}
