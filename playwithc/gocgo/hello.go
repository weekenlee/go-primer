package main

import "C"
import "fmt"

//export HelloFromGo
func HelloFromGo() {
	fmt.Printf("hello from go\n")
}
