package main

import (
	"fmt"
	"time"
)

func f1(c chan int) {
	time.Sleep(1 * time.Second)
	c <- 5
}

func f2(c chan int) {
	time.Sleep(15 * time.Second)
	c <- 15 
}

func main() {
	c := make(chan int, 2)

	ticker := time.NewTicker(3 * time.Second) 
	defer ticker.Stop()

	go f1(c)
	go f2(c)

	count := 0

	for {
		select {
		case c1:= <-c:
			fmt.Println()
			fmt.Println(c1)
			count++
			if count == 2 {
				goto done
			}
		case c2:= <-c:
			fmt.Println()
			fmt.Println(c2)
			count++
			if count == 2 {
				goto done
			}
		default:
			fmt.Printf(".")
		}

		select {
		case <- ticker.C:
			fmt.Println("\nTime out!")
			return
		default:
		}
	}

done:
	fmt.Println("done")
}
