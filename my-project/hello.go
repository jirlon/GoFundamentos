package main

import "fmt"

const spanish = "espanhol"
const french = "frances"
const japonese = "japones"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "
const japoneseHelloPrefix = "Kon'nichiwa, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	prefix := englishHelloPrefix

	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	case japonese:
		prefix = japoneseHelloPrefix
	}
	/*
		if language == spanish {
			return spanishHelloPrefix + name
		}

		if language == "frances" {
			return frenchHelloPrefix + name
		}

		return englishHelloPrefix + name
	*/
	return prefix + name
}

func main() {
	fmt.Println(Hello("world", ""))
}
