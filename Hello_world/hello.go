package main

import (
	"fmt"
)

const greeting = "Hello, "

func Hello(name string) string {
	if name == "" {
		return greeting + "world"
	}
	return greeting + name
}

func main() {
	fmt.Println(Hello(""))
	fmt.Println(Hello("Sejin"))
}
