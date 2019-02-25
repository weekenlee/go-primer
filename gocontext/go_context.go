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

	go sleepRandom("sleepRandomContext", sleeptimechan)

	select {
	case <-ctx.Done():
		fmt.Println("sleepRandomContext: time to return")
	case sleeptime := <-sleeptimechan:
		fmt.Println("Slept for ", sleeptime, "ms")
	}
}

func doWorkContext(ctx context.Context) {
	ctxWithTimeout, cancelFunction := context.WithTimeout(ctx, time.Duration(150)*time.Millisecond)

	defer func() {
		fmt.Println("doWorkContext complete")
		cancelFunction()
	}()

	ch := make(chan bool)
	go sleepRandomContext(ctxWithTimeout, ch)

	select {
	case <-ctx.Done():
		fmt.Println("doWorkContext: time to return ")
	case <-ch:
		fmt.Println("sleepRandomContext returned")
	}
}

func main() {
	ctx := context.Background()

	ctxWithCancle, cancelFunction := context.WithCancel(ctx)

	defer func() {
		fmt.Println("main defer: canceling context")
	}()

	go func() {
		sleepRandom("main", nil)
		cancelFunction()
		fmt.Println("Main Sleep complete, canceling, context")
	}()

	doWorkContext(ctxWithCancle)
}
