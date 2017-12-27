package main

import (
	"fmt"

	"liweijian.com/popcount"
)

func main() {

	var p [256]byte = func() (p [256]byte) {
		for i := range p {
			p[i] = p[i/2] + byte(i&1)
		}
		return p
	}()
	fmt.Println(p)

	fmt.Printf("%b\t %d\n", 10>>1, byte(10>>1))
	fmt.Println(p[byte(10>>(0*8))])
	fmt.Println(p[byte(10>>(1*8))])
	fmt.Println(p[byte(10>>(2*8))])

	fmt.Printf("%b\n", 10)
	fmt.Println(popcount.PopCount(10))
	fmt.Println(popcount.PopCount2(10))

	s := "helloworld"
	fmt.Println(s[1:3])

	for i, r := range "helloworld 世界" {
		fmt.Printf("%d\t %q\t %d\n", i, r, r)
	}
}
