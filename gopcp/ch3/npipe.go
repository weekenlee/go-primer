package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

func main() {
	fileBasedPipe()
	inMemorySyncPipe()
}

func fileBasedPipe() {
	reader, writer, err := os.Pipe()
	if err != nil {
		fmt.Printf("Error: could not create the named pipe: %s\n", err)
	}
	go func() {
		output := make([]byte, 10000)
		n, err := reader.Read(output)
		if err != nil {
			fmt.Printf("Error: could not read data from the named pipe: %s\n", err)
		}
		fmt.Printf("Read %d bytes. file base\n", n);
		//fmt.Println(output)
	}()

	input := make([]byte, 10000)
	for i:= 65; i < 10065; i++ {
		input[i - 65] = byte(i)
	}

	n, err := writer.Write(input)
	if err != nil {
		fmt.Printf("Error: could not write data to the name pipe %s \n", err)
	}

	fmt.Printf("Written %d bytes. file base pipe\n", n)
	time.Sleep(200 * time.Millisecond)
}

func inMemorySyncPipe() {
	reader, writer := io.Pipe()
	go func() {
		output := make([]byte, 100)
		n, err := reader.Read(output)
		if err != nil {
			fmt.Printf("Error: could not read data from name pipe %s \n", err)
		}
		fmt.Printf("Read %d bytes . in memory pipe\n", n)
	}()

	input := make([]byte, 26)
	for i := 65; i <= 90; i++ {
		input[i - 65] = byte(i)
	}

	n, err := writer.Write(input)
	if err != nil {
		fmt.Printf("Error: could not write data from name pipe %s \n", err)
	}
	fmt.Printf("Written %d bytes . in memory pipe\n", n)
	time.Sleep(200*time.Millisecond)
}
