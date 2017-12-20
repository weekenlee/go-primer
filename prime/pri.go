package main

import (
	"fmt"
)

const SIZE = 100

func generate(ch chan<- int) {
	for i := 2; ; i++ {
		ch <- i
	}
}

func filter(src <-chan int, des chan<- int, prime int) {
	for i := range src {
		if i%prime != 0 {
			des <- i
		}
	}
}

func main() {
	ch := make(chan int)
	go generate(ch)
	for {
		prime := <-ch
		fmt.Println(prime)
		if prime > SIZE {
			break
		}
		ch1 := make(chan int)
		go filter(ch, ch1, prime)
		ch = ch1
	}
}
