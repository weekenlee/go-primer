//go build -o hello.so -buildmode = c-shared .
package main

/*
extern int helloFromC();
*/
import "C"
import "fmt"

//export HelloFromGo
func HelloFromGo() {
	fmt.Printf("hello from go \n")
	C.helloFromC()
}

func main() {

}
