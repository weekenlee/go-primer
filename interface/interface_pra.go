package main

import "fmt"

type Animal interface {
	Speak() string
}

type Dog struct {
}

func (d Dog) Speak() string {
	return "Woof!"
}

type Cat struct {
}

func (c Cat) Speak() string {
	return "Meow!"
}

func doSomething(v interface{}) {
	switch v.(type) {
	case Dog:
		fmt.Println("Dog")
	case Cat:
		fmt.Println("Cat")
	default:
		fmt.Println("Defalut")
	}
}


func doSomethingAll(vs []interface{}) {
	for _, v := range vs {
		switch v.(type) {
		case Dog:
			fmt.Println("Dog")
		case Cat:
			fmt.Println("Cat")
		default:
			fmt.Println("Defalut")
		}
	}
}


func main() {
	animals := []Animal{
		Dog{},
		Cat{},
	}
	for _, an := range animals {
		fmt.Println(an.Speak())
	}
	for _, an := range animals {
		doSomething(an)
	}

	//error
	//doSomethingAll(animals)
	inanimals := make([] interface{}, len(animals))
	for i, v := range animals {
		inanimals[i] = v
	}
	doSomethingAll(inanimals)
}
