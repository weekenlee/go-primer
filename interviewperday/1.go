package main

import "fmt"

func defer_call() {
	defer func() { fmt.Println("1")} ()
	defer func() { fmt.Println("2")} ()
	defer func() { fmt.Println("3")} ()

	panic("error")
}

func main() {
	defer_call()
}
