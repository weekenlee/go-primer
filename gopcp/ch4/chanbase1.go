package main

import (
	"fmt"
	"time"
)

var strChan = make(chan string, 3)

func main() {
	syncChan1 := make(chan struct{}, 1)
	syncChan2 := make(chan struct{}, 2)

	go func() {
		<-syncChan1
		fmt.Println("Received a sync signal and wait a second....[recevier]")
		time.Sleep(time.Second)
		for {
			if elem, ok := <-strChan; ok {
				fmt.Println("Received: ", elem, "[recevier]")
			} else {
				break
			}
		}
		fmt.Println("Stopped. [receiver]")
		syncChan2 <- struct{}{}
	}()

	go func() {
		for _, elem := range []string {"a", "b", "c", "d"} {
			strChan <- elem
			fmt.Println("Send: ", elem, "[sender]")
			if elem == "c" {
				syncChan1 <- struct{}{}	
				fmt.Println("Send a sync signal. [sender]")
			}
		}
		fmt.Println("Wait 2 seconds .. [sender]")
		time.Sleep(time.Second * 2)
		close(strChan)
		syncChan2 <- struct{}{}
	}()

	<-syncChan2
	<-syncChan2
}
