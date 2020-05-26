package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func ListDir(dirPth string, datestr string) {
	path := dirPth

	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		//fmt.Println(f.Name(), "开始修改")
		if strings.Contains(f.Name(), ".exe") {
			continue
		}
		err = os.Rename(path+"//"+f.Name(), path+"//"+
			datestr+f.Name())
		if err != nil {
			fmt.Println(err)
		}
	}
}

func main() {
	f := bufio.NewReader(os.Stdin) //读取输入的内容
	fmt.Print("请输入日期(如20200526)>")
	var Input string
	Input, _ = f.ReadString('\n') //定义一行输入的内容分隔符。
	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	ListDir(dir, Input[:len(Input)-1])
}
