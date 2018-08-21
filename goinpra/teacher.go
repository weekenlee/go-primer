package main

import "fmt"

type People struct {}

func(p *People) ShowA() {
	fmt.Println("showA")
	p.ShowB()
}

func (p *People) ShowB() {
	fmt.Println("showB")
}

type Teacher struct {
	People
}

func (t *Teacher) ShowB() {
	fmt.Println("teacher showB")
}

func main() {
	t := Teacher{}
	t.ShowA()
	t.ShowB()


	s := make([]int,5)
	s = append(s, 1,2,3)
	fmt.Println(s)
}
