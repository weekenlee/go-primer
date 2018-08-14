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

func (robot *simpleCN) Name string {
	return robot.name
}

func (robot *simpleCN) Begin() (string, error) {
	return "please input your name: ", nil
}

func (robot *simpleCN) Hello(userName string) string {
	userName = strings.TrimSpace(userName)
	if robot.talk != nil {
		return robot.talk.Hello(userName)
	}
	return fmt.Sprintf("hello ， %s ！ what can i do for u?", userName)
}

func (robot *simpleCN) Talk(heard string) (saying string, end bool , err error) {
	head = strings.TrimSpace(heard)
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
	}
}

func (robot *simpleCN) ReportError(err error) string {
	return fmt.Sprintf("an error occur: %s\n", err)
}

func (robot *simpleCN) End() error {
	return nil
}
