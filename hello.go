package main

import "fmt"

const englishHelloPrefix = "Hello, "

func Hello(name string) string {
	return englishHelloPrefix + name
	// return "Hello, " + name
	// return "Hello, world" //got "Hello, world" want "Hello, Chris"
}

func main() {
	fmt.Println(Hello("world"))
}
