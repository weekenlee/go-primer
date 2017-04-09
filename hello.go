package main

import (
	"fmt"
	"os"
	"pra"
	"strings"
)

func main() {
	var hi pra.Mytype = 10.0
	fmt.Println(hi)
	fmt.Println("hello world")
	fmt.Println(pra.Even(2))

	//打印命令行 a
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)

	//打印命令行 b
	sep, s = "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)

	//打印命令行 c
	fmt.Println(strings.Join(os.Args[1:], " "))

	var iu8 uint8 = 255
	fmt.Println(iu8, iu8+1)
	var i int8 = 127
	fmt.Println(i, i+1)

	var x uint8 = 1<<2 | 1<<3
	fmt.Println(x)
	fmt.Printf("%08b\n", x)
	a := fmt.Sprintf("%08b\n", x)
	fmt.Println(a)
}
