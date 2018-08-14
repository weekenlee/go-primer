package chatbot

import (
	"fmt"
	"strings"
)

type simpleCN struct {
	name string
	talk Talk
}

func NewSimpleCN(name string, talk Talk) Chatbot {
	return &simpleCN {
		name: name,
		talk: talk,
	}
}

func (robot *simpleCN) Name() string {
	return robot.name
}

func (robot *simpleCN) Begin() (string, error) {
	return "请输入您的名字: ", nil
}

func (robot *simpleCN) Hello(userName string) string {
	userName = strings.TrimSpace(userName)
	if robot.talk != nil {
		return robot.talk.Hello(userName)
	}
	return fmt.Sprintf("你好， %s ！ 我可以为你做些什么?", userName)
}

func (robot *simpleCN) Talk(heard string) (saying string, end bool , err error) {
	heard = strings.TrimSpace(heard)
	if robot.talk != nil {
		return robot.talk.Talk(heard)
	}

	switch heard {
	case "":
		return 
	case "没有", "再见":
		saying = "再见"
		end = true
		return
	default:
		saying = "对不起， 我没有听懂您说的。"
		return 
	}
}

func (robot *simpleCN) ReportError(err error) string {
	return fmt.Sprintf("发生了一个错误: %s\n", err)
}

func (robot *simpleCN) End() error {
	return nil
}
