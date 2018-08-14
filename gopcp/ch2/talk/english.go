package chatbot

import (
	"fmt"
	"strings"
)

type simpleEN struct {
	name string
	talk Talk
}

func NewSimpleEN(name string, talk Talk) Chatbot {
	return &simpleEN {
		name: name,
		talk: talk,
	}
}

func (robot *simpleEN) Name() string {
	return robot.name
}

func (robot *simpleEN) Begin() (string, error) {
	return "please input your name: ", nil
}

func (robot *simpleEN) Hello(userName string) string {
	userName = strings.TrimSpace(userName)
	if robot.talk != nil {
		return robot.talk.Hello(userName)
	}
	return fmt.Sprintf("hello ， %s ！ what can i do for u?", userName)
}

func (robot *simpleEN) Talk(heard string) (saying string, end bool , err error) {
	heard = strings.TrimSpace(heard)
	if robot.talk != nil {
		return robot.talk.Talk(heard)
	}

	switch heard {
	case "":
		return 
	case "nothing", "bye":
		saying = "bye"
		end = true
		return
	default:
		saying = "sorry， i did not catch you。"
		return
	}
}

func (robot *simpleEN) ReportError(err error) string {
	return fmt.Sprintf("an error occur: %s\n", err)
}

func (robot *simpleEN) End() error {
	return nil
}
