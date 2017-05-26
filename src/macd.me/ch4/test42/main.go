package main

import (
	"crypto/sha256"
	"fmt"
	"os"
)

func main() {
	for i, arg := range os.Args[1:] {
		fmt.Printf("%d: %x\n", i, sha256.Sum256([]byte(arg)))
	}
}
