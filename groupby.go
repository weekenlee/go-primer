package main

import (
	"fmt"
)
type id = interface{}

func groupBy(f func(id) id)  func([]id) map[id][]id {
	return func(list []id) map[id][]id {
		rmap := make(map[id][]id)
		for _, v := range list {
			rmap[f(v)] = append(rmap[f(v)] , v)
		}
		return rmap
	}
}

func f1(v id) id {
	if v=="1" {
		return "YES"
	} else {
		return "NO"
	}
}

func f2(v id) id {
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
	r1 := groupBy(f1)([]interface{}{"1","2","3"})
	fmt.Println(r1)

	r2 := groupBy(f2)([]interface{}{1,2,3,4,5,6})
	fmt.Println(r2)

}
