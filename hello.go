package main

import (
	"fmt"
	"os"
	"pra"
)

func main() {
	fmt.Println("hello world")
	fmt.Println(pra.Even(2))

	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}
