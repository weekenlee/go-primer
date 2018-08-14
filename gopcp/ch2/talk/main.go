package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"runtime/debug"

	"liweijian/chatbot"
)

var chatbotName string

func init() {
	flag.StringVar(&chatbotName, "chatbot", "simple.en", "the chatbot's name for dialogue.'")
}
