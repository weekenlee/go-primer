package main

import "strings"
import "math"
import "fmt"

func myAtoi(str string) int	{
	s := strings.TrimSpace(str)
	if len(s) == 0 {
		return 0
	}

	sign, x := getSign(s)

	x = trim(x)

	return convert(sign, x)
}

func getSign(s string) (int, string) {
	sign := 1
	switch s[0] {
	case '-':
		s = s[1:]
		sign = -1.0
	case '+':
		s = s[1:]
	default:
	}
	return sign , s
}

func trim(s string) string {
	for i := range s {
		if s[i] < '0' || '9' < s[i] {
			return s[:i]
		}
	}

	return s
}

func convert(sign int, s string) int {
	base := 1 * sign
	res := 0
	yes := false

	for i := len(s) - 1; i >= 0 ; i-- {
		res, yes = isOverflow(res + (int(s[i]) - 48)*base)
		if yes {
			return res
		}
		base *= 10
	}

	return res
}

func isOverflow(i int) (int , bool) {
	switch {
	case i > math.MaxInt32:
		return math.MaxInt32, true
	case i < math.MinInt32:
		return math.MinInt32, true
	default:
		return i, false
	}
}

func main() {
	fmt.Println(myAtoi("  -101"))
}
