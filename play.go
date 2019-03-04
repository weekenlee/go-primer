package main

import "fmt"

func main() {
	var a interface{}
	var b string
	a = "123"
	b = a.(string)
	fmt.Println(b)
}
