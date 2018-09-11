package main
import (
	"fmt"
	"sync/atomic"
)

func main() {
	var countVal atomic.Value
	countVal.Store([]int{1,3,5,7})
	anotherStore(&countVal) //传引用可以， 值不行
	fmt.Printf("The count value : %+v \n", countVal.Load())
}

func anotherStore(countVal *atomic.Value) {
	countVal.Store([]int{2,4,6,8})
}
