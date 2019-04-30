package main

/*
#include <stdio.h>

void hi() {
	printf("hello world from c!\n");
}
*/
import "C"
import "fmt"

func main() {
	C.hi()
	fmt.Println("hi from go")
}
