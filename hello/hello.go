package main

import "fmt"

const englishPrefix= "Hello, "
const spanishPrefix = "Holla, "
const frenchPrefix = "Bonjour, "

func Hello(message, language string) string {
	if message == "" {
		message = "World"
	}
	return getPrefix(language) + message
}

func getPrefix(language string) (prefix string){
	switch language {
	case "French":
		prefix = frenchPrefix
	case "Spanish":
		prefix = spanishPrefix
	default:
		prefix = englishPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("World", "French"))
}