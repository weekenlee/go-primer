package main

import "fmt"

/*
#cgo CFLAGS: -I./
#cgo LDFLAGS: -L./ -lhi
#include "hi.c" // 可以是头文件，也可以是c文件, 严格说要是c文件。
*/
import "C"

func main() {
	C.hi()
	fmt.Println("Hi, vim-go")
}
