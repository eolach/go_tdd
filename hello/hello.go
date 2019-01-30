package main

import "fmt"

const spanish = "spanish"
const french = "french"
const helloPrefix = "Hello, "
const helloPrefixSpanish = "Hola, "
const helloPrefixFrench = "Bonjour, "

func Hello(name string, language string) string {
	if name == "" {
		name = "world"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case "spanish":
		prefix = helloPrefixSpanish
	case "french":
		prefix = helloPrefixFrench
	default:
		prefix = helloPrefix
	}

	return
}

func main() {
	fmt.Println(Hello("world", ""))
}
