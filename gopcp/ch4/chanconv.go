package main

import "fmt"

func main() {
	var ok bool
	ch := make(chan int, 1)
	_, ok = interface{}(ch).(<-chan int)
	fmt.Println("chan int => <-chan int:", ok)
	_, ok = interface{}(ch).(chan<- int)
	fmt.Println("chan int => chan<- int:", ok)

	sch := make(chan<-int, 1)
	_, ok = interface{}(sch).(chan int)
	fmt.Println("chan<-int => chan int:", ok)

	cch := make(<-chan int, 1)
	_, ok = interface{}(cch).(chan int)
	fmt.Println("<-chan int => chan int:", ok)
}
