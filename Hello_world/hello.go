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
	if lang == Spanish {
		return SpanishGreeting + name
	}
	if lang == French {
		return FrenchGreeting + name
	}
	return EngGreeting + name
}

func main() {
	fmt.Println(Hello("", ""))
	fmt.Println(Hello("Sejin", ""))
	fmt.Println(Hello("Elodie", "Spain"))
	fmt.Println(Hello("Macron", "France"))

}
