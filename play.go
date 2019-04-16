package main

import (
	"fmt"
	"unicode"
)

func main() {
	var a interface{}
	var b string
	a = "123"
	b = a.(string)
	fmt.Println(b)
	fmt.Println(b[0])
	fmt.Println(unicode.IsDigit(rune(b[0])))
}
