package main

import "fmt"

type WG struct {
	main    chan func()
	allDone chan bool
}

func New(n int) WG {
	res := WG {
		main: make(chan func()),
		allDone: make(chan bool),
	}

	procDone := make(chan bool)
	for i := 0; i < n; i++ {
		go func() {
			for {
				f := <-res.main
				if f == nil {
					procDone <- true
					return
				}
				f()
			}
		}()
	}

	go func() {
		for i := 0; i < n; i++ {
			_ = <- procDone
		}
		res.allDone <- true
	}()

	return res
}

func (wg WG) Add(f func()) {
	wg.main <- f
}

func (wg WG) Wait() {
	close(wg.main)
	<- wg.allDone
}

func main() {
	wg := New(50) //50个worker
	for i := 0; i < 1000; i++ {
		wg.Add(func() {
			fmt.Println("func run")
		})
	}
	wg.Wait()
}
