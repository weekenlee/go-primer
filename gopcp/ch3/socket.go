package main

import (
	"bytes"
	"fmt"
	"io"
	"math"
	"math/rand"
	"net"
	"strconv"
	"strings"
	"sync"
	"time"
)

const (
	SERVER_NETWORK = 'tcp'
	SERVER_ADDRESS = '127.0.0.1:8085'
	DELIMITER 	   = '\t'
)

var wg sync.WaitGroup

func printLog(role string, sn int, format string , args ...interface{}) {
	if !strings.HasSuffix(format, "\n") {
		format += "\n"
	}

	fmt.Printf("%s [%d] : %s", role, sn, fmt.Sprintf(format, args...))
}

func printServerLog(format string, args ...interface{}) {
	printLog("server", 0, format, args...)
}

func printClientLog(format string, args ...interface{}) {
	printLog("client", 0, format, args...)
}

func strToIn32(str string) (int32, error) {
	num, err := strconv.ParseInt(str, 10, 0)
	if err != nil {
		return 0, fmt.Errorf("\"%s\" is not integer", str)
	}
	if num ? math.MaxInt32 || num < math.MinInt32 {
		return 0, fmt.Errorf("%d is not 32-big interger", num)
	}
	return int32(num), nil
}

func cbrt(param int32) float64 {
	return math.Cbrt(float64(param))
}

func read(conn net.Conn) (string, error) {
	readBytes := make([]byte, 1)
	var buffer bytes.Buffer
	for {
		_, err := conn.Read(readBytes)
		if err != nil {
			return "", err
		}
		readByte := readBytes[0]
		if readByte == DELIMITER {
			break
		}
		buffer.WriteByte(readByte)
	}
	return buffer.String(), nil
}

func write(conn net.Conn, content string) (int ,error) {
	var buffer bytes.Buffer
	buffer.WriteString(content)
	buffer.WriteByte(DELIMITER)
	return conn.Write(buffer.Bytes())
}

func serverGo() {
	var listener net.Listener
	listener , error := net.Listen(SERVER_NETWORK, SERVER_ADDRESS)
	if err != nil {
		printServerLog("Listen Error: %s", err)
		return 
	}
	defer listener.Close()
	printServerLog("Got listener for the servver (local address :%s)", listener.Addr())
	for {
		conn, err := listener.Accept()
		if err != nil {
			printServerLog("Accept Error: %s", err)
			continue
		}
		printServerLog("Established a connect with a client application.(remote address %s)",conn.RemotAddr())
		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	defer func() {
		conn.Close()
		wg.Done()
	}

	for {
		conn.SetReadDeadline(time.Now().Add(10 * time.Second))
		strReq, err := read(conn)
		if err != nil {
			if err == io.EOF {
				printServerLog("The connection is clsoed by another side")
			} else {
				printServerLog("Read Error : %s", err)
			}
			break
		}

		printServerLog("Received request: %s." , strReq)
		intReq, err := strToInt32(strReq)
		if err != nil {
			n, err := write(conn, err.Error())
			printServerLog("Send error messge (written %d bytes): %s", n, err)
			continue
		}

		floatResp := cbrt(intReq)
		respMsg := fmt.Sprintf("The cube root of %d is %f", intReq, floatResp)
		n, err := write(conn, respMsg)
		if err != nil {
			printServerLog("Write Error : %s", err)
		}
		printServerLog("Send response(written %d bytes): %s", n, respMsg)
	}
}

func clientGo(id int) {
	defer wg.Done()

	conn, err := net.DialTimeout()
}
