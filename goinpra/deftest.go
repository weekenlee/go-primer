package main

import "fmt"

func main() {
	defer_call()
}

//defer 以stack形式压入运行
func defer_call() {
	defer func() {fmt.Println("打印前")}()
	defer func() {fmt.Println("打印中")}()
	defer func() {fmt.Println("打印后")}()
	panic("异常")
}
