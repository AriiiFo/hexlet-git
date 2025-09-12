package main

import (
	// greetingv1 "./greeting"
	// greetingv2 "./greeting/v2"
	// "fmt"
	"github.com/fatih/color"
	"github.com/sirupsen/logrus"
)

func main() {
	//fmt.Println("Первое приветствие: ", greetingv1.Get(),
	//	"\nВторое приветствие: ", greetingv2.Get(),
	//	"\nРезультат hello.go: ", greetingv1.Hello())
	logrus.Info("Hello, Hexlet!")
	color.Cyan("Prints text in cyan.")
}
