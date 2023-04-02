package main

import "fmt"

type bot interface {
	getGreeting() string // in order to qualify as type bot, you have to have a function call getGreeting and return type string
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (eb englishBot) getGreeting() string {
	// very custom logic for generating and English greeting
	return "hello"
}

func (sb spanishBot) getGreeting() string {
	// very custom logic for generating and Spanish greeting
	return "hola"
}
