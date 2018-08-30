package main

import "fmt"

var quit chan int

func foo(id int) {
	fmt.Println(id)
	quit <- 0
}

func foo2(ch chan int) {
	id := <-ch
	fmt.Println(id)
	quit <- 0
}

func main() {
	count := 5
	quit = make(chan int , 20)
	ch := make(chan int)

	for i:=0; i<count; i++ {
		go foo2(ch)
		ch <- i
	}

	for i:=0; i<count; i++ {
		go foo(i)
	}

	for i:=0; i< count*2; i++ {
		<-quit
	}
}
