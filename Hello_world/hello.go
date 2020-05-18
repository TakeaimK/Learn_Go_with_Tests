package main

import (
	"fmt"
)

const EngGreeting = "Hello, "
const FrenchGreeting = "Bonjour, "
const SpanishGreeting = "Hola, "

const Spanish = "Spain"
const French = "France"

func Hello(name string, lang string) string {
	if name == "" {
		name = "world"
	}

	prefix := EngGreeting
	switch lang {
	case Spanish:
		prefix = SpanishGreeting
	case French:
		prefix = FrenchGreeting
	default:
		prefix = EngGreeting
	}

	return prefix + name
}

func main() {
	fmt.Println(Hello("", ""))
	fmt.Println(Hello("Sejin", ""))
	fmt.Println(Hello("Elodie", "Spain"))
	fmt.Println(Hello("Macron", "France"))

}
