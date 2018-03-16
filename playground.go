package main

import (
	"fmt"
)

func groupBy(f func(string) string)  func([]string) map[string][]string {
	return func(list []string) map[string][]string {
		rmap := make(map[string][]string)
		for _, v := range list {
			rmap[f(v)] = append(rmap[f(v)] , v)
		}
		return rmap
	}
}

func groupBy2(f func(interface{}) interface{})  func([]interface{}) map[interface{}][]interface{} {
	return func(list []interface{}) map[interface{}][]interface{} {
		rmap := make(map[interface{}][]interface{})
		for _, v := range list {
			rmap[f(v)] = append(rmap[f(v)] , v)
		}
		return rmap
	}
}

func f1(v interface{}) interface{} {
	if v=="1" {
		return "YES"
	} else {
		return "NO"
	}
}

func f2(v string) string {
	if v=="1" {
		return "你好"
	} else {
		return "我好"
	}
}

func f3(v interface{}) interface{} {
	switch v := v.(type) {
	case int:
		if int(v)%2 == 0 {
			return "偶数"
		} else {
			return "奇数"
		}
	}
	return "NULL"
}

func main() {

	r := groupBy(func(v string) string {
		if v=="1" {
			return "YES"
		} else {
			return "NO"
		}
	})([]string{"1","2","3"})
	fmt.Println(r)

	r1 := groupBy(f2)([]string{"1","2","3"})
	fmt.Println(r1)



	r3 := groupBy2(f1)([]interface{}{"1","2","3"})
	fmt.Println(r3)


	r4 := groupBy2(f3)([]interface{}{1,2,3,4,5,6})
	fmt.Println(r4)
}
