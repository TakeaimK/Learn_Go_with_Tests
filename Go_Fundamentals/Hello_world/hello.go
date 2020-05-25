package main

import (
	"fmt"
)

const EngGreeting = "Hello, "
const FrenchGreeting = "Bonjour, "
const SpanishGreeting = "Hola, "

const Spanish = "Spain"
const French = "France"

//public
func Hello(name string, lang string) string {
	if name == "" {
		name = "world"
	}

	return greet(lang) + name
}

//private
func greet(lang string) (prefix string) {
	switch lang {
	case Spanish:
		prefix = SpanishGreeting
	case French:
		prefix = FrenchGreeting
	default:
		prefix = EngGreeting
	}
	return
}

func main() {
	fmt.Println(Hello("", ""))
	fmt.Println(Hello("Sejin", ""))
	fmt.Println(Hello("Elodie", "Spain"))
	fmt.Println(Hello("Macron", "France"))

}
