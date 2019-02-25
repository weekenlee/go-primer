package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func sleepRandom(fromFunction string, ch chan int) {
	defer func() { 
		fmt.Println(fromFunction, "sleepRandom complete")
	}()

	seed := time.Now().UnixNano()
	r := rand.New(rand.NewSource(seed))
	randomNumber := r.Intn(100)
	sleeptime := randomNumber + 100
	fmt.Println(fromFunction, "starting sleep for", sleeptime, "ms")
	time.Sleep(time.Duration(sleeptime) * time.Millisecond)
	fmt.Println(fromFunction, "walking up , slept for", sleeptime, "ms")
	
	if ch != nil {
		ch <- sleeptime
	}
}

func sleepRandomContext(ctx context.Context, ch chan bool) {
	defer func() {
		fmt.Println("SleepRandomContext complete")
		ch <- true
	}()

	sleeptimechan := make(chan int)

	go sleep
}