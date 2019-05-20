package main

import "fmt"

//GetLetterByNum 根据数字获取字符
func GetLetterByNum(num int) (s string) {
	if num <= 0 {
		panic("num参数取值范围为大于零的整数")
	}

	var temp []rune

	yu := 0
	shang := num
	for {
		yu = shang % 26
		shang = shang / 26

		if yu == 0 {
			yu = 26
			shang--
		}

		temp = append(temp, rune(yu+'A'-1))

		if shang == 0 {
			break
		}
	}

	for i := len(temp) - 1; i >= 0; i-- {
		s += string(temp[i])
	}

	return
}

func main() {
	fmt.Println(GetLetterByNum(1))
	fmt.Println(GetLetterByNum(28))
}
