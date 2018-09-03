package main

import (
	"fmt"
	"time"
)

func main() {
	names := []string { "Eric", "harry", "Robert", "Jim", "Mark"}
	for _, name := range names {
		go func(who string) {
			fmt.Printf("Hello %s\n", who)
		}(name)
	}
	time.Sleep(time.Millisecond)
}
