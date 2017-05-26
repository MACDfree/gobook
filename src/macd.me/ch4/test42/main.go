package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"os"
)

var t = flag.String("t", "256", "哈希算法类型，默认sha256（参数值为256），可选sha384（参数值为384）、sha512（参数值为512）")

func main() {
	flag.Parse()
	for _, arg := range flag.Args() {
		switch *t {
		case "256":
			fmt.Printf("%s: %x\n", arg, sha256.Sum256([]byte(arg)))
		case "384":
			fmt.Printf("%s: %x\n", arg, sha512.Sum384([]byte(arg)))
		case "512":
			fmt.Printf("%s: %x\n", arg, sha512.Sum512([]byte(arg)))
		default:
			os.Exit(1)
		}
	}
}
