package main

import "fmt"

const bolingF = 212.0

func main() {
	var f1, f2 = bolingF, 32.0
	var c1, c2 = FtoC(f1), FtoC(f2)
	fmt.Printf("boling point =  %g °F or %g °C\n", f1, c1)
	fmt.Printf("freezing point =  %g °F or %g °C\n", f2, c2)
}

func FtoC(f float64) float64 {
	return (f - 32) * 5 / 9
}
