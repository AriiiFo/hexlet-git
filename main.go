package main

import (
  greetingv1 "./greeting"
  greetingv2 "./greeting/v2"
  "fmt"
)

func main() {
  fmt.Println("Первое приветствие: ", greetingv1.Get(),
	"\nВторое приветствие: ", greetingv2.Get(), 
	"\nРезультат hello.go:", greetingv1.Hello())
}
