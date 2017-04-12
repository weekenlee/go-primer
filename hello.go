package main

import (
	"fmt"
	"os"
	"pra"
	"strings"
	"unicode/utf8"
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

	八进制 := 0666
	fmt.Printf("%d %[1]o %#[1]o\n", 八进制)

	十六 := 0xabc
	fmt.Printf("%d %[1]x %#[1]x\n", 十六)

	sw := "Hello, 世界"
	fmt.Println(len(sw))                    // "13"
	fmt.Println(utf8.RuneCountInString(sw)) // "9"

	for i := 0; i < len(sw); {
		r, size := utf8.DecodeRuneInString(sw[i:])
		fmt.Printf("%d\t%c\n", i, r)
		i += size
	}

	//Go语言的range循环在处理字符串的时候，会自动隐式解码UTF8字符串
	for i, r := range "Hello, 世界" {
		fmt.Printf("%d\t%q\t%d\n", i, r, r)
	}

	filepath := "/a/b/c.go"
	fmt.Println(basename(filepath))
}

func basename(s string) string {
	for i := len(s) - 1; i > 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}

	return s
}
