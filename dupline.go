package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	/* dup 1
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		if input.Text() == "quit" {
			break
		}
		counts[input.Text()]++
	}
	*/

	//dup 2
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		coutline(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			coutline(f, counts)
			f.Close()
		}
	}

	fmt.Println(counts)
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func coutline(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
}
