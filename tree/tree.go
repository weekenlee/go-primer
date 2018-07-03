package main

import (
    "fmt"
	"container/list"
)

type tree interface {
	do()
}

type node struct {
	*list.List
	name string
}

type leaf struct {
	name string
}

func (n node) do() {
	for e:=n.Front(); e != nil; e=e.Next() {
		e.Value.(tree).do()
	}
	fmt.Println(n.name + " node do something.")
}

func (n node)addSub(sub tree) {
	n.PushBack(sub)
}

func (l leaf)do() {
	fmt.Println(l.name + " leaf do something.")
}


func main() {
	n1 := node{list.New(), "n1"}
	n2 := node{list.New(), "n2"}
	l1 :=leaf{"l1"}
	l2 :=leaf{"l2"}
	n2.addSub(l2)
	n1.addSub(n2)
	n1.addSub(l1)

	n1.do()
}
