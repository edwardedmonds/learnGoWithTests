package main

import "fmt"

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "
const portugueseHelloPrefix = "Olá, "
const polishHelloPrefix = "Cześć, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	prefix := englishHelloPrefix

	switch language {
	case "French":
		prefix = frenchHelloPrefix
	case "Spanish":
		prefix = spanishHelloPrefix
	case "Portuguese":
		prefix = portugueseHelloPrefix
	case "Polish":
		prefix = polishHelloPrefix
	}

	return prefix + name
}

func main() {
	fmt.Println(Hello("Edward", "Polish"))
}
