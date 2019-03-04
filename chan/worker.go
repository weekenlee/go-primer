package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, result chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "start job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finish job", j)
		result <- j * 2
	}
}

func main() {
	jobs := make(chan int, 100)
	result := make(chan int, 100)

	for w := 1; w <= 3; w++ {
		go worker(w, jobs, result)
	}

	for i := 1; i < 5; i++ {
		jobs <- i
	}

	close(jobs)

	for i := 1; i < 5; i++ {
		<-result
	}
}
