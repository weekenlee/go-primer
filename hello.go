package main

import (
	"fmt"
	"os"
	"pra"
	"regexp"
	"strings"
	"unicode/utf8"
)

func main() {
	var hi pra.Mytype = 10.0
	fmt.Println(hi)
	fmt.Println("hello world")
	fmt.Println(pra.Even(2))

	//打印命令行 a
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)

	//打印命令行 b
	sep, s = "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)

	//打印命令行 c
	fmt.Println(strings.Join(os.Args[1:], " "))

	var iu8 uint8 = 255
	fmt.Println(iu8, iu8+1)
	var i int8 = 127
	fmt.Println(i, i+1)

	var x uint8 = 1<<2 | 1<<3
	fmt.Println(x)
	fmt.Printf("%08b\n", x)
	a := fmt.Sprintf("%08b\n", x)
	fmt.Println(a)

	八进制 := 0666
	fmt.Printf("%d %[1]o %#[1]o\n", 八进制)

	十六 := 0xabc
	fmt.Printf("%d %[1]x %#[1]x\n", 十六)

	sw := "Hello, 世界"
	fmt.Println(len(sw))                    // "13"
	fmt.Println(utf8.RuneCountInString(sw)) // "9"

	for i := 0; i < len(sw); {
		r, size := utf8.DecodeRuneInString(sw[i:])
		fmt.Printf("%d\t%c\n", i, r)
		i += size
	}

	//Go语言的range循环在处理字符串的时候，会自动隐式解码UTF8字符串
	for i, r := range "Hello, 世界" {
		fmt.Printf("%d\t%q\t%d\n", i, r, r)
	}

	filepath := "/a/b/c.go"
	fmt.Println(basename(filepath), basename2(filepath))
	filepath2 := "/a/b/c.go.go"
	fmt.Println(basename(filepath2), basename2(filepath2))

	text := "Hello 世界！123 Go."

	// 查找连续的小写字母
	reg := regexp.MustCompile(`[a-z]+`)
	fmt.Printf("%q\n", reg.FindAllString(text, -1))
	// ["ello" "o"]

	// 查找连续的非小写字母
	reg = regexp.MustCompile(`[^a-z]+`)
	fmt.Printf("%q\n", reg.FindAllString(text, -1))
	// ["H" " 世界！123 G" "."]

	// 查找连续的单词字母
	reg = regexp.MustCompile(`[\w]+`)
	fmt.Printf("%q\n", reg.FindAllString(text, -1))
	// ["Hello" "123" "Go"]

	// 查找连续的非单词字母、非空白字符
	reg = regexp.MustCompile(`[^\w\s]+`)
	fmt.Printf("%q\n", reg.FindAllString(text, -1))
	// ["世界！" "."]

	// 查找连续的大写字母
	reg = regexp.MustCompile(`[[:upper:]]+`)
	fmt.Printf("%q\n", reg.FindAllString(text, -1))
	// ["H" "G"]

	// 查找连续的非 ASCII 字符
	reg = regexp.MustCompile(`[[:^ascii:]]+`)
	fmt.Printf("%q\n", reg.FindAllString(text, -1))
	// ["世界！"]

	// 查找连续的标点符号
	reg = regexp.MustCompile(`[\pP]+`)
	fmt.Printf("%q\n", reg.FindAllString(text, -1))
	// ["！" "."]

	// 查找连续的非标点符号字符
	reg = regexp.MustCompile(`[\PP]+`)
	fmt.Printf("%q\n", reg.FindAllString(text, -1))
	// ["Hello 世界" "123 Go"]

	// 查找连续的汉字
	reg = regexp.MustCompile(`[\p{Han}]+`)
	fmt.Printf("%q\n", reg.FindAllString(text, -1))
	// ["世界"]

	// 查找连续的非汉字字符
	reg = regexp.MustCompile(`[\P{Han}]+`)
	fmt.Printf("%q\n", reg.FindAllString(text, -1))
	// ["Hello " "！123 Go."]

	// 查找 Hello 或 Go
	reg = regexp.MustCompile(`Hello|Go`)
	fmt.Printf("%q\n", reg.FindAllString(text, -1))
	// ["Hello" "Go"]

	// 查找行首以 H 开头，以空格结尾的字符串
	reg = regexp.MustCompile(`^H.*\s`)
	fmt.Printf("%q\n", reg.FindAllString(text, -1))
	// ["Hello 世界！123 "]

	// 查找行首以 H 开头，以空白结尾的字符串（非贪婪模式）
	reg = regexp.MustCompile(`(?U)^H.*\s`)
	fmt.Printf("%q\n", reg.FindAllString(text, -1))
	// ["Hello "]

	// 查找以 hello 开头（忽略大小写），以 Go 结尾的字符串
	reg = regexp.MustCompile(`(?i:^hello).*Go`)
	fmt.Printf("%q\n", reg.FindAllString(text, -1))
	// ["Hello 世界！123 Go"]

	// 查找 Go.
	reg = regexp.MustCompile(`\QGo.\E`)
	fmt.Printf("%q\n", reg.FindAllString(text, -1))
	// ["Go."]

	// 查找从行首开始，以空格结尾的字符串（非贪婪模式）
	reg = regexp.MustCompile(`(?U)^.* `)
	fmt.Printf("%q\n", reg.FindAllString(text, -1))
	// ["Hello "]

	// 查找以空格开头，到行尾结束，中间不包含空格字符串
	reg = regexp.MustCompile(` [^ ]*$`)
	fmt.Printf("%q\n", reg.FindAllString(text, -1))
	// [" Go."]

	// 查找“单词边界”之间的字符串
	reg = regexp.MustCompile(`(?U)\b.+\b`)
	fmt.Printf("%q\n", reg.FindAllString(text, -1))
	// ["Hello" " 世界！" "123" " " "Go"]

	// 查找连续 1 次到 4 次的非空格字符，并以 o 结尾的字符串
	reg = regexp.MustCompile(`[^ ]{1,4}o`)
	fmt.Printf("%q\n", reg.FindAllString(text, -1))
	// ["Hello" "Go"]

	// 查找 Hello 或 Go
	reg = regexp.MustCompile(`(?:Hell|G)o`)
	fmt.Printf("%q\n", reg.FindAllString(text, -1))
	// ["Hello" "Go"]

	// 查找 Hello 或 Go，替换为 Hellooo、Gooo
	reg = regexp.MustCompile(`(?:Hell|G)o`)
	fmt.Printf("%q\n", reg.ReplaceAllString(text, "${n}ooo"))
	// "Hellooo 世界！123 Gooo."

	// 交换 Hello 和 Go
	reg = regexp.MustCompile(`(Hello)(.*)(Go)`)
	fmt.Printf("%q\n", reg.ReplaceAllString(text, "$3$2$1"))
	// "Go 世界！123 Hello."

	// 特殊字符的查找
	reg = regexp.MustCompile(`[\f\t\n\r\v\123\x7F\x{10FFFF}\\\^\$\.\*\+\?\{\}\(\)\[\]|]`)
	fmt.Printf("%q\n", reg.ReplaceAllString("\f\t\n\r\v\123\x7F\U0010FFFF\\^$.*+?{}()[]|", "-"))
	// "----------------------"
}

func basename(s string) string {
	for i := len(s) - 1; i > 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}

	return s
}

func basename2(s string) string {
	slash := strings.LastIndex(s, "/")
	s = s[slash+1:]
	if point := strings.LastIndex(s, "."); point > 0 {
		s = s[:point]
	}
	return s
}
