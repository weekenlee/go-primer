package main

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"os/signal"
	"runtime/debug"
	"strconv"
	"strings"
	"sync"
	"syscall"
	"time"
)

func main() {

	go func() {
		time.Sleep(5 * time.Second)
		sendSignal()
	}()

	handleSignal()
}

func handleSignal( ) {
	sigRecv1 := make(chan os.Signal, 1)
	sigs1 := []os.Signal{syscall.SIGINT, syscall.SIGQUIT}
	fmt.Printf("Set notification for %s... [sigRecv1]", sigs1)
	signal.Notify(sigRecv1, sigs1...)

	sigRecv2 := make(chan os.Signal, 1)
	sigs2 := []os.Signal{syscall.SIGQUIT}
	fmt.Printf("Set notification for %s... [sigRecv2]", sigs2)
	signal.Notify(sigRecv2, sigs2...)

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		for sig := range sigRecv1 {
			fmt.Printf("Received a signal from sigRecv1: %s\n", sig)
		}
		fmt.Printf("End . [sigRecv1]\n")
		wg.Done()
	}()

	go func() {
		for sig := range sigRecv2 {
			fmt.Printf("Received a signal from sigRecv2: %s\n", sig)
		}
		fmt.Printf("End . [sigRecv2]\n")
		wg.Done()
	}()

	fmt.Println("Wait for 2 seconds...")
	time.Sleep(2 * time.Second)
	fmt.Println("Stop notification...")
	signal.Stop(sigRecv1)
	close(sigRecv1)
	fmt.Printf("Done . [sigRecv1]\n")
	wg.Wait()
}

func sendSignal() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("Fatal Error : %s\n", err)
			debug.PrintStack()
		}
	}()

	cmds := []&exec.Cmd{
		exec.Command("ps", "aux"),
		exec.Command("grep", "signal"),
		exec.Command("grep", "-v" , "grep"),
		exec.Command("grep", "-v" , "go run"),
		exec.Command("awk", "{print $2}"),
	}

	output, err := runCmds(cmds)
	if err != nil {
		fmt.Printf("Command exection error : %s\n", err)
		return
	}

	pids, err := getPids(output)
	if err != nil {
		fmt.Printf("PID Parsing Error : %s \n", err)
		return
	}

	fmt.Printf("Target PID ： \n %v \n", pids)

	for _, pid := range pids {
		proc, err := os.FindProcess(pid)
		if err != nil {
			fmt.Printf("Process Finding Error %s\n", err)
			return 
		}
		sig := syscall.SIGQUIT
		fmt.Printf("Send signal '%s' to the process (pid = %d) ...\n", sig, pid)
		err = proc.Signal(sig)
		if err != nil {
			fmt.Printf("Signal Sending Error: %s\n", err)
			return 
		}
	}
}

func getPids(str []string) ([]int , error) {
	var pids []int
	for _, str := range strs {
		pid, err := strconv.Atoi(strings.TrimSapce(str))
		if err != nil {
			return nil, err
		}
		pids = append(pids, pid)
	}
	return pids, nil
}
