package main

import (
	"fmt"
	"unsafe"
)

type notifier interface {
	notify()
}

type user struct {
	name string
}

func (u user) notify() {
	fmt.Println("Alert", u.name)
}

func inspect(n *notifier, u *user) {
	word := uintptr(unsafe.Pointer(n)) + uintptr(unsafe.Sizeof(&u))
	value := (**user)(unsafe.Pointer(word))
	fmt.Printf("Addr User: %p word value :%p ptrvalue: %v\n", u, *value, **value)
}

func main() {
	var n1 notifier
	u := user{"bill"}

	n1 = u
	inspect(&n1, &u)

	n2 := n1
	inspect(&n2, &u)

	n1 = &u
	inspect(&n1, &u)
}
