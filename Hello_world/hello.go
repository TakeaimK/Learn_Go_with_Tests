package main

import (
	"fmt"
)

const greeting = "Hello, "

func Hello(name string) string {
	if name == "" {
		name = "world"
	}
	return greeting + name
}

func main() {
	fmt.Println(Hello(""))
	fmt.Println(Hello("Sejin"))
}
