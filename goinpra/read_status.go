package main

import (
    "bufio"
    "net"
    "fmt"
)

func main() {
    conn, _ := net.Dial("tcp", "csdn.net:80")
    fmt.Fprintf(conn, "Get / HTTP/1.0\r\n\r\n")
    status, _ := bufio.NewReader(conn).ReadString('\n')
    fmt.Println(status)
}
