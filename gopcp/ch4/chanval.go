package main

import (
	"fmt"
	"time"
)

type Counter struct {
	count int
}

var mapChan = make(chan map[string]Counter , 1)

func main() {
	syncChan := make(chan struct {}, 2)

	go func() {
		for {
			if elem, ok := <- mapChan; ok {
				counter := elem["count"]
				counter.count++
				fmt.Printf("%p\n", &counter)
			} else {
				break
			}
		}

		fmt.Println("stoped [receiver]")
		syncChan <- struct{}{}
	}()

	go func() {
		countMap := map[string]Counter {
			"count": Counter{},
		}

		c := countMap["count"]
		fmt.Printf("%p\n", &c)
		for i := 0;  i< 5; i++ 	{
			mapChan <- countMap
			time.Sleep(time.Millisecond)
			fmt.Printf("The count map: %v. [sender]\n",countMap)
		}
		close(mapChan)
		syncChan <- struct{}{}
	}()

	<-syncChan
	<-syncChan
}
