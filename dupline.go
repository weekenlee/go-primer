package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		if input.Text() == "quit" {
			break
		}
		counts[input.Text()]++
	}
	fmt.Println(counts)
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", line, n)
		}
	}
}
