package main

import "fmt"

type bot interface {
	getGreeting() string
}

func sayHello(b bot) {
	fmt.Println(b.getGreeting())
}

type englishBot struct {
	Greeting string
}

type spanishBot struct {
	Greeting string
}

func (b englishBot) getGreeting() string {
	return b.Greeting
}

func (b spanishBot) getGreeting() string {
	return b.Greeting
}

func main() {
	eb := englishBot{
		Greeting: "Hello there, General Kenobi!",
	}
	sb := spanishBot{
		Greeting: "Hola como estas mi amigo!",
	}

	sayHello(eb)
	sayHello(sb)
}
