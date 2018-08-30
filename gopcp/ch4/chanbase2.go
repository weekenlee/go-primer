package main

import (
	"fmt"
	"time"
)
var strChan = make(chan string , 3)

func main() {
	syncChan1 := make(chan struct{}, 1)
	syncChan2 := make(chan struct{}, 2)
	go receive(strChan, syncChan1, syncChan2)
	go send(strChan, syncChan1, syncChan2)
	<-syncChan2
	<-syncChan2
}

func receive(strChan <-chan string,
			syncChan1 <-chan struct{},
			syncChan2 chan<- struct{}) {
	fmt.Println("Received a sync signal and wait a second...[recevier]")
	time.Sleep(time.Second)
	for {
		if elem, ok := <- strChan; ok {
			fmt.Println("Received: ", elem, "[receiver]")
		} else {
			break
		}
	}
	fmt.Println("stop. [receiver]")
	syncChan2 <- struct{}{}
}

func send(strChan chan<- string, 
syncChan1 chan<- struct{}, 
syncChan2 chan<- struct{}) {
	for _, elem := range []string{"a", "b", "c", "d"} {
		strChan <- elem
		fmt.Println("Send: ", elem, "[sender]")
		if elem == "c" {
			syncChan1 <- struct{} {}
			fmt.Println("Send a sync signal .[sender]")
		}
	}
	fmt.Println("wait 2 seconds...[sender]")
	time.Sleep(time.Second * 2)
	close(strChan)
	syncChan2 <- struct{}{}
}
