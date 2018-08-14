package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os/exec"
)

func main() {
	runCmd()
	fmt.Println()
	runCmdWithPipe()
}

func runCmdWithPipe() {
	fmt.Println("Run command 'ps aux|grep apipe':")
	cmd1 := exec.Command("ps", "aux")
	cmd2 := exec.Command("grep", "apipe")
	var outputBuf1 bytes.Buffer
	cmd1.Stdout = &outputBuf1
	if err := cmd1.Start(); err != nil {
		fmt.Printf("Error: Couldnot wait for the first command: %s\n", err)
		return 
	}
	if err := cmd1.Wait(); err != nil {
		fmt.Printf("Error: Could not wait for the first command: %s\n", err)
		return 
	}
	cmd2.Stdin = &outputBuf1
	var outputBuf2 bytes.Buffer
	cmd2.Stdout = &outputBuf2
	if err := cmd2.Start(); err != nil {
		fmt.Printf("Error: Couldnot wait for the first command: %s\n", err)
		return 
	}
	if err := cmd2.Wait(); err != nil {
		fmt.Printf("Error: Could not wait for the first command: %s\n", err)
		return 
	}
	fmt.Printf("%s\n", outputBuf2.Bytes())
}

func runCmd() {
	useBufferedIO := false
	fmt.Println("Run command 'echo -n \"my first command come from golang.\"':")
	cmd0 := exec.Command("echo", "-n", "my first comand come from golang.")
	stdout0, err := cmd0.StdoutPipe()
	if err != nil {
		fmt.Printf("Error: Could not obtain the stdout pipe for command No.0 : %s\n", err)
		return 
	}
	if err := cmd0.Start(); err != nil {
		fmt.Printf("Error: the command No.0 can not be startup :%s\n", err)
		return
	}

	if !useBufferedIO {
		var outputBuf0 bytes.Buffer 
		for {
			tempOutput := make([]byte, 5)
			n, err := stdout0.Read(tempOutput)
			if err != nil {
				if err == io.EOF {
					break
				} else {
					fmt.Printf("Error: could not read data from the pip： %s\n", err)
					return
				}
			}
			if n > 0 {
				outputBuf0.Write(tempOutput[:n])
			}
		}
		fmt.Printf("%s\n", outputBuf0.String())
	}else {
		outputBuf0 := bufio.NewReader(stdout0)
		output0, _, err := outputBuf0.ReadLine()
		if err != nil {
			fmt.Printf("Error : Couldnot read data from the pipe: %s\n", err)
			return 
		}
		fmt.Printf("%s\n", string(output0))
	}
}
